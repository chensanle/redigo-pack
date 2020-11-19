package redigo_pack

import (
	"github.com/gomodule/redigo/redis"
)

type RedigoPack struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}

var RedigoConn = new(RedigoPack)

func NewConnectionByPool(pool2 *redis.Pool) *RedigoPack {
	initPool(pool2)

	return new(RedigoPack)
}

var pool *redis.Pool

func initPool(old *redis.Pool) {
	pool = old
}
