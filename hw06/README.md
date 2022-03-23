## Installation

Create a namespace

```shell
kubectl create namespace otus-hw06
```

Select created namespace

```shell
kubectl config set-context --current --namespace=otus-hw06
```

Install and setup kafka
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install kafka bitnami/kafka -f kafka/config.yaml
```

Install postgres and setup user-app

```shell
cd user
helm install user-db bitnami/postgresql -f database/postgres/config.yaml
helm install user-app application/.helm
```

Install postgres and setup billing-app
```shell
cd billing
helm install billing-db bitnami/postgresql -f database/postgres/config.yaml
helm install billing-app application/.helm
```

Install `Order-Service`
```shell
cd order
helm install otus-hw06-order-service .helm
```

## Testing

Run Postman test scenario

```shell
bash .postman-test.sh
```

---
- Deployment
    - Dockerfile for migrations
    - Helm charts
- Set up probes
- Set up HPA
- Metrics for prometheus