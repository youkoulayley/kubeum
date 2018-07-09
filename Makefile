.PHONY: all fmt vendor build

fmt:
	go fmt api/*.go
	go fmt api/bootstrap/*.go
	go fmt api/controllers/*.go
	go fmt api/models/*.go
	go fmt api/tmpl/*.go
	go fmt cli/*.go
	go fmt cli/cmd/*.go

vendor:
	glide update

build: fmt
	GOOS=linux go build -o ./api-kubeum ./api/
	GOOS=linux go build -o ./kubeum ./cli/

docker: build
	docker build -t youkoulayley/kube-user-mgmt .
	docker push youkoulayley/kube-user-mgmt:latest

run: docker
	docker run -p 8080:8080 --rm -it --name=get-kubeconfig youkoulayley/kube-user-mgmt

run-kubernetes: docker
	kubectl apply -f my_manifests/
	kubectl patch deployment kube-get-kubeconfig -n kube-system -p \
	  "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"

docker-minikube: build
	eval $(minikube docker-env) && \
	docker build -t youkoulayley/kube-user-mgmt .
	docker push youkoulayley/kube-user-mgmt:latest

run-minikube: docker-minikube
	kubectl apply -f manifests/
	kubectl patch deployment kube-user-mgmt -n kube-system -p \
      "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"