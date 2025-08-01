# EasyDev sAPI

---
## Навигация

- [Описание](#Описание)
- [Контакты](#Контакты)
- [Лицензия](#Лицензия)
- [Хост](#хост)
- [Безопасность](#безопасность)
- [Swagger](#swagger)
- [User API](#user-api)
    - [Регистрация пользователя](#регистрация-пользователя)
    - [Аутентификация пользователя](#аутентификация-пользователя)
    - [Обновление токена](#обновление-токена)
    - [Получение профиля пользователя](#получение-профиля-пользователя)
    - [Обновление профиля пользователя](#обновление-профиля-пользователя)
    - [Изменение пароля](#изменение-пароля)
- [Admin API](#admin-api)
    - [Получение всех пользователей](#получение-всех-пользователей)
    - [Получение профиля пользователя](#получение-профиля-пользователя-1)
    - [Обновление прав пользователя](#обновление-прав-пользователя)
    - [Обновление данных пользователя](#обновление-данных-пользователя)
    - [Блокировка/разблокировка пользователя](#блокировкаразблокировка-пользователя)
    - [Удаление пользователя](#удаление-пользователя)
- [Управление задачами (Todo)](#управление-задачами-todo)
    - [Создание задачи](#создание-задачи)
    - [Получение всех задач](#получение-всех-задач)
    - [Получение задачи по ID](#получение-задачи-по-id)
    - [Обновление задачи](#обновление-задачи)
    - [Удаление задачи](#удаление-задачи)

---

## Описание

RESTful API сервис для EasyDev. Предоставляет функции управления пользователями и задачами.

## Контакты

- **Имя**: s4bb4t
- **Email**: s4bb4t@yandex.ru

## Хост

- **URL**: [http://easydev.club/api/v1](http://easydev.club/api/v1)

## Безопасность

- **Описание**: Для доступа к защищенным маршрутам требуется JWT Bearer токен. Формат: `Bearer <token>`

### Swagger

- **Путь**: [Swagger документация](http://easydev.club/api/v1/swagger/index.html#)

---

## User API

### Регистрация пользователя

- **Путь**: `/auth/signup`
- **Метод**: POST
- **Описание**: Регистрирует нового пользователя.
- **Параметры**:
    - **User** (тело запроса): Полные данные пользователя для регистрации.
      ```json
      {
        "login": "string",
        "username": "string",
        "password": "string",
        "email": "string",
        "phoneNumber": "string"
      }
      ```
- **Ответы**:
    - **201 Created**: Успешная регистрация. Возвращает данные пользователя.
    - **400 Bad Request**: Ошибка десериализации запроса или неверный ввод.
    - **409 Conflict**: Пользователь уже существует.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Аутентификация пользователя

- **Путь**: `/auth/signin`
- **Метод**: POST
- **Описание**: Аутентифицирует пользователя и возвращает JWT токены.
- **Параметры**:
    - **AuthData** (тело запроса): Данные для аутентификации.
      ```json
      {
        "login": "string",
        "password": "string"
      }
      ```
- **Ответы**:
    - **200 OK**: Успешная аутентификация. Возвращает JWT токены.
      ```json
      {
        "access": "string",
        "refresh": "string"
      }
      ```
    - **400 Bad Request**: Ошибка десериализации запроса или неверный ввод.
    - **401 Unauthorized**: Неверные учетные данные.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Обновление токена

- **Путь**: `/auth/refresh`
- **Метод**: POST
- **Описание**: Обновляет токен доступа на основе предоставленного refresh токена.
- **Параметры**:
    - **RefreshToken** (тело запроса): Refresh токен пользователя.
      ```json
      {
        "refreshToken": "string"
      }
      ```
- **Ответы**:
    - **200 OK**: Успешное обновление токена. Возвращает новые JWT токены.
      ```json
      {
        "access": "string",
        "refresh": "string"
      }
      ```
    - **400 Bad Request**: Ошибка десериализации запроса.
    - **401 Unauthorized**: Неверные учетные данные или токен истек.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Получение профиля пользователя

- **Путь**: `/user/profile`
- **Метод**: GET
- **Описание**: Возвращает профиль текущего аутентифицированного пользователя.
- **Ответы**:
    - **200 OK**: Возвращает данные профиля пользователя.
      ```json
      {
        "id": 1,
        "username": "string",
        "email": "string@string.com",
        "date": "2024-09-15 16:06:15",
        "isBlocked": false,
        "isAdmin": true,
        "phoneNumber": "+79134210880"
      }
      ```
    - **400 Bad Request**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Обновление профиля пользователя

- **Путь**: `/user/profile`
- **Метод**: PUT
- **Описание**: Обновляет профиль пользователя с новыми данными.
- **Параметры**:
    - **PutUser** (тело запроса): Обновленные данные пользователя.
      ```json
      {
        "username": "string",
        "email": "string",
        "phoneNumber": "string"
      }
      ```
- **Ответы**:
    - **200 OK**: Профиль успешно обновлен.
    - **400 Bad Request**: Ошибка десериализации запроса или логин/электронная почта уже используются.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Изменение пароля

- **Путь**: `/user/profile/reset-password`
- **Метод**: PUT
- **Описание**: Обновляет пароль пользователя.
- **Параметры**:
    - **Pwd** (тело запроса): Новый пароль.
      ```json
      {
        "password": "string"
      }
      ```
- **Ответы**:
    - **200 OK**: Пароль успешно изменен.
    - **400 Bad Request**: Ошибка десериализации запроса.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

---

## Admin API

### Получение всех пользователей

- **Путь**: `/admin/users`
- **Метод**: GET
- **Описание**: Получает список пользователей с возможностью фильтрации и сортировки.
- **Параметры запроса**:
    - **search** (строка, необязательно): Фильтр по ключевому слову в имени или электронной почте.
    - **sortBy** (строка, необязательно): Поле для сортировки (например, "username", "email").
    - **sortOrder** (строка, необязательно): Направление сортировки ("asc" или "desc").
    - **isBlocked** (логическое, необязательно): Фильтрация по статусу блокировки.
    - **limit** (целое число, необязательно): Количество элементов на странице (по умолчанию 20).
    - **offset** (целое число, необязательно): Смещение для пагинации (по умолчанию 0).
- **Ответы**:
    - **200 OK**: Возвращает список пользователей с метаинформацией.
      ```json
      {
        "data": [
          {
            "id": 1,
            "username": "string",
            "email": "string",
            "date": "string",
            "isBlocked": false,
            "isAdmin": false,
            "phoneNumber": "string"
          }
        ],
        "meta": {
          "totalAmount": 1,
          "sortBy": "id",
          "sortOrder": "asc"
        }
      }
      ```
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Получение профиля пользователя

- **Путь**: `/admin/users/{id}`
- **Метод**: GET
- **Описание**: Получает профиль пользователя по его ID.
- **Параметры**:
    - **id** (путь): ID пользователя.
- **Ответы**:
    - **200 OK**: Возвращает данные профиля пользователя.
      ```json
      {
        "id": 1,
        "username": "string",
        "email": "string@string.com",
        "date": "2024-09-15 16:06:15",
        "isBlocked": false,
        "isAdmin": true,
        "phoneNumber": "+79134210880"
      }
      ```
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Обновление прав пользователя

- **Путь**: `/admin/users/{id}/rights`
- **Метод**: PUT
- **Описание**: Обновляет права доступа пользователя.
- **Параметры**:
    - **id** (путь): ID пользователя.
    - **Roles** (тело запроса): Обновленные права доступа.
      ```json
      {
        "isAdmin": true,
        "isBlocked": false
      }
      ```
- **Ответы**:
    - **200 OK**: Права успешно обновлены.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Обновление данных пользователя

- **Путь**: `/admin/users/{id}`
- **Метод**: PUT
- **Описание**: Обновляет данные пользователя.
- **Параметры**:
    - **id** (путь): ID пользователя.
    - **User** (тело запроса): Новые данные пользователя.
      ```json
      {
        "username": "string",
        "email": "string@string.com",
        "phoneNumber": "string"
      }
      ```
- **Ответы**:
    - **200 OK**: Данные успешно обновлены.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Блокировка/разблокировка пользователя

- **Путь**: `/admin/users/{id}/block`
- **Метод**: POST
- **Описание**: Блокирует пользователя.
- **Путь**: `/admin/users/{id}/unblock`
- **Метод**: POST
- **Описание**: Разблокирует пользователя.
- **Параметры**:
    - **id** (путь): ID пользователя.
- **Ответы**:
    - **200 OK**: Статус успешно обновлен.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Удаление пользователя

- **Путь**: `/admin/users/{id}`
- **Метод**: DELETE
- **Описание**: Удаляет пользователя по его ID.
- **Параметры**:
    - **id** (путь): ID пользователя.
- **Ответы**:
    - **200 OK**: Пользователь успешно удален.
    - **404 Not Found**: Пользователь не найден.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

---

## Управление задачами (Todo)

### Создание задачи

- **Путь**: `/todos`
- **Метод**: POST
- **Описание**: Создает новую задачу.
- **Параметры**:
    - **Todo** (тело запроса): Данные задачи.
      ```json
      {
        "title": "string",
        "isDone": false
      }
      ```
- **Ответы**:
    - **201 Created**: Задача успешно создана.
    - **400 Bad Request**: Ошибка десериализации запроса.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Получение всех задач

- **Путь**: `/todos`
- **Метод**: GET
- **Описание**: Получает список всех задач.
- **Параметры запроса**:
    - **status** (строка, необязательно): Фильтрация по статусу.
    - **limit** (целое число, необязательно): Количество элементов на странице (по умолчанию 20).
    - **offset** (целое число, необязательно): Смещение для пагинации (по умолчанию 0).
- **Ответы**:
    - **200 OK**: Возвращает список задач.
      ```json
      {
        "data": [
          {
            "id": 1,
            "title": "string",
            "isDone": false,
            "created": "2024-09-15T16:06:15Z"
          }
        ],
        "meta": {
          "total": 100,
          "limit": 20,
          "offset": 0
        }
      }
      ```
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Получение задачи по ID

- **Путь**: `/todos/{id}`
- **Метод**: GET
- **Описание**: Получает задачу по ее ID.
- **Параметры**:
    - **id** (путь): ID задачи.
- **Ответы**:
    - **200 OK**: Возвращает данные задачи.
      ```json
      {
        "id": 1,
        "title": "string",
        "isDone": false,
        "created": "2024-09-15T16:06:15Z"
      }
      ```
    - **404 Not Found**: Задача не найдена.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Обновление задачи

- **Путь**: `/todos/{id}`
- **Метод**: PUT
- **Описание**: Обновляет данные задачи.
- **Параметры**:
    - **id** (путь): ID задачи.
    - **Todo** (тело запроса): Новые данные задачи.
      ```json
      {
        "title": "string",
        "isDone": false
      }
      ```
- **Ответы**:
    - **200 OK**: Задача успешно обновлена.
    - **400 Bad Request**: Ошибка десериализации запроса.
    - **404 Not Found**: Задача не найдена.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.

### Удаление задачи

- **Путь**: `/todos/{id}`
- **Метод**: DELETE
- **Описание**: Удаляет задачу по ее ID.
- **Параметры**:
    - **id** (путь): ID задачи.
- **Ответы**:
    - **200 OK**: Задача успешно удалена.
    - **404 Not Found**: Задача не найдена.
    - **500 Internal Server Error**: Внутренняя ошибка сервера.