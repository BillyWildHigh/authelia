package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ActiveDirectorySuite struct {
	*RodSuite
}

func NewActiveDirectorySuite() *ActiveDirectorySuite {
	return &ActiveDirectorySuite{RodSuite: new(RodSuite)}
}

func (s *ActiveDirectorySuite) TestOneFactorScenario() {
	suite.Run(s.T(), NewOneFactorScenario())
}

func (s *ActiveDirectorySuite) TestTwoFactorScenario() {
	suite.Run(s.T(), NewTwoFactorScenario())
}

func (s *ActiveDirectorySuite) TestResetPassword() {
	suite.Run(s.T(), NewResetPasswordScenario())
}

func (s *ActiveDirectorySuite) TestPasswordComplexity() {
	suite.Run(s.T(), NewPasswordComplexityScenario())
}

func (s *ActiveDirectorySuite) TestSigninEmailScenario() {
	suite.Run(s.T(), NewSigninEmailScenario())
}

func TestActiveDirectorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewActiveDirectorySuite())
}
