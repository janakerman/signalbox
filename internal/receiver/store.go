package receiver

import (
	"fmt"
	"strings"
)

type NotificationStore interface {
	GetByObject(kind, namespace, name, revision string) ([]Notification, error)
	GetByRevision(revision string) ([]Notification, error)
	GetByHash(revision string) ([]Notification, error)
	Write(n Notification) error
}

var _ NotificationStore = (*InMemoryStore)(nil)

type notificationKey struct {
	kind, namespace, name, revision string
}

type InMemoryStore struct {
	byObject   map[notificationKey][]Notification
	byRevision map[string][]Notification
	byHash     map[string][]Notification
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		byObject:   map[notificationKey][]Notification{},
		byRevision: map[string][]Notification{},
		byHash:     map[string][]Notification{},
	}
}

func (s *InMemoryStore) GetByObject(kind, namespace, name, revision string) ([]Notification, error) {
	key := notificationKey{
		kind:      kind,
		namespace: namespace,
		name:      name,
		revision:  revision,
	}

	ns := s.byObject[key]
	if ns == nil {
		ns = []Notification{}
	}

	return ns, nil
}

func (s *InMemoryStore) GetByRevision(revision string) ([]Notification, error) {
	ns := s.byRevision[revision]
	if ns == nil {
		ns = []Notification{}
	}
	return ns, nil
}

func (s *InMemoryStore) GetByHash(revision string) ([]Notification, error) {
	ns := s.byHash[revision]
	if ns == nil {
		ns = []Notification{}
	}
	return ns, nil
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

	s.byObject[key] = append(s.byObject[key], n)
	s.byRevision[revision] = append(s.byRevision[revision], n)

	split := strings.Split(revision, "/")
	if len(split) != 2 {
		return fmt.Errorf("unexpected revision format '%s'", revision)
	}

	s.byHash[split[1]] = append(s.byHash[revision], n)

	return nil
}
