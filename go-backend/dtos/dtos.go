package dtos

import uuid "github.com/satori/go.uuid"

type Franchise struct {
	Name		string		`json:"name"`
	StartDate	string		`json:"start_date"`
	EndDate		string		`json:"end_date"`
	Image		string		`json:"image"`
	Description	string		`json:"description"`
	CreatedBy	string		`json:"created_by"`
}

type GetFranchise struct {
	Name		*string		`json:"name"`
	StartDate	*string		`json:"start_date"`
	EndDate		*string		`json:"end_date"`
	Image		*string		`json:"image"`
	Description	*string		`json:"description"` 
}

type Profile struct {
	Id			string		`json:"id"`
}

type Place struct {
	Name		string		`json:"name"`
	Franchise	string		`json:"franchise"`
	CreatedBy	string		`json:"created_by"`
	Description	*string		`json:"description"`
	Image		*string		`json:"image"`
}

type GetPlace struct {
	Name		*string		`json:"name"`
	Franchise	*GetFranchise	`json:"franchise"`
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

type GetCharacter struct {
	Name		*string		`json:"name"`
	Franchise	*GetFranchise	`json:"franchise"`
	Description	*string		`json:"description"`
	Image		*string		`json:"image"`
	ExpressiveColor	*string	`json:"expressive_color"`
	Specie		*GetSpecie	`json:"specie"`
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

type GetSpecie struct {
	Name	*string		`json:"name"`
	NickName	*string	`json:"nick_name"`
	Franchise	*string	`json:"franchise"`
	Description	*string	`json:"description"`
	Image		*string	`json:"image"`
	Place		*GetPlace `json:"place"`
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

type GetInventor struct {
	Name	*string		`json:"name"`
	Image	*string		`json:"image"`
	ID		*uuid.UUID	`json:"id"`
}

type GetGadget struct {
	Name	*string		`json:"name"`
	NickName *string		`json:"nick_name"`
	Type	*string		`json:"type"`
	Image	*string		`json:"image"`
	ExpressiveColor	*string	`json:"expressive_color"`
	Description	*string	`json:"description"`
	Franchise	*GetFranchise	`json:"franchise"`
	Inventor	*GetInventor	`json:"inventor"`
}



