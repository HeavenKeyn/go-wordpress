package wordpress

import "testing"

func TestNewClient(t *testing.T) {
	t.Log(NewClient(&Options{
		Username: "aaa",
		Password: "bbb",
	}))
}
