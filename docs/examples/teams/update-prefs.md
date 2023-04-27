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
    client.SetJWT("") // Your secret JSON Web Token

    var service := appwrite.Teams{
        client: &client
    }

    var response, error := service.UpdatePrefs("[TEAM_ID]", nil)

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
