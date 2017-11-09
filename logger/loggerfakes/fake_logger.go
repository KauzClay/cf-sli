// Code generated by counterfeiter. DO NOT EDIT.
package loggerfakes

import (
	"sync"

	"github.com/pivotal-cloudops/cf-sli/logger"
)

type FakeLogger struct {
	PrintfStub        func(format string, v ...interface{})
	printfMutex       sync.RWMutex
	printfArgsForCall []struct {
		format string
		v      []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Printf(format string, v ...interface{}) {
	fake.printfMutex.Lock()
	fake.printfArgsForCall = append(fake.printfArgsForCall, struct {
		format string
		v      []interface{}
	}{format, v})
	fake.recordInvocation("Printf", []interface{}{format, v})
	fake.printfMutex.Unlock()
	if fake.PrintfStub != nil {
		fake.PrintfStub(format, v...)
	}
}

func (fake *FakeLogger) PrintfCallCount() int {
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	return len(fake.printfArgsForCall)
}

func (fake *FakeLogger) PrintfArgsForCall(i int) (string, []interface{}) {
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	return fake.printfArgsForCall[i].format, fake.printfArgsForCall[i].v
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.printfMutex.RLock()
	defer fake.printfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ logger.Logger = new(FakeLogger)
