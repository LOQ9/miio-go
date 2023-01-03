# miio-go

**Warning** This is a heavily modified version of miio-go!  

Forked from https://github.com/nickw444/miio-go which is no longer maintained  

## Example

Query information about the product "zhimi.airpurifier.v7"

```
miio-go control --address IP_ADDRESS --token TOKEN airpurifier
```

```
{
  "aqi": 13,
  "average_aqi": 3,
  "f1_hour": 3500,
  "f1_hour_used": 11034,
  "favorite_level": 14,
  "filter1_life": 0,
  "humidity": 79,
  "led": "on",
  "mode": "silent",
  "motor1_speed": 0,
  "motor2_speed": null,
  "power": "off",
  "purify_volume": 273753,
  "temp_dec": 159,
  "use_time": 39644100
}
```