package wordpress

import "testing"

func TestPosts_Get(t *testing.T) {
	w := initClient()
	p, err := w.Posts().Get(62)
	if err != nil {
		t.Fatal(p)
	}
	t.Log(p)
	t.Log(p.Content)
}
