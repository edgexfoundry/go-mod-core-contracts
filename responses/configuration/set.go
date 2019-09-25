package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

//SetConfigResponse is for SMA to use when responding to an incoming request to PUT (i.e. UPDATE) a resource.
type SetConfigResponse struct {
	Success     bool   `json:"success"`               // Success or failure of request itself.
	Description string `json:"description,omitempty"` // Info allied with the preceding result.
	isValidated bool   // internal member used for validation check
}

// Custom marshalling to make empty strings null
func (sc SetConfigResponse) MarshalJSON() ([]byte, error) {
	test := struct {
		Success     bool    `json:"success"`
		Description *string `json:"description,omitempty"`
		isValidated bool
	}{}

	test.Success = sc.Success

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
		Success     bool    `json:"success"`
		Description *string `json:"description,omitempty"`
	}{}

	//Verify that incoming string will unmarshal successfully
	if err = json.Unmarshal(data, &test); err != nil {
		return err
	}

	sc.Success = test.Success

	if test.Description != nil {
		sc.Description = *test.Description
	}

	sc.isValidated, err = sc.Validate()

	return err
}

// Validate satisfies the Validator interface
func (sc SetConfigResponse) Validate() (bool, error) {
	if sc.Description == "" {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid Description %s", sc.Description))
	}
	return true, nil
}
