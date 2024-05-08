package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

const my_number string = "+xxxxx"
const boss_number string = "+yyyyy"

func main() {
	output1, err := exec.Command("who").Output()
	output2 := os.Getenv("USER")
	users := string(output1[:])
	current_user := string(output2[:])
	if strings.Contains(users, current_user) {
		return
	}
	//create the reasons slice and append reasons to it
	reasons := make([]string, 0)
	reasons = append(reasons,
		"Locked out",
		"Pipes broke",
		"Food poisoning",
		"Not feeling well")

	// Create and send message using your signal service with api
	message := fmt.Sprint("Gonna work from home...", reasons[rand.Intn(len(reasons))])

	response, err := exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "localhost:8080/v2/send", "-d", fmt.Sprintf(`{"message": "%s", "number": "%s", "recipients": ["%s"]}`, message, my_number, boss_number)).Output()

	fmt.Printf("Command output:\n%s\n", response)

	if err != nil {
		fmt.Printf("Failed to send SMS: %s", err)
		return
	}
}
