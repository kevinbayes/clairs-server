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
	"text/template"
	"bytes"
)

type ShieldsService struct { }

func ShieldsServiceInstance() *ShieldsService {

	return &ShieldsService{}
}


func (s *ShieldsService) GetShield(containerId int64) (bytes.Buffer, error) {

	badge := &model.Shield{
		Subject: model.Text{
			Value: "clair",
		},
		Status: model.Text{
			Value: "not implemented",
		},
		Colour: "blue",
		Template: "flat",
	}

	_width1 := s.textWidth(badge.Subject.Value)
	_width2 := s.textWidth(badge.Status.Value)

	badge.Subject.Width = _width1 + 10
	badge.Status.Width = _width2 + 10
	badge.Width = badge.Subject.Width + badge.Status.Width

	badge.Subject.X = badge.Subject.Width / 2
	badge.Status.X = badge.Subject.Width + ( badge.Status.Width / 2 - 1 )

	tmpl, err := template.New("badge").ParseFiles("./src/shield/shield.flat.svg")

	if err != nil { return bytes.Buffer{}, err }

	var doc bytes.Buffer

	err = tmpl.ExecuteTemplate(&doc, "shield.flat.svg", badge)

	return doc, err
}

func (s *ShieldsService) textWidth(text string) int {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: gopdf.Rect{W: 595.28, H: 841.89}})

	err := pdf.AddTTFFont("DejaVu Serif", "./src/shield/DejaVuSerif.ttf")
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