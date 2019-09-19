// FromValue from a build-in type value to bytes
func FromValue(value any) (b []byte, err error) {
	switch value.(type) {

    {{range $type := .}}
        {{if isString $type}}
            {{template "case" makeMap "type" $type "name" $type}}
        {{else}}
            {{template "case" $type}}
        {{end}}
    {{end}}

	default:
		err = fmt.Errorf("Transfer Error: Can not solve the unknown type: %T", value)
	}
	return
}

{{define "case"}}
case {{.type}}:
		v, _ := value.({{.type}})
		b, err = From{{upperFirstChar .name}}(v)
{{end}}