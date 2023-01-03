package subscription

import (
	"miio-go/subscription/common"
	"miio-go/subscription/target"
)

func NewTarget() common.SubscriptionTarget {
	return target.NewTarget()
}

type SubscriptionTarget = common.SubscriptionTarget
type Subscription = common.Subscription
