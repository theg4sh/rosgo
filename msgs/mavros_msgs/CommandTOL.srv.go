// Code generated by ros-gen-go.
// source: CommandTOL.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvCommandTOL struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvCommandTOL) Name() string                  { return t.name }
func (t *_SrvCommandTOL) MD5Sum() string                { return t.md5sum }
func (t *_SrvCommandTOL) Text() string                  { return t.text }
func (t *_SrvCommandTOL) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvCommandTOL) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvCommandTOL) NewService() ros.Service {
	return new(CommandTOL)
}

var (
	SrvCommandTOL = &_SrvCommandTOL{
		"mavros_msgs/CommandTOL",
		"93ff4eaa9907f58c0e7a909cddce23e2",
		`# Common type for Take Off and Landing

float32 min_pitch	# used by takeoff
float32 yaw
float32 latitude
float32 longitude
float32 altitude
---
bool success
uint8 result
`,
		MsgCommandTOLRequest,
		MsgCommandTOLResponse,
	}
)

type CommandTOL struct {
	Request  CommandTOLRequest
	Response CommandTOLResponse
}

func (s *CommandTOL) ReqMessage() ros.Message { return &s.Request }
func (s *CommandTOL) ResMessage() ros.Message { return &s.Response }

// CommandTOLRequest

type _MsgCommandTOLRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandTOLRequest) Text() string {
	return t.text
}

func (t *_MsgCommandTOLRequest) Name() string {
	return t.name
}

func (t *_MsgCommandTOLRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandTOLRequest) NewMessage() ros.Message {
	m := new(CommandTOLRequest)

	return m
}

var (
	MsgCommandTOLRequest = &_MsgCommandTOLRequest{
		`# Common type for Take Off and Landing

float32 min_pitch	# used by takeoff
float32 yaw
float32 latitude
float32 longitude
float32 altitude
`,
		"mavros_msgs/CommandTOLRequest",
		"",
	}
)

type CommandTOLRequest struct {
	MinPitch  float32
	Yaw       float32
	Latitude  float32
	Longitude float32
	Altitude  float32
}

func (m *CommandTOLRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float32", &m.MinPitch); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Yaw); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Latitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Longitude); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Altitude); err != nil {
		return err
	}

	return
}

func (m *CommandTOLRequest) Deserialize(r io.Reader) (err error) {
	// MinPitch
	if err = ros.DeserializeMessageField(r, "float32", &m.MinPitch); err != nil {
		return err
	}

	// Yaw
	if err = ros.DeserializeMessageField(r, "float32", &m.Yaw); err != nil {
		return err
	}

	// Latitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Latitude); err != nil {
		return err
	}

	// Longitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Longitude); err != nil {
		return err
	}

	// Altitude
	if err = ros.DeserializeMessageField(r, "float32", &m.Altitude); err != nil {
		return err
	}

	return
}

// CommandTOLResponse

type _MsgCommandTOLResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCommandTOLResponse) Text() string {
	return t.text
}

func (t *_MsgCommandTOLResponse) Name() string {
	return t.name
}

func (t *_MsgCommandTOLResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCommandTOLResponse) NewMessage() ros.Message {
	m := new(CommandTOLResponse)

	return m
}

var (
	MsgCommandTOLResponse = &_MsgCommandTOLResponse{
		`
bool success
uint8 result
`,
		"mavros_msgs/CommandTOLResponse",
		"",
	}
)

type CommandTOLResponse struct {
	Success bool
	Result  uint8
}

func (m *CommandTOLResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Result); err != nil {
		return err
	}

	return
}

func (m *CommandTOLResponse) Deserialize(r io.Reader) (err error) {
	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// Result
	if err = ros.DeserializeMessageField(r, "uint8", &m.Result); err != nil {
		return err
	}

	return
}