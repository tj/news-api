# News API

This tiny API provides subscribing and unsubscribing of users using the [go-news](https://github.com/tj/go-news) package. 

## Routes

### GET /subscribers

Responds with all subscribers for a `newsletter`, provided in the query-string parameter. This route requires the __API_TOKEN__.

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

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
