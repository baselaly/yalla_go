package request

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

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
