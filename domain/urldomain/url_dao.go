package urldomain

import (
	"fmt"
	"strings"

	"github.com/hariprathap-hp/tinyURL_URL_API/dataResources/postgres"
	"github.com/hariprathap-hp/utils_repo/errors"
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
		fmt.Println(err)
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
	query, err := postgres.Client.Prepare(deleteURLQuery)
	if err != nil {
		return errors.NewInternalServerError("DB Error: statement preparation failed")
	}
	defer query.Close()
	if _, deleteErr := query.Exec(url.TinyURL); deleteErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to delete the tiny url %s", url.TinyURL))
	}
	return nil
}

func (url *Url) Redirect() (*string, *errors.RestErr) {
	//the function redirect is to fetch the original url for a tinyurl so that it can be redirected by our controller layer
	query, err := postgres.Client.Prepare(redirectURLQuery)
	if err != nil {
		return nil, errors.NewInternalServerError("DB Error: statement preparation failed")
	}
	defer query.Close()
	var orig_url string
	scanErr := query.QueryRow(url.TinyURL).Scan(&orig_url)
	if scanErr != nil {
		return nil, errors.NewInternalServerError("DB Error: error while trying to query database")
	}
	return &orig_url, nil
}

func (url *Url) List() (UrlsList, *errors.RestErr) {
	//the function list gets all the tinyURLs created for any particular user
	query, err := postgres.Client.Prepare(listURLQuery)
	if err != nil {
		return nil, errors.NewInternalServerError("DB Error: statement preparation failed")
	}
	defer query.Close()
	rows, listErr := query.Query(url.UserEmail)
	if listErr != nil {
		return nil, errors.NewInternalServerError("DB Error: error while trying to query database")
	}
	result := make([]ListURLs, 0)
	for rows.Next() {
		var res ListURLs
		if scanErr := rows.Scan(&res.OriginalURL, &res.TinyURL); scanErr != nil {
			return nil, errors.NewInternalServerError("DB Error: error while scanning the query result")
		}
		result = append(result, res)
	}
	return result, nil
}
