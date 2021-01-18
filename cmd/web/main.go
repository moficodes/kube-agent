/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/moficodes/kube-get-agent/internals/handler"
	"github.com/moficodes/kube-get-agent/pkg/k8s"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//fmt.Println(*kubeconfig)
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	clientset, err := k8s.K8sClient(kubeconfig)
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(clientset)

	e := echo.New()
	e.Use(
		middleware.CORS(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
	)

	e.Use(middleware.Secure())
	api := e.Group("/api/v1")
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	// http://localhost:8111/api/v1/pods
	// http://localhost:8111/api/v1/pods?podprefix=calico
	api.GET("/pods", h.GetAllPods)
	// http://localhost:8111/api/v1/pods/kube-system
	// http://localhost:8111/api/v1/pods/kube-system?podname=calico-node
	// http://localhost:8111/api/v1/pods/kube-system?label=k8s-app=calico-node
	api.GET("/pods/:namespace", h.GetAllPodsInNamespace)

	api.GET("/nodes", h.GetAllNodes)

	api.GET("/namespaces", h.GetAllNamespaces)

	api.GET("/ingresses", h.GetAllIngress)

	api.GET("/deployments", h.GetAllDeployment)
	api.GET("/deployments/:namespace", h.GetAllDeploymentInNamespace)

	e.Logger.Fatal(e.Start(":8111"))
}
