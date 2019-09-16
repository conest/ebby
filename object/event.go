package object

import "time"

// Event : 事件
type Event struct {
	Touchable   bool
	Times       int
	CoolTime    float64
	fn          func(*Event)
	lastTrigger time.Time
}

// NewEvent : 返回新事件
func NewEvent(fn func(*Event)) Event {
	return Event{
		Touchable: true,
		fn:        fn,
	}
}

// Trigger : 触发事件
func (e *Event) Trigger() {
	if !e.Touchable {
		return
	}
	if time.Since(e.lastTrigger).Seconds() < e.CoolTime {
		return
	}
	e.lastTrigger = time.Now()
	e.Times++
	e.fn(e)
}

// Events : 事件序列
type Events []Event

// Push : Push新事件，返回事件序号
func (e *Events) Push(event Event) int {
	*e = append(*e, event)
	return len(*e) - 1
}

// PushFn : Push新事件with function，返回事件序号与新的event指针
func (e *Events) PushFn(fn func(*Event)) (int, *Event) {
	event := Event{
		Touchable: true,
		fn:        fn,
	}
	*e = append(*e, event)
	return len(*e) - 1, &event
}
