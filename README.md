# Сервис подсчёта арифметических выражений
Привет!👋

Небольшой проект для курса Яндекс Лицея. Имеет несложную структуру и всего один эндпоинт.
#
**Запуск:**

    go run ./cmd/main.go

#
### Примеры работы

Успешно:

    curl --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "2+2*2"
    }'

Ошибка 422:
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": 2+2*2
}'
```

Ошибка 500:
```

    curl --location 'localhost/api/v1/calculate' \
    --header 'Content-Type: application/json'
    ```
