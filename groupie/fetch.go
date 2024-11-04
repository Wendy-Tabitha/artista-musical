package groupie

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
)

// ApiCache stores the fetched data and uses a mutex to handle concurrency.
var ApiCache Cache

// Cache struct with mutex to handle concurrency.
type Cache struct {
    Artists   []Artist
    Dates     Dates
    Locations Location
    Relations Relations
    sync.Mutex
}

// FetchArtists fetches the artists data from the API.
func FetchArtists(wg *sync.WaitGroup) {
    defer wg.Done()

    ApiCache.Lock()
    defer ApiCache.Unlock()

    if len(ApiCache.Artists) > 0 {
        return // Already fetched
    }

    url := "https://groupietrackers.herokuapp.com/api/artists"
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        var artists []Artist
        if err := json.NewDecoder(resp.Body).Decode(&artists); err == nil {
            ApiCache.Artists = artists
        } else {
            fmt.Println("Error decoding artists data:", err)
        }
    } else {
        fmt.Println("Error fetching artists:", err)
    }
}

// FetchDates fetches the dates data from the API.
func FetchDates(wg *sync.WaitGroup) {
    defer wg.Done()

    ApiCache.Lock()
    defer ApiCache.Unlock()

    if len(ApiCache.Dates.Index) > 0 {
        return // Already fetched
    }

    url := "https://groupietrackers.herokuapp.com/api/dates"
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        var dates Dates
        if err := json.NewDecoder(resp.Body).Decode(&dates); err == nil {
            ApiCache.Dates = dates
        } else {
            fmt.Println("Error decoding dates data:", err)
        }
    } else {
        fmt.Println("Error fetching dates:", err)
    }
}

// FetchLocations fetches the locations data from the API.
func FetchLocations(wg *sync.WaitGroup) {
    defer wg.Done()

    ApiCache.Lock()
    defer ApiCache.Unlock()

    if len(ApiCache.Locations.Index) > 0 {
        return // Already fetched
    }

    url := "https://groupietrackers.herokuapp.com/api/locations"
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        var locations Location
        if err := json.NewDecoder(resp.Body).Decode(&locations); err == nil {
            ApiCache.Locations = locations
        } else {
            fmt.Println("Error decoding locations data:", err)
        }
    } else {
        fmt.Println("Error fetching locations:", err)
    }
}

// FetchRelations fetches the relations data from the API.
func FetchRelations(wg *sync.WaitGroup) {
    defer wg.Done()

    ApiCache.Lock()
    defer ApiCache.Unlock()

    if len(ApiCache.Relations.Index) > 0 {
        return // Already fetched
    }

    url := "https://groupietrackers.herokuapp.com/api/relation"
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        var relations Relations
        if err := json.NewDecoder(resp.Body).Decode(&relations); err == nil {
            ApiCache.Relations = relations
        } else {
            fmt.Println("Error decoding relations data:", err)
        }
    } else {
        fmt.Println("Error fetching relations:", err)
    }
}
