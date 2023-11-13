package gcs

import "github.com/getkin/kin-openapi/openapi3"

type InMemoryStore struct{ pathToFile map[string][]byte }

func NewInMemoryStore(pathToFile map[string][]byte) Client {

	return &InMemoryStore{pathToFile: pathToFile}
}

func (m *InMemoryStore) UploadSpec(tenant string, webhook string, name string, spec *openapi3.T) error {

	return nil
}

func (m *InMemoryStore) Upload(string, []byte) error { return nil }

func (m *InMemoryStore) Read(path string) ([]byte, error) {

	return m.pathToFile[path], nil
}

func (m *InMemoryStore) Close() error { return nil }
