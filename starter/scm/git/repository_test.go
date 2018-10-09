package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"testing"
)

func TestClone(t *testing.T) {
	repo := NewRepository(func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
		return nil, nil
	})

	testCases := []struct {
		title        string
		url          string
		branch       string
		dstDir       string
		username     string
		password     string
		expectedPath string
		err          error
	}{
		{
			title:        "should return file path correctly",
			url:          "https://github.com/hidevopsio/hiboot.git",
			branch:       "master",
			dstDir:       "/path/to/dir",
			expectedPath: "/path/to/dir/hiboot",
			err:          nil,
		},
		{
			title:        "should err is nil",
			url:          "https://github.com/hidevopsio/hiboot",
			branch:       "master",
			dstDir:       "/path/to/dir/",
			expectedPath: "/path/to/dir/hiboot",
			err:          nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.title, func(t *testing.T) {
			path, err := repo.Clone(&git.CloneOptions{URL:testCase.url,
			ReferenceName:plumbing.Master,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Auth:transport.AuthMethod(&http.BasicAuth{
				Username: testCase.username,
				Password: testCase.password})},testCase.dstDir)
			assert.Equal(t, testCase.expectedPath, path)
			assert.Equal(t, testCase.err, err)
		})
	}

}

//should return err
func TestCloneErr(t *testing.T) {
	repo := NewRepository(func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
		return nil, fmt.Errorf("clone faile")
	})

	_, err := repo.Clone(&git.CloneOptions{}, "")
	assert.Equal(t, err, fmt.Errorf("clone faile"))
}
