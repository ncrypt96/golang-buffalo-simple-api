package actions

import (
	"encoding/json"
	"net/http"
	"simple_api/shapes"

	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
)

// USignup default implementation.
func USignup(c buffalo.Context) error {
	// declare a reference to struct
	var s *shapes.SignUpData
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}

	log.Info(*s)

	return c.Render(http.StatusOK, r.JSON(&s))
}
