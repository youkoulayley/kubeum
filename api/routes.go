package main

import (
	"github.com/youkoulayley/kubeum/api/controllers"
	"github.com/youkoulayley/kubeum/api/models"
)

var routes = models.Routes{
	models.Route{
		"health",
		"GET",
		"/health",
		controllers.GetHealth,
	},
	models.Route{
		"users.index",
		"GET",
		"/users",
		controllers.GetUsers,
	},
	models.Route{
		"users.kubeconfig",
		"POST",
		"/users/kubeconfig",
		controllers.GetKubeconfig,
	},
}
