package bootstrap

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var client *kubernetes.Clientset

func SetupClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Error(err.Error())
		panic(err.Error())
	}

	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error(err.Error())
	}

	return client
}

func GetClient() *kubernetes.Clientset {
	return client
}
