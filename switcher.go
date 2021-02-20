package switcher

import "context"

type IStoreEngine interface {
	IsEnable(context.Context, string, bool) bool
}

type Switcher struct {
	e IStoreEngine
}

func New(e IStoreEngine) {
	return &Switcher{e: e}
}

func (s *Switcher) IsEnable(ctx context.Context, key string) {
	return s.e.IsEnable(ctx, key)
}
