// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/calendar.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	// Opaque identifier of the event.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Status of the event. Possible values are:
	// "confirmed" - The event is confirmed. This is the default status.
	// "tentative" - The event is tentatively confirmed.
	// "cancelled" - The event is cancelled (deleted).
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// An absolute link to this event in the Google Calendar Web UI.
	HtmlLink string `protobuf:"bytes,3,opt,name=html_link,json=htmlLink,proto3" json:"html_link,omitempty"`
	// Creation time of the event (as a RFC3339 timestamp). Read-only.
	CreateTime string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Last modification time of the event (as a RFC3339 timestamp). Read-only.
	UpdateTime string `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Title of the event.
	Summary string `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
	// Description of the event. Can contain HTML. Optional.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Geographic location of the event as free-form text. Optional.
	Location string `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	// The creator of the event.
	Creator string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	// The organizer of the event.
	Organizer string `protobuf:"bytes,10,opt,name=organizer,proto3" json:"organizer,omitempty"`
	// The (inclusive) start time of the event (formatted according to RFC3339).
	StartTime string `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The (exclusive) end time of the event (formatted according to RFC3339).
	EndTime string `protobuf:"bytes,12,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Event unique identifier as defined in RFC5545.
	ICalUid string `protobuf:"bytes,13,opt,name=i_cal_uid,json=iCalUid,proto3" json:"i_cal_uid,omitempty"`
	// The attendees of the event.
	Attendees            []*Attendee `protobuf:"bytes,14,rep,name=attendees,proto3" json:"attendees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Event) GetHtmlLink() string {
	if m != nil {
		return m.HtmlLink
	}
	return ""
}

func (m *Event) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Event) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Event) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Event) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Event) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *Event) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Event) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Event) GetICalUid() string {
	if m != nil {
		return m.ICalUid
	}
	return ""
}

func (m *Event) GetAttendees() []*Attendee {
	if m != nil {
		return m.Attendees
	}
	return nil
}

