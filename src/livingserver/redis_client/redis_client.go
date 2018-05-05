package redis_client

import (
	"errors"
	"fmt"
	"log"
	"redis"
	"strconv"

	"github.com/astaxie/beego/logs"
	// "strconv"
	// "github.com/astaxie/beego/logs"
)

var client redis.Client

const (
	LIKE_CNT_PREFIX    = "like_cnt"
	COMMENT_CNT_PREFIX = "comment_cnt"
)

var (
	ErrRedisOp = errors.New("<Redis> error operation")
)

func init() {
	spec := redis.DefaultSpec().Db(0).Host("127.0.0.1").Password("123456")
	c, err := redis.NewSynchClientWithSpec(spec)
	if err != nil {
		log.Fatal("Init redis client failed")
		// logs.Critical("Init redis client failed")
		// os.Exit(-1)
	}
	err = c.Ping()
	if err != nil {
		log.Fatal("Init redis client failed")
		// logs.Critical("Ping redis server failed")
		// os.Exit(-1)
	}
	client = c
	log.Println("Init redis client successful")
	// logs.Info("Init redis client successful")

}

func CreateLikeCnt(id int) bool {
	key := fmt.Sprintf("%s_%d", LIKE_CNT_PREFIX, id)
	err := client.Set(key, []byte("0"))
	if err != nil {
		logs.Warn("err on Incr, key: [%v], err: [%v]\n", key, err)
		return false
	}
	logs.Debug("like cnt key: ", key)
	return true
}

func IncrLikeCnt(id int) bool {
	key := fmt.Sprintf("%s_%d", LIKE_CNT_PREFIX, id)
	_, err := client.Incr(key)
	if err != nil {
		logs.Warn("err on Incr, key: [%v], err: [%v]\n", key, err)
		return false
	}
	return true
}

func GetLikeCnt(id int) int {
	key := fmt.Sprintf("%s_%d", LIKE_CNT_PREFIX, id)
	v, e := client.Get(key)
	if e != nil {
		logs.Warn("err on Incr, key: [%v]\n", key)
		return -1
	}

	cnt, err := strconv.Atoi(string(v))
	if err != nil {
		logs.Warn("convert value[%v] of key[%v] failed, err: [%v]\n", v, key, err)
		return -1
	}
	return cnt
}

func CreateCommentCnt(id int) bool {
	key := fmt.Sprintf("%s_%d", COMMENT_CNT_PREFIX, id)
	err := client.Set(key, []byte("0"))
	if err != nil {
		logs.Warn("err on Incr, key: [%v], err: [%v]\n", key, err)
		return false
	}
	return true
}

func IncrCommentCnt(id int) bool {
	key := fmt.Sprintf("%s_%d", COMMENT_CNT_PREFIX, id)
	_, err := client.Incr(key)
	if err != nil {
		logs.Warn("err on Incr, key: [%v], err: [%v]\n", key, err)
		return false
	}
	return true
}

func GetCommentCnt(id int) int {
	key := fmt.Sprintf("%s_%d", COMMENT_CNT_PREFIX, id)
	v, e := client.Get(key)
	if e != nil {
		logs.Warn("err on Incr, key: [%v]\n", key)
		return -1
	}

	cnt, err := strconv.Atoi(string(v))
	if err != nil {
		logs.Warn("convert value[%v] of key[%v] failed, err: [%v]\n", v, key, err)
		return -1
	}
	return cnt
}
