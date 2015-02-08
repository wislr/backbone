// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package app

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of ControllerInterface interface
type MockControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerInterfaceRecorder
}

// Recorder for MockControllerInterface (not exported)
type _MockControllerInterfaceRecorder struct {
	mock *MockControllerInterface
}

func NewMockControllerInterface(ctrl *gomock.Controller) *MockControllerInterface {
	mock := &MockControllerInterface{ctrl: ctrl}
	mock.recorder = &_MockControllerInterfaceRecorder{mock}
	return mock
}

func (_m *MockControllerInterface) EXPECT() *_MockControllerInterfaceRecorder {
	return _m.recorder
}

// Mock of ResourceHandlerInterface interface
type MockResourceHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockResourceHandlerInterfaceRecorder
}

// Recorder for MockResourceHandlerInterface (not exported)
type _MockResourceHandlerInterfaceRecorder struct {
	mock *MockResourceHandlerInterface
}

func NewMockResourceHandlerInterface(ctrl *gomock.Controller) *MockResourceHandlerInterface {
	mock := &MockResourceHandlerInterface{ctrl: ctrl}
	mock.recorder = &_MockResourceHandlerInterfaceRecorder{mock}
	return mock
}

func (_m *MockResourceHandlerInterface) EXPECT() *_MockResourceHandlerInterfaceRecorder {
	return _m.recorder
}

func (_m *MockResourceHandlerInterface) Download(url string, target string) error {
	ret := _m.ctrl.Call(_m, "Download", url, target)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceHandlerInterfaceRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Download", arg0, arg1)
}

func (_m *MockResourceHandlerInterface) Upload(source string, url string) error {
	ret := _m.ctrl.Call(_m, "Upload", source, url)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceHandlerInterfaceRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Upload", arg0, arg1)
}

// Mock of HostInterface interface
type MockHostInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockHostInterfaceRecorder
}

// Recorder for MockHostInterface (not exported)
type _MockHostInterfaceRecorder struct {
	mock *MockHostInterface
}

func NewMockHostInterface(ctrl *gomock.Controller) *MockHostInterface {
	mock := &MockHostInterface{ctrl: ctrl}
	mock.recorder = &_MockHostInterfaceRecorder{mock}
	return mock
}

func (_m *MockHostInterface) EXPECT() *_MockHostInterfaceRecorder {
	return _m.recorder
}

func (_m *MockHostInterface) StorageAvailable() (int64, error) {
	ret := _m.ctrl.Call(_m, "StorageAvailable")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHostInterfaceRecorder) StorageAvailable() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StorageAvailable")
}

func (_m *MockHostInterface) ClearData(app *App) error {
	ret := _m.ctrl.Call(_m, "ClearData", app)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) ClearData(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClearData", arg0)
}

func (_m *MockHostInterface) RestoreData(app *App, dataFile string) error {
	ret := _m.ctrl.Call(_m, "RestoreData", app, dataFile)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) RestoreData(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RestoreData", arg0, arg1)
}

func (_m *MockHostInterface) BackupData(app *App) (string, error) {
	ret := _m.ctrl.Call(_m, "BackupData", app)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHostInterfaceRecorder) BackupData(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BackupData", arg0)
}

func (_m *MockHostInterface) AppExtensionsDir(bundle_id string) string {
	ret := _m.ctrl.Call(_m, "AppExtensionsDir", bundle_id)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) AppExtensionsDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppExtensionsDir", arg0)
}

func (_m *MockHostInterface) Install(app *App, path string) error {
	ret := _m.ctrl.Call(_m, "Install", app, path)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Install", arg0, arg1)
}

func (_m *MockHostInterface) Uninstall(app *App) error {
	ret := _m.ctrl.Call(_m, "Uninstall", app)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) Uninstall(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Uninstall", arg0)
}

func (_m *MockHostInterface) IsInstalled(app *App) (bool, error) {
	ret := _m.ctrl.Call(_m, "IsInstalled", app)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHostInterfaceRecorder) IsInstalled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsInstalled", arg0)
}

func (_m *MockHostInterface) SetEnabled(app *App, enabled bool) error {
	ret := _m.ctrl.Call(_m, "SetEnabled", app, enabled)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) SetEnabled(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetEnabled", arg0, arg1)
}

func (_m *MockHostInterface) StartApp(app *App, params AppParams) error {
	ret := _m.ctrl.Call(_m, "StartApp", app, params)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) StartApp(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartApp", arg0, arg1)
}

func (_m *MockHostInterface) StopApp(app *App) error {
	ret := _m.ctrl.Call(_m, "StopApp", app)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) StopApp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopApp", arg0)
}

func (_m *MockHostInterface) AppPid(app *App) (int, error) {
	ret := _m.ctrl.Call(_m, "AppPid", app)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHostInterfaceRecorder) AppPid(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AppPid", arg0)
}

func (_m *MockHostInterface) PausePid(pid int) error {
	ret := _m.ctrl.Call(_m, "PausePid", pid)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) PausePid(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PausePid", arg0)
}

func (_m *MockHostInterface) ResumePid(pid int) error {
	ret := _m.ctrl.Call(_m, "ResumePid", pid)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) ResumePid(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResumePid", arg0)
}

func (_m *MockHostInterface) SetInternetEnabled(enabled bool) error {
	ret := _m.ctrl.Call(_m, "SetInternetEnabled", enabled)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) SetInternetEnabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetInternetEnabled", arg0)
}

func (_m *MockHostInterface) RemoveFiles(files ...string) error {
	_s := []interface{}{}
	for _, _x := range files {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RemoveFiles", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) RemoveFiles(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveFiles", arg0...)
}

func (_m *MockHostInterface) SetAndroidId(androidId string) error {
	ret := _m.ctrl.Call(_m, "SetAndroidId", androidId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockHostInterfaceRecorder) SetAndroidId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAndroidId", arg0)
}
