/*
Copyright © 2023 Manan Patel - github.com/immnan
*/
package cmd

//import (
//	"fmt"
//	"os"

//"github.com/immnan/goinv/cmd/batch"
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lvgo",
	Short: "A Blazemeter CLI tool for Blazemeter Admins",
	Long:  "A Cli tool used for performing actions within Blazemeter platform",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println("Version: lvgo-0.1.0_beta")
			fmt.Println("Author:  https://github.com/ImMnan/")
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommand() {
	//rootCmd.AddCommand(get.GetCmd)
}
func init() {
	//	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bmgo.yaml)")
	// rootCmd.PersistentFlags().StringP("author", "a", "Manan Patel", "author name for copyright attribution")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("version", "v", false, "Version of the installed bmgo!")
	addSubCommand()
}
