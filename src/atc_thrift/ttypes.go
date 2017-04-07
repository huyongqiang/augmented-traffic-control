// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package atc_thrift

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type PlatformType int64

const (
	PlatformType_OTHER PlatformType = 0
	PlatformType_LINUX PlatformType = 1
)

func (p PlatformType) String() string {
	switch p {
	case PlatformType_OTHER:
		return "OTHER"
	case PlatformType_LINUX:
		return "LINUX"
	}
	return "<UNSET>"
}

func PlatformTypeFromString(s string) (PlatformType, error) {
	switch s {
	case "OTHER":
		return PlatformType_OTHER, nil
	case "LINUX":
		return PlatformType_LINUX, nil
	}
	return PlatformType(0), fmt.Errorf("not a valid PlatformType string")
}

func PlatformTypePtr(v PlatformType) *PlatformType { return &v }

func (p PlatformType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *PlatformType) UnmarshalText(text []byte) error {
	q, err := PlatformTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - Delay
//  - Jitter
//  - Correlation
type Delay struct {
	Delay       int32   `thrift:"delay,1" json:"delay"`
	Jitter      int32   `thrift:"jitter,2" json:"jitter,omitempty"`
	Correlation float64 `thrift:"correlation,3" json:"correlation,omitempty"`
}

func NewDelay() *Delay {
	return &Delay{}
}

func (p *Delay) GetDelay() int32 {
	return p.Delay
}

var Delay_Jitter_DEFAULT int32 = 0

func (p *Delay) GetJitter() int32 {
	return p.Jitter
}

var Delay_Correlation_DEFAULT float64 = 0

func (p *Delay) GetCorrelation() float64 {
	return p.Correlation
}
func (p *Delay) IsSetJitter() bool {
	return p.Jitter != Delay_Jitter_DEFAULT
}

func (p *Delay) IsSetCorrelation() bool {
	return p.Correlation != Delay_Correlation_DEFAULT
}

func (p *Delay) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Delay) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Delay = v
	}
	return nil
}

func (p *Delay) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Jitter = v
	}
	return nil
}

func (p *Delay) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Correlation = v
	}
	return nil
}

func (p *Delay) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Delay"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Delay) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("delay", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:delay: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Delay)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.delay (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:delay: ", p), err)
	}
	return err
}

func (p *Delay) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetJitter() {
		if err := oprot.WriteFieldBegin("jitter", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:jitter: ", p), err)
		}
		if err := oprot.WriteI32(int32(p.Jitter)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.jitter (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:jitter: ", p), err)
		}
	}
	return err
}

func (p *Delay) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetCorrelation() {
		if err := oprot.WriteFieldBegin("correlation", thrift.DOUBLE, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:correlation: ", p), err)
		}
		if err := oprot.WriteDouble(float64(p.Correlation)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.correlation (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:correlation: ", p), err)
		}
	}
	return err
}

func (p *Delay) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Delay(%+v)", *p)
}

// Attributes:
//  - Percentage
//  - Correlation
type Loss struct {
	Percentage  float64 `thrift:"percentage,1" json:"percentage"`
	Correlation float64 `thrift:"correlation,2" json:"correlation,omitempty"`
}

func NewLoss() *Loss {
	return &Loss{}
}

func (p *Loss) GetPercentage() float64 {
	return p.Percentage
}

var Loss_Correlation_DEFAULT float64 = 0

func (p *Loss) GetCorrelation() float64 {
	return p.Correlation
}
func (p *Loss) IsSetCorrelation() bool {
	return p.Correlation != Loss_Correlation_DEFAULT
}

func (p *Loss) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Loss) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Percentage = v
	}
	return nil
}

func (p *Loss) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Correlation = v
	}
	return nil
}

func (p *Loss) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Loss"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Loss) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("percentage", thrift.DOUBLE, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:percentage: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Percentage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.percentage (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:percentage: ", p), err)
	}
	return err
}

func (p *Loss) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetCorrelation() {
		if err := oprot.WriteFieldBegin("correlation", thrift.DOUBLE, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:correlation: ", p), err)
		}
		if err := oprot.WriteDouble(float64(p.Correlation)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.correlation (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:correlation: ", p), err)
		}
	}
	return err
}

