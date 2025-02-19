package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/storage"
)

func main() {
    client := client.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    service := storage.NewStorage(client)
    response, error := service.ListBuckets(
        storage.WithListBucketsQueries([]interface{}{}),
        storage.WithListBucketsSearch("[SEARCH]"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
