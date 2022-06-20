package entities

type Todo struct {
	ID          int64  `json:"id"` //Quando estiver na resposta do json ele ira aparecer id
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
