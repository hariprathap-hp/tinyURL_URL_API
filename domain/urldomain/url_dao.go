package urldomain

import (
	"fmt"
	"strings"
	"test3/hariprathap-hp/DesignTinyURL/tinyURL_URL_API/dataResources/postgres"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
)

const (
	checkDuplication = "duplicate key value violates"
	insertURLQuery   = "insert into url (hash,originalurl,creationdate,expirationdate,userid) values ($1,$2,$3,$4,$5)"
	deleteURLQuery   = "delete from url where hash=$1"
	listURLQuery     = "select originalurl, hash from url where userid=$1"
	redirectURLQuery = "select originalurl from url where hash=$1"
)

func (url *Url) Create() *errors.RestErr {
	//the function create is to insert the tinyURL along with other values into the DB
	query, err := postgres.Client.Prepare(insertURLQuery)
	if err != nil {
		return errors.NewInternalServerError("DB Error: statement preparation failed")
	}
	defer query.Close()
	if _, insertErr := query.Exec(url.TinyURL, url.OriginalURL, url.CreationDate, url.ExpirationDate, url.UserEmail); insertErr != nil {
		if strings.Contains(insertErr.Error(), checkDuplication) {
			return errors.NewInternalServerError(fmt.Sprintf("DB error : User %s already present in the database", url.UserEmail))
		}
		return errors.NewInternalServerError("DB Error: error while inserting the record to the database")
	}
	return nil
}

func (url *Url) Delete() *errors.RestErr {
	//the function delete is to delete a tinyURL from the DB
	return nil
}

func (url *Url) Redirect() (*string, *errors.RestErr) {
	//the function redirect is to fetch the original url for a tinyurl so that it can be redirected by our controller layer
	return nil, nil
}

func (url *Url) List() (UrlsList, *errors.RestErr) {
	//the function list gets all the tinyURLs created for any particular user
	return nil, nil
}
