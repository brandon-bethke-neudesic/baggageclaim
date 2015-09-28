// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/baggageclaim"
	"github.com/pivotal-golang/lager"
)

type FakeClient struct {
	CreateVolumeStub        func(lager.Logger, baggageclaim.VolumeSpec) (baggageclaim.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 baggageclaim.VolumeSpec
	}
	createVolumeReturns struct {
		result1 baggageclaim.Volume
		result2 error
	}
	ListVolumesStub        func(lager.Logger, baggageclaim.VolumeProperties) (baggageclaim.Volumes, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 lager.Logger
		arg2 baggageclaim.VolumeProperties
	}
	listVolumesReturns struct {
		result1 baggageclaim.Volumes
		result2 error
	}
	LookupVolumeStub        func(lager.Logger, string) (baggageclaim.Volume, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 baggageclaim.Volume
		result2 error
	}
}

func (fake *FakeClient) CreateVolume(arg1 lager.Logger, arg2 baggageclaim.VolumeSpec) (baggageclaim.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 baggageclaim.VolumeSpec
	}{arg1, arg2})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeClient) CreateVolumeArgsForCall(i int) (lager.Logger, baggageclaim.VolumeSpec) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].arg1, fake.createVolumeArgsForCall[i].arg2
}

func (fake *FakeClient) CreateVolumeReturns(result1 baggageclaim.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListVolumes(arg1 lager.Logger, arg2 baggageclaim.VolumeProperties) (baggageclaim.Volumes, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 lager.Logger
		arg2 baggageclaim.VolumeProperties
	}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2)
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeClient) ListVolumesArgsForCall(i int) (lager.Logger, baggageclaim.VolumeProperties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].arg1, fake.listVolumesArgsForCall[i].arg2
}

func (fake *FakeClient) ListVolumesReturns(result1 baggageclaim.Volumes, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 baggageclaim.Volumes
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) LookupVolume(arg1 lager.Logger, arg2 string) (baggageclaim.Volume, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2
	}
}

func (fake *FakeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeClient) LookupVolumeReturns(result1 baggageclaim.Volume, result2 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

var _ baggageclaim.Client = new(FakeClient)
