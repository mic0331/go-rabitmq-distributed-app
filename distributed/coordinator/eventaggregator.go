package coordinator

import "time"

type EventAggregator struct {
	listeners map[string][]func(EventData)
}

func NewEventAggregator() *EventAggregator {
	ea := EventAggregator{
		listeners: make(map[string][]func(EventData)),
	}
	return &ea
}

func (ae *EventAggregator) AddLitener(name string, f func(EventData)) {
	ae.listeners[name] = append(ae.listeners[name], f)
}

func (ae *EventAggregator) PublishEvent(name string, eventData EventData) {
	if ae.listeners[name] != nil {
		for _, r := range ae.listeners[name] {
			r(eventData)
		}
	}
}

type EventData struct {
	Name      string
	Value     float64
	Timestamp time.Time
}
