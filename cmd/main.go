package main

import (
  "context"
  "fmt"
  "github.com/slavamuravey/philosophers/pkg/dinner"
  "golang.org/x/sync/semaphore"
  "sync"
)

func main() {
  t := new(dinner.Table)
  wg := new(sync.WaitGroup)
  mtx := new(sync.Mutex)
  var maxPhilosophers int

  maxPhilosophers = 1
  fmt.Printf("%d philosophers can eat \n", maxPhilosophers)

  wg.Add(5)
  go t.DinnerMtx("Sokrat", wg, mtx)
  go t.DinnerMtx("Fales", wg, mtx)
  go t.DinnerMtx("Pifagor", wg, mtx)
  go t.DinnerMtx("Ploton", wg, mtx)
  go t.DinnerMtx("Arkhimed", wg, mtx)
  wg.Wait()

  maxPhilosophers = 2
  fmt.Printf("%d philosophers can eat \n", maxPhilosophers)

  smf := semaphore.NewWeighted(int64(maxPhilosophers))
  wg.Add(5)
  ctx := context.Background()
  go t.DinnerSmf(ctx, "Sokrat", wg, smf)
  go t.DinnerSmf(ctx, "Fales", wg, smf)
  go t.DinnerSmf(ctx, "Pifagor", wg, smf)
  go t.DinnerSmf(ctx, "Ploton", wg, smf)
  go t.DinnerSmf(ctx, "Arkhimed", wg, smf)
  wg.Wait()

  maxPhilosophers = 3
  fmt.Printf("%d philosophers can eat \n", maxPhilosophers)

  smfCh := make(chan struct{}, maxPhilosophers)
  wg.Add(5)
  go t.DinnerSmfCh("Sokrat", wg, smfCh)
  go t.DinnerSmfCh("Fales", wg, smfCh)
  go t.DinnerSmfCh("Pifagor", wg, smfCh)
  go t.DinnerSmfCh("Ploton", wg, smfCh)
  go t.DinnerSmfCh("Arkhimed", wg, smfCh)
  wg.Wait()
}
