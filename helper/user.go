package helper

// thats how we create structure, btw type key word is needed to create new type of variable in this case its 'user'
// all need to be written in uppercase cause it need to be visible in main
type User struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	UserTickets uint32 `json:"UserTickets"`
}
