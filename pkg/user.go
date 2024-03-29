package pkg

type User struct {
    id      int
    name    string
    password string
}

func NewUser(id int, name, password string) *User {
    return &User {
        id: id,
        name: name,
        password: password,
    }       
}
