package cli

import (
	"fmt"
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// EspeakCommand returns a cobra command
func EspeakCommand() *cobra.Command {
	var amplitude string
	var pitch string
	var speed string
	var capital string
	var voice string

	cmd := &cobra.Command{
		Use:   "espeak",
		Short: "pipe output to espeak TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("espeak",
					fmt.Sprintf("-a%s", amplitude),
					fmt.Sprintf("-p%s", pitch),
					fmt.Sprintf("-v%s", voice),
					fmt.Sprintf("-k%s", capital),
					fmt.Sprintf("-s%s", speed),
					sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}

	cmd.Flags().StringVarP(&amplitude, "amplitude", "a", "10", "Amplitude, 0 to 20")
	cmd.Flags().StringVarP(&pitch, "pitch", "p", "50", "Pitch adjustment, 0 to 99")
	cmd.Flags().StringVarP(&speed, "speed", "s", "160", "Speed in words per minute")
	cmd.Flags().StringVarP(&capital, "capital", "k", "5", "Indicate capital letters with")
	cmd.Flags().StringVarP(&voice, "voice", "v", "en+f3", "Use voice file from espeak-data/voices")

	return cmd
}
