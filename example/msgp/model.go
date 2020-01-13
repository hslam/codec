package msgp

//Object is a test struct
type Object struct {
	A uint32   `msg:"A"`
	B uint64   `msg:"B"`
	C float32  `msg:"C"`
	D float64  `msg:"D"`
	E string   `msg:"E"`
	F bool     `msg:"F"`
	G []byte   `msg:"G"`
	H [][]byte `msg:"H"`
}
