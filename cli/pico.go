package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// PicoCommand returns a cobra command
func PicoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pico",
		Short: "pipe output to pico TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("pico2wav", "-l", "en-GB", "-w", "/tmp/output.wav", sentence)
				c1.Run()
				c2 := exec.Command("aplay", "/tmp/output.wav")
				c2.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}
	return cmd
}
