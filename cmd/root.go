package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/symball/excemplate/config"
	"github.com/symball/excemplate/lib"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.ConfigInit(cfgFile)

			templateGen, err := lib.NewAgent(lib.Config{
				File:            args[0],
				ConfigSheet:     viper.GetString(config.SheetConfigControl),
				GeneratorsSheet: viper.GetString(config.SheetGenerators),
				TemplateSheet:   viper.GetString(config.SheetTemplates),
			})
			defer templateGen.CloseExcel()
			if err != nil {
				log.Fatalf("Unable to create template agent: %s", err)
			}

			var outputFile string
			if len(args) > 1 {
				outputFile = args[1]
			} else {
				outputFile = viper.GetString(config.OutputFile)
			}

			_, err = templateGen.SheetConfigParse()
			if err != nil {
				log.Fatalf("Error processing Sheet state: %s \n", err)
			}
			log.Println("Parsed sheet state")

			// Parse the templates
			_, err = templateGen.TemplateParse()
			if err != nil {
				log.Fatalf("Error processing templates: %s \n", err)
			}
			log.Println("Parsed templates")

			// Parse the generator items
			_, err = templateGen.GeneratorParse()
			if err != nil {
				log.Fatalf("Error processing generator items: %s \n", err)
			}
			err = templateGen.ProcessExcelFile(outputFile)
			if err != nil {
				log.Fatalf("Error processing Excel sheet: %s \n", err)
			}
		},
		Use:   "excemplate",
		Short: `A simple templating engine that allows the user to mass produce Excel content based on templates`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path (default is ./config.yaml)")

	rootCmd.PersistentFlags().String(
		config.SheetConfigControl,
		"Sheet Control",
		"Determine what rows to start writing to sheets on",
	)
	if err := viper.BindPFlag(config.SheetConfigControl, rootCmd.PersistentFlags().Lookup(config.SheetConfigControl)); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String(
		config.SheetTemplates,
		"Template",
		"Determine what rows to start writing to sheets on",
	)
	if err := viper.BindPFlag(config.SheetTemplates, rootCmd.PersistentFlags().Lookup(config.SheetTemplates)); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String(
		config.SheetGenerators,
		"Things To Generate",
		"Content to be rendered",
	)
	if err := viper.BindPFlag(config.SheetGenerators, rootCmd.PersistentFlags().Lookup(config.SheetGenerators)); err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().String(
		config.OutputFile,
		"output.xlsx",
		"File to create after completion",
	)
	if err := viper.BindPFlag(config.OutputFile, rootCmd.PersistentFlags().Lookup(config.OutputFile)); err != nil {
		log.Fatal(err)
	}
}
