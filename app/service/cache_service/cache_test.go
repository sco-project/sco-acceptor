/**
    package: sco_tracers
    filename: cache_service
    author: diogo@gmail.com
    time: 2021/9/14 16:01
**/
package cache_service

import "testing"

func TestCacheTagService_Set(t *testing.T) {

	cache := New()

	list := "diogoxiang@gmail.com"

	cache.Set("name", list, 0, "user")

	//t.Log(list)

	clist := cache.Get("name")

	t.Log(clist)

	t.Log(cache.Data())

	cache.RemoveByTag("user")
	t.Log(cache.Data())
}


func TestCacheTagService_SetTag(t *testing.T) {

	cache := New()

	ukey := "name"
	list := "diogoxiang@gmail.com"

	cache.Set(ukey,list,0,"mail")


	t.Log(cache.Contains(ukey))
	//cache.Contains(ukey)
	//t.Log(cache.Contains("nameDio"))
	//cache.GetOrSetFunc()
}

func TestCacheTagService_GetOrSetFunc(t *testing.T) {

	cache := New()


	ukey := "name"
	//list := "diogoxiang@gmail.com"
	//cache.Set(ukey,list,0,"mail")

	v := cache.GetOrSetFunc(ukey,func() (interface{}, error) {

		rstr := "diogoxiang"

		return rstr, nil
	},0,"mail")

	t.Log(v)

	uValue := cache.Get(ukey)

	t.Log(uValue)


}