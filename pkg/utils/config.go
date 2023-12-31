package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/bryanck29/be-test/internal/config"
)

func LoadEnv() {
	cfgFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer cfgFile.Close()

	byteValue, _ := io.ReadAll(cfgFile)
	err = json.Unmarshal(byteValue, &config.Core)
	if err != nil {
		log.Fatalln(err)
	}
}
