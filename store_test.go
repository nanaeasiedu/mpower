package mpower

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteTestStore struct {
	suite.Suite
	mpowerStore *Store
}

func (s *SuiteTestStore) SetupSuite() {
	s.mpowerStore = NewStore("Awesome Store", "Easy shopping", "0272271893", "P.0. Box MP555, Accra", "http://www.awesomestore.com.gh/logo.png")
}

func (s *SuiteTestStore) TestNewSetup() {
	assert.Equal(s.T(), "Awesome Store", s.mpowerStore.Name, "Master Keys are equal")
	assert.Equal(s.T(), "Easy shopping", s.mpowerStore.Tagline, "Private Keys Keys are equal")
	assert.Equal(s.T(), "0272271893", s.mpowerStore.PhoneNumber, "Public Keys are equal")
	assert.Equal(s.T(), "P.0. Box MP555, Accra", s.mpowerStore.PostalAddress, "Tokens are equal")
	assert.Equal(s.T(), "http://www.awesomestore.com.gh/logo.png", s.mpowerStore.LogoURL, "Urls are correct")
}

func TestStoreRunSuite(t *testing.T) {
	setupTester := new(SuiteTestStore)
	suite.Run(t, setupTester)
}
