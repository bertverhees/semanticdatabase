// Code generated from OdinParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // OdinParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type OdinParser struct {
	*antlr.BaseParser
}

var OdinParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func odinparserParserInit() {
	staticData := &OdinParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'...'", "'='", "'/'", "'::'", "','", "", "", "'>'",
		"'<'", "", "'+'", "'-'", "'%'", "'^'", "'|'", "'..'", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "'.'", "';'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "CMT_LINE", "EOL", "WS", "ODIN_URI", "SYM_LIST_CONTINUE", "SYM_EQ",
		"SYM_SLASH", "SYM_NAMESPACE_SEP", "SYM_COMMA", "SYM_LE", "SYM_GE", "SYM_GT",
		"SYM_LT", "SYM_PLUS_OR_MINUS", "SYM_PLUS", "SYM_MINUS", "SYM_PERCENT",
		"SYM_CARET", "SYM_IVL_DELIM", "SYM_IVL_SEP", "OBJECT_VERSION_ID", "ARCHETYPE_HRID",
		"ARCHETYPE_REF", "VERSION_ID", "TERM_CODE_REF", "ROOT_ID_CODE", "ID_CODE",
		"AT_CODE", "AC_CODE", "ISO8601_DATE_AUGMENTED", "ISO8601_TIME_AUGMENTED",
		"ISO8601_DATE_TIME_AUGMENTED", "ISO8601_DURATION", "SYM_TRUE", "SYM_FALSE",
		"GUID", "UUID", "INTEGER", "REAL", "SCI_INTEGER", "SCI_REAL", "STRING",
		"CHARACTER", "SYM_DOT", "SYM_SEMI_COLON", "SYM_LPAREN", "SYM_RPAREN",
		"SYM_LBRACKET", "SYM_RBRACKET", "SYM_LCURLY", "SYM_RCURLY", "UC_ID",
		"LC_ID", "QUALIFIED_TERM_CODE_REF",
	}
	staticData.RuleNames = []string{
		"odinObject", "odinAttrVal", "odinAttrName", "odinObjectBlock", "odinObjectValueBlock",
		"rmTypeSpec", "odinKeyedObject", "odinKeySpec", "odinObjectReferenceBlock",
		"odinPathList", "odinPath", "odinPathSegment", "rmTypeId", "primitiveObject",
		"primitiveValue", "primitiveListValue", "primitiveIntervalValue", "stringValue",
		"stringListValue", "integerValue", "integerListValue", "integerIntervalValue",
		"integerIntervalListValue", "realValue", "realListValue", "realIntervalValue",
		"realIntervalListValue", "booleanValue", "booleanListValue", "characterValue",
		"characterListValue", "dateValue", "dateListValue", "dateIntervalValue",
		"dateIntervalListValue", "timeValue", "timeListValue", "timeIntervalValue",
		"timeIntervalListValue", "dateTimeValue", "dateTimeListValue", "dateTimeIntervalValue",
		"dateTimeIntervalListValue", "durationValue", "durationListValue", "durationIntervalValue",
		"durationIntervalListValue", "termCodeValue", "termCodeListValue", "relop",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 604, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 4, 0, 102, 8, 0, 11, 0, 12, 0,
		103, 1, 0, 3, 0, 107, 8, 0, 1, 0, 3, 0, 110, 8, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 116, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 4, 3, 4,
		125, 8, 4, 1, 4, 1, 4, 1, 4, 4, 4, 130, 8, 4, 11, 4, 12, 4, 131, 1, 4,
		4, 4, 135, 8, 4, 11, 4, 12, 4, 136, 1, 4, 3, 4, 140, 8, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 165, 8, 9,
		11, 9, 12, 9, 166, 3, 9, 169, 8, 9, 1, 10, 3, 10, 172, 8, 10, 1, 10, 4,
		10, 175, 8, 10, 11, 10, 12, 10, 176, 1, 11, 1, 11, 1, 11, 3, 11, 182, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 189, 8, 12, 10, 12, 12, 12,
		192, 9, 12, 1, 12, 1, 12, 3, 12, 196, 8, 12, 1, 13, 1, 13, 1, 13, 3, 13,
		201, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 213, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 225, 8, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 233, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		4, 18, 240, 8, 18, 11, 18, 12, 18, 241, 1, 18, 1, 18, 3, 18, 246, 8, 18,
		1, 19, 3, 19, 249, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 4, 20, 256,
		8, 20, 11, 20, 12, 20, 257, 1, 20, 1, 20, 3, 20, 262, 8, 20, 1, 21, 1,
		21, 3, 21, 266, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 271, 8, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 278, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 289, 8, 21, 1, 22, 1, 22, 1,
		22, 4, 22, 294, 8, 22, 11, 22, 12, 22, 295, 1, 22, 1, 22, 3, 22, 300, 8,
		22, 1, 23, 3, 23, 303, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 4, 24,
		310, 8, 24, 11, 24, 12, 24, 311, 1, 24, 1, 24, 3, 24, 316, 8, 24, 1, 25,
		1, 25, 3, 25, 320, 8, 25, 1, 25, 1, 25, 1, 25, 3, 25, 325, 8, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 332, 8, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 343, 8, 25, 1, 26, 1, 26,
		1, 26, 4, 26, 348, 8, 26, 11, 26, 12, 26, 349, 1, 26, 1, 26, 3, 26, 354,
		8, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 4, 28, 361, 8, 28, 11, 28, 12,
		28, 362, 1, 28, 1, 28, 3, 28, 367, 8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		30, 4, 30, 374, 8, 30, 11, 30, 12, 30, 375, 1, 30, 1, 30, 3, 30, 380, 8,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 4, 32, 387, 8, 32, 11, 32, 12, 32,
		388, 1, 32, 1, 32, 3, 32, 393, 8, 32, 1, 33, 1, 33, 3, 33, 397, 8, 33,
		1, 33, 1, 33, 1, 33, 3, 33, 402, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 409, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 3, 33, 420, 8, 33, 1, 34, 1, 34, 1, 34, 4, 34, 425, 8, 34,
		11, 34, 12, 34, 426, 1, 34, 1, 34, 3, 34, 431, 8, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 4, 36, 438, 8, 36, 11, 36, 12, 36, 439, 1, 36, 1, 36,
		3, 36, 444, 8, 36, 1, 37, 1, 37, 3, 37, 448, 8, 37, 1, 37, 1, 37, 1, 37,
		3, 37, 453, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 460, 8, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 471,
		8, 37, 1, 38, 1, 38, 1, 38, 4, 38, 476, 8, 38, 11, 38, 12, 38, 477, 1,
		38, 1, 38, 3, 38, 482, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 4, 40,
		489, 8, 40, 11, 40, 12, 40, 490, 1, 40, 1, 40, 3, 40, 495, 8, 40, 1, 41,
		1, 41, 3, 41, 499, 8, 41, 1, 41, 1, 41, 1, 41, 3, 41, 504, 8, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 511, 8, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 522, 8, 41, 1, 42, 1, 42,
		1, 42, 4, 42, 527, 8, 42, 11, 42, 12, 42, 528, 1, 42, 1, 42, 3, 42, 533,
		8, 42, 1, 43, 3, 43, 536, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 4,
		44, 543, 8, 44, 11, 44, 12, 44, 544, 1, 44, 1, 44, 3, 44, 549, 8, 44, 1,
		45, 1, 45, 3, 45, 553, 8, 45, 1, 45, 1, 45, 1, 45, 3, 45, 558, 8, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 565, 8, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 576, 8, 45, 1, 46, 1,
		46, 1, 46, 4, 46, 581, 8, 46, 11, 46, 12, 46, 582, 1, 46, 1, 46, 3, 46,
		587, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 4, 48, 594, 8, 48, 11, 48,
		12, 48, 595, 1, 48, 1, 48, 3, 48, 600, 8, 48, 1, 49, 1, 49, 1, 49, 0, 0,
		50, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 6, 1, 0, 52,
		53, 1, 0, 15, 16, 2, 0, 38, 38, 40, 40, 2, 0, 39, 39, 41, 41, 1, 0, 34,
		35, 1, 0, 10, 13, 663, 0, 106, 1, 0, 0, 0, 2, 111, 1, 0, 0, 0, 4, 117,
		1, 0, 0, 0, 6, 121, 1, 0, 0, 0, 8, 124, 1, 0, 0, 0, 10, 143, 1, 0, 0, 0,
		12, 147, 1, 0, 0, 0, 14, 151, 1, 0, 0, 0, 16, 155, 1, 0, 0, 0, 18, 159,
		1, 0, 0, 0, 20, 171, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0, 24, 183, 1, 0, 0,
		0, 26, 200, 1, 0, 0, 0, 28, 212, 1, 0, 0, 0, 30, 224, 1, 0, 0, 0, 32, 232,
		1, 0, 0, 0, 34, 234, 1, 0, 0, 0, 36, 236, 1, 0, 0, 0, 38, 248, 1, 0, 0,
		0, 40, 252, 1, 0, 0, 0, 42, 288, 1, 0, 0, 0, 44, 290, 1, 0, 0, 0, 46, 302,
		1, 0, 0, 0, 48, 306, 1, 0, 0, 0, 50, 342, 1, 0, 0, 0, 52, 344, 1, 0, 0,
		0, 54, 355, 1, 0, 0, 0, 56, 357, 1, 0, 0, 0, 58, 368, 1, 0, 0, 0, 60, 370,
		1, 0, 0, 0, 62, 381, 1, 0, 0, 0, 64, 383, 1, 0, 0, 0, 66, 419, 1, 0, 0,
		0, 68, 421, 1, 0, 0, 0, 70, 432, 1, 0, 0, 0, 72, 434, 1, 0, 0, 0, 74, 470,
		1, 0, 0, 0, 76, 472, 1, 0, 0, 0, 78, 483, 1, 0, 0, 0, 80, 485, 1, 0, 0,
		0, 82, 521, 1, 0, 0, 0, 84, 523, 1, 0, 0, 0, 86, 535, 1, 0, 0, 0, 88, 539,
		1, 0, 0, 0, 90, 575, 1, 0, 0, 0, 92, 577, 1, 0, 0, 0, 94, 588, 1, 0, 0,
		0, 96, 590, 1, 0, 0, 0, 98, 601, 1, 0, 0, 0, 100, 102, 3, 2, 1, 0, 101,
		100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 107, 3, 8, 4, 0, 106, 101, 1, 0,
		0, 0, 106, 105, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 110, 5, 0, 0, 1,
		109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 1, 1, 0, 0, 0, 111, 112,
		3, 4, 2, 0, 112, 113, 5, 6, 0, 0, 113, 115, 3, 6, 3, 0, 114, 116, 5, 45,
		0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 3, 1, 0, 0, 0, 117,
		118, 7, 0, 0, 0, 118, 5, 1, 0, 0, 0, 119, 122, 3, 8, 4, 0, 120, 122, 3,
		16, 8, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 7, 1, 0, 0,
		0, 123, 125, 3, 10, 5, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 139, 5, 13, 0, 0, 127, 140, 3, 26, 13, 0, 128, 130,
		3, 2, 1, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 131, 132, 1, 0, 0, 0, 132, 140, 1, 0, 0, 0, 133, 135, 3, 12, 6, 0,
		134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136,
		137, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 140, 5, 4, 0, 0, 139, 127,
		1, 0, 0, 0, 139, 129, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 139, 138, 1, 0,
		0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 5, 12, 0, 0,
		142, 9, 1, 0, 0, 0, 143, 144, 5, 46, 0, 0, 144, 145, 3, 24, 12, 0, 145,
		146, 5, 47, 0, 0, 146, 11, 1, 0, 0, 0, 147, 148, 3, 14, 7, 0, 148, 149,
		5, 6, 0, 0, 149, 150, 3, 6, 3, 0, 150, 13, 1, 0, 0, 0, 151, 152, 5, 48,
		0, 0, 152, 153, 3, 28, 14, 0, 153, 154, 5, 49, 0, 0, 154, 15, 1, 0, 0,
		0, 155, 156, 5, 13, 0, 0, 156, 157, 3, 18, 9, 0, 157, 158, 5, 12, 0, 0,
		158, 17, 1, 0, 0, 0, 159, 168, 3, 20, 10, 0, 160, 161, 5, 9, 0, 0, 161,
		169, 5, 5, 0, 0, 162, 163, 5, 9, 0, 0, 163, 165, 3, 20, 10, 0, 164, 162,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0,
		0, 0, 167, 169, 1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 164, 1, 0, 0, 0,
		168, 169, 1, 0, 0, 0, 169, 19, 1, 0, 0, 0, 170, 172, 3, 14, 7, 0, 171,
		170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 175,
		3, 22, 11, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1,
		0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0, 178, 179, 5, 7, 0,
		0, 179, 181, 5, 53, 0, 0, 180, 182, 3, 14, 7, 0, 181, 180, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 23, 1, 0, 0, 0, 183, 195, 5, 52, 0, 0, 184,
		185, 5, 13, 0, 0, 185, 190, 3, 24, 12, 0, 186, 187, 5, 9, 0, 0, 187, 189,
		3, 24, 12, 0, 188, 186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1,
		0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0,
		0, 193, 194, 5, 12, 0, 0, 194, 196, 1, 0, 0, 0, 195, 184, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 25, 1, 0, 0, 0, 197, 201, 3, 28, 14, 0, 198, 201,
		3, 30, 15, 0, 199, 201, 3, 32, 16, 0, 200, 197, 1, 0, 0, 0, 200, 198, 1,
		0, 0, 0, 200, 199, 1, 0, 0, 0, 201, 27, 1, 0, 0, 0, 202, 213, 3, 34, 17,
		0, 203, 213, 3, 38, 19, 0, 204, 213, 3, 46, 23, 0, 205, 213, 3, 54, 27,
		0, 206, 213, 3, 58, 29, 0, 207, 213, 3, 94, 47, 0, 208, 213, 3, 62, 31,
		0, 209, 213, 3, 70, 35, 0, 210, 213, 3, 78, 39, 0, 211, 213, 3, 86, 43,
		0, 212, 202, 1, 0, 0, 0, 212, 203, 1, 0, 0, 0, 212, 204, 1, 0, 0, 0, 212,
		205, 1, 0, 0, 0, 212, 206, 1, 0, 0, 0, 212, 207, 1, 0, 0, 0, 212, 208,
		1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 211, 1, 0,
		0, 0, 213, 29, 1, 0, 0, 0, 214, 225, 3, 36, 18, 0, 215, 225, 3, 40, 20,
		0, 216, 225, 3, 48, 24, 0, 217, 225, 3, 56, 28, 0, 218, 225, 3, 60, 30,
		0, 219, 225, 3, 96, 48, 0, 220, 225, 3, 64, 32, 0, 221, 225, 3, 72, 36,
		0, 222, 225, 3, 80, 40, 0, 223, 225, 3, 88, 44, 0, 224, 214, 1, 0, 0, 0,
		224, 215, 1, 0, 0, 0, 224, 216, 1, 0, 0, 0, 224, 217, 1, 0, 0, 0, 224,
		218, 1, 0, 0, 0, 224, 219, 1, 0, 0, 0, 224, 220, 1, 0, 0, 0, 224, 221,
		1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 223, 1, 0, 0, 0, 225, 31, 1, 0,
		0, 0, 226, 233, 3, 42, 21, 0, 227, 233, 3, 50, 25, 0, 228, 233, 3, 66,
		33, 0, 229, 233, 3, 74, 37, 0, 230, 233, 3, 82, 41, 0, 231, 233, 3, 90,
		45, 0, 232, 226, 1, 0, 0, 0, 232, 227, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0,
		232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233,
		33, 1, 0, 0, 0, 234, 235, 5, 42, 0, 0, 235, 35, 1, 0, 0, 0, 236, 245, 3,
		34, 17, 0, 237, 238, 5, 9, 0, 0, 238, 240, 3, 34, 17, 0, 239, 237, 1, 0,
		0, 0, 240, 241, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0,
		242, 246, 1, 0, 0, 0, 243, 244, 5, 9, 0, 0, 244, 246, 5, 5, 0, 0, 245,
		239, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 37, 1, 0, 0, 0, 247, 249, 7,
		1, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0,
		0, 250, 251, 7, 2, 0, 0, 251, 39, 1, 0, 0, 0, 252, 261, 3, 38, 19, 0, 253,
		254, 5, 9, 0, 0, 254, 256, 3, 38, 19, 0, 255, 253, 1, 0, 0, 0, 256, 257,
		1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 262, 1, 0,
		0, 0, 259, 260, 5, 9, 0, 0, 260, 262, 5, 5, 0, 0, 261, 255, 1, 0, 0, 0,
		261, 259, 1, 0, 0, 0, 262, 41, 1, 0, 0, 0, 263, 265, 5, 19, 0, 0, 264,
		266, 5, 12, 0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267,
		1, 0, 0, 0, 267, 268, 3, 38, 19, 0, 268, 270, 5, 20, 0, 0, 269, 271, 5,
		13, 0, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0,
		0, 272, 273, 3, 38, 19, 0, 273, 274, 5, 19, 0, 0, 274, 289, 1, 0, 0, 0,
		275, 277, 5, 19, 0, 0, 276, 278, 3, 98, 49, 0, 277, 276, 1, 0, 0, 0, 277,
		278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 3, 38, 19, 0, 280, 281,
		5, 19, 0, 0, 281, 289, 1, 0, 0, 0, 282, 283, 5, 19, 0, 0, 283, 284, 3,
		38, 19, 0, 284, 285, 5, 14, 0, 0, 285, 286, 3, 38, 19, 0, 286, 287, 5,
		19, 0, 0, 287, 289, 1, 0, 0, 0, 288, 263, 1, 0, 0, 0, 288, 275, 1, 0, 0,
		0, 288, 282, 1, 0, 0, 0, 289, 43, 1, 0, 0, 0, 290, 299, 3, 42, 21, 0, 291,
		292, 5, 9, 0, 0, 292, 294, 3, 42, 21, 0, 293, 291, 1, 0, 0, 0, 294, 295,
		1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 300, 1, 0,
		0, 0, 297, 298, 5, 9, 0, 0, 298, 300, 5, 5, 0, 0, 299, 293, 1, 0, 0, 0,
		299, 297, 1, 0, 0, 0, 300, 45, 1, 0, 0, 0, 301, 303, 7, 1, 0, 0, 302, 301,
		1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 7, 3,
		0, 0, 305, 47, 1, 0, 0, 0, 306, 315, 3, 46, 23, 0, 307, 308, 5, 9, 0, 0,
		308, 310, 3, 46, 23, 0, 309, 307, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311,
		309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 316, 1, 0, 0, 0, 313, 314,
		5, 9, 0, 0, 314, 316, 5, 5, 0, 0, 315, 309, 1, 0, 0, 0, 315, 313, 1, 0,
		0, 0, 316, 49, 1, 0, 0, 0, 317, 319, 5, 19, 0, 0, 318, 320, 5, 12, 0, 0,
		319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		322, 3, 46, 23, 0, 322, 324, 5, 20, 0, 0, 323, 325, 5, 13, 0, 0, 324, 323,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 3, 46,
		23, 0, 327, 328, 5, 19, 0, 0, 328, 343, 1, 0, 0, 0, 329, 331, 5, 19, 0,
		0, 330, 332, 3, 98, 49, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0,
		332, 333, 1, 0, 0, 0, 333, 334, 3, 46, 23, 0, 334, 335, 5, 19, 0, 0, 335,
		343, 1, 0, 0, 0, 336, 337, 5, 19, 0, 0, 337, 338, 3, 46, 23, 0, 338, 339,
		5, 14, 0, 0, 339, 340, 3, 46, 23, 0, 340, 341, 5, 19, 0, 0, 341, 343, 1,
		0, 0, 0, 342, 317, 1, 0, 0, 0, 342, 329, 1, 0, 0, 0, 342, 336, 1, 0, 0,
		0, 343, 51, 1, 0, 0, 0, 344, 353, 3, 50, 25, 0, 345, 346, 5, 9, 0, 0, 346,
		348, 3, 50, 25, 0, 347, 345, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347,
		1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 354, 1, 0, 0, 0, 351, 352, 5, 9,
		0, 0, 352, 354, 5, 5, 0, 0, 353, 347, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0,
		354, 53, 1, 0, 0, 0, 355, 356, 7, 4, 0, 0, 356, 55, 1, 0, 0, 0, 357, 366,
		3, 54, 27, 0, 358, 359, 5, 9, 0, 0, 359, 361, 3, 54, 27, 0, 360, 358, 1,
		0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0,
		0, 363, 367, 1, 0, 0, 0, 364, 365, 5, 9, 0, 0, 365, 367, 5, 5, 0, 0, 366,
		360, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 57, 1, 0, 0, 0, 368, 369, 5,
		43, 0, 0, 369, 59, 1, 0, 0, 0, 370, 379, 3, 58, 29, 0, 371, 372, 5, 9,
		0, 0, 372, 374, 3, 58, 29, 0, 373, 371, 1, 0, 0, 0, 374, 375, 1, 0, 0,
		0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 380, 1, 0, 0, 0, 377,
		378, 5, 9, 0, 0, 378, 380, 5, 5, 0, 0, 379, 373, 1, 0, 0, 0, 379, 377,
		1, 0, 0, 0, 380, 61, 1, 0, 0, 0, 381, 382, 5, 30, 0, 0, 382, 63, 1, 0,
		0, 0, 383, 392, 3, 62, 31, 0, 384, 385, 5, 9, 0, 0, 385, 387, 3, 62, 31,
		0, 386, 384, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 388,
		389, 1, 0, 0, 0, 389, 393, 1, 0, 0, 0, 390, 391, 5, 9, 0, 0, 391, 393,
		5, 5, 0, 0, 392, 386, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 65, 1, 0,
		0, 0, 394, 396, 5, 19, 0, 0, 395, 397, 5, 12, 0, 0, 396, 395, 1, 0, 0,
		0, 396, 397, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 3, 62, 31, 0,
		399, 401, 5, 20, 0, 0, 400, 402, 5, 13, 0, 0, 401, 400, 1, 0, 0, 0, 401,
		402, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 404, 3, 62, 31, 0, 404, 405,
		5, 19, 0, 0, 405, 420, 1, 0, 0, 0, 406, 408, 5, 19, 0, 0, 407, 409, 3,
		98, 49, 0, 408, 407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0,
		0, 0, 410, 411, 3, 62, 31, 0, 411, 412, 5, 19, 0, 0, 412, 420, 1, 0, 0,
		0, 413, 414, 5, 19, 0, 0, 414, 415, 3, 62, 31, 0, 415, 416, 5, 14, 0, 0,
		416, 417, 3, 86, 43, 0, 417, 418, 5, 19, 0, 0, 418, 420, 1, 0, 0, 0, 419,
		394, 1, 0, 0, 0, 419, 406, 1, 0, 0, 0, 419, 413, 1, 0, 0, 0, 420, 67, 1,
		0, 0, 0, 421, 430, 3, 66, 33, 0, 422, 423, 5, 9, 0, 0, 423, 425, 3, 66,
		33, 0, 424, 422, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0,
		426, 427, 1, 0, 0, 0, 427, 431, 1, 0, 0, 0, 428, 429, 5, 9, 0, 0, 429,
		431, 5, 5, 0, 0, 430, 424, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 69, 1,
		0, 0, 0, 432, 433, 5, 31, 0, 0, 433, 71, 1, 0, 0, 0, 434, 443, 3, 70, 35,
		0, 435, 436, 5, 9, 0, 0, 436, 438, 3, 70, 35, 0, 437, 435, 1, 0, 0, 0,
		438, 439, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440,
		444, 1, 0, 0, 0, 441, 442, 5, 9, 0, 0, 442, 444, 5, 5, 0, 0, 443, 437,
		1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 444, 73, 1, 0, 0, 0, 445, 447, 5, 19,
		0, 0, 446, 448, 5, 12, 0, 0, 447, 446, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0,
		448, 449, 1, 0, 0, 0, 449, 450, 3, 70, 35, 0, 450, 452, 5, 20, 0, 0, 451,
		453, 5, 13, 0, 0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454,
		1, 0, 0, 0, 454, 455, 3, 70, 35, 0, 455, 456, 5, 19, 0, 0, 456, 471, 1,
		0, 0, 0, 457, 459, 5, 19, 0, 0, 458, 460, 3, 98, 49, 0, 459, 458, 1, 0,
		0, 0, 459, 460, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461, 462, 3, 70, 35,
		0, 462, 463, 5, 19, 0, 0, 463, 471, 1, 0, 0, 0, 464, 465, 5, 19, 0, 0,
		465, 466, 3, 70, 35, 0, 466, 467, 5, 14, 0, 0, 467, 468, 3, 86, 43, 0,
		468, 469, 5, 19, 0, 0, 469, 471, 1, 0, 0, 0, 470, 445, 1, 0, 0, 0, 470,
		457, 1, 0, 0, 0, 470, 464, 1, 0, 0, 0, 471, 75, 1, 0, 0, 0, 472, 481, 3,
		74, 37, 0, 473, 474, 5, 9, 0, 0, 474, 476, 3, 74, 37, 0, 475, 473, 1, 0,
		0, 0, 476, 477, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0,
		478, 482, 1, 0, 0, 0, 479, 480, 5, 9, 0, 0, 480, 482, 5, 5, 0, 0, 481,
		475, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 77, 1, 0, 0, 0, 483, 484, 5,
		32, 0, 0, 484, 79, 1, 0, 0, 0, 485, 494, 3, 78, 39, 0, 486, 487, 5, 9,
		0, 0, 487, 489, 3, 78, 39, 0, 488, 486, 1, 0, 0, 0, 489, 490, 1, 0, 0,
		0, 490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 495, 1, 0, 0, 0, 492,
		493, 5, 9, 0, 0, 493, 495, 5, 5, 0, 0, 494, 488, 1, 0, 0, 0, 494, 492,
		1, 0, 0, 0, 495, 81, 1, 0, 0, 0, 496, 498, 5, 19, 0, 0, 497, 499, 5, 12,
		0, 0, 498, 497, 1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0,
		500, 501, 3, 78, 39, 0, 501, 503, 5, 20, 0, 0, 502, 504, 5, 13, 0, 0, 503,
		502, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 506,
		3, 78, 39, 0, 506, 507, 5, 19, 0, 0, 507, 522, 1, 0, 0, 0, 508, 510, 5,
		19, 0, 0, 509, 511, 3, 98, 49, 0, 510, 509, 1, 0, 0, 0, 510, 511, 1, 0,
		0, 0, 511, 512, 1, 0, 0, 0, 512, 513, 3, 78, 39, 0, 513, 514, 5, 19, 0,
		0, 514, 522, 1, 0, 0, 0, 515, 516, 5, 19, 0, 0, 516, 517, 3, 78, 39, 0,
		517, 518, 5, 14, 0, 0, 518, 519, 3, 86, 43, 0, 519, 520, 5, 19, 0, 0, 520,
		522, 1, 0, 0, 0, 521, 496, 1, 0, 0, 0, 521, 508, 1, 0, 0, 0, 521, 515,
		1, 0, 0, 0, 522, 83, 1, 0, 0, 0, 523, 532, 3, 82, 41, 0, 524, 525, 5, 9,
		0, 0, 525, 527, 3, 82, 41, 0, 526, 524, 1, 0, 0, 0, 527, 528, 1, 0, 0,
		0, 528, 526, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 533, 1, 0, 0, 0, 530,
		531, 5, 9, 0, 0, 531, 533, 5, 5, 0, 0, 532, 526, 1, 0, 0, 0, 532, 530,
		1, 0, 0, 0, 533, 85, 1, 0, 0, 0, 534, 536, 7, 1, 0, 0, 535, 534, 1, 0,
		0, 0, 535, 536, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537, 538, 5, 33, 0, 0,
		538, 87, 1, 0, 0, 0, 539, 548, 3, 86, 43, 0, 540, 541, 5, 9, 0, 0, 541,
		543, 3, 86, 43, 0, 542, 540, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 542,
		1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 549, 1, 0, 0, 0, 546, 547, 5, 9,
		0, 0, 547, 549, 5, 5, 0, 0, 548, 542, 1, 0, 0, 0, 548, 546, 1, 0, 0, 0,
		549, 89, 1, 0, 0, 0, 550, 552, 5, 19, 0, 0, 551, 553, 5, 12, 0, 0, 552,
		551, 1, 0, 0, 0, 552, 553, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 555,
		3, 86, 43, 0, 555, 557, 5, 20, 0, 0, 556, 558, 5, 13, 0, 0, 557, 556, 1,
		0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 560, 3, 86, 43,
		0, 560, 561, 5, 19, 0, 0, 561, 576, 1, 0, 0, 0, 562, 564, 5, 19, 0, 0,
		563, 565, 3, 98, 49, 0, 564, 563, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565,
		566, 1, 0, 0, 0, 566, 567, 3, 86, 43, 0, 567, 568, 5, 19, 0, 0, 568, 576,
		1, 0, 0, 0, 569, 570, 5, 19, 0, 0, 570, 571, 3, 86, 43, 0, 571, 572, 5,
		14, 0, 0, 572, 573, 3, 86, 43, 0, 573, 574, 5, 19, 0, 0, 574, 576, 1, 0,
		0, 0, 575, 550, 1, 0, 0, 0, 575, 562, 1, 0, 0, 0, 575, 569, 1, 0, 0, 0,
		576, 91, 1, 0, 0, 0, 577, 586, 3, 90, 45, 0, 578, 579, 5, 9, 0, 0, 579,
		581, 3, 90, 45, 0, 580, 578, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 580,
		1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 587, 1, 0, 0, 0, 584, 585, 5, 9,
		0, 0, 585, 587, 5, 5, 0, 0, 586, 580, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0,
		587, 93, 1, 0, 0, 0, 588, 589, 5, 54, 0, 0, 589, 95, 1, 0, 0, 0, 590, 599,
		3, 94, 47, 0, 591, 592, 5, 9, 0, 0, 592, 594, 3, 94, 47, 0, 593, 591, 1,
		0, 0, 0, 594, 595, 1, 0, 0, 0, 595, 593, 1, 0, 0, 0, 595, 596, 1, 0, 0,
		0, 596, 600, 1, 0, 0, 0, 597, 598, 5, 9, 0, 0, 598, 600, 5, 5, 0, 0, 599,
		593, 1, 0, 0, 0, 599, 597, 1, 0, 0, 0, 600, 97, 1, 0, 0, 0, 601, 602, 7,
		5, 0, 0, 602, 99, 1, 0, 0, 0, 79, 103, 106, 109, 115, 121, 124, 131, 136,
		139, 166, 168, 171, 176, 181, 190, 195, 200, 212, 224, 232, 241, 245, 248,
		257, 261, 265, 270, 277, 288, 295, 299, 302, 311, 315, 319, 324, 331, 342,
		349, 353, 362, 366, 375, 379, 388, 392, 396, 401, 408, 419, 426, 430, 439,
		443, 447, 452, 459, 470, 477, 481, 490, 494, 498, 503, 510, 521, 528, 532,
		535, 544, 548, 552, 557, 564, 575, 582, 586, 595, 599,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// OdinParserInit initializes any static state used to implement OdinParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOdinParser(). You can call this function if you wish to initialize the static state ahead
// of ISO8601.
func OdinParserInit() {
	staticData := &OdinParserParserStaticData
	staticData.once.Do(odinparserParserInit)
}

// NewOdinParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOdinParser(input antlr.TokenStream) *OdinParser {
	OdinParserInit()
	this := new(OdinParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &OdinParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "OdinParser.g4"

	return this
}

// OdinParser tokens.
const (
	OdinParserEOF                         = antlr.TokenEOF
	OdinParserCMT_LINE                    = 1
	OdinParserEOL                         = 2
	OdinParserWS                          = 3
	OdinParserODIN_URI                    = 4
	OdinParserSYM_LIST_CONTINUE           = 5
	OdinParserSYM_EQ                      = 6
	OdinParserSYM_SLASH                   = 7
	OdinParserSYM_NAMESPACE_SEP           = 8
	OdinParserSYM_COMMA                   = 9
	OdinParserSYM_LE                      = 10
	OdinParserSYM_GE                      = 11
	OdinParserSYM_GT                      = 12
	OdinParserSYM_LT                      = 13
	OdinParserSYM_PLUS_OR_MINUS           = 14
	OdinParserSYM_PLUS                    = 15
	OdinParserSYM_MINUS                   = 16
	OdinParserSYM_PERCENT                 = 17
	OdinParserSYM_CARET                   = 18
	OdinParserSYM_IVL_DELIM               = 19
	OdinParserSYM_IVL_SEP                 = 20
	OdinParserOBJECT_VERSION_ID           = 21
	OdinParserARCHETYPE_HRID              = 22
	OdinParserARCHETYPE_REF               = 23
	OdinParserVERSION_ID                  = 24
	OdinParserTERM_CODE_REF               = 25
	OdinParserROOT_ID_CODE                = 26
	OdinParserID_CODE                     = 27
	OdinParserAT_CODE                     = 28
	OdinParserAC_CODE                     = 29
	OdinParserISO8601_DATE_AUGMENTED      = 30
	OdinParserISO8601_TIME_AUGMENTED      = 31
	OdinParserISO8601_DATE_TIME_AUGMENTED = 32
	OdinParserISO8601_DURATION            = 33
	OdinParserSYM_TRUE                    = 34
	OdinParserSYM_FALSE                   = 35
	OdinParserGUID                        = 36
	OdinParserUUID                        = 37
	OdinParserINTEGER                     = 38
	OdinParserREAL                        = 39
	OdinParserSCI_INTEGER                 = 40
	OdinParserSCI_REAL                    = 41
	OdinParserSTRING                      = 42
	OdinParserCHARACTER                   = 43
	OdinParserSYM_DOT                     = 44
	OdinParserSYM_SEMI_COLON              = 45
	OdinParserSYM_LPAREN                  = 46
	OdinParserSYM_RPAREN                  = 47
	OdinParserSYM_LBRACKET                = 48
	OdinParserSYM_RBRACKET                = 49
	OdinParserSYM_LCURLY                  = 50
	OdinParserSYM_RCURLY                  = 51
	OdinParserUC_ID                       = 52
	OdinParserLC_ID                       = 53
	OdinParserQUALIFIED_TERM_CODE_REF     = 54
)

// OdinParser rules.
const (
	OdinParserRULE_odinObject                = 0
	OdinParserRULE_odinAttrVal               = 1
	OdinParserRULE_odinAttrName              = 2
	OdinParserRULE_odinObjectBlock           = 3
	OdinParserRULE_odinObjectValueBlock      = 4
	OdinParserRULE_rmTypeSpec                = 5
	OdinParserRULE_odinKeyedObject           = 6
	OdinParserRULE_odinKeySpec               = 7
	OdinParserRULE_odinObjectReferenceBlock  = 8
	OdinParserRULE_odinPathList              = 9
	OdinParserRULE_odinPath                  = 10
	OdinParserRULE_odinPathSegment           = 11
	OdinParserRULE_rmTypeId                  = 12
	OdinParserRULE_primitiveObject           = 13
	OdinParserRULE_primitiveValue            = 14
	OdinParserRULE_primitiveListValue        = 15
	OdinParserRULE_primitiveIntervalValue    = 16
	OdinParserRULE_stringValue               = 17
	OdinParserRULE_stringListValue           = 18
	OdinParserRULE_integerValue              = 19
	OdinParserRULE_integerListValue          = 20
	OdinParserRULE_integerIntervalValue      = 21
	OdinParserRULE_integerIntervalListValue  = 22
	OdinParserRULE_realValue                 = 23
	OdinParserRULE_realListValue             = 24
	OdinParserRULE_realIntervalValue         = 25
	OdinParserRULE_realIntervalListValue     = 26
	OdinParserRULE_booleanValue              = 27
	OdinParserRULE_booleanListValue          = 28
	OdinParserRULE_characterValue            = 29
	OdinParserRULE_characterListValue        = 30
	OdinParserRULE_dateValue                 = 31
	OdinParserRULE_dateListValue             = 32
	OdinParserRULE_dateIntervalValue         = 33
	OdinParserRULE_dateIntervalListValue     = 34
	OdinParserRULE_timeValue                 = 35
	OdinParserRULE_timeListValue             = 36
	OdinParserRULE_timeIntervalValue         = 37
	OdinParserRULE_timeIntervalListValue     = 38
	OdinParserRULE_dateTimeValue             = 39
	OdinParserRULE_dateTimeListValue         = 40
	OdinParserRULE_dateTimeIntervalValue     = 41
	OdinParserRULE_dateTimeIntervalListValue = 42
	OdinParserRULE_durationValue             = 43
	OdinParserRULE_durationListValue         = 44
	OdinParserRULE_durationIntervalValue     = 45
	OdinParserRULE_durationIntervalListValue = 46
	OdinParserRULE_termCodeValue             = 47
	OdinParserRULE_termCodeListValue         = 48
	OdinParserRULE_relop                     = 49
)

// IOdinObjectContext is an interface to support dynamic dispatch.
type IOdinObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OdinObjectValueBlock() IOdinObjectValueBlockContext
	EOF() antlr.TerminalNode
	AllOdinAttrVal() []IOdinAttrValContext
	OdinAttrVal(i int) IOdinAttrValContext

	// IsOdinObjectContext differentiates from other interfaces.
	IsOdinObjectContext()
}

type OdinObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectContext() *OdinObjectContext {
	var p = new(OdinObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObject
	return p
}

func InitEmptyOdinObjectContext(p *OdinObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObject
}

func (*OdinObjectContext) IsOdinObjectContext() {}

func NewOdinObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectContext {
	var p = new(OdinObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObject

	return p
}

func (s *OdinObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectContext) OdinObjectValueBlock() IOdinObjectValueBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinObjectValueBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinObjectValueBlockContext)
}

func (s *OdinObjectContext) EOF() antlr.TerminalNode {
	return s.GetToken(OdinParserEOF, 0)
}

func (s *OdinObjectContext) AllOdinAttrVal() []IOdinAttrValContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOdinAttrValContext); ok {
			len++
		}
	}

	tst := make([]IOdinAttrValContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOdinAttrValContext); ok {
			tst[i] = t.(IOdinAttrValContext)
			i++
		}
	}

	return tst
}

func (s *OdinObjectContext) OdinAttrVal(i int) IOdinAttrValContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinAttrValContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinAttrValContext)
}

func (s *OdinObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinObject(s)
	}
}

func (s *OdinObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinObject(s)
	}
}

func (p *OdinParser) OdinObject() (localctx IOdinObjectContext) {
	localctx = NewOdinObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OdinParserRULE_odinObject)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case OdinParserUC_ID, OdinParserLC_ID:
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserUC_ID || _la == OdinParserLC_ID {
			{
				p.SetState(100)
				p.OdinAttrVal()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserSYM_LT, OdinParserSYM_LPAREN:
		{
			p.SetState(105)
			p.OdinObjectValueBlock()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(108)
			p.Match(OdinParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinAttrValContext is an interface to support dynamic dispatch.
type IOdinAttrValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OdinAttrName() IOdinAttrNameContext
	SYM_EQ() antlr.TerminalNode
	OdinObjectBlock() IOdinObjectBlockContext
	SYM_SEMI_COLON() antlr.TerminalNode

	// IsOdinAttrValContext differentiates from other interfaces.
	IsOdinAttrValContext()
}

type OdinAttrValContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinAttrValContext() *OdinAttrValContext {
	var p = new(OdinAttrValContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrVal
	return p
}

func InitEmptyOdinAttrValContext(p *OdinAttrValContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrVal
}

func (*OdinAttrValContext) IsOdinAttrValContext() {}

func NewOdinAttrValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinAttrValContext {
	var p = new(OdinAttrValContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinAttrVal

	return p
}

func (s *OdinAttrValContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinAttrValContext) OdinAttrName() IOdinAttrNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinAttrNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinAttrNameContext)
}

func (s *OdinAttrValContext) SYM_EQ() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_EQ, 0)
}

func (s *OdinAttrValContext) OdinObjectBlock() IOdinObjectBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinObjectBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinObjectBlockContext)
}

