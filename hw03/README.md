## ДЗ 2

- Создать простой RESTful CRUD-сервис.
- Добавить базу данных для приложения.
- Конфигурация приложения должна хранится в Configmaps.
- Доступы к БД должны храниться в Secrets.
- Первоначальные миграции должны быть оформлены в качестве Job-ы, если это требуется.
- Ingress-ы должны также вести на url arch.homework/
- Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя. Важно: в postman коллекции использовать базовый url - arch.homework.

*Команда установки БД из helm, вместе с файлом values.yaml.* (+5 балла за шаблонизацию приложения в helm чартах)

```shell
helm install otus-task2 .helm
```

*Команда для нагрузочного тестирования*
```shell
while true; do ab -n 50 -c 5 http://arch.homework/; sleep 2; done;
```

## Todo
2. make screenshots for PRS, latency
3. install and configure ingress

