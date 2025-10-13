package model


type GlassesPartialResponse struct {
	ID     uint   
	Name   string 
	Color  string 
	Status string 
	Brand  string 
	Drawer string
}

type GlassesSingleResponse struct{
	ID          uint   
	Name        string 
	Type        string 
	Color       string 
	Description string 
	RFID        string 
	Status      string 
	Drawer      string 
	Brand       string 
	Company     string 
	CreatedAt string
	UpdatedAt string

}