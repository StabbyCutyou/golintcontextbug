package golintcontextbug

import "context"

// TwoContexts demonstrates a bug in golints rule about context placement in a function
func TwoContexts(ctx context.Context, anotherCtx context.Context) context.Context {
	return ctx
}

// TwoContextsAndAString demonstrates a bug in golints rule about context placement in a function
func TwoContextsAndAString(ctx context.Context, anotherCtx context.Context, s string) context.Context {
	return ctx
}
