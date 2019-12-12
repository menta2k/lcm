package subscribers

import (
	"github.com/lileio/pubsub/v2"
)

type LcmServiceSubscriber struct{}

func (s *LcmServiceSubscriber) Setup(c *pubsub.Client) {
	// https://godoc.org/github.com/lileio/pubsub#Client.On
	// c.On(pubsub.HandlerOptions{
	// 	Topic:    "some_topic",
	// 	Name:     "service_name",
	// 	Handler:  s.SomeMethod,
	// 	Deadline: 10 * time.Second,
	// 	Concurrency: 10,
	// 	AutoAck:  true,
	// })
}

// func (s *LcmServiceSubscriber) SomeMethod(ctx context.Context, req *proto.Message, _ *pubsub.Msg) error {
// 	return nil
// }
