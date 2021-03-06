// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-cli/pkg/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

type MSP struct {
	AddAffiliationStub        func(request *msp.AffiliationRequest) (*msp.AffiliationResponse, error)
	addAffiliationMutex       sync.RWMutex
	addAffiliationArgsForCall []struct {
		request *msp.AffiliationRequest
	}
	addAffiliationReturns struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	addAffiliationReturnsOnCall map[int]struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	CreateIdentityStub        func(request *msp.IdentityRequest) (*msp.IdentityResponse, error)
	createIdentityMutex       sync.RWMutex
	createIdentityArgsForCall []struct {
		request *msp.IdentityRequest
	}
	createIdentityReturns struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	createIdentityReturnsOnCall map[int]struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	CreateSigningIdentityStub        func(opts ...mspctx.SigningIdentityOption) (mspctx.SigningIdentity, error)
	createSigningIdentityMutex       sync.RWMutex
	createSigningIdentityArgsForCall []struct {
		opts []mspctx.SigningIdentityOption
	}
	createSigningIdentityReturns struct {
		result1 mspctx.SigningIdentity
		result2 error
	}
	createSigningIdentityReturnsOnCall map[int]struct {
		result1 mspctx.SigningIdentity
		result2 error
	}
	EnrollStub        func(enrollmentID string, opts ...msp.EnrollmentOption) error
	enrollMutex       sync.RWMutex
	enrollArgsForCall []struct {
		enrollmentID string
		opts         []msp.EnrollmentOption
	}
	enrollReturns struct {
		result1 error
	}
	enrollReturnsOnCall map[int]struct {
		result1 error
	}
	GetAffiliationStub        func(affiliation string, options ...msp.RequestOption) (*msp.AffiliationResponse, error)
	getAffiliationMutex       sync.RWMutex
	getAffiliationArgsForCall []struct {
		affiliation string
		options     []msp.RequestOption
	}
	getAffiliationReturns struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	getAffiliationReturnsOnCall map[int]struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	GetAllAffiliationsStub        func(options ...msp.RequestOption) (*msp.AffiliationResponse, error)
	getAllAffiliationsMutex       sync.RWMutex
	getAllAffiliationsArgsForCall []struct {
		options []msp.RequestOption
	}
	getAllAffiliationsReturns struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	getAllAffiliationsReturnsOnCall map[int]struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	GetAllIdentitiesStub        func(options ...msp.RequestOption) ([]*msp.IdentityResponse, error)
	getAllIdentitiesMutex       sync.RWMutex
	getAllIdentitiesArgsForCall []struct {
		options []msp.RequestOption
	}
	getAllIdentitiesReturns struct {
		result1 []*msp.IdentityResponse
		result2 error
	}
	getAllIdentitiesReturnsOnCall map[int]struct {
		result1 []*msp.IdentityResponse
		result2 error
	}
	GetCAInfoStub        func() (*msp.GetCAInfoResponse, error)
	getCAInfoMutex       sync.RWMutex
	getCAInfoArgsForCall []struct{}
	getCAInfoReturns     struct {
		result1 *msp.GetCAInfoResponse
		result2 error
	}
	getCAInfoReturnsOnCall map[int]struct {
		result1 *msp.GetCAInfoResponse
		result2 error
	}
	GetIdentityStub        func(ID string, options ...msp.RequestOption) (*msp.IdentityResponse, error)
	getIdentityMutex       sync.RWMutex
	getIdentityArgsForCall []struct {
		ID      string
		options []msp.RequestOption
	}
	getIdentityReturns struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	getIdentityReturnsOnCall map[int]struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	GetSigningIdentityStub        func(id string) (mspctx.SigningIdentity, error)
	getSigningIdentityMutex       sync.RWMutex
	getSigningIdentityArgsForCall []struct {
		id string
	}
	getSigningIdentityReturns struct {
		result1 mspctx.SigningIdentity
		result2 error
	}
	getSigningIdentityReturnsOnCall map[int]struct {
		result1 mspctx.SigningIdentity
		result2 error
	}
	ModifyAffiliationStub        func(request *msp.ModifyAffiliationRequest) (*msp.AffiliationResponse, error)
	modifyAffiliationMutex       sync.RWMutex
	modifyAffiliationArgsForCall []struct {
		request *msp.ModifyAffiliationRequest
	}
	modifyAffiliationReturns struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	modifyAffiliationReturnsOnCall map[int]struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	ModifyIdentityStub        func(request *msp.IdentityRequest) (*msp.IdentityResponse, error)
	modifyIdentityMutex       sync.RWMutex
	modifyIdentityArgsForCall []struct {
		request *msp.IdentityRequest
	}
	modifyIdentityReturns struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	modifyIdentityReturnsOnCall map[int]struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	ReenrollStub        func(enrollmentID string, opts ...msp.EnrollmentOption) error
	reenrollMutex       sync.RWMutex
	reenrollArgsForCall []struct {
		enrollmentID string
		opts         []msp.EnrollmentOption
	}
	reenrollReturns struct {
		result1 error
	}
	reenrollReturnsOnCall map[int]struct {
		result1 error
	}
	RegisterStub        func(request *msp.RegistrationRequest) (string, error)
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		request *msp.RegistrationRequest
	}
	registerReturns struct {
		result1 string
		result2 error
	}
	registerReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RemoveAffiliationStub        func(request *msp.AffiliationRequest) (*msp.AffiliationResponse, error)
	removeAffiliationMutex       sync.RWMutex
	removeAffiliationArgsForCall []struct {
		request *msp.AffiliationRequest
	}
	removeAffiliationReturns struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	removeAffiliationReturnsOnCall map[int]struct {
		result1 *msp.AffiliationResponse
		result2 error
	}
	RemoveIdentityStub        func(request *msp.RemoveIdentityRequest) (*msp.IdentityResponse, error)
	removeIdentityMutex       sync.RWMutex
	removeIdentityArgsForCall []struct {
		request *msp.RemoveIdentityRequest
	}
	removeIdentityReturns struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	removeIdentityReturnsOnCall map[int]struct {
		result1 *msp.IdentityResponse
		result2 error
	}
	RevokeStub        func(request *msp.RevocationRequest) (*msp.RevocationResponse, error)
	revokeMutex       sync.RWMutex
	revokeArgsForCall []struct {
		request *msp.RevocationRequest
	}
	revokeReturns struct {
		result1 *msp.RevocationResponse
		result2 error
	}
	revokeReturnsOnCall map[int]struct {
		result1 *msp.RevocationResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MSP) AddAffiliation(request *msp.AffiliationRequest) (*msp.AffiliationResponse, error) {
	fake.addAffiliationMutex.Lock()
	ret, specificReturn := fake.addAffiliationReturnsOnCall[len(fake.addAffiliationArgsForCall)]
	fake.addAffiliationArgsForCall = append(fake.addAffiliationArgsForCall, struct {
		request *msp.AffiliationRequest
	}{request})
	fake.recordInvocation("AddAffiliation", []interface{}{request})
	fake.addAffiliationMutex.Unlock()
	if fake.AddAffiliationStub != nil {
		return fake.AddAffiliationStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addAffiliationReturns.result1, fake.addAffiliationReturns.result2
}

func (fake *MSP) AddAffiliationCallCount() int {
	fake.addAffiliationMutex.RLock()
	defer fake.addAffiliationMutex.RUnlock()
	return len(fake.addAffiliationArgsForCall)
}

func (fake *MSP) AddAffiliationArgsForCall(i int) *msp.AffiliationRequest {
	fake.addAffiliationMutex.RLock()
	defer fake.addAffiliationMutex.RUnlock()
	return fake.addAffiliationArgsForCall[i].request
}

func (fake *MSP) AddAffiliationReturns(result1 *msp.AffiliationResponse, result2 error) {
	fake.AddAffiliationStub = nil
	fake.addAffiliationReturns = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) AddAffiliationReturnsOnCall(i int, result1 *msp.AffiliationResponse, result2 error) {
	fake.AddAffiliationStub = nil
	if fake.addAffiliationReturnsOnCall == nil {
		fake.addAffiliationReturnsOnCall = make(map[int]struct {
			result1 *msp.AffiliationResponse
			result2 error
		})
	}
	fake.addAffiliationReturnsOnCall[i] = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) CreateIdentity(request *msp.IdentityRequest) (*msp.IdentityResponse, error) {
	fake.createIdentityMutex.Lock()
	ret, specificReturn := fake.createIdentityReturnsOnCall[len(fake.createIdentityArgsForCall)]
	fake.createIdentityArgsForCall = append(fake.createIdentityArgsForCall, struct {
		request *msp.IdentityRequest
	}{request})
	fake.recordInvocation("CreateIdentity", []interface{}{request})
	fake.createIdentityMutex.Unlock()
	if fake.CreateIdentityStub != nil {
		return fake.CreateIdentityStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIdentityReturns.result1, fake.createIdentityReturns.result2
}

func (fake *MSP) CreateIdentityCallCount() int {
	fake.createIdentityMutex.RLock()
	defer fake.createIdentityMutex.RUnlock()
	return len(fake.createIdentityArgsForCall)
}

func (fake *MSP) CreateIdentityArgsForCall(i int) *msp.IdentityRequest {
	fake.createIdentityMutex.RLock()
	defer fake.createIdentityMutex.RUnlock()
	return fake.createIdentityArgsForCall[i].request
}

func (fake *MSP) CreateIdentityReturns(result1 *msp.IdentityResponse, result2 error) {
	fake.CreateIdentityStub = nil
	fake.createIdentityReturns = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) CreateIdentityReturnsOnCall(i int, result1 *msp.IdentityResponse, result2 error) {
	fake.CreateIdentityStub = nil
	if fake.createIdentityReturnsOnCall == nil {
		fake.createIdentityReturnsOnCall = make(map[int]struct {
			result1 *msp.IdentityResponse
			result2 error
		})
	}
	fake.createIdentityReturnsOnCall[i] = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) CreateSigningIdentity(opts ...mspctx.SigningIdentityOption) (mspctx.SigningIdentity, error) {
	fake.createSigningIdentityMutex.Lock()
	ret, specificReturn := fake.createSigningIdentityReturnsOnCall[len(fake.createSigningIdentityArgsForCall)]
	fake.createSigningIdentityArgsForCall = append(fake.createSigningIdentityArgsForCall, struct {
		opts []mspctx.SigningIdentityOption
	}{opts})
	fake.recordInvocation("CreateSigningIdentity", []interface{}{opts})
	fake.createSigningIdentityMutex.Unlock()
	if fake.CreateSigningIdentityStub != nil {
		return fake.CreateSigningIdentityStub(opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createSigningIdentityReturns.result1, fake.createSigningIdentityReturns.result2
}

func (fake *MSP) CreateSigningIdentityCallCount() int {
	fake.createSigningIdentityMutex.RLock()
	defer fake.createSigningIdentityMutex.RUnlock()
	return len(fake.createSigningIdentityArgsForCall)
}

func (fake *MSP) CreateSigningIdentityArgsForCall(i int) []mspctx.SigningIdentityOption {
	fake.createSigningIdentityMutex.RLock()
	defer fake.createSigningIdentityMutex.RUnlock()
	return fake.createSigningIdentityArgsForCall[i].opts
}

func (fake *MSP) CreateSigningIdentityReturns(result1 mspctx.SigningIdentity, result2 error) {
	fake.CreateSigningIdentityStub = nil
	fake.createSigningIdentityReturns = struct {
		result1 mspctx.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) CreateSigningIdentityReturnsOnCall(i int, result1 mspctx.SigningIdentity, result2 error) {
	fake.CreateSigningIdentityStub = nil
	if fake.createSigningIdentityReturnsOnCall == nil {
		fake.createSigningIdentityReturnsOnCall = make(map[int]struct {
			result1 mspctx.SigningIdentity
			result2 error
		})
	}
	fake.createSigningIdentityReturnsOnCall[i] = struct {
		result1 mspctx.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) Enroll(enrollmentID string, opts ...msp.EnrollmentOption) error {
	fake.enrollMutex.Lock()
	ret, specificReturn := fake.enrollReturnsOnCall[len(fake.enrollArgsForCall)]
	fake.enrollArgsForCall = append(fake.enrollArgsForCall, struct {
		enrollmentID string
		opts         []msp.EnrollmentOption
	}{enrollmentID, opts})
	fake.recordInvocation("Enroll", []interface{}{enrollmentID, opts})
	fake.enrollMutex.Unlock()
	if fake.EnrollStub != nil {
		return fake.EnrollStub(enrollmentID, opts...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.enrollReturns.result1
}

func (fake *MSP) EnrollCallCount() int {
	fake.enrollMutex.RLock()
	defer fake.enrollMutex.RUnlock()
	return len(fake.enrollArgsForCall)
}

func (fake *MSP) EnrollArgsForCall(i int) (string, []msp.EnrollmentOption) {
	fake.enrollMutex.RLock()
	defer fake.enrollMutex.RUnlock()
	return fake.enrollArgsForCall[i].enrollmentID, fake.enrollArgsForCall[i].opts
}

func (fake *MSP) EnrollReturns(result1 error) {
	fake.EnrollStub = nil
	fake.enrollReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) EnrollReturnsOnCall(i int, result1 error) {
	fake.EnrollStub = nil
	if fake.enrollReturnsOnCall == nil {
		fake.enrollReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.enrollReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) GetAffiliation(affiliation string, options ...msp.RequestOption) (*msp.AffiliationResponse, error) {
	fake.getAffiliationMutex.Lock()
	ret, specificReturn := fake.getAffiliationReturnsOnCall[len(fake.getAffiliationArgsForCall)]
	fake.getAffiliationArgsForCall = append(fake.getAffiliationArgsForCall, struct {
		affiliation string
		options     []msp.RequestOption
	}{affiliation, options})
	fake.recordInvocation("GetAffiliation", []interface{}{affiliation, options})
	fake.getAffiliationMutex.Unlock()
	if fake.GetAffiliationStub != nil {
		return fake.GetAffiliationStub(affiliation, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAffiliationReturns.result1, fake.getAffiliationReturns.result2
}

func (fake *MSP) GetAffiliationCallCount() int {
	fake.getAffiliationMutex.RLock()
	defer fake.getAffiliationMutex.RUnlock()
	return len(fake.getAffiliationArgsForCall)
}

func (fake *MSP) GetAffiliationArgsForCall(i int) (string, []msp.RequestOption) {
	fake.getAffiliationMutex.RLock()
	defer fake.getAffiliationMutex.RUnlock()
	return fake.getAffiliationArgsForCall[i].affiliation, fake.getAffiliationArgsForCall[i].options
}

func (fake *MSP) GetAffiliationReturns(result1 *msp.AffiliationResponse, result2 error) {
	fake.GetAffiliationStub = nil
	fake.getAffiliationReturns = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetAffiliationReturnsOnCall(i int, result1 *msp.AffiliationResponse, result2 error) {
	fake.GetAffiliationStub = nil
	if fake.getAffiliationReturnsOnCall == nil {
		fake.getAffiliationReturnsOnCall = make(map[int]struct {
			result1 *msp.AffiliationResponse
			result2 error
		})
	}
	fake.getAffiliationReturnsOnCall[i] = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetAllAffiliations(options ...msp.RequestOption) (*msp.AffiliationResponse, error) {
	fake.getAllAffiliationsMutex.Lock()
	ret, specificReturn := fake.getAllAffiliationsReturnsOnCall[len(fake.getAllAffiliationsArgsForCall)]
	fake.getAllAffiliationsArgsForCall = append(fake.getAllAffiliationsArgsForCall, struct {
		options []msp.RequestOption
	}{options})
	fake.recordInvocation("GetAllAffiliations", []interface{}{options})
	fake.getAllAffiliationsMutex.Unlock()
	if fake.GetAllAffiliationsStub != nil {
		return fake.GetAllAffiliationsStub(options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllAffiliationsReturns.result1, fake.getAllAffiliationsReturns.result2
}

func (fake *MSP) GetAllAffiliationsCallCount() int {
	fake.getAllAffiliationsMutex.RLock()
	defer fake.getAllAffiliationsMutex.RUnlock()
	return len(fake.getAllAffiliationsArgsForCall)
}

func (fake *MSP) GetAllAffiliationsArgsForCall(i int) []msp.RequestOption {
	fake.getAllAffiliationsMutex.RLock()
	defer fake.getAllAffiliationsMutex.RUnlock()
	return fake.getAllAffiliationsArgsForCall[i].options
}

func (fake *MSP) GetAllAffiliationsReturns(result1 *msp.AffiliationResponse, result2 error) {
	fake.GetAllAffiliationsStub = nil
	fake.getAllAffiliationsReturns = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetAllAffiliationsReturnsOnCall(i int, result1 *msp.AffiliationResponse, result2 error) {
	fake.GetAllAffiliationsStub = nil
	if fake.getAllAffiliationsReturnsOnCall == nil {
		fake.getAllAffiliationsReturnsOnCall = make(map[int]struct {
			result1 *msp.AffiliationResponse
			result2 error
		})
	}
	fake.getAllAffiliationsReturnsOnCall[i] = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetAllIdentities(options ...msp.RequestOption) ([]*msp.IdentityResponse, error) {
	fake.getAllIdentitiesMutex.Lock()
	ret, specificReturn := fake.getAllIdentitiesReturnsOnCall[len(fake.getAllIdentitiesArgsForCall)]
	fake.getAllIdentitiesArgsForCall = append(fake.getAllIdentitiesArgsForCall, struct {
		options []msp.RequestOption
	}{options})
	fake.recordInvocation("GetAllIdentities", []interface{}{options})
	fake.getAllIdentitiesMutex.Unlock()
	if fake.GetAllIdentitiesStub != nil {
		return fake.GetAllIdentitiesStub(options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllIdentitiesReturns.result1, fake.getAllIdentitiesReturns.result2
}

func (fake *MSP) GetAllIdentitiesCallCount() int {
	fake.getAllIdentitiesMutex.RLock()
	defer fake.getAllIdentitiesMutex.RUnlock()
	return len(fake.getAllIdentitiesArgsForCall)
}

func (fake *MSP) GetAllIdentitiesArgsForCall(i int) []msp.RequestOption {
	fake.getAllIdentitiesMutex.RLock()
	defer fake.getAllIdentitiesMutex.RUnlock()
	return fake.getAllIdentitiesArgsForCall[i].options
}

func (fake *MSP) GetAllIdentitiesReturns(result1 []*msp.IdentityResponse, result2 error) {
	fake.GetAllIdentitiesStub = nil
	fake.getAllIdentitiesReturns = struct {
		result1 []*msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetAllIdentitiesReturnsOnCall(i int, result1 []*msp.IdentityResponse, result2 error) {
	fake.GetAllIdentitiesStub = nil
	if fake.getAllIdentitiesReturnsOnCall == nil {
		fake.getAllIdentitiesReturnsOnCall = make(map[int]struct {
			result1 []*msp.IdentityResponse
			result2 error
		})
	}
	fake.getAllIdentitiesReturnsOnCall[i] = struct {
		result1 []*msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetCAInfo() (*msp.GetCAInfoResponse, error) {
	fake.getCAInfoMutex.Lock()
	ret, specificReturn := fake.getCAInfoReturnsOnCall[len(fake.getCAInfoArgsForCall)]
	fake.getCAInfoArgsForCall = append(fake.getCAInfoArgsForCall, struct{}{})
	fake.recordInvocation("GetCAInfo", []interface{}{})
	fake.getCAInfoMutex.Unlock()
	if fake.GetCAInfoStub != nil {
		return fake.GetCAInfoStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getCAInfoReturns.result1, fake.getCAInfoReturns.result2
}

func (fake *MSP) GetCAInfoCallCount() int {
	fake.getCAInfoMutex.RLock()
	defer fake.getCAInfoMutex.RUnlock()
	return len(fake.getCAInfoArgsForCall)
}

func (fake *MSP) GetCAInfoReturns(result1 *msp.GetCAInfoResponse, result2 error) {
	fake.GetCAInfoStub = nil
	fake.getCAInfoReturns = struct {
		result1 *msp.GetCAInfoResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetCAInfoReturnsOnCall(i int, result1 *msp.GetCAInfoResponse, result2 error) {
	fake.GetCAInfoStub = nil
	if fake.getCAInfoReturnsOnCall == nil {
		fake.getCAInfoReturnsOnCall = make(map[int]struct {
			result1 *msp.GetCAInfoResponse
			result2 error
		})
	}
	fake.getCAInfoReturnsOnCall[i] = struct {
		result1 *msp.GetCAInfoResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetIdentity(ID string, options ...msp.RequestOption) (*msp.IdentityResponse, error) {
	fake.getIdentityMutex.Lock()
	ret, specificReturn := fake.getIdentityReturnsOnCall[len(fake.getIdentityArgsForCall)]
	fake.getIdentityArgsForCall = append(fake.getIdentityArgsForCall, struct {
		ID      string
		options []msp.RequestOption
	}{ID, options})
	fake.recordInvocation("GetIdentity", []interface{}{ID, options})
	fake.getIdentityMutex.Unlock()
	if fake.GetIdentityStub != nil {
		return fake.GetIdentityStub(ID, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getIdentityReturns.result1, fake.getIdentityReturns.result2
}

func (fake *MSP) GetIdentityCallCount() int {
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	return len(fake.getIdentityArgsForCall)
}

func (fake *MSP) GetIdentityArgsForCall(i int) (string, []msp.RequestOption) {
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	return fake.getIdentityArgsForCall[i].ID, fake.getIdentityArgsForCall[i].options
}

func (fake *MSP) GetIdentityReturns(result1 *msp.IdentityResponse, result2 error) {
	fake.GetIdentityStub = nil
	fake.getIdentityReturns = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetIdentityReturnsOnCall(i int, result1 *msp.IdentityResponse, result2 error) {
	fake.GetIdentityStub = nil
	if fake.getIdentityReturnsOnCall == nil {
		fake.getIdentityReturnsOnCall = make(map[int]struct {
			result1 *msp.IdentityResponse
			result2 error
		})
	}
	fake.getIdentityReturnsOnCall[i] = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetSigningIdentity(id string) (mspctx.SigningIdentity, error) {
	fake.getSigningIdentityMutex.Lock()
	ret, specificReturn := fake.getSigningIdentityReturnsOnCall[len(fake.getSigningIdentityArgsForCall)]
	fake.getSigningIdentityArgsForCall = append(fake.getSigningIdentityArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetSigningIdentity", []interface{}{id})
	fake.getSigningIdentityMutex.Unlock()
	if fake.GetSigningIdentityStub != nil {
		return fake.GetSigningIdentityStub(id)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSigningIdentityReturns.result1, fake.getSigningIdentityReturns.result2
}

func (fake *MSP) GetSigningIdentityCallCount() int {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	return len(fake.getSigningIdentityArgsForCall)
}

func (fake *MSP) GetSigningIdentityArgsForCall(i int) string {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	return fake.getSigningIdentityArgsForCall[i].id
}

func (fake *MSP) GetSigningIdentityReturns(result1 mspctx.SigningIdentity, result2 error) {
	fake.GetSigningIdentityStub = nil
	fake.getSigningIdentityReturns = struct {
		result1 mspctx.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) GetSigningIdentityReturnsOnCall(i int, result1 mspctx.SigningIdentity, result2 error) {
	fake.GetSigningIdentityStub = nil
	if fake.getSigningIdentityReturnsOnCall == nil {
		fake.getSigningIdentityReturnsOnCall = make(map[int]struct {
			result1 mspctx.SigningIdentity
			result2 error
		})
	}
	fake.getSigningIdentityReturnsOnCall[i] = struct {
		result1 mspctx.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *MSP) ModifyAffiliation(request *msp.ModifyAffiliationRequest) (*msp.AffiliationResponse, error) {
	fake.modifyAffiliationMutex.Lock()
	ret, specificReturn := fake.modifyAffiliationReturnsOnCall[len(fake.modifyAffiliationArgsForCall)]
	fake.modifyAffiliationArgsForCall = append(fake.modifyAffiliationArgsForCall, struct {
		request *msp.ModifyAffiliationRequest
	}{request})
	fake.recordInvocation("ModifyAffiliation", []interface{}{request})
	fake.modifyAffiliationMutex.Unlock()
	if fake.ModifyAffiliationStub != nil {
		return fake.ModifyAffiliationStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.modifyAffiliationReturns.result1, fake.modifyAffiliationReturns.result2
}

func (fake *MSP) ModifyAffiliationCallCount() int {
	fake.modifyAffiliationMutex.RLock()
	defer fake.modifyAffiliationMutex.RUnlock()
	return len(fake.modifyAffiliationArgsForCall)
}

func (fake *MSP) ModifyAffiliationArgsForCall(i int) *msp.ModifyAffiliationRequest {
	fake.modifyAffiliationMutex.RLock()
	defer fake.modifyAffiliationMutex.RUnlock()
	return fake.modifyAffiliationArgsForCall[i].request
}

func (fake *MSP) ModifyAffiliationReturns(result1 *msp.AffiliationResponse, result2 error) {
	fake.ModifyAffiliationStub = nil
	fake.modifyAffiliationReturns = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) ModifyAffiliationReturnsOnCall(i int, result1 *msp.AffiliationResponse, result2 error) {
	fake.ModifyAffiliationStub = nil
	if fake.modifyAffiliationReturnsOnCall == nil {
		fake.modifyAffiliationReturnsOnCall = make(map[int]struct {
			result1 *msp.AffiliationResponse
			result2 error
		})
	}
	fake.modifyAffiliationReturnsOnCall[i] = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) ModifyIdentity(request *msp.IdentityRequest) (*msp.IdentityResponse, error) {
	fake.modifyIdentityMutex.Lock()
	ret, specificReturn := fake.modifyIdentityReturnsOnCall[len(fake.modifyIdentityArgsForCall)]
	fake.modifyIdentityArgsForCall = append(fake.modifyIdentityArgsForCall, struct {
		request *msp.IdentityRequest
	}{request})
	fake.recordInvocation("ModifyIdentity", []interface{}{request})
	fake.modifyIdentityMutex.Unlock()
	if fake.ModifyIdentityStub != nil {
		return fake.ModifyIdentityStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.modifyIdentityReturns.result1, fake.modifyIdentityReturns.result2
}

func (fake *MSP) ModifyIdentityCallCount() int {
	fake.modifyIdentityMutex.RLock()
	defer fake.modifyIdentityMutex.RUnlock()
	return len(fake.modifyIdentityArgsForCall)
}

func (fake *MSP) ModifyIdentityArgsForCall(i int) *msp.IdentityRequest {
	fake.modifyIdentityMutex.RLock()
	defer fake.modifyIdentityMutex.RUnlock()
	return fake.modifyIdentityArgsForCall[i].request
}

func (fake *MSP) ModifyIdentityReturns(result1 *msp.IdentityResponse, result2 error) {
	fake.ModifyIdentityStub = nil
	fake.modifyIdentityReturns = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) ModifyIdentityReturnsOnCall(i int, result1 *msp.IdentityResponse, result2 error) {
	fake.ModifyIdentityStub = nil
	if fake.modifyIdentityReturnsOnCall == nil {
		fake.modifyIdentityReturnsOnCall = make(map[int]struct {
			result1 *msp.IdentityResponse
			result2 error
		})
	}
	fake.modifyIdentityReturnsOnCall[i] = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) Reenroll(enrollmentID string, opts ...msp.EnrollmentOption) error {
	fake.reenrollMutex.Lock()
	ret, specificReturn := fake.reenrollReturnsOnCall[len(fake.reenrollArgsForCall)]
	fake.reenrollArgsForCall = append(fake.reenrollArgsForCall, struct {
		enrollmentID string
		opts         []msp.EnrollmentOption
	}{enrollmentID, opts})
	fake.recordInvocation("Reenroll", []interface{}{enrollmentID, opts})
	fake.reenrollMutex.Unlock()
	if fake.ReenrollStub != nil {
		return fake.ReenrollStub(enrollmentID, opts...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.reenrollReturns.result1
}

func (fake *MSP) ReenrollCallCount() int {
	fake.reenrollMutex.RLock()
	defer fake.reenrollMutex.RUnlock()
	return len(fake.reenrollArgsForCall)
}

func (fake *MSP) ReenrollArgsForCall(i int) (string, []msp.EnrollmentOption) {
	fake.reenrollMutex.RLock()
	defer fake.reenrollMutex.RUnlock()
	return fake.reenrollArgsForCall[i].enrollmentID, fake.reenrollArgsForCall[i].opts
}

func (fake *MSP) ReenrollReturns(result1 error) {
	fake.ReenrollStub = nil
	fake.reenrollReturns = struct {
		result1 error
	}{result1}
}

func (fake *MSP) ReenrollReturnsOnCall(i int, result1 error) {
	fake.ReenrollStub = nil
	if fake.reenrollReturnsOnCall == nil {
		fake.reenrollReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reenrollReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *MSP) Register(request *msp.RegistrationRequest) (string, error) {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		request *msp.RegistrationRequest
	}{request})
	fake.recordInvocation("Register", []interface{}{request})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.registerReturns.result1, fake.registerReturns.result2
}

func (fake *MSP) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *MSP) RegisterArgsForCall(i int) *msp.RegistrationRequest {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].request
}

func (fake *MSP) RegisterReturns(result1 string, result2 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *MSP) RegisterReturnsOnCall(i int, result1 string, result2 error) {
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *MSP) RemoveAffiliation(request *msp.AffiliationRequest) (*msp.AffiliationResponse, error) {
	fake.removeAffiliationMutex.Lock()
	ret, specificReturn := fake.removeAffiliationReturnsOnCall[len(fake.removeAffiliationArgsForCall)]
	fake.removeAffiliationArgsForCall = append(fake.removeAffiliationArgsForCall, struct {
		request *msp.AffiliationRequest
	}{request})
	fake.recordInvocation("RemoveAffiliation", []interface{}{request})
	fake.removeAffiliationMutex.Unlock()
	if fake.RemoveAffiliationStub != nil {
		return fake.RemoveAffiliationStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.removeAffiliationReturns.result1, fake.removeAffiliationReturns.result2
}

func (fake *MSP) RemoveAffiliationCallCount() int {
	fake.removeAffiliationMutex.RLock()
	defer fake.removeAffiliationMutex.RUnlock()
	return len(fake.removeAffiliationArgsForCall)
}

func (fake *MSP) RemoveAffiliationArgsForCall(i int) *msp.AffiliationRequest {
	fake.removeAffiliationMutex.RLock()
	defer fake.removeAffiliationMutex.RUnlock()
	return fake.removeAffiliationArgsForCall[i].request
}

func (fake *MSP) RemoveAffiliationReturns(result1 *msp.AffiliationResponse, result2 error) {
	fake.RemoveAffiliationStub = nil
	fake.removeAffiliationReturns = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) RemoveAffiliationReturnsOnCall(i int, result1 *msp.AffiliationResponse, result2 error) {
	fake.RemoveAffiliationStub = nil
	if fake.removeAffiliationReturnsOnCall == nil {
		fake.removeAffiliationReturnsOnCall = make(map[int]struct {
			result1 *msp.AffiliationResponse
			result2 error
		})
	}
	fake.removeAffiliationReturnsOnCall[i] = struct {
		result1 *msp.AffiliationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) RemoveIdentity(request *msp.RemoveIdentityRequest) (*msp.IdentityResponse, error) {
	fake.removeIdentityMutex.Lock()
	ret, specificReturn := fake.removeIdentityReturnsOnCall[len(fake.removeIdentityArgsForCall)]
	fake.removeIdentityArgsForCall = append(fake.removeIdentityArgsForCall, struct {
		request *msp.RemoveIdentityRequest
	}{request})
	fake.recordInvocation("RemoveIdentity", []interface{}{request})
	fake.removeIdentityMutex.Unlock()
	if fake.RemoveIdentityStub != nil {
		return fake.RemoveIdentityStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.removeIdentityReturns.result1, fake.removeIdentityReturns.result2
}

func (fake *MSP) RemoveIdentityCallCount() int {
	fake.removeIdentityMutex.RLock()
	defer fake.removeIdentityMutex.RUnlock()
	return len(fake.removeIdentityArgsForCall)
}

func (fake *MSP) RemoveIdentityArgsForCall(i int) *msp.RemoveIdentityRequest {
	fake.removeIdentityMutex.RLock()
	defer fake.removeIdentityMutex.RUnlock()
	return fake.removeIdentityArgsForCall[i].request
}

func (fake *MSP) RemoveIdentityReturns(result1 *msp.IdentityResponse, result2 error) {
	fake.RemoveIdentityStub = nil
	fake.removeIdentityReturns = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) RemoveIdentityReturnsOnCall(i int, result1 *msp.IdentityResponse, result2 error) {
	fake.RemoveIdentityStub = nil
	if fake.removeIdentityReturnsOnCall == nil {
		fake.removeIdentityReturnsOnCall = make(map[int]struct {
			result1 *msp.IdentityResponse
			result2 error
		})
	}
	fake.removeIdentityReturnsOnCall[i] = struct {
		result1 *msp.IdentityResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) Revoke(request *msp.RevocationRequest) (*msp.RevocationResponse, error) {
	fake.revokeMutex.Lock()
	ret, specificReturn := fake.revokeReturnsOnCall[len(fake.revokeArgsForCall)]
	fake.revokeArgsForCall = append(fake.revokeArgsForCall, struct {
		request *msp.RevocationRequest
	}{request})
	fake.recordInvocation("Revoke", []interface{}{request})
	fake.revokeMutex.Unlock()
	if fake.RevokeStub != nil {
		return fake.RevokeStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.revokeReturns.result1, fake.revokeReturns.result2
}

func (fake *MSP) RevokeCallCount() int {
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	return len(fake.revokeArgsForCall)
}

func (fake *MSP) RevokeArgsForCall(i int) *msp.RevocationRequest {
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	return fake.revokeArgsForCall[i].request
}

func (fake *MSP) RevokeReturns(result1 *msp.RevocationResponse, result2 error) {
	fake.RevokeStub = nil
	fake.revokeReturns = struct {
		result1 *msp.RevocationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) RevokeReturnsOnCall(i int, result1 *msp.RevocationResponse, result2 error) {
	fake.RevokeStub = nil
	if fake.revokeReturnsOnCall == nil {
		fake.revokeReturnsOnCall = make(map[int]struct {
			result1 *msp.RevocationResponse
			result2 error
		})
	}
	fake.revokeReturnsOnCall[i] = struct {
		result1 *msp.RevocationResponse
		result2 error
	}{result1, result2}
}

func (fake *MSP) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addAffiliationMutex.RLock()
	defer fake.addAffiliationMutex.RUnlock()
	fake.createIdentityMutex.RLock()
	defer fake.createIdentityMutex.RUnlock()
	fake.createSigningIdentityMutex.RLock()
	defer fake.createSigningIdentityMutex.RUnlock()
	fake.enrollMutex.RLock()
	defer fake.enrollMutex.RUnlock()
	fake.getAffiliationMutex.RLock()
	defer fake.getAffiliationMutex.RUnlock()
	fake.getAllAffiliationsMutex.RLock()
	defer fake.getAllAffiliationsMutex.RUnlock()
	fake.getAllIdentitiesMutex.RLock()
	defer fake.getAllIdentitiesMutex.RUnlock()
	fake.getCAInfoMutex.RLock()
	defer fake.getCAInfoMutex.RUnlock()
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	fake.modifyAffiliationMutex.RLock()
	defer fake.modifyAffiliationMutex.RUnlock()
	fake.modifyIdentityMutex.RLock()
	defer fake.modifyIdentityMutex.RUnlock()
	fake.reenrollMutex.RLock()
	defer fake.reenrollMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.removeAffiliationMutex.RLock()
	defer fake.removeAffiliationMutex.RUnlock()
	fake.removeIdentityMutex.RLock()
	defer fake.removeIdentityMutex.RUnlock()
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MSP) recordInvocation(key string, args []interface{}) {
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

var _ fabric.MSP = new(MSP)
