package main

import (
	"doc-parser/classes"
	"fmt"
	"log"
	"os"
	"strings"
)

func WriteBasicDefinitions(packageName, directory string) {
	w, err := os.Create(directory + "/BasicDefinitions.go")

	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	WriteLine("package "+packageName, w)
	WriteLine("", w)
	WriteLine(CommentLines("Defines globally used constant values.", 80), w)
	WriteLine("", w)
	WriteLine("type BasicDefinition struct {", w)
	WriteLine(CommentLines("\tCarriage return character.", 80), w)
	WriteLine("\tCR\trune\t`yaml:\"cr\" json:\"cr\" xml:\"cr\"`", w)
	WriteLine(CommentLines("\tLine feed character.", 80), w)
	WriteLine("\tLF\trune\t`yaml:\"lf\" json:\"lf\" xml:\"lf\"`", w)
	WriteLine("\tAnyTypeName\tstring\t`yaml:\"anynametype\" json:\"anynametype\" xml:\"anynametype\"`", w)
	WriteLine("\tRegexAnyPattern\tstring\t`yaml:\"regexanypattern\" json:\"regexanypattern\" xml:\"regexanypattern\"`", w)
	WriteLine("\tDefaultEncoding\tstring\t`yaml:\"defaultencoding\" json:\"defaultencoding\" xml:\"defaultencoding\"`", w)
	WriteLine("\tNoneTypeName\tstring\t`yaml:\"nonetypename\" json:\"nonetypename\" xml:\"nonetypename\"`", w)
	WriteLine("}", w)
}

func WriteLine(line string, f *os.File) {
	f.WriteString(line + "\n")
}

func CreateFiles(packageName, directory string, model *classes.Model) {
	for _, c := range model.Classes {
		createFile(packageName, directory, c, nil)
	}
	for _, e := range model.Enumerations {
		createFile(packageName, directory, nil, e)
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
	WriteLine("package "+packageName, w)
	WriteLine("", w)
	WriteLine(CommentLines(topComment, 80), w)
	WriteLine("", w)
	WriteLine("type "+name+"er interface {", w)
	WriteLine("}", w)
	WriteLine("", w)
	WriteLine("type "+name+" struct {", w)
	WriteLine("}", w)

	if c != nil {
	}
	if e != nil {
	}
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

func CommentLines(line string, lineWidth int) string {
	if len(line) < 80 {
		if strings.HasPrefix(line, "\t\t") {
			return "\t\t// " + line[2:]
		}
		if strings.HasPrefix(line, "\t") {
			return "\t// " + line[1:]
		}
		return "// " + line
	}
	return fmt.Sprintf("/**\n%s\n*/", wordWrap(line, lineWidth))
}

func wordWrap(text string, lineWidth int) (wrapped string) {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped = words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return
}
