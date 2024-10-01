package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CobraCliServer struct {
	rootCmd *cobra.Command
}

func New() *CobraCliServer {
	return &CobraCliServer{
		rootCmd: &cobra.Command{
			Use:   "asmparser",
			Short: "AVR assembly parser",
			Long: `asmparser is a CLI tool for parsing and 
disassembling AVR assembly code written in Intel HEX format`,
			Run: func(cmd *cobra.Command, args []string) {},
		},
	}
}

func (c *CobraCliServer) Configure(cl *UsecaseCollector) {
	const defaultFilePath = "none"
	var filePath string
	var parseCmd = &cobra.Command{
		Use:     "parse",
		Aliases: []string{"p"},
		Short:   "Parse HEX from file or from stdin",
		Run: func(cmd *cobra.Command, args []string) {
			var progStrs []string
			var err error
			if filePath == defaultFilePath {
				progStrs, err = cl.ParseStdinUsecase.Do()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error while parse from stdin '%s'", err)
					os.Exit(1)
				}
			} else {
				progStrs, err = cl.ParseFileUsecase.Do(filePath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error while parse from file '%s'", err)
					os.Exit(1)
				}
			}
			fmt.Println("Your ATmega assembly code:")
			for _, progStr := range progStrs {
				fmt.Println(progStr)
			}
		},
	}
	parseCmd.Flags().StringVarP(&filePath, "file", "f", defaultFilePath, "Parse HEX from file")
	c.rootCmd.AddCommand(parseCmd)
}

func (c *CobraCliServer) Serve() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
