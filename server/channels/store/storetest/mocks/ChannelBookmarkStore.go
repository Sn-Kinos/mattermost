// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ChannelBookmarkStore is an autogenerated mock type for the ChannelBookmarkStore type
type ChannelBookmarkStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: bookmarkId
func (_m *ChannelBookmarkStore) Delete(bookmarkId string) error {
	ret := _m.Called(bookmarkId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bookmarkId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ErrorIfBookmarkFileInfoAlreadyAttached provides a mock function with given fields: fileId
func (_m *ChannelBookmarkStore) ErrorIfBookmarkFileInfoAlreadyAttached(fileId string) error {
	ret := _m.Called(fileId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: Id, includeDeleted
func (_m *ChannelBookmarkStore) Get(Id string, includeDeleted bool) (*model.ChannelBookmarkWithFileInfo, error) {
	ret := _m.Called(Id, includeDeleted)

	var r0 *model.ChannelBookmarkWithFileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (*model.ChannelBookmarkWithFileInfo, error)); ok {
		return rf(Id, includeDeleted)
	}
	if rf, ok := ret.Get(0).(func(string, bool) *model.ChannelBookmarkWithFileInfo); ok {
		r0 = rf(Id, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ChannelBookmarkWithFileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(Id, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookmarksForChannelSince provides a mock function with given fields: channelId, since
func (_m *ChannelBookmarkStore) GetBookmarksForChannelSince(channelId string, since int64) ([]*model.ChannelBookmarkWithFileInfo, error) {
	ret := _m.Called(channelId, since)

	var r0 []*model.ChannelBookmarkWithFileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) ([]*model.ChannelBookmarkWithFileInfo, error)); ok {
		return rf(channelId, since)
	}
	if rf, ok := ret.Get(0).(func(string, int64) []*model.ChannelBookmarkWithFileInfo); ok {
		r0 = rf(channelId, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelBookmarkWithFileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(channelId, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: bookmark, increaseSortOrder
func (_m *ChannelBookmarkStore) Save(bookmark *model.ChannelBookmark, increaseSortOrder bool) (*model.ChannelBookmarkWithFileInfo, error) {
	ret := _m.Called(bookmark, increaseSortOrder)

	var r0 *model.ChannelBookmarkWithFileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ChannelBookmark, bool) (*model.ChannelBookmarkWithFileInfo, error)); ok {
		return rf(bookmark, increaseSortOrder)
	}
	if rf, ok := ret.Get(0).(func(*model.ChannelBookmark, bool) *model.ChannelBookmarkWithFileInfo); ok {
		r0 = rf(bookmark, increaseSortOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ChannelBookmarkWithFileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ChannelBookmark, bool) error); ok {
		r1 = rf(bookmark, increaseSortOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: bookmark
func (_m *ChannelBookmarkStore) Update(bookmark *model.ChannelBookmark) error {
	ret := _m.Called(bookmark)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ChannelBookmark) error); ok {
		r0 = rf(bookmark)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSortOrder provides a mock function with given fields: bookmarkId, channelId, newIndex
func (_m *ChannelBookmarkStore) UpdateSortOrder(bookmarkId string, channelId string, newIndex int64) ([]*model.ChannelBookmarkWithFileInfo, error) {
	ret := _m.Called(bookmarkId, channelId, newIndex)

	var r0 []*model.ChannelBookmarkWithFileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) ([]*model.ChannelBookmarkWithFileInfo, error)); ok {
		return rf(bookmarkId, channelId, newIndex)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) []*model.ChannelBookmarkWithFileInfo); ok {
		r0 = rf(bookmarkId, channelId, newIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelBookmarkWithFileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(bookmarkId, channelId, newIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewChannelBookmarkStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewChannelBookmarkStore creates a new instance of ChannelBookmarkStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChannelBookmarkStore(t mockConstructorTestingTNewChannelBookmarkStore) *ChannelBookmarkStore {
	mock := &ChannelBookmarkStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
