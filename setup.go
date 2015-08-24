package mpower

const (
	baseLive = "https://app.mpowerpayments.com/api/v1"
	baseTest = "https://app.mpowerpayments.com/sandbox-api/v1"
)

// Setup as defined by mpower docs with the exception of the BaseURL
type Setup struct {
	MasterKey  string
	PrivateKey string
	PublicKey  string
	Token      string
	Headers    map[string]string
	BaseURL    string
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

// NewSetupFromEnv creates a setup from your environment keys
func NewSetupFromEnv() *Setup {
	setup := &Setup{
		MasterKey:  env("MP_Master_Key"),
		PrivateKey: env("MP_Private_Key"),
		PublicKey:  env("MP_Public_Key"),
		Token:      env("MP_Token"),
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
