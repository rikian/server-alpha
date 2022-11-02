package entities

type ResponseInsertProduct struct {
	UrlImage string
	Status   string
}

type ResponseDeleteProduct struct {
	Status    string
	ProductId string
}

type ResponseUpdateProduct struct {
	Status  string
	Product Product
}
