package dtos

type Franchise struct {
	Name		string		`json:"name"`
	StartDate	string		`json:"start_date"`
	EndDate		string		`json:"end_date"`
	Image		string		`json:"image"`
	Description	string		`json:"description"`
	CreatedOn	string		`json:"created_on"`
	CreatedBy	string		`json:"created_by"`
	Status		string		`json:"status"`
}