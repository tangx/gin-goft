package goft

import "reflect"

type Annotation interface {
	SetTag(reflect.StructTag)
}

var AnnotationList []Annotation

func IsAnnotation(t reflect.Type) bool {
	for _, anno := range AnnotationList {
		if reflect.TypeOf(anno) == t {
			return true
		}
	}

	return false
}

func init() {
	AnnotationList = make([]Annotation, 0)
	AnnotationList = append(AnnotationList, new(Value))
}

type Value struct {
	tag reflect.StructTag
}

func (value *Value) SetTag(tag reflect.StructTag) {
	value.tag = tag
}

func (value *Value) String() string {
	return "32"
}
