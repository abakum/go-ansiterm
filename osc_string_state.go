package ansiterm

type oscStringState struct {
	baseState
}

func (oscState oscStringState) Handle(b byte) (s state, e error) {
	oscState.parser.logf("OscString::Handle %#x", b)
	defer func() {
		oscState.parser.escapeInOscString = b == ANSI_ESCAPE_PRIMARY
	}()

	if oscState.parser.strictECMA {
		nextState, err := oscState.baseState.Handle(b)
		if nextState != nil || err != nil {
			return nextState, err
		}
	}

	switch {
	case isOscStringTerminator(b, oscState):
		return oscState.parser.ground, nil
	}

	return oscState, nil
}

// See below for OSC string terminators for linux
// http://man7.org/linux/man-pages/man4/console_codes.4.html
func isOscStringTerminator(b byte, oscState oscStringState) bool {
	st := ANSI_BEL
	if oscState.parser.strictECMA {
		st = ANSI_ST
	}

	return (int(b) == st) || (oscState.parser.escapeInOscString && b == ANSI_CMD_STR_TERM)
}

/*
https://vt100.net/emu/dec_ansi_parser
https://ecma-international.org/wp-content/uploads/ECMA-48_5th_edition_june_1991.pdf

8.3.89 OSC - OPERATING SYSTEM COMMAND
Notation: (C1)
Representation: 09/13 or ESC 05/13
OSC is used as the opening delimiter of a control string for operating system use. The command string
following may consist of a sequence of bit combinations in the range 00/08 to 00/13 and 02/00 to 07/14.
The control string is closed by the terminating delimiter STRING TERMINATOR (ST). The
interpretation of the command string depends on the relevant operating system.

8.3.143 ST - STRING TERMINATOR
Notation: (C1)
Representation: 09/12 or ESC 05/12
ST is used as the closing delimiter of a control string opened by APPLICATION PROGRAM
COMMAND (APC), DEVICE CONTROL STRING (DCS), OPERATING SYSTEM COMMAND
(OSC), PRIVACY MESSAGE (PM), or START OF STRING (SOS).
*/
