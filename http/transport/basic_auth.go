package transport

import "net/http"

type basicAuth struct {
	http.RoundTripper
	username string
	password string
}

func (b *basicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(b.username, b.password)
	return b.RoundTripper.RoundTrip(req)
}

func BasicAuth(parent http.RoundTripper, username, password string) http.RoundTripper {
	return &basicAuth{
		RoundTripper: parent,
		username:     username,
		password:     password,
	}
}
