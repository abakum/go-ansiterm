package ansiterm

type oscStringState struct {
	baseState
}

var escapeInOscString bool

func (oscState oscStringState) Handle(b byte) (s state, e error) {
	oscState.parser.logf("OscString::Handle %#x", b)
	defer func() {
		escapeInOscString = b == ANSI_ESCAPE_PRIMARY
	}()

	// nextState, err := oscState.baseState.Handle(b)
	// if nextState != nil || err != nil {
	// 	return nextState, err
	// }

	switch {
	case isOscStringTerminator(b, escapeInOscString):
		return oscState.parser.ground, nil
	}

	return oscState, nil
}

// See below for OSC string terminators for linux
// http://man7.org/linux/man-pages/man4/console_codes.4.html
func isOscStringTerminator(b byte, wasEscape bool) bool {
	//  []byte{ANSI_ESCAPE_PRIMARY, ANSI_CMD_STR_TERM}
	if b == ANSI_BEL || (wasEscape && b == ANSI_CMD_STR_TERM) {
		return true
	}

	return false
}
