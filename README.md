# vk_challenge_rest
REST-API маркетплейса


## База данных

Использовал для реализации БД на постгресе:

![alt text](image.png)

По сущностям: есть таблица юзеров, ролей, их маппинг и таблица объявлений

## Модель

В качестве модели выбрал `MVC`. Итого, у меня есть слой сервисов в папке `service`, слой моделей `model`, контроллеры `controller`. Ручки я создавал через `http.HandleFunc` в `main.go`

## Авторизация

Пользователь в теле запроса передает json с нужными полями:
```
{
    "login":"Misha",
    "password":"qwerty"
}
```


Пароли и логины хранятся в бд в таком же виде в каком их передает пользователь. При попытке авторизации создается jwt токен (реализация в `jwt`) и возвращается в теле ответа в виде json:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE0NzU2MzQsInVzZXIiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAifQ.nTa2aUuFXrHBz1SBVqf5Y83FF7ovKRBZXivCYPBWWlo"
}
```

Если пользователь не был зарегестрирован - возвращаем в теле ответа ошибку 

## Регистрация

Если пользователь был зарегестрирован - возвращаем его в теле ответа. В бд ничего не меняем

Если пользователь новый, то добавляем его в табличку `users`, смотрим есть ли роль `USER` в таблице `role`. И если ее нет, то добавляем ее в таблицу. Затем в талицу `user_x_roles` добавляем маппинг юзера и роли. Последняя таблица нужна, если у пользователя будет несколько ролей. В итоге возвращаем добавленного пользователя

Добавлена транзакция на случай ошибки в добавление строк в таблицы `role` и `user_x_roles`. При появлении ошибки откатываем все изменения назад

Токен в данном случае не узнать, можно было писать в лог, но я не стал. Пользователю надо авторизоваться, после этого покажется токен

Пользователь
```
{
    "login":"third_new_new",
    "password":"qwerty"
}
```

Ответ:
```
{
    "id": "e69a4197-a0c2-48ed-9f77-445a29d31bb4",
    "login": "third_new_new",
    "password": "qwerty"
}
```

## Добавление объявлений

Из хедера запроса достаем значение токена. Проверяем его на валидность, если все ок, то добавляем объявление

Но, прежде чем добавить объявление проверяем валидность полей. В `constant` указаны ограничения на них

В ответе возвращаем добавленное объявление

Клиент:
```
{
    "title": "Example Ad",
    "content": "This is an example ad.",
    "image_path": "https://example.com/image.jpg",
    "price": 12
}
```
```
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE0NzY2OTUsInVzZXIiOiJkY2JhZDU3Yy1iNTU0LTQ4NGYtOGM4Mi02ZDFkYzIyMmU1MGUifQ.sQMXvDIG0_HQe3LY0vAGr9rEBMefdZhirGIz0Ea8N2M
```

Ответ:
```
{
    "title": "Example Ad",
    "content": "This is an example ad.",
    "image_path": "https://example.com/image.jpg",
    "price": 12
}
```

## Просмотр объявлений

Вначале получаем объявления, затем смотрим есть ли валидный пользователь, получаем его id и возвращаем результат в порядке возрастания даты

```
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE0NzY2OTUsInVzZXIiOiJkY2JhZDU3Yy1iNTU0LTQ4NGYtOGM4Mi02ZDFkYzIyMmU1MGUifQ.sQMXvDIG0_HQe3LY0vAGr9rEBMefdZhirGIz0Ea8N2M
```

Ответ:
```
[
    {
        "Adv": {
            "title": "Example Ad",
            "content": "This is an example ad.",
            "image_path": "https://example.com/image.jpg",
            "price": 12
        },
        "IsUserAdv": false
    },
    {
        "Adv": {
            "title": "Example ad another",
            "content": "This is an example ad.",
            "image_path": "https://example.com/img.jpg",
            "price": 14
        },
        "IsUserAdv": true
    }
]
```

## JWT токен

Использовал алгоритм хеширования `SigningMethodHS256`


## Клиентский путь:

1) В ручке `http://localhost:8080/register` зарегистрироваться с телом:
```
{
    "login":"third_new_new",
    "password":"qwerty"
}
```

2) В ручке `http://localhost:8080/login` авторизоваться и получить токен:
```
{
    "login":"third_new_new",
    "password":"qwerty"
}
```

Ответ:
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE0Nzk0NjksInVzZXIiOiJlNjlhNDE5Ny1hMGMyLTQ4ZWQtOWY3Ny00NDVhMjlkMzFiYjQifQ.2Ha3D2UKtJMoAivRh6aLVoe4zE43KJDXJ5sZshv0dP4"
}
```

3) С токеном в хедере по ключу `Authorization` в ручке: `http://localhost:8080/add/adv` добавить объявление:

```
{
    "title": "Example Ad",
    "content": "This is an example ad.",
    "image_path": "https://example.com/image.jpg",
    "price": 12
}
```

4) С тем же токеном в ручке `http://localhost:8080/show/adv` получить объявления