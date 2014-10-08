package safemap

type SafeMap chan commandData

type commandData struct {
	action  int
	key     interface{}
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[interface{}]interface{}
	updater UpdateFunc
}

const (
	remove = iota
	end
	find
	insert
	length
	update
	keys
)

type findResult struct {
	value interface{}
	found bool
}

type UpdateFunc func(interface{}, bool) interface{}

func New() SafeMap {
	sm := make(SafeMap)
	go sm.run()
	return sm
}

func (sm SafeMap) run() {
	store := make(map[interface{}]interface{})
	for command := range sm {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			delete(store, command.key)
		case find:
			value, found := store[command.key]
			command.result <- findResult{value, found}
		case length:
			command.result <- len(store)
		case keys:
			keys := make([]interface{}, 0)
			for key, _ := range store {
				keys = append(keys, key)
			}
			command.result <- keys
		case update:
			value, found := store[command.key]
			store[command.key] = command.updater(value, found)
		case end:
			close(sm)
			command.data <- store
		}
	}
}

func (sm SafeMap) Insert(key interface{}, value interface{}) {
	sm <- commandData{action: insert, key: key, value: value}
}

func (sm SafeMap) Delete(key interface{}) {
	sm <- commandData{action: remove, key: key}
}

func (sm SafeMap) Find(key interface{}) (value interface{}, found bool) {
	reply := make(chan interface{})
	sm <- commandData{action: find, key: key, result: reply}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (sm SafeMap) Len() int {
	reply := make(chan interface{})
	sm <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (sm SafeMap) Update(key interface{}, updater UpdateFunc) {
	sm <- commandData{action: update, key: key, updater: updater}
}

func (sm SafeMap) Close() map[interface{}]interface{} {
	reply := make(chan map[interface{}]interface{})
	sm <- commandData{action: end, data: reply}
	return <-reply
}

func (sm SafeMap) Keys() []interface{} {
	reply := make(chan interface{})
	sm <- commandData{action: keys, result: reply}
	return (<-reply).([]interface{})
}
