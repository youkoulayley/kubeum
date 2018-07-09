package models

type Config struct {
	APIServerURL string `env:"API_SERVER_URL" envDefault:"https://127.0.0.1:6443"`
	CaFile       string `env:"CA_FILE" envDefault:"/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"`
	ClusterName  string `env:"CLUSTER_NAME" envDefault:"kubernetes"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"Info"`
	Port         int    `env:"PORT" envDefault:"8080"`
}
