package wordpress

import "fmt"

type Client struct {
	opt   *Options
	posts *posts
}

type Options struct {
	Username, Password, BasicUrl string
}

func (o Options) url(a interface{}) string {
	return fmt.Sprint(o.BasicUrl, "/wp-json/wp/v2/", a)
}

func NewClient(opt *Options) *Client {
	return &Client{opt: opt}
}

func (c *Client) Posts() *posts {
	if c.posts == nil {
		c.posts = new(posts)
		c.posts.Client = c
	}
	return c.posts
}
