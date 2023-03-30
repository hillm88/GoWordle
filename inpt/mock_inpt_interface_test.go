// Code generated by pegomock. DO NOT EDIT.
// Source: inpt_interface.go

package inpt

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockprivateMethodInterface struct {
	fail func(message string, callerSkip ...int)
}

func NewMockprivateMethodInterface(options ...pegomock.Option) *MockprivateMethodInterface {
	mock := &MockprivateMethodInterface{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockprivateMethodInterface) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockprivateMethodInterface) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockprivateMethodInterface) getInput() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockprivateMethodInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("getInput", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockprivateMethodInterface) isValidInput(input string, exactLength int) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockprivateMethodInterface().")
	}
	params := []pegomock.Param{input, exactLength}
	result := pegomock.GetGenericMockFrom(mock).Invoke("isValidInput", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockprivateMethodInterface) retrieveCLIInput() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockprivateMethodInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("retrieveCLIInput", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockprivateMethodInterface) VerifyWasCalledOnce() *VerifierMockprivateMethodInterface {
	return &VerifierMockprivateMethodInterface{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockprivateMethodInterface) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockprivateMethodInterface {
	return &VerifierMockprivateMethodInterface{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockprivateMethodInterface) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockprivateMethodInterface {
	return &VerifierMockprivateMethodInterface{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockprivateMethodInterface) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockprivateMethodInterface {
	return &VerifierMockprivateMethodInterface{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockprivateMethodInterface struct {
	mock                   *MockprivateMethodInterface
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockprivateMethodInterface) getInput() *MockprivateMethodInterface_getInput_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "getInput", params, verifier.timeout)
	return &MockprivateMethodInterface_getInput_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockprivateMethodInterface_getInput_OngoingVerification struct {
	mock              *MockprivateMethodInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockprivateMethodInterface_getInput_OngoingVerification) GetCapturedArguments() {
}

func (c *MockprivateMethodInterface_getInput_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockprivateMethodInterface) isValidInput(input string, exactLength int) *MockprivateMethodInterface_isValidInput_OngoingVerification {
	params := []pegomock.Param{input, exactLength}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "isValidInput", params, verifier.timeout)
	return &MockprivateMethodInterface_isValidInput_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockprivateMethodInterface_isValidInput_OngoingVerification struct {
	mock              *MockprivateMethodInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockprivateMethodInterface_isValidInput_OngoingVerification) GetCapturedArguments() (string, int) {
	input, exactLength := c.GetAllCapturedArguments()
	return input[len(input)-1], exactLength[len(exactLength)-1]
}

func (c *MockprivateMethodInterface_isValidInput_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockprivateMethodInterface) retrieveCLIInput() *MockprivateMethodInterface_retrieveCLIInput_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "retrieveCLIInput", params, verifier.timeout)
	return &MockprivateMethodInterface_retrieveCLIInput_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockprivateMethodInterface_retrieveCLIInput_OngoingVerification struct {
	mock              *MockprivateMethodInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockprivateMethodInterface_retrieveCLIInput_OngoingVerification) GetCapturedArguments() {
}

func (c *MockprivateMethodInterface_retrieveCLIInput_OngoingVerification) GetAllCapturedArguments() {
}