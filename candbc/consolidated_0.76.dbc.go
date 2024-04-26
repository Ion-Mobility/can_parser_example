// Package consolidated076can provides primitives for encoding and decoding consolidated_0.76 CAN messages.
//
// Source: consolidated_0.76.dbc
package consolidated076can

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/candebug"
	"go.einride.tech/can/pkg/canrunner"
	"go.einride.tech/can/pkg/cantext"
	"go.einride.tech/can/pkg/descriptor"
	"go.einride.tech/can/pkg/generated"
	"go.einride.tech/can/pkg/socketcan"
)

// prevent unused imports
var (
	_ = context.Background
	_ = fmt.Print
	_ = net.Dial
	_ = http.Error
	_ = sync.Mutex{}
	_ = time.Now
	_ = socketcan.Dial
	_ = candebug.ServeMessagesHTTP
	_ = canrunner.Run
)

// Generated code. DO NOT EDIT.
// vcu_status_pkt_3Reader provides read access to a vcu_status_pkt_3 message.
type vcu_status_pkt_3Reader interface {
	// display_SOC returns the physical value of the display_SOC signal.
	display_SOC() float64
	// range_km returns the value of the range_km signal.
	range_km() uint8
	// target_charge_hours_rem returns the physical value of the target_charge_hours_rem signal.
	target_charge_hours_rem() float64
	// target_charge_mins_rem returns the physical value of the target_charge_mins_rem signal.
	target_charge_mins_rem() float64
	// target_charge_range_km returns the value of the target_charge_range_km signal.
	target_charge_range_km() uint8
}

// vcu_status_pkt_3Writer provides write access to a vcu_status_pkt_3 message.
type vcu_status_pkt_3Writer interface {
	// CopyFrom copies all values from vcu_status_pkt_3Reader.
	CopyFrom(vcu_status_pkt_3Reader) *vcu_status_pkt_3
	// Setdisplay_SOC sets the physical value of the display_SOC signal.
	Setdisplay_SOC(float64) *vcu_status_pkt_3
	// Setrange_km sets the value of the range_km signal.
	Setrange_km(uint8) *vcu_status_pkt_3
	// Settarget_charge_hours_rem sets the physical value of the target_charge_hours_rem signal.
	Settarget_charge_hours_rem(float64) *vcu_status_pkt_3
	// Settarget_charge_mins_rem sets the physical value of the target_charge_mins_rem signal.
	Settarget_charge_mins_rem(float64) *vcu_status_pkt_3
	// Settarget_charge_range_km sets the value of the target_charge_range_km signal.
	Settarget_charge_range_km(uint8) *vcu_status_pkt_3
}

type vcu_status_pkt_3 struct {
	xxx_display_SOC             uint8
	xxx_range_km                uint8
	xxx_target_charge_hours_rem uint8
	xxx_target_charge_mins_rem  uint8
	xxx_target_charge_range_km  uint8
}

func Newvcu_status_pkt_3() *vcu_status_pkt_3 {
	m := &vcu_status_pkt_3{}
	m.Reset()
	return m
}

func (m *vcu_status_pkt_3) Reset() {
	m.xxx_display_SOC = 0
	m.xxx_range_km = 0
	m.xxx_target_charge_hours_rem = 0
	m.xxx_target_charge_mins_rem = 0
	m.xxx_target_charge_range_km = 0
}

func (m *vcu_status_pkt_3) CopyFrom(o vcu_status_pkt_3Reader) *vcu_status_pkt_3 {
	m.Setdisplay_SOC(o.display_SOC())
	m.xxx_range_km = o.range_km()
	m.Settarget_charge_hours_rem(o.target_charge_hours_rem())
	m.Settarget_charge_mins_rem(o.target_charge_mins_rem())
	m.xxx_target_charge_range_km = o.target_charge_range_km()
	return m
}

