package cli

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// SayCommand returns a cobra command
func SayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "say",
		Short: "pipe output to OSX say TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("say", sentence)
				c1.Run()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}
	return cmd
}
