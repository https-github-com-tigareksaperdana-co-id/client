// Auto-generated by avdl-compiler v1.3.11 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/common.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type ThreadID []byte
type MessageID uint
type TopicID []byte
type ConversationID []byte
type TLFID []byte
type Hash []byte
type InboxVers uint64
type OutboxID []byte
type MessageType int

const (
	MessageType_NONE               MessageType = 0
	MessageType_TEXT               MessageType = 1
	MessageType_ATTACHMENT         MessageType = 2
	MessageType_EDIT               MessageType = 3
	MessageType_DELETE             MessageType = 4
	MessageType_METADATA           MessageType = 5
	MessageType_TLFNAME            MessageType = 6
	MessageType_HEADLINE           MessageType = 7
	MessageType_ATTACHMENTUPLOADED MessageType = 8
)

var MessageTypeMap = map[string]MessageType{
	"NONE":               0,
	"TEXT":               1,
	"ATTACHMENT":         2,
	"EDIT":               3,
	"DELETE":             4,
	"METADATA":           5,
	"TLFNAME":            6,
	"HEADLINE":           7,
	"ATTACHMENTUPLOADED": 8,
}

var MessageTypeRevMap = map[MessageType]string{
	0: "NONE",
	1: "TEXT",
	2: "ATTACHMENT",
	3: "EDIT",
	4: "DELETE",
	5: "METADATA",
	6: "TLFNAME",
	7: "HEADLINE",
	8: "ATTACHMENTUPLOADED",
}

type TopicType int

const (
	TopicType_NONE TopicType = 0
	TopicType_CHAT TopicType = 1
	TopicType_DEV  TopicType = 2
)

var TopicTypeMap = map[string]TopicType{
	"NONE": 0,
	"CHAT": 1,
	"DEV":  2,
}

var TopicTypeRevMap = map[TopicType]string{
	0: "NONE",
	1: "CHAT",
	2: "DEV",
}

type ConversationStatus int

const (
	ConversationStatus_UNFILED  ConversationStatus = 0
	ConversationStatus_FAVORITE ConversationStatus = 1
	ConversationStatus_IGNORED  ConversationStatus = 2
	ConversationStatus_BLOCKED  ConversationStatus = 3
	ConversationStatus_MUTED    ConversationStatus = 4
)

var ConversationStatusMap = map[string]ConversationStatus{
	"UNFILED":  0,
	"FAVORITE": 1,
	"IGNORED":  2,
	"BLOCKED":  3,
	"MUTED":    4,
}

var ConversationStatusRevMap = map[ConversationStatus]string{
	0: "UNFILED",
	1: "FAVORITE",
	2: "IGNORED",
	3: "BLOCKED",
	4: "MUTED",
}

