package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type User struct {
	FirstName string
	LastName  string
}

// USignup default implementation.
func USignup(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(&User{"Nithin", "Stephan"}))
}
