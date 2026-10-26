package todo

import (
	"slices"
)

type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

var userTodos = make(map[string][]ToDo)

func AddTodo(userId string, title string, content string, isCompleted bool) *ToDo {
	if userTodos[userId] == nil {
		userTodos[userId] = []ToDo{}
	}
	todos := userTodos[userId]
	var id int
	if len(todos) >= 1 {
		id = todos[len(todos)-1].Id + 1
	} else {
		id = 1
	}
	newTodo := ToDo{
		Id:          id,
		Title:       title,
		Content:     content,
		IsCompleted: isCompleted,
	}
	userTodos[userId] = append(userTodos[userId], newTodo)
	return &newTodo
}

func GetAllTodos(userId string) []ToDo {
	todos := userTodos[userId]
	return todos
}

func GetTodoById(userId string, todoId int) *ToDo {
	var todos = userTodos[userId]
	for _, todo := range todos {
		if todo.Id == todoId {
			return &todo
		}
	}
	return nil
}

func DeleteTodoById(userId string, todoId int) bool {
	var todos = userTodos[userId]
	for index, todo := range todos {
		if todo.Id == todoId {
			slice := slices.Delete(todos, index, index+1)
			userTodos[userId] = slice
			return true
		}
	}
	return false

}

func UpdateTodoById(userId string, todoId int, title string, content string, isCompleted bool) *ToDo {
	var todos = userTodos[userId]
	for index, todo := range todos {
		if todo.Id == todoId {
			userTodos[userId][index].Content = content
			userTodos[userId][index].Title = title
			userTodos[userId][index].IsCompleted = isCompleted
			return &userTodos[userId][index]
		}
	}
	return nil
}
