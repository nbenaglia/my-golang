package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

func main() {
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Use the Olivere package to get the Elasticsearch version number
	fmt.Println("Version:", elastic.Version)

	// Create a context object for the API calls
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("https://es.sand.cuebiq.ai:443"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		fmt.Println("elastic. NewClient() ERROR:", err)
		log.Fatalf("quiting connection..")
	} else {
		// Print client information
		fmt.Println("client:", client)
		fmt.Println("client TYPE:", reflect.TypeOf(client))
	}
	// Declare index name as string
	indexName := "nbenaglia_index"

	// Declare a slice of strings for the index names to be checked
	indices := []string{indexName}

	// Use the append() function to add more index names to slice
	indices = append(indices, indexName)
	indices = append(indices, "second_index")
	indices = append(indices, "last_index")

	// Instantiate a new *elastic. IndicesExistsService
	existService := elastic.NewIndicesExistsService(client)
	fmt.Println("existService TYPE:", reflect.TypeOf(existService))

	// Iterate over the IndicesExistsService object's methods
	t := reflect.TypeOf(existService)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println("nexist METHOD NAME:", i, method.Name)
		fmt.Println("method:", method)
	}

	// Iterate over the slice of Elasticsearch documents
	for _, index := range indices {
		indexSlice := []string{index}
		// Pass the string slice with the index name to the Index() method
		existService.Index(indexSlice)

		// Bool option for checking just local node or master node as well
		existService.Local(true)

		// Have Do() return an API response by passing the Context object to the method call
		exist, err := existService.Do(ctx)

		// Check if the IndicesExistsService. Do() method returned any errors
		if err != nil {
			log.Fatalf("IndicesExistsService. Do() ERROR:", err)

			// Print the index exists boolean response
		} else {
			fmt.Println("nIndicesExistsService. Do():", index, "exists:", exist)
		}
	}
}
