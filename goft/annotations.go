package goft

import (
	"reflect"
)

type IAnnotation interface {
	SetTag(reflect.StructTag)
	String() string
}

var IAnnotationList []IAnnotation

func IsAnnotation(t reflect.Type) bool {
	for _, anno := range IAnnotationList {
		if reflect.TypeOf(anno) == t {
			return true
		}
	}

	return false
}

func init() {
	IAnnotationList = make([]IAnnotation, 0)
}
