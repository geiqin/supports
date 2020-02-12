package paginate

import "github.com/geiqin/supports/helper"

type Paginator struct {
	Total int32
	PageSize int32
	Paged int32
}

func New(paged int32,pageSize ...int32) *Paginator {
	var psize int32
	if pageSize !=nil {
		if pageSize[0] >0{
			psize =pageSize[0]
		}
	}
	if paged <1 {
		paged =1
	}

	if psize < 1{
		psize =20
	}

	entity :=&Paginator{
		PageSize:psize,
		Paged:paged,
	}
	return entity
}


func (a *Paginator) Offset() int32 {
	offset :=(a.Paged-1) * a.PageSize
	return offset
}

func (a *Paginator) Limit() int32 {
	return a.PageSize
}

func (a *Paginator) ToPager(pbPager interface{}) *interface{} {
	helper.StructCopy(pbPager,a)
	return &pbPager
}

func Top(top int32) int32  {
	if top > 0{
		return top
	}
	return 20
}