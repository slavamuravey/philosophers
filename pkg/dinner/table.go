package dinner

import (
  "context"
  "github.com/fatih/color"
  "golang.org/x/sync/semaphore"
  "sync"
  "time"
)

// Table is thread safe struct holds philosophers count
type Table struct {
  philosophersCnt int
  sync.Mutex
}

// addPhilosopher puts the philosopher at the table
func (t *Table) addPhilosopher(name string) {
  t.Lock()
  defer t.Unlock()
  t.philosophersCnt++
  color.New(color.FgBlue).Printf("Philosopher %s starts to eat, total %d\n", name, t.philosophersCnt)
}

// removePhilosopher drops the philosopher from the table
func (t *Table) removePhilosopher(name string) {
  t.Lock()
  defer t.Unlock()
  t.philosophersCnt--
  color.New(color.FgRed).Printf("Philosopher %s stops to eat, total %d\n", name, t.philosophersCnt)
}

// DinnerMtx runs dinner process for philosopher using mutex
func (t *Table) DinnerMtx(name string, wg *sync.WaitGroup, mtx *sync.Mutex) {
  mtx.Lock()
  defer mtx.Unlock()
  defer wg.Done()
  t.addPhilosopher(name)
  time.Sleep(time.Millisecond * 500)
  t.removePhilosopher(name)
}

// DinnerSmf runs dinner process for philosopher using semaphore
func (t *Table) DinnerSmf(ctx context.Context, name string, wg *sync.WaitGroup, smf *semaphore.Weighted) {
  err := smf.Acquire(ctx, 1)
  if err != nil {
    panic(err.Error())
  }
  defer smf.Release(1)
  defer wg.Done()
  t.addPhilosopher(name)
  time.Sleep(time.Millisecond * 500)
  t.removePhilosopher(name)
}

// DinnerSmfCh runs dinner process for philosopher using buffered channel
func (t *Table) DinnerSmfCh(name string, wg *sync.WaitGroup, smfCh chan struct{}) {
  smfCh <- struct{}{}
  defer func() {
    <-smfCh
  }()
  defer wg.Done()
  t.addPhilosopher(name)
  time.Sleep(time.Millisecond * 500)
  t.removePhilosopher(name)
}
