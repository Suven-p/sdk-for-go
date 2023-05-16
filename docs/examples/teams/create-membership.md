package main

import (
    "fmt"
    "github.com/appwrite/sdk-for-go/client"
    "github.com/appwrite/sdk-for-go/teams"
)

func main() {
    client := client.NewClient()

    client.SetEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    service := teams.NewTeams(client)
    response, error := service.CreateMembership(
        "[TEAM_ID]",
        []interface{}{},
        "https://example.com",
        teams.WithCreateMembershipEmail("email@example.com"),
        teams.WithCreateMembershipUserId("[USER_ID]"),
        teams.WithCreateMembershipPhone("+12065550100"),
        teams.WithCreateMembershipName("[NAME]"),
    )

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
