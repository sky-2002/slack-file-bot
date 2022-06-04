package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/sqweek/dialog"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "---your-bot-user-auth-token---")
	os.Setenv("CHANNEL_ID", "---your-channel-id---")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{}                                      // create a slice to hold strings
	filename, err := dialog.File().Load()                      // select file from dialog and load it
	fmt.Printf("File is: %s and error is %v\n", filename, err) // filename contains path of selected file

	if err != nil {
		fmt.Println("Error occured while loading file.")
		return
	}

	fileArr = append(fileArr, filename) // add filename to fileArr

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
