# Run #
```make up```

# Down #
```make down```

# Requests #
- Register new user

```curl -d '{"email": "test@test.com", "user_name": "test user", "city": "Limassol", "country": "Cyprus", "currency_code": "EUR"}'  -H "Content-Type: application/json" -X POST http://:9090/user```
Expected response:

```{"user_id":"fd0345cf-b6f9-4456-8933-f6755b83a26b"}```

- Update currency rate

```curl -d '{"currency_code": "EUR", "rate": 90, "valid_date": "2021-10-02T01:00:00Z"}' -H "Content-Type: application/json" -X POST http://:9090/rate_update```

Expected response:

```{"result":"ok"}```

- Deposit (only user's wallet currency)

```curl -d '{"user_id": "fd0345cf-b6f9-4456-8933-f6755b83a26b", "amount": 50}' -H "Content-Type: application/json" -X POST http://:9090/deposit```

Expected response:

```{"result":"ok"}```

- Transfer

```curl -d '{"user_id_from": "fd0345cf-b6f9-4456-8933-f6755b83a26b", "user_id_to": "eca6ae47-4e7a-4d2a-ad39-710ad0aa357e", "amount": 300, "currency_code": "EUR"}'  -H "Content-Type: application/json" -X POST http://:9090/transfer```

Expected response:

```{"result":"ok"}```

- Report generation

```make report user_id=eca6ae47-4e7a-4d2a-ad39-710ad0aa357e begin_time=2021-10-03T08:43:01Z filename=/full/path/to/report.csv```

If filename is omited, program use stdout.

# Заметки #
- предполагается что у юзера только 1 кошелек
- все amount указаны в "копейках". По стандарту (en.wikipedia.org/wiki/ISO_4217) для некотрых валют может быть 3 или 4 знака после запятой - для таких валют здесь логика не реализована
- rate также задан через int. Реализован rate для 100 ед валюты. Т.е. если rate = 110, это означается что за 110 ед валюты можно купить 100 USD. В  реальной ситуации можно было бы изменить коэффициент - скажем, считать за 1000 USD или выше, чтобы избегать дробей для удобства работы. 
- при конвертации валют округление происходит по математическим правилам
- во многих местах не проверяется существование пользователя и/или других сущностей (проверка сделана только для currency rate, для остального можно будет сделать аналогично)
- по-хорошему отчет лучше делать через постановку таска, например, в реббит и формировать его отдельным сервисом асинхронно. Здесь для упращения все сделано отдельной программой
