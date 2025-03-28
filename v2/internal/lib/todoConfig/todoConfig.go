package todoconfig

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Created     string `json:"created"`
	IsDone      bool   `json:"isDone"`
	Executor    string `json:"executor"`
	Description string `json:"description"`
}

type Todos []Todo

type TodoRequest struct {
	Title       string `json:"title,omitempty"`
	IsDone      *bool  `json:"isDone,omitempty"`
	Executor    string `json:"executor,omitempty"`
	Description string `json:"description,omitempty"`
}

type TodoInfo struct {
	All       int `json:"all"`
	Completed int `json:"completed"`
	InWork    int `json:"inWork"`
}

type Meta struct {
	TotalAmount int `json:"totalAmount"`
}

type MetaResponse struct {
	Data []Todo   `json:"data"`
	Info TodoInfo `json:"info"`
	Meta Meta     `json:"meta"`
}
