// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	"sync"
	"time"

	"github.com/concourse/baggageclaim/volume"
)

type FakeFilesystemLiveVolume struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	DataPathStub        func() string
	dataPathMutex       sync.RWMutex
	dataPathArgsForCall []struct{}
	dataPathReturns     struct {
		result1 string
	}
	dataPathReturnsOnCall map[int]struct {
		result1 string
	}
	LoadPropertiesStub        func() (volume.Properties, error)
	loadPropertiesMutex       sync.RWMutex
	loadPropertiesArgsForCall []struct{}
	loadPropertiesReturns     struct {
		result1 volume.Properties
		result2 error
	}
	loadPropertiesReturnsOnCall map[int]struct {
		result1 volume.Properties
		result2 error
	}
	StorePropertiesStub        func(volume.Properties) error
	storePropertiesMutex       sync.RWMutex
	storePropertiesArgsForCall []struct {
		arg1 volume.Properties
	}
	storePropertiesReturns struct {
		result1 error
	}
	storePropertiesReturnsOnCall map[int]struct {
		result1 error
	}
	LoadTTLStub        func() (volume.TTL, time.Time, error)
	loadTTLMutex       sync.RWMutex
	loadTTLArgsForCall []struct{}
	loadTTLReturns     struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}
	loadTTLReturnsOnCall map[int]struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}
	StoreTTLStub        func(volume.TTL) (time.Time, error)
	storeTTLMutex       sync.RWMutex
	storeTTLArgsForCall []struct {
		arg1 volume.TTL
	}
	storeTTLReturns struct {
		result1 time.Time
		result2 error
	}
	storeTTLReturnsOnCall map[int]struct {
		result1 time.Time
		result2 error
	}
	LoadPrivilegedStub        func() (bool, error)
	loadPrivilegedMutex       sync.RWMutex
	loadPrivilegedArgsForCall []struct{}
	loadPrivilegedReturns     struct {
		result1 bool
		result2 error
	}
	loadPrivilegedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	StorePrivilegedStub        func(bool) error
	storePrivilegedMutex       sync.RWMutex
	storePrivilegedArgsForCall []struct {
		arg1 bool
	}
	storePrivilegedReturns struct {
		result1 error
	}
	storePrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	ParentStub        func() (volume.FilesystemLiveVolume, bool, error)
	parentMutex       sync.RWMutex
	parentArgsForCall []struct{}
	parentReturns     struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}
	parentReturnsOnCall map[int]struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct{}
	destroyReturns     struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	SizeInBytesStub        func() (int64, error)
	sizeInBytesMutex       sync.RWMutex
	sizeInBytesArgsForCall []struct{}
	sizeInBytesReturns     struct {
		result1 int64
		result2 error
	}
	sizeInBytesReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	NewSubvolumeStub        func(handle string) (volume.FilesystemInitVolume, error)
	newSubvolumeMutex       sync.RWMutex
	newSubvolumeArgsForCall []struct {
		handle string
	}
	newSubvolumeReturns struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}
	newSubvolumeReturnsOnCall map[int]struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilesystemLiveVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handleReturns.result1
}

