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

// NewSetup - returns a new setup object
func NewSetup(masterKey, privateKey, publicKey, token string) *Setup {
	setup := &Setup{
		MasterKey:  masterKey,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Token:      token,
	}

	setup.setupHeaders()

	return setup
}

func NewSetupFromEnv() *Setup {
	setup := &Setup{
		MasterKey:  env("MP-Master-Key"),
		PrivateKey: env("MP-Private-Key"),
		PublicKey:  env("MP-Public-Key"),
		Token:      env("MP-Token"),
	}

	setup.setupHeaders()

	return setup
}

func (s *Setup) setupHeaders() {
	s.Headers = make(map[string]string)
	s.Headers["MP-Master-Key"] = s.MasterKey
	s.Headers["MP-Private-Key"] = s.PrivateKey
	s.Headers["MP-Public-Key"] = s.PublicKey
	s.Headers["MP-Token"] = s.Token
	s.Headers["Content-Type"] = "application/json"
}
