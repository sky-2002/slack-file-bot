package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"github.com/sqweek/dialog"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3291137781171-3607791117671-GPIrB7hUmg35BTxfcjXTXrvk")
	os.Setenv("CHANNEL_ID", "C03K041GQM6")

	// app token for age bot
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03J3P3RYDU-3624781024820-ff85d7e78faaa63249e1d89a3105fc69ad3100926b787d0af90b5f16a35dbed8")

	// Upload("SLACK_BOT_TOKEN", "CHANNEL_ID")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{

		Description: "yob calculator",
		Example:     "My yob is 2002",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func Upload(slack_bot_token string, channel_id string) {
	api := slack.New(os.Getenv(slack_bot_token))
	channelArr := []string{os.Getenv(channel_id)}

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
