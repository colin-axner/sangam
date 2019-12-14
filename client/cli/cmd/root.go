package cmd

import (
	"fmt"
	ipfsCmds "github.com/ipfs/go-ipfs/core/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// default home directory
var homeDir = os.ExpandEnv("$HOME/.sangamcli/")

// RootCmd returns the initialized root cmd for sangamcli
func RootCmd() *cobra.Command {
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().String(flags.Home, homeDir, "home directory for sangamcli")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(
		ipfsCmds.KeyCmd(),
	)

	//	rootCmd.AddCommand(
	//	tx.RootCmd(),
	//	query.RootCmd(),
	//	client.LineBreak,

	//	RestServerCmd(),
	//	client.LineBreak,

	//	keys.RootCmd(),
	//	client.LineBreak,

	//	VersionCmd(),
	//	)

	return rootCmd
}

var rootCmd = &cobra.Command{
	Use:           "sangamcli",
	Short:         "Sangam Client",
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		//		config.RegisterViperAndEnv()
		homeDir := viper.GetString(flags.Home)

		//		store.InitKeystore(homeDir)

		//		configFilepath := filepath.Join(homeDir, "config.toml")
		//		if _, err := os.Stat(configFilepath); os.IsNotExist(err) {
		//			if err := config.WriteConfigFile(configFilepath, config.DefaultConfig()); err != nil {
		//				return err
		//			}
		//		} else if err != nil {
		//			return err
		//		}

		//		viper.AddConfigPath(homeDir)
		//		viper.SetConfigName("config")
		//		if err := viper.MergeInConfig(); err != nil {
		//			return err
		//		}

		return nil
	},
}
