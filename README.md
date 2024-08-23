# Банкомат (Калькулятор размена валюты)

## Оформление
- Фреймворк Echo
- Конфигурация (хост, порт)
- Логирование (*TODO добавить уровень логирования*)
- Graceful shutdown
- *TODO unit тесты алгоритма*

## Описание
#### Принимаем запрос:
- amount - сумма денег
- banknotes - доступные номиналы банкнот
#### Ответ:
- все варианты размена отправленный суммы


## Алгоритм
Рекурсивный алгоритм, который реализует функциональность расчета всех вариантов размена для указанной суммы денег

## Адреса 
- ".../" - главная страница (приветствие "Welcome to the ATM!")
- ".../b2b" - адрес для отправки GET-запросов обмена


