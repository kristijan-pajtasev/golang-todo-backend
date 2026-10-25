package todo

type ToDo struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}

func AddTodo() {}

func GetAllTodos() {}

func GetTodoById() {}

func DeleteTodoById() {}

func UpdateTodoById() {}
