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

var packPool *PackPool

type PackPool struct {
	*redis.Pool
}

var RedigoConn = new(RedigoPack)

func NewConnectionByPool(p *PackPool) *RedigoPack {
	packPool = p

	return RedigoConn
}

// Note! release the conn
func GetConn() redis.Conn {
	if packPool == nil {
		return nil
	}
	return packPool.Get()
}

// Note! release the conn
func GetPubSubConn() *redis.PubSubConn {
	if packPool == nil {
		return nil
	}

	psc := new(redis.PubSubConn)
	psc.Conn = packPool.Get()

	return psc
}