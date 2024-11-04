package main

import (
	"fmt"
	"sync"

	"groupie/groupie" // replace with actual package name if different
)

func main() {
	// Start Go routines to fetch data from each endpoint.
	var wg sync.WaitGroup
	wg.Add(4) // Four endpoints to fetch

	go groupie.FetchArtists(&wg)
	go groupie.FetchDates(&wg)
	go groupie.FetchLocations(&wg)
	go groupie.FetchRelations(&wg)

	wg.Wait() // Wait for all Go routines to complete

	// Serve data from cache without redundant API calls.
	fmt.Printf("Cached Artists: %v\n", groupie.ApiCache.Artists)
	fmt.Printf("Cached Artists: %v\n", groupie.ApiCache.Dates)
	fmt.Printf("Cached Artists: %v\n", groupie.ApiCache.Locations)
	fmt.Printf("Cached Artists: %v\n", groupie.ApiCache.Relations)


}
