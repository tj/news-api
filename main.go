package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tj/go/env"
	"github.com/tj/go/http/response"

	"github.com/tj/go-news"
	"github.com/tj/news-api/token"
)

// secret for token signing.
var secret = env.Get("TOKEN_SECRET")

// apiToken for authentication.
var apiToken = env.Get("API_TOKEN")

// subscribeURL is the redirect URL used after a subscribe.
var subscribeURL = env.Get("SUBSCRIBE_REDIRECT_URL")

// unsubscribeURL is the redirect URL used after an unsubscribe.
var unsubscribeURL = env.Get("UNSUBSCRIBE_REDIRECT_URL")

// newsletters storage.
var newsletters = news.New(env.GetDefault("DYNAMO_TABLE", "news"))

// port number to bind on.
var port = env.GetDefault("PORT", "3000")

func main() {
	http.HandleFunc("/subscribers", auth(subscribers))
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/unsubscribe", unsubscribe)
	http.HandleFunc("/_health", health)
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// auth middleware.
func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, pass, ok := r.BasicAuth()
		if !ok {
			response.Unauthorized(w, "Route requires api token")
			return
		}

		if pass != apiToken {
			response.Forbidden(w, "Invalid api token")
			return
		}

		next.ServeHTTP(w, r)
	}
}

// subscribers route.
func subscribers(w http.ResponseWriter, r *http.Request) {
	newsletter := r.URL.Query().Get("newsletter")

	if newsletter == "" {
		response.BadRequest(w, "The newsletter query-string parameter is required")
		return
	}

	emails, err := newsletters.GetSubscribers(newsletter)
	if err != nil {
		log.Printf("error fetching subscribers: %v\n", err)
		return
	}

	response.OK(w, emails)
}

// subscribe route.
func subscribe(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing form: %v", err)
		response.BadRequest(w, "Error parsing form values")
		return
	}

	newsletter := r.FormValue("newsletter")
	email := r.FormValue("email")

	err = newsletters.AddSubscriber(newsletter, email)
	if err != nil {
		log.Printf("error subscribing email %q to %q: %v", email, newsletter, err)
		response.InternalServerError(w, "Error subscribing to newsletter")
		return
	}

	log.Printf("subscribed email %q to %q", email, newsletter)
	w.Header().Set("Location", subscribeURL)
	response.Found(w, "Redirecting to "+subscribeURL)
}

// unsubscribe route.
func unsubscribe(w http.ResponseWriter, r *http.Request) {
	newsletter := r.URL.Query().Get("newsletter")
	unsubscribeToken := r.URL.Query().Get("token")

	if newsletter == "" {
		response.BadRequest(w, "The newsletter query-string parameter is required")
		return
	}

	if unsubscribeToken == "" {
		response.BadRequest(w, "The token query-string parameter is required")
		return
	}

	email, ok := token.Unsign(secret, unsubscribeToken)
	if !ok {
		log.Printf("error unsigning %q", unsubscribeToken)
		response.BadRequest(w, "Invalid unsubscribe token")
		return
	}

	err := newsletters.RemoveSubscriber(newsletter, email)
	if err != nil {
		log.Printf("error unsubscribing %q: %v", email, err)
		response.InternalServerError(w, "Error unsubscribing from newsletter")
		return
	}

	log.Printf("unsubscribed email %q from %q", email, newsletter)
	w.Header().Set("Location", unsubscribeURL)
	response.Found(w, "Redirecting to "+unsubscribeURL)
}

// health route.
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, ":)")
}
