package client

import "fmt"

type RssClient struct {
	id int
}

func (r RssClient) ID() int {
	return r.id
}

func (r RssClient) Notify(msg string) {
	fmt.Printf("RssClient with ID=%v received msg=%v", r.ID(), msg)
}

func NewRssClient(id int) RssClient {
	return RssClient{id: id}
}
