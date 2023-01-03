package capability

import (
	"encoding/json"
	"fmt"

	"miio-go/common"
	"miio-go/subscription"

	"miio-go/protocol/transport"
)

type AirPurifierV7 struct {
	subscriptionTarget subscription.SubscriptionTarget
	outbound           transport.Outbound
}

type AirPurifierV7Response struct {
	Result []common.PowerState `json:"result"`
}

func NewAirPurifierV7(target subscription.SubscriptionTarget, transport transport.Outbound) *AirPurifierV7 {
	return &AirPurifierV7{
		subscriptionTarget: target,
		outbound:           transport,
	}
}

func (p *AirPurifierV7) Info() error {
	var resp transport.Response
	props := []string{"power", "aqi", "average_aqi", "humidity", "temp_dec", "mode", "favorite_level", "filter1_life", "f1_hour_used", "use_time", "motor1_speed", "motor2_speed", "purify_volume", "f1_hour", "led"}

	err := p.outbound.CallAndDeserialize("get_prop", props, &resp)
	if err != nil {
		return err
	}

	// Combine props with result
	resultList := map[string]interface{}{}
	for i, result := range resp.Result.([]interface{}) {
		propName := props[i]
		resultList[propName] = result
	}

	// Print results
	jsonResponse, err := json.Marshal(resultList)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonResponse))

	return nil
}
