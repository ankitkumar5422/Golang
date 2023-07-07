package main

import (
	"fmt"
	"sync"

	"github.com/aerospike/aerospike-client-go"
)

func main() {
	// Establish Aerospike client connection
	client, err := aerospike.NewClient("localhost", 3000)
	if err != nil {
		fmt.Println("Error connecting to Aerospike:", err)
		return
	}
	defer client.Close()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of sets to insert
	numSets := 1000

	// Add the number of goroutines to the wait group
	wg.Add(numSets)

	for i := 0; i < numSets; i++ {
		go func(index int) {
			defer wg.Done()

			// Generate a unique key for the set
			key, err := aerospike.NewKey("namespace", "set", index)
			if err != nil {
				fmt.Println("Error creating key:", err)
				return
			}

			// Create a slice of aerospike.Bin values
			bins := []*aerospike.Bin{
				{
					Name:  "bin",
					Value: aerospike.NewValue(fmt.Sprintf("value%d", index)),
				},
			}

			// Write the set to Aerospike
			err = client.PutBins(nil, key, bins...)
			if err != nil {
				fmt.Println("Error inserting set:", err)
				return
			}

			fmt.Println("Set inserted:", index)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All sets inserted successfully")
}
