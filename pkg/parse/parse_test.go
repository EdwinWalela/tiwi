package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildProject(t *testing.T) {
	args := []string{"../../test-site"}
	Build(args)
	assert.DirExists(t, fmt.Sprintf("%s/static", args[0]))
}
