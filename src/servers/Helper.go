package servers

import "strconv"

// 分页大集合
type Paging struct {
	Total   int           // 总计页数
	Current int           // 当前页码
	Size    int           // 页码尺寸
	PageNum int           // 一共多少页
	Data    []interface{} // 数据
	Ext     interface{}   // 预留信息
}

// 链式调用返回自身
func (pa *Paging) SetEXT(ext interface{}) *Paging {
	pa.Ext = ext
	return pa
}

//
type Helper struct{}

func NewHelper() *Helper {
	return &Helper{}
}

// 字符串int
func (*Helper) StrTOint(str string, def int) int {
	re, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return re
}

// 分页实现
func (*Helper) PageResource() *Paging {
	return nil
}
