// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeTracker struct {
	InitStub        func(lager.Logger, resource.Metadata, resource.Session, resource.ResourceType, atc.Tags, atc.ResourceTypes, worker.ImageFetchingDelegate) (resource.Resource, error)
	initMutex       sync.RWMutex
	initArgsForCall []struct {
		arg1 lager.Logger
		arg2 resource.Metadata
		arg3 resource.Session
		arg4 resource.ResourceType
		arg5 atc.Tags
		arg6 atc.ResourceTypes
		arg7 worker.ImageFetchingDelegate
	}
	initReturns struct {
		result1 resource.Resource
		result2 error
	}
	InitWithSourcesStub        func(lager.Logger, resource.Metadata, resource.Session, resource.ResourceType, atc.Tags, map[string]resource.ArtifactSource, atc.ResourceTypes, worker.ImageFetchingDelegate) (resource.Resource, []string, error)
	initWithSourcesMutex       sync.RWMutex
	initWithSourcesArgsForCall []struct {
		arg1 lager.Logger
		arg2 resource.Metadata
		arg3 resource.Session
		arg4 resource.ResourceType
		arg5 atc.Tags
		arg6 map[string]resource.ArtifactSource
		arg7 atc.ResourceTypes
		arg8 worker.ImageFetchingDelegate
	}
	initWithSourcesReturns struct {
		result1 resource.Resource
		result2 []string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTracker) Init(arg1 lager.Logger, arg2 resource.Metadata, arg3 resource.Session, arg4 resource.ResourceType, arg5 atc.Tags, arg6 atc.ResourceTypes, arg7 worker.ImageFetchingDelegate) (resource.Resource, error) {
	fake.initMutex.Lock()
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
		arg1 lager.Logger
		arg2 resource.Metadata
		arg3 resource.Session
		arg4 resource.ResourceType
		arg5 atc.Tags
		arg6 atc.ResourceTypes
		arg7 worker.ImageFetchingDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Init", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	} else {
		return fake.initReturns.result1, fake.initReturns.result2
	}
}

func (fake *FakeTracker) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeTracker) InitArgsForCall(i int) (lager.Logger, resource.Metadata, resource.Session, resource.ResourceType, atc.Tags, atc.ResourceTypes, worker.ImageFetchingDelegate) {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return fake.initArgsForCall[i].arg1, fake.initArgsForCall[i].arg2, fake.initArgsForCall[i].arg3, fake.initArgsForCall[i].arg4, fake.initArgsForCall[i].arg5, fake.initArgsForCall[i].arg6, fake.initArgsForCall[i].arg7
}

func (fake *FakeTracker) InitReturns(result1 resource.Resource, result2 error) {
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeTracker) InitWithSources(arg1 lager.Logger, arg2 resource.Metadata, arg3 resource.Session, arg4 resource.ResourceType, arg5 atc.Tags, arg6 map[string]resource.ArtifactSource, arg7 atc.ResourceTypes, arg8 worker.ImageFetchingDelegate) (resource.Resource, []string, error) {
	fake.initWithSourcesMutex.Lock()
	fake.initWithSourcesArgsForCall = append(fake.initWithSourcesArgsForCall, struct {
		arg1 lager.Logger
		arg2 resource.Metadata
		arg3 resource.Session
		arg4 resource.ResourceType
		arg5 atc.Tags
		arg6 map[string]resource.ArtifactSource
		arg7 atc.ResourceTypes
		arg8 worker.ImageFetchingDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("InitWithSources", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.initWithSourcesMutex.Unlock()
	if fake.InitWithSourcesStub != nil {
		return fake.InitWithSourcesStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	} else {
		return fake.initWithSourcesReturns.result1, fake.initWithSourcesReturns.result2, fake.initWithSourcesReturns.result3
	}
}

func (fake *FakeTracker) InitWithSourcesCallCount() int {
	fake.initWithSourcesMutex.RLock()
	defer fake.initWithSourcesMutex.RUnlock()
	return len(fake.initWithSourcesArgsForCall)
}

func (fake *FakeTracker) InitWithSourcesArgsForCall(i int) (lager.Logger, resource.Metadata, resource.Session, resource.ResourceType, atc.Tags, map[string]resource.ArtifactSource, atc.ResourceTypes, worker.ImageFetchingDelegate) {
	fake.initWithSourcesMutex.RLock()
	defer fake.initWithSourcesMutex.RUnlock()
	return fake.initWithSourcesArgsForCall[i].arg1, fake.initWithSourcesArgsForCall[i].arg2, fake.initWithSourcesArgsForCall[i].arg3, fake.initWithSourcesArgsForCall[i].arg4, fake.initWithSourcesArgsForCall[i].arg5, fake.initWithSourcesArgsForCall[i].arg6, fake.initWithSourcesArgsForCall[i].arg7, fake.initWithSourcesArgsForCall[i].arg8
}

func (fake *FakeTracker) InitWithSourcesReturns(result1 resource.Resource, result2 []string, result3 error) {
	fake.InitWithSourcesStub = nil
	fake.initWithSourcesReturns = struct {
		result1 resource.Resource
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTracker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.initWithSourcesMutex.RLock()
	defer fake.initWithSourcesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTracker) recordInvocation(key string, args []interface{}) {
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

var _ resource.Tracker = new(FakeTracker)
