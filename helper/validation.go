package helper

var CustomMessages = map[string]string{
	"required": "The %s field is required.",
	"email":    "The %s field must be a valid email address.",
	"alpha":    "The %s field must contain only letters.",
	"numeric":  "The %s field must contain only numbers.",
	"min":      "The %s field must be at least %param% characters.",
	"max":      "The %s field may not be greater than %param% characters.",
	"unique":   "The %s field has already been taken.",
	"exists":   "The selected %s is invalid.",
	"between":  "The %s field must be between %param% and %param%.",
	"same":     "The %s and %s fields must match.",
	"size":     "The %s must be %param% characters.",
	"len":      "The %s must be %param% characters.",
}
