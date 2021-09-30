package wordpress

import "fmt"

type posts struct {
	*Client
}

const PostUrl = "posts"

//func (p *posts) Create(post Post) (int, error) {
//	return 0, nil
//}

func (p *posts) Get(id int) (Post, error) {
	var post Post
	err := getResponse(fmt.Sprintf("%s/%d", p.opt.url(PostUrl), id), &post)
	return post, err
}
