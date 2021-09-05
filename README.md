# Conways Game of Life in Golang

## Run on Docker Container

```sh
Conways-Game-of-Life-in-Golang$ docker run -it -v `pwd`:/app/ golang:1.15.6-alpine3.12 /bin/sh
/go # cd ../app/
```

## Sequential

### Time to complete

```sh
/app # go build game-of-life-sequence.go 
/app # time ./game-of-life-sequence
real    1m 32.93s
user    1m 31.97s
sys     0m 0.94s
```

### View World

```sh
/app # go run game-of-life-sequence-display-world.go 

 ----- ----- ----- ----- -----
:  1  :  0  :  0  :  1  :  0  :
 ----- ----- ----- ----- -----
:  1  :  0  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  1  :  0  :  0  :  0  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  0  :  1  :
 ----- ----- ----- ----- -----
:  1  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
Next Generation:

 ----- ----- ----- ----- -----
:  0  :  1  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  1  :  0  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  1  :  1  :  0  :  1  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
```

## Concurrency

### Time to complete

```sh
/app # go build game-of-life-concurrency-fast.go
/app # time ./game-of-life-concurrency-fast
real    1m 5.70s
user    8m 10.29s
sys     0m 12.98s
```

### View World

```sh
/app # go run game-of-life-concurrency-fast-display-world.go

 ----- ----- ----- ----- -----
:  1  :  0  :  0  :  1  :  0  :
 ----- ----- ----- ----- -----
:  1  :  0  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  1  :  0  :  0  :  0  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  0  :  1  :
 ----- ----- ----- ----- -----
:  1  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
Next Generation:

 ----- ----- ----- ----- -----
:  0  :  1  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  1  :  0  :  1  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  1  :  1  :  0  :  1  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
:  0  :  0  :  0  :  1  :  1  :
 ----- ----- ----- ----- -----
```