# News API

This tiny API provides subscribing and unsubscribing of users using the [go-news](https://github.com/tj/go-news) package. 

## Routes

### GET /subscribers

Responds with all subscribers for a `newsletter`, provided in the query-string parameter.

### POST /subscribe

Accepts a form body with a `newsletter` name and subscriber and `email`.

### GET /unsubscribe

Requires the `newsletter` and `token` query-string parameters. The token must be signed with the __TOKEN_SECRET__ and the included [token](./token) package, this ensures that only people who received an email with this token in their unsubscribe link can perform an unsubscribe.

## Setup

- Create a DynamoDB table with a __Partition Key__ of "newsletter", and a __Sort Key__ of "email".
- Deploy this app however you prefer, you could use [Up](https://github.com/apex/up) for example
- Define the required environment variables:
    - __TOKEN_SECRET__: The secret used for signing unsubscribe tokens
    - __API_TOKEN__: The API token used to secure sensitive routes (`GET /subscribers`)
    - __SUBSCRIBE_REDIRECT_URL__: The redirect URL used for a successful subscription (thank you page)
    - __UNSUBSCRIBE_REDIRECT_URL__: The redirect URL used for a successful unsubscribe (bye bye page)
    - __DYNAMO_TABLE__: An optional table name, defaults to "news"

## Auth

The __API_TOKEN__ is used to restrict access to sensitive routes. You can specify it as the Basic Auth password, no username is necessary, for example:

```
$ export API_TOKEN=hello
$ go run main.go &
$ curl -u :slothy localhost:3000/subscribers?newsletter=blog
```

---

[![GoDoc](https://godoc.org/github.com/tj/news-api?status.svg)](https://godoc.org/github.com/tj/news-api)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

## Sponsors

This project is sponsored by my [GitHub sponsors](https://github.com/sponsors/tj):

[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/0" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/0)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/1" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/1)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/2" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/2)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/3" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/3)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/4" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/4)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/5" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/5)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/6" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/6)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/7" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/7)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/8" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/8)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/9" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/9)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/10" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/10)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/11" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/11)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/12" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/12)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/13" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/13)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/14" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/14)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/15" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/15)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/16" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/16)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/17" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/17)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/18" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/18)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/19" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/19)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/20" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/20)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/21" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/21)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/22" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/22)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/23" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/23)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/24" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/24)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/25" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/25)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/26" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/26)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/27" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/27)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/28" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/28)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/29" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/29)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/30" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/30)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/31" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/31)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/32" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/32)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/33" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/33)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/34" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/34)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/35" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/35)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/36" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/36)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/37" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/37)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/38" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/38)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/39" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/39)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/40" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/40)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/41" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/41)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/42" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/42)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/43" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/43)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/44" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/44)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/45" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/45)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/46" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/46)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/47" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/47)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/48" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/48)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/49" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/49)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/50" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/50)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/51" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/51)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/52" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/52)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/53" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/53)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/54" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/54)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/55" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/55)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/56" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/56)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/57" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/57)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/58" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/58)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/59" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/59)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/60" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/60)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/61" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/61)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/62" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/62)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/63" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/63)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/64" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/64)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/65" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/65)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/66" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/66)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/67" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/67)
[<img src="https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/avatar/68" width="35">](https://sponsors-api-u2fftug6kq-uc.a.run.app/sponsor/profile/68)

