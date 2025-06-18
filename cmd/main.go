package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Morizz00/leet-helper/problem"
	"github.com/Morizz00/leet-helper/scrapper"
	"github.com/joho/godotenv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	fmt.Println("API KEY LOADED:", apiKey)

	if apiKey == "" {
		log.Fatal("Api key not found,can't fetch shi")
	}
	agent := scrapper.NewAgent(apiKey)

	fmt.Println("Agent ready to roll! Enter Command With 'dsa' ")

	for {
		fmt.Print("-->")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace((scanner.Text()))
		if input == "" || !strings.HasPrefix(input, "dsa") {
			fmt.Println("Invalidated,start over lil bro")
			continue
		}

		args := strings.SplitN(input, " ", 2)

		if len(args) != 2 {
			fmt.Println("dsa <name-prob>")
			continue
		}
		problemName := strings.TrimSpace(args[1])

		fmt.Println("Processing...")
		code, err := agent.GetSolution(problemName)
		if err != nil {
			log.Printf("Error fetching :%v\n", err)
			continue
		}
		err = problem.CreateProblemFile(problemName, code)
		if err != nil {
			log.Printf("Err creating the file :%v\n", err)
			continue
		}
		fmt.Printf("Problem File created for :%s\n", problemName)
	}

}
