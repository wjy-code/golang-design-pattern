package option

import (
	"testing"
)

func TestXxx(t *testing.T) {
	userDefault := NewUser()
	t.Logf("user: %T", userDefault)

	user := NewUser(WithName("wangjy"), WithAge(28))
	t.Logf("user: %v", user)
}
