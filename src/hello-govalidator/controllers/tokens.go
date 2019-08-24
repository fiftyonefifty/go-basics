package controllers

import (
	"encoding/json"
	"fmt"
	"hello-govalidator/resolvers"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

const minLength = 4
const maxLength = 32

var betweenConstraint = fmt.Sprintf("between:%d,%d", minLength, maxLength)
var scopeBetweenConstraint = "between:4,128"

// Get all books
var GetTokens = func(w http.ResponseWriter, r *http.Request) {

	rules := govalidator.MapData{
		"grant_type": []string{"required", betweenConstraint},
	}
	messages := govalidator.MapData{
		"grant_type": []string{"required:grant_type", betweenConstraint},
	}
	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.Validate()
	if len(e) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		// then
		err := map[string]interface{}{"validationError": e}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	for index, element := range r.Form {
		fmt.Println(index)
		fmt.Println(element)
	}
	switch r.FormValue("grant_type") {
	case "arbitrary_resource_owner":
		resolvers.ProcessArbitraryResourceOwner(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)

	}
}
