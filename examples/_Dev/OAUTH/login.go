package oauthexample

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/context"


	"github.com/nu7hatch/gouuid"
)

func init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/github-login", handleGithubLogin)
	http.HandleFunc("/oauth2callback", handleOauth2Callback)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, `<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <a href="/github-login">LOGIN WITH GITHUB</a>
  </body>
</html>`)
}

func handleGithubLogin(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// get the session
	id, _ := uuid.NewV4()

	redirectURI := "http://localhost:8080/oauth2callback"

	values := make(url.Values)
	values.Add("client_id", "0ccd33716940f347065e")
	values.Add("redirect_uri", redirectURI)
	values.Add("scope", "user:email")
	values.Add("state", id.String())

	// save the session
	session.State = id.String()
	putSession(ctx, res, session)

	http.Redirect(res, req, fmt.Sprintf(
		"https://github.com/login/oauth/authorize?%s",
		values.Encode(),
	), 302)
}
