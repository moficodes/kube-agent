package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllPods(c echo.Context) error{
	pods, err := k8s.GetAllPods(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pods)
}