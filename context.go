package ansiterm

type ansiContext struct {
	currentChar  byte
	paramBuffer  []byte
	interBuffer  []byte
	previousChar byte // only for OSC
}
