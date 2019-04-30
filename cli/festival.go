package cli

import (
	"io"
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
	"github.com/spf13/cobra"
)

// FestivalCommand returns a cobra command
func FestivalCommand() *cobra.Command {
	var language string

	cmd := &cobra.Command{
		Use:   "festival",
		Short: "pipe output to festival TTS",
		Run: func(cmd *cobra.Command, args []string) {
			for {
				sentence := bs.Sentence()
				c1 := exec.Command("echo", sentence)
				c2 := exec.Command("festival", "--tts", "--language", language)
				r, w := io.Pipe()
				c1.Stdout = w
				c2.Stdin = r
				c1.Start()
				c2.Start()
				c1.Wait()
				w.Close()
				c2.Wait()
				time.Sleep(1000 * time.Millisecond)
			}
		},
	}

	// Supported languages :
	// english
	// spanish
	// welsh
	cmd.Flags().StringVarP(&language, "langugage", "l", "english", "the language to use")

	return cmd
}
