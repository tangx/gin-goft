package goft

type ClassController interface {
	// Build(goft *GoftGroup)
	Method() string
	Path() string
	Handler() (interface{}, error)
}

type HandlerFunc = func() (interface{}, error)
