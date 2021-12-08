package utils

import (
	"bq_init/config"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

/*
	Connects to BigQuery
*/
func ConnectToBigQuery() (context.Context, *bigquery.Client) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatal("Could not connect to BigQuery.")
	}

	fmt.Printf("Client datatype: %T\n", client)
	defer client.Close()

	return ctx, client
}

/*
	Runs a query in BigQuery
*/
func RunQuery(ctx context.Context, client *bigquery.Client, query_text string) (*bigquery.RowIterator, error) {

	query := client.Query(
		query_text)
	query.Location = "US"

	// Run the query and print results when the query job is completed.
	job, err := query.Run(ctx)
	if err != nil {
		log.Println("Job failed:", err)
		return nil, err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Println("Problem while waiting for the job", err)
		return nil, err
	}
	if err := status.Err(); err != nil {
		log.Println("Deu xabugo no status:", err)
		return nil, err
	}

	// Job.Read returns a bigQuery.RowIterator
	it, err := job.Read(ctx)

	return it, nil
}
