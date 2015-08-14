package mpower

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteTestSetup struct {
	suite.Suite
	mpowerSetup *Setup
}

func (s *SuiteTestSetup) SetupSuite() {
	s.mpowerSetup = NewSetup(map[string]string{
		"masterKey":  "43434-54545-45454-545432",
		"privateKey": "test_private_auhidaudvbirbyyrieoib",
		"publicKey":  "test_public_iopjasdioppdadipjoasd",
		"token":      "ioapdojdifouw8h",
		"mode":       "test",
	})
}

func (s *SuiteTestSetup) TestNewSetup() {
	assert.Equal(s.T(), "43434-54545-45454-545432", s.mpowerSetup.Get("MasterKey"), "Master Keys are equal")
	assert.Equal(s.T(), "test_private_auhidaudvbirbyyrieoib", s.mpowerSetup.Get("PrivateKey"), "Private Keys Keys are equal")
	assert.Equal(s.T(), "test_public_iopjasdioppdadipjoasd", s.mpowerSetup.Get("PublicKey"), "Public Keys are equal")
	assert.Equal(s.T(), "ioapdojdifouw8h", s.mpowerSetup.Get("Token"), "Tokens are equal")
	assert.Equal(s.T(), "https://app.mpowerpayments.com/sandbox-api/v1", s.mpowerSetup.Get("BASE_URL"), "Urls are correct")
}

func (s *SuiteTestSetup) TestNewSetupGetHeaders() {
	headers := s.mpowerSetup.Headers
	assert.Equal(s.T(), "43434-54545-45454-545432", headers["MP-Master-Key"], "Master Keys are equal")
	assert.Equal(s.T(), "test_private_auhidaudvbirbyyrieoib", headers["MP-Private-Key"], "Private Keys Keys are equal")
	assert.Equal(s.T(), "test_public_iopjasdioppdadipjoasd", headers["MP-Public-Key"], "Public Keys are equal")
	assert.Equal(s.T(), "ioapdojdifouw8h", headers["MP-Token"], "Tokens are equal")
	assert.Equal(s.T(), "application/json", headers["Content-Type"], "Content type is application/json")
}

func TestSetupRunSuite(t *testing.T) {
	setupTester := new(SuiteTestSetup)
	suite.Run(t, setupTester)
}
