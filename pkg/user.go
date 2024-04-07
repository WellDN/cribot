package pkg

func NewUser(id int, name, password string) *User {
    return &User {
        id: id,
        name: name,
        password: password,
    }       
}

type User struct {
    id      int
    name    string
    password string
}

func GetUserById(u *User) int {
    return u.id
}

func GetUserByName(u *User) string {
    return u.name
}
