# ToDo List Web

Учебное приложение - сайт с todo list

## Шаги

1. добавляем окружение. Сайт с докер образами - hub.docker.com
2. добавляем Makefile
3. добавляем миграции
4. инициализируем приложение
``` bash
   go mod init github.com/zinovev-dm/golang-todoapp
```

5. добавляем пакеты
``` bash
   go get -u go.uber.org/zap
   go get github.com/kelseyhightower/envconfig
   go get github.com/google/uuid
   go get github.com/go-playground/validator/v10
   go get github.com/jackc/pgx/v5
   go get github.com/jackc/pgx/v5/pgxpool
```

6. проверяем курлом
``` bash
curl -X POST localhost:5050/api/v1/users

curl --location '127.0.0.1:5050/api/v1/users' \
--header 'Content-Type: application/json' \
--data '{
    "full_name": "Петров Иван"
}'

curl --location '127.0.0.1:5050/api/v1/users'

curl --location '127.0.0.1:5050/api/v1/users?limit=2'

curl --location '127.0.0.1:5050/api/v1/users/1' 
```
