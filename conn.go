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

var pool *redis.Pool



var RedigoConn = new(RedigoPack)

func NewConnectionByPool(p *redis.Pool) *RedigoPack {
	pool = p

	return RedigoConn
}

// Note! release the conn
func GetConn() redis.Conn {
	if pool == nil {
		return nil
	}
	return pool.Get()
}

// Note! release the conn
func GetPubSubConn() *redis.PubSubConn {
	if pool == nil {
		return nil
	}

	psc := new(redis.PubSubConn)
	psc.Conn = pool.Get()

	return psc
}