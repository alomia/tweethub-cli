package cmd

import (
	"github.com/spf13/cobra"
)

// followCmd represents the follow command
var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Follow or unfollow a Twitter user.",
	Long: `The follow command allows you to interact with Twitter user relationships.
You can use it to follow or unfollow a specified user. If the "--undo" flag is provided,
it will undo the action, i.e., unfollow the user.

Examples:
- Follow a user:
  tweethub follow --username <target-username>

- Unfollow a user:
  tweethub follow --username <target-username> --undo`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case undo:
			cancel := tweetHub.UnFollow(username)
			defer cancel()
		default:
			cancel := tweetHub.Follow(username)
			defer cancel()
		}
	},
}

func init() {
	followCmd.Flags().StringVarP(&username, "username", "u", "", "Specify the target Twitter username.")
	followCmd.Flags().BoolVar(&undo, "undo", false, "Undo the follow action (unfollow).")

	followCmd.MarkFlagRequired("username")

	rootCmd.AddCommand(followCmd)
}
