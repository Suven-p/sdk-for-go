package main

import (
    "fmt"
    "time"
    "github.com/appwrite/sdk-for-go"
)

func main() {
    client := appwrite.NewClient()

    client.SetEndpoint("https://[HOSTNAME_OR_IP]/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    var service := appwrite.Storage{
        client: &client
    }

    var response, error := service.GetFilePreview("[BUCKET_ID]", "[FILE_ID]", 0, 0, "center", 0, 0, "", 0, 0, -360, "", "jpg")

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
