package main

// Arguments to format are:
//
//	[1]: type name
const textMethods = `
// MarshalText implements the encoding.TextMarshaler interface for %[1]s
func (i %[1]s) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for %[1]s
func (i *%[1]s) UnmarshalText(text []byte) error {
	val, ok := %[1]sString(string(text))
	if !ok {
		return fmt.Errorf("%%s does not belong to %[1]s values", string(text))
	}

	*i = val
	return nil
}
`

func (g *Generator) buildTextMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(textMethods, typeName)
}
