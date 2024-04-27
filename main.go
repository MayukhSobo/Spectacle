package main

import (
	"log"
	"os"
	"spectacle/cmd"
	// "github.com/MayukhSobo/spectacle/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
