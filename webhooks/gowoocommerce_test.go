package webhooks

import (
	"testing"
)

func TestValidate(t *testing.T) {
	if Validate("secret", "this is my signature", []byte("some data")) {
		t.Fail()
	}

	if !Validate("secret", "ov+1ADPF9Fbr081JlqnG0CgxuM0Tg0oSt9ODpTs2nZ8=", []byte("some data")) {
		t.Fail()
	}
}
