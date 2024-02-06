package ansiterm

import (
	"bytes"
)

func Strip(b bytes.Buffer, opt ...Option) (string, error) {
	h := handler{}
	parser := CreateParser("Ground", &h, opt...)
	_, err := parser.Parse(b.Bytes())
	if err != nil {
		return "", err
	}
	return h.buf.String(), nil
}

type handler struct {
	buf bytes.Buffer
}

func (h *handler) Print(b byte) error {
	h.buf.WriteByte(b)
	return nil
}

func (h *handler) Execute(b byte) error {
	return nil
}

func (h *handler) CUU(int) error {
	return nil
}

func (h *handler) CUD(int) error {
	return nil
}

func (h *handler) CUF(int) error {
	return nil
}

func (h *handler) CUB(int) error {
	return nil
}

func (h *handler) CNL(int) error {
	return nil
}

func (h *handler) CPL(int) error {
	return nil
}

func (h *handler) CHA(int) error {
	return nil
}

func (h *handler) VPA(int) error {
	return nil
}

func (h *handler) CUP(int, int) error {
	return nil
}

func (h *handler) HVP(int, int) error {
	return nil
}

func (h *handler) DECTCEM(bool) error {
	return nil
}

func (h *handler) DECOM(bool) error {
	return nil
}

func (h *handler) DECCOLM(bool) error {
	return nil
}

func (h *handler) ED(int) error {
	return nil
}

func (h *handler) EL(int) error {
	return nil
}

func (h *handler) IL(int) error {
	return nil
}

func (h *handler) DL(int) error {
	return nil
}

func (h *handler) ICH(int) error {
	return nil
}

func (h *handler) DCH(int) error {
	return nil
}

func (h *handler) SGR([]int) error {
	return nil
}

func (h *handler) SU(int) error {
	return nil
}

func (h *handler) SD(int) error {
	return nil
}

func (h *handler) DA([]string) error {
	return nil
}

func (h *handler) DECSTBM(int, int) error {
	return nil
}

func (h *handler) IND() error {
	return nil
}

func (h *handler) RI() error {
	return nil
}

func (h *handler) Flush() error {
	return nil
}
