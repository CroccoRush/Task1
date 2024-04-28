package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	rdb *redis.Client
	ctx context.Context
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // use default Addr
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
	log.Println("redis:6379")
	ctx = context.TODO()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
	req, err := json.Marshal("pong")
	if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(req)
	w.WriteHeader(http.StatusOK)
}

func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	key := mux.Vars(r)["key"]
	value := mux.Vars(r)["value"]

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeRes(w, "Success")
	w.WriteHeader(http.StatusOK)
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	key := mux.Vars(r)["key"]

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
		writeRes(w, "key does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("key", val)
	writeRes(w, val)
	w.WriteHeader(http.StatusOK)
}

func writeRes(w http.ResponseWriter, val string) {
	req, err := json.Marshal(val)
	if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(req)
}
