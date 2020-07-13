package models

type UserRegister struct {
	UserId    string    `json:"user_id"`
	Email  string    `json:"email"`	
	Password string    `json:"password"`
	Name string    `json:"name"`
	About string    `json:"about"`		
	createdAt      int 	`json:"created_at"`
	updatedAt      int 	`json:"updated_at"`
}

type UserProfileUpdate struct {
	Name string    `json:"name"`
	About string    `json:"about"`		
	updatedAt      int 	`json:"updated_at"`
}

type UserPasswordUpdate struct {
	Password string    `json:"password"`
	NewPassword string    `json:"new_password"`
}

type UserLogin struct {
	Email  string    `json:"email"`	
	Password string    `json:"password"`
}