func (s *OdinAttrValContext) SYM_SEMI_COLON() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_SEMI_COLON, 0)
}

func (s *OdinAttrValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinAttrValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinAttrValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinAttrVal(s)
	}
}

func (s *OdinAttrValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinAttrVal(s)
	}
}

func (p *OdinParser) OdinAttrVal() (localctx IOdinAttrValContext) {
	localctx = NewOdinAttrValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OdinParserRULE_odinAttrVal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.OdinAttrName()
	}
	{
		p.SetState(112)
		p.Match(OdinParserSYM_EQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.OdinObjectBlock()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_SEMI_COLON {
		{
			p.SetState(114)
			p.Match(OdinParserSYM_SEMI_COLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinAttrNameContext is an interface to support dynamic dispatch.
type IOdinAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UC_ID() antlr.TerminalNode
	LC_ID() antlr.TerminalNode

	// IsOdinAttrNameContext differentiates from other interfaces.
	IsOdinAttrNameContext()
}

type OdinAttrNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinAttrNameContext() *OdinAttrNameContext {
	var p = new(OdinAttrNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrName
	return p
}

func InitEmptyOdinAttrNameContext(p *OdinAttrNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrName
}

func (*OdinAttrNameContext) IsOdinAttrNameContext() {}

func NewOdinAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinAttrNameContext {
	var p = new(OdinAttrNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinAttrName

	return p
}

func (s *OdinAttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinAttrNameContext) UC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserUC_ID, 0)
}

func (s *OdinAttrNameContext) LC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserLC_ID, 0)
}

func (s *OdinAttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinAttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinAttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinAttrName(s)
	}
}

func (s *OdinAttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinAttrName(s)
	}
}

func (p *OdinParser) OdinAttrName() (localctx IOdinAttrNameContext) {
	localctx = NewOdinAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OdinParserRULE_odinAttrName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserUC_ID || _la == OdinParserLC_ID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinObjectBlockContext is an interface to support dynamic dispatch.
type IOdinObjectBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OdinObjectValueBlock() IOdinObjectValueBlockContext
	OdinObjectReferenceBlock() IOdinObjectReferenceBlockContext

	// IsOdinObjectBlockContext differentiates from other interfaces.
	IsOdinObjectBlockContext()
}

type OdinObjectBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectBlockContext() *OdinObjectBlockContext {
	var p = new(OdinObjectBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectBlock
	return p
}

func InitEmptyOdinObjectBlockContext(p *OdinObjectBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectBlock
}

func (*OdinObjectBlockContext) IsOdinObjectBlockContext() {}

func NewOdinObjectBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectBlockContext {
	var p = new(OdinObjectBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectBlock

	return p
}

func (s *OdinObjectBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectBlockContext) OdinObjectValueBlock() IOdinObjectValueBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinObjectValueBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinObjectValueBlockContext)
}

func (s *OdinObjectBlockContext) OdinObjectReferenceBlock() IOdinObjectReferenceBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinObjectReferenceBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinObjectReferenceBlockContext)
}

func (s *OdinObjectBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinObjectBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinObjectBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinObjectBlock(s)
	}
}

func (s *OdinObjectBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinObjectBlock(s)
	}
}

