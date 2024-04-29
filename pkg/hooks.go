package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

type BellRingEvent struct {
	Volume uint8
	Sound  string
	Path   string
}

type HandlerFunc[T any] func(e T) error
type Handler[T any] struct {
	id      string
	handler HandlerFunc[T]
}

type Hook[T any] struct {
	event    T
	handlers []*Handler[T]
}

func (h *Hook[T]) Add(f HandlerFunc[T]) string {
	id := RandStr(16)
	h.handlers = append(h.handlers, &Handler[T]{
		id:      id,
		handler: f,
	})
	return id
}

func (h *Hook[T]) Remove(id string) {
	for i := len(h.handlers) - 1; i >= 0; i-- {
		if h.handlers[i].id == id {
			h.handlers = append(h.handlers[:i], h.handlers[i+1:]...)
			return
		}
	}
}

func (h *Hook[T]) RemoveAll() {
	h.handlers = make([]*Handler[T], 0)
}

func (h *Hook[T]) Trigger(e T) error {
	for i := range h.handlers {
		err := h.handlers[i].handler(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func RandStr(size int) string {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
