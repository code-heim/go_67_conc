package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sourcegraph/conc/pool"
)

func SendEmail(email string, subject string, body string) {
	fmt.Printf("Sending email to %s\n", email)
	fmt.Printf("Subject %s\n Body: %s\n", subject, body)
	// Simulate sending email
	time.Sleep(2 * time.Second)
}

func main() {
	pool := pool.New().WithMaxGoroutines(10)

	for i := 1; i <= 100; i++ {
		email := "email" + strconv.Itoa(i+1) + "@codeheim.io"
		subject := "Welcome!"
		body := "Thank you for signing up."
		pool.Go(func() {
			SendEmail(email, subject, body)
		})
	}

	pool.Wait()
	fmt.Println("All tasks completed")
}
