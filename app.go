package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	bs "github.com/russelltsherman/new-age-bullshit/generate"
)

func main() {
	for {
		sentence := bs.Sentence()
		fmt.Println(sentence)

		cmd := exec.Command("say", sentence)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		time.Sleep(800 * time.Millisecond)
	}
}
