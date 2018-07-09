package tmpl

type KubeconfigData struct {
	ClusterName string
	ClusterCA   string
	ClusterURL  string
	Username    string
	Namespace   string
	Token       string
}

var Kubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{.ClusterCA}}
    server: {{.ClusterURL}}
  name: {{.ClusterName}}
contexts:
- context:
    cluster: {{.ClusterName}}
    namespace: {{.Namespace}}
    user: {{.Username}}.{{.ClusterName}}
  name: {{.ClusterName}}
current-context: {{.ClusterName}}
kind: Config
preferences: {}
users:
- name: {{.Username}}.{{.ClusterName}}
  user:
    token: {{.Token}}
`
