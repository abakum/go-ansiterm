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
		o             string
		err           error
	)
	crash := 0
	fe := 0
	switch 1 {
	case 0:
		// ECMA-13 ISO 8859-5 chcp 28595
		crash = 0
		fe = 0
		cp := charmap.ISO8859_5
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	case 1:
		crash = 1
		fe = 1
		cp := charmap.CodePage866
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	case 2:
		crash = 0
		fe = 0
		cp := charmap.Windows1251
		e = cp.NewEncoder()
		d = cp.NewDecoder()
	case 3:
		crash = 1
		fe = 1
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
	switch crash {
	case 1:
		b.Write(title.Bytes()) //crash
	}

	fmt.Printf("%#v\n", b)
	fmt.Println(b.String())

	switch fe {
	case 0:
		o, err = ansiterm.StripBuffer(&b)
	case 1:
		o, err = ansiterm.StripBuffer(&b, ansiterm.WithFe(true))
	}
	fmt.Println(o, err)

	b.Write(title.Bytes())
	b.WriteString("123")
	fmt.Println(ansiterm.StripBuffer(&b, ansiterm.WithFe(true)))

	// decode from cp to utf8
	rInUTF8 := transform.NewReader(strings.NewReader(o), d)
	decBytes, _ := io.ReadAll(rInUTF8)
	fmt.Println(string(decBytes))
	fmt.Scanln()
}
