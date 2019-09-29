package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

const (
	scanCount = 1000
)

func ReplyToMap(reply interface{}, values interface{}) error {
	array, err := redis.Values(reply, nil)
	if err != nil {
		return err
	}

	return ArrayToMap(array, values)
}

func ArrayToMap(array []interface{}, values interface{}) error {
	typ := reflect.TypeOf(values)
	if typ.Kind() != reflect.Map {
		return errors.New("param is not map")
	}

	val := reflect.ValueOf(values)
	keyType := typ.Key().Kind()
	valType := typ.Elem().Kind()

	length := len(array)
	for i := 0; i < length; i += 2 {
		k, err := toValue(keyType, array[i])
		if err != nil {
			return err
		}

		v, err := toValue(valType, array[i+1])
		if err != nil {
			return err
		}

		val.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
	}

	return nil
}

func toValue(kind reflect.Kind, value interface{}) (interface{}, error) {
	switch kind {
	case reflect.Int64:
		return redis.Int64(value, nil)
	case reflect.Int:
		return redis.Int(value, nil)
	case reflect.String:
		return redis.String(value, nil)
	case reflect.Uint64:
		return redis.Uint64(value, nil)
	case reflect.Bool:
		return redis.Bool(value, nil)
	case reflect.Float64:
		return redis.Float64(value, nil)
	}

	return nil, errors.New("not support type")
}

func ScanHash(key string) ([]interface{}, error) {
	var cursor = "0"
	values := make([]interface{}, 0, scanCount)

	for {
		reply, err := Do("HSCAN", key, cursor, "COUNT", scanCount)
		if err != nil {
			return nil, err
		}

		array, err := redis.Values(reply, nil)
		if err != nil {
			return nil, err
		}

		if len(array) != 2 {
			return nil, errors.New("redis hscan error")
		}

		cursor, err = redis.String(array[0], nil)
		if err != nil {
			return nil, err
		}

		values = append(values, array[1].([]interface{})...)

		if cursor == "0" {
			break
		}
	}

	return values, nil
}
