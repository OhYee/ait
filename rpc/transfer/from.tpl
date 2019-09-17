func fromValue(value any) (b []byte, err error) {
	switch value.(type) {

    {{range $type := .}}
    {{template "case" $type}}
    {{end}}

	default:
		err = fmt.Errorf("Can not solve the unknown type: %T", value)
	}
	return
}
{{define "case"}}
case {{.}}:
		v, _ := value.({{.}})
		b, err = From{{upperFirstChar .}}(v)
{{end}}