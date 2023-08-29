package util

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func ToInterfaceSlice(slice interface{}) []interface{} {

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		log.Fatalf("failed to cast interface into slice with 'given a non-slice type %s'", reflect.TypeOf(slice).String())
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