func (p *Loss) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Loss(%+v)", *p)
}

// Attributes:
//  - Percentage
//  - Gap
//  - Correlation
type Reorder struct {
	Percentage  float64 `thrift:"percentage,1" json:"percentage"`
	Gap         int32   `thrift:"gap,2" json:"gap"`
	Correlation float64 `thrift:"correlation,3" json:"correlation,omitempty"`
}

func NewReorder() *Reorder {
	return &Reorder{}
}

func (p *Reorder) GetPercentage() float64 {
	return p.Percentage
}

func (p *Reorder) GetGap() int32 {
	return p.Gap
}

var Reorder_Correlation_DEFAULT float64 = 0

func (p *Reorder) GetCorrelation() float64 {
	return p.Correlation
}
func (p *Reorder) IsSetCorrelation() bool {
	return p.Correlation != Reorder_Correlation_DEFAULT
}

func (p *Reorder) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Reorder) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Percentage = v
	}
	return nil
}

func (p *Reorder) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Gap = v
	}
	return nil
}

func (p *Reorder) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Correlation = v
	}
	return nil
}

func (p *Reorder) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Reorder"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Reorder) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("percentage", thrift.DOUBLE, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:percentage: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Percentage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.percentage (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:percentage: ", p), err)
	}
	return err
}

func (p *Reorder) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("gap", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gap: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Gap)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.gap (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gap: ", p), err)
	}
	return err
}

func (p *Reorder) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetCorrelation() {
		if err := oprot.WriteFieldBegin("correlation", thrift.DOUBLE, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:correlation: ", p), err)
		}
		if err := oprot.WriteDouble(float64(p.Correlation)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.correlation (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:correlation: ", p), err)
		}
	}
	return err
}

func (p *Reorder) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Reorder(%+v)", *p)
}

// Attributes:
//  - Percentage
//  - Correlation
type Corruption struct {
	Percentage  float64 `thrift:"percentage,1" json:"percentage"`
	Correlation float64 `thrift:"correlation,2" json:"correlation,omitempty"`
}

func NewCorruption() *Corruption {
	return &Corruption{}
}

func (p *Corruption) GetPercentage() float64 {
	return p.Percentage
}

var Corruption_Correlation_DEFAULT float64 = 0

func (p *Corruption) GetCorrelation() float64 {
	return p.Correlation
}
func (p *Corruption) IsSetCorrelation() bool {
	return p.Correlation != Corruption_Correlation_DEFAULT
}

func (p *Corruption) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Corruption) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Percentage = v
	}
	return nil
}

func (p *Corruption) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Correlation = v
	}
	return nil
}

func (p *Corruption) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Corruption"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Corruption) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("percentage", thrift.DOUBLE, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:percentage: ", p), err)
	}
	if err := oprot.WriteDouble(float64(p.Percentage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.percentage (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:percentage: ", p), err)
	}
	return err
}

func (p *Corruption) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetCorrelation() {
		if err := oprot.WriteFieldBegin("correlation", thrift.DOUBLE, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:correlation: ", p), err)
		}
		if err := oprot.WriteDouble(float64(p.Correlation)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.correlation (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:correlation: ", p), err)
		}
	}
	return err
}

func (p *Corruption) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Corruption(%+v)", *p)
}

// Attributes:
//  - Rate
//  - Delay
//  - Loss
//  - Reorder
//  - Corruption
type LinkShaping struct {
	Rate       int32       `thrift:"rate,1" json:"rate"`
	Delay      *Delay      `thrift:"delay,2" json:"delay,omitempty"`
	Loss       *Loss       `thrift:"loss,3" json:"loss,omitempty"`
	Reorder    *Reorder    `thrift:"reorder,4" json:"reorder,omitempty"`
	Corruption *Corruption `thrift:"corruption,5" json:"corruption,omitempty"`
}

func NewLinkShaping() *LinkShaping {
	return &LinkShaping{}
}

func (p *LinkShaping) GetRate() int32 {
	return p.Rate
}

