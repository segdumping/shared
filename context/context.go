package context

import "context"

type Context struct {
	context.Context
}

type ctxKey struct {}

func NewContext(ctx context.Context, values map[string]interface{}) *Context {
	return &Context{
		context.WithValue(ctx, ctxKey{}, values),
	}
}

func FromContext(ctx context.Context) (map[string]interface{}, bool) {
	val, ok := ctx.Value(ctxKey{}).(map[string]interface{})
	return val, ok
}
