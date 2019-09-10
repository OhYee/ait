{{if (eq .length "1")}}
    // from{{upperFirstChar .type}} transfer from `{{.type}}` to `[]byte`
    func from{{upperFirstChar .type}} (value {{.type}}) (b []byte, err error) {
        b = []byte{Type{{upperFirstChar .type}}, u{{.type}}(value)}
        return
    }

    // fromU{{.type}} transfer from `u{{.type}}` to `[]byte`
    func fromU{{.type}} (value u{{.type}}) (b []byte, err error) {
        b = []byte{TypeU{{.type}}, value}
        return
    }
{{else}}
    // from{{upperFirstChar .type}} transfer from `{{.type}}` to `[]byte`
    func from{{upperFirstChar .type}} (value {{.type}}) (b []byte, err error) {
        b = make([]byte, {{.length}})
        byteOrder.PutU{{.type}}(b, u{{.type}}(value))
        b = append([]byte{Type{{upperFirstChar .type}}}, b...)
        return
    }

    // fromU{{.type}} transfer from `u{{.type}}` to `[]byte`
    func fromU{{.type}} (value u{{.type}}) (b []byte, err error) {
        b = make([]byte, {{.length}})
        byteOrder.PutU{{.type}}(b, value)
        b = append([]byte{TypeU{{.type}}}, b...)

        return
    }
{{end}}