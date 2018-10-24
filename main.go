package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
	flags "github.com/jessevdk/go-flags"
)

var strurl = "https://api.twitter.com/1.1/statuses/update.json"
var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

type Options struct {
	Longs bool `short:"l" long:"long" description:"Long text"`
}

var opts Options

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
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	var tweet string
	if opts.Longs {
		fmt.Println("Please Long Text.")
		tweet = long()
	} else {
		tweet = short()
	}

	httpClient := http.DefaultClient
	credentials := &oauth.Credentials{
		Token:  access_token,
		Secret: access_token_secret,
	}

	oauthClient.Credentials.Token = consumer_key
	oauthClient.Credentials.Secret = consumer_secret

	err = HttpPost(httpClient, credentials, tweet)
	if err != nil {
		fmt.Printf("Error")
	}
}

func short() string {
	return os.Args[1]
}

func long() string {
	return os.Args[1]
}
