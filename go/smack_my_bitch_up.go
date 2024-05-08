package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	output1, err := exec.Command("who").Output()
	output2 := os.Getenv("USER")
	users := string(output1[:])
	current_user := string(output2[:])
	if !strings.Contains(users, current_user) {
		return
	}

	reasons := []string{"Working hard", "Gotta ship this feature", "Someone fucked the system again"}

	rand.Seed(time.Now().UTC().UnixNano())
	message := "Late at work. " + reasons[rand.Intn(len(reasons))]

	MY_NUMBER := string(os.Getenv("MY_NUMBER"))
	HER_NUMBER := string(os.Getenv("HER_NUMBER"))

	response, err := exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "localhost:8080/v2/send", "-d", fmt.Sprintf(`{"message": "%s", "number": "%s", "recipients": ["%s"]}`, message, MY_NUMBER, HER_NUMBER)).Output()

	fmt.Printf("Command output:\n%s\n", response)

	if err != nil {
		fmt.Printf("Failed to send SMS: %s", err)
		return
	}

	fmt.Printf("Message Sent Successfully with response: %s ", response)
}
