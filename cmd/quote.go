package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Quote a tweet with a custom message.",
	Long: `The quote command allows you to quote a tweet on Twitter with a custom message.
You can specify the tweet's URL using the "--url" flag and provide a custom message
using the "--message" flag.

Examples:
- Quote a tweet with a custom message:
  tweethub quote --url <tweet-url> --message "Your custom message here"`,
	Run: func(cmd *cobra.Command, args []string) {
		messages := viper.GetStringSlice("messages")
		messagesLenght := len(messages)

		switch {
		case allAccounts:
			for _, user := range accounts {
				tweetHub.SetUsername(user.Username)
				tweetHub.SetPassword(user.Password)

				if useMessages {
					message = viper.GetString("messages.0")

					if random {
						if messagesLenght > 0 {
							idx := rand.Int63n(int64(messagesLenght))
							message = messages[idx]
						}
					}
				}

				cancel := tweetHub.Quote(url, message)
				cancel()
			}
		case useMessages:
			message = viper.GetString("messages.0")

			if random {
				if messagesLenght > 0 {
					idx := rand.Int63n(int64(messagesLenght))
					message = messages[idx]
				}
			}

			cancel := tweetHub.Quote(url, message)
			defer cancel()
		default:
			cancel := tweetHub.Quote(url, message)
			defer cancel()
		}
	},
}

func init() {
	quoteCmd.Flags().StringVarP(&message, "message", "m", "", "Specify a custom message for the quoted tweet.")
	quoteCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet to be quoted.")
	quoteCmd.Flags().BoolVar(&random, "random", false, "Radom tweet.")
	quoteCmd.Flags().BoolVar(&allAccounts, "all-accounts", false, "Use all accounts")
	quoteCmd.Flags().BoolVar(&useMessages, "use-messages", false, "Use predefined messages from the configuration file")

	quoteCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(quoteCmd)
}
