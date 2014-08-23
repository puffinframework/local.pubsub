package pubsub_test

import (
	. "github.com/puffinframework/local.pubsub"
	. "github.com/puffinframework/pubsub/impltests"
	"testing"
)

func TestRemote(t *testing.T) {
	pb := NewPubSub()
	TestSubscribe(t, pb)
	pb.Close()

	pb = NewPubSub()
	TestUnsubscribe(t, pb)
	pb.Close()

	pb = NewPubSub()
	TestSubscribeSync(t, pb)
	pb.Close()
}
