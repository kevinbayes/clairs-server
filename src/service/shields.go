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
package service

import (
	"github.com/signintech/gopdf"
	"log"
	"github.com/signintech/gopdf/fontmaker/core"
	"../model"
	"../config"
	"text/template"
	"bytes"
	"fmt"
)

type ShieldsService struct { }

func ShieldsServiceInstance() *ShieldsService {

	return &ShieldsService{}
}


func (s *ShieldsService) GetShield(containerId int64) (bytes.Buffer, error) {

	shield := &model.Shield{
		Subject: model.Text{
			Value: "clair",
		},
		Status: model.Text{
			Value: "not implemented",
		},
		Colour: "blue",
		Template: "flat",
	}

	_width1 := s.textWidth(shield.Subject.Value)
	_width2 := s.textWidth(shield.Status.Value)

	shield.Subject.Width = _width1 + 10
	shield.Status.Width = _width2 + 10
	shield.Width = shield.Subject.Width + shield.Status.Width

	shield.Subject.X = shield.Subject.Width / 2
	shield.Status.X = shield.Subject.Width + ( shield.Status.Width / 2 - 1 )

	tmpl, err := template.New(shield.Template).ParseFiles(fmt.Sprintf("%s/shield/shield.flat.svg", config.GetConfig().Server.Filepath))

	if err != nil { return bytes.Buffer{}, err }

	var doc bytes.Buffer

	err = tmpl.ExecuteTemplate(&doc, "shield.flat.svg", shield)

	return doc, err
}

func (s *ShieldsService) textWidth(text string) int {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}})

	err := pdf.AddTTFFont("DejaVu Serif", fmt.Sprintf("%s/shield/DejaVuSerif.ttf", config.GetConfig().Server.Filepath))
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	err2 := pdf.SetFont("DejaVu Serif", "", 12)

	if(err2 != nil) {

		panic(err2)
	} else {

		_width, _ := pdf.MeasureTextWidth(text)

		return core.Round(_width)
	}
}