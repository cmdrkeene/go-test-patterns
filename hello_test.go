import "testing"

func TestGreeting(t *testing.T) {
	if greeting() != "hello world" {
		t.Fail()
	}
}
