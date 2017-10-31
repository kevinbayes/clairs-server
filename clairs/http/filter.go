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
package http

import (
	"github.com/julienschmidt/httprouter"
)

type FilterChain struct {

	filters []Filter `Filters for the chain.`
}

type Filter struct {

	Name string `Logical name of the filter.`
	Description string `Description of the filter.`
	Path string `Path that the filter will apply to.`
	Methods []string `Filtered methods.`
	Handler FilterHandler `Handler that will act as the filter.`
	Order int8 `Order to execute.`
}

type FilterHandler func(httprouter.Handle) httprouter.Handle

func (chain *FilterChain) addFilter(filter Filter) *FilterChain {

	chain.filters = append(chain.filters, filter);
	return chain;
}

func (filter *Filter) IsApplicable(method string, path string) bool {

	_valid := false;

	for _, verb := range filter.Methods {

		if(verb == method) {

			_valid = true;
			break;
		}
	}

	return  _valid // TODO: Path Validation && filter.Path == "/**"
}

