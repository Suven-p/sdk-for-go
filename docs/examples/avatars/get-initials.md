package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/avatars"
)

func main() {
    client := client.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    service := avatars.NewAvatars(client)
    response, error := service.GetInitials(
        avatars.WithGetInitialsName("[NAME]"),
        avatars.WithGetInitialsWidth(0),
        avatars.WithGetInitialsHeight(0),
        avatars.WithGetInitialsBackground(""),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
