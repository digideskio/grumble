// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

package freezer

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Server struct {
	Config           []*ConfigKeyValuePair `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
	BanList          *BanList              `protobuf:"bytes,3,opt,name=ban_list" json:"ban_list,omitempty"`
	Channels         []*Channel            `protobuf:"bytes,4,rep,name=channels" json:"channels,omitempty"`
	Users            []*User               `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *Server) Reset()         { *this = Server{} }
func (this *Server) String() string { return proto.CompactTextString(this) }
func (*Server) ProtoMessage()       {}

func (this *Server) GetBanList() *BanList {
	if this != nil {
		return this.BanList
	}
	return nil
}

type ConfigKeyValuePair struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ConfigKeyValuePair) Reset()         { *this = ConfigKeyValuePair{} }
func (this *ConfigKeyValuePair) String() string { return proto.CompactTextString(this) }
func (*ConfigKeyValuePair) ProtoMessage()       {}

func (this *ConfigKeyValuePair) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *ConfigKeyValuePair) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type Ban struct {
	Ip               []byte  `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask             *uint32 `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	Username         *string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	CertHash         *string `protobuf:"bytes,4,opt,name=cert_hash" json:"cert_hash,omitempty"`
	Reason           *string `protobuf:"bytes,5,opt,name=reason" json:"reason,omitempty"`
	Start            *int64  `protobuf:"varint,6,opt,name=start" json:"start,omitempty"`
	Duration         *uint32 `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Ban) Reset()         { *this = Ban{} }
func (this *Ban) String() string { return proto.CompactTextString(this) }
func (*Ban) ProtoMessage()       {}

func (this *Ban) GetIp() []byte {
	if this != nil {
		return this.Ip
	}
	return nil
}

func (this *Ban) GetMask() uint32 {
	if this != nil && this.Mask != nil {
		return *this.Mask
	}
	return 0
}

func (this *Ban) GetUsername() string {
	if this != nil && this.Username != nil {
		return *this.Username
	}
	return ""
}

func (this *Ban) GetCertHash() string {
	if this != nil && this.CertHash != nil {
		return *this.CertHash
	}
	return ""
}

func (this *Ban) GetReason() string {
	if this != nil && this.Reason != nil {
		return *this.Reason
	}
	return ""
}

func (this *Ban) GetStart() int64 {
	if this != nil && this.Start != nil {
		return *this.Start
	}
	return 0
}

func (this *Ban) GetDuration() uint32 {
	if this != nil && this.Duration != nil {
		return *this.Duration
	}
	return 0
}

type BanList struct {
	Bans             []*Ban `protobuf:"bytes,1,rep,name=bans" json:"bans,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *BanList) Reset()         { *this = BanList{} }
func (this *BanList) String() string { return proto.CompactTextString(this) }
func (*BanList) ProtoMessage()       {}

type User struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password         *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	CertHash         *string `protobuf:"bytes,4,opt,name=cert_hash" json:"cert_hash,omitempty"`
	Email            *string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	TextureBlob      *string `protobuf:"bytes,6,opt,name=texture_blob" json:"texture_blob,omitempty"`
	CommentBlob      *string `protobuf:"bytes,7,opt,name=comment_blob" json:"comment_blob,omitempty"`
	LastChannelId    *uint32 `protobuf:"varint,8,opt,name=last_channel_id" json:"last_channel_id,omitempty"`
	LastActive       *uint64 `protobuf:"varint,9,opt,name=last_active" json:"last_active,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *User) Reset()         { *this = User{} }
func (this *User) String() string { return proto.CompactTextString(this) }
func (*User) ProtoMessage()       {}

func (this *User) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *User) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *User) GetPassword() string {
	if this != nil && this.Password != nil {
		return *this.Password
	}
	return ""
}

func (this *User) GetCertHash() string {
	if this != nil && this.CertHash != nil {
		return *this.CertHash
	}
	return ""
}

func (this *User) GetEmail() string {
	if this != nil && this.Email != nil {
		return *this.Email
	}
	return ""
}

func (this *User) GetTextureBlob() string {
	if this != nil && this.TextureBlob != nil {
		return *this.TextureBlob
	}
	return ""
}

func (this *User) GetCommentBlob() string {
	if this != nil && this.CommentBlob != nil {
		return *this.CommentBlob
	}
	return ""
}

func (this *User) GetLastChannelId() uint32 {
	if this != nil && this.LastChannelId != nil {
		return *this.LastChannelId
	}
	return 0
}

func (this *User) GetLastActive() uint64 {
	if this != nil && this.LastActive != nil {
		return *this.LastActive
	}
	return 0
}

type UserRemove struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *UserRemove) Reset()         { *this = UserRemove{} }
func (this *UserRemove) String() string { return proto.CompactTextString(this) }
func (*UserRemove) ProtoMessage()       {}

func (this *UserRemove) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type Channel struct {
	Id               *uint32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ParentId         *uint32  `protobuf:"varint,3,opt,name=parent_id" json:"parent_id,omitempty"`
	Position         *int64   `protobuf:"varint,4,opt,name=position" json:"position,omitempty"`
	InheritAcl       *bool    `protobuf:"varint,5,opt,name=inherit_acl" json:"inherit_acl,omitempty"`
	Links            []uint32 `protobuf:"varint,6,rep,name=links" json:"links,omitempty"`
	Acl              []*ACL   `protobuf:"bytes,7,rep,name=acl" json:"acl,omitempty"`
	Groups           []*Group `protobuf:"bytes,8,rep,name=groups" json:"groups,omitempty"`
	DescriptionBlob  *string  `protobuf:"bytes,9,opt,name=description_blob" json:"description_blob,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Channel) Reset()         { *this = Channel{} }
func (this *Channel) String() string { return proto.CompactTextString(this) }
func (*Channel) ProtoMessage()       {}

func (this *Channel) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Channel) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Channel) GetParentId() uint32 {
	if this != nil && this.ParentId != nil {
		return *this.ParentId
	}
	return 0
}

