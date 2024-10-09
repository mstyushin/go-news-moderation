GoNews Moderation Service
=========================
Сервис пре-модерации для комментариев.

---

Принимает на модерацию текст комментария и никнейм автора в теле POST запроса, в случае успешной валидации отвечает HTTP 200, в противном случае HTTP 400.

Пример объёкта на валидацию:

        {
            "author": "Alice",
            "text": "Nice one, keep it up!"
        }

# Требования

-   docker >=23.0.0

-   golang 1.22

# Начало работы

Приложение поддерживает конфиг-файлы в формате yaml:

        $ make build
        $ ./bin/go-news-moderation -print-config > config.yaml
        $ ./bin/go-news-moderation -config ./config.yaml

---

Для быстрого запуска с дефолтным конфигом:

        $ make run

Логи будут писаться сюда:

        $ tail -f log/go-news-moderation.log

Остановить приложение:

        $ make clean

## Примеры запросов

Валидный запрос:

        $ curl -v http://127.0.0.1:8083/moderation -d '{"author":"Jane", "text":"Totally legal comment"}'

Невалидный запрос (если использовался блэклист слов из конфига по-умолчанию):

        $ curl -v http://127.0.0.1:8083/moderation -d '{"author":"Bob", "text":"Absolutely illegal asdfg comment"}'

# Тесты
 
Полный прогон имеющихся тестов:

        $ make test

## Docker
TBD
