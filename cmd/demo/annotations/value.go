package annotations

import (
	"reflect"

	"github.com/tangx-labs/gin-goft/goft"
)

var _ goft.IAnnotation = &Value{}

type Value struct {
	tag reflect.StructTag
}

func NewValue() *Value {
	return &Value{}
}

func (value *Value) SetTag(tag reflect.StructTag) {
	value.tag = tag
}

func (value *Value) String() string {
	// 防御性变成， 注入 value 是 nil 会 panic
	if value == nil {
		return ""
	}

	pTag, ok := value.tag.Lookup("prefix")
	if !ok {
		return "没有 tag"
	}

	return pTag
}
