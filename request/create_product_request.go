package request

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatePostProductRequest validate post product request
func ValidatePostProductRequest(c *gin.Context) url.Values {
	rules := govalidator.MapData{
		"name": []string{"required", "max:250"},
	}

	messages := govalidator.MapData{
		"name": []string{"required:name field is required", "max:name must not be morethan 250 characters"},
	}

	description := c.PostForm("description")
	if description != "" {
		rules["description"] = []string{"required", "max:2000"}
		messages["description"] = []string{"required:description fieild is required", "max:description must not morethan 2000 characters"}
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
