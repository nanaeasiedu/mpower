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
	_, s.mpowerStore = NewStore(map[string]string{
		"name":          "Awesome Store",
		"tagline":       "Easy shopping",
		"phoneNumber":   "0272271893",
		"postalAddress": "P.0. Box MP555, Accra",
		"logoURL":       "http://www.awesomestore.com.gh/logo.png",
	})
}

func (s *SuiteTestStore) TestNewSetup() {
	assert.Equal(s.T(), "Awesome Store", s.mpowerStore.Get("Name"), "Master Keys are equal")
	assert.Equal(s.T(), "Easy shopping", s.mpowerStore.Get("Tagline"), "Private Keys Keys are equal")
	assert.Equal(s.T(), "0272271893", s.mpowerStore.Get("PhoneNumber"), "Public Keys are equal")
	assert.Equal(s.T(), "P.0. Box MP555, Accra", s.mpowerStore.Get("PostalAddress"), "Tokens are equal")
	assert.Equal(s.T(), "http://www.awesomestore.com.gh/logo.png", s.mpowerStore.Get("LogoURL"), "Urls are correct")
}

func TestStoreSRunSuite(t *testing.T) {
	setupTester := new(SuiteTestStore)
	suite.Run(t, setupTester)
}
