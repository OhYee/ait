// ToValue from bytes to build-in type value
func ToValue(r io.Reader) (value goutil.Any, err error) {
	var t []byte
	t,err = goutil.ReadNBytes(r,1)
	switch t[0] {
	{{range $type := .}}
    {{template "case" $type}}
    {{end}}
	default:
		err = fmt.Errorf("Transfer Error: Can not identify type %d",t[0])
	}
	return
}
{{define "case"}}
case Type{{upperFirstChar .}}:
		value, err = to{{upperFirstChar .}}(r)
{{end}}