func (p *OdinParser) OdinObjectBlock() (localctx IOdinObjectBlockContext) {
	localctx = NewOdinObjectBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, OdinParserRULE_odinObjectBlock)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.OdinObjectValueBlock()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.OdinObjectReferenceBlock()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinObjectValueBlockContext is an interface to support dynamic dispatch.
type IOdinObjectValueBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_LT() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	RmTypeSpec() IRmTypeSpecContext
	PrimitiveObject() IPrimitiveObjectContext
	ODIN_URI() antlr.TerminalNode
	AllOdinAttrVal() []IOdinAttrValContext
	OdinAttrVal(i int) IOdinAttrValContext
	AllOdinKeyedObject() []IOdinKeyedObjectContext
	OdinKeyedObject(i int) IOdinKeyedObjectContext

	// IsOdinObjectValueBlockContext differentiates from other interfaces.
	IsOdinObjectValueBlockContext()
}

type OdinObjectValueBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectValueBlockContext() *OdinObjectValueBlockContext {
	var p = new(OdinObjectValueBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectValueBlock
	return p
}

func InitEmptyOdinObjectValueBlockContext(p *OdinObjectValueBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectValueBlock
}

func (*OdinObjectValueBlockContext) IsOdinObjectValueBlockContext() {}

func NewOdinObjectValueBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectValueBlockContext {
	var p = new(OdinObjectValueBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectValueBlock

	return p
}

func (s *OdinObjectValueBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectValueBlockContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *OdinObjectValueBlockContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *OdinObjectValueBlockContext) RmTypeSpec() IRmTypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmTypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmTypeSpecContext)
}

func (s *OdinObjectValueBlockContext) PrimitiveObject() IPrimitiveObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveObjectContext)
}

func (s *OdinObjectValueBlockContext) ODIN_URI() antlr.TerminalNode {
	return s.GetToken(OdinParserODIN_URI, 0)
}

func (s *OdinObjectValueBlockContext) AllOdinAttrVal() []IOdinAttrValContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOdinAttrValContext); ok {
			len++
		}
	}

	tst := make([]IOdinAttrValContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOdinAttrValContext); ok {
			tst[i] = t.(IOdinAttrValContext)
			i++
		}
	}

	return tst
}

func (s *OdinObjectValueBlockContext) OdinAttrVal(i int) IOdinAttrValContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinAttrValContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinAttrValContext)
}

func (s *OdinObjectValueBlockContext) AllOdinKeyedObject() []IOdinKeyedObjectContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOdinKeyedObjectContext); ok {
			len++
		}
	}

	tst := make([]IOdinKeyedObjectContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOdinKeyedObjectContext); ok {
			tst[i] = t.(IOdinKeyedObjectContext)
			i++
		}
	}

	return tst
}

func (s *OdinObjectValueBlockContext) OdinKeyedObject(i int) IOdinKeyedObjectContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinKeyedObjectContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinKeyedObjectContext)
}

func (s *OdinObjectValueBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinObjectValueBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinObjectValueBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinObjectValueBlock(s)
	}
}

func (s *OdinObjectValueBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinObjectValueBlock(s)
	}
}

func (p *OdinParser) OdinObjectValueBlock() (localctx IOdinObjectValueBlockContext) {
	localctx = NewOdinObjectValueBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, OdinParserRULE_odinObjectValueBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LPAREN {
		{
			p.SetState(123)
			p.RmTypeSpec()
		}

	}
	{
		p.SetState(126)
		p.Match(OdinParserSYM_LT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case OdinParserSYM_PLUS, OdinParserSYM_MINUS, OdinParserSYM_IVL_DELIM, OdinParserISO8601_DATE_AUGMENTED, OdinParserISO8601_TIME_AUGMENTED, OdinParserISO8601_DATE_TIME_AUGMENTED, OdinParserISO8601_DURATION, OdinParserSYM_TRUE, OdinParserSYM_FALSE, OdinParserINTEGER, OdinParserREAL, OdinParserSCI_INTEGER, OdinParserSCI_REAL, OdinParserSTRING, OdinParserCHARACTER, OdinParserQUALIFIED_TERM_CODE_REF:
		{
			p.SetState(127)
			p.PrimitiveObject()
		}

	case OdinParserUC_ID, OdinParserLC_ID:
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserUC_ID || _la == OdinParserLC_ID {
			{
				p.SetState(128)
				p.OdinAttrVal()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserSYM_LBRACKET:
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_LBRACKET {
			{
				p.SetState(133)
				p.OdinKeyedObject()
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserODIN_URI:
		{
			p.SetState(138)
			p.Match(OdinParserODIN_URI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case OdinParserSYM_GT:

	default:
	}
	{
		p.SetState(141)
		p.Match(OdinParserSYM_GT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmTypeSpecContext is an interface to support dynamic dispatch.
type IRmTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_LPAREN() antlr.TerminalNode
	RmTypeId() IRmTypeIdContext
	SYM_RPAREN() antlr.TerminalNode

	// IsRmTypeSpecContext differentiates from other interfaces.
	IsRmTypeSpecContext()
}

type RmTypeSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmTypeSpecContext() *RmTypeSpecContext {
	var p = new(RmTypeSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeSpec
	return p
}

func InitEmptyRmTypeSpecContext(p *RmTypeSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeSpec
}

func (*RmTypeSpecContext) IsRmTypeSpecContext() {}

func NewRmTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmTypeSpecContext {
	var p = new(RmTypeSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_rmTypeSpec

	return p
}

func (s *RmTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *RmTypeSpecContext) SYM_LPAREN() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LPAREN, 0)
}

func (s *RmTypeSpecContext) RmTypeId() IRmTypeIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmTypeIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmTypeIdContext)
}

func (s *RmTypeSpecContext) SYM_RPAREN() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_RPAREN, 0)
}

func (s *RmTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRmTypeSpec(s)
	}
}

func (s *RmTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRmTypeSpec(s)
	}
}

func (p *OdinParser) RmTypeSpec() (localctx IRmTypeSpecContext) {
	localctx = NewRmTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OdinParserRULE_rmTypeSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(OdinParserSYM_LPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.RmTypeId()
	}
	{
		p.SetState(145)
		p.Match(OdinParserSYM_RPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinKeyedObjectContext is an interface to support dynamic dispatch.
type IOdinKeyedObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OdinKeySpec() IOdinKeySpecContext
	SYM_EQ() antlr.TerminalNode
	OdinObjectBlock() IOdinObjectBlockContext

	// IsOdinKeyedObjectContext differentiates from other interfaces.
	IsOdinKeyedObjectContext()
}

type OdinKeyedObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinKeyedObjectContext() *OdinKeyedObjectContext {
	var p = new(OdinKeyedObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeyedObject
	return p
}

func InitEmptyOdinKeyedObjectContext(p *OdinKeyedObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeyedObject
}

func (*OdinKeyedObjectContext) IsOdinKeyedObjectContext() {}

func NewOdinKeyedObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinKeyedObjectContext {
	var p = new(OdinKeyedObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinKeyedObject

	return p
}

func (s *OdinKeyedObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinKeyedObjectContext) OdinKeySpec() IOdinKeySpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinKeySpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinKeySpecContext)
}

func (s *OdinKeyedObjectContext) SYM_EQ() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_EQ, 0)
}

func (s *OdinKeyedObjectContext) OdinObjectBlock() IOdinObjectBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinObjectBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinObjectBlockContext)
}

func (s *OdinKeyedObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinKeyedObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinKeyedObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinKeyedObject(s)
	}
}

func (s *OdinKeyedObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinKeyedObject(s)
	}
}

func (p *OdinParser) OdinKeyedObject() (localctx IOdinKeyedObjectContext) {
	localctx = NewOdinKeyedObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, OdinParserRULE_odinKeyedObject)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.OdinKeySpec()
	}
	{
		p.SetState(148)
		p.Match(OdinParserSYM_EQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.OdinObjectBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinKeySpecContext is an interface to support dynamic dispatch.
type IOdinKeySpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_LBRACKET() antlr.TerminalNode
	PrimitiveValue() IPrimitiveValueContext
	SYM_RBRACKET() antlr.TerminalNode

	// IsOdinKeySpecContext differentiates from other interfaces.
	IsOdinKeySpecContext()
}

type OdinKeySpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinKeySpecContext() *OdinKeySpecContext {
	var p = new(OdinKeySpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeySpec
	return p
}

func InitEmptyOdinKeySpecContext(p *OdinKeySpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeySpec
}

func (*OdinKeySpecContext) IsOdinKeySpecContext() {}

func NewOdinKeySpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinKeySpecContext {
	var p = new(OdinKeySpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinKeySpec

	return p
}

func (s *OdinKeySpecContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinKeySpecContext) SYM_LBRACKET() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LBRACKET, 0)
}

func (s *OdinKeySpecContext) PrimitiveValue() IPrimitiveValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
}

func (s *OdinKeySpecContext) SYM_RBRACKET() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_RBRACKET, 0)
}

func (s *OdinKeySpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinKeySpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinKeySpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinKeySpec(s)
	}
}

func (s *OdinKeySpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinKeySpec(s)
	}
}

func (p *OdinParser) OdinKeySpec() (localctx IOdinKeySpecContext) {
	localctx = NewOdinKeySpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, OdinParserRULE_odinKeySpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(OdinParserSYM_LBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.PrimitiveValue()
	}
	{
		p.SetState(153)
		p.Match(OdinParserSYM_RBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinObjectReferenceBlockContext is an interface to support dynamic dispatch.
type IOdinObjectReferenceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_LT() antlr.TerminalNode
	OdinPathList() IOdinPathListContext
	SYM_GT() antlr.TerminalNode

	// IsOdinObjectReferenceBlockContext differentiates from other interfaces.
	IsOdinObjectReferenceBlockContext()
}

type OdinObjectReferenceBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectReferenceBlockContext() *OdinObjectReferenceBlockContext {
	var p = new(OdinObjectReferenceBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectReferenceBlock
	return p
}

func InitEmptyOdinObjectReferenceBlockContext(p *OdinObjectReferenceBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectReferenceBlock
}

func (*OdinObjectReferenceBlockContext) IsOdinObjectReferenceBlockContext() {}

func NewOdinObjectReferenceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectReferenceBlockContext {
	var p = new(OdinObjectReferenceBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectReferenceBlock

	return p
}

func (s *OdinObjectReferenceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectReferenceBlockContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *OdinObjectReferenceBlockContext) OdinPathList() IOdinPathListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinPathListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinPathListContext)
}

func (s *OdinObjectReferenceBlockContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *OdinObjectReferenceBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinObjectReferenceBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinObjectReferenceBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinObjectReferenceBlock(s)
	}
}

func (s *OdinObjectReferenceBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinObjectReferenceBlock(s)
	}
}

func (p *OdinParser) OdinObjectReferenceBlock() (localctx IOdinObjectReferenceBlockContext) {
	localctx = NewOdinObjectReferenceBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OdinParserRULE_odinObjectReferenceBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(OdinParserSYM_LT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.OdinPathList()
	}
	{
		p.SetState(157)
		p.Match(OdinParserSYM_GT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinPathListContext is an interface to support dynamic dispatch.
type IOdinPathListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOdinPath() []IOdinPathContext
	OdinPath(i int) IOdinPathContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsOdinPathListContext differentiates from other interfaces.
	IsOdinPathListContext()
}

type OdinPathListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathListContext() *OdinPathListContext {
	var p = new(OdinPathListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathList
	return p
}

func InitEmptyOdinPathListContext(p *OdinPathListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathList
}

func (*OdinPathListContext) IsOdinPathListContext() {}

func NewOdinPathListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathListContext {
	var p = new(OdinPathListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPathList

	return p
}

func (s *OdinPathListContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathListContext) AllOdinPath() []IOdinPathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOdinPathContext); ok {
			len++
		}
	}

	tst := make([]IOdinPathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOdinPathContext); ok {
			tst[i] = t.(IOdinPathContext)
			i++
		}
	}

	return tst
}

func (s *OdinPathListContext) OdinPath(i int) IOdinPathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinPathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinPathContext)
}

func (s *OdinPathListContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *OdinPathListContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *OdinPathListContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *OdinPathListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinPathListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinPathListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinPathList(s)
	}
}

func (s *OdinPathListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinPathList(s)
	}
}

func (p *OdinParser) OdinPathList() (localctx IOdinPathListContext) {
	localctx = NewOdinPathListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, OdinParserRULE_odinPathList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.OdinPath()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 2 {
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(162)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(163)
				p.OdinPath()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinPathContext is an interface to support dynamic dispatch.
type IOdinPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OdinKeySpec() IOdinKeySpecContext
	AllOdinPathSegment() []IOdinPathSegmentContext
	OdinPathSegment(i int) IOdinPathSegmentContext

	// IsOdinPathContext differentiates from other interfaces.
	IsOdinPathContext()
}

type OdinPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathContext() *OdinPathContext {
	var p = new(OdinPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPath
	return p
}

func InitEmptyOdinPathContext(p *OdinPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPath
}

func (*OdinPathContext) IsOdinPathContext() {}

func NewOdinPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathContext {
	var p = new(OdinPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPath

	return p
}

func (s *OdinPathContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathContext) OdinKeySpec() IOdinKeySpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinKeySpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinKeySpecContext)
}

func (s *OdinPathContext) AllOdinPathSegment() []IOdinPathSegmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOdinPathSegmentContext); ok {
			len++
		}
	}

	tst := make([]IOdinPathSegmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOdinPathSegmentContext); ok {
			tst[i] = t.(IOdinPathSegmentContext)
			i++
		}
	}

	return tst
}

func (s *OdinPathContext) OdinPathSegment(i int) IOdinPathSegmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinPathSegmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinPathSegmentContext)
}

func (s *OdinPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinPath(s)
	}
}

func (s *OdinPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinPath(s)
	}
}

func (p *OdinParser) OdinPath() (localctx IOdinPathContext) {
	localctx = NewOdinPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, OdinParserRULE_odinPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LBRACKET {
		{
			p.SetState(170)
			p.OdinKeySpec()
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OdinParserSYM_SLASH {
		{
			p.SetState(173)
			p.OdinPathSegment()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOdinPathSegmentContext is an interface to support dynamic dispatch.
type IOdinPathSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_SLASH() antlr.TerminalNode
	LC_ID() antlr.TerminalNode
	OdinKeySpec() IOdinKeySpecContext

	// IsOdinPathSegmentContext differentiates from other interfaces.
	IsOdinPathSegmentContext()
}

type OdinPathSegmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathSegmentContext() *OdinPathSegmentContext {
	var p = new(OdinPathSegmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathSegment
	return p
}

func InitEmptyOdinPathSegmentContext(p *OdinPathSegmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathSegment
}

func (*OdinPathSegmentContext) IsOdinPathSegmentContext() {}

func NewOdinPathSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathSegmentContext {
	var p = new(OdinPathSegmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPathSegment

	return p
}

func (s *OdinPathSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathSegmentContext) SYM_SLASH() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_SLASH, 0)
}

func (s *OdinPathSegmentContext) LC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserLC_ID, 0)
}

func (s *OdinPathSegmentContext) OdinKeySpec() IOdinKeySpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdinKeySpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdinKeySpecContext)
}

