package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/sourcegraph/conc/pool"
)

// simulateAPI simulates an API call to fetch product prices
func simulateAPI(apiID int) int {
	// Simulate varying response times
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random duration
	randomDuration := time.Duration(rng.Int63n(int64(2 * time.Second)))

	time.Sleep(randomDuration)

	// Return a mock price
	return apiID*10 + rand.Intn(10)
}

func main() {
	// Create a result pool
	pool := pool.NewWithResults[int]().WithMaxGoroutines(5)

	// Simulate fetching prices
	for apiID := 1; apiID <= 10; apiID++ {
		apiID := apiID // Capture variable for goroutine
		pool.Go(func() int {
			fmt.Printf("Fetching price from API %d...\n", apiID)
			return simulateAPI(apiID)
		})
	}

	// Wait for all API calls to finish
	prices := pool.Wait()

	// Display the collected prices
	fmt.Println("Collected Prices:", prices)
}
