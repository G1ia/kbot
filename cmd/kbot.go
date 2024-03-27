/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	//TeleToken
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {


			payload := m.Text()

			switch payload {
			case "/help":
				response := "Доступні команди:\n" +
							"/start - Почати спілкування з ботом\n" +
							"/help - Показати список доступних команд\n" +
							"hello - Привітання від бота\n" +
							"bye - Прощання з ботом і завершення спілкування\n" +
							"howareyou - Запитати бота, як він себе почуває\n" +
							"thanks - Подякувати боту\n" +
							"sorry - Вибачитися перед ботом\n" +
							"yes - Висловити згоду\n" +
							"no - Висловити відмову\n" +
							"please - Ввічливо попросити бота про щось\n" +
							"congrats - Вітати бота з досягненням\n" +
							"good - Висловити позитивні емоції перед ботом"
				err = m.Send(response)
			case "hello":
				err = m.Send(fmt.Sprintf("Hello! I`am Kbot %s", appVersion))

			case "bye":
				err = m.Send(fmt.Sprintf("Oh, leaving already? Well, until next time!"))

			case "howareyou":
				err = m.Send(fmt.Sprintf("How am I? Just dandy, thanks for asking!"))

			case "thanks":
				err = m.Send(fmt.Sprintf("You're welcome! It's not like I have anything better to do."))

			case "sorry":
				err = m.Send(fmt.Sprintf("Apology accepted. This time."))

			case "yes":
				err = m.Send(fmt.Sprintf("Yes, of course! As if there was ever any doubt."))

			case "no":
				err = m.Send(fmt.Sprintf("No? Well, I guess it's my job to be disappointed then."))

			case "please":
				err = m.Send(fmt.Sprintf("Please? Oh, alright then. Just this once."))

			case "congrats":
				err = m.Send(fmt.Sprintf("Congratulations! Now, what's next on the agenda?"))

			case "good":
				err = m.Send(fmt.Sprintf("Good? Well, I suppose that's better than bad, isn't it?"))
	

			}
			return err

		})
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
