// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: emailmessages.proto

package emailmessages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define the Request message
type EmailReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subject       string                 `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Recipients    []string               `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
	Body          string                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	SenderName    string                 `protobuf:"bytes,4,opt,name=senderName,proto3" json:"senderName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailReq) Reset() {
	*x = EmailReq{}
	mi := &file_emailmessages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailReq) ProtoMessage() {}

func (x *EmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_emailmessages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailReq.ProtoReflect.Descriptor instead.
func (*EmailReq) Descriptor() ([]byte, []int) {
	return file_emailmessages_proto_rawDescGZIP(), []int{0}
}

func (x *EmailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailReq) GetRecipients() []string {
	if x != nil {
		return x.Recipients
	}
	return nil
}

func (x *EmailReq) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailReq) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

type LetterMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LetterMessage) Reset() {
	*x = LetterMessage{}
	mi := &file_emailmessages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LetterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LetterMessage) ProtoMessage() {}

func (x *LetterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_emailmessages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LetterMessage.ProtoReflect.Descriptor instead.
func (*LetterMessage) Descriptor() ([]byte, []int) {
	return file_emailmessages_proto_rawDescGZIP(), []int{1}
}

func (x *LetterMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *LetterMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Broadcast struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BroadcastId   []byte                 `protobuf:"bytes,1,opt,name=broadcastId,proto3" json:"broadcastId,omitempty"`
	Sender        string                 `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Subscriber    string                 `protobuf:"bytes,3,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Broadcast) Reset() {
	*x = Broadcast{}
	mi := &file_emailmessages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Broadcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Broadcast) ProtoMessage() {}

func (x *Broadcast) ProtoReflect() protoreflect.Message {
	mi := &file_emailmessages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Broadcast.ProtoReflect.Descriptor instead.
func (*Broadcast) Descriptor() ([]byte, []int) {
	return file_emailmessages_proto_rawDescGZIP(), []int{2}
}

func (x *Broadcast) GetBroadcastId() []byte {
	if x != nil {
		return x.BroadcastId
	}
	return nil
}

func (x *Broadcast) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Broadcast) GetSubscriber() string {
	if x != nil {
		return x.Subscriber
	}
	return ""
}

func (x *Broadcast) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SubscriberList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subscribers   []string               `protobuf:"bytes,1,rep,name=subscribers,proto3" json:"subscribers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscriberList) Reset() {
	*x = SubscriberList{}
	mi := &file_emailmessages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberList) ProtoMessage() {}

func (x *SubscriberList) ProtoReflect() protoreflect.Message {
	mi := &file_emailmessages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberList.ProtoReflect.Descriptor instead.
func (*SubscriberList) Descriptor() ([]byte, []int) {
	return file_emailmessages_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriberList) GetSubscribers() []string {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

type GenericResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     []byte                 `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	mi := &file_emailmessages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emailmessages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_emailmessages_proto_rawDescGZIP(), []int{4}
}

func (x *GenericResponse) GetMessageId() []byte {
	if x != nil {
		return x.MessageId
	}
	return nil
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_emailmessages_proto protoreflect.FileDescriptor

var file_emailmessages_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x78, 0x0a, 0x08, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x0d, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x7f, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x32, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x81, 0x02, 0x0a, 0x13, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x1e, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x53, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x1d, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x6d, 0x63, 0x6e, 0x61, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_emailmessages_proto_rawDescOnce sync.Once
	file_emailmessages_proto_rawDescData []byte
)

func file_emailmessages_proto_rawDescGZIP() []byte {
	file_emailmessages_proto_rawDescOnce.Do(func() {
		file_emailmessages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_emailmessages_proto_rawDesc), len(file_emailmessages_proto_rawDesc)))
	})
	return file_emailmessages_proto_rawDescData
}

var file_emailmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_emailmessages_proto_goTypes = []any{
	(*EmailReq)(nil),        // 0: emailmessages.EmailReq
	(*LetterMessage)(nil),   // 1: emailmessages.LetterMessage
	(*Broadcast)(nil),       // 2: emailmessages.Broadcast
	(*SubscriberList)(nil),  // 3: emailmessages.SubscriberList
	(*GenericResponse)(nil), // 4: emailmessages.GenericResponse
}
var file_emailmessages_proto_depIdxs = []int32{
	0, // 0: emailmessages.EmailMessageService.SendEmail:input_type -> emailmessages.EmailReq
	1, // 1: emailmessages.EmailMessageService.BroadcastLetter:input_type -> emailmessages.LetterMessage
	3, // 2: emailmessages.EmailMessageService.UpdateScribers:input_type -> emailmessages.SubscriberList
	4, // 3: emailmessages.EmailMessageService.SendEmail:output_type -> emailmessages.GenericResponse
	2, // 4: emailmessages.EmailMessageService.BroadcastLetter:output_type -> emailmessages.Broadcast
	4, // 5: emailmessages.EmailMessageService.UpdateScribers:output_type -> emailmessages.GenericResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emailmessages_proto_init() }
func file_emailmessages_proto_init() {
	if File_emailmessages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_emailmessages_proto_rawDesc), len(file_emailmessages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emailmessages_proto_goTypes,
		DependencyIndexes: file_emailmessages_proto_depIdxs,
		MessageInfos:      file_emailmessages_proto_msgTypes,
	}.Build()
	File_emailmessages_proto = out.File
	file_emailmessages_proto_goTypes = nil
	file_emailmessages_proto_depIdxs = nil
}
