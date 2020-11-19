package main

import (
	v1 "github.com/zeros-key-generator/src/rsa/v1"
	"log"
)

func main() {

	pri,pub,err:=v1.New().Generate(struct {
		ConvertToString bool
		ReturnByteArray bool
	}{ConvertToString:false , ReturnByteArray: false})

	if err!=nil{
		log.Println(err)
	}

	log.Print(pub)
	log.Println("")
	log.Println(pri)

}

