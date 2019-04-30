package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// EspeakCommand returns a cobra command
func EspeakCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "espeak",
		Short: "pipe output to espeak TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("espeak", "-ven+f3", "-k5", "-s150", sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}
	return cmd
}
