// Code generated by ros-gen-go.
// source: WaypointSetCurrent.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvWaypointSetCurrent struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvWaypointSetCurrent) Name() string                  { return t.name }
func (t *_SrvWaypointSetCurrent) MD5Sum() string                { return t.md5sum }
func (t *_SrvWaypointSetCurrent) Text() string                  { return t.text }
func (t *_SrvWaypointSetCurrent) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvWaypointSetCurrent) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvWaypointSetCurrent) NewService() ros.Service {
	return new(WaypointSetCurrent)
}

var (
	SrvWaypointSetCurrent = &_SrvWaypointSetCurrent{
		"mavros_msgs/WaypointSetCurrent",
		"f99aa1a911a80ab546ea470f4e90c807",
		`# Request set current waypoint
#
# wp_seq - index in waypoint array

uint16 wp_seq
---
bool success
`,
		MsgWaypointSetCurrentRequest,
		MsgWaypointSetCurrentResponse,
	}
)

type WaypointSetCurrent struct {
	Request  WaypointSetCurrentRequest
	Response WaypointSetCurrentResponse
}

func (s *WaypointSetCurrent) ReqMessage() ros.Message { return &s.Request }
func (s *WaypointSetCurrent) ResMessage() ros.Message { return &s.Response }

// WaypointSetCurrentRequest

type _MsgWaypointSetCurrentRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointSetCurrentRequest) Text() string {
	return t.text
}

func (t *_MsgWaypointSetCurrentRequest) Name() string {
	return t.name
}

func (t *_MsgWaypointSetCurrentRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointSetCurrentRequest) NewMessage() ros.Message {
	m := new(WaypointSetCurrentRequest)

	return m
}

var (
	MsgWaypointSetCurrentRequest = &_MsgWaypointSetCurrentRequest{
		`# Request set current waypoint
#
# wp_seq - index in waypoint array

uint16 wp_seq
`,
		"mavros_msgs/WaypointSetCurrentRequest",
		"",
	}
)

type WaypointSetCurrentRequest struct {
	WpSeq uint16
}

func (m *WaypointSetCurrentRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint16", &m.WpSeq); err != nil {
		return err
	}

	return
}

func (m *WaypointSetCurrentRequest) Deserialize(r io.Reader) (err error) {
	// WpSeq
	if err = ros.DeserializeMessageField(r, "uint16", &m.WpSeq); err != nil {
		return err
	}

	return
}

// WaypointSetCurrentResponse

type _MsgWaypointSetCurrentResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWaypointSetCurrentResponse) Text() string {
	return t.text
}

func (t *_MsgWaypointSetCurrentResponse) Name() string {
	return t.name
}

func (t *_MsgWaypointSetCurrentResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWaypointSetCurrentResponse) NewMessage() ros.Message {
	m := new(WaypointSetCurrentResponse)

	return m
}

var (
	MsgWaypointSetCurrentResponse = &_MsgWaypointSetCurrentResponse{
		`
bool success
`,
		"mavros_msgs/WaypointSetCurrentResponse",
		"",
	}
)

type WaypointSetCurrentResponse struct {
	Success bool
}

func (m *WaypointSetCurrentResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	return
}

func (m *WaypointSetCurrentResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	return
}
