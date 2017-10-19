// Code generated by ros-gen-go.
// source: Empty.srv
// DO NOT EDIT!
package std_srvs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvEmpty struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvEmpty) Name() string                  { return t.name }
func (t *_SrvEmpty) MD5Sum() string                { return t.md5sum }
func (t *_SrvEmpty) Text() string                  { return t.text }
func (t *_SrvEmpty) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvEmpty) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvEmpty) NewService() ros.Service {
	return new(Empty)
}

var (
	SrvEmpty = &_SrvEmpty{
		"std_srvs/Empty",
		"d41d8cd98f00b204e9800998ecf8427e",
		`---`,
		MsgEmptyRequest,
		MsgEmptyResponse,
	}
)

type Empty struct {
	Request  EmptyRequest
	Response EmptyResponse
}

func (s *Empty) ReqMessage() ros.Message { return &s.Request }
func (s *Empty) ResMessage() ros.Message { return &s.Response }

// EmptyRequest

type _MsgEmptyRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgEmptyRequest) Text() string {
	return t.text
}

func (t *_MsgEmptyRequest) Name() string {
	return t.name
}

func (t *_MsgEmptyRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgEmptyRequest) NewMessage() ros.Message {
	m := new(EmptyRequest)

	return m
}

var (
	MsgEmptyRequest = &_MsgEmptyRequest{
		``,
		"std_srvs/EmptyRequest",
		"",
	}
)

type EmptyRequest struct {
}

func (m *EmptyRequest) Serialize(w io.Writer) (err error) {
	return
}

func (m *EmptyRequest) Deserialize(r io.Reader) (err error) {
	return
}

// EmptyResponse

type _MsgEmptyResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgEmptyResponse) Text() string {
	return t.text
}

func (t *_MsgEmptyResponse) Name() string {
	return t.name
}

func (t *_MsgEmptyResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgEmptyResponse) NewMessage() ros.Message {
	m := new(EmptyResponse)

	return m
}

var (
	MsgEmptyResponse = &_MsgEmptyResponse{
		``,
		"std_srvs/EmptyResponse",
		"",
	}
)

type EmptyResponse struct {
}

func (m *EmptyResponse) Serialize(w io.Writer) (err error) {
	return
}

func (m *EmptyResponse) Deserialize(r io.Reader) (err error) {
	return
}
