## ДЗ 1

1. Создать сервис который отвечает на порту 8000
2. Добавить Endpoint `/health` в сервисе который вернет `{"status": "OK"}`
3. Залить образ на DockerHub
4. Написать все манифесты для сущностей Deployment, Service, Ingress
5. Добавить Liveness, Readiness пробы
6. Количество реплик сервиса должно быть не меньше 2
7. Хост в Ingress должен быть `arch.homework`

Задание со звёздочкой: В Ingress-е должно быть правило, которое форвардит все запросы с /otusapp/{student name}/* на сервис с rewrite-ом пути. Где {student name} - это имя студента.

### Инструкция по запуску

```shell
kubectl apply -f deployment.yaml -f service.yaml -f ingress.yaml
```

