package main

import (
	"doc-parser/classes"
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteLine(line string, f *os.File) {
	f.WriteString(line + "\n")
}

func CreateFiles(packageName, directory string, model *classes.Model) {
	for _, c := range model.Classes {
		if packageName == "v2" {
			if strings.HasPrefix(c.Name, "P_") {
				createFile(packageName, directory, c, nil)
			}
		} else {
			createFile(packageName, directory, c, nil)
		}
	}
	for _, e := range model.Enumerations {
		if packageName == "v2" {
			if strings.HasPrefix(e.Name, "P_") {
				createFile(packageName, directory, nil, e)
			}
		} else {
			createFile(packageName, directory, nil, e)
		}
	}
}

/*
*
This is a rather long line that needs
word wrapping to an arbirtary line
lenght so it's easier to read it.
*/
func createFile(packageName, directory string, c *classes.Class, e *classes.Enumeration) {
	name := ""
	topComment := ""
	if c != nil {
		name = UpperCamelCase(c.Name)
		topComment = c.Comment
	}
	if e != nil {
		name = UpperCamelCase(e.Name)
		topComment = e.Comment
	}
	w, err := os.Create(directory + "/" + name + ".go")

	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	if c != nil && e == nil {
		WriteLine("package "+packageName, w)
		WriteLine("", w)
		WriteLine(CommentLines(topComment, 80, 0), w)
		WriteLine("", w)
		WriteLine("type I"+name+" interface {", w)
		for _, f := range c.Functions {
			pIn, pOut := ConstructFunctions(f)
			WriteLine(CommentLines(f.Comment, 80, 1), w)
			WriteLine("\t"+f.Name+" ( "+pIn+" ) "+pOut, w)
		}
		WriteLine("}", w)
		WriteLine("", w)
		WriteLine("type "+name+" struct {", w)
		WriteLine("}", w)
	}
	if e != nil {
	}
}

func ConstructFunctions(f *classes.Function) (in, out string) {
	var parameterIn string
	for j, p := range f.In {
		var genericPart string
		if p.IsGeneric {
			for k, g := range p.GenericPartTypes {
				if k == 0 {
					genericPart = g
				} else {
					genericPart = genericPart + ", " + g
				}
			}
		}
		if j == 0 {
			if p.IsGeneric {
				parameterIn = p.Name + " < " + genericPart + " > "
			} else {
				parameterIn = p.Name + " " + p.Type
			}
		} else {
			if p.IsGeneric {
				parameterIn = parameterIn + ", " + p.Name + " < " + genericPart + " > "
			} else {
				parameterIn = parameterIn + ", " + p.Name + " " + p.Type
			}
		}
	}
	var parameterOut string
	for j, p := range f.Out {
		var genericPart string
		if p.IsGeneric {
			for k, g := range p.GenericPartTypes {
				if k == 0 {
					genericPart = g
				} else {
					genericPart = genericPart + ", " + g
				}
			}
		}
		if j == 0 {
			if p.IsGeneric {
				parameterOut = p.Name + " < " + genericPart + " > "
			} else {
				parameterOut = p.Name + " " + p.Type
			}
		} else {
			if p.IsGeneric {
				parameterOut = parameterIn + ", " + p.Name + " < " + genericPart + " > "
			} else {
				parameterOut = parameterIn + ", " + p.Name + " " + p.Type
			}
		}
	}
	return parameterIn, parameterOut
}

func UpperCamelCase(s string) string {
	return camelCase(s, true)
}

// LowerCamelCase converts a string into camel case starting with a lower case letter.
func LowerCamelCase(s string) string {
	return camelCase(s, false)
}

func camelCase(s string, upper bool) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	stringIter(s, func(prev, curr, next rune) {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (upper && prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else if isLower(prev) {
				buffer = append(buffer, curr)
			} else if isUpper(prev) && isUpper(curr) && isLower(next) {
				// Assume a case like "R" for "XRequestId"
				buffer = append(buffer, curr)
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
	})

	return string(buffer)
}

func isLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// toLower converts a character in the range of ASCII characters 'A' to 'Z' to its lower
// case counterpart. Other characters remain the same.
func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

// isLower checks if a character is upper case. More precisely it evaluates if it is
// in the range of ASCII characters 'A' to 'Z'.
func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// toLower converts a character in the range of ASCII characters 'a' to 'z' to its lower
// case counterpart. Other characters remain the same.
func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

// isSpace checks if a character is some kind of whitespace.
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isDelimiter checks if a character is some kind of whitespace or '_' or '-'.
func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

// iterFunc is a callback that is called fro a specific position in a string. Its arguments are the
// rune at the respective string position as well as the previous and the next rune. If curr is at the
// first position of the string prev is zero. If curr is at the end of the string next is zero.
type iterFunc func(prev, curr, next rune)

// stringIter iterates over a string, invoking the callback for every single rune in the string.
func stringIter(s string, callback iterFunc) {
	var prev rune
	var curr rune
	for _, next := range s {
		if curr == 0 {
			prev = curr
			curr = next
			continue
		}

		callback(prev, curr, next)

		prev = curr
		curr = next
	}

	if len(s) > 0 {
		callback(prev, curr, 0)
	}
}

func CommentLines(line string, lineWidth int, tab int) string {
	if len(line) < 80 {
		if strings.HasPrefix(line, "\t\t") {
			return "\t\t// " + line[2:]
		}
		if strings.HasPrefix(line, "\t") {
			return "\t// " + line[1:]
		}
		return "// " + line
	}
	var tabString string
	i := 0
	for i < tab {
		tabString = tabString + "\t"
		i = i + 1
	}
	return fmt.Sprintf(tabString+"/**\n%s\n"+tabString+"*/", tabString+"\t"+wordWrap(line, lineWidth, tabString+"\t"))
}

func wordWrap(text string, lineWidth int, tabString string) (wrapped string) {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return tabString + text
	}
	wrapped = words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + tabString + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return
}
