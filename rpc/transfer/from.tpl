// FromValue from a build-in type value to bytes
func FromValue(value any) (b []byte, err error) {
	switch value.(type) {

    {{range $type := .}}
    {{template "case" $type}}
    {{end}}

	default:
		err = fmt.Errorf("Transfer Error: Can not solve the unknown type: %T", value)
	}
	return
}
{{define "case"}}
case {{.}}:
		v, _ := value.({{.}})
		b, err = From{{upperFirstChar .}}(v)
{{end}}