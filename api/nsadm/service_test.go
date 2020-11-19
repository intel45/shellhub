package nsadm

import (
	"context"
	"testing"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/api/store/mocks"
	"github.com/shellhub-io/shellhub/pkg/api/paginator"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestListNamespaces(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	namespaces := []models.Namespace{
		{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713"},
		{Name: "group2", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf48471i4"},
	}
	query := paginator.Query{Page: 1, PerPage: 10}
	mock.On("ListNamespaces", ctx, query, []models.Filter(nil), false).Return(namespaces, len(namespaces), nil).Once()
	returnedNamespaces, count, err := s.ListNamespaces(ctx, query, "", false)
	assert.NoError(t, err)
	assert.Equal(t, namespaces, returnedNamespaces)
	assert.Equal(t, count, len(namespaces))
	mock.AssertExpectations(t)
}

func TestGetNamespace(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713"}

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()

	returnNamespace, err := s.GetNamespace(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, namespace, returnNamespace)

	mock.AssertExpectations(t)
}

func TestCreateNamespace(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	user := &models.User{Name: "user1", Username: "hash1", ID: "hash1"}

	namespace := &models.Namespace{Name: "group1", Owner: "hash1"}

	mock.On("GetUserByUsername", ctx, user.Username).Return(user, nil).Once()
	mock.On("CreateNamespace", ctx, namespace).Return(namespace, nil).Once()

	returnedNamespace, err := s.CreateNamespace(ctx, namespace, namespace.Owner)
	assert.NoError(t, err)
	assert.Equal(t, namespace, returnedNamespace)
	mock.AssertExpectations(t)
}

func TestEditNamespace(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	user := &models.User{Name: "user1", Username: "hash1", ID: "hash1"}

	newName := "newname"
	namespace := &models.Namespace{Name: "oldname", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713"}
	namespaceWithNewName := &models.Namespace{Name: "newname", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713"}

	mock.On("GetUserByUsername", ctx, user.Username).Return(user, nil).Once()
	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Twice()
	mock.On("EditNamespace", ctx, namespace.TenantID, newName).Return(namespaceWithNewName, nil).Once()
	_, err := s.EditNamespace(ctx, namespace.TenantID, newName, namespace.Owner)

	assert.NoError(t, err)
	returnedNamespace, err := s.GetNamespace(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, newName, returnedNamespace.Name)
	mock.AssertExpectations(t)
}

func TestDeleteNamespace(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	user := &models.User{Name: "user1", Username: "hash1", ID: "hash1"}
	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713"}

	mock.On("GetUserByUsername", ctx, user.Username).Return(user, nil).Once()
	mock.On("DeleteNamespace", ctx, namespace.TenantID).Return(nil).Once()
	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()

	err := s.DeleteNamespace(ctx, namespace.TenantID, namespace.Owner)
	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func TestAddNamespaceUser(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))
	ctx := context.TODO()
	user := &models.User{Name: "user1", Username: "username1", ID: "hash1"}
	member := &models.User{Name: "user2", Username: "username2", ID: "hash2"}
	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Members: []string{"hash1"}}
	namespace2 := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Members: []string{"hash1", "hash2"}}

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()
	mock.On("GetUserByUsername", ctx, user.Username).Return(user, nil).Once()
	mock.On("GetUserByUsername", ctx, member.Username).Return(member, nil).Once()
	mock.On("AddNamespaceUser", ctx, namespace.TenantID, member.ID).Return(namespace2, nil).Once()

	_, err := s.AddNamespaceUser(ctx, namespace.TenantID, member.Username, user.Username)
	assert.NoError(t, err)

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace2, nil).Once()

	returnedNamespace, err := s.GetNamespace(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, namespace2, returnedNamespace)
}

func TestRemoveNamespaceUser(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))
	ctx := context.TODO()
	user := &models.User{Name: "user1", Username: "username1", ID: "hash1"}
	member := &models.User{Name: "user2", Username: "username2", ID: "hash2"}
	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Members: []string{"hash1", "hash2"}}
	namespace2 := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Members: []string{"hash1"}}

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()
	mock.On("GetUserByUsername", ctx, user.Username).Return(user, nil).Once()
	mock.On("GetUserByUsername", ctx, member.Username).Return(member, nil).Once()
	mock.On("RemoveNamespaceUser", ctx, namespace.TenantID, member.ID).Return(namespace2, nil).Once()

	_, err := s.RemoveNamespaceUser(ctx, namespace.TenantID, member.Username, user.Username)
	assert.NoError(t, err)

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace2, nil).Once()

	returnedNamespace, err := s.GetNamespace(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, namespace2, returnedNamespace)
}

func TestGetDataUserSecurity(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Settings: &models.NamespaceSettings{SessionRecord: true}}

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()
	mock.On("GetDataUserSecurity", ctx, namespace.TenantID).
		Return(namespace.Settings.SessionRecord, nil).Once()

	returnedUserSecurity, err := s.GetDataUserSecurity(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, returnedUserSecurity, namespace.Settings.SessionRecord)

	mock.AssertExpectations(t)
}

func TestUpdateDataUserSecurity(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	namespace := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Settings: &models.NamespaceSettings{SessionRecord: true}}
	namespace2 := &models.Namespace{Name: "group1", Owner: "hash1", TenantID: "a736a52b-5777-4f92-b0b8-e359bf484713", Settings: &models.NamespaceSettings{SessionRecord: false}}

	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace, nil).Once()
	mock.On("GetNamespace", ctx, namespace.TenantID).Return(namespace2, nil).Once()
	mock.On("UpdateDataUserSecurity", ctx, !namespace.Settings.SessionRecord, namespace.TenantID).
		Return(nil).Once()
	mock.On("GetDataUserSecurity", ctx, namespace.TenantID).
		Return(!namespace.Settings.SessionRecord, nil).Once()

	err := s.UpdateDataUserSecurity(ctx, !namespace.Settings.SessionRecord, namespace.TenantID)
	assert.NoError(t, err)

	returnedUserSecurity, err := s.GetDataUserSecurity(ctx, namespace.TenantID)
	assert.NoError(t, err)
	assert.Equal(t, !namespace.Settings.SessionRecord, returnedUserSecurity)

	mock.AssertExpectations(t)
}
