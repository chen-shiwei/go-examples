package redis

import (
	"fmt"
	"time"
)

func StringSet(key string, val interface{}, expiration time.Duration) error {
	status := cli.Set(key, val, expiration)
	if err := status.Err(); err != nil {
		err = fmt.Errorf("set err:%s \n", status.Err().Error())
		return err
	}
	fmt.Printf("status: name:%s value:%s", status.Name(), status.Val())
	return nil
}

func StringGet(key string) (val string, err error) {
	val, err = cli.Get(key).Result()
	if err != nil {
		err = fmt.Errorf("get err:%s", err.Error())
		return
	}
	return
}

func StringDel(key string) (rows int64, err error) {
	rows, err = cli.Del(key).Result()
	if err != nil {
		err = fmt.Errorf("del err:%s", err.Error())
		return
	}
	return
}
