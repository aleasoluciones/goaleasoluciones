package safemap

type Key interface{}
type Value interface{}
type store map[Key]Value
type SafeMap chan func(store)

type findResult struct {
	value Value
	found bool
}

func NewSafeMap() SafeMap {
	sm := make(SafeMap)
	go sm.run()
	return sm
}

func (sm SafeMap) run() {
	st := make(store)
	for command := range sm {
		command(st)
	}
}

func (sm SafeMap) Insert(key Key, value Value) {
	sm <- func(st store) {
		st[key] = value
	}
}

func (sm SafeMap) Delete(key Key) {
	sm <- func(st store) {
		delete(st, key)
	}
}

func (sm SafeMap) Find(key Key) (value Value, found bool) {
	reply := make(chan findResult)
	sm <- func(st store) {
		value, found := st[key]
		reply <- findResult{value, found}
	}
	result := <-reply
	return result.value, result.found
}

func (sm SafeMap) Len() int {
	reply := make(chan int)
	sm <- func(st store) {
		reply <- len(st)
	}
	return <-reply
}

func (sm SafeMap) Update(key Key, updater func(Value, bool) Value) {
	sm <- func(st store) {
		value, found := st[key]
		st[key] = updater(value, found)
	}
}

func (sm SafeMap) Close() {
	sm <- func(st store) {
		close(sm)
	}
}

func (sm SafeMap) Keys() []Key {
	reply := make(chan []Key)
	sm <- func(st store) {
		keys := make([]Key, 0)
		for key := range st {
			keys = append(keys, key)
		}
		reply <- keys
	}
	return <-reply
}
