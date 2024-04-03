package event

// Producer is an interface for producing events to the event bus
type Producer interface {
	// Publish publishes an event to the event bus
	Publish(event Event) error
}

// Consumer is an interface for consuming events from the event bus
type Consumer interface {
	// Consume consumes an event from the event bus
	Consume() (Event, error)
}

// EventBus is an interface for an event bus
type EventBus interface {
	Producer
	Consumer
}

// Event is a generic event
type Event struct {
	// Type is the type of the event
	Type string `json:"type"`
	// Payload is the payload of the event
	Payload interface{} `json:"payload"`
}

// NewEvent returns a new event
func NewEvent(eventType string, payload interface{}) Event {
	return Event{
		Type:    eventType,
		Payload: payload,
	}
}
