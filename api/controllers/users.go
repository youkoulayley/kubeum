package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/youkoulayley/kubeum/api/bootstrap"
	"github.com/youkoulayley/kubeum/api/models"
	"github.com/youkoulayley/kubeum/api/tmpl"
	"html/template"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strings"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	c := bootstrap.GetClient()

	sa, err := c.CoreV1().ServiceAccounts("").List(metav1.ListOptions{})
	if err != nil {
		logrus.Error(err.Error())
	}

	users := models.Users{}

	for _, sa := range sa.Items {
		user := models.User{
			Name:      sa.Name,
			Namespace: sa.Namespace,
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func GetKubeconfig(w http.ResponseWriter, r *http.Request) {
	c := bootstrap.GetClient()

	cfg := models.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	clusterCA, err := ioutil.ReadFile(cfg.CaFile)
	if err != nil {
		logrus.Error("CA File not found : " + err.Error())
	}
	base64ClusterCA := base64.StdEncoding.EncodeToString(clusterCA)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			logrus.Error(err.Error())
		}
	} else {
		user = user.Validate()

		if len(user.Errors) <= 0 {
			if user.Exists() {
				getUser, err := c.CoreV1().ServiceAccounts(user.Namespace).Get(user.Name, metav1.GetOptions{})
				if err != nil {
					logrus.Error(err)
				}

				for _, secret := range getUser.Secrets {
					if !strings.Contains(secret.Name, "token") {
						continue
					}
					userSecret, err := c.CoreV1().Secrets(user.Namespace).Get(secret.Name, metav1.GetOptions{})
					if err != nil {
						logrus.Debug(err)
					}
					token := string(userSecret.Data["token"])
					data := tmpl.KubeconfigData{
						ClusterName: cfg.ClusterName,
						ClusterCA:   base64ClusterCA,
						ClusterURL:  cfg.APIServerURL,
						Username:    user.Name,
						Namespace:   user.Namespace,
						Token:       token,
					}
					tpl, err := template.New("").Parse(tmpl.Kubeconfig)
					if err != nil {
						logrus.Error("Unable to parse template kubeconfig : " + err.Error())
					}

					w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
					w.WriteHeader(http.StatusOK)
					tpl.Execute(w, data)
				}
			} else {
				user.Errors["exists"] = "The user does not exists"
				w.Header().Set("Content-Type", "application/json;charset=UTF-8")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(user)
			}
		}
	}
}
