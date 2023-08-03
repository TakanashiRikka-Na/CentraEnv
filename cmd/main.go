package main

import (
	"fmt"
	"github.com/TakanashiRikka-Na/CentraEnv/internal/config"
)

func main() {
	path := config.ReadPathEnv()
	fmt.Println(path)
}
