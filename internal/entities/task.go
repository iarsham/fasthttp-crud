package entities

type TaskRequest struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
