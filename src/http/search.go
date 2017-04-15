package http

type ListMeta struct {

	Size int
	Page int
	Pages int
}

type ResponseList struct {
	Meta ListMeta
	Entities interface{}
	Links []Link
}

func MakeSearchResult(size int, page int, pages int, entities interface{}, links []Link) *ResponseList {

	return &ResponseList{
		Meta: ListMeta{
			Size: size,
			Page: page,
			Pages: pages,
		},
		Entities: entities,
		Links: links,
	}
}