package user

// TransformedUser for user returns
type TransformedUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Cover     string `json:"cover"`
	Email     string `json:"email"`
}

// TransformUser transform user for view
func (u User) TransformUser() TransformedUser {
	var Tuser TransformedUser

	Tuser.ID = u.ID
	Tuser.Email = u.GetEmail()
	Tuser.FirstName = u.GetFirstName()
	Tuser.LastName = u.GetLastName()
	Tuser.Image = u.GetImage()
	Tuser.Cover = u.GetCover()

	return Tuser
}
