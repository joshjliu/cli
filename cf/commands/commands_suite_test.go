package commands_test

import (
	"github.com/cloudfoundry/cli/cf/commands"
	"github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCommands(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = commands.API{}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}

type passingRequirement struct {
	Name string
}

func (r passingRequirement) Execute() error {
	return nil
}
