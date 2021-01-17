package k8s

import (
	"context"
	"k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllDeploy(ctx context.Context, clientset *kubernetes.Clientset) ([]v1beta1.Deployment, error) {
	deploys, err := clientset.AppsV1beta1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deploys.Items, nil
}