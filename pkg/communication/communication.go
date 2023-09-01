package communication

type Channels struct {
	messages chan mqtt
}

type mqtt struct {
	payload string
}
