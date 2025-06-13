package application

import (
	"sync"
)

type Context struct {
	Tasks sync.WaitGroup
}