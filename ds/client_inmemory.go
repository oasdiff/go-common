package ds

import "reflect"

type InMemoryClient struct {
	payload map[Kind]interface{}
}

func NewInMemoryClient(data map[Kind]interface{}) Client {

	if data == nil {
		data = make(map[Kind]interface{})
	}

	return &InMemoryClient{payload: data}
}

func (c *InMemoryClient) Get(kind Kind, id string, dst interface{}) error {

	if kind == KindTenant {
		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(Tenant{
			Id:   id,
			Name: "test-123",
		}))
	}

	return nil
}

func (c *InMemoryClient) GetAll(kind Kind, dst interface{}) error {

	if res, ok := c.payload[kind]; ok {
		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(res))
	}

	return nil
}

func (c *InMemoryClient) Put(kind Kind, id string, src interface{}) error { return nil }
