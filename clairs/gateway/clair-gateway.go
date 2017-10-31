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
package gateway

import (
	"net/http"
	"../config"
	"../model"
	clairDto "../clairpb"
	"fmt"
	"time"
	"log"
	grpc "google.golang.org/grpc"
	"context"
)

var clairHttp = &http.Client{
	Timeout: time.Minute * 10,
}

type ClairClient struct { }

func ClairClientInstance() *ClairClient {

	return &ClairClient{}
}



func (c * ClairClient) AnalyzeImage(container *model.ContainerImage, _layers []string) {

	conf := config.GetConfig()

	serverAddr := fmt.Sprintf("%s:%s", conf.Clair.Host, conf.Clair.Port)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	grpc.WithInsecure()
	cc, err := grpc.Dial(serverAddr, opts...)
	if err != nil {

		log.Printf("Error connecting to gRPC [%s]", err)
	}
	defer cc.Close()
	
	client := clairDto.NewAncestryServiceClient(cc)

	var layers []*clairDto.PostAncestryRequest_PostLayer

	for _, item := range _layers {

		postLayer := &clairDto.PostAncestryRequest_PostLayer {
			Hash: item,
			Path: fmt.Sprintf("%s/tmp/0/%s/layer.tar", conf.Clair.Share, item),
		}

		layers = append(layers, postLayer)
	}

	req := &clairDto.PostAncestryRequest{
		AncestryName: container.Image,
		Format: "docker",
		Layers: layers,
	}

	res, err := client.PostAncestry(context.Background(), req)

	log.Printf("Response: %s", res)
	log.Printf("Error: %s", err)

	log.Printf("Completed analyzing %s", container.Image)
}

func (c * ClairClient) GetAnalysis(container *model.ContainerImage) (*clairDto.Layer, error) {

	serverAddr := "localhost:6060"

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	cc, err := grpc.Dial(serverAddr, opts...)
	if err != nil {

		log.Printf("Error connecting to gRPC [%s]", err)
	}
	defer cc.Close()

	client := clairDto.NewAncestryServiceClient(cc)

	req := &clairDto.GetAncestryRequest {
		AncestryName: container.Image,
		WithFeatures: true,
		WithVulnerabilities: true,
	}

	res, err := client.GetAncestry(context.Background(), req)
	if err != nil {

		log.Printf("Error connecting to gRPC [%s]", err)
	}

	log.Printf("Response [%s]", res)
	log.Printf("Layers [%s]", res.GetAncestry().Layers)
	log.Printf("Vuln [%s]", res.GetAncestry().Features[0].Vulnerabilities)

	for _,item := range res.Ancestry.Features {

		log.Printf("Vulns for [%s]: [%s]", item.Name, item.Vulnerabilities)
	}


	return nil, nil
}
