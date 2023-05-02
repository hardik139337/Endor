package objectdb

import (
	object "Endor/Object"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type redisObjectDB struct {
	client *redis.Client
}

func NewRedisObjectDB(Addr string) *redisObjectDB {
	if Addr == "" {
		Addr = "localhost:6379"

	}

	client := redis.NewClient(&redis.Options{Addr: Addr})

	return &redisObjectDB{client}
}

func (db *redisObjectDB) Store(ctx context.Context, object object.Object) error {
	if object.GetID() != "" {
		return errors.New("object already has an ID")
	}

	object.SetID(uuid.New().String())
	reflect.ValueOf(object).Elem().FieldByName("Type").SetString(object.GetKind())

	data, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = db.client.Set(ctx, object.GetID(), data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
func (db *redisObjectDB) GetObjectByID(ctx context.Context, id string) (object.Object, error) {
	data, err := db.client.Get(ctx, id).Result()

	if err != nil {
		return nil, err
	}

	var obj object.Object
	if strings.Contains(data, "object.Person") {
		obj = &object.Person{}
	}
	if strings.Contains(data, "object.Animal") {
		obj = &object.Animal{}
	}
	err = json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
func (r *redisObjectDB) GetObjectByName(ctx context.Context, name string) (obj object.Object, err error) {
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return
	}
	for _, key := range keys {
		obj, err := r.GetObjectByID(ctx, key)
		if err != nil {
			panic(err)
		}
		if obj.GetName() == name {
			return obj, nil
		}
	}
	return
}

func (r *redisObjectDB) ListObjects(ctx context.Context, kind string) (results []object.Object, err error) {
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return
	}

	for _, key := range keys {
		o, err := r.GetObjectByID(ctx, key)
		if err != nil {
			panic(err)
		}

		if o.GetKind() == kind {
			results = append(results, o)
		}
	}
	return
}

func (r *redisObjectDB) DeleteObject(ctx context.Context, id string) error {
	err := r.client.Del(ctx, id).Err()
	if err != nil {
		return err
	}
	return nil
}
