// Code generated by MockGen. DO NOT EDIT.
// Source: loop.go
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mocks.go -package=mocks -source=loop.go -typed
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	scrape "github.com/florianloch/prom2mqtt/internal/scrape"
	gomock "go.uber.org/mock/gomock"
)

// MockMQTTer is a mock of MQTTer interface.
type MockMQTTer struct {
	ctrl     *gomock.Controller
	recorder *MockMQTTerMockRecorder
}

// MockMQTTerMockRecorder is the mock recorder for MockMQTTer.
type MockMQTTerMockRecorder struct {
	mock *MockMQTTer
}

// NewMockMQTTer creates a new mock instance.
func NewMockMQTTer(ctrl *gomock.Controller) *MockMQTTer {
	mock := &MockMQTTer{ctrl: ctrl}
	mock.recorder = &MockMQTTerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMQTTer) EXPECT() *MockMQTTerMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockMQTTer) Publish(ctx context.Context, topic, payload string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, topic, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockMQTTerMockRecorder) Publish(ctx, topic, payload any) *MQTTerPublishCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockMQTTer)(nil).Publish), ctx, topic, payload)
	return &MQTTerPublishCall{Call: call}
}

// MQTTerPublishCall wrap *gomock.Call
type MQTTerPublishCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MQTTerPublishCall) Return(arg0 error) *MQTTerPublishCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MQTTerPublishCall) Do(f func(context.Context, string, string) error) *MQTTerPublishCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MQTTerPublishCall) DoAndReturn(f func(context.Context, string, string) error) *MQTTerPublishCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockScraper is a mock of Scraper interface.
type MockScraper struct {
	ctrl     *gomock.Controller
	recorder *MockScraperMockRecorder
}

// MockScraperMockRecorder is the mock recorder for MockScraper.
type MockScraperMockRecorder struct {
	mock *MockScraper
}

// NewMockScraper creates a new mock instance.
func NewMockScraper(ctrl *gomock.Controller) *MockScraper {
	mock := &MockScraper{ctrl: ctrl}
	mock.recorder = &MockScraperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScraper) EXPECT() *MockScraperMockRecorder {
	return m.recorder
}

// ScrapeURL mocks base method.
func (m *MockScraper) ScrapeURL(ctx context.Context, targetURL string) (*scrape.Metrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrapeURL", ctx, targetURL)
	ret0, _ := ret[0].(*scrape.Metrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScrapeURL indicates an expected call of ScrapeURL.
func (mr *MockScraperMockRecorder) ScrapeURL(ctx, targetURL any) *ScraperScrapeURLCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrapeURL", reflect.TypeOf((*MockScraper)(nil).ScrapeURL), ctx, targetURL)
	return &ScraperScrapeURLCall{Call: call}
}

// ScraperScrapeURLCall wrap *gomock.Call
type ScraperScrapeURLCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ScraperScrapeURLCall) Return(arg0 *scrape.Metrics, arg1 error) *ScraperScrapeURLCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ScraperScrapeURLCall) Do(f func(context.Context, string) (*scrape.Metrics, error)) *ScraperScrapeURLCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ScraperScrapeURLCall) DoAndReturn(f func(context.Context, string) (*scrape.Metrics, error)) *ScraperScrapeURLCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}