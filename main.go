package main

import (
	"homestead/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	cmd.Execute()
}
