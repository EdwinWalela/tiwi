package create

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	args := []string{"../../test-site"}
	CreateSite(args)
	assert.DirExists(t, args[0])
}
