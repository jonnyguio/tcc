package coroutines

type Coroutine struct {
	base          func(...interface{}) []interface{}
	yield         chan interface{}
	resume        chan interface{}
	dead, started bool
}

func Create(base func(*Coroutine, ...interface{}) []interface{}) *Coroutine {
	coro := &Coroutine{

		yield:   make(chan interface{}),
		resume:  make(chan interface{}),
		dead:    false,
		started: false,
	}
	coro.base = func(args ...interface{}) []interface{} {
		rets := base(coro, args...)
		coro.dead = true
		return coro.Yield(rets...)
	}
	return coro
}

func (c *Coroutine) Resume(args ...interface{}) []interface{} {
	if c.dead {
		panic("Cannot resume a dead coroutine")
	}
	if !c.started {
		c.started = true
		go c.base(args...)
	} else {
		for _, value := range args {
			c.resume <- value
		}
		close(c.resume)
	}
	list := []interface{}{}
	for yieldedValue := range c.yield {
		list = append(list, yieldedValue)
	}
	c.yield = make(chan interface{})
	return list
}

func (c *Coroutine) Yield(rets ...interface{}) (list []interface{}) {
	for _, value := range rets {
		c.yield <- value
	}
	close(c.yield)
	if len(rets) > 0 {
		list = []interface{}{}
		for newArg := range c.resume {
			list = append(list, newArg)
		}
	}
	c.resume = make(chan interface{})
	return list
}