func (s *OdinPathSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdinPathSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdinPathSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterOdinPathSegment(s)
	}
}

func (s *OdinPathSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitOdinPathSegment(s)
	}
}

func (p *OdinParser) OdinPathSegment() (localctx IOdinPathSegmentContext) {
	localctx = NewOdinPathSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OdinParserRULE_odinPathSegment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(OdinParserSYM_SLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(OdinParserLC_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LBRACKET {
		{
			p.SetState(180)
			p.OdinKeySpec()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmTypeIdContext is an interface to support dynamic dispatch.
type IRmTypeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UC_ID() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	AllRmTypeId() []IRmTypeIdContext
	RmTypeId(i int) IRmTypeIdContext
	SYM_GT() antlr.TerminalNode
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode

	// IsRmTypeIdContext differentiates from other interfaces.
	IsRmTypeIdContext()
}

type RmTypeIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmTypeIdContext() *RmTypeIdContext {
	var p = new(RmTypeIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeId
	return p
}

func InitEmptyRmTypeIdContext(p *RmTypeIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeId
}

func (*RmTypeIdContext) IsRmTypeIdContext() {}

func NewRmTypeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmTypeIdContext {
	var p = new(RmTypeIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_rmTypeId

	return p
}

func (s *RmTypeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *RmTypeIdContext) UC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserUC_ID, 0)
}

func (s *RmTypeIdContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *RmTypeIdContext) AllRmTypeId() []IRmTypeIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRmTypeIdContext); ok {
			len++
		}
	}

	tst := make([]IRmTypeIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRmTypeIdContext); ok {
			tst[i] = t.(IRmTypeIdContext)
			i++
		}
	}

	return tst
}

func (s *RmTypeIdContext) RmTypeId(i int) IRmTypeIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmTypeIdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmTypeIdContext)
}

func (s *RmTypeIdContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *RmTypeIdContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *RmTypeIdContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *RmTypeIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmTypeIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmTypeIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRmTypeId(s)
	}
}

func (s *RmTypeIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRmTypeId(s)
	}
}

func (p *OdinParser) RmTypeId() (localctx IRmTypeIdContext) {
	localctx = NewRmTypeIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OdinParserRULE_rmTypeId)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(OdinParserUC_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LT {
		{
			p.SetState(184)
			p.Match(OdinParserSYM_LT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.RmTypeId()
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == OdinParserSYM_COMMA {
			{
				p.SetState(186)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(187)
				p.RmTypeId()
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(193)
			p.Match(OdinParserSYM_GT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveObjectContext is an interface to support dynamic dispatch.
type IPrimitiveObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimitiveValue() IPrimitiveValueContext
	PrimitiveListValue() IPrimitiveListValueContext
	PrimitiveIntervalValue() IPrimitiveIntervalValueContext

	// IsPrimitiveObjectContext differentiates from other interfaces.
	IsPrimitiveObjectContext()
}

type PrimitiveObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveObjectContext() *PrimitiveObjectContext {
	var p = new(PrimitiveObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveObject
	return p
}

func InitEmptyPrimitiveObjectContext(p *PrimitiveObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveObject
}

func (*PrimitiveObjectContext) IsPrimitiveObjectContext() {}

func NewPrimitiveObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveObjectContext {
	var p = new(PrimitiveObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveObject

	return p
}

func (s *PrimitiveObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveObjectContext) PrimitiveValue() IPrimitiveValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
}

func (s *PrimitiveObjectContext) PrimitiveListValue() IPrimitiveListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveListValueContext)
}

func (s *PrimitiveObjectContext) PrimitiveIntervalValue() IPrimitiveIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveIntervalValueContext)
}

func (s *PrimitiveObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterPrimitiveObject(s)
	}
}

func (s *PrimitiveObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitPrimitiveObject(s)
	}
}

func (p *OdinParser) PrimitiveObject() (localctx IPrimitiveObjectContext) {
	localctx = NewPrimitiveObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, OdinParserRULE_primitiveObject)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.PrimitiveValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.PrimitiveListValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)
			p.PrimitiveIntervalValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveValueContext is an interface to support dynamic dispatch.
type IPrimitiveValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringValue() IStringValueContext
	IntegerValue() IIntegerValueContext
	RealValue() IRealValueContext
	BooleanValue() IBooleanValueContext
	CharacterValue() ICharacterValueContext
	TermCodeValue() ITermCodeValueContext
	DateValue() IDateValueContext
	TimeValue() ITimeValueContext
	DateTimeValue() IDateTimeValueContext
	DurationValue() IDurationValueContext

	// IsPrimitiveValueContext differentiates from other interfaces.
	IsPrimitiveValueContext()
}

type PrimitiveValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveValueContext() *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveValue
	return p
}

func InitEmptyPrimitiveValueContext(p *PrimitiveValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveValue
}

func (*PrimitiveValueContext) IsPrimitiveValueContext() {}

func NewPrimitiveValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveValue

	return p
}

func (s *PrimitiveValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveValueContext) StringValue() IStringValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *PrimitiveValueContext) IntegerValue() IIntegerValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *PrimitiveValueContext) RealValue() IRealValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
}

func (s *PrimitiveValueContext) BooleanValue() IBooleanValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *PrimitiveValueContext) CharacterValue() ICharacterValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterValueContext)
}

func (s *PrimitiveValueContext) TermCodeValue() ITermCodeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermCodeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermCodeValueContext)
}

func (s *PrimitiveValueContext) DateValue() IDateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *PrimitiveValueContext) TimeValue() ITimeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *PrimitiveValueContext) DateTimeValue() IDateTimeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
}

func (s *PrimitiveValueContext) DurationValue() IDurationValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *PrimitiveValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterPrimitiveValue(s)
	}
}

func (s *PrimitiveValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitPrimitiveValue(s)
	}
}

func (p *OdinParser) PrimitiveValue() (localctx IPrimitiveValueContext) {
	localctx = NewPrimitiveValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, OdinParserRULE_primitiveValue)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.StringValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.IntegerValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(204)
			p.RealValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(205)
			p.BooleanValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(206)
			p.CharacterValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(207)
			p.TermCodeValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(208)
			p.DateValue()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(209)
			p.TimeValue()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(210)
			p.DateTimeValue()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(211)
			p.DurationValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveListValueContext is an interface to support dynamic dispatch.
type IPrimitiveListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringListValue() IStringListValueContext
	IntegerListValue() IIntegerListValueContext
	RealListValue() IRealListValueContext
	BooleanListValue() IBooleanListValueContext
	CharacterListValue() ICharacterListValueContext
	TermCodeListValue() ITermCodeListValueContext
	DateListValue() IDateListValueContext
	TimeListValue() ITimeListValueContext
	DateTimeListValue() IDateTimeListValueContext
	DurationListValue() IDurationListValueContext

	// IsPrimitiveListValueContext differentiates from other interfaces.
	IsPrimitiveListValueContext()
}

type PrimitiveListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveListValueContext() *PrimitiveListValueContext {
	var p = new(PrimitiveListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveListValue
	return p
}

func InitEmptyPrimitiveListValueContext(p *PrimitiveListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveListValue
}

func (*PrimitiveListValueContext) IsPrimitiveListValueContext() {}

func NewPrimitiveListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveListValueContext {
	var p = new(PrimitiveListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveListValue

	return p
}

func (s *PrimitiveListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveListValueContext) StringListValue() IStringListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringListValueContext)
}

func (s *PrimitiveListValueContext) IntegerListValue() IIntegerListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerListValueContext)
}

func (s *PrimitiveListValueContext) RealListValue() IRealListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealListValueContext)
}

func (s *PrimitiveListValueContext) BooleanListValue() IBooleanListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanListValueContext)
}

func (s *PrimitiveListValueContext) CharacterListValue() ICharacterListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterListValueContext)
}

func (s *PrimitiveListValueContext) TermCodeListValue() ITermCodeListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermCodeListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermCodeListValueContext)
}

func (s *PrimitiveListValueContext) DateListValue() IDateListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateListValueContext)
}

func (s *PrimitiveListValueContext) TimeListValue() ITimeListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeListValueContext)
}

func (s *PrimitiveListValueContext) DateTimeListValue() IDateTimeListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeListValueContext)
}

func (s *PrimitiveListValueContext) DurationListValue() IDurationListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationListValueContext)
}

func (s *PrimitiveListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterPrimitiveListValue(s)
	}
}

func (s *PrimitiveListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitPrimitiveListValue(s)
	}
}

func (p *OdinParser) PrimitiveListValue() (localctx IPrimitiveListValueContext) {
	localctx = NewPrimitiveListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, OdinParserRULE_primitiveListValue)
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.StringListValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.IntegerListValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.RealListValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.BooleanListValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.CharacterListValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(219)
			p.TermCodeListValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(220)
			p.DateListValue()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(221)
			p.TimeListValue()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(222)
			p.DateTimeListValue()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(223)
			p.DurationListValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitiveIntervalValueContext is an interface to support dynamic dispatch.
type IPrimitiveIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerIntervalValue() IIntegerIntervalValueContext
	RealIntervalValue() IRealIntervalValueContext
	DateIntervalValue() IDateIntervalValueContext
	TimeIntervalValue() ITimeIntervalValueContext
	DateTimeIntervalValue() IDateTimeIntervalValueContext
	DurationIntervalValue() IDurationIntervalValueContext

	// IsPrimitiveIntervalValueContext differentiates from other interfaces.
	IsPrimitiveIntervalValueContext()
}

type PrimitiveIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveIntervalValueContext() *PrimitiveIntervalValueContext {
	var p = new(PrimitiveIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveIntervalValue
	return p
}

func InitEmptyPrimitiveIntervalValueContext(p *PrimitiveIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveIntervalValue
}

func (*PrimitiveIntervalValueContext) IsPrimitiveIntervalValueContext() {}

func NewPrimitiveIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveIntervalValueContext {
	var p = new(PrimitiveIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveIntervalValue

	return p
}

func (s *PrimitiveIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveIntervalValueContext) IntegerIntervalValue() IIntegerIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) RealIntervalValue() IRealIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DateIntervalValue() IDateIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) TimeIntervalValue() ITimeIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DateTimeIntervalValue() IDateTimeIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DurationIntervalValue() IDurationIntervalValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationIntervalValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterPrimitiveIntervalValue(s)
	}
}

func (s *PrimitiveIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitPrimitiveIntervalValue(s)
	}
}

func (p *OdinParser) PrimitiveIntervalValue() (localctx IPrimitiveIntervalValueContext) {
	localctx = NewPrimitiveIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, OdinParserRULE_primitiveIntervalValue)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.IntegerIntervalValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.RealIntervalValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			p.DateIntervalValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			p.TimeIntervalValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.DateTimeIntervalValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			p.DurationIntervalValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_stringValue
	return p
}

func InitEmptyStringValueContext(p *StringValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_stringValue
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(OdinParserSTRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *OdinParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, OdinParserRULE_stringValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(OdinParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringListValueContext is an interface to support dynamic dispatch.
type IStringListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStringValue() []IStringValueContext
	StringValue(i int) IStringValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsStringListValueContext differentiates from other interfaces.
	IsStringListValueContext()
}

type StringListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListValueContext() *StringListValueContext {
	var p = new(StringListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_stringListValue
	return p
}

func InitEmptyStringListValueContext(p *StringListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_stringListValue
}

func (*StringListValueContext) IsStringListValueContext() {}

func NewStringListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListValueContext {
	var p = new(StringListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_stringListValue

	return p
}

func (s *StringListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListValueContext) AllStringValue() []IStringValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringValueContext); ok {
			len++
		}
	}

	tst := make([]IStringValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringValueContext); ok {
			tst[i] = t.(IStringValueContext)
			i++
		}
	}

	return tst
}

func (s *StringListValueContext) StringValue(i int) IStringValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *StringListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *StringListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *StringListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *StringListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterStringListValue(s)
	}
}

func (s *StringListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitStringListValue(s)
	}
}

func (p *OdinParser) StringListValue() (localctx IStringListValueContext) {
	localctx = NewStringListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, OdinParserRULE_stringListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.StringValue()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(237)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)
				p.StringValue()
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(243)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	SCI_INTEGER() antlr.TerminalNode
	SYM_PLUS() antlr.TerminalNode
	SYM_MINUS() antlr.TerminalNode

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerValue
	return p
}

func InitEmptyIntegerValueContext(p *IntegerValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerValue
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerValue

	return p
}

func (s *IntegerValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(OdinParserINTEGER, 0)
}

func (s *IntegerValueContext) SCI_INTEGER() antlr.TerminalNode {
	return s.GetToken(OdinParserSCI_INTEGER, 0)
}

func (s *IntegerValueContext) SYM_PLUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS, 0)
}

func (s *IntegerValueContext) SYM_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_MINUS, 0)
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (p *OdinParser) IntegerValue() (localctx IIntegerValueContext) {
	localctx = NewIntegerValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, OdinParserRULE_integerValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		{
			p.SetState(247)
			_la = p.GetTokenStream().LA(1)

			if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserINTEGER || _la == OdinParserSCI_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerListValueContext is an interface to support dynamic dispatch.
type IIntegerListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntegerValue() []IIntegerValueContext
	IntegerValue(i int) IIntegerValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsIntegerListValueContext differentiates from other interfaces.
	IsIntegerListValueContext()
}

type IntegerListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerListValueContext() *IntegerListValueContext {
	var p = new(IntegerListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerListValue
	return p
}

func InitEmptyIntegerListValueContext(p *IntegerListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerListValue
}

func (*IntegerListValueContext) IsIntegerListValueContext() {}

func NewIntegerListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerListValueContext {
	var p = new(IntegerListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerListValue

	return p
}

func (s *IntegerListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerListValueContext) AllIntegerValue() []IIntegerValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntegerValueContext); ok {
			len++
		}
	}

	tst := make([]IIntegerValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntegerValueContext); ok {
			tst[i] = t.(IIntegerValueContext)
			i++
		}
	}

	return tst
}

