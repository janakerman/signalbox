package receiver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	GitRepositoryEvent = Notification{
		Object: InvolvedObject{
			Kind:            "GitRepository",
			Namespace:       "flux-system",
			Name:            "podinfo",
			Uid:             "708ae727-68cb-48ae-922f-fb06162be279",
			ApiVersion:      "source.toolkit.fluxcd.io/v1beta1",
			ResourceVersion: "28730",
		},
		Severity:            "info",
		Timestamp:           "2021-12-19T16:16:36Z",
		Message:             "Fetched revision: v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9",
		Reason:              "info",
		ReportingController: "source-controller",
		ReportingInstance:   "source-controller-565f8fbbff-j4sjz",
	}
)

func Test_InMemoryStore_Empty(t *testing.T) {
	s := NewInMemoryStore()

	ns, err := s.GetByObject("kind", "namespace", "name", "revision")
	require.NoError(t, err)
	assert.Len(t, ns, 0)

	ns, err = s.GetByRevision("revision")
	require.NoError(t, err)
	assert.Len(t, ns, 0)

	ns, err = s.GetByHash("revision")
	require.NoError(t, err)
	assert.Len(t, ns, 0)
}

func Test_InMemoryStore_Write(t *testing.T) {
	s := NewInMemoryStore()

	err := s.Write(GitRepositoryEvent)
	require.NoError(t, err)

	ns, err := s.GetByRevision("v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, ns, 1)
}

func Test_InMemoryStore_GetByRevision(t *testing.T) {
	s := NewInMemoryStore()

	err := s.Write(GitRepositoryEvent)
	require.NoError(t, err)

	ns, err := s.GetByRevision("v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, ns, 1)
	assert.Equal(t, GitRepositoryEvent, ns[0])
}

func Test_InMemoryStore_GetByHash(t *testing.T) {
	s := NewInMemoryStore()

	err := s.Write(GitRepositoryEvent)
	require.NoError(t, err)

	ns, err := s.GetByHash("6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, ns, 1)
	assert.Equal(t, GitRepositoryEvent, ns[0])
}
