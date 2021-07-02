package main

import (
	"fmt"
	"net"
	"encoding/gob"


)
type Parametros struct {
	KNearest     string
	GRUPO        string 
	EDAD         string
	SEXO         string
	DOSIS        string
	UBIGEO       string
	Eleccion     string
	//RESULTADO    sortedClassVotes	
	 
}


func cliente(parametros Parametros){
	c , err := net.Dial("tcp", ":9999")
	if err != nil{
		fmt.Println(err)
		return

	}
	
	err = gob.NewEncoder(c).Encode(parametros)
	if err != nil{
		fmt.Println(err)
	}
	c.Close()

}

func main(){
	parametro := Parametros{
		KNearest: "5",
		GRUPO: "3",
		EDAD: "14",
		SEXO: "1",
		DOSIS: "0",
		UBIGEO: "14",
	}
	go cliente(parametro)
	var input string
	fmt.Scanln(&input)

}

