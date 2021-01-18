package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllNodes(c echo.Context) error{
	nodes, err := k8s.GetAllNodes(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nodes)
}
