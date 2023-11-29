package cmd

import (
	"fmt"
	"os"

	"github.com/alomia/tweethub-cli/internal/tweethub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Account struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

var (
	username    string
	message     string
	url         string
	undo        bool
	allAccounts bool

	accounts []Account
	tweetHub *tweethub.TweetHub
	cfgFile  = "tweethub.yaml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tweethub-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	initConfig()

	if err := viper.UnmarshalKey("accounts", &accounts); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		os.Exit(1)
	}

	tweetHub = tweethub.New()
	tweetHub.SetUsername(accounts[0].Username)
	tweetHub.SetPassword(accounts[0].Password)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("tweethub")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error config file: The file %s does not exist in the directory\n", cfgFile)
		return
	}
}
