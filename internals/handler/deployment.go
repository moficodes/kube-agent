package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllDeployment(c echo.Context) error {
	deployments, err := k8s.GetAllDeploy(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, deployments)
}

func (h *Handler) GetAllDeploymentInNamespace(c echo.Context) error {
	namespace := c.Param("namespace")
	deployments, err := k8s.GetAllDeployInNamespace(c.Request().Context(), h.clientset, namespace)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, deployments)
}