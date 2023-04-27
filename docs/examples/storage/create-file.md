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

    var response, error := service.CreateFile("[BUCKET_ID]", "[FILE_ID]", NewInputFile("/path/to/file.png", "file.png"), ["read("any")"])

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
