/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func StreamResponse(query string, postUrl string, model string) {
	jsonBody := fmt.Sprintf(`{"model":"%s","prompt":"%s"}`, model, query)
	fmt.Println(jsonBody)
	var jsonData = []byte(jsonBody)
	request, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	fullResponse := ""
	reader := bufio.NewReader(response.Body)
	for {
		var tempResult map[string]any
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		json.Unmarshal([]byte(line), &tempResult)
		streamData := tempResult["response"]
		fmt.Print(streamData)
		fullResponse = fullResponse + streamData.(string)
	}
}

// setserverCmd represents the setserver command
var setserverCmd = &cobra.Command{
	Use:   "setserver",
	Short: "This is to enter the cli chat window with the local hosted llm",
	Long: `This tool is to give the command in the following format 

	Example:

	jpt setserver <host>:<port> <model>

	Please follow the following format to enter the chat window
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setserver called")
		arg_url := args[0]
		model := args[1]
		fmt.Println(arg_url)
		postUrl := fmt.Sprintf("http://%s/api/generate", arg_url)
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')

			if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP Client exiting.......")
				return
			}
			StreamResponse(strings.TrimSpace(string(text)), postUrl, model)
			fmt.Print("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(setserverCmd)
}
