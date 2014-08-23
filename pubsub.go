package pubsub

import (
    . "github.com/puffinframework/pubsub"
	"time"
)

type localPubSub struct {
    localSubscriptions []*localSubscription
}

func NewPubSub() PubSub {
	return &localPubSub{}
}

func (self *localPubSub) Close() {
}

func (self *localPubSub) Subscribe(topic string, callback Callback) (Subscription, error) {
    localSubscription := &localSubscription{ localPubSub: self, topic: topic, callback: callback }
    self.localSubscriptions = append(self.localSubscriptions, localSubscription)
	return localSubscription, nil
}

func (self *localPubSub) SubscribeSync(topic string, callbackSync CallbackSync) (Subscription, error) {
    localSubscription := &localSubscription{ localPubSub: self, topic: topic, callbackSync: callbackSync }
    self.localSubscriptions = append(self.localSubscriptions, localSubscription)
	return localSubscription, nil
}

func (self *localPubSub) Publish(topic string, data []byte) error {
    for _, localSubscription := range self.localSubscriptions {
        if localSubscription.topic == topic {
            localSubscription.callback(data)
        }
    }
    return nil
}

func (self *localPubSub) PublishSync(topic string, data []byte, timeout time.Duration) (result []byte, err error) {
    for _, localSubscription := range self.localSubscriptions {
        if localSubscription.topic == topic {
            if localSubscription.callbackSync != nil {
                result, err = localSubscription.callbackSync(data)
            }
            if localSubscription.callback != nil {
                localSubscription.callback(data)
            }
        }
    }
    return result, err
}

type localSubscription struct {
    localPubSub *localPubSub
    topic string
    callback Callback
    callbackSync CallbackSync
}

func (self *localSubscription) Unsubscribe() error {
    for i, localSubscription := range self.localPubSub.localSubscriptions {
        if localSubscription == self {
            self.localPubSub.localSubscriptions = append(self.localPubSub.localSubscriptions[:i], self.localPubSub.localSubscriptions[i+1:]...)
            return nil
        }
    }
    return nil
}
