package models

type Task struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
	Notes  string    `json:"notes"`
	notifyType string `json:"notify_type"`
	Interval  string    `json:"interval"`
	Contact  string    `json:"contact"`		
	Date      int 	`json:"date"`	
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type TaskUpdate struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
	Notes  string    `json:"notes"`
	Interval  string    `json:"interval"`
	Date      int 	`json:"date"`	
	UpdatedAt int `json:"updated_at"`
}

type TaskDelete struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
}

type TaskInfo struct {
	OwnerId    int    `json:"owner_id"`
	RefId string `json:"ref_id"`
}