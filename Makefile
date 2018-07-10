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
	docker build -t youkoulayley/api-kubeum .
	docker push youkoulayley/api-kubeum:latest

run: docker
	docker run -p 8080:8080 --rm -it --name=get-kubeconfig youkoulayley/api-kubeum

run-kubernetes: docker
	kubectl apply -f my_manifests/
	kubectl patch deployment api-kubeum -n kube-system -p \
	  "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"

docker-minikube: build
	eval $(minikube docker-env) && \
	docker build -t youkoulayley/api-kubeum .
	docker push youkoulayley/api-kubeum:latest

run-minikube: docker-minikube
	kubectl apply -f manifests/
	kubectl patch deployment api-kubeum -n kube-system -p \
      "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}"