package main

import "fmt"
import "github.com/go-redis/redis"

// The init function.
func init() {
	fmt.Print("I am very pleased to be initializing, for your delight - deep joy!\n")
}

// Function to reset the cache when we start afresh.
func resetCache(rclient *redis.Client) {
	if rclient.LLen(devFuncKey).Val() == 0 {
		// Already empty cache so no action required
		return
	}

	if _, err := rclient.Del(devFuncKey).Result(); err != nil {
		panic(err)
	}
}

// The name of the cache key.
const devFuncKey = "DeviceFunction"

// Whether the cache should be preserved or zeroed for the next use.
const cachePreserve = false

func main() {
	// Connect to Redis cache layer:
	rclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:32525",
		Password: "",
		DB:       0,
	})

	// Check to see if the list of device functions is already cached. If not,
	// add the values of the device functions to the cache key:
	fmt.Println("Populating device function cache item:")
	ff := []string{"ml", "se", "db", "px", "ad", "fw", "fs", "as", "vo", "mn", "bk", "ns", "cf", "mq", "wb"}
	for f := range ff {
		fmt.Printf("- Caching device function tag %s\n", ff[f])
		rclient.RPush(devFuncKey, ff[f])
	}

	fmt.Printf("Length of deviceFunction list: %d\n", rclient.LLen(devFuncKey).Val())

	// Clean out the cache at the end:
	if cachePreserve == false {
		resetCache(rclient)
	}
	// Somewhere here we have to handle the payload which will have been
	// passed in a query parameter: the payload will probably contain some
	// information which will come from e.g, metadata for provisioning, which
	// can be use to determine which device function is most appropriate for
	// the host we want to create

}
