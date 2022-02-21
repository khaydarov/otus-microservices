## Prerequisites

- Kubernetes
- Helm
- Nginx Ingress
- Docker

## Installation

Auth service setup

```shell
cd auth-app
helm install otus-hw05-auth-app .helm
```

User service setup

```shell
cd user-app
helm install otus-hw05-user-app .helm
```

### Ingress setup

```shell
cd api-gateway
kubectl apply -f nginx-ingress/ingress.yaml
```

### Traefik setup

```shell
cd api-gateway
...
```

### Forward-Auth
```shell
kubectl apply -f nginx-ingress/auth.yaml
```

Service communation schema

![](auth-schema.png)

## TODO

- Add migration for auth-service
- Postman tests