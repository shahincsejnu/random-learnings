# Maps in Go

# Basics

- `make` : allocate and initialize the memory --> non zeroed storage
- `new` : only allocates the memory --> no initialization of memory
- `var mp map[string]int` or `mp := map[string]int`: 
    - here only memory allocated but not initialized
    - so we cannot do like :  `mp["oka"] = 100`, it will show an error (assignment to entry in nil map)
    - If we declare our map like this way then we need to do like:
        - ```
          mp = map[string]int{
              "oka" : 100,
          }
          ```
- `mp := make(map[string]int)`:
    - here, `mp["oka"] = 100` will not show error, as the map memory allocated and also initialized because of make

- `mp["new one"] = 200` : to add new entry
- `mp["oka"] = 300` : to update existing entry value
- `delete(mp, "oka")` : to delete an existing entry
- `val, ok := mp["oka"]` : to get the value (if the key "oka" exists then `val` will be the value of that key and `ok` will be true, else `val` will be default zero value and `ok` will be false)
- ```
  for key, value := range mp {
      // key is the key of the map named mp
      // value will be the value of the key
  }
  ```

# [Concurrency](https://blog.golang.org/maps#TOC_6.)
- Maps are not safe for concurrent use
- It's not defined what happens when you read and write to them simultaneously. 
- If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism
- One common way to protect maps is with [sync.RWMutex]()


## Note
- if we range/traverse in the map, every time we can get different order of key-value pair, cause map in golang does not maintain the order


# Resources
- [x] [Golang Tutorial #15 - Maps](https://www.youtube.com/watch?v=yJE2RC37BF4)
- [x] [Maps are super easy in golang | golang series](https://www.youtube.com/watch?v=p4LS3UdgJA4)
- [x] [Golang Official Tutorial](https://blog.golang.org/maps)