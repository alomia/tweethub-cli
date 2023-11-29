package cmd

import (
	"github.com/spf13/cobra"
)

// likeCmd represents the like command
var likeCmd = &cobra.Command{
	Use:   "like",
	Short: "Like or unlike a tweet.",
	Long: `The like command allows you to like or unlike a tweet on Twitter.
You can specify the tweet's URL using the "--url" flag. If the "--undo" flag is provided,
it will undo the action, i.e., unlike the tweet. If the "--all-accounts" flag is used,
the action will be performed across all linked Twitter accounts.

Examples:
- Like a tweet:
  tweethub like --url <tweet-url>

- Unlike a tweet:
  tweethub like --url <tweet-url> --undo

- Like a tweet across all linked accounts:
  tweethub like --url <tweet-url> --all-accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case undo:
			cancel := tweetHub.UnLike(url)
			defer cancel()
		default:
			cancel := tweetHub.Like(url)
			defer cancel()
		}
	},
}

func init() {
	likeCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet.")
	likeCmd.Flags().BoolVar(&undo, "undo", false, "Undo the like action (unlike).")
	likeCmd.Flags().BoolVar(&allAccounts, "all-accounts", false, "Perform the action across all linked accounts.")

	likeCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(likeCmd)
}
