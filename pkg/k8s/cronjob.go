package k8s

import (
	"context"
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllCronjobs(ctx context.Context, clientset *kubernetes.Clientset) ([]v1beta1.CronJob, error) {
	cronjobs, err := clientset.BatchV1beta1().CronJobs("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return cronjobs.Items, nil
}