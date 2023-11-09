package gcs

type InMemoryStore struct{ pathToFile map[string][]byte }

func NewInMemoryStore(pathToFile map[string][]byte) Client {

	return &InMemoryStore{pathToFile: pathToFile}
}

func (m *InMemoryStore) UploadSpec(string, string, []byte) error { return nil }

func (m *InMemoryStore) Read(path string) ([]byte, error) {

	return m.pathToFile[path], nil
}

func (m *InMemoryStore) Close() error { return nil }
