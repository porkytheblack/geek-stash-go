package dtos

type Franchise struct {
	Name		string		`json:"name"`
	StartDate	string		`json:"start_date"`
	EndDate		string		`json:"end_date"`
	Image		string		`json:"image"`
	Description	string		`json:"description"`
	CreatedBy	string		`json:"created_by"`
}

type Profile struct {
	UserName	string		`json:"username"`
	PicUrl		string		`json:"pic_url"`
}


