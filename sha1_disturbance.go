package detectcoll

import "fmt"

type Sha1_dv struct {
	Type          int
	K             int
	B             uint
	message_block sha1_mb
}

var SHA1_STANDARD_DV []Sha1_dv = []Sha1_dv{
	Sha1_dv{Type: 1, K: 46, B: 0},
	Sha1_dv{Type: 1, K: 47, B: 0},
	Sha1_dv{Type: 1, K: 48, B: 0},
	Sha1_dv{Type: 1, K: 49, B: 0},
	Sha1_dv{Type: 1, K: 50, B: 0},
	Sha1_dv{Type: 1, K: 51, B: 0},
	Sha1_dv{Type: 1, K: 48, B: 2},
	Sha1_dv{Type: 1, K: 49, B: 2},
	Sha1_dv{Type: 1, K: 50, B: 2},
	Sha1_dv{Type: 2, K: 45, B: 0},
	Sha1_dv{Type: 2, K: 46, B: 0},
	Sha1_dv{Type: 2, K: 47, B: 0},
	Sha1_dv{Type: 2, K: 48, B: 0},
	Sha1_dv{Type: 2, K: 49, B: 0},
	Sha1_dv{Type: 2, K: 50, B: 0},
	Sha1_dv{Type: 2, K: 51, B: 0},
	Sha1_dv{Type: 2, K: 52, B: 0},
	Sha1_dv{Type: 2, K: 53, B: 0},
	Sha1_dv{Type: 2, K: 54, B: 0},
	Sha1_dv{Type: 2, K: 56, B: 0},
	Sha1_dv{Type: 2, K: 57, B: 0},
	Sha1_dv{Type: 2, K: 58, B: 0},
}

var SHA1_COMPLETE_DV []Sha1_dv = []Sha1_dv{
	Sha1_dv{Type: 1, K: 42, B: 0},
	Sha1_dv{Type: 1, K: 43, B: 0},
	Sha1_dv{Type: 1, K: 44, B: 0},
	Sha1_dv{Type: 1, K: 45, B: 0},
	Sha1_dv{Type: 1, K: 46, B: 0},
	Sha1_dv{Type: 1, K: 47, B: 0},
	Sha1_dv{Type: 1, K: 48, B: 0},
	Sha1_dv{Type: 1, K: 49, B: 0},
	Sha1_dv{Type: 1, K: 50, B: 0},
	Sha1_dv{Type: 1, K: 51, B: 0},
	Sha1_dv{Type: 1, K: 52, B: 0},
	Sha1_dv{Type: 1, K: 53, B: 0},
	Sha1_dv{Type: 1, K: 54, B: 0},
	Sha1_dv{Type: 1, K: 55, B: 0},
	Sha1_dv{Type: 1, K: 56, B: 0},
	Sha1_dv{Type: 1, K: 42, B: 2},
	Sha1_dv{Type: 1, K: 43, B: 2},
	Sha1_dv{Type: 1, K: 44, B: 2},
	Sha1_dv{Type: 1, K: 45, B: 2},
	Sha1_dv{Type: 1, K: 46, B: 2},
	Sha1_dv{Type: 1, K: 47, B: 2},
	Sha1_dv{Type: 1, K: 48, B: 2},
	Sha1_dv{Type: 1, K: 49, B: 2},
	Sha1_dv{Type: 1, K: 50, B: 2},
	Sha1_dv{Type: 1, K: 51, B: 2},
	Sha1_dv{Type: 1, K: 52, B: 2},
	Sha1_dv{Type: 1, K: 53, B: 2},
	Sha1_dv{Type: 1, K: 54, B: 2},
	Sha1_dv{Type: 1, K: 55, B: 2},
	Sha1_dv{Type: 1, K: 56, B: 2},
	Sha1_dv{Type: 2, K: 44, B: 0},
	Sha1_dv{Type: 2, K: 45, B: 0},
	Sha1_dv{Type: 2, K: 46, B: 0},
	Sha1_dv{Type: 2, K: 47, B: 0},
	Sha1_dv{Type: 2, K: 48, B: 0},
	Sha1_dv{Type: 2, K: 49, B: 0},
	Sha1_dv{Type: 2, K: 50, B: 0},
	Sha1_dv{Type: 2, K: 51, B: 0},
	Sha1_dv{Type: 2, K: 52, B: 0},
	Sha1_dv{Type: 2, K: 53, B: 0},
	Sha1_dv{Type: 2, K: 54, B: 0},
	Sha1_dv{Type: 2, K: 55, B: 0},
	Sha1_dv{Type: 2, K: 56, B: 0},
	Sha1_dv{Type: 2, K: 45, B: 2},
	Sha1_dv{Type: 2, K: 46, B: 2},
	Sha1_dv{Type: 2, K: 47, B: 2},
	Sha1_dv{Type: 2, K: 48, B: 2},
	Sha1_dv{Type: 2, K: 49, B: 2},
	Sha1_dv{Type: 2, K: 50, B: 2},
	Sha1_dv{Type: 2, K: 51, B: 2},
	Sha1_dv{Type: 2, K: 52, B: 2},
	Sha1_dv{Type: 2, K: 53, B: 2},
	Sha1_dv{Type: 2, K: 54, B: 2},
	Sha1_dv{Type: 2, K: 55, B: 2},
	Sha1_dv{Type: 2, K: 56, B: 2},
}

