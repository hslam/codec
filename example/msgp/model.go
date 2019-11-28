package msgp

type Student struct {
	Name string 		`msg:"Name"`
	Age int32 			`msg:"Age"`
	Address string 		`msg:"Address"`
}
