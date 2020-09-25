package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router) {

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, msg interface{}) {
		log.Println("Panic: msg")
		SendResponse(w, Response{
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		})
	}

	// Todo Routes
	router.POST("/todo", CreateTodo)
	router.GET("/todo", GetTodo)
	router.PATCH("/todo", UpdateTodo)
	router.DELETE("/todo", DeleteTodo)

	// List Item Routes

	router.POST("/todo/:id/item", CreateListItem)
	router.GET("/todo/:id/item", GetListItem)
	router.PATCH("/todo/:id/item", UpdateListItem)
	router.DELETE("/todo/:id/item", DeleteListItem)
}

func SendResponse(w http.ResponseWriter, data Response) {
	w.WriteHeader(data.StatusCode)
	data.Status = http.StatusText(data.StatusCode)
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		log.Println(err)
	}
}
