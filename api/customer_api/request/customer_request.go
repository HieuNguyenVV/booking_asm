package request

type CreateCustomerRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UpdateCustomerRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	License_id string `json:"licenseId"`
	Active     bool   `json:"active"`
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
type FindCustomerRequest struct {
	ID string `json:"id"`
}
type ChangePasswordRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
