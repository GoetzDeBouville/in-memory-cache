# memory-cache

Library for handling cache in single thread.

Contains functions for get, set and remove elements from cache.

## Usage

1. Plug in library to your project

```shell
go get github.com/zhashkevych/scheduler
```

2. Create new cache entity in your function using cache package:

```go
    newCach := cache.New()
```

3. Use functions Set(), Get() and Delete() for handling cache.

```go
	newCach.Set("userId1", 12)
	newCach.Set("userId2", "413")
	newCach.Set("userId3", "413")

	fmt.Println(newCach.Get("userId1"))

	newCach.Delete("userId2")
	fmt.Println(newCach)
```

Output:

```shell
12 true
&{map[userId1:12 userId3:413]}
```
