package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/makaryb/nashare/http-rest-api/internal/app/apiserver"
	"log"
)

// путь к файлу, который будет являться нашим конфигом, будет задаваться в кач-ве флага при запуске бинарника
var (
	configPath string
)

func init() {
	// что хотим парсить
	flag.StringVar(&configPath,
		"config-path",
		"configs/apiserver.toml",
		"path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	// распарсить наш файл и записать все, что нужно в переменную config
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
