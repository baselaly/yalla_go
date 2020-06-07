package product

// TransformedProduct struct for transformed product
type TransformedProduct struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

// TransformProduct function to transform product
func (p Product) TransformProduct() TransformedProduct {
	var Tproduct TransformedProduct

	Tproduct.ID = p.ID
	Tproduct.Name = p.GetName()
	Tproduct.Description = p.GetDescription()
	Tproduct.UserID = p.GetUserID()

	return Tproduct
}
