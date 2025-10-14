package model

type GlassesPartialResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Status string `json:"status"`
	Brand  string `json:"brand"`
	Drawer string `json:"drawer"`
}

type GlassesSingleResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Color       string `json:"color"`
	Description string `json:"description"`
	RFID        string `json:"rfid"`
	Status      string `json:"status"`
	Drawer      string `json:"drawer"`
	Brand       string `json:"brand"`
	Company     string `json:"company"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
