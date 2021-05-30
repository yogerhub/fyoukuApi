// Code generated by protoc-gen-go. DO NOT EDIT.
// source: video.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	video.proto

It has these top-level messages:
	RequestChannelAdvert
	ResponseChannelAdvert
	ChannelAdvertData
	RequestChannelHotList
	ResponseChannelHotList
	ChannelHotList
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestChannelAdvert struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channelId" json:"channelId,omitempty"`
}

func (m *RequestChannelAdvert) Reset()                    { *m = RequestChannelAdvert{} }
func (m *RequestChannelAdvert) String() string            { return proto1.CompactTextString(m) }
func (*RequestChannelAdvert) ProtoMessage()               {}
func (*RequestChannelAdvert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestChannelAdvert) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type ResponseChannelAdvert struct {
	Code  int64                `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg   string               `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Items []*ChannelAdvertData `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Count int64                `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *ResponseChannelAdvert) Reset()                    { *m = ResponseChannelAdvert{} }
func (m *ResponseChannelAdvert) String() string            { return proto1.CompactTextString(m) }
func (*ResponseChannelAdvert) ProtoMessage()               {}
func (*ResponseChannelAdvert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseChannelAdvert) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseChannelAdvert) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseChannelAdvert) GetItems() []*ChannelAdvertData {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ResponseChannelAdvert) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ChannelAdvertData struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	SubTitle string `protobuf:"bytes,3,opt,name=subTitle" json:"subTitle,omitempty"`
	AddTime  int64  `protobuf:"varint,4,opt,name=addTime" json:"addTime,omitempty"`
	Img      string `protobuf:"bytes,5,opt,name=img" json:"img,omitempty"`
	Url      string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
}

func (m *ChannelAdvertData) Reset()                    { *m = ChannelAdvertData{} }
func (m *ChannelAdvertData) String() string            { return proto1.CompactTextString(m) }
func (*ChannelAdvertData) ProtoMessage()               {}
func (*ChannelAdvertData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChannelAdvertData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChannelAdvertData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ChannelAdvertData) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *ChannelAdvertData) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

func (m *ChannelAdvertData) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *ChannelAdvertData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type RequestChannelHotList struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channelId" json:"channelId,omitempty"`
}

func (m *RequestChannelHotList) Reset()                    { *m = RequestChannelHotList{} }
func (m *RequestChannelHotList) String() string            { return proto1.CompactTextString(m) }
func (*RequestChannelHotList) ProtoMessage()               {}
func (*RequestChannelHotList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestChannelHotList) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type ResponseChannelHotList struct {
	Code  int64             `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg   string            `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Items []*ChannelHotList `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Count int64             `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *ResponseChannelHotList) Reset()                    { *m = ResponseChannelHotList{} }
func (m *ResponseChannelHotList) String() string            { return proto1.CompactTextString(m) }
func (*ResponseChannelHotList) ProtoMessage()               {}
func (*ResponseChannelHotList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseChannelHotList) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseChannelHotList) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseChannelHotList) GetItems() []*ChannelHotList {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ResponseChannelHotList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ChannelHotList struct {
	Id            int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	SubTitle      string `protobuf:"bytes,3,opt,name=subTitle" json:"subTitle,omitempty"`
	AddTime       int64  `protobuf:"varint,4,opt,name=addTime" json:"addTime,omitempty"`
	Img           string `protobuf:"bytes,5,opt,name=img" json:"img,omitempty"`
	Img1          string `protobuf:"bytes,6,opt,name=img1" json:"img1,omitempty"`
	EpisodesCount int64  `protobuf:"varint,7,opt,name=episodesCount" json:"episodesCount,omitempty"`
	IsEnd         int64  `protobuf:"varint,8,opt,name=isEnd" json:"isEnd,omitempty"`
	Comment       int64  `protobuf:"varint,9,opt,name=comment" json:"comment,omitempty"`
}

func (m *ChannelHotList) Reset()                    { *m = ChannelHotList{} }
func (m *ChannelHotList) String() string            { return proto1.CompactTextString(m) }
func (*ChannelHotList) ProtoMessage()               {}
func (*ChannelHotList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChannelHotList) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChannelHotList) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ChannelHotList) GetSubTitle() string {
	if m != nil {
		return m.SubTitle
	}
	return ""
}

func (m *ChannelHotList) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

func (m *ChannelHotList) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *ChannelHotList) GetImg1() string {
	if m != nil {
		return m.Img1
	}
	return ""
}

func (m *ChannelHotList) GetEpisodesCount() int64 {
	if m != nil {
		return m.EpisodesCount
	}
	return 0
}

func (m *ChannelHotList) GetIsEnd() int64 {
	if m != nil {
		return m.IsEnd
	}
	return 0
}

func (m *ChannelHotList) GetComment() int64 {
	if m != nil {
		return m.Comment
	}
	return 0
}

func init() {
	proto1.RegisterType((*RequestChannelAdvert)(nil), "proto.RequestChannelAdvert")
	proto1.RegisterType((*ResponseChannelAdvert)(nil), "proto.ResponseChannelAdvert")
	proto1.RegisterType((*ChannelAdvertData)(nil), "proto.ChannelAdvertData")
	proto1.RegisterType((*RequestChannelHotList)(nil), "proto.RequestChannelHotList")
	proto1.RegisterType((*ResponseChannelHotList)(nil), "proto.ResponseChannelHotList")
	proto1.RegisterType((*ChannelHotList)(nil), "proto.ChannelHotList")
}

