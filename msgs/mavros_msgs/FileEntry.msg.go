// Code generated by ros-gen-go.
// source: FileEntry.msg
// DO NOT EDIT!
package mavros_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgFileEntry struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFileEntry) Text() string {
	return t.text
}

func (t *_MsgFileEntry) Name() string {
	return t.name
}

func (t *_MsgFileEntry) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFileEntry) NewMessage() ros.Message {
	m := new(FileEntry)

	return m
}

var (
	MsgFileEntry = &_MsgFileEntry{
		`# File/Dir information

uint8 TYPE_FILE = 0
uint8 TYPE_DIRECTORY = 1

string name
uint8 type
uint64 size

# Not supported by MAVLink FTP
#time atime
#int32 access_flags
`,
		"mavros_msgs/FileEntry",
		"5ed706bccb946c5b3a5087569cc53ac3",
	}
)

type FileEntry struct {
	Name string
	Type uint8
	Size uint64
}

func (m *FileEntry) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.Type); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint64", &m.Size); err != nil {
		return err
	}

	return
}

func (m *FileEntry) Deserialize(r io.Reader) (err error) {
	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	// Type
	if err = ros.DeserializeMessageField(r, "uint8", &m.Type); err != nil {
		return err
	}

	// Size
	if err = ros.DeserializeMessageField(r, "uint64", &m.Size); err != nil {
		return err
	}

	return
}
