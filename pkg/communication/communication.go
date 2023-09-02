package communication

type Channels struct {
	Messages chan Mqtt
}

type Mqtt struct {
	Topic string
	Msg   string
}
