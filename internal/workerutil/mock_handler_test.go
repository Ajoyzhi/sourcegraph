// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package workerutil

import (
	"context"
	"sync"
)

// MockHandler is a mock implementation of the Handler interface (from the
// package github.com/sourcegraph/sourcegraph/internal/workerutil) used for
// unit testing.
type MockHandler struct {
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *HandlerHandleFunc
}

// NewMockHandler creates a new mock of the Handler interface. All methods
// return zero values for all results, unless overwritten.
func NewMockHandler() *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: func(context.Context, Store, Record) error {
				return nil
			},
		},
	}
}

// NewMockHandlerFrom creates a new mock of the MockHandler interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockHandlerFrom(i Handler) *MockHandler {
	return &MockHandler{
		HandleFunc: &HandlerHandleFunc{
			defaultHook: i.Handle,
		},
	}
}

// HandlerHandleFunc describes the behavior when the Handle method of the
// parent MockHandler instance is invoked.
type HandlerHandleFunc struct {
	defaultHook func(context.Context, Store, Record) error
	hooks       []func(context.Context, Store, Record) error
	history     []HandlerHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) Handle(v0 context.Context, v1 Store, v2 Record) error {
	r0 := m.HandleFunc.nextHook()(v0, v1, v2)
	m.HandleFunc.appendCall(HandlerHandleFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerHandleFunc) SetDefaultHook(hook func(context.Context, Store, Record) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockHandler instance inovkes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *HandlerHandleFunc) PushHook(hook func(context.Context, Store, Record) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, Store, Record) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, Store, Record) error {
		return r0
	})
}

func (f *HandlerHandleFunc) nextHook() func(context.Context, Store, Record) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerHandleFunc) appendCall(r0 HandlerHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerHandleFuncCall objects describing
// the invocations of this function.
func (f *HandlerHandleFunc) History() []HandlerHandleFuncCall {
	f.mutex.Lock()
	history := make([]HandlerHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerHandleFuncCall is an object that describes an invocation of method
// Handle on an instance of MockHandler.
type HandlerHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 Store
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 Record
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