func init() { proto1.RegisterFile("video.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xdd, 0xea, 0xd3, 0x40,
	0x10, 0xc5, 0xff, 0x69, 0x9a, 0x7e, 0x4c, 0x6d, 0xd1, 0xa5, 0x95, 0xa5, 0x56, 0x28, 0xc1, 0x8b,
	0x82, 0x50, 0xb0, 0xea, 0x03, 0x48, 0x15, 0x14, 0x0a, 0x42, 0x2c, 0xde, 0xa7, 0xd9, 0x21, 0x2e,
	0x64, 0xb3, 0x35, 0xbb, 0xe9, 0xbd, 0xf8, 0x08, 0x3e, 0x87, 0xef, 0xe5, 0x63, 0xc8, 0x7e, 0xb4,
	0x9a, 0x36, 0xa8, 0x57, 0x5e, 0x75, 0x67, 0x4e, 0xe7, 0x9c, 0xdd, 0x1f, 0x19, 0x18, 0x9d, 0x38,
	0x43, 0xb9, 0x3e, 0x56, 0x52, 0x4b, 0x12, 0xd9, 0x9f, 0xf8, 0x05, 0x4c, 0x13, 0xfc, 0x5c, 0xa3,
	0xd2, 0xdb, 0x4f, 0x69, 0x59, 0x62, 0xf1, 0x8a, 0x9d, 0xb0, 0xd2, 0x64, 0x01, 0xc3, 0xcc, 0x35,
	0xde, 0x31, 0x1a, 0x2c, 0x83, 0xd5, 0x30, 0xf9, 0xd5, 0x88, 0xbf, 0x06, 0x30, 0x4b, 0x50, 0x1d,
	0x65, 0xa9, 0xb0, 0x39, 0x47, 0xa0, 0x9b, 0x49, 0x86, 0x76, 0x24, 0x4c, 0xec, 0x99, 0xdc, 0x87,
	0x50, 0xa8, 0x9c, 0x76, 0xac, 0x8b, 0x39, 0x92, 0x35, 0x44, 0x5c, 0xa3, 0x50, 0x34, 0x5c, 0x86,
	0xab, 0xd1, 0x86, 0xba, 0x3b, 0xad, 0x1b, 0x56, 0xaf, 0x53, 0x9d, 0x26, 0xee, 0x6f, 0x64, 0x0a,
	0x51, 0x26, 0xeb, 0x52, 0xd3, 0xae, 0xb5, 0x75, 0x45, 0xfc, 0x2d, 0x80, 0x07, 0x37, 0x23, 0x64,
	0x02, 0x1d, 0xce, 0x7c, 0x7e, 0x87, 0x33, 0x33, 0xab, 0xb9, 0x2e, 0xd0, 0xe7, 0xbb, 0x82, 0xcc,
	0x61, 0xa0, 0xea, 0xc3, 0xde, 0x0a, 0xa1, 0x15, 0x2e, 0x35, 0xa1, 0xd0, 0x4f, 0x19, 0xdb, 0x73,
	0x81, 0x3e, 0xef, 0x5c, 0x9a, 0x97, 0x70, 0x91, 0xd3, 0xc8, 0xbd, 0x84, 0x8b, 0xdc, 0x74, 0xea,
	0xaa, 0xa0, 0x3d, 0xd7, 0xa9, 0xab, 0x22, 0x7e, 0x69, 0xd0, 0xfc, 0x4e, 0xf4, 0xad, 0xd4, 0x3b,
	0xae, 0xfe, 0x86, 0xf4, 0x4b, 0x00, 0x0f, 0xaf, 0x90, 0x9e, 0x07, 0xff, 0x8d, 0xe9, 0xd3, 0x26,
	0xd3, 0x59, 0x93, 0xa9, 0xf7, 0xfa, 0x33, 0xd0, 0x1f, 0x01, 0x4c, 0xae, 0xb2, 0xff, 0x2f, 0x4d,
	0x02, 0x5d, 0x2e, 0xf2, 0x67, 0x1e, 0xa7, 0x3d, 0x93, 0x27, 0x30, 0xc6, 0x23, 0x57, 0x92, 0xa1,
	0xda, 0xda, 0x2b, 0xf7, 0xad, 0x4b, 0xb3, 0x69, 0xee, 0xc5, 0xd5, 0x9b, 0x92, 0xd1, 0x81, 0x7b,
	0x90, 0x2d, 0x4c, 0x76, 0x26, 0x85, 0xc0, 0x52, 0xd3, 0xa1, 0xcb, 0xf6, 0xe5, 0xe6, 0x7b, 0x00,
	0xf7, 0x3e, 0x9a, 0x75, 0xf8, 0x80, 0xd5, 0x89, 0x67, 0x48, 0x76, 0x30, 0x6e, 0x7e, 0xc9, 0x8f,
	0x3c, 0xc0, 0xb6, 0xf5, 0x98, 0x2f, 0x2e, 0x62, 0xcb, 0x12, 0xc4, 0x77, 0xe4, 0xfd, 0x0d, 0xc8,
	0x45, 0xab, 0x9d, 0x57, 0xe7, 0x8f, 0xdb, 0xfd, 0xbc, 0x1c, 0xdf, 0x1d, 0x7a, 0x56, 0x7f, 0xfe,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xd1, 0x8f, 0x70, 0xc4, 0x03, 0x00, 0x00,
}
