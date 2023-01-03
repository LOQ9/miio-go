package model

import (
	"encoding/hex"
	"net"
	"sync"
	"time"

	"miio-go/common"

	"miio-go/subscription"

	"miio-go/protocol"
	"miio-go/protocol/tokens"
)

type Client struct {
	sync.RWMutex
	subscription.SubscriptionTarget

	protocol          protocol.Protocol
	discoveryInterval time.Duration
	quitChan          chan struct{}
	events            chan interface{}
}

func CreateClient(local bool, address string, token string, deviceId int) (*Client, error) {
	addr := net.IPv4bcast
	if local {
		addr = net.IPv4(127, 0, 0, 1)
	}

	if address != "" {
		addr = net.ParseIP(address)
	}

	// tokenStore, err := tokens.FromFile("tokens.txt")
	// if err != nil {
	// 	panic(err)
	// }

	parsedToken, err := hex.DecodeString(token)

	if err != nil {
		return nil, err
	}

	tokenStore := tokens.New()
	tokenStore.AddDevice(uint32(deviceId), parsedToken)

	protocolConfig := protocol.ProtocolConfig{
		BroadcastIP: addr,
		TokenStore:  tokenStore,
	}

	if deviceId == 0 {
		protocolConfig.AutoDevice = true
	}

	proto, err := protocol.NewProtocol(protocolConfig)
	if err != nil {
		return nil, err
	}

	return NewClientWithProtocol(proto)
}

// NewClient creates a new default Client with the protocol.
func NewClient() (*Client, error) {
	tokenStore, err := tokens.FromFile("tokens.txt")
	if err != nil {
		return nil, err
	}

	protocolConfig := protocol.ProtocolConfig{
		BroadcastIP: net.IPv4bcast,
		TokenStore:  tokenStore,
	}

	p, err := protocol.NewProtocol(protocolConfig)
	if err != nil {
		return nil, err
	}

	return NewClientWithProtocol(p)
}

func NewClientWithProtocol(protocol protocol.Protocol) (*Client, error) {
	c := &Client{
		SubscriptionTarget: subscription.NewTarget(),
		protocol:           protocol,
		quitChan:           make(chan struct{}),
	}

	c.SetDiscoveryInterval(time.Second * 15)

	return c, c.init()
}

func (c *Client) init() error {
	if err := c.subscribe(); err != nil {
		return err
	}
	return c.discover()
}

func (c *Client) SetDiscoveryInterval(interval time.Duration) {
	c.discoveryInterval = interval
	c.protocol.SetExpiryTime(interval * 2)
}

func (c *Client) discover() error {
	if c.discoveryInterval == 0 {
		common.Log.Debugf("Discovery interval is zero, discovery will only be performed once")
		return c.protocol.Discover()
	}

	_ = c.protocol.Discover()

	go func() {
		c.RLock()
		tick := time.Tick(c.discoveryInterval)
		c.RUnlock()
		for {
			select {
			case <-c.quitChan:
				common.Log.Debugf("Quitting discovery loop")
				return
			default:
			}
			select {
			case <-c.quitChan:
				common.Log.Debugf("Quitting discovery loop")
				return
			case <-tick:
				common.Log.Debugf("Performing discovery")
				_ = c.protocol.Discover()
			}
		}
	}()

	return nil
}

// Proxy events from protocol level
func (c *Client) subscribe() error {
	sub, err := c.protocol.NewSubscription()
	if err != nil {
		return err
	}

	go func() {
		for event := range sub.Events() {
			c.Publish(event)
		}
	}()
	return nil
}