func (s *IntegerListValueContext) IntegerValue(i int) IIntegerValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *IntegerListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *IntegerListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *IntegerListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *IntegerListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterIntegerListValue(s)
	}
}

func (s *IntegerListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitIntegerListValue(s)
	}
}

func (p *OdinParser) IntegerListValue() (localctx IIntegerListValueContext) {
	localctx = NewIntegerListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, OdinParserRULE_integerListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.IntegerValue()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(253)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(254)
				p.IntegerValue()
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(259)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerIntervalValueContext is an interface to support dynamic dispatch.
type IIntegerIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllIntegerValue() []IIntegerValueContext
	IntegerValue(i int) IIntegerValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode

	// IsIntegerIntervalValueContext differentiates from other interfaces.
	IsIntegerIntervalValueContext()
}

type IntegerIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerIntervalValueContext() *IntegerIntervalValueContext {
	var p = new(IntegerIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalValue
	return p
}

func InitEmptyIntegerIntervalValueContext(p *IntegerIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalValue
}

func (*IntegerIntervalValueContext) IsIntegerIntervalValueContext() {}

func NewIntegerIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerIntervalValueContext {
	var p = new(IntegerIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerIntervalValue

	return p
}

func (s *IntegerIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *IntegerIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *IntegerIntervalValueContext) AllIntegerValue() []IIntegerValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntegerValueContext); ok {
			len++
		}
	}

	tst := make([]IIntegerValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntegerValueContext); ok {
			tst[i] = t.(IIntegerValueContext)
			i++
		}
	}

	return tst
}

func (s *IntegerIntervalValueContext) IntegerValue(i int) IIntegerValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *IntegerIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *IntegerIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *IntegerIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *IntegerIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *IntegerIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *IntegerIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterIntegerIntervalValue(s)
	}
}

func (s *IntegerIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitIntegerIntervalValue(s)
	}
}

func (p *OdinParser) IntegerIntervalValue() (localctx IIntegerIntervalValueContext) {
	localctx = NewIntegerIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, OdinParserRULE_integerIntervalValue)
	var _la int

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(264)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(267)
			p.IntegerValue()
		}
		{
			p.SetState(268)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(269)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(272)
			p.IntegerValue()
		}
		{
			p.SetState(273)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(276)
				p.Relop()
			}

		}
		{
			p.SetState(279)
			p.IntegerValue()
		}
		{
			p.SetState(280)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.IntegerValue()
		}
		{
			p.SetState(284)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.IntegerValue()
		}
		{
			p.SetState(286)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerIntervalListValueContext is an interface to support dynamic dispatch.
type IIntegerIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntegerIntervalValue() []IIntegerIntervalValueContext
	IntegerIntervalValue(i int) IIntegerIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsIntegerIntervalListValueContext differentiates from other interfaces.
	IsIntegerIntervalListValueContext()
}

type IntegerIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerIntervalListValueContext() *IntegerIntervalListValueContext {
	var p = new(IntegerIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalListValue
	return p
}

func InitEmptyIntegerIntervalListValueContext(p *IntegerIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalListValue
}

func (*IntegerIntervalListValueContext) IsIntegerIntervalListValueContext() {}

func NewIntegerIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerIntervalListValueContext {
	var p = new(IntegerIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerIntervalListValue

	return p
}

func (s *IntegerIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerIntervalListValueContext) AllIntegerIntervalValue() []IIntegerIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntegerIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]IIntegerIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntegerIntervalValueContext); ok {
			tst[i] = t.(IIntegerIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *IntegerIntervalListValueContext) IntegerIntervalValue(i int) IIntegerIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerIntervalValueContext)
}

func (s *IntegerIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *IntegerIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *IntegerIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *IntegerIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterIntegerIntervalListValue(s)
	}
}

func (s *IntegerIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitIntegerIntervalListValue(s)
	}
}

func (p *OdinParser) IntegerIntervalListValue() (localctx IIntegerIntervalListValueContext) {
	localctx = NewIntegerIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, OdinParserRULE_integerIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.IntegerIntervalValue()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(291)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(292)
				p.IntegerIntervalValue()
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(297)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealValueContext is an interface to support dynamic dispatch.
type IRealValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REAL() antlr.TerminalNode
	SCI_REAL() antlr.TerminalNode
	SYM_PLUS() antlr.TerminalNode
	SYM_MINUS() antlr.TerminalNode

	// IsRealValueContext differentiates from other interfaces.
	IsRealValueContext()
}

type RealValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealValueContext() *RealValueContext {
	var p = new(RealValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realValue
	return p
}

func InitEmptyRealValueContext(p *RealValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realValue
}

func (*RealValueContext) IsRealValueContext() {}

func NewRealValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealValueContext {
	var p = new(RealValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realValue

	return p
}

func (s *RealValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealValueContext) REAL() antlr.TerminalNode {
	return s.GetToken(OdinParserREAL, 0)
}

func (s *RealValueContext) SCI_REAL() antlr.TerminalNode {
	return s.GetToken(OdinParserSCI_REAL, 0)
}

func (s *RealValueContext) SYM_PLUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS, 0)
}

func (s *RealValueContext) SYM_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_MINUS, 0)
}

func (s *RealValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRealValue(s)
	}
}

func (s *RealValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRealValue(s)
	}
}

func (p *OdinParser) RealValue() (localctx IRealValueContext) {
	localctx = NewRealValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, OdinParserRULE_realValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		{
			p.SetState(301)
			_la = p.GetTokenStream().LA(1)

			if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(304)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserREAL || _la == OdinParserSCI_REAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealListValueContext is an interface to support dynamic dispatch.
type IRealListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRealValue() []IRealValueContext
	RealValue(i int) IRealValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsRealListValueContext differentiates from other interfaces.
	IsRealListValueContext()
}

type RealListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealListValueContext() *RealListValueContext {
	var p = new(RealListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realListValue
	return p
}

func InitEmptyRealListValueContext(p *RealListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realListValue
}

func (*RealListValueContext) IsRealListValueContext() {}

func NewRealListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealListValueContext {
	var p = new(RealListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realListValue

	return p
}

func (s *RealListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealListValueContext) AllRealValue() []IRealValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRealValueContext); ok {
			len++
		}
	}

	tst := make([]IRealValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRealValueContext); ok {
			tst[i] = t.(IRealValueContext)
			i++
		}
	}

	return tst
}

func (s *RealListValueContext) RealValue(i int) IRealValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
}

func (s *RealListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *RealListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *RealListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *RealListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRealListValue(s)
	}
}

func (s *RealListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRealListValue(s)
	}
}

func (p *OdinParser) RealListValue() (localctx IRealListValueContext) {
	localctx = NewRealListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, OdinParserRULE_realListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.RealValue()
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(307)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(308)
				p.RealValue()
			}

			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(313)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealIntervalValueContext is an interface to support dynamic dispatch.
type IRealIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllRealValue() []IRealValueContext
	RealValue(i int) IRealValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode

	// IsRealIntervalValueContext differentiates from other interfaces.
	IsRealIntervalValueContext()
}

type RealIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealIntervalValueContext() *RealIntervalValueContext {
	var p = new(RealIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalValue
	return p
}

func InitEmptyRealIntervalValueContext(p *RealIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalValue
}

func (*RealIntervalValueContext) IsRealIntervalValueContext() {}

func NewRealIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealIntervalValueContext {
	var p = new(RealIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realIntervalValue

	return p
}

func (s *RealIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *RealIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *RealIntervalValueContext) AllRealValue() []IRealValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRealValueContext); ok {
			len++
		}
	}

	tst := make([]IRealValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRealValueContext); ok {
			tst[i] = t.(IRealValueContext)
			i++
		}
	}

	return tst
}

func (s *RealIntervalValueContext) RealValue(i int) IRealValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
}

func (s *RealIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *RealIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *RealIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *RealIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *RealIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *RealIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRealIntervalValue(s)
	}
}

func (s *RealIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRealIntervalValue(s)
	}
}

func (p *OdinParser) RealIntervalValue() (localctx IRealIntervalValueContext) {
	localctx = NewRealIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, OdinParserRULE_realIntervalValue)
	var _la int

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(318)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(321)
			p.RealValue()
		}
		{
			p.SetState(322)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(323)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(326)
			p.RealValue()
		}
		{
			p.SetState(327)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(330)
				p.Relop()
			}

		}
		{
			p.SetState(333)
			p.RealValue()
		}
		{
			p.SetState(334)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.RealValue()
		}
		{
			p.SetState(338)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.RealValue()
		}
		{
			p.SetState(340)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealIntervalListValueContext is an interface to support dynamic dispatch.
type IRealIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRealIntervalValue() []IRealIntervalValueContext
	RealIntervalValue(i int) IRealIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsRealIntervalListValueContext differentiates from other interfaces.
	IsRealIntervalListValueContext()
}

type RealIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealIntervalListValueContext() *RealIntervalListValueContext {
	var p = new(RealIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalListValue
	return p
}

func InitEmptyRealIntervalListValueContext(p *RealIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalListValue
}

func (*RealIntervalListValueContext) IsRealIntervalListValueContext() {}

func NewRealIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealIntervalListValueContext {
	var p = new(RealIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realIntervalListValue

	return p
}

func (s *RealIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealIntervalListValueContext) AllRealIntervalValue() []IRealIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRealIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]IRealIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRealIntervalValueContext); ok {
			tst[i] = t.(IRealIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *RealIntervalListValueContext) RealIntervalValue(i int) IRealIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealIntervalValueContext)
}

func (s *RealIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *RealIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *RealIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *RealIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRealIntervalListValue(s)
	}
}

func (s *RealIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRealIntervalListValue(s)
	}
}

func (p *OdinParser) RealIntervalListValue() (localctx IRealIntervalListValueContext) {
	localctx = NewRealIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, OdinParserRULE_realIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.RealIntervalValue()
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(345)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(346)
				p.RealIntervalValue()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(351)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_TRUE() antlr.TerminalNode
	SYM_FALSE() antlr.TerminalNode

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_booleanValue
	return p
}

func InitEmptyBooleanValueContext(p *BooleanValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_booleanValue
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueContext) SYM_TRUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_TRUE, 0)
}

func (s *BooleanValueContext) SYM_FALSE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_FALSE, 0)
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (p *OdinParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, OdinParserRULE_booleanValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserSYM_TRUE || _la == OdinParserSYM_FALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanListValueContext is an interface to support dynamic dispatch.
type IBooleanListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBooleanValue() []IBooleanValueContext
	BooleanValue(i int) IBooleanValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsBooleanListValueContext differentiates from other interfaces.
	IsBooleanListValueContext()
}

type BooleanListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanListValueContext() *BooleanListValueContext {
	var p = new(BooleanListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_booleanListValue
	return p
}

func InitEmptyBooleanListValueContext(p *BooleanListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_booleanListValue
}

func (*BooleanListValueContext) IsBooleanListValueContext() {}

func NewBooleanListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanListValueContext {
	var p = new(BooleanListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_booleanListValue

	return p
}

func (s *BooleanListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanListValueContext) AllBooleanValue() []IBooleanValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBooleanValueContext); ok {
			len++
		}
	}

	tst := make([]IBooleanValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBooleanValueContext); ok {
			tst[i] = t.(IBooleanValueContext)
			i++
		}
	}

	return tst
}

func (s *BooleanListValueContext) BooleanValue(i int) IBooleanValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *BooleanListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *BooleanListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *BooleanListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *BooleanListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterBooleanListValue(s)
	}
}

func (s *BooleanListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitBooleanListValue(s)
	}
}

func (p *OdinParser) BooleanListValue() (localctx IBooleanListValueContext) {
	localctx = NewBooleanListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, OdinParserRULE_booleanListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.BooleanValue()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(358)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(359)
				p.BooleanValue()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(364)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharacterValueContext is an interface to support dynamic dispatch.
type ICharacterValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHARACTER() antlr.TerminalNode

	// IsCharacterValueContext differentiates from other interfaces.
	IsCharacterValueContext()
}

type CharacterValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterValueContext() *CharacterValueContext {
	var p = new(CharacterValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_characterValue
	return p
}

func InitEmptyCharacterValueContext(p *CharacterValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_characterValue
}

func (*CharacterValueContext) IsCharacterValueContext() {}

func NewCharacterValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterValueContext {
	var p = new(CharacterValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_characterValue

	return p
}

func (s *CharacterValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterValueContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(OdinParserCHARACTER, 0)
}

func (s *CharacterValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterCharacterValue(s)
	}
}

func (s *CharacterValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitCharacterValue(s)
	}
}

func (p *OdinParser) CharacterValue() (localctx ICharacterValueContext) {
	localctx = NewCharacterValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, OdinParserRULE_characterValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(OdinParserCHARACTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharacterListValueContext is an interface to support dynamic dispatch.
type ICharacterListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCharacterValue() []ICharacterValueContext
	CharacterValue(i int) ICharacterValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsCharacterListValueContext differentiates from other interfaces.
	IsCharacterListValueContext()
}

type CharacterListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterListValueContext() *CharacterListValueContext {
	var p = new(CharacterListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_characterListValue
	return p
}

func InitEmptyCharacterListValueContext(p *CharacterListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_characterListValue
}

func (*CharacterListValueContext) IsCharacterListValueContext() {}

func NewCharacterListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterListValueContext {
	var p = new(CharacterListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_characterListValue

	return p
}

func (s *CharacterListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterListValueContext) AllCharacterValue() []ICharacterValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharacterValueContext); ok {
			len++
		}
	}

	tst := make([]ICharacterValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharacterValueContext); ok {
			tst[i] = t.(ICharacterValueContext)
			i++
		}
	}

	return tst
}

func (s *CharacterListValueContext) CharacterValue(i int) ICharacterValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterValueContext)
}

func (s *CharacterListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *CharacterListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *CharacterListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *CharacterListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterCharacterListValue(s)
	}
}

func (s *CharacterListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitCharacterListValue(s)
	}
}

func (p *OdinParser) CharacterListValue() (localctx ICharacterListValueContext) {
	localctx = NewCharacterListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, OdinParserRULE_characterListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.CharacterValue()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(371)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(372)
				p.CharacterValue()
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(377)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateValueContext is an interface to support dynamic dispatch.
type IDateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISO8601_DATE_AUGMENTED() antlr.TerminalNode

	// IsDateValueContext differentiates from other interfaces.
	IsDateValueContext()
}

type DateValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateValueContext() *DateValueContext {
	var p = new(DateValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateValue
	return p
}

func InitEmptyDateValueContext(p *DateValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateValue
}

func (*DateValueContext) IsDateValueContext() {}

func NewDateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateValueContext {
	var p = new(DateValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateValue

	return p
}

func (s *DateValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateValueContext) ISO8601_DATE_AUGMENTED() antlr.TerminalNode {
	return s.GetToken(OdinParserISO8601_DATE_AUGMENTED, 0)
}

func (s *DateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateValue(s)
	}
}

func (s *DateValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateValue(s)
	}
}

func (p *OdinParser) DateValue() (localctx IDateValueContext) {
	localctx = NewDateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, OdinParserRULE_dateValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(OdinParserISO8601_DATE_AUGMENTED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateListValueContext is an interface to support dynamic dispatch.
type IDateListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDateValue() []IDateValueContext
	DateValue(i int) IDateValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDateListValueContext differentiates from other interfaces.
	IsDateListValueContext()
}

type DateListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateListValueContext() *DateListValueContext {
	var p = new(DateListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateListValue
	return p
}

func InitEmptyDateListValueContext(p *DateListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateListValue
}

func (*DateListValueContext) IsDateListValueContext() {}

func NewDateListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateListValueContext {
	var p = new(DateListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateListValue

	return p
}

func (s *DateListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateListValueContext) AllDateValue() []IDateValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateValueContext); ok {
			len++
		}
	}

	tst := make([]IDateValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateValueContext); ok {
			tst[i] = t.(IDateValueContext)
			i++
		}
	}

	return tst
}

func (s *DateListValueContext) DateValue(i int) IDateValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *DateListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DateListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DateListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DateListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateListValue(s)
	}
}

func (s *DateListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateListValue(s)
	}
}

func (p *OdinParser) DateListValue() (localctx IDateListValueContext) {
	localctx = NewDateListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, OdinParserRULE_dateListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.DateValue()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(384)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(385)
				p.DateValue()
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(390)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateIntervalValueContext is an interface to support dynamic dispatch.
type IDateIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllDateValue() []IDateValueContext
	DateValue(i int) IDateValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode
	DurationValue() IDurationValueContext

	// IsDateIntervalValueContext differentiates from other interfaces.
	IsDateIntervalValueContext()
}

type DateIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateIntervalValueContext() *DateIntervalValueContext {
	var p = new(DateIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalValue
	return p
}

func InitEmptyDateIntervalValueContext(p *DateIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalValue
}

func (*DateIntervalValueContext) IsDateIntervalValueContext() {}

func NewDateIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateIntervalValueContext {
	var p = new(DateIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateIntervalValue

	return p
}

func (s *DateIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *DateIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *DateIntervalValueContext) AllDateValue() []IDateValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateValueContext); ok {
			len++
		}
	}

	tst := make([]IDateValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateValueContext); ok {
			tst[i] = t.(IDateValueContext)
			i++
		}
	}

	return tst
}

func (s *DateIntervalValueContext) DateValue(i int) IDateValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *DateIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *DateIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DateIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DateIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *DateIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *DateIntervalValueContext) DurationValue() IDurationValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *DateIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateIntervalValue(s)
	}
}

func (s *DateIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateIntervalValue(s)
	}
}

func (p *OdinParser) DateIntervalValue() (localctx IDateIntervalValueContext) {
	localctx = NewDateIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, OdinParserRULE_dateIntervalValue)
	var _la int

	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(395)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(398)
			p.DateValue()
		}
		{
			p.SetState(399)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(400)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(403)
			p.DateValue()
		}
		{
			p.SetState(404)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(407)
				p.Relop()
			}

		}
		{
			p.SetState(410)
			p.DateValue()
		}
		{
			p.SetState(411)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.DateValue()
		}
		{
			p.SetState(415)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.DurationValue()
		}
		{
			p.SetState(417)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateIntervalListValueContext is an interface to support dynamic dispatch.
type IDateIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDateIntervalValue() []IDateIntervalValueContext
	DateIntervalValue(i int) IDateIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDateIntervalListValueContext differentiates from other interfaces.
	IsDateIntervalListValueContext()
}

type DateIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateIntervalListValueContext() *DateIntervalListValueContext {
	var p = new(DateIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalListValue
	return p
}

func InitEmptyDateIntervalListValueContext(p *DateIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalListValue
}

func (*DateIntervalListValueContext) IsDateIntervalListValueContext() {}

func NewDateIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateIntervalListValueContext {
	var p = new(DateIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateIntervalListValue

	return p
}

func (s *DateIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateIntervalListValueContext) AllDateIntervalValue() []IDateIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]IDateIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateIntervalValueContext); ok {
			tst[i] = t.(IDateIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *DateIntervalListValueContext) DateIntervalValue(i int) IDateIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalValueContext)
}

func (s *DateIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DateIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DateIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DateIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateIntervalListValue(s)
	}
}

func (s *DateIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateIntervalListValue(s)
	}
}

func (p *OdinParser) DateIntervalListValue() (localctx IDateIntervalListValueContext) {
	localctx = NewDateIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, OdinParserRULE_dateIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.DateIntervalValue()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(422)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(423)
				p.DateIntervalValue()
			}

			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(428)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeValueContext is an interface to support dynamic dispatch.
type ITimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISO8601_TIME_AUGMENTED() antlr.TerminalNode

	// IsTimeValueContext differentiates from other interfaces.
	IsTimeValueContext()
}

type TimeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeValueContext() *TimeValueContext {
	var p = new(TimeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeValue
	return p
}

func InitEmptyTimeValueContext(p *TimeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeValue
}

func (*TimeValueContext) IsTimeValueContext() {}

func NewTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeValueContext {
	var p = new(TimeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeValue

	return p
}

func (s *TimeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeValueContext) ISO8601_TIME_AUGMENTED() antlr.TerminalNode {
	return s.GetToken(OdinParserISO8601_TIME_AUGMENTED, 0)
}

func (s *TimeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTimeValue(s)
	}
}

func (s *TimeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTimeValue(s)
	}
}

func (p *OdinParser) TimeValue() (localctx ITimeValueContext) {
	localctx = NewTimeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, OdinParserRULE_timeValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(OdinParserISO8601_TIME_AUGMENTED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeListValueContext is an interface to support dynamic dispatch.
type ITimeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTimeValue() []ITimeValueContext
	TimeValue(i int) ITimeValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsTimeListValueContext differentiates from other interfaces.
	IsTimeListValueContext()
}

type TimeListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeListValueContext() *TimeListValueContext {
	var p = new(TimeListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeListValue
	return p
}

func InitEmptyTimeListValueContext(p *TimeListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeListValue
}

func (*TimeListValueContext) IsTimeListValueContext() {}

func NewTimeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeListValueContext {
	var p = new(TimeListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeListValue

	return p
}

func (s *TimeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeListValueContext) AllTimeValue() []ITimeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeValueContext); ok {
			len++
		}
	}

	tst := make([]ITimeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeValueContext); ok {
			tst[i] = t.(ITimeValueContext)
			i++
		}
	}

	return tst
}

func (s *TimeListValueContext) TimeValue(i int) ITimeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *TimeListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *TimeListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *TimeListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *TimeListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTimeListValue(s)
	}
}

func (s *TimeListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTimeListValue(s)
	}
}

func (p *OdinParser) TimeListValue() (localctx ITimeListValueContext) {
	localctx = NewTimeListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, OdinParserRULE_timeListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.TimeValue()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(435)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(436)
				p.TimeValue()
			}

			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(441)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeIntervalValueContext is an interface to support dynamic dispatch.
type ITimeIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllTimeValue() []ITimeValueContext
	TimeValue(i int) ITimeValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode
	DurationValue() IDurationValueContext

	// IsTimeIntervalValueContext differentiates from other interfaces.
	IsTimeIntervalValueContext()
}

type TimeIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalValueContext() *TimeIntervalValueContext {
	var p = new(TimeIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalValue
	return p
}

func InitEmptyTimeIntervalValueContext(p *TimeIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalValue
}

func (*TimeIntervalValueContext) IsTimeIntervalValueContext() {}

func NewTimeIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalValueContext {
	var p = new(TimeIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeIntervalValue

	return p
}

func (s *TimeIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *TimeIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *TimeIntervalValueContext) AllTimeValue() []ITimeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeValueContext); ok {
			len++
		}
	}

	tst := make([]ITimeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeValueContext); ok {
			tst[i] = t.(ITimeValueContext)
			i++
		}
	}

	return tst
}

func (s *TimeIntervalValueContext) TimeValue(i int) ITimeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *TimeIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *TimeIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *TimeIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *TimeIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *TimeIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *TimeIntervalValueContext) DurationValue() IDurationValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *TimeIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTimeIntervalValue(s)
	}
}

func (s *TimeIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTimeIntervalValue(s)
	}
}

func (p *OdinParser) TimeIntervalValue() (localctx ITimeIntervalValueContext) {
	localctx = NewTimeIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, OdinParserRULE_timeIntervalValue)
	var _la int

	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(446)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(449)
			p.TimeValue()
		}
		{
			p.SetState(450)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(451)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(454)
			p.TimeValue()
		}
		{
			p.SetState(455)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(458)
				p.Relop()
			}

		}
		{
			p.SetState(461)
			p.TimeValue()
		}
		{
			p.SetState(462)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(464)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.TimeValue()
		}
		{
			p.SetState(466)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.DurationValue()
		}
		{
			p.SetState(468)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeIntervalListValueContext is an interface to support dynamic dispatch.
type ITimeIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTimeIntervalValue() []ITimeIntervalValueContext
	TimeIntervalValue(i int) ITimeIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsTimeIntervalListValueContext differentiates from other interfaces.
	IsTimeIntervalListValueContext()
}

type TimeIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalListValueContext() *TimeIntervalListValueContext {
	var p = new(TimeIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalListValue
	return p
}

func InitEmptyTimeIntervalListValueContext(p *TimeIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalListValue
}

func (*TimeIntervalListValueContext) IsTimeIntervalListValueContext() {}

func NewTimeIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalListValueContext {
	var p = new(TimeIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeIntervalListValue

	return p
}

func (s *TimeIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalListValueContext) AllTimeIntervalValue() []ITimeIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITimeIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]ITimeIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITimeIntervalValueContext); ok {
			tst[i] = t.(ITimeIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *TimeIntervalListValueContext) TimeIntervalValue(i int) ITimeIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalValueContext)
}

func (s *TimeIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *TimeIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *TimeIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *TimeIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTimeIntervalListValue(s)
	}
}

func (s *TimeIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTimeIntervalListValue(s)
	}
}

func (p *OdinParser) TimeIntervalListValue() (localctx ITimeIntervalListValueContext) {
	localctx = NewTimeIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, OdinParserRULE_timeIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.TimeIntervalValue()
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(473)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(474)
				p.TimeIntervalValue()
			}

			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(479)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateTimeValueContext is an interface to support dynamic dispatch.
type IDateTimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISO8601_DATE_TIME_AUGMENTED() antlr.TerminalNode

	// IsDateTimeValueContext differentiates from other interfaces.
	IsDateTimeValueContext()
}

type DateTimeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeValueContext() *DateTimeValueContext {
	var p = new(DateTimeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeValue
	return p
}

func InitEmptyDateTimeValueContext(p *DateTimeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeValue
}

func (*DateTimeValueContext) IsDateTimeValueContext() {}

func NewDateTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeValueContext {
	var p = new(DateTimeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeValue

	return p
}

func (s *DateTimeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeValueContext) ISO8601_DATE_TIME_AUGMENTED() antlr.TerminalNode {
	return s.GetToken(OdinParserISO8601_DATE_TIME_AUGMENTED, 0)
}

func (s *DateTimeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateTimeValue(s)
	}
}

func (s *DateTimeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateTimeValue(s)
	}
}

func (p *OdinParser) DateTimeValue() (localctx IDateTimeValueContext) {
	localctx = NewDateTimeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, OdinParserRULE_dateTimeValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(OdinParserISO8601_DATE_TIME_AUGMENTED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateTimeListValueContext is an interface to support dynamic dispatch.
type IDateTimeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDateTimeValue() []IDateTimeValueContext
	DateTimeValue(i int) IDateTimeValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDateTimeListValueContext differentiates from other interfaces.
	IsDateTimeListValueContext()
}

type DateTimeListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeListValueContext() *DateTimeListValueContext {
	var p = new(DateTimeListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeListValue
	return p
}

func InitEmptyDateTimeListValueContext(p *DateTimeListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeListValue
}

func (*DateTimeListValueContext) IsDateTimeListValueContext() {}

func NewDateTimeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeListValueContext {
	var p = new(DateTimeListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeListValue

	return p
}

func (s *DateTimeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeListValueContext) AllDateTimeValue() []IDateTimeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateTimeValueContext); ok {
			len++
		}
	}

	tst := make([]IDateTimeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateTimeValueContext); ok {
			tst[i] = t.(IDateTimeValueContext)
			i++
		}
	}

	return tst
}

func (s *DateTimeListValueContext) DateTimeValue(i int) IDateTimeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
}

func (s *DateTimeListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DateTimeListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DateTimeListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DateTimeListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimeListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateTimeListValue(s)
	}
}

func (s *DateTimeListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateTimeListValue(s)
	}
}

func (p *OdinParser) DateTimeListValue() (localctx IDateTimeListValueContext) {
	localctx = NewDateTimeListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, OdinParserRULE_dateTimeListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.DateTimeValue()
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(486)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(487)
				p.DateTimeValue()
			}

			p.SetState(490)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(492)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateTimeIntervalValueContext is an interface to support dynamic dispatch.
type IDateTimeIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllDateTimeValue() []IDateTimeValueContext
	DateTimeValue(i int) IDateTimeValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode
	DurationValue() IDurationValueContext

	// IsDateTimeIntervalValueContext differentiates from other interfaces.
	IsDateTimeIntervalValueContext()
}

type DateTimeIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeIntervalValueContext() *DateTimeIntervalValueContext {
	var p = new(DateTimeIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalValue
	return p
}

func InitEmptyDateTimeIntervalValueContext(p *DateTimeIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalValue
}

func (*DateTimeIntervalValueContext) IsDateTimeIntervalValueContext() {}

func NewDateTimeIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeIntervalValueContext {
	var p = new(DateTimeIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeIntervalValue

	return p
}

func (s *DateTimeIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *DateTimeIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *DateTimeIntervalValueContext) AllDateTimeValue() []IDateTimeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateTimeValueContext); ok {
			len++
		}
	}

	tst := make([]IDateTimeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateTimeValueContext); ok {
			tst[i] = t.(IDateTimeValueContext)
			i++
		}
	}

	return tst
}

