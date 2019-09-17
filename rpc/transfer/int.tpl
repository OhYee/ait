{{if (eq .length "1")}}
    // From{{upperFirstChar .type}} transfer from `{{.type}}` to `[]byte`
    func From{{upperFirstChar .type}} (value {{.type}}) (b []byte, err error) {
        b = []byte{Type{{upperFirstChar .type}}, u{{.type}}(value)}
        return
    }

    // To{{upperFirstChar .type}} transfer from `[]byte` to `{{.type}}`
    func To{{upperFirstChar .type}}(r io.Reader) (value {{.type}}, err error) {
        var t []byte
        t, err = goutil.ReadNBytes(r, 1)
        if t[0] != Type{{upperFirstChar .type}} {
            err = fmt.Errorf("Transfer error: want %d(%T), got %d", Type{{upperFirstChar .type}}, value, t[0])
        }
        value, err = to{{upperFirstChar .type}}(r)
        return
    }

    func to{{upperFirstChar .type}}(r io.Reader) (value {{.type}}, err error) {
        var data []byte
        data, err = goutil.ReadNBytes(r, 1)
        if err != nil {
            return
        }
        value = {{.type}}(data[0])
        return
    }

    // FromU{{.type}} transfer from `u{{.type}}` to `[]byte`
    func FromU{{.type}} (value u{{.type}}) (b []byte, err error) {
        b = []byte{TypeU{{.type}}, value}
        return
    }

    // ToU{{.type}} transfer from `[]byte` to `u{{.type}}`
    func ToU{{.type}}(r io.Reader) (value u{{.type}}, err error) {
        var t []byte
        t, err = goutil.ReadNBytes(r, 1)
        if t[0] != TypeU{{.type}} {
            err = fmt.Errorf("Transfer error: want %d(%T), got %d", TypeU{{.type}}, value, t[0])
        }
        value, err = toU{{.type}}(r)
        return
    }

    func toU{{.type}}(r io.Reader) (value u{{.type}}, err error) {
        var data []byte
        data, err = goutil.ReadNBytes(r, 1)
        if err != nil {
            return
        }
        value = u{{.type}}(data[0])
        return
    }
{{else}}
    // From{{upperFirstChar .type}} transfer from `{{.type}}` to `[]byte`
    func From{{upperFirstChar .type}} (value {{.type}}) (b []byte, err error) {
        b = make([]byte, {{.length}})
        byteOrder.PutU{{.type}}(b, u{{.type}}(value))
        b = append([]byte{Type{{upperFirstChar .type}}}, b...)
        return
    }

    // To{{upperFirstChar .type}} transfer from `[]byte` to `{{.type}}`
    func To{{upperFirstChar .type}}(r io.Reader) (value {{.type}}, err error) {
        var t []byte
        t, err = goutil.ReadNBytes(r, 1)
        if t[0] != Type{{upperFirstChar .type}} {
            err = fmt.Errorf("Transfer error: want %d(%T), got %d", Type{{upperFirstChar .type}}, value, t[0])
        }
        value, err = to{{upperFirstChar .type}}(r)
        return
    }

    func to{{upperFirstChar .type}}(r io.Reader) (value {{.type}}, err error) {
        var data []byte
        data, err = goutil.ReadNBytes(r, {{.length}})
        if err != nil {
            return
        }
        value = {{.type}}(byteOrder.U{{.type}}(data))
        return
    }

    // FromU{{.type}} transfer from `u{{.type}}` to `[]byte`
    func FromU{{.type}} (value u{{.type}}) (b []byte, err error) {
        b = make([]byte, {{.length}})
        byteOrder.PutU{{.type}}(b, value)
        b = append([]byte{TypeU{{.type}}}, b...)

        return
    }

    // ToU{{.type}} transfer from `[]byte` to `u{{.type}}`
    func ToU{{.type}}(r io.Reader) (value u{{.type}}, err error) {
        var t []byte
        t, err = goutil.ReadNBytes(r, 1)
        if t[0] != TypeU{{.type}} {
            err = fmt.Errorf("Transfer error: want %d(%T), got %d", TypeU{{.type}}, value, t[0])
        }
        value, err = toU{{.type}}(r)
        return
    }

    func toU{{.type}}(r io.Reader) (value u{{.type}}, err error) {
        var data []byte
        data, err = goutil.ReadNBytes(r, {{.length}})
        if err != nil {
            return
        }
        value = byteOrder.U{{.type}}(data)
        return
    }
{{end}}