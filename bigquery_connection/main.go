package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// TODO: Add dataset variables to the config file (which uses struct)
const projectID = "YOUR_PROJECT_ID"
const datasetID = "YOUR_DATASET_ID"
const tableID = "YOUR_TABLE_NAME"

//tablePath := projectID + "." + datasetID + "." + tableID

func main() {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Println("BigQuery Client deu xabugo:", err)
		//return fmt.Errorf("bigquery.NewClient: %v", err)
	}

	fmt.Printf("Client datatype: %T\n", client)
	defer client.Close()

	q := client.Query(
		"SELECT * FROM `mlopsplatform.model_catalog.model_catalog_table_t` " +
			"WHERE modelStatus = \"Unprocessed\" ")
	q.Location = "US"
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		log.Println("Job failed:", err)
		//return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Println("Problem while waiting for the job", err)
		//return err
	}
	if err := status.Err(); err != nil {
		log.Println("Deu xabugo no status:", err)
		//return err
	}

	// Job.Read returns a bigQuery.RowIterator
	it, err := job.Read(ctx)

	// https://github.com/googleapis/google-cloud-go/blob/f91a0c3581e771b8ccc0fd801a0893636dda3c5f/bigquery/iterator.go#L43
	log.Println("Total number of rows:", it.TotalRows)
	//log.Println("Table Schema:", *it.Schema)

	if it.TotalRows == 0 {
		log.Println("There wasn't any rows retrieved.")
		//http.Error(rw, "expecting multipart file", http.StatusBadRequest)
		//fmt.Errorf("")
	} else {
		for {
			var row []bigquery.Value
			err := it.Next(&row)
			if err == iterator.Done {
				log.Println("Read all rows.")
				break
			}
			if err != nil {
				log.Println("Error while reading table's rows:", err)
				//return err
			}
			log.Println(&row)
		}
	}

}
