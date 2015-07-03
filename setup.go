package mpower

const (
	BASE_URL_LIVE = "https://app.mpowerpayments.com/api/v1"
	BASE_URL_TEST = "https://app.mpowerpayments.com/sandbox-api/v1"
)

// The Setup as defined by mpower docs with the exception of the BASE_URL
type Setup struct {
	MasterKey   string
	PrivateKey  string
	PublicKey   string
	Token       string
	ContentType string
	BASE_URL    string
}

// Get - gets a value from the struct by using its field name
//
// Example.
//    key := newSetup.Get("MasterKey")
func (setup *Setup) Get(fieldName string) string {
	return get(setup, fieldName)
}

// GetHeaders - gets the respective headers to set on a request for an mpower transaction
func (setup *Setup) GetHeaders() map[string]string {
	headers := make(map[string]string)

	headers["MP-Master-Key"] = setup.MasterKey
	headers["MP-Private-Key"] = setup.PrivateKey
	headers["MP-Public-Key"] = setup.PublicKey
	headers["MP-Token"] = setup.Token
	headers["Content-Type"] = setup.ContentType

	return headers
}

// NewSetup - returns a new setup object
//
// Example.
//    newSetup := mpower.NewSetup(map[string]string{
//        "masterKey":  "55647970-22e1-4e7e-8fb4-56eca2b3b006",
//        "privateKey": "test_private_B8EiE1AGWpb4tVMzVTyFDu9rYoc",
//        "publicKey":  "test_public_B1wo2UVmxUrvwzZuPqpLrWqlA74",
//        "token":      "a6d96e2586c8bbae7c28",
//        "mode":       "test",
//    })
func NewSetup(setupInfo map[string]string) *Setup {
	setup := &Setup{
		MasterKey:   envOr("MP-Master-Key", setupInfo["masterKey"]),
		PrivateKey:  envOr("MP-Private-Key", setupInfo["privateKey"]),
		PublicKey:   envOr("MP-Public-Key", setupInfo["publicKey"]),
		Token:       envOr("MP-Token", setupInfo["token"]),
		ContentType: "application/json",
	}

	if val, ok := setupInfo["mode"]; ok && val == "live" {
		setup.BASE_URL = BASE_URL_LIVE
	} else {
		setup.BASE_URL = BASE_URL_TEST
	}

	return setup
}
