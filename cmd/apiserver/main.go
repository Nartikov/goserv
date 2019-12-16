package main

import (

	"log"
	"flag"
	"github.com/Nartikov/goserv/internal/app/apiserver"
	"github.com/BurntSushi/toml"

)

var (
	configPath string

)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file for server")
}

func main(){
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	 if err != nil {
		log.Fatal(err)
	 }
	s := apiserver.New()
	if err:=s.Start(); err != nil {
		log.Fatal(err)
	}
}