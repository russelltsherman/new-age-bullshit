package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// SayCommand returns a cobra command
func SayCommand() *cobra.Command {
	var voice string
	var rate string

	cmd := &cobra.Command{
		Use:   "say",
		Short: "pipe output to OSX say TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("say", "-v", voice, "-r", rate, sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}

	// voice options
	// Alex Fred Samantha Victoria and more
	cmd.Flags().StringVarP(&voice, "voice", "v", "Alex", "the voice to use")
	// rate of speech
	cmd.Flags().StringVarP(&rate, "rate", "r", "170", "the rate of speech")

	return cmd
}
