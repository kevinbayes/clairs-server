// Copyright 2017 Kevin Bayes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"./gateway"
	"./model"
	"log"
)

func main() {

	container := &model.Container{
		Id: 1,
		Image: "kibana:latest",
	}

	//gateway.DockerClientInstance().PullImage(registry, container)


	//imageId, parentId, err := gateway.DockerClientInstance().ImageId(container)

	err4 := gateway.DockerClientInstance().SaveImage(container)


	if(err4 != nil) {

		log.Print(err4)
		panic(err4.Error())
	}

	layers, err5 := gateway.DockerClientInstance().ImageLayers(container)

	if(err5 != nil) {

		log.Print(err5)
		panic(err5.Error())
	}

	log.Print(layers)

	//log.Printf(path)

	gateway.ClairClientInstance().AnalyzeImage(container, "1f7dc3c631cf050d003954694badabe318cc5bbd86956d872e7da14e54039a10", layers)

	layer, err3 := gateway.ClairClientInstance().GetLayer(layers[0]);

	/*log.Print(err)
	log.Print(imageId)
	log.Print(parentId)*/
	//log.Print(err2)
	log.Print(err3)
	log.Print(layer)

	vulnerabilities := 0;
	for _, v := range layer.Features {

		if(v.Vulnerabilities != nil && len(v.Vulnerabilities) > 0) {

			vulnerabilities += len(v.Vulnerabilities)
		}
	}

	log.Print(vulnerabilities)
}