// Descriptor returns the vcu_status_pkt_3 descriptor.
func (m *vcu_status_pkt_3) Descriptor() *descriptor.Message {
	return Messages().vcu_status_pkt_3.Message
}

// String returns a compact string representation of the message.
func (m *vcu_status_pkt_3) String() string {
	return cantext.MessageString(m)
}

func (m *vcu_status_pkt_3) display_SOC() float64 {
	return Messages().vcu_status_pkt_3.display_SOC.ToPhysical(float64(m.xxx_display_SOC))
}

func (m *vcu_status_pkt_3) Setdisplay_SOC(v float64) *vcu_status_pkt_3 {
	m.xxx_display_SOC = uint8(Messages().vcu_status_pkt_3.display_SOC.FromPhysical(v))
	return m
}

func (m *vcu_status_pkt_3) range_km() uint8 {
	return m.xxx_range_km
}

func (m *vcu_status_pkt_3) Setrange_km(v uint8) *vcu_status_pkt_3 {
	m.xxx_range_km = uint8(Messages().vcu_status_pkt_3.range_km.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *vcu_status_pkt_3) target_charge_hours_rem() float64 {
	return Messages().vcu_status_pkt_3.target_charge_hours_rem.ToPhysical(float64(m.xxx_target_charge_hours_rem))
}

func (m *vcu_status_pkt_3) Settarget_charge_hours_rem(v float64) *vcu_status_pkt_3 {
	m.xxx_target_charge_hours_rem = uint8(Messages().vcu_status_pkt_3.target_charge_hours_rem.FromPhysical(v))
	return m
}

func (m *vcu_status_pkt_3) target_charge_mins_rem() float64 {
	return Messages().vcu_status_pkt_3.target_charge_mins_rem.ToPhysical(float64(m.xxx_target_charge_mins_rem))
}

func (m *vcu_status_pkt_3) Settarget_charge_mins_rem(v float64) *vcu_status_pkt_3 {
	m.xxx_target_charge_mins_rem = uint8(Messages().vcu_status_pkt_3.target_charge_mins_rem.FromPhysical(v))
	return m
}

func (m *vcu_status_pkt_3) target_charge_range_km() uint8 {
	return m.xxx_target_charge_range_km
}

func (m *vcu_status_pkt_3) Settarget_charge_range_km(v uint8) *vcu_status_pkt_3 {
	m.xxx_target_charge_range_km = uint8(Messages().vcu_status_pkt_3.target_charge_range_km.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *vcu_status_pkt_3) Frame() can.Frame {
	md := Messages().vcu_status_pkt_3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.display_SOC.MarshalUnsigned(&f.Data, uint64(m.xxx_display_SOC))
	md.range_km.MarshalUnsigned(&f.Data, uint64(m.xxx_range_km))
	md.target_charge_hours_rem.MarshalUnsigned(&f.Data, uint64(m.xxx_target_charge_hours_rem))
	md.target_charge_mins_rem.MarshalUnsigned(&f.Data, uint64(m.xxx_target_charge_mins_rem))
	md.target_charge_range_km.MarshalUnsigned(&f.Data, uint64(m.xxx_target_charge_range_km))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *vcu_status_pkt_3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *vcu_status_pkt_3) UnmarshalFrame(f can.Frame) error {
	md := Messages().vcu_status_pkt_3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal vcu_status_pkt_3: expects ID 418247930 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal vcu_status_pkt_3: expects length 5 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal vcu_status_pkt_3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal vcu_status_pkt_3: expects extended ID (got %s with standard ID)", f.String(),
		)
	}
	m.xxx_display_SOC = uint8(md.display_SOC.UnmarshalUnsigned(f.Data))
	m.xxx_range_km = uint8(md.range_km.UnmarshalUnsigned(f.Data))
	m.xxx_target_charge_hours_rem = uint8(md.target_charge_hours_rem.UnmarshalUnsigned(f.Data))
	m.xxx_target_charge_mins_rem = uint8(md.target_charge_mins_rem.UnmarshalUnsigned(f.Data))
	m.xxx_target_charge_range_km = uint8(md.target_charge_range_km.UnmarshalUnsigned(f.Data))
	return nil
}

