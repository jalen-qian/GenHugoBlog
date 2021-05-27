package main

import (
	"log"
	"os"

	"github.com/jalen-qian/GenHugoBlog/cmd"
)

func main() {
	if err := cmd.NewApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