func (fake *FakeFilesystemLiveVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) DataPath() string {
	fake.dataPathMutex.Lock()
	ret, specificReturn := fake.dataPathReturnsOnCall[len(fake.dataPathArgsForCall)]
	fake.dataPathArgsForCall = append(fake.dataPathArgsForCall, struct{}{})
	fake.recordInvocation("DataPath", []interface{}{})
	fake.dataPathMutex.Unlock()
	if fake.DataPathStub != nil {
		return fake.DataPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dataPathReturns.result1
}

func (fake *FakeFilesystemLiveVolume) DataPathCallCount() int {
	fake.dataPathMutex.RLock()
	defer fake.dataPathMutex.RUnlock()
	return len(fake.dataPathArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) DataPathReturns(result1 string) {
	fake.DataPathStub = nil
	fake.dataPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) DataPathReturnsOnCall(i int, result1 string) {
	fake.DataPathStub = nil
	if fake.dataPathReturnsOnCall == nil {
		fake.dataPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dataPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) LoadProperties() (volume.Properties, error) {
	fake.loadPropertiesMutex.Lock()
	ret, specificReturn := fake.loadPropertiesReturnsOnCall[len(fake.loadPropertiesArgsForCall)]
	fake.loadPropertiesArgsForCall = append(fake.loadPropertiesArgsForCall, struct{}{})
	fake.recordInvocation("LoadProperties", []interface{}{})
	fake.loadPropertiesMutex.Unlock()
	if fake.LoadPropertiesStub != nil {
		return fake.LoadPropertiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadPropertiesReturns.result1, fake.loadPropertiesReturns.result2
}

func (fake *FakeFilesystemLiveVolume) LoadPropertiesCallCount() int {
	fake.loadPropertiesMutex.RLock()
	defer fake.loadPropertiesMutex.RUnlock()
	return len(fake.loadPropertiesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) LoadPropertiesReturns(result1 volume.Properties, result2 error) {
	fake.LoadPropertiesStub = nil
	fake.loadPropertiesReturns = struct {
		result1 volume.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) LoadPropertiesReturnsOnCall(i int, result1 volume.Properties, result2 error) {
	fake.LoadPropertiesStub = nil
	if fake.loadPropertiesReturnsOnCall == nil {
		fake.loadPropertiesReturnsOnCall = make(map[int]struct {
			result1 volume.Properties
			result2 error
		})
	}
	fake.loadPropertiesReturnsOnCall[i] = struct {
		result1 volume.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) StoreProperties(arg1 volume.Properties) error {
	fake.storePropertiesMutex.Lock()
	ret, specificReturn := fake.storePropertiesReturnsOnCall[len(fake.storePropertiesArgsForCall)]
	fake.storePropertiesArgsForCall = append(fake.storePropertiesArgsForCall, struct {
		arg1 volume.Properties
	}{arg1})
	fake.recordInvocation("StoreProperties", []interface{}{arg1})
	fake.storePropertiesMutex.Unlock()
	if fake.StorePropertiesStub != nil {
		return fake.StorePropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.storePropertiesReturns.result1
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesCallCount() int {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	return len(fake.storePropertiesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesArgsForCall(i int) volume.Properties {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	return fake.storePropertiesArgsForCall[i].arg1
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesReturns(result1 error) {
	fake.StorePropertiesStub = nil
	fake.storePropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) StorePropertiesReturnsOnCall(i int, result1 error) {
	fake.StorePropertiesStub = nil
	if fake.storePropertiesReturnsOnCall == nil {
		fake.storePropertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storePropertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) LoadTTL() (volume.TTL, time.Time, error) {
	fake.loadTTLMutex.Lock()
	ret, specificReturn := fake.loadTTLReturnsOnCall[len(fake.loadTTLArgsForCall)]
	fake.loadTTLArgsForCall = append(fake.loadTTLArgsForCall, struct{}{})
	fake.recordInvocation("LoadTTL", []interface{}{})
	fake.loadTTLMutex.Unlock()
	if fake.LoadTTLStub != nil {
		return fake.LoadTTLStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.loadTTLReturns.result1, fake.loadTTLReturns.result2, fake.loadTTLReturns.result3
}

func (fake *FakeFilesystemLiveVolume) LoadTTLCallCount() int {
	fake.loadTTLMutex.RLock()
	defer fake.loadTTLMutex.RUnlock()
	return len(fake.loadTTLArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) LoadTTLReturns(result1 volume.TTL, result2 time.Time, result3 error) {
	fake.LoadTTLStub = nil
	fake.loadTTLReturns = struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) LoadTTLReturnsOnCall(i int, result1 volume.TTL, result2 time.Time, result3 error) {
	fake.LoadTTLStub = nil
	if fake.loadTTLReturnsOnCall == nil {
		fake.loadTTLReturnsOnCall = make(map[int]struct {
			result1 volume.TTL
			result2 time.Time
			result3 error
		})
	}
	fake.loadTTLReturnsOnCall[i] = struct {
		result1 volume.TTL
		result2 time.Time
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) StoreTTL(arg1 volume.TTL) (time.Time, error) {
	fake.storeTTLMutex.Lock()
	ret, specificReturn := fake.storeTTLReturnsOnCall[len(fake.storeTTLArgsForCall)]
	fake.storeTTLArgsForCall = append(fake.storeTTLArgsForCall, struct {
		arg1 volume.TTL
	}{arg1})
	fake.recordInvocation("StoreTTL", []interface{}{arg1})
	fake.storeTTLMutex.Unlock()
	if fake.StoreTTLStub != nil {
		return fake.StoreTTLStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.storeTTLReturns.result1, fake.storeTTLReturns.result2
}

func (fake *FakeFilesystemLiveVolume) StoreTTLCallCount() int {
	fake.storeTTLMutex.RLock()
	defer fake.storeTTLMutex.RUnlock()
	return len(fake.storeTTLArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) StoreTTLArgsForCall(i int) volume.TTL {
	fake.storeTTLMutex.RLock()
	defer fake.storeTTLMutex.RUnlock()
	return fake.storeTTLArgsForCall[i].arg1
}

func (fake *FakeFilesystemLiveVolume) StoreTTLReturns(result1 time.Time, result2 error) {
	fake.StoreTTLStub = nil
	fake.storeTTLReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) StoreTTLReturnsOnCall(i int, result1 time.Time, result2 error) {
	fake.StoreTTLStub = nil
	if fake.storeTTLReturnsOnCall == nil {
		fake.storeTTLReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 error
		})
	}
	fake.storeTTLReturnsOnCall[i] = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) LoadPrivileged() (bool, error) {
	fake.loadPrivilegedMutex.Lock()
	ret, specificReturn := fake.loadPrivilegedReturnsOnCall[len(fake.loadPrivilegedArgsForCall)]
	fake.loadPrivilegedArgsForCall = append(fake.loadPrivilegedArgsForCall, struct{}{})
	fake.recordInvocation("LoadPrivileged", []interface{}{})
	fake.loadPrivilegedMutex.Unlock()
	if fake.LoadPrivilegedStub != nil {
		return fake.LoadPrivilegedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadPrivilegedReturns.result1, fake.loadPrivilegedReturns.result2
}

func (fake *FakeFilesystemLiveVolume) LoadPrivilegedCallCount() int {
	fake.loadPrivilegedMutex.RLock()
	defer fake.loadPrivilegedMutex.RUnlock()
	return len(fake.loadPrivilegedArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) LoadPrivilegedReturns(result1 bool, result2 error) {
	fake.LoadPrivilegedStub = nil
	fake.loadPrivilegedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) LoadPrivilegedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.LoadPrivilegedStub = nil
	if fake.loadPrivilegedReturnsOnCall == nil {
		fake.loadPrivilegedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.loadPrivilegedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) StorePrivileged(arg1 bool) error {
	fake.storePrivilegedMutex.Lock()
	ret, specificReturn := fake.storePrivilegedReturnsOnCall[len(fake.storePrivilegedArgsForCall)]
	fake.storePrivilegedArgsForCall = append(fake.storePrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("StorePrivileged", []interface{}{arg1})
	fake.storePrivilegedMutex.Unlock()
	if fake.StorePrivilegedStub != nil {
		return fake.StorePrivilegedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.storePrivilegedReturns.result1
}

func (fake *FakeFilesystemLiveVolume) StorePrivilegedCallCount() int {
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	return len(fake.storePrivilegedArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) StorePrivilegedArgsForCall(i int) bool {
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	return fake.storePrivilegedArgsForCall[i].arg1
}

func (fake *FakeFilesystemLiveVolume) StorePrivilegedReturns(result1 error) {
	fake.StorePrivilegedStub = nil
	fake.storePrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) StorePrivilegedReturnsOnCall(i int, result1 error) {
	fake.StorePrivilegedStub = nil
	if fake.storePrivilegedReturnsOnCall == nil {
		fake.storePrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storePrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) Parent() (volume.FilesystemLiveVolume, bool, error) {
	fake.parentMutex.Lock()
	ret, specificReturn := fake.parentReturnsOnCall[len(fake.parentArgsForCall)]
	fake.parentArgsForCall = append(fake.parentArgsForCall, struct{}{})
	fake.recordInvocation("Parent", []interface{}{})
	fake.parentMutex.Unlock()
	if fake.ParentStub != nil {
		return fake.ParentStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.parentReturns.result1, fake.parentReturns.result2, fake.parentReturns.result3
}

func (fake *FakeFilesystemLiveVolume) ParentCallCount() int {
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	return len(fake.parentArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) ParentReturns(result1 volume.FilesystemLiveVolume, result2 bool, result3 error) {
	fake.ParentStub = nil
	fake.parentReturns = struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) ParentReturnsOnCall(i int, result1 volume.FilesystemLiveVolume, result2 bool, result3 error) {
	fake.ParentStub = nil
	if fake.parentReturnsOnCall == nil {
		fake.parentReturnsOnCall = make(map[int]struct {
			result1 volume.FilesystemLiveVolume
			result2 bool
			result3 error
		})
	}
	fake.parentReturnsOnCall[i] = struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemLiveVolume) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeFilesystemLiveVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) DestroyReturnsOnCall(i int, result1 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemLiveVolume) SizeInBytes() (int64, error) {
	fake.sizeInBytesMutex.Lock()
	ret, specificReturn := fake.sizeInBytesReturnsOnCall[len(fake.sizeInBytesArgsForCall)]
	fake.sizeInBytesArgsForCall = append(fake.sizeInBytesArgsForCall, struct{}{})
	fake.recordInvocation("SizeInBytes", []interface{}{})
	fake.sizeInBytesMutex.Unlock()
	if fake.SizeInBytesStub != nil {
		return fake.SizeInBytesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sizeInBytesReturns.result1, fake.sizeInBytesReturns.result2
}

func (fake *FakeFilesystemLiveVolume) SizeInBytesCallCount() int {
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	return len(fake.sizeInBytesArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) SizeInBytesReturns(result1 int64, result2 error) {
	fake.SizeInBytesStub = nil
	fake.sizeInBytesReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) SizeInBytesReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SizeInBytesStub = nil
	if fake.sizeInBytesReturnsOnCall == nil {
		fake.sizeInBytesReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.sizeInBytesReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) NewSubvolume(handle string) (volume.FilesystemInitVolume, error) {
	fake.newSubvolumeMutex.Lock()
	ret, specificReturn := fake.newSubvolumeReturnsOnCall[len(fake.newSubvolumeArgsForCall)]
	fake.newSubvolumeArgsForCall = append(fake.newSubvolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("NewSubvolume", []interface{}{handle})
	fake.newSubvolumeMutex.Unlock()
	if fake.NewSubvolumeStub != nil {
		return fake.NewSubvolumeStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newSubvolumeReturns.result1, fake.newSubvolumeReturns.result2
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeCallCount() int {
	fake.newSubvolumeMutex.RLock()
	defer fake.newSubvolumeMutex.RUnlock()
	return len(fake.newSubvolumeArgsForCall)
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeArgsForCall(i int) string {
	fake.newSubvolumeMutex.RLock()
	defer fake.newSubvolumeMutex.RUnlock()
	return fake.newSubvolumeArgsForCall[i].handle
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeReturns(result1 volume.FilesystemInitVolume, result2 error) {
	fake.NewSubvolumeStub = nil
	fake.newSubvolumeReturns = struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) NewSubvolumeReturnsOnCall(i int, result1 volume.FilesystemInitVolume, result2 error) {
	fake.NewSubvolumeStub = nil
	if fake.newSubvolumeReturnsOnCall == nil {
		fake.newSubvolumeReturnsOnCall = make(map[int]struct {
			result1 volume.FilesystemInitVolume
			result2 error
		})
	}
	fake.newSubvolumeReturnsOnCall[i] = struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemLiveVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.dataPathMutex.RLock()
	defer fake.dataPathMutex.RUnlock()
	fake.loadPropertiesMutex.RLock()
	defer fake.loadPropertiesMutex.RUnlock()
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	fake.loadTTLMutex.RLock()
	defer fake.loadTTLMutex.RUnlock()
	fake.storeTTLMutex.RLock()
	defer fake.storeTTLMutex.RUnlock()
	fake.loadPrivilegedMutex.RLock()
	defer fake.loadPrivilegedMutex.RUnlock()
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.sizeInBytesMutex.RLock()
	defer fake.sizeInBytesMutex.RUnlock()
	fake.newSubvolumeMutex.RLock()
	defer fake.newSubvolumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFilesystemLiveVolume) recordInvocation(key string, args []interface{}) {
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

var _ volume.FilesystemLiveVolume = new(FakeFilesystemLiveVolume)
