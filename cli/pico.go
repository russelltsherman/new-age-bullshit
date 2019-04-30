package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// PicoCommand returns a cobra command
func PicoCommand() *cobra.Command {
	var language string

	cmd := &cobra.Command{
		Use:   "pico",
		Short: "pipe output to pico TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("pico2wav", "-l", language, "-w", "/tmp/output.wav", sentence)
				c1.Run()
				c2 := exec.Command("aplay", "/tmp/output.wav")
				c2.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}

	// Supported languages :
	// English en-US
	// English en-GB
	// French fr-FR
	// Spanish es-ES
	// German de-DE
	// Italian it-IT
	cmd.Flags().StringVarP(&language, "langugage", "l", "en-GB", "the language to use")

	return cmd
}
