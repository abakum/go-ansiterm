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

	//encode from utf8 to cp
	wInUTF8.Write(title.Bytes())
	wInUTF8.Write(msg.Bytes())
	wInUTF8.Close()

	//write in utf8
	b.Write(title.Bytes())

	fmt.Printf("%#v\n", b)
	fmt.Println(b.String())

	o, err := ansiterm.Strip(b, ansiterm.WithFe(true))
	fmt.Println(o, err)

	// decode from cp to utf8
	rInUTF8 := transform.NewReader(strings.NewReader(o), d)
	decBytes, _ := io.ReadAll(rInUTF8)
	fmt.Println(string(decBytes))
}
