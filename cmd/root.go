package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "go-todo",
	Short: "A stupid todo app in go",
	Long:  `this is a simple todo app I wrote to learn go`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if customId {
			return customId()
		}

		return standardPrint()
	},
}

func addTodo() {

}

func init() {
	rootCmd.Flags().BoolVarP(&customId, "custonId", false, "give todo a custom Id")
	rootCmd.Flags().BoolVarP(&customId, "custonId", false, "give todo a custom Id")
	rootCmd.Flags().BoolVarP(&customId, "custonId", false, "give todo a custom Id")
}

func customId() {

}

func search() {

}

func standardPrint()
