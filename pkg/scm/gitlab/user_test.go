package gitlab

import (
	"testing"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestUser_GetUser(t *testing.T) {
	token := os.Getenv("SCM_TOKEN")
	baseUrl := os.Getenv("SCM_URL")
	log.Debugf("accessToken: %v", token)
	user := new(User)
	u, err := user.GetUser(baseUrl, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "chulei", u.Username)
}
