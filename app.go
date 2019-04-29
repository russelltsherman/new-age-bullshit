package main

import (
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
)

func main() {
	for {
		sentence := bs.Sentence()

		// executing using osx say
		// say is good sound but only available on OSX
		// cmd := exec.Command("say", sentence)
		// cmd.Run()

		// executing using espeak
		// voices are a bit screechy
		// playback on pi zero frequently stops and slurs
		// cmd := exec.Command("espeak", "-ven+f3", "-k5", "-s150", sentence)
		// cmd.Run()

		// executing using Cepstral voice
		// Cepstral is certainly the most natural sounding voice
		// playback on pi zero frequently stops and slurs
		// cmd := exec.Command("aoss", "swift", sentence)
		// cmd.Run()

		// executing using festival
		// voice quality is ok
		// playback on pi zero is very stable
		// c1 := exec.Command("echo", sentence)
		// c2 := exec.Command("festival", "--tts")
		// r, w := io.Pipe()
		// c1.Stdout = w
		// c2.Stdin = r
		// c1.Start()
		// c2.Start()
		// c1.Wait()
		// w.Close()
		// c2.Wait()

		// executing using Pico
		// voice quality is ok
		// playback on pi zero is the most stable of all tested
		c1 := exec.Command("pico2wav", "-l", "en-GB", "-w", "/tmp/output.wav", sentence)
		c1.Run()
		c2 := exec.Command("aplay", "/tmp/output.wav")
		c2.Run()

		time.Sleep(1000 * time.Millisecond)
	}
}