type Attendee struct {
	// The attendee's Profile ID, if available.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The attendee's email address, if available.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// The attendee's name, if available.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Whether the attendee is the organizer of the event.
	Organizer bool `protobuf:"varint,4,opt,name=organizer,proto3" json:"organizer,omitempty"`
	// Whether this is an optional attendee.
	Optional bool `protobuf:"varint,5,opt,name=optional,proto3" json:"optional,omitempty"`
	// The attendee's response status. Possible values are:
	// "needsAction" - The attendee has not responded to the invitation.
	// "declined" - The attendee has declined the invitation.
	// "tentative" - The attendee has tentatively accepted the invitation.
	// "accepted" - The attendee has accepted the invitation.
	ResponseStatus string `protobuf:"bytes,6,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	// The attendee's response comment.
	Comment string `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	// Number of additional guests.
	AdditionalGuests     int64    `protobuf:"varint,8,opt,name=additionalGuests,proto3" json:"additionalGuests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attendee) Reset()         { *m = Attendee{} }
func (m *Attendee) String() string { return proto.CompactTextString(m) }
func (*Attendee) ProtoMessage()    {}
func (*Attendee) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{1}
}

func (m *Attendee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attendee.Unmarshal(m, b)
}
func (m *Attendee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attendee.Marshal(b, m, deterministic)
}
func (m *Attendee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attendee.Merge(m, src)
}
func (m *Attendee) XXX_Size() int {
	return xxx_messageInfo_Attendee.Size(m)
}
func (m *Attendee) XXX_DiscardUnknown() {
	xxx_messageInfo_Attendee.DiscardUnknown(m)
}

var xxx_messageInfo_Attendee proto.InternalMessageInfo

func (m *Attendee) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Attendee) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Attendee) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Attendee) GetOrganizer() bool {
	if m != nil {
		return m.Organizer
	}
	return false
}

func (m *Attendee) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *Attendee) GetResponseStatus() string {
	if m != nil {
		return m.ResponseStatus
	}
	return ""
}

func (m *Attendee) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Attendee) GetAdditionalGuests() int64 {
	if m != nil {
		return m.AdditionalGuests
	}
	return 0
}

type ListEventsRequest struct {
	// The parent resource name.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventsRequest) Reset()         { *m = ListEventsRequest{} }
func (m *ListEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEventsRequest) ProtoMessage()    {}
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{2}
}

func (m *ListEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventsRequest.Unmarshal(m, b)
}
func (m *ListEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventsRequest.Marshal(b, m, deterministic)
}
func (m *ListEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventsRequest.Merge(m, src)
}
func (m *ListEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEventsRequest.Size(m)
}
func (m *ListEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventsRequest proto.InternalMessageInfo

func (m *ListEventsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListEventsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListEventsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListEventsResponse struct {
	// The field name should match the noun "events" in the method name.  There
	// will be a maximum number of items returned based on the page_size field
	// in the request.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventsResponse) Reset()         { *m = ListEventsResponse{} }
func (m *ListEventsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEventsResponse) ProtoMessage()    {}
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{3}
}

func (m *ListEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventsResponse.Unmarshal(m, b)
}
func (m *ListEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventsResponse.Marshal(b, m, deterministic)
}
func (m *ListEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventsResponse.Merge(m, src)
}
func (m *ListEventsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEventsResponse.Size(m)
}
func (m *ListEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventsResponse proto.InternalMessageInfo

func (m *ListEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ListEventsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetEventRequest struct {
	// The field will contain name of the resource requested.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRequest) Reset()         { *m = GetEventRequest{} }
func (m *GetEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventRequest) ProtoMessage()    {}
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{4}
}

func (m *GetEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRequest.Unmarshal(m, b)
}
func (m *GetEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRequest.Marshal(b, m, deterministic)
}
func (m *GetEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRequest.Merge(m, src)
}
func (m *GetEventRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventRequest.Size(m)
}
func (m *GetEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRequest proto.InternalMessageInfo

func (m *GetEventRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateEventRequest struct {
	// The parent resource name where the event is to be created.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The text describing the event to be created.
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{5}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateEventRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type DeleteEventRequest struct {
	// The resource name of the event to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9149431be0552851, []int{6}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "kordevaus.Event")
	proto.RegisterType((*Attendee)(nil), "kordevaus.Attendee")
	proto.RegisterType((*ListEventsRequest)(nil), "kordevaus.ListEventsRequest")
	proto.RegisterType((*ListEventsResponse)(nil), "kordevaus.ListEventsResponse")
	proto.RegisterType((*GetEventRequest)(nil), "kordevaus.GetEventRequest")
	proto.RegisterType((*CreateEventRequest)(nil), "kordevaus.CreateEventRequest")
	proto.RegisterType((*DeleteEventRequest)(nil), "kordevaus.DeleteEventRequest")
}

func init() { proto.RegisterFile("pb/calendar.proto", fileDescriptor_9149431be0552851) }

var fileDescriptor_9149431be0552851 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcb, 0x4e, 0x23, 0x47,
	0x14, 0x95, 0x9f, 0xb4, 0xaf, 0x01, 0x43, 0x25, 0x22, 0x9d, 0x06, 0x04, 0xb4, 0x94, 0x60, 0x79,
	0xe1, 0x06, 0x22, 0x36, 0x48, 0x91, 0x42, 0x08, 0x62, 0x83, 0xa2, 0xc8, 0x90, 0x4d, 0x36, 0xad,
	0xb2, 0xfb, 0x62, 0x4a, 0xee, 0xae, 0xee, 0x74, 0x95, 0x11, 0x30, 0x62, 0x33, 0xbf, 0x30, 0x1f,
	0x31, 0xff, 0x32, 0xdb, 0x99, 0x4f, 0x98, 0xcf, 0x98, 0xc5, 0xa8, 0x1e, 0x6d, 0x1a, 0x0c, 0x62,
	0xe5, 0xbe, 0xe7, 0x5c, 0xdd, 0x5b, 0xe7, 0x9c, 0x72, 0xc1, 0x6a, 0x36, 0x0c, 0x46, 0x34, 0x46,
	0x1e, 0xd1, 0xbc, 0x9f, 0xe5, 0xa9, 0x4c, 0x49, 0x6b, 0x92, 0xe6, 0x11, 0xde, 0xd0, 0xa9, 0xf0,
	0x36, 0xc6, 0x69, 0x3a, 0x8e, 0x31, 0xa0, 0x19, 0x0b, 0x28, 0xe7, 0xa9, 0xa4, 0x92, 0xa5, 0x5c,
	0x98, 0x46, 0xef, 0xa7, 0x12, 0x3b, 0x8a, 0x19, 0x72, 0x69, 0x89, 0x75, 0x4b, 0xe8, 0x6a, 0x38,
	0xbd, 0x0a, 0x30, 0xc9, 0xe4, 0x9d, 0x21, 0xfd, 0x8f, 0x35, 0x68, 0x9c, 0xde, 0x20, 0x97, 0x64,
	0x19, 0xaa, 0x2c, 0x72, 0x2b, 0xdb, 0x95, 0x6e, 0x6b, 0x50, 0x65, 0x11, 0x59, 0x83, 0xa6, 0x90,
	0x54, 0x4e, 0x85, 0x5b, 0xd5, 0x98, 0xad, 0xc8, 0x3a, 0xb4, 0xae, 0x65, 0x12, 0x87, 0x31, 0xe3,
	0x13, 0xb7, 0xa6, 0x29, 0x47, 0x01, 0xe7, 0x8c, 0x4f, 0xc8, 0x16, 0xb4, 0x47, 0x39, 0x52, 0x89,
	0xa1, 0x64, 0x09, 0xba, 0x75, 0x4d, 0x83, 0x81, 0x2e, 0x59, 0x82, 0xaa, 0x61, 0x9a, 0x45, 0xb3,
	0x86, 0x86, 0x69, 0x30, 0x90, 0x6e, 0x70, 0x61, 0x41, 0x4c, 0x93, 0x84, 0xe6, 0x77, 0x6e, 0x53,
	0x93, 0x45, 0x49, 0xb6, 0xa1, 0x1d, 0xa1, 0x18, 0xe5, 0x2c, 0x53, 0xb2, 0xdd, 0x05, 0xcd, 0x96,
	0x21, 0xe2, 0x81, 0x13, 0xa7, 0x23, 0xed, 0x8a, 0xeb, 0x98, 0x93, 0x15, 0xb5, 0x9a, 0xab, 0x8f,
	0x91, 0xe6, 0x6e, 0xcb, 0xcc, 0xb5, 0x25, 0xd9, 0x80, 0x56, 0x9a, 0x8f, 0x29, 0x67, 0xf7, 0x98,
	0xbb, 0xa0, 0xb9, 0x47, 0x80, 0x6c, 0x02, 0x08, 0x49, 0x73, 0x69, 0xce, 0xdb, 0x36, 0xb4, 0x46,
	0xf4, 0x71, 0x7f, 0x06, 0x07, 0x79, 0x64, 0xc8, 0x45, 0x33, 0x17, 0x79, 0xa4, 0x29, 0x0f, 0x5a,
	0x2c, 0x1c, 0xd1, 0x38, 0x9c, 0xb2, 0xc8, 0x5d, 0x32, 0x1c, 0x3b, 0xa1, 0xf1, 0xbf, 0x2c, 0x22,
	0xfb, 0xd0, 0xa2, 0x52, 0x22, 0x8f, 0x10, 0x85, 0xbb, 0xbc, 0x5d, 0xeb, 0xb6, 0x0f, 0x7e, 0xe8,
	0xcf, 0x92, 0xee, 0x1f, 0x5b, 0x6e, 0xf0, 0xd8, 0xe5, 0x7f, 0xab, 0x80, 0x53, 0xe0, 0x73, 0x61,
	0xfd, 0x08, 0x0d, 0x4c, 0x28, 0x8b, 0x6d, 0x56, 0xa6, 0x20, 0x3b, 0xb0, 0x18, 0x31, 0x91, 0xc5,
	0xf4, 0x2e, 0xe4, 0x34, 0x41, 0x9b, 0x56, 0xdb, 0x62, 0x7f, 0xd3, 0x04, 0x9f, 0x8a, 0x57, 0x71,
	0x39, 0x65, 0xf1, 0x1e, 0x38, 0xa9, 0xb6, 0x96, 0xc6, 0x3a, 0x2a, 0x67, 0x30, 0xab, 0xc9, 0x2e,
	0x74, 0x72, 0x14, 0x59, 0xca, 0x05, 0x86, 0xf6, 0xa2, 0x98, 0xc0, 0x96, 0x0b, 0xf8, 0xc2, 0x5c,
	0x18, 0xe5, 0x7c, 0x9a, 0x24, 0xc8, 0xa5, 0xcd, 0xac, 0x28, 0x49, 0x0f, 0x56, 0x68, 0x14, 0x31,
	0x33, 0xf0, 0x6c, 0x8a, 0x42, 0x0a, 0x9d, 0x5b, 0x6d, 0x30, 0x87, 0xfb, 0x63, 0x58, 0x3d, 0x67,
	0x42, 0xea, 0xbb, 0x2a, 0x06, 0xf8, 0xbf, 0x42, 0xd5, 0x1d, 0xcd, 0x68, 0xae, 0x26, 0x1b, 0x2b,
	0x6c, 0xa5, 0xee, 0x68, 0x46, 0xc7, 0x18, 0x0a, 0x76, 0x8f, 0xda, 0x92, 0xc6, 0xc0, 0x51, 0xc0,
	0x05, 0xbb, 0x47, 0x95, 0xa8, 0x26, 0x65, 0x3a, 0x41, 0x6e, 0x3d, 0xd1, 0xed, 0x97, 0x0a, 0xf0,
	0xaf, 0x80, 0x94, 0x17, 0x19, 0x29, 0xa4, 0x0b, 0x4d, 0xd4, 0x88, 0x5b, 0xd1, 0x69, 0xad, 0x94,
	0xd2, 0xd2, 0xad, 0x03, 0xcb, 0x93, 0x5f, 0xa1, 0xc3, 0xf1, 0x56, 0x86, 0xa5, 0x1d, 0x26, 0x94,
	0x25, 0x05, 0xff, 0x33, 0xdb, 0xf3, 0x0b, 0x74, 0xce, 0xd0, 0xac, 0x29, 0xe4, 0x10, 0xa8, 0xeb,
	0x9c, 0x8c, 0x18, 0xfd, 0xed, 0xff, 0x01, 0xe4, 0x44, 0xff, 0x7d, 0x9e, 0x74, 0xbe, 0x26, 0x9c,
	0x40, 0x5d, 0xe2, 0xad, 0xb4, 0x1b, 0xf5, 0xb7, 0xdf, 0x05, 0xf2, 0x17, 0xc6, 0xf8, 0x6c, 0xc2,
	0x0b, 0xbb, 0x0e, 0xbe, 0xd4, 0xc0, 0x39, 0xb1, 0xcf, 0x0f, 0xc9, 0x01, 0x1e, 0x7d, 0x20, 0x1b,
	0x25, 0xbd, 0x73, 0x39, 0x78, 0x9b, 0xaf, 0xb0, 0xc6, 0x3c, 0x7f, 0xf7, 0xfd, 0xe7, 0xaf, 0x1f,
	0xaa, 0x3b, 0x64, 0x2b, 0xb8, 0xd9, 0x0f, 0xde, 0x99, 0xa3, 0xfe, 0x5e, 0xbc, 0x73, 0x22, 0xe8,
	0x3d, 0x04, 0xd6, 0x3b, 0x0a, 0x4e, 0xe1, 0x09, 0xf1, 0x4a, 0x33, 0x9f, 0x19, 0xe5, 0xcd, 0xb9,
	0xff, 0x6c, 0x85, 0x52, 0x53, 0x5e, 0x60, 0xe7, 0x07, 0xbd, 0x07, 0xc2, 0xa1, 0x5d, 0xf2, 0x93,
	0x94, 0x4f, 0x3e, 0xef, 0xf3, 0x0b, 0x8b, 0xfa, 0x7a, 0x51, 0xd7, 0x7f, 0x4b, 0xcb, 0x51, 0x43,
	0xff, 0x92, 0x04, 0xda, 0x25, 0xf7, 0x9f, 0xec, 0x9b, 0x4f, 0xc5, 0x5b, 0xeb, 0x9b, 0xc7, 0xba,
	0x5f, 0x3c, 0xd6, 0xfd, 0x53, 0xf5, 0x58, 0x17, 0xf2, 0x7a, 0x6f, 0xc9, 0xf3, 0xc8, 0xa7, 0xe3,
	0x8e, 0x7a, 0xf3, 0xe2, 0xeb, 0x54, 0xc8, 0xa3, 0xc3, 0xbd, 0xbd, 0xc3, 0xfd, 0x3f, 0xeb, 0xff,
	0x55, 0xb3, 0xe1, 0xb0, 0xa9, 0x47, 0xfe, 0xf6, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x86, 0x65, 0x9b,
	0x68, 0x64, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	// Lists events in a calendar.
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	// Gets an event.
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error)
	// Creates an event in a calendar.
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error)
	// Deletes an event.
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type calendarClient struct {
	cc *grpc.ClientConn
}

func NewCalendarClient(cc *grpc.ClientConn) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/kordevaus.Calendar/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/kordevaus.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/kordevaus.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kordevaus.Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	// Lists events in a calendar.
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	// Gets an event.
	GetEvent(context.Context, *GetEventRequest) (*Event, error)
	// Creates an event in a calendar.
	CreateEvent(context.Context, *CreateEventRequest) (*Event, error)
	// Deletes an event.
	DeleteEvent(context.Context, *DeleteEventRequest) (*empty.Empty, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) ListEvents(ctx context.Context, req *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *GetEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kordevaus.Calendar/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kordevaus.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kordevaus.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kordevaus.Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kordevaus.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEvents",
			Handler:    _Calendar_ListEvents_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/calendar.proto",
}
