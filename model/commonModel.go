package model

type Info struct {
	Number  int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number"`
	Passed  bool   `protobuf:"varint,2,opt,name=passed,proto3" json:"passed"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

type Error struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

type DataResponse struct {
	Entity interface{}    `protobuf:"bytes,2,rep,name=entity,proto3" json:"entity,omitempty"`
	Items  []*interface{} `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *Error         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *Info          `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}
