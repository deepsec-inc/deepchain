// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	deliver "deepchain/common/deliver"
	blockledger "deepchain/common/ledger/blockledger"
	policies "deepchain/common/policies"
)

type Chain struct {
	ErroredStub        func() <-chan struct{}
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
	}
	erroredReturns struct {
		result1 <-chan struct{}
	}
	erroredReturnsOnCall map[int]struct {
		result1 <-chan struct{}
	}
	PolicyManagerStub        func() policies.Manager
	policyManagerMutex       sync.RWMutex
	policyManagerArgsForCall []struct {
	}
	policyManagerReturns struct {
		result1 policies.Manager
	}
	policyManagerReturnsOnCall map[int]struct {
		result1 policies.Manager
	}
	ReaderStub        func() blockledger.Reader
	readerMutex       sync.RWMutex
	readerArgsForCall []struct {
	}
	readerReturns struct {
		result1 blockledger.Reader
	}
	readerReturnsOnCall map[int]struct {
		result1 blockledger.Reader
	}
	SequenceStub        func() uint64
	sequenceMutex       sync.RWMutex
	sequenceArgsForCall []struct {
	}
	sequenceReturns struct {
		result1 uint64
	}
	sequenceReturnsOnCall map[int]struct {
		result1 uint64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Chain) Errored() <-chan struct{} {
	fake.erroredMutex.Lock()
	ret, specificReturn := fake.erroredReturnsOnCall[len(fake.erroredArgsForCall)]
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
	}{})
	fake.recordInvocation("Errored", []interface{}{})
	fake.erroredMutex.Unlock()
	if fake.ErroredStub != nil {
		return fake.ErroredStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.erroredReturns
	return fakeReturns.result1
}

func (fake *Chain) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *Chain) ErroredCalls(stub func() <-chan struct{}) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *Chain) ErroredReturns(result1 <-chan struct{}) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = nil
	fake.erroredReturns = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *Chain) ErroredReturnsOnCall(i int, result1 <-chan struct{}) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = nil
	if fake.erroredReturnsOnCall == nil {
		fake.erroredReturnsOnCall = make(map[int]struct {
			result1 <-chan struct{}
		})
	}
	fake.erroredReturnsOnCall[i] = struct {
		result1 <-chan struct{}
	}{result1}
}

func (fake *Chain) PolicyManager() policies.Manager {
	fake.policyManagerMutex.Lock()
	ret, specificReturn := fake.policyManagerReturnsOnCall[len(fake.policyManagerArgsForCall)]
	fake.policyManagerArgsForCall = append(fake.policyManagerArgsForCall, struct {
	}{})
	fake.recordInvocation("PolicyManager", []interface{}{})
	fake.policyManagerMutex.Unlock()
	if fake.PolicyManagerStub != nil {
		return fake.PolicyManagerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.policyManagerReturns
	return fakeReturns.result1
}

func (fake *Chain) PolicyManagerCallCount() int {
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	return len(fake.policyManagerArgsForCall)
}

func (fake *Chain) PolicyManagerCalls(stub func() policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = stub
}

func (fake *Chain) PolicyManagerReturns(result1 policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = nil
	fake.policyManagerReturns = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *Chain) PolicyManagerReturnsOnCall(i int, result1 policies.Manager) {
	fake.policyManagerMutex.Lock()
	defer fake.policyManagerMutex.Unlock()
	fake.PolicyManagerStub = nil
	if fake.policyManagerReturnsOnCall == nil {
		fake.policyManagerReturnsOnCall = make(map[int]struct {
			result1 policies.Manager
		})
	}
	fake.policyManagerReturnsOnCall[i] = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *Chain) Reader() blockledger.Reader {
	fake.readerMutex.Lock()
	ret, specificReturn := fake.readerReturnsOnCall[len(fake.readerArgsForCall)]
	fake.readerArgsForCall = append(fake.readerArgsForCall, struct {
	}{})
	fake.recordInvocation("Reader", []interface{}{})
	fake.readerMutex.Unlock()
	if fake.ReaderStub != nil {
		return fake.ReaderStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.readerReturns
	return fakeReturns.result1
}

func (fake *Chain) ReaderCallCount() int {
	fake.readerMutex.RLock()
	defer fake.readerMutex.RUnlock()
	return len(fake.readerArgsForCall)
}

func (fake *Chain) ReaderCalls(stub func() blockledger.Reader) {
	fake.readerMutex.Lock()
	defer fake.readerMutex.Unlock()
	fake.ReaderStub = stub
}

func (fake *Chain) ReaderReturns(result1 blockledger.Reader) {
	fake.readerMutex.Lock()
	defer fake.readerMutex.Unlock()
	fake.ReaderStub = nil
	fake.readerReturns = struct {
		result1 blockledger.Reader
	}{result1}
}

func (fake *Chain) ReaderReturnsOnCall(i int, result1 blockledger.Reader) {
	fake.readerMutex.Lock()
	defer fake.readerMutex.Unlock()
	fake.ReaderStub = nil
	if fake.readerReturnsOnCall == nil {
		fake.readerReturnsOnCall = make(map[int]struct {
			result1 blockledger.Reader
		})
	}
	fake.readerReturnsOnCall[i] = struct {
		result1 blockledger.Reader
	}{result1}
}

func (fake *Chain) Sequence() uint64 {
	fake.sequenceMutex.Lock()
	ret, specificReturn := fake.sequenceReturnsOnCall[len(fake.sequenceArgsForCall)]
	fake.sequenceArgsForCall = append(fake.sequenceArgsForCall, struct {
	}{})
	fake.recordInvocation("Sequence", []interface{}{})
	fake.sequenceMutex.Unlock()
	if fake.SequenceStub != nil {
		return fake.SequenceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sequenceReturns
	return fakeReturns.result1
}

func (fake *Chain) SequenceCallCount() int {
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	return len(fake.sequenceArgsForCall)
}

func (fake *Chain) SequenceCalls(stub func() uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = stub
}

func (fake *Chain) SequenceReturns(result1 uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = nil
	fake.sequenceReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *Chain) SequenceReturnsOnCall(i int, result1 uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = nil
	if fake.sequenceReturnsOnCall == nil {
		fake.sequenceReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.sequenceReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *Chain) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.policyManagerMutex.RLock()
	defer fake.policyManagerMutex.RUnlock()
	fake.readerMutex.RLock()
	defer fake.readerMutex.RUnlock()
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Chain) recordInvocation(key string, args []interface{}) {
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

var _ deliver.Chain = new(Chain)
