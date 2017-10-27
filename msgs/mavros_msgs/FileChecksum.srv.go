// Code generated by ros-gen-go.
// source: FileChecksum.srv
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

// Service type metadata
type _SrvFileChecksum struct {
	name    string
	md5sum  string
	text    string
	reqType ros.MessageType
	resType ros.MessageType
}

func (t *_SrvFileChecksum) Name() string                  { return t.name }
func (t *_SrvFileChecksum) MD5Sum() string                { return t.md5sum }
func (t *_SrvFileChecksum) Text() string                  { return t.text }
func (t *_SrvFileChecksum) RequestType() ros.MessageType  { return t.reqType }
func (t *_SrvFileChecksum) ResponseType() ros.MessageType { return t.resType }
func (t *_SrvFileChecksum) NewService() ros.Service {
	return new(FileChecksum)
}

var (
	SrvFileChecksum = &_SrvFileChecksum{
		"mavros_msgs/FileChecksum",
		"c32158e17d9a1a1c682fe19adbd29fac",
		`# FTP::Checksum
#
# :file_path:	file to calculate checksum
# :crc32:	file checksum
# :success:	indicates success end of request
# :r_errno:	remote errno if applicapable

string file_path
---
uint32 crc32
bool success
int32 r_errno
`,
		MsgFileChecksumRequest,
		MsgFileChecksumResponse,
	}
)

type FileChecksum struct {
	Request  FileChecksumRequest
	Response FileChecksumResponse
}

func (s *FileChecksum) ReqMessage() ros.Message { return &s.Request }
func (s *FileChecksum) ResMessage() ros.Message { return &s.Response }

// FileChecksumRequest

type _MsgFileChecksumRequest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFileChecksumRequest) Text() string {
	return t.text
}

func (t *_MsgFileChecksumRequest) Name() string {
	return t.name
}

func (t *_MsgFileChecksumRequest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFileChecksumRequest) NewMessage() ros.Message {
	m := new(FileChecksumRequest)

	return m
}

var (
	MsgFileChecksumRequest = &_MsgFileChecksumRequest{
		`# FTP::Checksum
#
# :file_path:	file to calculate checksum
# :crc32:	file checksum
# :success:	indicates success end of request
# :r_errno:	remote errno if applicapable

string file_path
`,
		"mavros_msgs/FileChecksumRequest",
		"",
	}
)

type FileChecksumRequest struct {
	FilePath string
}

func (m *FileChecksumRequest) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.FilePath); err != nil {
		return err
	}

	return
}

func (m *FileChecksumRequest) Deserialize(r io.Reader) (err error) {
	// FilePath
	if err = ros.DeserializeMessageField(r, "string", &m.FilePath); err != nil {
		return err
	}

	return
}

// FileChecksumResponse

type _MsgFileChecksumResponse struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFileChecksumResponse) Text() string {
	return t.text
}

func (t *_MsgFileChecksumResponse) Name() string {
	return t.name
}

func (t *_MsgFileChecksumResponse) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFileChecksumResponse) NewMessage() ros.Message {
	m := new(FileChecksumResponse)

	return m
}

var (
	MsgFileChecksumResponse = &_MsgFileChecksumResponse{
		`
uint32 crc32
bool success
int32 r_errno
`,
		"mavros_msgs/FileChecksumResponse",
		"",
	}
)

type FileChecksumResponse struct {
	Crc32   uint32
	Success bool
	RErrno  int32
}

func (m *FileChecksumResponse) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint32", &m.Crc32); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.Success); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "int32", &m.RErrno); err != nil {
		return err
	}

	return
}

func (m *FileChecksumResponse) Deserialize(r io.Reader) (err error) {
	// Crc32
	if err = ros.DeserializeMessageField(r, "uint32", &m.Crc32); err != nil {
		return err
	}

	// Success
	if err = ros.DeserializeMessageField(r, "bool", &m.Success); err != nil {
		return err
	}

	// RErrno
	if err = ros.DeserializeMessageField(r, "int32", &m.RErrno); err != nil {
		return err
	}

	return
}