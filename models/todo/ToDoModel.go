package todo

import "slices"

type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

var userTodos = make(map[string][]ToDo)

func init() {
	userTodos["user"] = []ToDo{
		{
			Id:          1,
			Title:       "Test title 1",
			Content:     "Test Content 1",
			IsCompleted: false,
		},
	}
}

func AddTodo(title string, content string, isCompleted bool) *ToDo {
	todos := userTodos["user"]
	id := todos[len(todos)-1].Id + 1
	newTodo := ToDo{
		Id:          id,
		Title:       title,
		Content:     content,
		IsCompleted: isCompleted,
	}
	userTodos["user"] = append(userTodos["user"], newTodo)
	return &newTodo
}

func GetAllTodos() []ToDo {
	todos := userTodos["user"]
	return todos
}

func GetTodoById(todoId int) *ToDo {
	var todos = userTodos["user"]
	for _, todo := range todos {
		if todo.Id == todoId {
			return &todo
		}
	}
	return nil
}

func DeleteTodoById(todoId int) bool {
	var todos = userTodos["user"]
	for index, todo := range todos {
		if todo.Id == todoId {
			slice := slices.Delete(todos, index, index+1)
			userTodos["user"] = slice
			return true
		}
	}
	return false

}

func UpdateTodoById(todoId int, title string, content string, isCompleted bool) bool {
	var todos = userTodos["user"]
	for _, todo := range todos {
		if todo.Id == todoId {
			todo.Title = title
			todo.Content = content
			todo.IsCompleted = isCompleted
			return true
		}
	}
	return false
}
