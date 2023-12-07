package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alomia/tweethub-cli/internal/tweethub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Account struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

var (
	cfgFile     string
	username    string
	message     string
	url         string
	undo        bool
	random      bool
	allAccounts bool
	useMessages bool

	accounts []Account
	tweetHub *tweethub.TweetHub
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tweethub-cli",
	Short: "A CLI tool for managing Twitter accounts and tweets.",
	Long: `Tweethub-CLI is a command-line interface (CLI) application designed for managing Twitter accounts and tweets efficiently.
It leverages the Cobra library in Go to provide a powerful and flexible user experience.

With Tweethub-CLI, you can perform various Twitter-related tasks, such as tweeting messages, deleting tweets, and managing multiple accounts.

Examples:
  - Tweet a message: tweethub-cli tweet -m "Hello, world!"
  - Delete a tweet: tweethub-cli tweet --undo --url <tweet-url>
  - Use predefined messages: tweethub-cli tweet --use-messages

Explore the full range of features by checking the available commands and their respective options.

Cobra is a CLI library for Go that empowers applications by providing a simple and elegant way to build powerful CLI tools.
This application aims to simplify the interaction with Twitter through the command line, making it a valuable tool for developers and Twitter enthusiasts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, func() {
		if err := viper.UnmarshalKey("accounts", &accounts); err != nil {
			fmt.Printf("Error unmarshaling config: %v\n", err)
			os.Exit(1)
		}
	}, func() {
		tweetHub = tweethub.New()
		tweetHub.SetUsername(accounts[0].Username)
		tweetHub.SetPassword(accounts[0].Password)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./tweethub.yaml)")
}

func initConfig() {
	baseDir, err := os.Getwd()
	cobra.CheckErr(err)

	fullPath := filepath.Join(baseDir, cfgFile)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(fullPath)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("tweethub")
	}

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(fmt.Errorf("Error reading config file: %v", err))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
