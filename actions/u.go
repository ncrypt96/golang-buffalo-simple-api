package actions

import (
	"encoding/json"

	"net/http"
	"simple_api/db"
	"simple_api/shapes"

	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
)

// Error Response
var errResponse = shapes.ErrorResponse{
	Error: &shapes.Error{400, "Eithier Name or Quote is missing"},
}

var successResponse = shapes.SuccessResponse{
	Data: &shapes.Data{"The provided name has been successfully added to the database"},
}

// UAdd default implementation.
func UAdd(c buffalo.Context) error {
	var u shapes.User
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if u.Name == "" || u.Quote == "" {
		log.Info("The request did not contain all necessary keys")
		return c.Render(http.StatusBadRequest, r.JSON(&errResponse))
	}
	{
		err := addUserToDB("users", u.Name, u.Quote)
		if err != nil {
			log.Error(err)
		}
	}

	return c.Render(http.StatusOK, r.JSON(&successResponse))
}

func addUserToDB(bucketName, name, quote string) error {
	dbRef, err := db.OpenDB("./bolt.db")
	if err != nil {
		log.Error(err)
		return err
	}

	defer dbRef.Close()

	err = db.PutData(dbRef, "users", name, quote)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Successfully entered name: " + name + " quote: " + quote + " in the database")
	return nil
}
