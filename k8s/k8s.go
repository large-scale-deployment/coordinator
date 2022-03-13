package k8s

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

type K8sInfoCollector struct {
	clientset *kubernetes.Clientset
}

type PodConditionStats struct {
	Type               v1.PodConditionType
	LastTransitionTime metav1.Time
}

type PodConditionStatsList struct {
	Items []PodConditionStats
}

func FindKubeConfigFile() string {
	if home := homedir.HomeDir(); home != "" {
		return filepath.Join(home, ".kube", "config")
	} else {
		return ""
	}
}

type KubernetesClientWrapper struct {
	Clientset *kubernetes.Clientset
}

func NewKubernetesClientWrapper(kubeconfig string) (*KubernetesClientWrapper, error) {
	if kubeconfig == "" {
		// use the current context in kubeconfig
		kubeconfig = FindKubeConfigFile()
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &KubernetesClientWrapper{Clientset: clientset}, nil
}

func (clientProxy *KubernetesClientWrapper) CreateDemoDeployment(deploymentName string) *schema.GroupVersionResource {
	configFile := FindKubeConfigFile()
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		return nil
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploymentRes := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": deploymentName,
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name":          "http",
										"protocol":      "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	namespace := "default"
	result, err := client.Resource(deploymentRes).Namespace(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetName())

	return &deploymentRes

}
func (clientProxy *KubernetesClientWrapper) DeleteDemoDeployment(deploymentName string, deploymentRes *schema.GroupVersionResource) {
	configFile := FindKubeConfigFile()
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		return
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleting deployment...")
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	namespace := "default"
	if err := client.Resource(*deploymentRes).Namespace(namespace).Delete(context.TODO(), deploymentName, deleteOptions); err != nil {
		panic(err)
	}

	fmt.Println("Deleted deployment.")
}

func NewK8sInfoCollector(kubeconfig string) (*K8sInfoCollector, error) {
	// create the clientset
	clientProxy, err := NewKubernetesClientWrapper(kubeconfig)
	if err != nil {
		return nil, err
	}
	return &K8sInfoCollector{clientset: clientProxy.Clientset}, nil
}

func (collector *K8sInfoCollector) GetPods(namePattern string) (selectedPods []v1.Pod, err error) {

	nameReg := regexp.MustCompile(namePattern)
	selectedPods = make([]v1.Pod, 0)

	pods, err := collector.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		if nameReg.MatchString(pod.Name) {
			selectedPods = append(selectedPods, pod)
			/*
				for _, condition := range pod.Status.Conditions {
					stats := &PodConditionStats{
						Type: condition.Type, LastTransitionTime: condition.LastTransitionTime,
					}
					fmt.Println(condition.Type, condition.LastTransitionTime)
				}
			*/
		}
	}
	return selectedPods, nil
}
