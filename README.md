![badge](https://drone.journeytotheit.ovh/api/badges/Youkoulayley/kubeum/status.svg)

# Kube User Management

## API

This is a simple app to run inside the cluster for managing Kubernetes Users.
The purpose of this app is to be a wrapper in front of the APIServer to provide data to the kubeum commandline and the UI.

If the Kubernetes evolves, juste the API need to be modify.

### Installation

You can install the app by doing those commands :

```bash
kubectl apply -f https://raw.githubusercontent.com/Youkoulayley/kubeum/master/manifests/rbac.yaml
kubectl apply -f https://raw.githubusercontent.com/Youkoulayley/kubeum/master/manifests/deployment.yaml
```

These resources will be created : 

* All `RBAC` resources needed ;
* A `deployment`.

### Environment variables
|**Name**|**Description**|**Type**|**Default**|
|--------|---------------|--------|-----------|
|*APP_SERVER_URL*|Define the URL that the kubeconfig will used to connect to the cluster|string|https://127.0.0.1:6443|
|*CA_FILE*|Path of the CA to connect to the cluster|string|/var/run/secrets/kubernetes.io/serviceaccount/ca.crt|
|*CLUSTER_NAME*|Name of the cluster to identify it in the kubeconfig|string|kubernetes|
|*LOG_LEVEL*|Define the log level of the application|Debug/Info/Warning/Error|Info|
|*PORT*|Port on which the application will listen|int|8080|

## CLI

This command line interfaces with the API and provides the same actions.

```bash
$ kubeum
Cli that permits you to manage your kubernetes users.

Usage:
  kubeum [command]

Available Commands:
  help        Help about any command
  kubeconfig  Generate a kubeconfig for a particular user.
  list        List all users of your Kubernetes Cluster.
  version     Print the version number of kubeum

Flags:
  -a, --api-kubeum string   Link To API Kubeum (default "http://localhost:31000")
      --config string       config file
  -h, --help                help for kubeum
  -t, --toggle              Help message for toggle

Use "kubeum [command] --help" for more information about a command.
```

## Todo

- [x] Setup the router
- [x] GET /health
- [x] GET /users
- [X] POST /kubeconfig
- [X] CLI
- [ ] UI
- [ ] Testing
- [ ] CI
- [X] Goreleaser

