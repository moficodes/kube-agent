package k8s

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllDeploy(ctx context.Context, clientset *kubernetes.Clientset) ([]v1.Deployment, error) {
	return GetAllDeployInNamespace(ctx, clientset, "")
}

func GetAllDeployInNamespace(ctx context.Context, clientset *kubernetes.Clientset, namespace string) ([]v1.Deployment, error) {
	deploys, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deploys.Items, nil
}