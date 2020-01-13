package model

//Object is a test struct
type Object struct {
	A uint32   `json:"A" xml:"A"`
	B uint64   `json:"B" xml:"B"`
	C float32  `json:"C" xml:"C"`
	D float64  `json:"D" xml:"D"`
	E string   `json:"E" xml:"E"`
	F bool     `json:"F" xml:"F"`
	G []byte   `json:"G" xml:"-"`
	H [][]byte `json:"H" xml:"-"`
}
