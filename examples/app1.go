package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"code.google.com/p/google-api-go-client/bigquery/v2"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

var (
	serviceEmail = flag.String("se", "", "the-service-email-of-authorization")
	pkLocation   = flag.String("pkl", "", "the-location-of-you-primary-key")
	projectID    = flag.String("pid", "", "the-project-id")
)

func main() {
	flag.Parse()

	if *serviceEmail == "" || *pkLocation == "" || *projectID == "" {
		flag.PrintDefaults()
		return
	}

	service, err := bigquery.New(client())

	if err != nil {
		log.Printf("could not create big query service %#v", err)
		panic(err)
	}

	query := `
	SELECT
		count(*)
	FROM
		TABLE_QUERY(githubarchive:year, 'true')
	`

	queryRequest := &bigquery.QueryRequest{
		Query: query,
		Kind:  "igquery#queryRequest",
	}

	jobsService := bigquery.NewJobsService(service)
	jobsQueryCall := jobsService.Query(*projectID, queryRequest)
	queryResponse, err := jobsQueryCall.Do()

	if err != nil {
		log.Printf("Could not call query job: %#v", err)
		panic(err)
	}

	log.Printf("Job completed?: %#v", queryResponse.JobComplete)
	log.Printf("Total rows: %d", queryResponse.TotalRows)
	log.Printf(queryResponse.Rows[0].F[0].V.(string))
}

func client() *http.Client {
	ctx := context.Background()

	pkBytes, err := ioutil.ReadFile(*pkLocation)

	if err != nil {
		panic(err)
	}

	conf := &jwt.Config{
		Email:      *serviceEmail,
		PrivateKey: pkBytes,
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
		},
		TokenURL: google.JWTTokenURL,
	}

	return conf.Client(ctx)
}
