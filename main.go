package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	for {
		err := loopCmd()
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Second)
	}
}

func loopCmd() error {
	fp, err := os.Create("arks-status.txt")
	if err != nil {
		return err
	}
	defer fp.Close()

	cmd := exec.Command("bash", "-c", "arkmanager status | sed -r \"s/\\x1B\\[([0-9]{1,2}(;[0-9]{1,2})?)?[mGK]//g\"")
	cmd.Stdout = fp
	cmd.Start()
	cmd.Wait()

	return nil
}
