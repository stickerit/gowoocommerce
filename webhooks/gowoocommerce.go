package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Signature(secret string, b []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(b)
	expectedMAC := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(expectedMAC)
}

func Validate(secret string, xWcWebhookSignature string, b []byte) bool {
	b64 := Signature(secret, b)

	return b64 == xWcWebhookSignature
}
