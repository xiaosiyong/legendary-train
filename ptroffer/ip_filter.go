package main

import "sync"

type ip struct {
	IP        string
	Timestamp int64
	Count     int64
}

func IpFilter(ip string, timestamp int64) bool {
	m := sync.Map{}
	v, ok := m.Load(ip)
	if !ok {
		return true
	}

}