func (dv *Sha1_dv) String() string {
	var t string
	switch dv.Type {
	case 1:
		t = "I"
	case 2:
		t = "II"
	}
	return fmt.Sprintf("%s(%d, %d)", t, dv.K, dv.B)
}

func (dv *Sha1_dv) create_messageblock() {
	switch dv.Type {
	case 1:
		dv.message_block[dv.K+0] = (1 << 4)
		dv.message_block[dv.K+1] = (1 << 29) | (1 << 31)
		dv.message_block[dv.K+3] = (1 << 29)
		dv.message_block[dv.K+4] = (1 << 29)
		dv.message_block[dv.K+15] = (1 << 0)
	case 2:
		dv.message_block[dv.K+0] = (1 << 29)
		dv.message_block[dv.K+1] = (1 << 31)
		dv.message_block[dv.K+2] = (1 << 4)
		dv.message_block[dv.K+4] = (1 << 4) | (1 << 29)
		dv.message_block[dv.K+5] = (1 << 29) | (1 << 31)
		dv.message_block[dv.K+7] = (1 << 29)
		dv.message_block[dv.K+8] = (1 << 29)
		dv.message_block[dv.K+15] = (1 << 0)
	default:
		panic("Not a supported disturbance-vector type")
	}
	if dv.B != 0 {
		for i := 0; i < 16; i++ {
			dv.message_block[dv.K+i] = rotl32(dv.message_block[dv.K+i], dv.B)
		}
	}
	for i := dv.K + 16; i < 80; i++ {
		dv.message_block[i] = rotl32(dv.message_block[i-3]^dv.message_block[i-8]^dv.message_block[i-14]^dv.message_block[i-16], 1)
	}
	for i := dv.K - 1; i >= 0; i-- {
		dv.message_block[i] = rotr32(dv.message_block[i+16], 1) ^ dv.message_block[i+13] ^ dv.message_block[i+8] ^ dv.message_block[i+2]
	}
}

func init() {
	for i, _ := range SHA1_STANDARD_DV {
		SHA1_STANDARD_DV[i].create_messageblock()
	}
	for i, _ := range SHA1_COMPLETE_DV {
		SHA1_COMPLETE_DV[i].create_messageblock()
	}
}

func (dv *Sha1_dv) disturb(mb *sha1_mb) sha1_mb {
	var ret sha1_mb
	for i := 0; i < 80; i++ {
		ret[i] = dv.message_block[i] ^ mb[i]
	}
	return ret
}
