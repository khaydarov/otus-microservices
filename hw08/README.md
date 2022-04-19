### Пользовательские сценарии

- Я как пользователь могу зарегистрироваться в интернет-магазине
- Я как пользователь могу авторизоваться в интернет-магазине
- Я как пользователь могу пополнить баланс своего счёта
- Я как пользователь могу получить баланс своего счёта
- Я как пользователь могу создать заказ
- Я как пользователь могу получить список уведомлений о заказах

### Предметрные области

- Пользователь
- Заказ
- Уведомление
- Счёт

### Системные действия

- Регистрация
- Авторизация
- Пополнение счёта
- Получение баланса 
- Создание заказа
- Получение уведомлений

### Таблица


### Сервисы

- User service
- Billing service
- Order service
- Notification service

### Методы API

Публичные
```
PUT  /api/v1/users/{id}
POST /api/v1/users/register
POST /api/v1/users/login
POST /api/v1/billing/account/deposit
POST /api/v1/orders
POST /api/v1/notifications
```

Внутренние
```
POST /api/v1/billing/account
POST /api/v1/billing/account/withdraw
```

### TODO

- Bounded Context Canvas
- Communication schema
- API methods best practices
