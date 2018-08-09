package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
)

var strurl = "https://api.twitter.com/1.1/statuses/update.json"
var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func HttpPost(client *http.Client, cred *oauth.Credentials, str string) error {
	from := url.Values{}
	from.Set("status", str+"\nFROM CLI")

	resp, err := oauthClient.Post(client, cred, strurl, from)
	if err != nil {
		return err
	}
	resp.Body.Close()

	fmt.Printf("received %s\n", resp)

	return nil
}

func main() {
	tweet := os.Args[1]
	httpClient := http.DefaultClient
	credentials := &oauth.Credentials{
		Token:  access_token,
		Secret: access_token_secret,
	}

	oauthClient.Credentials.Token = consumer_key
	oauthClient.Credentials.Secret = consumer_secret

	err := HttpPost(httpClient, credentials, tweet)
	if err != nil {
		fmt.Printf("Error")
	}
}
