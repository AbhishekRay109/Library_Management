package main

type detail struct {
	PersonName     string `json:"personname" bson:"personname"`
	LastIssueDate  string `json:"issue" bson:"issue"`
	LastReturnDate string `json:"return" bson:"return"`
}

// Iniatialising all the structure of data
type User struct {
	//ID        int     `json:"id" bson:"user_id"`
	BookName   string `json:"bookname" bson:"bookname"`
	Catergory  string `json:"category" bson:"category"`
	BookDetail detail `json:"bookdetail" bson:"bookdetail"`
	RentPerDay string `json:"rentperday" bson:"rentperday"`
}
