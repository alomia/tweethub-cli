package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Send a tweet on Twitter",
	Long: `The "tweet" command allows you to post tweets on Twitter.

You can specify the content of the tweet using the "--message" flag. If you want to use predefined messages from the configuration file, provide the "--use-messages" flag. Additionally, you can choose to send a random message using the "--random" flag.

Examples:
  tweethub-cli tweet --message "Hello, world!"
  tweethub-cli tweet --use-messages
  tweethub-cli tweet --undo --url <tweet-url>`,
	Run: func(cmd *cobra.Command, args []string) {
		messages := viper.GetStringSlice("messages")
		messagesLenght := len(messages)

		switch {
		case undo:
			cancel := tweetHub.UnTweet(url)
			defer cancel()

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

				cancel := tweetHub.Tweet(message)
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

			cancel := tweetHub.Tweet(message)
			defer cancel()
		default:
			cancel := tweetHub.Tweet(message)
			defer cancel()
		}
	},
}

func init() {
	tweetCmd.Flags().StringVarP(&message, "message", "m", "", "Specify the content of the tweet.")
	tweetCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet to be deleted.")
	tweetCmd.Flags().BoolVar(&random, "random", false, "Radom tweet.")
	tweetCmd.Flags().BoolVar(&undo, "undo", false, "Delete the specified tweet.")
	tweetCmd.Flags().BoolVar(&allAccounts, "all-accounts", false, "Use all accounts")
	tweetCmd.Flags().BoolVar(&useMessages, "use-messages", false, "Use predefined messages from the configuration file")

	rootCmd.AddCommand(tweetCmd)
}
