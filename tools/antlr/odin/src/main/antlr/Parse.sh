#!/bin/sh

alias antlr='java -jar ~/Development/antlr4/antlr-4.13.1-complete.jar'

antlr -Dlanguage=Go -o ../../../../../../vocabulary/gen/odin OdinParser.g4
antlr -Dlanguage=Go -o ../../../../../../vocabulary/gen/odin OdinLexer.g4
