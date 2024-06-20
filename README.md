### Выполнил Степанов Алексей Александрович

Для проекта я использовал Go 1.22.2, БД PostgreSQL, фреймворк [Fiber](https://github.com/gofiber/fiber)

## Запуск:

Запуск dev-среды с помощью docker-compose:

```bash
docker-compose up
```
Среда, в которой происходит запуск, должна иметь переменные окружения:
- DB_HOST - хост базы данных.
- DB_PORT - порт для доступа к базе данных.
- DB_USER - логин пользователя БД.
- DB_PASSWORD - пароль пользователя БД.
- DB_NAME - имя пользователя БД.

Запуск с помощью go run:

```bash
go run ./cmd/web/main.go
# Флаги:
# -override_tables=true : запуск с автоматическим созданием таблиц в БД
# -config=config.toml : путь к конфигурационному файлу
# -use_env=false : флаг на использование данных о БД из окружения, а не файла конфигурации
```

При запуске в файл конфигурации формата [TOML](https://toml.io) необходимо поместить:

```toml
addr        # используемый в проекте порт
# данные о БД также можно хранить в окружении, в дапнном случае при запуске необходимо установить -use_env=true
db.host     # хост БД
db.port     # порт для доступа к БД
db.user     # логин пользователя БД
db.password # логин пользователя БД
db.name     # имя пользователя БД
db.driver   # используемый driver БД (в настоящий момент доступен только postgres)
```

## PostgreSQL Query для создания таблицы в БД вручную:

```sql
CREATE TABLE IF NOT EXISTS links (
	id SERIAL PRIMARY KEY,
	link TEXT
);
```

## UI Swagger доступен по адресу `/swagger`

Документация располагается в папке [docs](./docs/)

Для генерации документации использовалась утилита [Swag](https://github.com/swaggo/swag):

```bash
# Установка Swag
go install github.com/swaggo/swag/cmd/swag@latest
# Генерация документации
swag init -d ./cmd/web,./internal -o ./docs
```
