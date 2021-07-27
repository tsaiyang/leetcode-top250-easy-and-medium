package main

import "container/list"

type Entry struct {
    key   int
    value int
}

type LRUCache struct {
    capacity int
    size     int
    ll       *list.List
    cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        size:     0,
        ll:       list.New(),
        cache:    make(map[int]*list.Element),
    }
}

func (this *LRUCache) Get(key int) int {
    element, ok := this.cache[key]
    if !ok {
        return -1
    }
    this.ll.MoveToBack(element)
    entry := element.Value.(*Entry)
    return entry.value
}

func (this *LRUCache) Put(key int, value int) {
    element, ok := this.cache[key]
    if ok {
        this.ll.MoveToBack(element)
        entry := element.Value.(*Entry)
        entry.value = value
        return
    }
    newElement := this.ll.PushBack(&Entry{key: key, value: value})
    this.cache[key] = newElement
    this.size++
    if this.size > this.capacity {
        removeElement := this.ll.Front()
        this.ll.Remove(removeElement)
        entry := removeElement.Value.(*Entry)
        delete(this.cache, entry.key)
        this.size--
    }
}
