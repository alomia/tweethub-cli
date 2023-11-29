package cmd

import (
	"github.com/spf13/cobra"
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
		cancel := tweetHub.Quote(url, message)
		defer cancel()
	},
}

func init() {
	quoteCmd.Flags().StringVarP(&message, "message", "m", "", "Specify a custom message for the quoted tweet.")
	quoteCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet to be quoted.")

	quoteCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(quoteCmd)
}
