// Package xevie is the X client API for the XEVIE extension.
package xevie

// This file is automatically generated from xevie.xml. Edit at your peril!

import (
	"github.com/Ptitlu42/xgb"

	"github.com/Ptitlu42/xgb/xproto"
)

// Init must be called before using the XEVIE extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 5, "XEVIE").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XEVIE could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["XEVIE"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["XEVIE"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XEVIE"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["XEVIE"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XEVIE"] = make(map[int]xgb.NewErrorFun)
}

const (
	DatatypeUnmodified = 0
	DatatypeModified   = 1
)

type Event struct {
	// padding: 32 bytes
}

// EventRead reads a byte slice into a Event value.
func EventRead(buf []byte, v *Event) int {
	b := 0

	b += 32 // padding

	return b
}

// EventReadList reads a byte slice into a list of Event values.
func EventReadList(buf []byte, dest []Event) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Event{}
		b += EventRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Event value to a byte slice.
func (v Event) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	b += 32 // padding

	return buf[:b]
}

// EventListBytes writes a list of Event values to a byte slice.
func EventListBytes(buf []byte, list []Event) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// EndCookie is a cookie used only for End requests.
type EndCookie struct {
	*xgb.Cookie
}

// End sends a checked request.
// If an error occurs, it will be returned with the reply by calling EndCookie.Reply()
func End(c *xgb.Conn, Cmap uint32) EndCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'End' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(endRequest(c, Cmap), cookie)
	return EndCookie{cookie}
}

// EndUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func EndUnchecked(c *xgb.Conn, Cmap uint32) EndCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'End' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(endRequest(c, Cmap), cookie)
	return EndCookie{cookie}
}

// EndReply represents the data returned from a End request.
type EndReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Reply blocks and returns the reply data for a End request.
func (cook EndCookie) Reply() (*EndReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return endReply(buf), nil
}

// endReply reads a byte slice into a EndReply value.
func endReply(buf []byte) *EndReply {
	v := new(EndReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return v
}

// Write request to wire for End
// endRequest writes a End request to a byte slice.
func endRequest(c *xgb.Conn, Cmap uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XEVIE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Cmap)
	b += 4

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// padding: 20 bytes
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = xgb.Get16(buf[b:])
	b += 2

	b += 20 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XEVIE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], ClientMajorVersion)
	b += 2

	xgb.Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// SelectInputCookie is a cookie used only for SelectInput requests.
type SelectInputCookie struct {
	*xgb.Cookie
}

// SelectInput sends a checked request.
// If an error occurs, it will be returned with the reply by calling SelectInputCookie.Reply()
func SelectInput(c *xgb.Conn, EventMask uint32) SelectInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(selectInputRequest(c, EventMask), cookie)
	return SelectInputCookie{cookie}
}

// SelectInputUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SelectInputUnchecked(c *xgb.Conn, EventMask uint32) SelectInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(selectInputRequest(c, EventMask), cookie)
	return SelectInputCookie{cookie}
}

// SelectInputReply represents the data returned from a SelectInput request.
type SelectInputReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Reply blocks and returns the reply data for a SelectInput request.
func (cook SelectInputCookie) Reply() (*SelectInputReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return selectInputReply(buf), nil
}

// selectInputReply reads a byte slice into a SelectInputReply value.
func selectInputReply(buf []byte) *SelectInputReply {
	v := new(SelectInputReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return v
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(c *xgb.Conn, EventMask uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XEVIE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], EventMask)
	b += 4

	return buf
}

// SendCookie is a cookie used only for Send requests.
type SendCookie struct {
	*xgb.Cookie
}

// Send sends a checked request.
// If an error occurs, it will be returned with the reply by calling SendCookie.Reply()
func Send(c *xgb.Conn, Event Event, DataType uint32) SendCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'Send' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(sendRequest(c, Event, DataType), cookie)
	return SendCookie{cookie}
}

// SendUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SendUnchecked(c *xgb.Conn, Event Event, DataType uint32) SendCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'Send' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(sendRequest(c, Event, DataType), cookie)
	return SendCookie{cookie}
}

// SendReply represents the data returned from a Send request.
type SendReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Reply blocks and returns the reply data for a Send request.
func (cook SendCookie) Reply() (*SendReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return sendReply(buf), nil
}

// sendReply reads a byte slice into a SendReply value.
func sendReply(buf []byte) *SendReply {
	v := new(SendReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return v
}

// Write request to wire for Send
// sendRequest writes a Send request to a byte slice.
func sendRequest(c *xgb.Conn, Event Event, DataType uint32) []byte {
	size := 104
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XEVIE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	{
		structBytes := Event.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	xgb.Put32(buf[b:], DataType)
	b += 4

	b += 64 // padding

	return buf
}

// StartCookie is a cookie used only for Start requests.
type StartCookie struct {
	*xgb.Cookie
}

// Start sends a checked request.
// If an error occurs, it will be returned with the reply by calling StartCookie.Reply()
func Start(c *xgb.Conn, Screen uint32) StartCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'Start' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(startRequest(c, Screen), cookie)
	return StartCookie{cookie}
}

// StartUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func StartUnchecked(c *xgb.Conn, Screen uint32) StartCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["XEVIE"]; !ok {
		panic("Cannot issue request 'Start' using the uninitialized extension 'XEVIE'. xevie.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(startRequest(c, Screen), cookie)
	return StartCookie{cookie}
}

// StartReply represents the data returned from a Start request.
type StartReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Reply blocks and returns the reply data for a Start request.
func (cook StartCookie) Reply() (*StartReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return startReply(buf), nil
}

// startReply reads a byte slice into a StartReply value.
func startReply(buf []byte) *StartReply {
	v := new(StartReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return v
}

// Write request to wire for Start
// startRequest writes a Start request to a byte slice.
func startRequest(c *xgb.Conn, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["XEVIE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Screen)
	b += 4

	return buf
}
