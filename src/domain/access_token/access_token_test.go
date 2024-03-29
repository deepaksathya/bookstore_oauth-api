package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")

	assert.True(t, at.UserID == 0, "new access token should not have associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "access tokem created with three hours expiry should NOT be expired")

}
