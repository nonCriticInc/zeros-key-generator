package v1

import "log"

func example1() {

	pri,pub,err:=New().Generate(struct {
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
