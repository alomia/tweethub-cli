package cmd

import (
	"github.com/spf13/cobra"
)

// repostCmd represents the repost command
var repostCmd = &cobra.Command{
	Use:   "repost",
	Short: "Repost or unrepost a tweet.",
	Long: `The repost command allows you to repost or unrepost a tweet on Twitter.
You can specify the tweet's URL using the "--url" flag. If the "--undo" flag is provided,
it will undo the action, i.e., unrepost the tweet.

Examples:
- Repost a tweet:
  tweethub repost --url <tweet-url>

- Unrepost a tweet:
  tweethub repost --url <tweet-url> --undo`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case undo:
			cancel := tweetHub.UnRepost(url)
			defer cancel()
		default:
			cancel := tweetHub.Repost(url)
			defer cancel()
		}
	},
}

func init() {
	repostCmd.Flags().StringVar(&url, "url", "", "Specify the URL of the tweet.")
	repostCmd.Flags().BoolVar(&undo, "undo", false, "Undo the repost action (unrepost).")

	repostCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(repostCmd)
}
