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
/app # time ./game-of-life-sequence --display_world=false --max_rows=5000 --max_cols=5000
real    0m 1.08s
user    0m 0.97s
sys     0m 0.11s
```

### View World

```sh
/app # ./game-of-life-sequence --display_world=true --max_rows=5000 --max_cols=5000

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
/app # time ./game-of-life-concurrency-fast --display_world=false --max_rows=5000 --max_cols=5000
real    0m 0.97s
user    0m 2.51s
sys     0m 0.43s
```

### View World

```sh
/app # ./game-of-life-concurrency-fast --display_world=true --max_rows=5000 --max_cols=5000

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