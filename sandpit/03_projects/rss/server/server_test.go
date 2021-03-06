package server

import (
	client2 "github.com/bogowroc/go-playground/sandpit/03_projects/rss/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeClient struct {
	id  int
	msg string
}

func (f *fakeClient) Notify(msg string) {
	f.msg = msg
}
func (f fakeClient) ID() int {
	return f.id
}

var _ = Describe("Hash tests", func() {

	var server RssServer

	BeforeEach(func() {
		server = NewRssServer()
	})

	It("should inform that client was not registered", func() {
		c := client2.NewRssClient(1)

		Expect(server.IsRegistered(c)).To(Equal(false))
	})

	It("should be possible to register client", func() {
		c := client2.NewRssClient(1)

		server.Register(c)

		Expect(server.IsRegistered(c)).To(Equal(true))
	})

	It("should be possible to register few clients", func() {
		c1 := client2.NewRssClient(1)
		c2 := client2.NewRssClient(2)

		server.Register(c1)
		server.Register(c2)

		Expect(server.IsRegistered(c1)).To(Equal(true))
		Expect(server.IsRegistered(c2)).To(Equal(true))
	})

	It("should not register client with the same ID once again", func() {
		c1 := client2.NewRssClient(1)
		c2 := client2.NewRssClient(1)

		server.Register(c1)
		server.Register(c2)

		Expect(len(server.clients)).To(Equal(1))
	})

	It("should be possible to send message to registered clients", func() {
		c := &fakeClient{id: 5}

		server.Register(c)
		server.broadcast("Super news!")

		Expect(c.msg).To(Equal("Super news!"))

	})
})
