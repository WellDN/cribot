package main

import (
    "bufio"
    "fmt"
    "os"
    "runtime"
    "strings"
    "time"
)

func main() {
    IRC()
}


func IRC() {
    in := bufio.NewScanner(os.Stdin)

    fmt.Println("\n\tChatty")
    fmt.Println(strings.Repeat("-", 25)+"\n")


    for {
        now := time.Now()
        chatTime := now.Format("[15:04] ")
        fmt.Print(chatTime)
        if !in.Scan() {
            break
        }
        cmd := in.Text()
        out := CriBot(cmd) 
        fmt.Print(chatTime + out)
    }
}

func CriBot(cmd string) string {
    // Stack scripts 
    switch cmd {
    case "!time":
        return time.Now().String()
    case "!os":
        os := runtime.GOOS
        return fmt.Sprintf("Your OS is: %s", os)
    default:
        return fmt.Sprintln("Give me a command (starting with '!')")
    }
}
