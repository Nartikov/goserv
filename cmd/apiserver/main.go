package main

import (

	"log"
	"flag"
	"github.com/Nartikov/goserv/internal/app/apiserver"

)

var (
	configPath string

)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file for server")
}

func main(){
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	if err:=s.Start(); err != nil {
		log.Fatal(err)
	}
}