var LinkShaping_Delay_DEFAULT *Delay = &Delay{
	Delay: 0,
}

func (p *LinkShaping) GetDelay() *Delay {
	if !p.IsSetDelay() {
		return LinkShaping_Delay_DEFAULT
	}
	return p.Delay
}

var LinkShaping_Loss_DEFAULT *Loss = &Loss{
	Percentage: 0,
}

func (p *LinkShaping) GetLoss() *Loss {
	if !p.IsSetLoss() {
		return LinkShaping_Loss_DEFAULT
	}
	return p.Loss
}

var LinkShaping_Reorder_DEFAULT *Reorder = &Reorder{
	Percentage: 0,
}

func (p *LinkShaping) GetReorder() *Reorder {
	if !p.IsSetReorder() {
		return LinkShaping_Reorder_DEFAULT
	}
	return p.Reorder
}

var LinkShaping_Corruption_DEFAULT *Corruption = &Corruption{
	Percentage: 0,
}

func (p *LinkShaping) GetCorruption() *Corruption {
	if !p.IsSetCorruption() {
		return LinkShaping_Corruption_DEFAULT
	}
	return p.Corruption
}
func (p *LinkShaping) IsSetDelay() bool {
	return p.Delay != nil
}

func (p *LinkShaping) IsSetLoss() bool {
	return p.Loss != nil
}

func (p *LinkShaping) IsSetReorder() bool {
	return p.Reorder != nil
}

func (p *LinkShaping) IsSetCorruption() bool {
	return p.Corruption != nil
}

func (p *LinkShaping) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LinkShaping) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Rate = v
	}
	return nil
}

func (p *LinkShaping) readField2(iprot thrift.TProtocol) error {
	p.Delay = &Delay{}
	if err := p.Delay.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Delay), err)
	}
	return nil
}

func (p *LinkShaping) readField3(iprot thrift.TProtocol) error {
	p.Loss = &Loss{}
	if err := p.Loss.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Loss), err)
	}
	return nil
}

func (p *LinkShaping) readField4(iprot thrift.TProtocol) error {
	p.Reorder = &Reorder{}
	if err := p.Reorder.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Reorder), err)
	}
	return nil
}

func (p *LinkShaping) readField5(iprot thrift.TProtocol) error {
	p.Corruption = &Corruption{}
	if err := p.Corruption.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Corruption), err)
	}
	return nil
}

func (p *LinkShaping) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LinkShaping"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LinkShaping) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("rate", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:rate: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Rate)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.rate (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:rate: ", p), err)
	}
	return err
}

func (p *LinkShaping) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDelay() {
		if err := oprot.WriteFieldBegin("delay", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:delay: ", p), err)
		}
		if err := p.Delay.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Delay), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:delay: ", p), err)
		}
	}
	return err
}

func (p *LinkShaping) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetLoss() {
		if err := oprot.WriteFieldBegin("loss", thrift.STRUCT, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:loss: ", p), err)
		}
		if err := p.Loss.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Loss), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:loss: ", p), err)
		}
	}
	return err
}

func (p *LinkShaping) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetReorder() {
		if err := oprot.WriteFieldBegin("reorder", thrift.STRUCT, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:reorder: ", p), err)
		}
		if err := p.Reorder.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Reorder), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:reorder: ", p), err)
		}
	}
	return err
}

func (p *LinkShaping) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetCorruption() {
		if err := oprot.WriteFieldBegin("corruption", thrift.STRUCT, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:corruption: ", p), err)
		}
		if err := p.Corruption.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Corruption), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:corruption: ", p), err)
		}
	}
	return err
}

func (p *LinkShaping) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LinkShaping(%+v)", *p)
}

// Attributes:
//  - Up
//  - Down
type Shaping struct {
	Up   *LinkShaping `thrift:"up,1" json:"up"`
	Down *LinkShaping `thrift:"down,2" json:"down"`
}

func NewShaping() *Shaping {
	return &Shaping{}
}

var Shaping_Up_DEFAULT *LinkShaping

