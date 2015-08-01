package resourcelimit

import ()

type SimpleLimit struct {
	resourceNum int
	limitPool   chan int
}

func NewLimit(ResourceNum int) *SimpleLimit {
	return &SimpleLimit{
		resourceNum: ResourceNum,
		limitPool:   make(chan int, ResourceNum),
	}
}

func (r *SimpleLimit) UseResource() {
	r.limitPool <- 1
	return
}

func (r *SimpleLimit) FreeResource() {
	<-r.limitPool
	return
}

func (r *SimpleLimit) GetResourceNum() int {
	return r.resourceNum
}
