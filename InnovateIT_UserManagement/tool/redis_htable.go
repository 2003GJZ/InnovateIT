package tool

import "InnovateIT_UserManagement/mylink"

type Redis_htable struct {
	Htabname   string            //哈希表名
	Redis_link *mylink.RedisLink //redis链接
}

// 查询缓存
func (htable *Redis_htable) Query_caching(key string) string {
	var value string
	htable.Redis_link.Client.HGet(htable.Redis_link.Ctx, htable.Htabname, key).Scan(&value)
	return value
}

// 插入缓存
func (htable *Redis_htable) Insert_caching(key string, value string) bool {
	htable.Redis_link.Client.HSet(htable.Redis_link.Ctx, htable.Htabname, key, value)
	return true
}

// 删除缓存
func (htable *Redis_htable) Delete_caching(key string) bool {
	htable.Redis_link.Client.HDel(htable.Redis_link.Ctx, htable.Htabname, key)
	return true
}
