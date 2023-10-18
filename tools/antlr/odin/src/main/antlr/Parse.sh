#!/bin/sh

alias antlr='java -jar ~/Development/antlr4/antlr-4.7-complete.jar'

antlr -Dlanguage=Go -o ../../../../../../gen OdinParser.g4
