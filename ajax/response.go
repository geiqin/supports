package ajax

type Info struct {
	Number  int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Passed  bool   `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type Error struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

type Pager struct {
	Paged     int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	Total     int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	PageCount int32 `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count"`
	PageSize  int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	PrevPage  int32 `protobuf:"varint,5,opt,name=prev_page,json=prevPage,proto3" json:"prev_page"`
	LastPage  int32 `protobuf:"varint,6,opt,name=last_page,json=lastPage,proto3" json:"last_page"`
}

//
type DataResponse struct {
	Entity *interface{}   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *Pager         `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*interface{} `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func NewError(code int32, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
