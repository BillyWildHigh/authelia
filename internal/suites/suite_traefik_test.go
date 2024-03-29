package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TraefikSuite struct {
	*RodSuite
}

func NewTraefikSuite() *TraefikSuite {
	return &TraefikSuite{RodSuite: new(RodSuite)}
}

func (s *TraefikSuite) TestOneFactorScenario() {
	suite.Run(s.T(), NewOneFactorScenario())
}

func (s *TraefikSuite) TestTwoFactorScenario() {
	suite.Run(s.T(), NewTwoFactorScenario())
}

func (s *TraefikSuite) TestRedirectionURLScenario() {
	suite.Run(s.T(), NewRedirectionURLScenario())
}

func (s *TraefikSuite) TestCustomHeaders() {
	suite.Run(s.T(), NewCustomHeadersScenario())
}

func TestTraefikSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewTraefikSuite())
}
