package handler

import "k8s.io/client-go/kubernetes"

type Handler struct {
	clientset *kubernetes.Clientset
}

func NewHandler(clientset *kubernetes.Clientset)(*Handler) {
	h := Handler{
		clientset: clientset,
	}
	return &h
}
