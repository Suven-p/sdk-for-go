package appwrite

type ID struct {}

func NewID() *ID {
    return &ID{}
}

func (i *ID) Custom(id string) string {
    return id
}

func (id *ID) Unique() string {
    return "unique()"
}