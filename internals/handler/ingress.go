package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllIngress(c echo.Context) error {
	ings, err := k8s.GetAllIngresses(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ings)
}
