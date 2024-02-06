package ansiterm

type stateID int

type state interface {
	Enter() error
	Exit() error
	Handle(byte) (state, error)
	Name() string
	Transition(state) error
}

type baseState struct {
	name   string
	parser *AnsiParser
}

func (base baseState) Enter() error {
	return nil
}

func (base baseState) Exit() error {
	return nil
}

func (base baseState) Handle(b byte) (s state, e error) {
	fe := base.parser.fe
	switch {
	case !fe && b == CSI_ENTRY:
		return base.parser.csiEntry, nil
	case !fe && b == DCS_ENTRY:
		return base.parser.dcsEntry, nil
	case b == ANSI_ESCAPE_PRIMARY:
		return base.parser.escape, nil
	case !fe && b == OSC_STRING:
		return base.parser.oscString, nil
	case sliceContains(getToGroundBytes(fe), b):
		return base.parser.ground, nil
	}

	return nil, nil
}

func (base baseState) Name() string {
	return base.name
}

func (base baseState) Transition(s state) error {
	if s == base.parser.ground {
		if sliceContains(getExecBytes(base.parser.fe), base.parser.context.currentChar) {
			return base.parser.execute()
		}
	}

	return nil
}

type dcsEntryState struct {
	baseState
}

type errorState struct {
	baseState
}
