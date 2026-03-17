package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Hello, world!\n")

	//config, err := loadConfig("config.toml")
	//if err != nil {
	//	log.Fatalf("failed to read config: %s", err.Error())
	//}

	// fmt.Printf("config: %v\n", config)

	err := runTmux("new-session")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

}
