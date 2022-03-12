## otus-microservices

1. Pros and cons of microservice architecture
    * Architecture and architect
    * Monoliths and microservices
    * Patterns of microservice architecture
    

2. Docker basics
   * Containerization. Overview
   * Docker components: engine, cli, registry
   * Building from Dockerfile
   * Practice: build, run, up, down, pull, push
    

3. Infrastructural patterns
   * CI/CD methodology
   * VM vs Containers
   * Deployment patterns
   * Service discovery 
   * Health check


4. Kubernetes basics (part 1)
   * Pods, ReplicaSets, Deployments


5. Kubernetes basics (part 2). [Homework #1](https://github.com/khaydarov/otus-microservices/tree/main/hw01)
   * ConfigMaps, Persistence Volumes, Persistence Volume Claims
   * Helm, Helm-dep, Ingress 


6. Kubernetes basics (part 3). [Homework #2](https://github.com/khaydarov/otus-microservices/tree/main/hw02)
   * Шаблонизация с помощью Helm
   * Сущности: Job, Secrets


7. Kubernetes. QA
8. Monitoring and alerting
   * USE, RED и Four Golden Signals
   * SLI, SLO, SLA
   * Metric collection patterns
    

9. Prometheus, Grafana. [Homework #3](https://github.com/khaydarov/otus-microservices/tree/main/hw03)
   * Prometheus
   * Grafana
   * AlertManager
   * PromQL 
    

10. Service mesh on the example of Istio [Homework #4](https://github.com/khaydarov/otus-microservices/tree/main/hw04)
    * Service Mesh architecture


11. Authorization and authentication in microservice architecture
    * Auth patterns in monoliths
    * Identity Provider и OIDC
    * Token-Based authentication, JWT
    * Auth-Proxy
    

12. Backend for frontends. API Gateway. [Homework #5](https://github.com/khaydarov/otus-microservices/tree/main/hw05)
    * API Gateway
    * Backend for Frontends
    * Auth patterns in API Gateway
    * Circuit Breaker, Retry


13. Asynchronous and synchronous API
    * Message Bus, Enterprise Service Bus
    * CQRS, Event Sourcing  
    * Orchestration and choreography
    * API versioning
    * IDL, API design first
    * Anemic API vs Rich API


14. Event Driven Architecture
    * Designing event driven patterns
    * Using event driven patterns


15. Distributed message brokers on the example of Kafka
    * Kafka    


16. Consistent data maintenance patterns (Stream processing). [Homework №6]()
    * Transactional Log
    * Stream processing
    * Event Sourcing
    * Change Data Capture


17. GraphQL, gRPC
18. RESTful


19. Идемпотентность и коммутативность API в HTTP и очередях. [Домашнее задание №7]()
20. Тестирование микросервисов (часть 1)
21. Тестирование микросервисов (часть 2)
22. DDD и модульные монолиты (часть 1)
23. DDD и модульные монолиты (часть 2)
24. Паттерны декомпозиции микросервисов. [Домашнее задание №8]()
25. От монолита к микросервису
26. Введение в распределенные системы
27. Распределенные транзакции. [Домашнее задание №9]()
28. Паттерны кэширования и основные принципы
29. Шардирование
30. CP системы
31. AP системы
32. Роль архитектора
33. Стомость архитектуры. Артефакты архитектуры
