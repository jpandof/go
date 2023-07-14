package main

import "sync"

type PostDatabase struct {
	sync.Mutex
	posts map[string]*BlogPost
}

func NewPostDatabase() *PostDatabase {
	return &PostDatabase{
		posts: make(map[string]*BlogPost),
	}
}

type CreatePostCommand struct {
	Post  *BlogPost
	Store *PostDatabase
}

func (c *CreatePostCommand) Execute() {
	c.Store.Lock()
	defer c.Store.Unlock()
	c.Store.posts[c.Post.ID] = c.Post
}
