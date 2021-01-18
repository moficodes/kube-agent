package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"net/http"
)

func (h *Handler) GetAllNamespaces(c echo.Context) error {
	ns, err := k8s.GetAllNamespaces(c.Request().Context(), h.clientset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ns)
}
