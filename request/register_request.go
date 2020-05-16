package request

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidateRegisterRequest validate register request
func ValidateRegisterRequest(c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"email":      []string{"required", "email", "max:250", "unique:users"},
		"first_name": []string{"required", "max:250"},
		"last_name":  []string{"required", "max:250"},
		"password":   []string{"required", "min:8", "max:200", "confirmed:" + c.PostForm("password_confirmation")},
	}

	image, imageHeader, _ := c.Request.FormFile("image")
	cover, coverHeader, _ := c.Request.FormFile("cover")

	if image != nil && imageHeader != nil {
		rules["file:image"] = []string{"ext:jpg,jpeg,png", "mime:image/jpg,image/png,image/jpeg"}
	}

	if cover != nil && coverHeader != nil {
		rules["file:cover"] = []string{"ext:jpg,jpeg,png", "mime:image/jpg,image/png,image/jpeg"}
	}

	opts := govalidator.Options{
		Request:         c.Request,
		Rules:           rules,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.Validate()
	return e
}
