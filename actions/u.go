package actions

import (
	"encoding/json"
	"strings"

	"net/http"
	"simple_api/db"
	"simple_api/responsecodes"
	"simple_api/shapes"

	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
)

// UAdd default implementation.
func UAdd(c buffalo.Context) error {
	var u shapes.User
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Quote) == "" {
		log.Info("The request did not contain all necessary keys")
		return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseOnAddMissing))
	}
	{
		err := addUserToDB("users", u.Name, u.Quote)
		if err != nil {
			log.Error(err)
			return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseDatabase))
		}
	}

	return c.Render(http.StatusOK, r.JSON(&responsecodes.SuccessResponseOnAdd))
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

// UGet default implementation.
func UGet(c buffalo.Context) error {
	n := c.Request().URL.Query().Get("name")
	if len(n) <= 0 {
		return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseOnGetInvalidParameter))
	}
	{
		val, err := getQuote("users", n)
		log.Info("iiiiiiiiii", val, "iiiiiiiiiiiiii")
		if err != nil {
			return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseDatabase))
		} else if len(val) <= 0 {
			return c.Render(http.StatusConflict, r.JSON(&responsecodes.ErrResponseOnGetNotExist))
		} else {
			return c.Render(http.StatusOK, r.JSON(responsecodes.SuccessResponseOnGet(val)))
		}
	}
}

func getQuote(bucketName, key string) (string, error) {

	dbRef, err := db.OpenDB("./bolt.db")

	if err != nil {
		return "", err
	}

	defer dbRef.Close()

	v, err := db.GetData(dbRef, bucketName, key)

	if err != nil {
		return "", err
	}

	return v, err
}
