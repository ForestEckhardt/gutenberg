package gutenberg_test

import (
	"bytes"
	"testing"

	"github.com/ForestEckhardt/gutenberg"
	"github.com/cloudfoundry/packit"
	"github.com/cloudfoundry/packit/scribe"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testGutenberg(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		press  gutenberg.Press
		buffer *bytes.Buffer
		logger scribe.Logger
	)

	it.Before(func() {
		buffer = bytes.NewBuffer(nil)
		logger = scribe.NewLogger(buffer)

		press = gutenberg.NewPress()
	})

	context("PrintEnv", func() {
		context("when the operation is override", func() {
			it("prints the env in a well formatted map", func() {
				envMap := press.MapEnv(packit.Environment{
					"ENV.override": "some-value",
				})
				logger.Title(envMap.String())
				Expect(buffer.String()).To(ContainSubstring(`ENV -> "some-value"`))
			})
		})

		context("when the operation is default", func() {
			it("prints the env in a well formatted map", func() {
				envMap := press.MapEnv(packit.Environment{
					"ENV.default": "some-value",
				})
				logger.Title(envMap.String())
				Expect(buffer.String()).To(ContainSubstring(`ENV -> "some-value"`))
			})
		})

		context("when the operation is prepend", func() {
			it("prints the env in a well formatted map", func() {
				envMap := press.MapEnv(packit.Environment{
					"ENV.prepend": "some-value",
					"ENV.delim":   ":",
				})
				logger.Title(envMap.String())
				Expect(buffer.String()).To(ContainSubstring(`ENV -> "some-value:$ENV"`))
			})
		})

		context("when the operation is append", func() {
			it("prints the env in a well formatted map", func() {
				envMap := press.MapEnv(packit.Environment{
					"ENV.append": "some-value",
					"ENV.delim":  ":",
				})
				logger.Title(envMap.String())
				Expect(buffer.String()).To(ContainSubstring(`ENV -> "$ENV:some-value"`))
			})
		})
	})
}
