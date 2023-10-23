package main

import "sync"

var wg sync.WaitGroup
var state = make(map[int64]map[string]interface{})
var goroutines = make(map[int]chan bool)
var goroutineId = 0