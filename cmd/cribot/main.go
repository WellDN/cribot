package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

)

func main() {
    runCmd()
}
// TODO: Make a login context
// TODO: Make chat a server

func runCmd() {
    db, err := DbConnection()
    if err != nil {
        log.Fatal("An error occurred while connecting database", err)
    }
    err = createTableUser(db)
    if err != nil {
        log.Fatal("An error occurred while creating a table", err)
    }

    IRC(db)
    
    defer db.Close()
}
// Just testing how i would do it, but this doesn't make sense, yet
// You have to pass in the unique user here
func IRC(db *sql.DB) {
    in := bufio.NewScanner(os.Stdin)

    fmt.Println("\n\tChatty")
    fmt.Println(strings.Repeat("-", 25)+"\n")

    fmt.Println("Enter the user's name:")
    var name string
    _, err := fmt.Scanln(&name)
    if err != nil {
        log.Fatal(err)
    }

    for {
        now := time.Now()
        // pass the name user accordingly with the name provided
        user, err := getUserByName(db, name)
        if err != nil {
            log.Fatal(err)
        }
        chatFormat := now.Format("[15:04] <" + user.Name + "> ")
        fmt.Print(chatFormat)
        if !in.Scan() {
            break
        }
        cmd := in.Text()
        out := CriBot(cmd) 
        chatFormat = now.Format("[15:04] <cribot> ")
        fmt.Print(chatFormat + out)
    }
}

func CriBot(cmd string) string {
    //TODO: Stack scripts 
    //TODO: Concurrency
    switch cmd {
    case "!time":
        return time.Now().String() + "\n"
    case "!os":
        os := runtime.GOOS
        return fmt.Sprintf("Your OS is: %s\n", os)
    case "!commands":
        return fmt.Sprintln("!time, !os")
    default:
        if strings.HasPrefix(cmd, "!") {
            return fmt.Sprintln("This command does not exist. to see the commands infer: !commands")
        } else {
            return fmt.Sprintln("Give me a command (starting with '!')")
        }
    }
}
