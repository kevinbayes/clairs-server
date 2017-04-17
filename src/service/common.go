package service

import (
	"log"
	"../model"
)

var newContainerChannel = make(chan *model.Container)

func init() {

	log.Print("Initiating Services")
	ContainerServiceSingleton().Init()
}