func (e ConversationStatus) String() string {
	if v, ok := ConversationStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type Pagination struct {
	Next     []byte `codec:"next" json:"next"`
	Previous []byte `codec:"previous" json:"previous"`
	Num      int    `codec:"num" json:"num"`
	Last     bool   `codec:"last" json:"last"`
}

type RateLimit struct {
	Name           string `codec:"name" json:"name"`
	CallsRemaining int    `codec:"callsRemaining" json:"callsRemaining"`
	WindowReset    int    `codec:"windowReset" json:"windowReset"`
	MaxCalls       int    `codec:"maxCalls" json:"maxCalls"`
}

type TLFVisibility int

const (
	TLFVisibility_ANY     TLFVisibility = 0
	TLFVisibility_PUBLIC  TLFVisibility = 1
	TLFVisibility_PRIVATE TLFVisibility = 2
)

var TLFVisibilityMap = map[string]TLFVisibility{
	"ANY":     0,
	"PUBLIC":  1,
	"PRIVATE": 2,
}

var TLFVisibilityRevMap = map[TLFVisibility]string{
	0: "ANY",
	1: "PUBLIC",
	2: "PRIVATE",
}

func (e TLFVisibility) String() string {
	if v, ok := TLFVisibilityRevMap[e]; ok {
		return v
	}
	return ""
}

type GetInboxQuery struct {
	ConvID            *ConversationID      `codec:"convID,omitempty" json:"convID,omitempty"`
	TopicType         *TopicType           `codec:"topicType,omitempty" json:"topicType,omitempty"`
	TlfID             *TLFID               `codec:"tlfID,omitempty" json:"tlfID,omitempty"`
	TlfVisibility     *TLFVisibility       `codec:"tlfVisibility,omitempty" json:"tlfVisibility,omitempty"`
	Before            *gregor1.Time        `codec:"before,omitempty" json:"before,omitempty"`
	After             *gregor1.Time        `codec:"after,omitempty" json:"after,omitempty"`
	OneChatTypePerTLF *bool                `codec:"oneChatTypePerTLF,omitempty" json:"oneChatTypePerTLF,omitempty"`
	Status            []ConversationStatus `codec:"status" json:"status"`
	UnreadOnly        bool                 `codec:"unreadOnly" json:"unreadOnly"`
	ReadOnly          bool                 `codec:"readOnly" json:"readOnly"`
	ComputeActiveList bool                 `codec:"computeActiveList" json:"computeActiveList"`
}

type ConversationIDTriple struct {
	Tlfid     TLFID     `codec:"tlfid" json:"tlfid"`
	TopicType TopicType `codec:"topicType" json:"topicType"`
	TopicID   TopicID   `codec:"topicID" json:"topicID"`
}

type ConversationFinalizeInfo struct {
	ResetUser      string       `codec:"resetUser" json:"resetUser"`
	ResetDate      string       `codec:"resetDate" json:"resetDate"`
	ResetFull      string       `codec:"resetFull" json:"resetFull"`
	ResetTimestamp gregor1.Time `codec:"resetTimestamp" json:"resetTimestamp"`
}

type ConversationResolveInfo struct {
	NewTLFName string `codec:"newTLFName" json:"newTLFName"`
}

type ConversationMetadata struct {
	IdTriple       ConversationIDTriple      `codec:"idTriple" json:"idTriple"`
	ConversationID ConversationID            `codec:"conversationID" json:"conversationID"`
	Visibility     TLFVisibility             `codec:"visibility" json:"visibility"`
	Status         ConversationStatus        `codec:"status" json:"status"`
	FinalizeInfo   *ConversationFinalizeInfo `codec:"finalizeInfo,omitempty" json:"finalizeInfo,omitempty"`
	Supersedes     []ConversationMetadata    `codec:"supersedes" json:"supersedes"`
	SupersededBy   []ConversationMetadata    `codec:"supersededBy" json:"supersededBy"`
	ActiveList     []gregor1.UID             `codec:"activeList" json:"activeList"`
}

type ConversationReaderInfo struct {
	Mtime     gregor1.Time `codec:"mtime" json:"mtime"`
	ReadMsgid MessageID    `codec:"readMsgid" json:"readMsgid"`
	MaxMsgid  MessageID    `codec:"maxMsgid" json:"maxMsgid"`
}

type Conversation struct {
	Metadata        ConversationMetadata    `codec:"metadata" json:"metadata"`
	ReaderInfo      *ConversationReaderInfo `codec:"readerInfo,omitempty" json:"readerInfo,omitempty"`
	MaxMsgSummaries []MessageSummary        `codec:"maxMsgSummaries" json:"maxMsgSummaries"`
}

type MessageSummary struct {
	MsgID       MessageID    `codec:"msgID" json:"msgID"`
	MessageType MessageType  `codec:"messageType" json:"messageType"`
	TlfName     string       `codec:"tlfName" json:"tlfName"`
	TlfPublic   bool         `codec:"tlfPublic" json:"tlfPublic"`
	Ctime       gregor1.Time `codec:"ctime" json:"ctime"`
}

type MessageServerHeader struct {
	MessageID    MessageID    `codec:"messageID" json:"messageID"`
	SupersededBy MessageID    `codec:"supersededBy" json:"supersededBy"`
	Ctime        gregor1.Time `codec:"ctime" json:"ctime"`
}

type MessagePreviousPointer struct {
	Id   MessageID `codec:"id" json:"id"`
	Hash Hash      `codec:"hash" json:"hash"`
}

type OutboxInfo struct {
	Prev        MessageID    `codec:"prev" json:"prev"`
	ComposeTime gregor1.Time `codec:"composeTime" json:"composeTime"`
}

type MessageClientHeader struct {
	Conv         ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName      string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic    bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType  MessageType              `codec:"messageType" json:"messageType"`
	Supersedes   MessageID                `codec:"supersedes" json:"supersedes"`
	Deletes      []MessageID              `codec:"deletes" json:"deletes"`
	Prev         []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender       gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	MerkleRoot   *MerkleRoot              `codec:"merkleRoot,omitempty" json:"merkleRoot,omitempty"`
	OutboxID     *OutboxID                `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	OutboxInfo   *OutboxInfo              `codec:"outboxInfo,omitempty" json:"outboxInfo,omitempty"`
}

type MessageClientHeaderVerified struct {
	Conv         ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName      string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic    bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType  MessageType              `codec:"messageType" json:"messageType"`
	Prev         []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender       gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	OutboxID     *OutboxID                `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	OutboxInfo   *OutboxInfo              `codec:"outboxInfo,omitempty" json:"outboxInfo,omitempty"`
}

type EncryptedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

type SignEncryptedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

type SealedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

type SignatureInfo struct {
	V int    `codec:"v" json:"v"`
	S []byte `codec:"s" json:"s"`
	K []byte `codec:"k" json:"k"`
}

type MerkleRoot struct {
	Seqno int64  `codec:"seqno" json:"seqno"`
	Hash  []byte `codec:"hash" json:"hash"`
}

type InboxResType int

const (
	InboxResType_VERSIONHIT InboxResType = 0
	InboxResType_FULL       InboxResType = 1
)

var InboxResTypeMap = map[string]InboxResType{
	"VERSIONHIT": 0,
	"FULL":       1,
}

var InboxResTypeRevMap = map[InboxResType]string{
	0: "VERSIONHIT",
	1: "FULL",
}

func (e InboxResType) String() string {
	if v, ok := InboxResTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type InboxViewFull struct {
	Vers          InboxVers      `codec:"vers" json:"vers"`
	Conversations []Conversation `codec:"conversations" json:"conversations"`
	Pagination    *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type InboxView struct {
	Rtype__ InboxResType   `codec:"rtype" json:"rtype"`
	Full__  *InboxViewFull `codec:"full,omitempty" json:"full,omitempty"`
}

func (o *InboxView) Rtype() (ret InboxResType, err error) {
	switch o.Rtype__ {
	case InboxResType_FULL:
		if o.Full__ == nil {
			err = errors.New("unexpected nil value for Full__")
			return ret, err
		}
	}
	return o.Rtype__, nil
}

func (o InboxView) Full() InboxViewFull {
	if o.Rtype__ != InboxResType_FULL {
		panic("wrong case accessed")
	}
	if o.Full__ == nil {
		return InboxViewFull{}
	}
	return *o.Full__
}

func NewInboxViewWithVersionhit() InboxView {
	return InboxView{
		Rtype__: InboxResType_VERSIONHIT,
	}
}

func NewInboxViewWithFull(v InboxViewFull) InboxView {
	return InboxView{
		Rtype__: InboxResType_FULL,
		Full__:  &v,
	}
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
