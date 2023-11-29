package cmd

import (
	"github.com/spf13/cobra"
)

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case undo:
			cancel := tweetHub.UnTweet(url)
			defer cancel()
		default:
			if allAccounts {
				for _, user := range accounts {
					tweetHub.SetUsername(user.Username)
					tweetHub.SetPassword(user.Password)

					cancel := tweetHub.Tweet(message)
					defer cancel()
				}
				return
			}
			cancel := tweetHub.Tweet(message)
			defer cancel()
		}
	},
}

func init() {
	tweetCmd.Flags().StringVarP(&message, "message", "m", "", "Specify the content of the tweet.")
	tweetCmd.Flags().BoolVar(&undo, "undo", false, "Delete the specified tweet.")
	tweetCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet to be deleted.")
	tweetCmd.Flags().BoolVar(&allAccounts, "all-accounts", false, "Use all accounts")

	rootCmd.AddCommand(tweetCmd)
}
