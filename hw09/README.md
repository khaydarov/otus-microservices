### Installation

## Prerequisite

### cert-manager installation

```shell
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.3/cert-manager.yaml
```

### Jaeger installation

```shell
kubectl create namespace observability
kubectl create -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.33.0/jaeger-operator.yaml -n observability
kubectl apply -f jaeger/simplest.yaml -n observability
```

## Application

PostgreSQL installation
```shell
helm install db bitnami/postgresql -f db/postgres/config.yaml
```

Install necessary databases and tables
```shell
kubectl exec -it db-postgresql-0 -- psql postgresql://postgres:postgres@localhost:5432 < db/postgres/databases.sql
kubectl exec -it db-postgresql-0 -- psql postgresql://postgres:postgres@localhost:5432/orders < db/postgres/orders.sql
kubectl exec -it db-postgresql-0 -- psql postgresql://postgres:postgres@localhost:5432/payments < db/postgres/payments.sql
kubectl exec -it db-postgresql-0 -- psql postgresql://postgres:postgres@localhost:5432/inventory < db/postgres/inventory.sql
kubectl exec -it db-postgresql-0 -- psql postgresql://postgres:postgres@localhost:5432/shipment < db/postgres/shipment.sql
```

Create a namespace

```shell
kubectl create namespace otus-hw09
```

Select created namespace

```shell
kubectl config set-context --current --namespace=otus-hw09
```

### Order Service

* Uses Saga orchestration approach to provide distributed transaction
```shell
cd order
helm install order-service .helm
```

### Payment Service
```shell
cd payments
helm install payments-service .helm
```

### Inventory Service
```shell
cd inventory
helm install inventory-service .helm
```

### Shipment Service
```shell
cd shipment
helm install shipment-service
```

### API Gateway

```shell
kubectl apply -f api-gateway/ingress.yaml
```

### Testing

Run Postman test scenario

```shell
bash .postman-test.sh
```

### TODO

* setup jaeger/tracing