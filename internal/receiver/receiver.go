package receiver

import "strings"

type NotificationStore interface {
	Get(kind, namespace, name, revision string) ([]Notification, error)
	Write(n Notification) error
}

var _ NotificationStore = (*InMemoryStore)(nil)

type notificationKey struct {
	kind, namespace, name, revision string
}

type InMemoryStore struct {
	s map[notificationKey][]Notification
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		s: map[notificationKey][]Notification{},
	}
}

func (s *InMemoryStore) Get(kind, namespace, name, revision string) ([]Notification, error) {
	key := notificationKey{
		kind:      kind,
		namespace: namespace,
		name:      name,
		revision:  revision,
	}
	return s.s[key], nil
}

func (s *InMemoryStore) Write(n Notification) error {
	var revision string

	switch n.Object.Kind {
	case "GitRepository":
		revision = strings.TrimPrefix(n.Message, "Fetched revision: ")
	case "Kustomization":
		revision = n.Metadata["revision"]
	}

	key := notificationKey{
		kind:      n.Object.Kind,
		namespace: n.Object.Namespace,
		name:      n.Object.Name,
		revision:  revision,
	}

	s.s[key] = append(s.s[key], n)
	return nil
}
