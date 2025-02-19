package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/databases"
)

func main() {
    client := client.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    service := databases.NewDatabases(client)
    response, error := service.UpdateDocument(
        "[DATABASE_ID]",
        "[COLLECTION_ID]",
        "[DOCUMENT_ID]",
        databases.WithUpdateDocumentData(map[string]interface{}{}),
        databases.WithUpdateDocumentPermissions(interface{}{"read("any")"}),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
