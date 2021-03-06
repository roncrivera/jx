// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/util (interfaces: CommandInterface)

package util_test

import (
	backoff "github.com/cenkalti/backoff"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	time "time"
)

type MockCommandInterface struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommandInterface() *MockCommandInterface {
	return &MockCommandInterface{fail: pegomock.GlobalFailHandler}
}

func (mock *MockCommandInterface) DidError() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidError", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommandInterface) DidFail() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidFail", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommandInterface) Error() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Error", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommandInterface) Run() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
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

func (mock *MockCommandInterface) RunWithoutRetry() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunWithoutRetry", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
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

func (mock *MockCommandInterface) SetArgs(_param0 []string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetArgs", params, []reflect.Type{})
}

func (mock *MockCommandInterface) SetDir(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetDir", params, []reflect.Type{})
}

func (mock *MockCommandInterface) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetExponentialBackOff", params, []reflect.Type{})
}

func (mock *MockCommandInterface) SetName(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetName", params, []reflect.Type{})
}

func (mock *MockCommandInterface) SetTimeout(_param0 time.Duration) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandInterface().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetTimeout", params, []reflect.Type{})
}

func (mock *MockCommandInterface) VerifyWasCalledOnce() *VerifierCommandInterface {
	return &VerifierCommandInterface{mock, pegomock.Times(1), nil}
}

func (mock *MockCommandInterface) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierCommandInterface {
	return &VerifierCommandInterface{mock, invocationCountMatcher, nil}
}

func (mock *MockCommandInterface) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierCommandInterface {
	return &VerifierCommandInterface{mock, invocationCountMatcher, inOrderContext}
}

type VerifierCommandInterface struct {
	mock                   *MockCommandInterface
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierCommandInterface) DidError() *CommandInterface_DidError_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidError", params)
	return &CommandInterface_DidError_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_DidError_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_DidError_OngoingVerification) GetCapturedArguments() {
}

func (c *CommandInterface_DidError_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommandInterface) DidFail() *CommandInterface_DidFail_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidFail", params)
	return &CommandInterface_DidFail_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_DidFail_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_DidFail_OngoingVerification) GetCapturedArguments() {
}

func (c *CommandInterface_DidFail_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommandInterface) Error() *CommandInterface_Error_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Error", params)
	return &CommandInterface_Error_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_Error_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_Error_OngoingVerification) GetCapturedArguments() {
}

func (c *CommandInterface_Error_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommandInterface) Run() *CommandInterface_Run_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params)
	return &CommandInterface_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_Run_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_Run_OngoingVerification) GetCapturedArguments() {
}

func (c *CommandInterface_Run_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommandInterface) RunWithoutRetry() *CommandInterface_RunWithoutRetry_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunWithoutRetry", params)
	return &CommandInterface_RunWithoutRetry_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_RunWithoutRetry_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_RunWithoutRetry_OngoingVerification) GetCapturedArguments() {
}

func (c *CommandInterface_RunWithoutRetry_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierCommandInterface) SetArgs(_param0 []string) *CommandInterface_SetArgs_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetArgs", params)
	return &CommandInterface_SetArgs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_SetArgs_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_SetArgs_OngoingVerification) GetCapturedArguments() []string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *CommandInterface_SetArgs_OngoingVerification) GetAllCapturedArguments() (_param0 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierCommandInterface) SetDir(_param0 string) *CommandInterface_SetDir_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetDir", params)
	return &CommandInterface_SetDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_SetDir_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_SetDir_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *CommandInterface_SetDir_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierCommandInterface) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) *CommandInterface_SetExponentialBackOff_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetExponentialBackOff", params)
	return &CommandInterface_SetExponentialBackOff_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_SetExponentialBackOff_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_SetExponentialBackOff_OngoingVerification) GetCapturedArguments() *backoff.ExponentialBackOff {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *CommandInterface_SetExponentialBackOff_OngoingVerification) GetAllCapturedArguments() (_param0 []*backoff.ExponentialBackOff) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*backoff.ExponentialBackOff, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*backoff.ExponentialBackOff)
		}
	}
	return
}

func (verifier *VerifierCommandInterface) SetName(_param0 string) *CommandInterface_SetName_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetName", params)
	return &CommandInterface_SetName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_SetName_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_SetName_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *CommandInterface_SetName_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierCommandInterface) SetTimeout(_param0 time.Duration) *CommandInterface_SetTimeout_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetTimeout", params)
	return &CommandInterface_SetTimeout_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommandInterface_SetTimeout_OngoingVerification struct {
	mock              *MockCommandInterface
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommandInterface_SetTimeout_OngoingVerification) GetCapturedArguments() time.Duration {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *CommandInterface_SetTimeout_OngoingVerification) GetAllCapturedArguments() (_param0 []time.Duration) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]time.Duration, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(time.Duration)
		}
	}
	return
}
