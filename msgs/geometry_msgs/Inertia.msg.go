// Code generated by ros-gen-go.
// source: Inertia.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgInertia struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInertia) Text() string {
	return t.text
}

func (t *_MsgInertia) Name() string {
	return t.name
}

func (t *_MsgInertia) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInertia) NewMessage() ros.Message {
	m := new(Inertia)

	return m
}

var (
	MsgInertia = &_MsgInertia{
		`# Mass [kg]
float64 m

# Center of mass [m]
geometry_msgs/Vector3 com

# Inertia Tensor [kg-m^2]
#     | ixx ixy ixz |
# I = | ixy iyy iyz |
#     | ixz iyz izz |
float64 ixx
float64 ixy
float64 ixz
float64 iyy
float64 iyz
float64 izz
`,
		"geometry_msgs/Inertia",
		"1d26e4bb6c83ff141c5cf0d883c2b0fe",
	}
)

type Inertia struct {
	M   float64
	Com Vector3
	Ixx float64
	Ixy float64
	Ixz float64
	Iyy float64
	Iyz float64
	Izz float64
}

func (m *Inertia) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "float64", &m.M); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Vector3", &m.Com); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Ixx); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Ixy); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Ixz); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Iyy); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Iyz); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float64", &m.Izz); err != nil {
		return err
	}

	return
}

func (m *Inertia) Deserialize(r io.Reader) (err error) {
	// M
	if err = ros.DeserializeMessageField(r, "float64", &m.M); err != nil {
		return err
	}

	// Com
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Vector3", &m.Com); err != nil {
		return err
	}

	// Ixx
	if err = ros.DeserializeMessageField(r, "float64", &m.Ixx); err != nil {
		return err
	}

	// Ixy
	if err = ros.DeserializeMessageField(r, "float64", &m.Ixy); err != nil {
		return err
	}

	// Ixz
	if err = ros.DeserializeMessageField(r, "float64", &m.Ixz); err != nil {
		return err
	}

	// Iyy
	if err = ros.DeserializeMessageField(r, "float64", &m.Iyy); err != nil {
		return err
	}

	// Iyz
	if err = ros.DeserializeMessageField(r, "float64", &m.Iyz); err != nil {
		return err
	}

	// Izz
	if err = ros.DeserializeMessageField(r, "float64", &m.Izz); err != nil {
		return err
	}

	return
}