func (this *Channel) GetPosition() int64 {
	if this != nil && this.Position != nil {
		return *this.Position
	}
	return 0
}

func (this *Channel) GetInheritAcl() bool {
	if this != nil && this.InheritAcl != nil {
		return *this.InheritAcl
	}
	return false
}

func (this *Channel) GetDescriptionBlob() string {
	if this != nil && this.DescriptionBlob != nil {
		return *this.DescriptionBlob
	}
	return ""
}

type ChannelRemove struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ChannelRemove) Reset()         { *this = ChannelRemove{} }
func (this *ChannelRemove) String() string { return proto.CompactTextString(this) }
func (*ChannelRemove) ProtoMessage()       {}

func (this *ChannelRemove) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type ACL struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	Group            *string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	ApplyHere        *bool   `protobuf:"varint,3,opt,name=apply_here" json:"apply_here,omitempty"`
	ApplySubs        *bool   `protobuf:"varint,4,opt,name=apply_subs" json:"apply_subs,omitempty"`
	Allow            *uint32 `protobuf:"varint,5,opt,name=allow" json:"allow,omitempty"`
	Deny             *uint32 `protobuf:"varint,6,opt,name=deny" json:"deny,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ACL) Reset()         { *this = ACL{} }
func (this *ACL) String() string { return proto.CompactTextString(this) }
func (*ACL) ProtoMessage()       {}

func (this *ACL) GetUserId() uint32 {
	if this != nil && this.UserId != nil {
		return *this.UserId
	}
	return 0
}

func (this *ACL) GetGroup() string {
	if this != nil && this.Group != nil {
		return *this.Group
	}
	return ""
}

func (this *ACL) GetApplyHere() bool {
	if this != nil && this.ApplyHere != nil {
		return *this.ApplyHere
	}
	return false
}

func (this *ACL) GetApplySubs() bool {
	if this != nil && this.ApplySubs != nil {
		return *this.ApplySubs
	}
	return false
}

func (this *ACL) GetAllow() uint32 {
	if this != nil && this.Allow != nil {
		return *this.Allow
	}
	return 0
}

func (this *ACL) GetDeny() uint32 {
	if this != nil && this.Deny != nil {
		return *this.Deny
	}
	return 0
}

type Group struct {
	Name             *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inherit          *bool    `protobuf:"varint,2,opt,name=inherit" json:"inherit,omitempty"`
	Inheritable      *bool    `protobuf:"varint,3,opt,name=inheritable" json:"inheritable,omitempty"`
	Add              []uint32 `protobuf:"varint,4,rep,name=add" json:"add,omitempty"`
	Remove           []uint32 `protobuf:"varint,5,rep,name=remove" json:"remove,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Group) Reset()         { *this = Group{} }
func (this *Group) String() string { return proto.CompactTextString(this) }
func (*Group) ProtoMessage()       {}

func (this *Group) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Group) GetInherit() bool {
	if this != nil && this.Inherit != nil {
		return *this.Inherit
	}
	return false
}

func (this *Group) GetInheritable() bool {
	if this != nil && this.Inheritable != nil {
		return *this.Inheritable
	}
	return false
}

func init() {
}
