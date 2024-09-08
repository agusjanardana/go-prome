package web

type NoteResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedBy string `json:"created_by"`
}
