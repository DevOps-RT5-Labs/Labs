package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type KV struct {
	*redis.Client
	context.Context
}

func NewKV(hostname string) (*KV, error) {
	kv := &KV{
		Client: redis.NewClient(&redis.Options{
			Addr:     hostname + ":6379",
			Password: "", 
			DB:       0,  
		}),
		Context: context.Background(),
	}

	_, err := kv.Ping(kv.Context).Result()
	if err != nil {
		return nil, err
	}

	return kv, nil
}

func (kv *KV) Get(key string) (string, error) {
	val, err := kv.Client.Get(kv.Context, key).Result()
    if err != nil {
		return "", err
    }

	return val, err
}

func (kv *KV) Set(key string, value string) error {
	err := kv.Client.Set(kv.Context, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}