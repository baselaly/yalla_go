package product

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

// SetName function to set product name
func (p *Product) SetName(name string) {
	p.Name = name
}

// GetName function to get name
func (p *Product) GetName() string {
	return p.Name
}

// SetDescription function to set description
func (p *Product) SetDescription(description string) {
	p.Description = description
}

// GetDescription function to get description
func (p *Product) GetDescription() string {
	return p.Description
}

// SetDescription function to set description
func (p *Product) SetUserID(userID uint) {
	p.UserID = userID
}

// GetDescription function to get description
func (p *Product) GetUserID() uint {
	return p.UserID
}
