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
		"masterKey":  "55647970-22e1-4e7e-8fb4-56eca2b3b006",
		"privateKey": "test_private_B8EiE1AGWpb4tVMzVTyFDu9rYoc",
		"publicKey":  "test_public_B1wo2UVmxUrvwzZuPqpLrWqlA74",
		"token":      "a6d96e2586c8bbae7c28",
		"mode":       "test",
	})
}

func (s *SuiteTestSetup) TestNewSetup() {
	assert.Equal(s.T(), "55647970-22e1-4e7e-8fb4-56eca2b3b006", s.mpowerSetup.Get("MasterKey"), "Master Keys are equal")
	assert.Equal(s.T(), "test_private_B8EiE1AGWpb4tVMzVTyFDu9rYoc", s.mpowerSetup.Get("PrivateKey"), "Private Keys Keys are equal")
	assert.Equal(s.T(), "test_public_B1wo2UVmxUrvwzZuPqpLrWqlA74", s.mpowerSetup.Get("PublicKey"), "Public Keys are equal")
	assert.Equal(s.T(), "a6d96e2586c8bbae7c28", s.mpowerSetup.Get("Token"), "Tokens are equal")
	assert.Equal(s.T(), "https://app.mpowerpayments.com/sandbox-api/v1", s.mpowerSetup.Get("BaseURL"), "Urls are correct")
}

func (s *SuiteTestSetup) TestNewSetupGetHeaders() {
	headers := s.mpowerSetup.GetHeaders()
	assert.Equal(s.T(), "55647970-22e1-4e7e-8fb4-56eca2b3b006", headers["MP-Master-Key"], "Master Keys are equal")
	assert.Equal(s.T(), "test_private_B8EiE1AGWpb4tVMzVTyFDu9rYoc", headers["MP-Private-Key"], "Private Keys Keys are equal")
	assert.Equal(s.T(), "test_public_B1wo2UVmxUrvwzZuPqpLrWqlA74", headers["MP-Public-Key"], "Public Keys are equal")
	assert.Equal(s.T(), "a6d96e2586c8bbae7c28", headers["MP-Token"], "Tokens are equal")
	assert.Equal(s.T(), "application/json", headers["Content-Type"], "Content type is application/json")
}

func TestSetupRunSuite(t *testing.T) {
	setupTester := new(SuiteTestSetup)
	suite.Run(t, setupTester)
}
