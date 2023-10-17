package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	SheetConfigControl = "sheet_config_control"
	SheetTemplates     = "sheet_templates"
	SheetGenerators    = "sheet_generators"
	OutputFile         = "output_file"
)

func ConfigInit(cfgFile string) {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		workingPath, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".viper" (without extension).
		viper.AddConfigPath(workingPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
