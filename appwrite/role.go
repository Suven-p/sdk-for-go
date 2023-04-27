package appwrite

import (
    "fmt"
)

type Role struct {}

func NewRole() *Role {
    return &Role{}
}

func (r *Role) Any() string {
    return "any"
}

func (r *Role) User(id, status string) string {
    if status != "" {
        return fmt.Sprintf("user:%s/%s", id, status)
    }
    return fmt.Sprintf("user:%s", id)
}

func (r *Role) Users(status string) string{
    if status != "" {
        return fmt.Sprintf("users/%s", status)
    }
    return "users"
}

func (r *Role) Guests() string {
    return "guests"
}

func (r *Role) Team(id, role string) string {
    if role != "" {
        return fmt.Sprintf("team:%s/%s", id, role)
    }
    return fmt.Sprintf("team:%s", id)
}

func (r *Role) Member(id string) string {
    return fmt.Sprintf("member:%s", id)
}