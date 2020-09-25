package handler

import (
	"aezion/internal/todo/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
CreateListItem will create a new list item in an existing todo list
*/
func CreateListItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	listIDStr := params.ByName("id")
	id, err := strconv.Atoi(listIDStr)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	listItem := model.ListItem{}
	err = json.NewDecoder(r.Body).Decode(&listItem)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	temp, ok := TodoList.Load(int32(id))
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "todo list does not exist",
		})

		return
	}

	todoList, ok := temp.(model.TodoList)
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusInternalServerError,
			Error:      "sorry, something went wrong.",
		})

		return
	}

	todoList.Items = append(todoList.Items, listItem)
	TodoList.Store(int32(id), todoList)

	log.Println(todoList.Items)
	SendResponse(w, Response{
		ID:         int32(len(todoList.Items) - 1),
		Message:    "list item created",
		StatusCode: http.StatusCreated,
	})
}

/*
UpdateListItem API will update the existing listItem in a thread
*/
func UpdateListItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	listIDStr := params.ByName("id")
	id, err := strconv.Atoi(listIDStr)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	listItem := model.ListItem{}
	err = json.NewDecoder(r.Body).Decode(&listItem)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	temp, ok := TodoList.Load(int32(id))
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "todo list does not exist",
		})

		return
	}

	todoList, ok := temp.(model.TodoList)
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusInternalServerError,
			Error:      "sorry, something went wrong.",
		})

		return
	}

	todoList.Items[listItem.Index] = listItem
	TodoList.Store(int32(id), listItem)

	SendResponse(w, Response{
		Message:    "list item created",
		StatusCode: http.StatusOK,
	})
}

/*
DeleteListItem will delete an existing list item
*/
func DeleteListItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	listIDStr := params.ByName("id")
	id, err := strconv.Atoi(listIDStr)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	itemIndexStr := r.URL.Query().Get("item_index")
	itemIndex, err := strconv.ParseInt(itemIndexStr, 10, 32)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	temp, ok := TodoList.Load(int32(id))
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "todo list does not exist",
		})

		return
	}

	todoList, ok := temp.(model.TodoList)
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusInternalServerError,
			Error:      "sorry, something went wrong.",
		})

		return
	}

	copy(todoList.Items[itemIndex:], todoList.Items[itemIndex+1:])
	todoList.Items = todoList.Items[:len(todoList.Items)-1]

	TodoList.Store(int32(id), todoList)
	SendResponse(w, Response{
		StatusCode: http.StatusOK,
		Error:      "item has been deleted.",
	})
}

/*
GetListItem API return a single list item in the given todo list

`id` in the request URI should be a  valid todo list
`item_index` should exist
*/
func GetListItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	listIDStr := params.ByName("id")
	id, err := strconv.Atoi(listIDStr)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	itemIndexStr := r.URL.Query().Get("item_index")
	itemIndex, err := strconv.Atoi(itemIndexStr)
	if err != nil {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})

		return
	}

	temp, ok := TodoList.Load(int32(id))
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusBadRequest,
			Error:      "todo list does not exist",
		})

		return
	}

	todoList, ok := temp.(model.TodoList)
	if !ok {
		SendResponse(w, Response{
			StatusCode: http.StatusInternalServerError,
			Error:      "sorry, something went wrong.",
		})

		return
	}

	if itemIndex > len(todoList.Items) {
		SendResponse(w, Response{
			StatusCode: http.StatusInternalServerError,
			Error:      "List item index is out of range ",
		})

		return
	}

	var data interface{}
	if len(todoList.Items) > 0 {
		data = todoList.Items[itemIndex]
	}

	SendResponse(w, Response{
		StatusCode: http.StatusOK,
		Message:    "item has been retrived.",
		Data:       data,
	})
}
