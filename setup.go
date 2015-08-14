package mpower

const (
	baseLive = "https://app.mpowerpayments.com/api/v1"
	baseTest = "https://app.mpowerpayments.com/sandbox-api/v1"
)

// The Setup as defined by mpower docs with the exception of the BaseURL
type Setup struct {
	MasterKey   string
	PrivateKey  string
	PublicKey   string
	Token       string
	ContentType string
	Headers     map[string]string
	BaseURL     string
}

// Get - gets a value from the struct by using its field name
//
// Example.
//    key := newSetup.Get("MasterKey")
func (setup *Setup) Get(fieldName string) string {
	return get(setup, fieldName)
}

// NewSetup - returns a new setup object
//
// Example.
//    newSetup := mpower.NewSetup(map[string]string{
//        "masterKey":  YOUR MASTER KEY,
//        "privateKey": YOUR PRIVATE KEY,
//        "publicKey":  "YOUR PUBLIC KEY,
//        "token":      YOUR TOKEN,
//        "mode":       MODE,
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
		setup.BaseURL = baseLive
	} else {
		setup.BaseURL = baseTest
	}

	setup.Headers = make(map[string]string)
	setup.Headers["MP-Master-Key"] = setup.MasterKey
	setup.Headers["MP-Private-Key"] = setup.PrivateKey
	setup.Headers["MP-Public-Key"] = setup.PublicKey
	setup.Headers["MP-Token"] = setup.Token
	setup.Headers["Content-Type"] = setup.ContentType

	return setup
}
