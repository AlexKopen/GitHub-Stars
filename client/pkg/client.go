package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GitHub Stargazer application started")

	for {
		fmt.Println("\nEnter a list of <organization>/<repository> inputs, separated by commas. Ex. angular/angular, twilio/twilio-python:")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		fmt.Printf(text)

	}

}
