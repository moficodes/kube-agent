package k8s

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func GetAllPods(ctx context.Context, clientset *kubernetes.Clientset) ([]v1.Pod, error) {
	return GetAllPodsInNamespace(ctx, clientset, "")
}

func GetAllPodsInNamespace(ctx context.Context, clientset *kubernetes.Clientset, namespace string) ([]v1.Pod, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func GetPodInNamespace(ctx context.Context, clientset *kubernetes.Clientset, namespace string, podName string) (*v1.Pod, error) {
	pod, err := clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func GetPodWithPrefix(ctx context.Context, clientset *kubernetes.Clientset, podPrefix string) ([]v1.Pod, error) {
	pods, err := GetAllPods(ctx, clientset)
	if err != nil {
		return nil, err
	}

	res := make([]v1.Pod, 0)
	for _, pod := range pods {
		if strings.HasPrefix(pod.Name, podPrefix) {
			res = append(res, pod)
		}
	}

	if len(res) == 0 {
		return nil, errors.New("no such pod found")
	}
	return res, nil
}

func GetPodsWithLabel(ctx context.Context, clientset *kubernetes.Clientset, namespace string, label string) ([]v1.Pod, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

