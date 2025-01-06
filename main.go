package main

import (
	"github.com/joho/godotenv"
	"github.com/tomaspavlatka/ptx-go-chef/cmd"
)

func main() {
  godotenv.Load()
	cmd.Execute()
}
