package dto

import (
	"../../model"
)

type NewRegistry struct {
	Name string
	Description string
	Uri string

	Credentials model.Credentials
}
