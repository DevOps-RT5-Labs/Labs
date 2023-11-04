
IMAGE_NAME = "devops-demo-app"

.PHONY: image
image:
	docker build -t $(IMAGE_NAME) 1-docker/apps

.PHONY: cluster
cluster:
	kind create cluster --config 2-k8s/kind.yml
	kind load docker-image devops-demo-app -n devops-lab
	kubectl apply -f 2-k8s/k8s/webapp.yml


.PHONY: argo
argo:
	k apply -f 2-k8s/tools/argo.yml