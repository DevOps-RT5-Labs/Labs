# Lab 3 → Kubernetes
> Prepared by Mohamed Sofiene Barka, Mohamed Rafraf, Jihene Ben Tekaya
> Note that in my examples I use `k` as an alias for `kubectl`

## Create a Cluster
We can create a Kubernetes cluster using various methods, from `kubeadm`, `k3s`, `minikube`, `microk8s`, `kind` or any managed service. For this lab, we will use `kind` to create a cluster with 3 nodes. `kind` is easily configurable so we will use the following YAML file to create our cluster:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: devops-lab
nodes:
- role: control-plane
  extraPortMappings:
    - containerPort: 31000
      hostPort: 31000
- role: worker
- role: worker
```
![](assets/kind.png)


## Deploy a web app on Kubernetes, and expose it
We will deploy the same application that we have made from the previous lab.
> We should load the docker image to the cluster nodes, as a requirement for `kind` to work.

```bash
kind load docker-image devops-demo-app -n devops-lab
```

We can deploy the application using `webapp.yml` file in the current directory. It contains the following:
- A namespace `web-app`
- A Deployment `web` with 2 replicas, with the strategy `RollingUpdate` and `maxSurge` and `maxUnavailable` being set
- A Service `web` that exposes the deployment on port `31000` for external access (NodePort)
- A Deployment `kv` with 1 replica (As a database, it would be more suited to use a StatefulSet, but since we are using a simple In-Memory database with no persistence, we can use a Deployment)
- A Service `kv` that exposes the deployment on port `6379` for internal access (ClusterIP)

![](assets/deps.png)

When we update the response of one of the endpoints, from "OK" to "DONE", we update the image and apply the changes to the cluster. We can see that the pods are being updated one by one, and the app overall is still available.

![](assets/rolling.png)

## Monitoring

### Grafana and Prometheus

<!--
- Test Metrics and Logging with Any tools.
-->

### GitOps

As part of a delivery pipeline for Kubernetes, we can use GitOps to deploy our application with ArgoCD.

We first need to install ArgoCD CRDs and the ArgoCD itself:

```bash
k create ns argocd
k apply -n argocd -f 2-k8s/tools/argo.yml
```

We change its service type to `NodePort` to be able to access it from outside the cluster:

```bash
k patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
```

We retrieve the password for the `admin` user:

```bash
k get secrets/argocd-initial-admin-secret -n argocd -o yaml | grep password | awk '{print $2}' | base64 -d
```

And finally we can export the service port to port `8080` on our machine:

```bash
k port-forward -n argocd service/argocd-server 8080:80
```

And we can access the ArgoCD UI on `localhost:8080` using `admin` as username and the password we retrieved earlier:
![](assets/argo.png)

If we try to delete the earlier deployment, and create an `Application` in ArgoCD (For simplicity, we have prepared a YAML file for the application in `2-k8s/tools/argo-app.yml`), we can see that the application is deployed again automatically:
![](assets/argo-app.png)


Note that the web apps are broken because of an old commit that we have pushed to the repository, but we can fix that by pushing a new commit to the repository. And we can see that the application is updated again. and up and running, thanks to ArgoCD.