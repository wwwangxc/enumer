package main

import "fmt"

// Arguments to format are:
//
//	[1]: type name
const stringNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, bool) {
	if val, ok := _%[1]sNameToValueMap[s]; ok {
		return val, true
	}

	if val, ok := _%[1]sNameToValueMap[strings.ToLower(s)]; ok {
		return val, true
	}

	return 0, false
}
`

// Arguments to format are:
//
//	[1]: type name
const stringValuesMethod = `// All%[1]s returns all values of the enum
func All%[1]s() []%[1]s {
	return _%[1]sValues
}
`

// Arguments to format are:
//
//	[1]: type name
const stringsMethod = `// All%[1]sString returns a slice of all String values of the enum
func All%[1]sString() []string {
	strs := make([]string, len(_%[1]sNames))
	copy(strs, _%[1]sNames)
	return strs
}
`

// Arguments to format are:
//
//	[1]: type name
const stringBelongsMethodLoop = `// IsValid returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsValid() bool {
	for _, v := range _%[1]sValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Arguments to format are:
//
//	[1]: type name
const stringBelongsMethodSet = `// IsValid returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsValid() bool {
	_, ok := _%[1]sMap[i] 
	return ok
}
`

// Arguments to format are:
//
//	[1]: type name
const altStringValuesMethod = `func (%[1]s) Values() []string {
	return All%[1]sString()
}
`

func (g *Generator) buildAltStringValuesMethod(typeName string) {
	g.Printf("\n")
	g.Printf(altStringValuesMethod, typeName)
}

func (g *Generator) buildBasicExtras(runs [][]Value, typeName string, runsThreshold int) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called

	// Print the slice of values
	g.Printf("\nvar _%sValues = []%s{", typeName, typeName)
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%s, ", value.originalName)
		}
	}
	g.Printf("}\n\n")

	// Print the map between name and value
	g.printValueMap(runs, typeName, runsThreshold)

	// Print the slice of names
	g.printNamesSlice(runs, typeName, runsThreshold)

	// Print the basic extra methods
	g.Printf(stringNameToValueMethod, typeName)
	g.Printf(stringValuesMethod, typeName)
	g.Printf(stringsMethod, typeName)
	if len(runs) <= runsThreshold {
		g.Printf(stringBelongsMethodLoop, typeName)
	} else { // There is a map of values, the code is simpler then
		g.Printf(stringBelongsMethodSet, typeName)
	}
}

func (g *Generator) printValueMap(runs [][]Value, typeName string, runsThreshold int) {
	thereAreRuns := len(runs) > 1 && len(runs) <= runsThreshold
	g.Printf("\nvar _%sNameToValueMap = map[string]%s{\n", typeName, typeName)

	var n int
	var runID string
	for i, values := range runs {
		if thereAreRuns {
			runID = "_" + fmt.Sprintf("%d", i)
			n = 0
		} else {
			runID = ""
		}

		for _, value := range values {
			g.Printf("\t_%sName%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), value.originalName)
			g.Printf("\t_%sLowerName%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), value.originalName)
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")
}
func (g *Generator) printNamesSlice(runs [][]Value, typeName string, runsThreshold int) {
	thereAreRuns := len(runs) > 1 && len(runs) <= runsThreshold
	g.Printf("\nvar _%sNames = []string{\n", typeName)

	var n int
	var runID string
	for i, values := range runs {
		if thereAreRuns {
			runID = "_" + fmt.Sprintf("%d", i)
			n = 0
		} else {
			runID = ""
		}

		for _, value := range values {
			g.Printf("\t_%sName%s[%d:%d],\n", typeName, runID, n, n+len(value.name))
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")
}
