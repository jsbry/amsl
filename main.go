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

	cmd := exec.Command("arkmanager", "status")
	cmd.Stdout = fp
	cmd.Stderr = fp
	cmd.Run()

	return nil
}
