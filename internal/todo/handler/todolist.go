package handler

import (
	"aezion/internal/todo/model"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/julienschmidt/httprouter"
)

var TodoList = sync.Map{}
var totalList int32 = 0

// Response is the common response format for the api request
type Response struct {
	ID         int32       `json:"id"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"-"`
	Status     string      `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

/*
CreateTodo API creates a new model.TodoList

Name for the TodoList is mandatory. Items for the TodoList are optional.
*/
func CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := model.TodoList{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	if todo.Name == "" {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "name field is required.",
		})

		return
	}

	//  Get a New List Index
	totalList = atomic.AddInt32(&totalList, 1)

	TodoList.Store(totalList, todo)
	SendResponse(w, Response{
		ID:         totalList,
		Message:    "TodoList has been created.",
		StatusCode: http.StatusCreated,
	})
}

/*
UpdateTodo API will update an existing todo list

ID of an existing todo list item should be valid, and the name field is required.
*/
func UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := model.TodoList{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	if todo.Name == "" {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "name field is required.",
		})

		return
	}

	_, ok := TodoList.Load(todo.ID)
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusNotFound,
			Error:      "list does not exist.",
		})

		return
	}

	TodoList.Store(todo.ID, todo.Items)
	SendResponse(w, Response{
		ID:         todo.ID,
		Message:    "TodoList has been updated.",
		StatusCode: http.StatusAccepted,
	})
}

/*
GetTodo this will return an existing todo list application

`id`  given in the URL query params should be a valid one
*/
func GetTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistID := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(todolistID, 10, 32)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusNotFound,
			Error:      err.Error(),
		})
		return
	}

	todoList, ok := TodoList.Load(int32(id))
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusNotFound,
			Error:      "list does not exist.",
		})
		return
	}

	SendResponse(w, Response{
		ID:         int32(id),
		Message:    "TodoList has been retrived.",
		StatusCode: http.StatusOK,
		Data:       todoList,
	})

}

/*
DeleteTodo API will delete an already existing todo list
Is there any conditions for inputs?

`id`  given in the URL query params should be a valid one
*/
func DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(todolistID)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusNotFound,
			Error:      err.Error(),
		})
		return
	}

	TodoList.Delete(id)
	SendResponse(w, Response{
		ID:         int32(id),
		Message:    "TodoList has been deleted.",
		StatusCode: http.StatusOK,
	})

}
