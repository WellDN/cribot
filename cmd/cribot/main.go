package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/welldn/cribot/pkg/database"
)

var (
	in          = bufio.NewScanner(os.Stdin)
	out         = make(chan string)
	currentName string
	mu          sync.Mutex
)

func main() {
	db, err := DbConnection()
	if err != nil {
		log.Fatal("An error occurred while connecting database", err)
	}
	defer db.Close()

    go IRC(db)

	http.HandleFunc("/", handleRequest)
	fmt.Println("ChatBot is running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IRC(db *sql.DB) {
	fmt.Println("\n\tChatty")
	fmt.Println(strings.Repeat("-", 25) + "\n")

	for {
		fmt.Printf("Enter the user's name (or enter 'guest' to enter as guest): ")
		if in.Scan() {
			mu.Lock()
			currentName = in.Text()
			mu.Unlock()

			name := strings.TrimSpace(strings.ToLower(currentName))

			if name == "guest" {
				fmt.Println("entered as guest user.")
				handleGuest()
				continue
			}
            
            _, err := database.GetUserDByName(db, name)
			if err != nil {
				log.Println("User not found in the database, please try again")
				continue
			}

			// Chat loop for registered users
			for {
				mu.Lock()
				name := currentName
				mu.Unlock()

				now := time.Now()

				user, err := database.GetUserDByName(db, name)
				if err != nil {
					log.Println("User not found, try again")
					break
				}

				chatFormat := now.Format("[15:04] <" + user.Name + "> ")
				fmt.Print(chatFormat)

				if in.Scan() {
                    // this works but refactor
                    cribot(db)

                    //cribot() // This should be what users could do, but for guests should be different
				} else {
					log.Fatal("Failed to scan input:", in.Err())
				}
			}
		} else {
			log.Fatal("Failed to scan input:", in.Err())
		}
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "User's name is required", http.StatusBadRequest)
		return
	}

	// Process the message
	cmd := r.FormValue("message")
	out <- fmt.Sprintf("%s: %s", name, cmd)

	w.Write([]byte("Message sent"))
}

// this works but refactor
func handleGuest() {
	for {
		now := time.Now()
		fmt.Printf("[%s] <guest> ", now.Format("15:04"))

		if in.Scan() {
			cmd := in.Text()
			switch cmd {
			case "!commands":
				 fmt.Println("Available commands for guest: !time, !os")
			case "!time":
				 fmt.Println("Current time:", now.Format("Mon Jan 2 15:04:05 2006"))
			case "!os":
				 fmt.Println("Your OS is:", runtime.GOOS)
			default:
				 fmt.Println("Unknown command. Available commands for guest: !time, !os")
			}
		} else {
			log.Fatal("Failed to scan input:", in.Err())
		}
	}
}

// This should have more priviliage and just the user gonna use but doing the copy just to see
// This works but refactor
func cribot(db *sql.DB) {
	for {
		now := time.Now()
        name := currentName
        user, err := database.GetUserDByName(db, name)
        if err != nil {
            log.Println("User not found, try again")
            break
        }
				chatFormat := now.Format("[15:04] <" + user.Name + "> ")
				fmt.Print(chatFormat)

		if in.Scan() {
			cmd := in.Text()
			switch cmd {
			case "!commands":
				 fmt.Println("Available commands for guest: !time, !os")
			case "!time":
				 fmt.Println("Current time:", now.Format("Mon Jan 2 15:04:05 2006"))
			case "!os":
				 fmt.Println("Your OS is:", runtime.GOOS)
			default:
				 fmt.Println("Unknown command. Available commands for guest: !time, !os")
			}
		} else {
			log.Fatal("Failed to scan input:", in.Err())
		}
	}
}
