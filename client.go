package wordpress

type Client struct {
	opt *Options
}

type Options struct {
	Username, Password string
}

func NewClient(opt *Options) *Client {
	return &Client{opt: opt}
}
