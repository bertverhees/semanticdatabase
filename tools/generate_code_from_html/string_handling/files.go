package string_handling

import (
	"doc-parser/classes"
	"log"
	"os"
	"strings"
)

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
		CommentLines(topComment, 80, 0, w)
		WriteLine("", w)
		WriteLine("type I"+name+" interface {", w)
		for _, f := range c.Functions {
			pIn, pOut := ConstructFunctions(f)
			//WriteLine(CommentLines(f.Comment, 80, 1), w)
			WriteLine("\t"+UpperCamelCase(f.Name)+" ( "+pIn+" ) "+pOut, w)
		}
		WriteLine("}", w)
		WriteLine("", w)
		WriteLine("type "+name+" struct {", w)
		for _, co := range c.Constants {
			CommentLines(co.Comment, 80, 1, w)
			WriteLine("\t"+UpperCamelCase(co.Name)+"\t"+formatType(co.Type)+"\t`yaml:\""+strings.ToLower(co.Name)+"\" json:\""+strings.ToLower(co.Name)+"\" xml:\""+strings.ToLower(co.Name)+"\"`", w)
		}
		WriteLine("}", w)
	}
	if e != nil {
	}
}
