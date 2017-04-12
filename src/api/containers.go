package api

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	middleware "../http"
)

func RegisterContainersHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering party handlers")

	//Containers
	router.POST("/api/containers", createContainerHandler)
	router.GET("/api/containers", readContainersHandler)

	//Container
	router.GET("/api/containers/:id", readContainerHandler)
	router.PUT("/api/containers/:id", updateContainerHandler)
	router.DELETE("/api/containers/:id", deleteContainerHandler)

	//Container Reports
	router.GET("/api/containers/:id/reports", readContainerReportsHandler)

	//Container Report
	router.GET("/api/containers/:id/reports/:reportId", readContainerReportHandler)

	//Container Re-evaluate
	router.PUT("/api/containers/:id/_evaluate", reevaluateContainerReportsHandler)
}

func createContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func updateContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func deleteContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainerReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainerReportHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func reevaluateContainerReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}


