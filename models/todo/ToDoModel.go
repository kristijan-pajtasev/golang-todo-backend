package todo

import "strconv"

type ToDo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

var userTodos = make(map[string][]ToDo)

func init() {

	// Punjenje početnih podataka pod ključ "posao"
	userTodos["user"] = []ToDo{
		{Title: "ToDo 1", Content: "ToDo Content 1", IsCompleted: true},
	}
}

func AddTodo() {}

func GetAllTodos() []ToDo {
	todos := []ToDo{
		{
			Id:          1,
			Title:       "Todo 1",
			Content:     "Test Content",
			IsCompleted: false,
		},
		{Id: 2,
			Title:       "Todo 2",
			Content:     "Test Content",
			IsCompleted: false,
		},
	}
	return todos
}

func GetTodoById(todoId int) ToDo {
	return ToDo{
		Id:          todoId,
		Title:       "Test title " + strconv.Itoa(todoId),
		Content:     "Test Content " + strconv.Itoa(todoId),
		IsCompleted: false,
	}
}

func DeleteTodoById(todoId int) {}

func UpdateTodoById() {}
