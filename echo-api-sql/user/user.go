package user

type User struct {
	Id   int    `json:"id"`
	Name string `json:"nickname"`
	Age  int    `json:"age"`
}

type Err struct {
	Message string `json:"message"`
}
