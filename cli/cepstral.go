package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// CepstralCommand returns a cobra command
func CepstralCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cepstral",
		Short: "pipe output to Cepstral TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("aoss", "swift", sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}
	return cmd
}
