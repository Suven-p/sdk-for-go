package appwrite

import "fmt"

type Permission struct {}

func NewPermission() *Permission {
	return &Permission{}
}

func (p *Permission) Read(role string) string {
	return fmt.Sprintf("read(\"%s\")", role)
}

func (p *Permission) Write(role string) string {
	return fmt.Sprintf("write(\"%s\")", role)
}

func (p *Permission) Create(role string) string {
	return fmt.Sprintf("create(\"%s\")", role)
}

func (p *Permission) Update(role string) string {
	return fmt.Sprintf("update(\"%s\")", role)
}

func (p *Permission) Delete(role string) string {
	return fmt.Sprintf("delete(\"%s\")", role)
}
