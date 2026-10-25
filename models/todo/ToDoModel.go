package todo

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

func GetTodoById(todoId int) *ToDo {
	var userTodos = userTodos["user"]
	for _, todo := range userTodos {
		if todo.Id == todoId {
			return &todo
		}
	}
	return nil
}

func DeleteTodoById(todoId int) {}

func UpdateTodoById() {}
