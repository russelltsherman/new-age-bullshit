package cli

import (
	"fmt"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// TextCommand returns a cobra command
func TextCommand() *cobra.Command {
	var sentences int

	cmd := &cobra.Command{
		Use:   "text",
		Short: "pipe output to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(bs.Text(sentences))
		},
	}

	cmd.Flags().IntVarP(&sentences, "sentences", "n", 1, "number of sentences to generate")

	return cmd
}
