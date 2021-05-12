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

type Helper struct{}

func NewHelper() *Helper {
	return &Helper{}
}

// 字符串转换int
func (*Helper) StrTOint(str string, def int) int {
	re, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return re
}

// 分页实现
func (*Helper) PageResource(cu, size,tot int,list []interface{}) *Paging {
	// 设置默认的分页为8个
	if size == 0 || size > tot{
		size = 8
	}
	// 设置默认的当前所在页码
	if cu <=0 {
		cu = 1
	}
	// 构造计算的变量
	pagemessage := &Paging{Total: tot,Size: size}
	// 计算总的页面数量
	pageNum := 1
	if pagemessage.Total > size {
		pageNum = pagemessage.Total/size
		if pagemessage.Total % size != 0 {
			pageNum ++
		}
	}
	// 设置当前页面尺寸
	if cu > pageNum {
		cu = 1
	}
	// 赋值当前页面
	pagemessage.Current = cu
	// 构造新的切片
	set := make([]interface{},0)
	// 处理当前的页面位置
	if cu*size > pagemessage.Total {
		set = append(set,list[(cu-1)*size:]...)
	}else {
		set = append(set,list[(cu-1)*size:(cu-1)*size+size]...)
	}
	// 赋予最新的数值
	pagemessage.Data=set
	pagemessage.PageNum=pageNum
	return pagemessage
}
