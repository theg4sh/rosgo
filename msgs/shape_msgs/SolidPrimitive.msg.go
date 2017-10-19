// Code generated by ros-gen-go.
// source: SolidPrimitive.msg
// DO NOT EDIT!
package shape_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cnord/rosgo/ros"
)

type _MsgSolidPrimitive struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSolidPrimitive) Text() string {
	return t.text
}

func (t *_MsgSolidPrimitive) Name() string {
	return t.name
}

func (t *_MsgSolidPrimitive) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSolidPrimitive) NewMessage() ros.Message {
	m := new(SolidPrimitive)

	return m
}

var (
	MsgSolidPrimitive = &_MsgSolidPrimitive{
		`# Define box, sphere, cylinder, cone 
# All shapes are defined to have their bounding boxes centered around 0,0,0.

uint8 BOX=1
uint8 SPHERE=2
uint8 CYLINDER=3
uint8 CONE=4

# The type of the shape
uint8 type


# The dimensions of the shape
float64[] dimensions

# The meaning of the shape dimensions: each constant defines the index in the 'dimensions' array

# For the BOX type, the X, Y, and Z dimensions are the length of the corresponding
# sides of the box.
uint8 BOX_X=0
uint8 BOX_Y=1
uint8 BOX_Z=2


# For the SPHERE type, only one component is used, and it gives the radius of
# the sphere.
uint8 SPHERE_RADIUS=0


# For the CYLINDER and CONE types, the center line is oriented along
# the Z axis.  Therefore the CYLINDER_HEIGHT (CONE_HEIGHT) component
# of dimensions gives the height of the cylinder (cone).  The
# CYLINDER_RADIUS (CONE_RADIUS) component of dimensions gives the
# radius of the base of the cylinder (cone).  Cone and cylinder
# primitives are defined to be circular. The tip of the cone is
# pointing up, along +Z axis.

uint8 CYLINDER_HEIGHT=0
uint8 CYLINDER_RADIUS=1

uint8 CONE_HEIGHT=0
uint8 CONE_RADIUS=1
`,
		"shape_msgs/SolidPrimitive",
		"d8f8cbc74c5ff283fca29569ccefb45d",
	}
)

type SolidPrimitive struct {
	Type       uint8
	Dimensions []float64
}

func (m *SolidPrimitive) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "uint8", &m.Type); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Dimensions)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Dimensions {
		if err = ros.SerializeMessageField(w, "float64", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *SolidPrimitive) Deserialize(r io.Reader) (err error) {
	// Type
	if err = ros.DeserializeMessageField(r, "uint8", &m.Type); err != nil {
		return err
	}

	// Dimensions
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Dimensions: %s", err)
		}
		m.Dimensions = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "float64", &m.Dimensions[i]); err != nil {
				return err
			}
		}
	}

	return
}
