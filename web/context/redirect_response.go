package context

import (
	"fmt"
	"net/http"
)

// RedirectResponse used to send a redirect response
type RedirectResponse struct {
	*Response
	location string
}

// SetLocation sets the redirect location
func (redirectResponse *RedirectResponse) SetLocation(location string) *RedirectResponse {
	redirectResponse.location = location
	return redirectResponse
}

// Send sends the response
func (redirectResponse *RedirectResponse) Send(w http.ResponseWriter, r *http.Request) {
	if redirectResponse.statusCode < http.StatusMultipleChoices || redirectResponse.statusCode > http.StatusPermanentRedirect {
		panic(fmt.Errorf("can not redirect with HTTP status code `%d`", redirectResponse.statusCode))
	}
	http.Redirect(w, r, redirectResponse.location, redirectResponse.statusCode)
}
