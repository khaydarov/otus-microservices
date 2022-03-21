## Installation

Create a namespace

```shell
kubectl create namespace otus-hw06
```

Select created namespace

```shell
kubectl config set-context --current --namespace=otus-hw06
```


## Testing

Run Postman test scenario

```shell
bash .postman-test.sh
```

---
- Deployment
    - Dockerfile for app
    - Dockerfile for migration
    - Helm charts
- Metrics for prom