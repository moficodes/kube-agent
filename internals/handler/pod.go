package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllPods(c echo.Context) error{
	podprefix := c.QueryParam("podprefix")
	if podprefix != "" {
		pods, err := k8s.GetPodWithPrefix(c.Request().Context(), h.clientset, podprefix)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("pod with prefix %s not found", podprefix))
		}
		return c.JSON(http.StatusOK, pods)
	}
	pods, err := k8s.GetAllPods(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pods)
}

func (h *Handler) GetAllPodsInNamespace(c echo.Context) error {
	namespace := c.Param("namespace")
	pod := c.QueryParam("podname")
	if pod != "" {
		pod, err := k8s.GetPodInNamespace(c.Request().Context(), h.clientset, namespace, pod)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("pod %s not found", pod))
		}
		return c.JSON(http.StatusOK, pod)
	}

	label := c.QueryParam("label")
	if label != "" {
		pods, err := k8s.GetPodsWithLabel(c.Request().Context(), h.clientset, namespace, label)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("pod with label %s not found", label))
		}
		return c.JSON(http.StatusOK, pods)
	}

	pods, err := k8s.GetAllPodsInNamespace(c.Request().Context(), h.clientset, namespace)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pods)
}
