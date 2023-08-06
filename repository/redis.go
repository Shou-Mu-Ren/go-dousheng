package repository

import (
	"fmt"
	"sync"
	"douyin/util"
	"github.com/gomodule/redigo/redis"
)

type RedisDao struct {
}

var redisDao *RedisDao
var redisOnce sync.Once

func NewRedisDaoInstance() *RedisDao {
	redisOnce.Do(
		func() {
			redisDao = &RedisDao{}
		})
	return redisDao
}

func (*RedisDao) AddFavorite(userId int64,videoId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("sadd",fmt.Sprintf("%s_%d", "Favorite", userId),videoId)
	if err != nil {
		util.Logger.Error("sadd favorite failed, err:" + err.Error())
    }
	return err
}

func (*RedisDao) RemFavorite(userId int64,videoId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("srem",fmt.Sprintf("%s_%d", "Favorite", userId), videoId)
	if err != nil {
		util.Logger.Error("srem favorite failed, err:" + err.Error()) 
    }
	return err
}

func (*RedisDao) IsFavorite(userId int64,videoId int64) (bool, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Bool(rdb.Do("sismember", fmt.Sprintf("%s_%d", "Favorite", userId), videoId)) 
    if err != nil {
		util.Logger.Error("is favorite err:" + err.Error()) 
        return false, err
    }
	return res, nil
}

func (*RedisDao) GetFavoriteVideoIds(userId int64) ([]string, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Strings(rdb.Do("smembers", fmt.Sprintf("%s_%d", "Favorite", userId))) 
    if err != nil {
		util.Logger.Error("get favorite ids err:" + err.Error()) 
        return nil, err
    }
	return res, nil
}

func (*RedisDao) AddFollow(userId int64,toUserId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("sadd",fmt.Sprintf("%s_%d", "Follow", userId),toUserId)
	if err != nil {
		util.Logger.Error("sadd follow failed, err:" + err.Error())
    }
	return err
}

func (*RedisDao) RemFollow(userId int64,toUserId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("srem",fmt.Sprintf("%s_%d", "Follow", userId), toUserId)
	if err != nil {
		util.Logger.Error("srem follow failed, err:" + err.Error()) 
    }
	return err
}

func (*RedisDao) IsFollow(userId int64, toUserId int64) (bool, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Bool(rdb.Do("sismember", fmt.Sprintf("%s_%d", "Follow", userId), toUserId)) 
    if err != nil {
		util.Logger.Error("is follow err:" + err.Error()) 
        return false, err
    }
	return res, nil
}

func (*RedisDao) GetFollowUserIds(userId int64) ([]string, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Strings(rdb.Do("smembers", fmt.Sprintf("%s_%d", "Follow", userId))) 
    if err != nil {
		util.Logger.Error("get follow ids err:" + err.Error()) 
        return nil, err
    }
	return res, nil
}

func (*RedisDao) AddFollower(userId int64, toUserId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("sadd",fmt.Sprintf("%s_%d", "Follower", toUserId), userId)
	if err != nil {
		util.Logger.Error("sadd follower failed, err:" + err.Error())
    }
	return err
}

func (*RedisDao) RemFollower(userId int64, toUserId int64) error{
	rdb := pool.Get()
	defer rdb.Close()
	_, err := rdb.Do("srem",fmt.Sprintf("%s_%d", "Follower", toUserId), userId)
	if err != nil {
		util.Logger.Error("srem follower failed, err:" + err.Error()) 
    }
	return err
}



func (*RedisDao) GetFollowerUserIds(toUserId int64) ([]string, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Strings(rdb.Do("smembers", fmt.Sprintf("%s_%d", "Follower", toUserId))) 
    if err != nil {
		util.Logger.Error("get follower ids err:" + err.Error()) 
        return nil, err
    }
	return res, nil
}

func (*RedisDao) GetFriendUserIds(userId int64) ([]string, error){
	rdb := pool.Get()
	defer rdb.Close()
	res, err := redis.Strings(rdb.Do("sinter", fmt.Sprintf("%s_%d", "Follower", userId), fmt.Sprintf("%s_%d", "Follow", userId))) 
    if err != nil {
		util.Logger.Error("get follower ids err:" + err.Error()) 
        return nil, err
    }
	return res, nil
}