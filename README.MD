### WALLET API
Обработка транзакций в существующих кошельках.

#### Выполненные эндпоинты
1. POST   /api/v1/wallets
```
{
    "wallet_id":"3422b448-2460-4fd2-9183-8000de6f8343",
    "operation_type":"DEPOSIT",
    "amount":122
}
```

wallet_id              - в формате UUID  
operation_type - DEPOSIT или WITHDRAW  
amount                  - число от 1 до 2^32  

2. GET    /api/v1/wallets/:id  
id - UUID кошелька

#### Запуск приложения
1. Клонировать репозиторий
2. Убрать 'example' в config.env файле
3. Выставить желаемые значения (дефолтные работают)
4. Выполнить команду
```docker compose --env-file config.env up```
