# Сервис подсчёта арифметических выражений
Привет!👋

Небольшой проект для курса Яндекс Лицея. 

Так как проект, на мой взгляд имеет несложную структуру:
 - Все, что связано с работой хэндлера помещено в pkg/handler
 - Все, что связано с бизнес-логикой (калькулятор) помещено в pkg/service
 - Запуск сервера осуществляется в cmd/main.go
#
**Установка:**

    git clone https://github.com/child6yo/y-lms-calculator
    cd y-lms-calculator

**Запуск:**

    go run ./cmd/main.go

 - Запуск осуществляется по порту 8000.

#

Сервер имеет один эндпоинт в /api/v1/calculate

Существуют три варианта ответа сервера:
 - 200 OK, если хэндлер успешно обработал запрос
 - 422 Unprocessable Content, если ошибка в запросе клиента (возвращается либо при ошибке в работе калькулятора, а значит при неверных данных, либо при невозможности считать json в структуру, а значит данные представлены не в виде строки).
 - 500 Internal Server Error, если ошибка на стороне сервера (на практике возвращается при пустом теле запроса)

### Примеры работы

Успешно:

    curl --location 'localhost:8000/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "2+2*2"
    }'

Ошибка 422:
```
curl --location 'localhost:8000/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": 2+2*2
}'
```

Ошибка 500:
```

    curl --location 'localhost:8000/api/v1/calculate' \
    --header 'Content-Type: application/json'
    ```
