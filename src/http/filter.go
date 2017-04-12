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

