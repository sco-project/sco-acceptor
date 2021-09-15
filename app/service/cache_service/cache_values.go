/**
    package: sco_tracers
    filename: cache_service
    author: diogo@gmail.com
    time: 2021/9/14 15:56
**/
package cache_service

//缓存前缀KEY
const (
	AdminAuthMenu = iota
	AdminAuthRole
	AdminCmsMenu
	AdminConfigDict
	AdminBlogClassification
)

//缓存TAG标签
const (
	AdminAuthTag = iota
	AdminCmsTag
	AdminSysConfigTag
	AdminModelTag
	AdminBlogTag
)
