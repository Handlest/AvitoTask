# Тестовое задание для стажёра Backend в Авито "Сервис динамического сегментирования пользователей"

## Выполненные задания:
- Основное задание(минимум) выполнено
- Доп. задание 1 (история попадания/выбывания пользователя) выполнено
- Доп. задание 2 (Возможность задавать TTL пользователя в сегменте) выполнено частично (Добавлен метод в БД, который можно вызывать с помощью cron и возможность задания поля TTL при добавлении пользователя в сегмент)
- Создан Swagger файл(доступен при запущенном сервере по ссылке localhost:8000/swagger/index.html )

## Дополнительная информация:
- Выбраная СУБД: PostgreSQL
- Испольованные библиотеки: gin-gonic, viper, sqlx, swaggo
- Для миграций использовалась утилита migrate
- Добавлен makefile(команды будут описаны далее)

## <a href="https://www.postman.com/rikomero/workspace/avito-trainee-app/documentation/29339192-eb4a0e30-a290-48dd-920a-6f0da26dbe14?entity=request-fc5473ac-7517-422c-9420-6d8acf9b1af0">Документация с описанием запросов и примерами в postman

### Для запуска приложения(unix):
```
make build
make run
```
Если приложение запускается впервые, необходимо применить миграции к базе данных:
```
make migrate
```


### Для запуска приложения(windows):
```
docker-compose up --build avito-task
docker-compose up avito-task
```
Если приложение запускается впервые, необходимо применить миграции к базе данных:
```
migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up
```
После успешного запуска контейнеров, можно обращаться к серверу с помощью запросов описанных в документации: https://www.postman.com/rikomero/workspace/avito-trainee-app/documentation/29339192-eb4a0e30-a290-48dd-920a-6f0da26dbe14?entity=request-fc5473ac-7517-422c-9420-6d8acf9b1af0


### Альтернативный способ запуска приложения(Если первый не сработал)
Необходимо запустить докер с базой данных в контейнере, применить миграции и затем запустить сам сервер вне контейнера
```
docker run --name=AvitoTask -e POSTGRES_PASSWORD=password -p 5436:5432 -d --rm postgres  

migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up

go run cmd/main.go   
```

### Примечание
Необходимость повторого вызова docker-compose up после build вызвана медленной инициализацией базы данных. Данную проблему можно исправить, если при запуске копировать внутрь контейнера и вызывать скрипт wait-for-postgres.sh
