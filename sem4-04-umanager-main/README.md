# 03-04-umanager

## Описание
В этом задании вам нужно подвести итог нашего приложения. К сожалению мы не можем рассмотреть множество тем, которые 
могли бы сделать наш проект более продакшен реди. Мы сосредоточились на сетевой связности и некоторых асинхронных 
функциях.

Что мы не можем разобрать в наших дз и семинарах. Это тему авторизации, нормального использования observability и DDD.
Вы можете самостоятельно разобрать эти темы и, например добавить такие вещи как хэширование паролей через bcrypt и 
организацию сессий. Сессии можно хранить в redis. Можно пойти по пути генерации jwt токенов.

* Во-первых, вам нужно залогировать ошибки там где это уместно. Сделать это нужно в http хэндлерах api gw. В grpc это 
делать не нужно. 
* Во-вторых Вам нужно продемонстрировать ваше умение писать тесты. Нужно написать несколько тестов

## Про тесты
Вы можете сделать это несколькими путями. Вы можете
написать тесты на хэндлеры в api gw с помощью httptest. Достаточно написать несколько тестов на POST, DELETE
операции. Вам может понадобится использовать mockgen для мокания grpc client. Вместо mockgen также можете
использовать собственный stub объект. Для этого вам нужно написать собственную реализацию интерфейсов grpc
клиентов.В реализации можете возвращать например фейковые данные. 

