package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	ansiterm "github.com/abakum/go-ansiterm"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var o bytes.Buffer

func main() {

	var (
		b, title, msg bytes.Buffer
		e             *encoding.Encoder
		d             *encoding.Decoder
	)
	switch 1 {
	case 1:
		cp := charmap.CodePage866
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	case 2:
		cp := charmap.Windows1251
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	case 3:
		cp := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	}
	t := "Это заголовок окна"
	// t = "This is the window title"
	title.WriteByte(ansiterm.ANSI_ESCAPE_PRIMARY)
	title.WriteByte(ansiterm.ANSI_OSC_STRING_ENTRY)
	title.WriteByte('0')
	title.WriteByte(ansiterm.ANSI_PARAMETER_SEP[0])
	title.WriteString(t)
	title.WriteByte(ansiterm.ANSI_BEL)

	title.WriteByte(ansiterm.ANSI_ESCAPE_PRIMARY)
	title.WriteByte(ansiterm.ANSI_OSC_STRING_ENTRY)
	title.WriteByte('0')
	title.WriteByte(ansiterm.ANSI_PARAMETER_SEP[0])
	title.WriteString(t)
	title.WriteByte(ansiterm.ANSI_ESCAPE_PRIMARY)
	title.WriteByte(ansiterm.ANSI_CMD_STR_TERM)

	s := "Для продолжения нажмите любую клавишу . . ."
	// s = "Press any key to continue . . ."
	msg.WriteString(s)

	fmt.Println(msg.String())

	wInUTF8 := transform.NewWriter(&b, e)

	//encode from utf8
	wInUTF8.Write(title.Bytes())
	wInUTF8.Write(msg.Bytes())
	wInUTF8.Close()

	//write in utf8
	b.Write(title.Bytes())

	fmt.Printf("%#v\n", b)
	fmt.Println(b.String())

	parser := ansiterm.CreateParser("Ground", &handler{})
	i, err := parser.Parse(b.Bytes())
	fmt.Println()
	fmt.Println(i, err)

	rInUTF8 := transform.NewReader(strings.NewReader(o.String()), d)
	decBytes, _ := io.ReadAll(rInUTF8)
	fmt.Println()
	fmt.Println(string(decBytes))

}

type handler struct{}

func (h *handler) Print(b byte) error {
	fmt.Printf("p%#x ", b)
	o.WriteByte(b)
	return nil
}

func (h *handler) Execute(b byte) error {
	fmt.Printf("e%#x ", b)
	return nil
}

func (h *handler) CUU(int) error {
	panic("not implemented")
}

func (h *handler) CUD(int) error {
	panic("not implemented")
}

func (h *handler) CUF(int) error {
	panic("not implemented")
}

func (h *handler) CUB(int) error {
	panic("not implemented")
}

func (h *handler) CNL(int) error {
	panic("not implemented")
}

func (h *handler) CPL(int) error {
	panic("not implemented")
}

func (h *handler) CHA(int) error {
	panic("not implemented")
}

func (h *handler) VPA(int) error {
	panic("not implemented")
}

func (h *handler) CUP(int, int) error {
	panic("not implemented")
}

func (h *handler) HVP(int, int) error {
	panic("not implemented")
}

func (h *handler) DECTCEM(bool) error {
	panic("not implemented")
}

func (h *handler) DECOM(bool) error {
	panic("not implemented")
}

func (h *handler) DECCOLM(bool) error {
	panic("not implemented")
}

func (h *handler) ED(int) error {
	panic("not implemented")
}

func (h *handler) EL(int) error {
	panic("not implemented")
}

func (h *handler) IL(int) error {
	panic("not implemented")
}

func (h *handler) DL(int) error {
	panic("not implemented")
}

func (h *handler) ICH(int) error {
	panic("not implemented")
}

func (h *handler) DCH(int) error {
	panic("not implemented")
}

func (h *handler) SGR([]int) error {
	panic("not implemented")
}

func (h *handler) SU(int) error {
	panic("not implemented")
}

func (h *handler) SD(int) error {
	panic("not implemented")
}

func (h *handler) DA([]string) error {
	panic("not implemented")
}

func (h *handler) DECSTBM(int, int) error {
	panic("not implemented")
}

func (h *handler) IND() error {
	panic("not implemented")
}

func (h *handler) RI() error {
	panic("not implemented")
}

func (h *handler) Flush() error {
	return nil
}
