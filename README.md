# Разработка методологии и прототипа системы рейтингования поставщиков

**Команда "Бычий Цепень"**

## Использованные технологии:

- Js (фронтенд)
- PostgreSQL (хранилище данных)
- Docker (запуск сервиса)
- Gin (веб-фреймворк)
- golang-migrate/migrate (миграции БД)
- sqlx (работа с БД)
- logrus (логгирование)
- pq (драйвер для работы с postgres)
- viper (работа с файлами конфигураций)

## Quickstart:

Быстрый старт - `make run`, порт `localhost:8000`. Запуск контейнера с бэком и контейнера с базой данных + миграция.

Для корректной работы необходимо перенести данные из csv файла в базу

```text
docker cp {example}.csv <database docker container>:/tmp

COPY ks(participant_inn, participant_kpp, is_winner, ks_id, publish_date, price, customer_inn, customer_kpp, customer_type, kpgz, name, items, region_code, violations)
FROM '/tmp/ks.csv'
DELIMITER ';'
CSV HEADER;

COPY contracts(ks_id, contract_id, conclusion_date, price, customer_inn, customer_kpp, supplier_inn, supplier_kpp, violations, status)
FROM '/tmp/contracts.csv'
DELIMITER ';'
CSV HEADER;

COPY blocking(supplier_inn, supplier_kpp, reason, blocking_start_date, blocking_end_date)
FROM '/tmp/blocking.csv'
DELIMITER ';'
CSV HEADER;

COPY contreactexec(contract_id, upd_id, scheduled_delivery_date, actual_delivery_date, supplier_inn, supplier_kpp, customer_inn, customer_kpp)
FROM '/tmp/contract_execution.csv'
DELIMITER ';'
CSV HEADER;

```

Опционально - настройка `configs/config.yaml` и `.env`

## Эндпоинты:

- Получение главной страницы сайта

```curl
http://127.0.0.1:8000/
```

- Получение метрик для всех поставщиков в выбранной категории

```curl
GET http://127.0.0.1:8000/filter/kpgz/{kpgz}
```
Пример ответа:
```json
{
  "data": [
    {
      "inn": "7708333320",
      "failed_dedlines": 0,
      "avg_udp_contract": 351.4,
      "activity": 0.02,
      "total": -1.2
    }
  ]
}
```

- Получение метрик для поставщика по заданному ИНН

```curl
GET http://127.0.0.1:8000/filter/inn/{inn}
```
