package gutenberg_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitGutenberg(t *testing.T) {
	suite := spec.New("Gutenberg", spec.Report(report.Terminal{}))
	suite("Gutenberg", testGutenberg)
	suite.Run(t)
}
