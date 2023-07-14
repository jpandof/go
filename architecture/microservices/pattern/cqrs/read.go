package main

type GetPostQuery struct {
	ID    string
	Store *PostDatabase
}

func (q *GetPostQuery) Execute() interface{} {
	q.Store.Lock()
	defer q.Store.Unlock()
	return q.Store.posts[q.ID]
}
