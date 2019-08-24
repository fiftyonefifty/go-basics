package resolvers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

const minLength = 4
const maxLength = 32

var betweenConstraint = fmt.Sprintf("between:%d,%d", minLength, maxLength)
var scopeBetweenConstraint = "between:4,128"

var ProcessArbitraryResourceOwner = func(w http.ResponseWriter, r *http.Request) {

	success, err := ValidateArbitraryResourceOwner(r)

	if !success {

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)

		return
	}
	w.WriteHeader(http.StatusOK)

}

func ValidateArbitraryResourceOwner(r *http.Request) (success bool, err map[string]interface{}) {
	rules := govalidator.MapData{
		"subject":             []string{"required", betweenConstraint},
		"client_id":           []string{"required", betweenConstraint},
		"client_secret":       []string{"required", betweenConstraint},
		"scope":               []string{"required", scopeBetweenConstraint},
		"arbitrary_claims":    []string{"json"},
		"arbitrary_amrs":      []string{"json"},
		"arbitrary_audiences": []string{"json"},
		"custom_payload":      []string{"json"},
	}
	messages := govalidator.MapData{
		"subject":             []string{"required", betweenConstraint},
		"client_id":           []string{"required", betweenConstraint},
		"client_secret":       []string{"required", betweenConstraint},
		"scope":               []string{"required", scopeBetweenConstraint},
		"arbitrary_claims":    []string{"json"},
		"arbitrary_amrs":      []string{"json"},
		"arbitrary_audiences": []string{"json"},
		"custom_payload":      []string{"json"},
	}
	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: false,    // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) > 0 {

		err := map[string]interface{}{"validationError": e}
		return false, err
	}
	return true, nil
}