func (p *Shaping) GetUp() *LinkShaping {
	if !p.IsSetUp() {
		return Shaping_Up_DEFAULT
	}
	return p.Up
}

var Shaping_Down_DEFAULT *LinkShaping

func (p *Shaping) GetDown() *LinkShaping {
	if !p.IsSetDown() {
		return Shaping_Down_DEFAULT
	}
	return p.Down
}
func (p *Shaping) IsSetUp() bool {
	return p.Up != nil
}

func (p *Shaping) IsSetDown() bool {
	return p.Down != nil
}

func (p *Shaping) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Shaping) readField1(iprot thrift.TProtocol) error {
	p.Up = &LinkShaping{}
	if err := p.Up.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Up), err)
	}
	return nil
}

func (p *Shaping) readField2(iprot thrift.TProtocol) error {
	p.Down = &LinkShaping{}
	if err := p.Down.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Down), err)
	}
	return nil
}

func (p *Shaping) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Shaping"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Shaping) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("up", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:up: ", p), err)
	}
	if err := p.Up.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Up), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:up: ", p), err)
	}
	return err
}

func (p *Shaping) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("down", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:down: ", p), err)
	}
	if err := p.Down.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Down), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:down: ", p), err)
	}
	return err
}

func (p *Shaping) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Shaping(%+v)", *p)
}

// Attributes:
//  - Platform
//  - Version
type AtcdInfo struct {
	Platform PlatformType `thrift:"platform,1" json:"platform"`
	Version  string       `thrift:"version,2" json:"version"`
}

func NewAtcdInfo() *AtcdInfo {
	return &AtcdInfo{}
}

func (p *AtcdInfo) GetPlatform() PlatformType {
	return p.Platform
}

func (p *AtcdInfo) GetVersion() string {
	return p.Version
}
func (p *AtcdInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AtcdInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := PlatformType(v)
		p.Platform = temp
	}
	return nil
}

func (p *AtcdInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Version = v
	}
	return nil
}

func (p *AtcdInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AtcdInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AtcdInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("platform", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:platform: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Platform)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.platform (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:platform: ", p), err)
	}
	return err
}

func (p *AtcdInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("version", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:version: ", p), err)
	}
	if err := oprot.WriteString(string(p.Version)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.version (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:version: ", p), err)
	}
	return err
}

func (p *AtcdInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AtcdInfo(%+v)", *p)
}

// Attributes:
//  - ID
//  - Members
//  - Shaping
type ShapingGroup struct {
	ID      int64    `thrift:"id,1" json:"id"`
	Members []string `thrift:"members,2" json:"members"`
	Shaping *Shaping `thrift:"shaping,3" json:"shaping,omitempty"`
}

func NewShapingGroup() *ShapingGroup {
	return &ShapingGroup{}
}

func (p *ShapingGroup) GetID() int64 {
	return p.ID
}

func (p *ShapingGroup) GetMembers() []string {
	return p.Members
}

var ShapingGroup_Shaping_DEFAULT *Shaping

func (p *ShapingGroup) GetShaping() *Shaping {
	if !p.IsSetShaping() {
		return ShapingGroup_Shaping_DEFAULT
	}
	return p.Shaping
}
func (p *ShapingGroup) IsSetShaping() bool {
	return p.Shaping != nil
}

func (p *ShapingGroup) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ShapingGroup) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *ShapingGroup) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Members = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Members = append(p.Members, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ShapingGroup) readField3(iprot thrift.TProtocol) error {
	p.Shaping = &Shaping{}
	if err := p.Shaping.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Shaping), err)
	}
	return nil
}

func (p *ShapingGroup) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ShapingGroup"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ShapingGroup) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *ShapingGroup) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("members", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:members: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Members)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Members {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:members: ", p), err)
	}
	return err
}

func (p *ShapingGroup) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetShaping() {
		if err := oprot.WriteFieldBegin("shaping", thrift.STRUCT, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:shaping: ", p), err)
		}
		if err := p.Shaping.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Shaping), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:shaping: ", p), err)
		}
	}
	return err
}

func (p *ShapingGroup) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ShapingGroup(%+v)", *p)
}
