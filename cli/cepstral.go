package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// CepstralCommand returns a cobra command
func CepstralCommand() *cobra.Command {
	var voice string

	cmd := &cobra.Command{
		Use:   "cepstral",
		Short: "pipe output to Cepstral TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("aoss", "swift", "-n", voice, sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}
	// https://www.cepstral.com
	// options depend on which voices you have licensed and installed
	cmd.Flags().StringVarP(&voice, "voice", "v", "David", "the voice to use")

	return cmd
}
