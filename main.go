package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func printModelResponse(resp *genai.GenerateContentResponse) {
	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				text := fmt.Sprintf("%s", part)
				fmt.Printf("%s", text)
			}
			fmt.Println()
		}
	}
}

func processCLIArgs(args []string, ctx context.Context, model *genai.GenerativeModel) error {
	start := time.Now()

	resp, err := model.GenerateContent(ctx, genai.Text(args[0]))
	if err != nil {
		return err
	}

	end := time.Now()
	duration := end.Sub(start).Seconds()
	log.Printf("debug: took %.2f secs to get a response from AI", duration)
	fmt.Println()

	printModelResponse(resp)

	return nil
}

func processUserInputLoop(ctx context.Context, model *genai.GenerativeModel) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		scanned := scanner.Scan()

		if scanned {
			userInput := scanner.Text()
			userInput = strings.TrimSpace(userInput)

			if len(userInput) == 0 {
				continue
			}

			start := time.Now()

			resp, err := model.GenerateContent(ctx, genai.Text(userInput))
			if err != nil {
				return err
			}

			end := time.Now()
			duration := end.Sub(start).Seconds()
			log.Printf("debug: took %.2f secs to get a response from AI", duration)
			fmt.Println()

			printModelResponse(resp)
		}

		if !scanned || scanner.Err() != nil {
			break
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	if len(os.Args) > 1 {
		if err := processCLIArgs(os.Args[1:], ctx, model); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := processUserInputLoop(ctx, model); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done...")
}
