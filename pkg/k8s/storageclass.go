package k8s

import (
	"context"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllStorageClasses(ctx context.Context, clientset *kubernetes.Clientset) ([]v1.StorageClass, error) {
	scs, err := clientset.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return scs.Items, nil
}