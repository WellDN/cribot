package common

type DBUser struct {  
    ID int
    Name string
    Password string
}
// lil query
const userQ = "SELECT id, name FROM user"
