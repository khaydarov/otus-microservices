### Installation

Create a namespace

```shell
kubectl create namespace otus-hw07
```

Select created namespace

```shell
kubectl config set-context --current --namespace=otus-hw07
```

Install and setup kafka
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install kafka bitnami/kafka -f kafka/config.yaml
```

Install postgres and setup billing-app
```shell
cd billing
helm install billing-db bitnami/postgresql -f database/postgres/config.yaml
helm install billing-app application/.helm
```

Install postgres and setup order-app
```shell
cd order
helm install order-db bitnami/postgresql -f database/postgres/config.yaml
helm install order-app application/.helm
```
