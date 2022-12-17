package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Mutex")

	/*
		In this example, we have a map of pointers to User objects, which are protected by a mutex. A mutex (short for "mutual exclusion") is a type of synchronization object that controls access to a shared resource. It is used to ensure that only one goroutine (a lightweight thread of execution) can access the resource at a time, preventing race conditions and other synchronization issues. A Mutex has two methods: Lock, which acquires the mutex, and Unlock, which releases the mutex. When a goroutine calls Lock, it will block until it can acquire the mutex. Once it has acquired the mutex, it can access the shared resource. When it is finished with the resource, it must call Unlock to release the mutex. This allows other goroutines to acquire the mutex and access the shared resource.

		In the following code, we have two goroutines that modify the map: one adds a new User object to the map, and the other removes the same object from the map. We then run the garbage collector to reclaim the memory used by the User object. Running the garbage collector can be useful in certain cases, such as when a program is using a lot of memory and needs to free up some resources. However, it is generally not recommended to run the garbage collector manually, as it can impact the program. In most cases, the Go runtime will automatically run the garbage collector as needed. Finally, we use a read lock to access the map and attempt to retrieve the User object. This will result in the output "user not found in map" every time the program is run. However, if you deactivate the locking, this will not be the case anymore. An educated reader with knowledge of garbage collection may point out that this is not the full story, as one can see when the manual triggering of garbage collection is deactivated. However, the code demonstrates in principle how to work with mutexes.
	*/

	type User struct {
		Name   string
		Email  string
		Active bool
	}

	// Declare a map to store pointers to User objects
	users := make(map[string]*User)

	// Create a mutual exclusion to protect the map
	var mu sync.RWMutex

	// Create a goroutine to add users to the map
	go func() {
		mu.Lock()
		defer mu.Unlock()

		// Create a new User object and add it to the map
		u := &User{
			Name:   "Michael Schmidt",
			Email:  "michael@marslabs.mars",
			Active: true,
		}
		users["michael"] = u

		// Set the User object to nil, which will release the memory used by the object
		u = nil
	}()

	// Create a goroutine to remove users from the map
	go func() {
		mu.Lock()
		defer mu.Unlock()

		// Remove the user from the map
		delete(users, "michael")
	}()

	// Run the garbage collector manually to reclaim the memory - (! generally not recommended)
	runtime.GC()

	// Use a read lock to access the map
	mu.RLock()
	defer mu.RUnlock()

	// Attempt to access the user in the map
	u, ok := users["michael"]
	if !ok {
		fmt.Println("User not found in map")
	} else {
		fmt.Println(u.Name)
	}
}
