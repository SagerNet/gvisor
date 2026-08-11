package stack

// PostDispatchLinkEndpoint is implemented by link endpoints that can invoke a
// callback on their dispatch goroutine after each dispatch, outside stack
// locks.
type PostDispatchLinkEndpoint interface {
	LinkEndpoint

	// SetPostDispatch installs the callback. Must be called before Attach.
	SetPostDispatch(postDispatch func())
}
