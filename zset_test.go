package redigo_pack

import (
	"github.com/gomodule/redigo/redis"
	"testing"
)

func Test_zSetRds_ZAdd(t *testing.T) {
	NewConnectionByPool(&redis.Pool{})

	err := RedigoConn.ZSet.ZAdd("key", map[interface{}]interface{}{3: 3, 1: 4, 5: 5}).error
	if err != nil {
		t.Error(err)
	}
}
