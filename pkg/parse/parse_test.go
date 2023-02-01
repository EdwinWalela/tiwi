package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPages(t *testing.T) {
	pages, err := getPages("../../test-site")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(pages))
}

func TestCreateOutputFolder(t *testing.T) {
	directory := "../../test-site"
	err := createOutputFolder(directory)

	assert.NoError(t, err)
	assert.DirExists(t, directory+"/static")
}

func TestReadMarkDown(t *testing.T) {

}

func TestParseMd(t *testing.T) {

}

func TestWriteHTML(t *testing.T) {

}

func TestBuildProject(t *testing.T) {
	args := []string{"../../test-site"}
	Build(args)
	assert.DirExists(t, fmt.Sprintf("%s/static", args[0]))
}
