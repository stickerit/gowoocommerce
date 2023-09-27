package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Validate(secret string, xWcWebhookSignature string, b []byte) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(b)
	expectedMAC := mac.Sum(nil)

	b64 := base64.StdEncoding.EncodeToString(expectedMAC)

	return b64 == xWcWebhookSignature
}
