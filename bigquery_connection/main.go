package main

import (
	"fmt"
	"log"

	"bq_init/utils"

	"google.golang.org/api/iterator"
)

type TableSchema struct {
	FirstName string //`bigquery:"firstName"`
	LastName  string //`bigquery:"lastName"`
	Age       int    //`bigquery:"age"`
}

func main() {
	ctx, client := utils.ConnectToBigQuery()
	query_text := "SELECT * FROM `mlopsplatform.model_catalog.golang_table_test`"
	it, err := utils.RunQuery(ctx, client, query_text)

	if err != nil {
		log.Fatal(err)
	}
	// https://github.com/googleapis/google-cloud-go/blob/f91a0c3581e771b8ccc0fd801a0893636dda3c5f/bigquery/iterator.go#L43
	log.Println("Total number of rows:", it.TotalRows)
	//log.Println("Table Schema:", *it.Schema)

	if it.TotalRows == 0 {
		log.Println("There wasn't any rows retrieved.")
		//http.Error(rw, "expecting multipart file", http.StatusBadRequest)
		//fmt.Errorf("")
	} else {
		for {
			var ts TableSchema
			err := it.Next(&ts)
			if err == iterator.Done {
				log.Println("Read all rows.")
				break
			}
			if err != nil {
				log.Println("Error while reading table's rows:", err)
				//return err
			}
			// Acessing each row value based on the column's names
			fmt.Println(ts.FirstName)
			fmt.Println(ts.LastName)
			//namesSlice := storeNames(ts.FirstName)
		}
	}

	//log.Println("Names retrieved:", namesSlice)

}

func storeNames(bigQueryRow string) []string {
	var namesSlice []string
	namesSlice = append(namesSlice, bigQueryRow)
	return namesSlice
}
