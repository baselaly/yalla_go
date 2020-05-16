package user

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func init() {
	// password_confirmation rule
	govalidator.AddCustomRule("confirmed", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		confirmation := strings.TrimPrefix(rule, "confirmed:")
		if val != confirmation {
			return fmt.Errorf("The %s field need confirmation", field)
		}
		return nil
	})
}

// ValidateLoginRequest validate login request
func ValidateLoginRequest(c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}

	messages := govalidator.MapData{
		"email":    []string{"required:email field is required", "email:email must be valid email"},
		"password": []string{"required:password field is required"},
	}

	opts := govalidator.Options{
		Request:         c.Request,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.Validate()
	return e
}

// ValidateRegisterRequest validate register request
func ValidateRegisterRequest(c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"email":      []string{"required", "email", "max:250"},
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
