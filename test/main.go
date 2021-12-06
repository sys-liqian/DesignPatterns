package main

import "sync"

type Broker interface {
	//推送消息
	publish(topicName string, msg interface{}) error
	//订阅主题
	subScribe(topicName string) (<-chan interface{}, error)
	//取消订阅
	unSubScribe(topicName string, sub <-chan interface{}) error
	//广播
	broadCast(msg interface{}, subs []chan interface{})
	//设置条件
	setConditions(cap int)
	//关闭
	close()
}

type BrokerImpl struct {
	exit   chan bool
	cap    int
	topics map[string][]chan interface{}
	sync.RWMutex
}

func newBrokerImpl() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}

type Client struct {
	brokerImpl *BrokerImpl
}

func NewClient() *Client {
	return &Client{brokerImpl: newBrokerImpl()}
}
func (c *Client) SetConditions(capacity int) {
	c.brokerImpl.cap = capacity
}

func main() {

}
