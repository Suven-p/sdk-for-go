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
    response, error := service.CreateRelationshipAttribute(
        "[DATABASE_ID]",
        "[COLLECTION_ID]",
        "[RELATED_COLLECTION_ID]",
        "oneToOne",
        databases.WithCreateRelationshipAttributeTwoWay(false),
        databases.WithCreateRelationshipAttributeKey(""),
        databases.WithCreateRelationshipAttributeTwoWayKey(""),
        databases.WithCreateRelationshipAttributeOnDelete("cascade"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
