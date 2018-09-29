package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
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
			path, err := repo.Clone(
				testCase.url,
				testCase.branch,
				testCase.dstDir,
				testCase.username,
				testCase.password,
			)

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

	_, err := repo.Clone("", "","","","")
	assert.Equal(t, err, fmt.Errorf("clone faile"))
}