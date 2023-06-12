package gcs

type InMemoryStore struct{}

func NewInMemoryStore() Client { return &InMemoryStore{} }

func (m *InMemoryStore) UploadSpec(string, string, []byte) error { return nil }

func (m *InMemoryStore) Read(string) ([]byte, error) { return []byte{}, nil }

func (m *InMemoryStore) Close() error { return nil }