// Nodes returns the consolidated_0.76 node descriptors.
func Nodes() *NodesDescriptor {
	return nd
}

// NodesDescriptor contains all consolidated_0.76 node descriptors.
type NodesDescriptor struct {
	MC       *descriptor.Node
	MCU      *descriptor.Node
	VCU      *descriptor.Node
	Vector_X *descriptor.Node
	Vector__ *descriptor.Node
}

// Messages returns the consolidated_0.76 message descriptors.
func Messages() *MessagesDescriptor {
	return md
}

// MessagesDescriptor contains all consolidated_0.76 message descriptors.
type MessagesDescriptor struct {
	vcu_status_pkt_3 *vcu_status_pkt_3Descriptor
}

// UnmarshalFrame unmarshals the provided consolidated_0.76 CAN frame.
func (md *MessagesDescriptor) UnmarshalFrame(f can.Frame) (generated.Message, error) {
	switch f.ID {
	case md.vcu_status_pkt_3.ID:
		var msg vcu_status_pkt_3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal consolidated_0.76 frame: %w", err)
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unmarshal consolidated_0.76 frame: ID not in database: %d", f.ID)
	}
}

type vcu_status_pkt_3Descriptor struct {
	*descriptor.Message
	display_SOC             *descriptor.Signal
	range_km                *descriptor.Signal
	target_charge_hours_rem *descriptor.Signal
	target_charge_mins_rem  *descriptor.Signal
	target_charge_range_km  *descriptor.Signal
}

// Database returns the consolidated_0.76 database descriptor.
func (md *MessagesDescriptor) Database() *descriptor.Database {
	return d
}

var nd = &NodesDescriptor{
	MC:       d.Nodes[0],
	MCU:      d.Nodes[1],
	VCU:      d.Nodes[2],
	Vector_X: d.Nodes[3],
	Vector__: d.Nodes[4],
}

var md = &MessagesDescriptor{
	vcu_status_pkt_3: &vcu_status_pkt_3Descriptor{
		Message:                 d.Messages[0],
		display_SOC:             d.Messages[0].Signals[0],
		range_km:                d.Messages[0].Signals[1],
		target_charge_hours_rem: d.Messages[0].Signals[2],
		target_charge_mins_rem:  d.Messages[0].Signals[3],
		target_charge_range_km:  d.Messages[0].Signals[4],
	},
}

var d = (*descriptor.Database)(&descriptor.Database{
	SourceFile: (string)("consolidated_0.76.dbc"),
	Version:    (string)(""),
	Messages: ([]*descriptor.Message)([]*descriptor.Message{
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("vcu_status_pkt_3"),
			ID:          (uint32)(418247930),
			IsExtended:  (bool)(true),
			Length:      (uint8)(5),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("display_SOC"),
					Start:             (uint8)(0),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(100),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("range_km"),
					Start:             (uint8)(8),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(500),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("target_charge_hours_rem"),
					Start:             (uint8)(16),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(23),
					Unit:              (string)("Hours"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("target_charge_mins_rem"),
					Start:             (uint8)(24),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(59),
					Unit:              (string)("Minutes"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("target_charge_range_km"),
					Start:             (uint8)(32),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(false),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(500),
					Unit:              (string)("Km"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("VCU"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
	}),
	Nodes: ([]*descriptor.Node)([]*descriptor.Node{
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("MC"),
			Description: (string)(""),
		}),
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("MCU"),
			Description: (string)(""),
		}),
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("VCU"),
			Description: (string)(""),
		}),
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("Vector_X"),
			Description: (string)(""),
		}),
		(*descriptor.Node)(&descriptor.Node{
			Name:        (string)("Vector__"),
			Description: (string)(""),
		}),
	}),
})
