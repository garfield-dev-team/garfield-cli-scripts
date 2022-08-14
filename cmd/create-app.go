package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(createAppCmd)
}

var createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "To create simple react component",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Component name must be specified")
		}
		name := args[0]
		template, err := os.ReadFile("./template/index.tsx")
		if err != nil {
			log.Fatal(err)
		}
		fileString := string(template)
		fileString = strings.ReplaceAll(fileString, "App", name)
		log.Printf("create react component: %s", fileString)
		err = os.WriteFile("./dist", []byte(fileString), 0666)
		if err != nil {
			log.Fatal(err)
		}
	},
}
