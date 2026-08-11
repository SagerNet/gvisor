package stack

// PostDispatchLinkEndpoint is implemented by link endpoints that can invoke a
// callback on their dispatch goroutine after each dispatch, outside stack
// locks.
type PostDispatchLinkEndpoint interface {
	LinkEndpoint

	// SetPostDispatch installs the callback and reports whether it will be
	// invoked. Wrappers forward to the wrapped endpoint and report its
	// result. Must be called before Attach.
	SetPostDispatch(postDispatch func()) bool
}
