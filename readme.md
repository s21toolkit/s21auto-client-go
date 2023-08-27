# s21client 🍻🫃

[![.github/workflows/test.yaml](https://github.com/s21toolkit/s21client/actions/workflows/test.yaml/badge.svg)](https://github.com/s21toolkit/s21client/actions/workflows/test.yaml)

Клиент для внутреннего GQL API платформы edu.21-school.ru.

```golang
client := s21client.New(s21client.DefaultAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))

user, err := client.R().GetCurrentUser(requests.Variables_GetCurrentUser{})
if err != nil {
  t.Error(err)
}

fmt.Println(user)
```
