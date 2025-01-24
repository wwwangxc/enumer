package main

// Arguments to format are:
//
//	[1]: type name
const yamlMethods = `
// MarshalYAML implements a YAML Marshaler for %[1]s
func (i %[1]s) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for %[1]s
func (i *%[1]s) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	val, ok := %[1]sString(s)
	if !ok {
		return fmt.Errorf("%%s does not belong to %[1]s values", s)
	}

	*i = val
	return nil
}
`

func (g *Generator) buildYAMLMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(yamlMethods, typeName)
}
