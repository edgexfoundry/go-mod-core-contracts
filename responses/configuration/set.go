package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

//SetConfigResponse is for SMA to use when responding to an incoming request to PUT (i.e. UPDATE) a resource.
type SetConfigResponse struct {
	Code           string   `json:"code" yaml:"code,omitempty"`                     //HTTP code to return (e.g. 200, 500, etc.)
	Description    string   `json:"description" yaml:"description,omitempty"`       //info allied with preceding HTTP code
	ExpectedValues []string `json:"expectedValues" yaml:"expectedValues,omitempty"` //expected values to return
	isValidated    bool     //internal member used for validation check
}

// Custom marshalling to make empty strings null
func (sc SetConfigResponse) MarshalJSON() ([]byte, error) {
	test := struct {
		Code           *string  `json:"code,omitempty"`
		Description    *string  `json:"description,omitempty"`
		ExpectedValues []string `json:"expectedValues,omitempty"`
		isValidated    bool
	}{
		ExpectedValues: sc.ExpectedValues,
	}

	// Empty strings are null
	if sc.Code != "" {
		test.Code = &sc.Code
	}
	if sc.Description != "" {
		test.Description = &sc.Description
	}

	return json.Marshal(test)
}

// The toString function for SetConfigResponse struct
func (sc SetConfigResponse) String() string {
	out, err := json.Marshal(sc)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

//Implements unmarshaling of JSON string to SetConfigResponse type instance
func (sc *SetConfigResponse) UnmarshalJSON(data []byte) error {
	var err error
	test := struct {
		Code           *string  `json:"code,omitempty"`
		Description    *string  `json:"description,omitempty"`
		ExpectedValues []string `json:"expectedValues,omitempty"`
	}{}

	//Verify that incoming string will unmarshal successfully
	if err = json.Unmarshal(data, &test); err != nil {
		return err
	}

	//If verified, copy the fields
	if test.Code != nil {
		sc.Code = *test.Code
	}

	if test.Description != nil {
		sc.Description = *test.Description
	}

	if test.ExpectedValues != nil {
		sc.ExpectedValues = test.ExpectedValues
	}

	sc.isValidated, err = sc.Validate()

	return err
}

// Validate satisfies the Validator interface
func (sc SetConfigResponse) Validate() (bool, error) {
	if sc.Code == "" {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid Code %s", sc.Code))
	}
	if sc.Description == "" {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid Description %s", sc.Description))
	}
	if sc.ExpectedValues == nil {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid ExpectedValues %v", sc.ExpectedValues))
	}
	return true, nil
}
