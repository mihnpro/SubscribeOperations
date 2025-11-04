# SubscriptionService

Микросервис для управления подписками пользователей с REST API, построенный на Go и PostgreSQL.

## Функциональность

### Основные возможности
- **CRUDL операции** над подписками пользователей
- **Расчет стоимости** подписок за выбранный период
- **Фильтрация** по пользователю и названию подписки
- **Swagger документация** API
- **Docker контейнеризация**

### Модель данных подписки
```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "2023-01-01T00:00:00Z",
  "end_date": "2023-12-31T00:00:00Z"
}
```
### Документация
- localhost:8080/swagger/index.html


### API Endpoints

#### Подписки (CRUDL)
- **POST	/createSub**	Создать новую подписку
- **GET	/getSub/{subscriptionID}**	Получить подписку по ID
- **PUT	/updateSub**	Обновить подписку
- **DELETE	/deleteSub**	Удалить подписку
- **GET	/getAllSubs**	Получить список всех подписок

#### Аналитика
- **POST	/getFullSubPriceByPeriod**	Суммарная стоимость подписок за период

## Быстрый старт

### Предварительные требования

- **Docker**
- **Docker Compose**

### Запуск приложения

#### Клонируйте репозиторий

```bash
git clone 
cd subscription-service
```

#### Запустите приложение
```bash
docker-compose up --build
```
#### Приложение будет доступно по адресу:

- **API:** http://localhost:8080
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **PostgreSQL:** localhost:5433

## Переменные окружения
Создайте файл .env в корне проекта:


```.env
# Database
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=
DB_TIMEZONE=

# PostgreSQL
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

# App
PORT=
```

## Технический стек

- **Backend:** Go 1.24+
- **Database:** PostgreSQL 15
- **ORM:** sqlx
- **Migrations:** golang-migrate
- **API Documentation:** Swagger/OpenAPI 3.0
- **Containerization:** Docker + Docker Compose
- **Configuration:** .env файлы