func (s *DateTimeIntervalValueContext) DateTimeValue(i int) IDateTimeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
}

func (s *DateTimeIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *DateTimeIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DateTimeIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DateTimeIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *DateTimeIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *DateTimeIntervalValueContext) DurationValue() IDurationValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *DateTimeIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimeIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateTimeIntervalValue(s)
	}
}

func (s *DateTimeIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateTimeIntervalValue(s)
	}
}

func (p *OdinParser) DateTimeIntervalValue() (localctx IDateTimeIntervalValueContext) {
	localctx = NewDateTimeIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, OdinParserRULE_dateTimeIntervalValue)
	var _la int

	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(497)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(500)
			p.DateTimeValue()
		}
		{
			p.SetState(501)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(502)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(505)
			p.DateTimeValue()
		}
		{
			p.SetState(506)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(509)
				p.Relop()
			}

		}
		{
			p.SetState(512)
			p.DateTimeValue()
		}
		{
			p.SetState(513)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(515)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.DateTimeValue()
		}
		{
			p.SetState(517)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.DurationValue()
		}
		{
			p.SetState(519)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateTimeIntervalListValueContext is an interface to support dynamic dispatch.
type IDateTimeIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDateTimeIntervalValue() []IDateTimeIntervalValueContext
	DateTimeIntervalValue(i int) IDateTimeIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDateTimeIntervalListValueContext differentiates from other interfaces.
	IsDateTimeIntervalListValueContext()
}

type DateTimeIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeIntervalListValueContext() *DateTimeIntervalListValueContext {
	var p = new(DateTimeIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalListValue
	return p
}

func InitEmptyDateTimeIntervalListValueContext(p *DateTimeIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalListValue
}

func (*DateTimeIntervalListValueContext) IsDateTimeIntervalListValueContext() {}

func NewDateTimeIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeIntervalListValueContext {
	var p = new(DateTimeIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeIntervalListValue

	return p
}

func (s *DateTimeIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeIntervalListValueContext) AllDateTimeIntervalValue() []IDateTimeIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateTimeIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]IDateTimeIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateTimeIntervalValueContext); ok {
			tst[i] = t.(IDateTimeIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *DateTimeIntervalListValueContext) DateTimeIntervalValue(i int) IDateTimeIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeIntervalValueContext)
}

func (s *DateTimeIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DateTimeIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DateTimeIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DateTimeIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimeIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDateTimeIntervalListValue(s)
	}
}

func (s *DateTimeIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDateTimeIntervalListValue(s)
	}
}

func (p *OdinParser) DateTimeIntervalListValue() (localctx IDateTimeIntervalListValueContext) {
	localctx = NewDateTimeIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, OdinParserRULE_dateTimeIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.DateTimeIntervalValue()
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(524)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(525)
				p.DateTimeIntervalValue()
			}

			p.SetState(528)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(530)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationValueContext is an interface to support dynamic dispatch.
type IDurationValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ISO8601_DURATION() antlr.TerminalNode
	SYM_PLUS() antlr.TerminalNode
	SYM_MINUS() antlr.TerminalNode

	// IsDurationValueContext differentiates from other interfaces.
	IsDurationValueContext()
}

type DurationValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationValueContext() *DurationValueContext {
	var p = new(DurationValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationValue
	return p
}

func InitEmptyDurationValueContext(p *DurationValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationValue
}

func (*DurationValueContext) IsDurationValueContext() {}

func NewDurationValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationValueContext {
	var p = new(DurationValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationValue

	return p
}

func (s *DurationValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationValueContext) ISO8601_DURATION() antlr.TerminalNode {
	return s.GetToken(OdinParserISO8601_DURATION, 0)
}

func (s *DurationValueContext) SYM_PLUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS, 0)
}

func (s *DurationValueContext) SYM_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_MINUS, 0)
}

func (s *DurationValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDurationValue(s)
	}
}

func (s *DurationValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDurationValue(s)
	}
}

func (p *OdinParser) DurationValue() (localctx IDurationValueContext) {
	localctx = NewDurationValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, OdinParserRULE_durationValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		{
			p.SetState(534)
			_la = p.GetTokenStream().LA(1)

			if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(537)
		p.Match(OdinParserISO8601_DURATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationListValueContext is an interface to support dynamic dispatch.
type IDurationListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDurationValue() []IDurationValueContext
	DurationValue(i int) IDurationValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDurationListValueContext differentiates from other interfaces.
	IsDurationListValueContext()
}

type DurationListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationListValueContext() *DurationListValueContext {
	var p = new(DurationListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationListValue
	return p
}

func InitEmptyDurationListValueContext(p *DurationListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationListValue
}

func (*DurationListValueContext) IsDurationListValueContext() {}

func NewDurationListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationListValueContext {
	var p = new(DurationListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationListValue

	return p
}

func (s *DurationListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationListValueContext) AllDurationValue() []IDurationValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationValueContext); ok {
			len++
		}
	}

	tst := make([]IDurationValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationValueContext); ok {
			tst[i] = t.(IDurationValueContext)
			i++
		}
	}

	return tst
}

func (s *DurationListValueContext) DurationValue(i int) IDurationValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *DurationListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DurationListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DurationListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DurationListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDurationListValue(s)
	}
}

func (s *DurationListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDurationListValue(s)
	}
}

func (p *OdinParser) DurationListValue() (localctx IDurationListValueContext) {
	localctx = NewDurationListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, OdinParserRULE_durationListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.DurationValue()
	}
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 70, p.GetParserRuleContext()) {
	case 1:
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(540)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(541)
				p.DurationValue()
			}

			p.SetState(544)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(546)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationIntervalValueContext is an interface to support dynamic dispatch.
type IDurationIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSYM_IVL_DELIM() []antlr.TerminalNode
	SYM_IVL_DELIM(i int) antlr.TerminalNode
	AllDurationValue() []IDurationValueContext
	DurationValue(i int) IDurationValueContext
	SYM_IVL_SEP() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode
	Relop() IRelopContext
	SYM_PLUS_OR_MINUS() antlr.TerminalNode

	// IsDurationIntervalValueContext differentiates from other interfaces.
	IsDurationIntervalValueContext()
}

type DurationIntervalValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationIntervalValueContext() *DurationIntervalValueContext {
	var p = new(DurationIntervalValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalValue
	return p
}

func InitEmptyDurationIntervalValueContext(p *DurationIntervalValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalValue
}

func (*DurationIntervalValueContext) IsDurationIntervalValueContext() {}

func NewDurationIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationIntervalValueContext {
	var p = new(DurationIntervalValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationIntervalValue

	return p
}

func (s *DurationIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationIntervalValueContext) AllSYM_IVL_DELIM() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_IVL_DELIM)
}

func (s *DurationIntervalValueContext) SYM_IVL_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_DELIM, i)
}

func (s *DurationIntervalValueContext) AllDurationValue() []IDurationValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationValueContext); ok {
			len++
		}
	}

	tst := make([]IDurationValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationValueContext); ok {
			tst[i] = t.(IDurationValueContext)
			i++
		}
	}

	return tst
}

func (s *DurationIntervalValueContext) DurationValue(i int) IDurationValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *DurationIntervalValueContext) SYM_IVL_SEP() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_IVL_SEP, 0)
}

func (s *DurationIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DurationIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DurationIntervalValueContext) Relop() IRelopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *DurationIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *DurationIntervalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationIntervalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationIntervalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDurationIntervalValue(s)
	}
}

func (s *DurationIntervalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDurationIntervalValue(s)
	}
}

func (p *OdinParser) DurationIntervalValue() (localctx IDurationIntervalValueContext) {
	localctx = NewDurationIntervalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, OdinParserRULE_durationIntervalValue)
	var _la int

	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(551)
				p.Match(OdinParserSYM_GT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(554)
			p.DurationValue()
		}
		{
			p.SetState(555)
			p.Match(OdinParserSYM_IVL_SEP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(556)
				p.Match(OdinParserSYM_LT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(559)
			p.DurationValue()
		}
		{
			p.SetState(560)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(562)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0 {
			{
				p.SetState(563)
				p.Relop()
			}

		}
		{
			p.SetState(566)
			p.DurationValue()
		}
		{
			p.SetState(567)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(569)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.DurationValue()
		}
		{
			p.SetState(571)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(572)
			p.DurationValue()
		}
		{
			p.SetState(573)
			p.Match(OdinParserSYM_IVL_DELIM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurationIntervalListValueContext is an interface to support dynamic dispatch.
type IDurationIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDurationIntervalValue() []IDurationIntervalValueContext
	DurationIntervalValue(i int) IDurationIntervalValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsDurationIntervalListValueContext differentiates from other interfaces.
	IsDurationIntervalListValueContext()
}

type DurationIntervalListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationIntervalListValueContext() *DurationIntervalListValueContext {
	var p = new(DurationIntervalListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalListValue
	return p
}

func InitEmptyDurationIntervalListValueContext(p *DurationIntervalListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalListValue
}

func (*DurationIntervalListValueContext) IsDurationIntervalListValueContext() {}

func NewDurationIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationIntervalListValueContext {
	var p = new(DurationIntervalListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationIntervalListValue

	return p
}

func (s *DurationIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationIntervalListValueContext) AllDurationIntervalValue() []IDurationIntervalValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationIntervalValueContext); ok {
			len++
		}
	}

	tst := make([]IDurationIntervalValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationIntervalValueContext); ok {
			tst[i] = t.(IDurationIntervalValueContext)
			i++
		}
	}

	return tst
}

func (s *DurationIntervalListValueContext) DurationIntervalValue(i int) IDurationIntervalValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationIntervalValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationIntervalValueContext)
}

func (s *DurationIntervalListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *DurationIntervalListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *DurationIntervalListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *DurationIntervalListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationIntervalListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationIntervalListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterDurationIntervalListValue(s)
	}
}

func (s *DurationIntervalListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitDurationIntervalListValue(s)
	}
}

func (p *OdinParser) DurationIntervalListValue() (localctx IDurationIntervalListValueContext) {
	localctx = NewDurationIntervalListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, OdinParserRULE_durationIntervalListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(577)
		p.DurationIntervalValue()
	}
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 76, p.GetParserRuleContext()) {
	case 1:
		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(578)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(579)
				p.DurationIntervalValue()
			}

			p.SetState(582)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(584)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermCodeValueContext is an interface to support dynamic dispatch.
type ITermCodeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUALIFIED_TERM_CODE_REF() antlr.TerminalNode

	// IsTermCodeValueContext differentiates from other interfaces.
	IsTermCodeValueContext()
}

type TermCodeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermCodeValueContext() *TermCodeValueContext {
	var p = new(TermCodeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeValue
	return p
}

func InitEmptyTermCodeValueContext(p *TermCodeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeValue
}

func (*TermCodeValueContext) IsTermCodeValueContext() {}

func NewTermCodeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermCodeValueContext {
	var p = new(TermCodeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_termCodeValue

	return p
}

func (s *TermCodeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TermCodeValueContext) QUALIFIED_TERM_CODE_REF() antlr.TerminalNode {
	return s.GetToken(OdinParserQUALIFIED_TERM_CODE_REF, 0)
}

func (s *TermCodeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermCodeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermCodeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTermCodeValue(s)
	}
}

func (s *TermCodeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTermCodeValue(s)
	}
}

func (p *OdinParser) TermCodeValue() (localctx ITermCodeValueContext) {
	localctx = NewTermCodeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, OdinParserRULE_termCodeValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(588)
		p.Match(OdinParserQUALIFIED_TERM_CODE_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermCodeListValueContext is an interface to support dynamic dispatch.
type ITermCodeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTermCodeValue() []ITermCodeValueContext
	TermCodeValue(i int) ITermCodeValueContext
	AllSYM_COMMA() []antlr.TerminalNode
	SYM_COMMA(i int) antlr.TerminalNode
	SYM_LIST_CONTINUE() antlr.TerminalNode

	// IsTermCodeListValueContext differentiates from other interfaces.
	IsTermCodeListValueContext()
}

type TermCodeListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermCodeListValueContext() *TermCodeListValueContext {
	var p = new(TermCodeListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeListValue
	return p
}

func InitEmptyTermCodeListValueContext(p *TermCodeListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeListValue
}

func (*TermCodeListValueContext) IsTermCodeListValueContext() {}

func NewTermCodeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermCodeListValueContext {
	var p = new(TermCodeListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_termCodeListValue

	return p
}

func (s *TermCodeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TermCodeListValueContext) AllTermCodeValue() []ITermCodeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermCodeValueContext); ok {
			len++
		}
	}

	tst := make([]ITermCodeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermCodeValueContext); ok {
			tst[i] = t.(ITermCodeValueContext)
			i++
		}
	}

	return tst
}

func (s *TermCodeListValueContext) TermCodeValue(i int) ITermCodeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermCodeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermCodeValueContext)
}

func (s *TermCodeListValueContext) AllSYM_COMMA() []antlr.TerminalNode {
	return s.GetTokens(OdinParserSYM_COMMA)
}

func (s *TermCodeListValueContext) SYM_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_COMMA, i)
}

func (s *TermCodeListValueContext) SYM_LIST_CONTINUE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LIST_CONTINUE, 0)
}

func (s *TermCodeListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermCodeListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermCodeListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterTermCodeListValue(s)
	}
}

func (s *TermCodeListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitTermCodeListValue(s)
	}
}

func (p *OdinParser) TermCodeListValue() (localctx ITermCodeListValueContext) {
	localctx = NewTermCodeListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, OdinParserRULE_termCodeListValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.TermCodeValue()
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(591)
				p.Match(OdinParserSYM_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(592)
				p.TermCodeValue()
			}

			p.SetState(595)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(597)
			p.Match(OdinParserSYM_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)
			p.Match(OdinParserSYM_LIST_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYM_LE() antlr.TerminalNode
	SYM_GE() antlr.TerminalNode
	SYM_GT() antlr.TerminalNode
	SYM_LT() antlr.TerminalNode

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_relop
	return p
}

func InitEmptyRelopContext(p *RelopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OdinParserRULE_relop
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) SYM_LE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LE, 0)
}

func (s *RelopContext) SYM_GE() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GE, 0)
}

func (s *RelopContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *RelopContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdinParserListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *OdinParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, OdinParserRULE_relop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
