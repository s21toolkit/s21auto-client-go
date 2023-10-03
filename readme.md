# s21client 🍻🫃

[![.github/workflows/test.yaml](https://github.com/s21toolkit/s21client/actions/workflows/test.yaml/badge.svg)](https://github.com/s21toolkit/s21client/actions/workflows/test.yaml)

Клиент для внутреннего GQL API платформы edu.21-school.ru.

```sh
go get github.com/s21toolkit/s21client
```

Пример использования:

```golang
client := s21client.New(
  s21client.DefaultAuth(
    os.Getenv("S21_USERNAME"),
    os.Getenv("S21_PASSWORD")
  )
)

user, err := client.R().GetCurrentUser(requests.GetCurrentUser_Variables{})
if err != nil {
  t.Error(err)
}

fmt.Println(user)
```

## Генерация методов

Методы клиента генерируются автоматически на основе запросов платформы к бекенду.

Для генерации запросов используется [s21auto](https://github.com/s21toolkit/s21auto):

```sh
s21auto client generate log.har -o s21client/requests
```

> Если какие-то методы не нужны, из папки requests можно удалить всё кроме [`context.go`](/requests/context.go).
