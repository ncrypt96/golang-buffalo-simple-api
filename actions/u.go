package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// USignup default implementation.
func USignup(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("u/signup.html"))
}

