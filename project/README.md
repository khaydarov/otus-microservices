## Система размещения рекламы на сайтах

Система позволяет создавать рекламу для продвижения своих товаров и услуг. 
Рекламодатели создают рекламу, настраивают таргетинг, устанавливают бюджет и стоимость показа, а владельцы веб-сайтов размещают специальный сгенерированный код для открутки рекламного блока.

### Пользовательские сценарии

- Пользователь регистрируется в нашей системе как рекламодатель или владелец веб-сайтов.
- При регистрации у пользователей создается счёт.
- Рекламодатель может пополнять счёт, а тратить только на открутку рекламы.
- Владелец веб-сайтов заработанные деньги может выводить на карту.
- Рекламодатель может создавать рекламу: загружает картинку, пишет заголовок и описание, настраивает таргетинг
- Рекламодатель и владелец веб-сайтов могут просматривать историю операций
- Рекламодатель может просматривать статистику показов созданных объявлений
- Владелец веб-сайтов может создавать разные сайты для показа рекламы

### Системные сценарии

- При создании рекламы деньги списываются сразу за количество показов
- Раз в сутки владельцам веб-сайтов начисляются деньги за показы рекламы

### Действующие лица

- Рекламодатель
- Владелец веб-сайтов

### Схема доменной модели

[UML Diagram](uml.drawio)

### Модель взаимодействия сервисов
 
- Использовать Microservice Canvas
- Добавить мониторинги и алерты
- Задокументировать сервис
- Добавить SLA | SLI | SLO
- Покрыть тестами
- Добавить трассировку запросов
- Service mesh: istio

### Публичные методы API системы

- [Спецификацию OpenAPI](api-composition/api.spec.yaml)

## Установка

Install Prometheus Stack

```shell
helm install prom prometheus-community/kube-prometheus-stack -f monitoring/prometheus.yaml -n monitoring
kubectl port-forward service/prom-grafana -n monitoring 9000:80 # expose grafana (pass: prom-operator)
kubectl port-forward service/prom-kube-prometheus-stack-prometheus -n monitoring 9090:9090 # expose prometheus
```

Install Istio addons
```shell
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/jaeger.yaml -n istio-system
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/kiali.yaml -n istio-system
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/prometheus.yaml -n istio-system
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/addons/grafana.yaml -n istio-system
```

PostgresSQL installation
```shell
helm install db bitnami/postgresql -f persistence/postgres/config.yaml
```