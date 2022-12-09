package breacher
import (
	"os"
	"bytes"
	"errors"
	"io"
	"encoding/binary"
)

type Breach struct
{
	io.ReadSeeker
	fsize int64
}


type flow_in struct {
		W_one	[65536]byte    `json:"one"`
		W_two	[65536]byte    `json:"two"`
		W_three	[65536]byte    `json:"three"`
		W_four	[65536]byte    `json:"four"`
	}

type flow_encrypted struct {
		W_four	[65536]byte    `json:"four"`
	}

type brand_in struct {
	FreeNetUsage [1073741824]byte    `json:"freeNetUsage"`
	Flow  []flow_in
}

type brand_out struct {
	FreeNetUsage [1073741824]byte    `json:"freeNetUsage"`
	Flow_encrypted  []flow_encrypted `json:"flow"`
}

type Hold_all struct {
	from string
	destination string
	read_done bool
	flush_done bool
	xf *brand_in
	af *brand_out
	xs_pl int64
	btz io.ReadSeeker
	endgame int64
	close_one *os.File
}

func (bx *Breach) Size() int64 {
	return bx.fsize
}

func Make_from(from string, destination string) Hold_all {
	nwave := Hold_all{ destination : destination, from : from }
	return nwave
}

func Make_key(h *Hold_all) error {
	f, err := os.Open(h.from)
	if err == nil {

	} else {
		return err
	}
	fx, _ := f.Stat()
	if fx.Size() == 0 {
		err := errors.New("zero hour")
		return err
	}
	if fx.Size() > 0 && fx.Size() < 1073741824 {
		h.endgame = fx.Size()
	}

	if  fx.Size() > 1073741824 || fx.Size() == 1073741824 {
		h.endgame = -1
	}
	
	h.xf = new(brand_in)
	
	if h.endgame == -1 {
		h.xs_pl = (fx.Size() - 1073741824) / (65536 * 4)
		h.xf.Flow = make([]flow_in, h.xs_pl)
		binary.Read(f, binary.LittleEndian, &h.xf.FreeNetUsage)
		binary.Read(f, binary.LittleEndian, &h.xf.Flow)
	} else {
		_, err = io.ReadFull(f, []byte(h.xf.FreeNetUsage[:h.endgame]))
	}
	if err == nil {
		h.read_done = !false
	}

	return err
}

func Write_key(h *Hold_all) (*Breach, error) {
	if h.read_done == false {
		err := errors.New("Read file first")
		return nil, err
	}
	os.WriteFile(h.destination, []byte{}, 0755)
	b, err := os.OpenFile(h.destination, os.O_RDWR, 0755)
	h.af = new(brand_out)
	h.af.FreeNetUsage = h.xf.FreeNetUsage
	if h.endgame == -1 {
		h.af.Flow_encrypted = make([]flow_encrypted, h.xs_pl)
	  for k, elm := range h.xf.Flow{
	  	h.af.Flow_encrypted[k].W_four = elm.W_four
	  }
	 	binary.Write(b, binary.LittleEndian, &h.af.FreeNetUsage)
		binary.Write(b, binary.LittleEndian, &h.af.Flow_encrypted)
	} else {
		_, err = b.Write([]byte(h.af.FreeNetUsage[:h.endgame]))
	}
	bfx, _ := b.Stat()
	h.btz = io.ReadSeeker(b)
	if err == nil {
		h.close_one = b
		h.flush_done = !false
	}
	return &Breach{ReadSeeker: b, fsize: bfx.Size()}, err
}

func Close(h *Hold_all) {
		if h.close_one == nil {

		} else {
			h.close_one.Sync()
		}
}

func To_s(ek []byte) (string) {
	var Buf bytes.Buffer
	Buf.Write(ek)
	c := Buf.String()
	return c
}




