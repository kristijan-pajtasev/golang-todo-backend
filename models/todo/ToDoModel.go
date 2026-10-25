package todo

type ToDo struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

func AddTodo() {}

func GetAllTodos() []ToDo {
	todos := []ToDo{
		{
			Title:       "Todo 1",
			Content:     "Test Content",
			IsCompleted: false,
		},
		{
			Title:       "Todo 2",
			Content:     "Test Content",
			IsCompleted: false,
		},
	}
	return todos
}

func GetTodoById(todoId int) ToDo {
	return ToDo{
		Title:       "Test title",
		Content:     "Test Content",
		IsCompleted: false,
	}
}

func DeleteTodoById(todoId int) {}

func UpdateTodoById() {}
