package main

import "fmt"
import "github.com/go-redis/redis"

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

const devFuncKey = "DeviceFunction"
const startAfresh = false

func main() {

	rclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:32525",
		Password: "",
		DB:       0,
	})

	if startAfresh == true {
		resetCache(rclient)
	}

	// Check to see if the list of device functions is already cached. If not,
	// add the values of the device functions to the cache key:
	if rclient.LLen(devFuncKey).Val() == 0 {
		fmt.Println("Populating device function cache item:")
		ff := []string{"ml", "se", "db", "px", "ad", "fw"}
		for f := range ff {
			fmt.Printf("- Caching device function tag %s\n", ff[f])
			rclient.LPush(devFuncKey, ff[f])
		}
	}
	fmt.Printf("Length of deviceFunction list: %d\n", rclient.LLen(devFuncKey).Val())
}
