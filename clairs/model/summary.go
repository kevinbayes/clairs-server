package model


type Summary struct {
	Registries RegistriesSummary
	Containers ContainersSummary
	Reports ReportsSummary
}


type RegistriesSummary struct {
	Total int
}

type ContainersSummary struct {
	Total int
}

type ReportsSummary struct {
	Total int
}