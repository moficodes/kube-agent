package k8s

import (
	"context"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllIngresses( ctx context.Context, clientset *kubernetes.Clientset) ([]v1beta1.Ingress, error) {
	ings, err := clientset.NetworkingV1beta1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return ings.Items, nil
}
