// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	ID          string `json:"id"`
	Name string `json:"companyname"`
	Location     string `json:"address"`
	Salary         string `json:"sal"`
}

type Job struct {
	ID         string `json:"id"`
	Cid        string `json:"cid"`
	Jobtitle   string `json:"jobtitle"`
	Salary     string `json:"salary"`
	Experience string `json:"experience"`
}

type NewCompany struct {
	Name string `json:"companyname"`
	Location     string `json:"address"`
	Salary         string `json:"sal"`
}

type NewJob struct {
	Cid        string `json:"cid"`
	Jobtitle   string `json:"jobtitle"`
	Salary     string `json:"salary"`
	Experience string `json:"experience"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
