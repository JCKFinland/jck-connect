package purchase

import "context"

// fakeTxManager executes the callback immediately.
//
// It mimics database.WithTransaction()
// without touching a real database.
type fakeTxManager struct{}

func (f *fakeTxManager) WithTransaction(
	ctx context.Context,
	fn func(any) error,
) error {
	return fn(nil)
}