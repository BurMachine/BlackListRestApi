# BlackListRestApi
Микросервис предназначен для хранения информации о пользователях, которые были добавлены в черный список. Микросервис обеспечивает возможность добавления, удаления и поиска пользователей в черном списке.

Сервис разработан на языке Go и использует в качестве хранилища PostgreSQL. 
В качестве веб фреймворка используется библиотека `gorilla/mux`. Для PostgreSQL использовал драйвер `jackc/pgx`.

### Пример использования сервиса
#### Получение токена доступа для использования api методов

`curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8080/auth`

```json
{
  "name": "artur"
}
```
Ответ:


```json
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

#### Добавление в черный список

`curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8080/add`

```json
{
  "user_phone": "string",
  "user_name": "string",
  "reason": "string",
  "admin_name": "string"
}
```

Ответ:

Content-Type: application/json

```json
OK
```

#### Удаление пользователя
`curl http://localhost:8080/delete?id=1`

Ответ
```json
OK
```

#### Поиск пользователей
`curl http://localhost:8080/serch?number=string`

Ответ:
```json
{
  [
    {
      "user_phone": "string",
      "user_name": "string",
      "reason": "string",
      "date": 01-01-2023
      "admin_name": "string"
    }
  ]
}
```

### Запуск сервиса:

`make build` - первая сборка

`make restart` - пересборка

Для того чтобы использовать сервис локально нужно поднять `PostgreSQL` на порту `5432` и `make local-build`

Для генерации swagger-документации `make swagger`

Для удаления докер контейнеров `make rm_containers`

Сервис распространяется в виде докер контенера и содержит помимо БД PostgreSQL еще и веб-версию pgAdmin для удобства мониторинга


### Swagger-документация:
![Image alt](pics/photo_5458794079509923653_w.jpg)


