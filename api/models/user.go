package models

import (
	"github.com/youkoulayley/kubeum/api/bootstrap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type User struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Errors    map[string]string `json:"errors"`
}

type Users []User

func (user User) Validate() User {
	user.Errors = make(map[string]string)

	if strings.TrimSpace(user.Name) == "" {
		user.Errors["name"] = "Please enter a valid email address"
	}
	if strings.TrimSpace(user.Namespace) == "" {
		user.Errors["namespace"] = "Please enter a valid namespace"
	}

	return user
}

func (user User) Exists() bool {
	c := bootstrap.GetClient()

	_, err := c.CoreV1().ServiceAccounts(user.Namespace).Get(user.Name, metav1.GetOptions{})
	if err != nil {
		return false
	}

	return true
}
