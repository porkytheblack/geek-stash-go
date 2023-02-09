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

type Place struct {
	Name		string		`json:"name"`
	Franchise	string		`json:"franchise"`
	CreatedBy	string		`json:"created_by"`
	Description	*string		`json:"description"`
	Image		*string		`json:"image"`
}

type Character struct {
	Name		string		`json:"name"`
	Bio			string		`json:"bio"`
	Attributes	string		`json:"attributes"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
	ExpressiveColor	string	`json:"expressive_color"`
	Specie		string		`json:"specie"`
	// Weapon		string		`json:"weapon"` -> still being worked on
}

type Specie struct {
	Name 		string		`json:"name"`
	NickName	string		`json:"nick_name"`
	Franchise	string		`json:"franchise"`
	CreatedBy	string		`json:"created_by"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
	Place		string		`json:"place"`
}

type Gadgets struct {
	Name		string		`json:"name"`
	NickName	string		`json:"nick_name"`
	Type		string		`json:"type"`
	Image		string		`json:"image"`
	ExpressiveColor	string	`json:"expressive_color"`
	CreatedBy	string		`json:"created_by"`
	Description	string		`json:"description"`
	Franchise	string		`json:"franchise"`
	Inventor	string		`json:"inventor"`
}



