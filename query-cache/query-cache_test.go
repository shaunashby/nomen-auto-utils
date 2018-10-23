package main

import (
	"testing"

	"github.com/go-redis/redis"
)

func Test_resetCache(t *testing.T) {
	type args struct {
		rclient *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetCache(tt.args.rclient)
		})
	}
}
