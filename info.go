package discovery

import (
	"time"
)

// ServiceInfo provides information about single service
type ServiceInfo struct {
	Name      string
	Address   string
	UpdatedAt time.Time
	GoodUntil time.Time
	updatedBy string
	next      *ServiceInfo
	prev      *ServiceInfo
}

func (left *ServiceInfo) equals(right *ServiceInfo) bool {
	return left.Name == right.Name &&
		left.Address == right.Address
}

type infoList struct {
	first *ServiceInfo
	last  *ServiceInfo
	size  int
}

func (l *infoList) find(item *ServiceInfo) *ServiceInfo {
	current := l.first
	for current != nil {
		if current.equals(item) {
			return current
		}
		current = current.next
	}
	return nil
}

func (l *infoList) insert(item *ServiceInfo) {
	item.prev = l.last
	item.next = nil
	if l.last != nil {
		l.last.next = item
	}
	l.last = item
	if l.first == nil {
		l.first = item
	}
	l.size++
}

func (l *infoList) remove(item *ServiceInfo) {
	if item.prev != nil {
		item.prev.next = item.next
	} else {
		l.first = item.prev
	}
	if item.next != nil {
		item.next.prev = item.prev
	} else {
		l.last = item
	}
	l.size--
}

func contains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func (l *infoList) clear() {
	item := l.first
	for item != nil {
		next := item.next
		item.next = nil
		item.prev = nil
		item = next
	}
	l.first = nil
	l.last = nil
}
