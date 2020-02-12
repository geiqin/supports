package memory
/*
import (
	"container/list"
	"sync"
)


import (
	"container/list"
	"github.com/geiqin/supports/sessionBak"
	"sync"
	"time"
)

var pder = &FromMemory{list: list.New()}

func MemoryInit() {
	pder.sessions = make(map[string]*list.Element, 0)
	//注册  memory 调用的时候一定有一致
	session.Register("memory", pder)
}

//session来自内存 实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

func (frommemory *FromMemory) SessionInit(sid string) (session.Session, error) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &session.SessionStore{sid: int(sid), LastAccessedTime: time.Now(), value: v}
	element := frommemory.list.PushBack(newsess)
	frommemory.sessions[sid] = element
	return newsess, nil
}

func (frommemory *FromMemory) SessionRead(sid string) (session.Session, error) {
	if element, ok := frommemory.sessions[sid]; ok {
		return element.Value.(*session.SessionStore), nil
	} else {
		sess, err := frommemory.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (frommemory *FromMemory) SessionDestroy(sid string) error {
	if element, ok := frommemory.sessions[sid]; ok {
		delete(frommemory.sessions, sid)
		frommemory.list.Remove(element)
		return nil
	}
	return nil
}

func (frommemory *FromMemory) SessionGC(maxLifeTime int64) {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	for {
		element := frommemory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*session.SessionStore).LastAccessedTime.Unix() + maxLifeTime) <
			time.Now().Unix() {
			frommemory.list.Remove(element)
			delete(frommemory.sessions, element.Value.(*session.SessionStore).sid)
		} else {
			break
		}
	}
}
func (frommemory *FromMemory) SessionUpdate(sid string) error {
	frommemory.lock.Lock()
	defer frommemory.lock.Unlock()
	if element, ok := frommemory.sessions[sid]; ok {
		element.Value.(*session.SessionStore).LastAccessedTime = time.Now()
		frommemory.list.MoveToFront(element)
		return nil
	}
	return nil
}

 */