// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: news/v1/news.proto

package newsv1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ImageUrl  string               `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	UserId    string               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	mi := &file_news_v1_news_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{0}
}

func (x *News) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *News) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *News) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *News) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *News) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *News) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *News) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetNewsFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetNewsFeedRequest) Reset() {
	*x = GetNewsFeedRequest{}
	mi := &file_news_v1_news_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNewsFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsFeedRequest) ProtoMessage() {}

func (x *GetNewsFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsFeedRequest.ProtoReflect.Descriptor instead.
func (*GetNewsFeedRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{1}
}

func (x *GetNewsFeedRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetNewsFeedRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetNewsFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News        []*News `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty"`
	TotalPages  int32   `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	CurrentPage int32   `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	TotalItems  int32   `protobuf:"varint,4,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
}

func (x *GetNewsFeedResponse) Reset() {
	*x = GetNewsFeedResponse{}
	mi := &file_news_v1_news_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNewsFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsFeedResponse) ProtoMessage() {}

func (x *GetNewsFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsFeedResponse.ProtoReflect.Descriptor instead.
func (*GetNewsFeedResponse) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{2}
}

func (x *GetNewsFeedResponse) GetNews() []*News {
	if x != nil {
		return x.News
	}
	return nil
}

func (x *GetNewsFeedResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *GetNewsFeedResponse) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *GetNewsFeedResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

type GetNewsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNewsByIdRequest) Reset() {
	*x = GetNewsByIdRequest{}
	mi := &file_news_v1_news_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNewsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsByIdRequest) ProtoMessage() {}

func (x *GetNewsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsByIdRequest.ProtoReflect.Descriptor instead.
func (*GetNewsByIdRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{3}
}

func (x *GetNewsByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ImageUrl string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *CreateNewsRequest) Reset() {
	*x = CreateNewsRequest{}
	mi := &file_news_v1_news_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewsRequest) ProtoMessage() {}

func (x *CreateNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewsRequest.ProtoReflect.Descriptor instead.
func (*CreateNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNewsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateNewsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNewsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateNewsRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type DeleteNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNewsRequest) Reset() {
	*x = DeleteNewsRequest{}
	mi := &file_news_v1_news_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNewsRequest) ProtoMessage() {}

func (x *DeleteNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNewsRequest.ProtoReflect.Descriptor instead.
func (*DeleteNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNewsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteNewsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id       string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Title    *string `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Content  *string `protobuf:"bytes,4,opt,name=content,proto3,oneof" json:"content,omitempty"`
	ImageUrl *string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
}

func (x *UpdateNewsRequest) Reset() {
	*x = UpdateNewsRequest{}
	mi := &file_news_v1_news_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewsRequest) ProtoMessage() {}

func (x *UpdateNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewsRequest.ProtoReflect.Descriptor instead.
func (*UpdateNewsRequest) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNewsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateNewsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateNewsRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateNewsRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *UpdateNewsRequest) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

type OperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	mi := &file_news_v1_news_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_news_v1_news_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_news_v1_news_proto_rawDescGZIP(), []int{7}
}

func (x *OperationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OperationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_news_v1_news_proto protoreflect.FileDescriptor

var file_news_v1_news_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x04, 0x6e, 0x65, 0x77, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x76, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x39, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22,
	0x47, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd4, 0x02, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x73, 0x46, 0x65, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12,
	0x1a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x1a, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x73, 0x12, 0x1a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x00, 0x42,
	0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x77,
	0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_news_v1_news_proto_rawDescOnce sync.Once
	file_news_v1_news_proto_rawDescData = file_news_v1_news_proto_rawDesc
)

func file_news_v1_news_proto_rawDescGZIP() []byte {
	file_news_v1_news_proto_rawDescOnce.Do(func() {
		file_news_v1_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_news_v1_news_proto_rawDescData)
	})
	return file_news_v1_news_proto_rawDescData
}

var file_news_v1_news_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_news_v1_news_proto_goTypes = []any{
	(*News)(nil),                // 0: news.v1.News
	(*GetNewsFeedRequest)(nil),  // 1: news.v1.GetNewsFeedRequest
	(*GetNewsFeedResponse)(nil), // 2: news.v1.GetNewsFeedResponse
	(*GetNewsByIdRequest)(nil),  // 3: news.v1.GetNewsByIdRequest
	(*CreateNewsRequest)(nil),   // 4: news.v1.CreateNewsRequest
	(*DeleteNewsRequest)(nil),   // 5: news.v1.DeleteNewsRequest
	(*UpdateNewsRequest)(nil),   // 6: news.v1.UpdateNewsRequest
	(*OperationResponse)(nil),   // 7: news.v1.OperationResponse
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_news_v1_news_proto_depIdxs = []int32{
	8, // 0: news.v1.News.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: news.v1.News.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: news.v1.GetNewsFeedResponse.news:type_name -> news.v1.News
	1, // 3: news.v1.NewsService.GetNewsFeed:input_type -> news.v1.GetNewsFeedRequest
	3, // 4: news.v1.NewsService.GetNewsById:input_type -> news.v1.GetNewsByIdRequest
	4, // 5: news.v1.NewsService.CreateNews:input_type -> news.v1.CreateNewsRequest
	5, // 6: news.v1.NewsService.DeleteNews:input_type -> news.v1.DeleteNewsRequest
	6, // 7: news.v1.NewsService.UpdateNews:input_type -> news.v1.UpdateNewsRequest
	2, // 8: news.v1.NewsService.GetNewsFeed:output_type -> news.v1.GetNewsFeedResponse
	0, // 9: news.v1.NewsService.GetNewsById:output_type -> news.v1.News
	0, // 10: news.v1.NewsService.CreateNews:output_type -> news.v1.News
	7, // 11: news.v1.NewsService.DeleteNews:output_type -> news.v1.OperationResponse
	0, // 12: news.v1.NewsService.UpdateNews:output_type -> news.v1.News
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_news_v1_news_proto_init() }
func file_news_v1_news_proto_init() {
	if File_news_v1_news_proto != nil {
		return
	}
	file_news_v1_news_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_news_v1_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_news_v1_news_proto_goTypes,
		DependencyIndexes: file_news_v1_news_proto_depIdxs,
		MessageInfos:      file_news_v1_news_proto_msgTypes,
	}.Build()
	File_news_v1_news_proto = out.File
	file_news_v1_news_proto_rawDesc = nil
	file_news_v1_news_proto_goTypes = nil
	file_news_v1_news_proto_depIdxs = nil
}
