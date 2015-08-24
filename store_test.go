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
	s.mpowerStore = NewStore("Awesome Store")
}

func (s *SuiteTestStore) TestNewSetup() {
	assert.Equal(s.T(), "Awesome Store", s.mpowerStore.Name, "Master Keys are equal")
}

func TestStoreRunSuite(t *testing.T) {
	setupTester := new(SuiteTestStore)
	suite.Run(t, setupTester)
}
