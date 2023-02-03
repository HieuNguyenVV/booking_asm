package response

type CreateCustomerResponse struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UpdateCustomerResponse struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type ChangePasswordResponse struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type FindCustomerResponse struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
