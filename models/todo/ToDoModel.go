package todo

import (
	"golang-todo-backend-project/config"
	"slices"
)

type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

var userTodos = make(map[string][]ToDo)

func init() {
	allowedUsersList := config.Config.AllowedUsersList
	for _, user := range allowedUsersList {
		userTodos[user] = []ToDo{
			{
				Id:          1,
				Title:       "Test title 1",
				Content:     "Test Content 1",
				IsCompleted: false,
			},
		}
	}
}

func AddTodo(userId string, title string, content string, isCompleted bool) *ToDo {
	todos := userTodos[userId]
	id := todos[len(todos)-1].Id + 1
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
