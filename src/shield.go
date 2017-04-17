package main

import (
	"./service"
)


func main() {

	_service := service.ShieldsServiceSingleton()

	output, err := _service.GetShield(1)

	if(err != nil) {

		println(err.Error())
	} else {

		print(output.String())
	}
}
