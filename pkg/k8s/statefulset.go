package k8s

import (
	"context"
	"k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllStatefulSets(ctx context.Context, clientset *kubernetes.Clientset) ([]v1beta1.StatefulSet, error) {
	sts, err := clientset.AppsV1beta1().StatefulSets("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return sts.Items, nil
}