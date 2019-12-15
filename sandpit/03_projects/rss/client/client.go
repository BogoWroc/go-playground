package client

type RssClient struct {
	id int
}

func (r RssClient) ID() int {
	return r.id
}

func NewRssClient(id int) RssClient {
	return RssClient{id: id}
}
