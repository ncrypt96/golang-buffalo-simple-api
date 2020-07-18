package actions

import (
	"encoding/json"
	"strings"

	"net/http"
	"simple_api/db"
	"simple_api/requestshapes"
	"simple_api/responsecodes"

	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
)

// UAdd default implementation.
func UAdd(c buffalo.Context) error {
	// get a reference to the response struct
	var u requestshapes.UserAdd
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	//If the parameters are only empty spaces
	tName := strings.TrimSpace(u.Name)
	tQuote := strings.TrimSpace(u.Quote)
	if tName == "" || tQuote == "" {
		log.Info("The request did not contain all necessary keys")
		return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseOnAddMissing))
	}
	{
		err := addUserToDB("users", tName, tQuote)
		if err != nil {
			log.Error(err)
			return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseDatabase))
		}
	}

	return c.Render(http.StatusOK, r.JSON(&responsecodes.SuccessResponseOnAdd))
}

//addUserToDB addes user to database
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
		if err != nil {
			// If there is an issue with the database
			return c.Render(http.StatusBadRequest, r.JSON(&responsecodes.ErrResponseDatabase))
		} else if len(val) <= 0 {
			return c.Render(http.StatusConflict, r.JSON(&responsecodes.ErrResponseOnGetNotExist))
		} else {
			return c.Render(http.StatusOK, r.JSON(responsecodes.SuccessResponseOnGet(val)))
		}
	}
}

//getQuote gets the quote from the database for the given name
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
