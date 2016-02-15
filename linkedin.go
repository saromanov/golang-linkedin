package linkedin

import (
   "encoding/json"
   "golang.org/x/oauth2"
   ln "golang.org/x/oauth2/linkedin"
   "io/ioutil"
   "net/http"
)


type Linkedin struct {

}

// Create of the new access for linkedin
func New(appId, appSecret, redirect string, scopes []string) {
	lnConfig := &oauth2.Config {
		ClientID: appId,
		ClientSecret: appSecret,
		RedirectURL: redirect,
		Scopes: scopes,
		Endpoint: ln.Endpoint
	}
}