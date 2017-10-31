package main

import (
	"./clairs/gateway"
	"./clairs/model"
	"log"
)

func main() {

	dto := &model.ContainerImage{}

	dto.Image = "337795043676.dkr.ecr.ap-southeast-2.amazonaws.com/data-platform/ml-base:latest"
	dto.Tags = []model.ContainerImageTag{}

	dockerClient := gateway.DockerClientInstance()

	dockerClient.SaveImage(dto)

	layers, err := dockerClient.ImageLayers(dto)
	if(err != nil) {

		log.Printf("Eish %s", err)
	}

	gateway.ClairClientInstance().AnalyzeImage(dto, layers)

	gateway.ClairClientInstance().GetAnalysis(dto)
}
