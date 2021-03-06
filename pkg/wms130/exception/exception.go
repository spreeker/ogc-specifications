package exception

import (
	"encoding/xml"
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

// Type and Version as constant
const (
	Service string = `WMS`
	Version string = `1.3.0`
)

// WMSServiceExceptionReport struct
// TODO exception restructuring
type WMSServiceExceptionReport struct {
	XMLName          xml.Name        `xml:"ServiceExceptionReport" yaml:"serviceexceptionreport"`
	Version          string          `xml:"version,attr" yaml:"version"`
	Xmlns            string          `xml:"xmlns,attr,omitempty"`
	Xsi              string          `xml:"xsi,attr,omitempty"`
	SchemaLocation   string          `xml:"schemaLocation,attr,omitempty"`
	ServiceException []ows.Exception `xml:"ServiceException"`
}

// Report returns WMSServiceExceptionReport
func (r WMSServiceExceptionReport) Report(errors []ows.Exception) []byte {
	r.Version = Version
	r.Xmlns = `http://www.opengis.net/ogc`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.SchemaLocation = `http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd`

	r.ServiceException = errors
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// WMSException grouping the error message variables together
type WMSException struct {
	ExceptionText string `xml:",chardata" yaml:"exception"`
	ExceptionCode string `xml:"code,attr" yaml:"code"`
	LocatorCode   string `xml:"locator,attr,omitempty" yaml:"locator,omitempty"`
}

// Error returns available ExceptionText
func (e WMSException) Error() string {
	return e.ExceptionText
}

// Code returns available ExceptionCode
func (e WMSException) Code() string {
	return e.ExceptionCode
}

// Locator returns available LocatorCode
func (e WMSException) Locator() string {
	return e.LocatorCode
}

// InvalidFormat exception
func InvalidFormat(unknownformat string) WMSException {
	return WMSException{
		ExceptionText: fmt.Sprintf("The format: %s, is a invalid image format", unknownformat),
		ExceptionCode: `InvalidFormat`,
	}
}

// InvalidCRS exception
func InvalidCRS(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("CRS is not known by this service: %s", s[0]),
			ExceptionCode: `InvalidCRS`,
		}
	}
	if len(s) == 2 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The CRS: %s is not known by the layer: %s", s[0], s[1]),
			ExceptionCode: `InvalidCRS`,
		}
	}
	return WMSException{
		ExceptionCode: `InvalidCRS`,
	}
}

// LayerNotDefined exception
func LayerNotDefined(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The layer: %s is not known by the server", s[0]),
			ExceptionCode: `LayerNotDefined`,
		}
	}
	return WMSException{
		ExceptionCode: `LayerNotDefined`,
	}
}

// StyleNotDefined exception
func StyleNotDefined(s ...string) WMSException {
	if len(s) == 2 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The style: %s is not known by the server for the layer: %s", s[0], s[1]),
			ExceptionCode: `StyleNotDefined`,
		}
	}
	return WMSException{
		ExceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		ExceptionCode: `StyleNotDefined`,
	}
}

// LayerNotQueryable exception
func LayerNotQueryable(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("Layer: %s, can not be queried", s[0]),
			ExceptionCode: `LayerNotQueryable`,
			LocatorCode:   s[0],
		}
	}
	return WMSException{
		ExceptionCode: `LayerNotQueryable`,
	}
}

// InvalidPoint exception
// i and j are strings so we can return none integer values in the exception
func InvalidPoint(i, j string) WMSException {
	// TODO provide giving WIDTH and HEIGHT values in Exception response
	return WMSException{
		ExceptionText: fmt.Sprintf("The parameters I and J are invalid, given: %s for I and %s for J", i, j),
		ExceptionCode: `InvalidPoint`,
	}
}

// CurrentUpdateSequence exception
func CurrentUpdateSequence() WMSException {
	return WMSException{
		ExceptionCode: `CurrentUpdateSequence`,
	}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() WMSException {
	return WMSException{
		ExceptionCode: `InvalidUpdateSequence`,
	}
}

// MissingDimensionValue exception
func MissingDimensionValue() WMSException {
	return WMSException{
		ExceptionCode: `MissingDimensionValue`,
	}
}

// InvalidDimensionValue exception
func InvalidDimensionValue() WMSException {
	return WMSException{
		ExceptionCode: `InvalidDimensionValue`,
	}
}
