### Problem
We have N philosophers. Only M of N philosophers can have dinner at the same time.

### Solution

##### Only one philosopher can have dinner at the same time
We can use mutex to allow the work of only one thread at the same time. 

##### More than one philosopher can have dinner at the same time
To allow more threads we can use semaphore or buffered channel (as alternative implementation of semaphore logic).

### Demo

To run the demonstration please do:

```shell script
go run ./cmd
```