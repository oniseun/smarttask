package models

type Task struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
	Notes  string    `json:"notes"`
	notifyType string `json:"notify_type"`
	Interval  string    `json:"interval"`
	Phone  string    `json:"phone"`
	Email  string    `json:"email"`		
	Date      int 	`json:"date"`	
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type TaskUpdate struct {
	Notes  string    `json:"notes"`
	Interval  string    `json:"interval"`
	Date      int 	`json:"date"`	
	UpdatedAt int `json:"updated_at"`
}

type TaskDelete struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
}