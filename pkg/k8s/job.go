package k8s

import (
	"context"
	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAlljobs(ctx context.Context, clientset *kubernetes.Clientset) ([]v1.Job, error) {
	jobs, err := clientset.BatchV1().Jobs("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return jobs.Items, nil
}