package main

import "./internal/api/apiserver"

func main(){
	s := apiserver.New()
	if err:=s.Start(); err != nil {
		log.Fatal(err)
	}
}