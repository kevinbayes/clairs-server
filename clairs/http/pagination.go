package http

import (
	"../repository"
	"net/http"
	"strconv"
)

func MakePagination(r *http.Request) *repository.Pagination {

	query := r.URL.Query()

	page := query.Get("p")
	size := query.Get("s")

	_page, err := strconv.ParseInt(page, 10, 0)
	if(err != nil) { _page = 0 }

	_size, err := strconv.ParseInt(size, 10, 0)
	if(err != nil) { _size = 10 }


	return &repository.Pagination{
		Page: int(_page),
		Size: int(_size),
		Offset: int(_page * _size),
	}
}
