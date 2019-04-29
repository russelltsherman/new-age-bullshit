package main

import (
	"io"
	"log"
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
)

func main() {
	for {
		sentence := bs.Sentence()
		log.Println(sentence)

		// executing using osx say
		// say is good sound but only available on OSX
		// cmd := exec.Command("say", sentence)
		// err := cmd.Run()
		// if err != nil {
		// 	log.Fatalf("cmd.Run() failed with %s\n", err)
		// }

		// executing using espeak
		// espeak frequently freezes
		// cmd := exec.Command("espeak", "-ven+f3", "-k5", "-s150", sentence)
		// err := cmd.Run()
		// if err != nil {
		// 	log.Fatalf("cmd.Run() failed with %s\n", err)
		// }

		// executing using festival
		// festival is most stable playback tested
		c1 := exec.Command("echo", sentence)
		c2 := exec.Command("festival", "--tts")
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
}
