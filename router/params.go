package router

type routeParams struct {
	index int
	keys  [8]string
	vals  [8]string
}

func NewRouteParams() *routeParams {
	return &routeParams{
		index: 0,
		keys:  [8]string{},
		vals:  [8]string{},
	}
}

func (rp *routeParams) Get(key string) string {
	for i := 0; i < len(rp.keys); i++ {
		if rp.keys[i] == key {
			return rp.vals[i]
		}
	}
	return ""
}

func (rp *routeParams) Add(key, val string) {
	rp.keys[rp.index] = key
	rp.vals[rp.index] = val
	rp.index++
}
