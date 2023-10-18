// Generated from OdinParser.g4 by ANTLR 4.7.

package parser // OdinParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 606,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 6, 2, 104, 10, 2, 13, 2, 14, 2, 105, 3,
	2, 5, 2, 109, 10, 2, 3, 2, 5, 2, 112, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 118, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 124, 10, 5, 3, 6, 5, 6, 127,
	10, 6, 3, 6, 3, 6, 3, 6, 6, 6, 132, 10, 6, 13, 6, 14, 6, 133, 3, 6, 6,
	6, 137, 10, 6, 13, 6, 14, 6, 138, 3, 6, 5, 6, 142, 10, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 167,
	10, 11, 13, 11, 14, 11, 168, 5, 11, 171, 10, 11, 3, 12, 5, 12, 174, 10,
	12, 3, 12, 6, 12, 177, 10, 12, 13, 12, 14, 12, 178, 3, 13, 3, 13, 3, 13,
	5, 13, 184, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 191, 10,
	14, 12, 14, 14, 14, 194, 11, 14, 3, 14, 3, 14, 5, 14, 198, 10, 14, 3, 15,
	3, 15, 3, 15, 5, 15, 203, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 215, 10, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 227, 10, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 235, 10, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 6, 20, 242, 10, 20, 13, 20, 14, 20, 243, 3, 20, 3,
	20, 5, 20, 248, 10, 20, 3, 21, 5, 21, 251, 10, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 6, 22, 258, 10, 22, 13, 22, 14, 22, 259, 3, 22, 3, 22, 5,
	22, 264, 10, 22, 3, 23, 3, 23, 5, 23, 268, 10, 23, 3, 23, 3, 23, 3, 23,
	5, 23, 273, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 280, 10,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	291, 10, 23, 3, 24, 3, 24, 3, 24, 6, 24, 296, 10, 24, 13, 24, 14, 24, 297,
	3, 24, 3, 24, 5, 24, 302, 10, 24, 3, 25, 5, 25, 305, 10, 25, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 6, 26, 312, 10, 26, 13, 26, 14, 26, 313, 3, 26,
	3, 26, 5, 26, 318, 10, 26, 3, 27, 3, 27, 5, 27, 322, 10, 27, 3, 27, 3,
	27, 3, 27, 5, 27, 327, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27,
	334, 10, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 5, 27, 345, 10, 27, 3, 28, 3, 28, 3, 28, 6, 28, 350, 10, 28, 13, 28,
	14, 28, 351, 3, 28, 3, 28, 5, 28, 356, 10, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 6, 30, 363, 10, 30, 13, 30, 14, 30, 364, 3, 30, 3, 30, 5, 30,
	369, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 6, 32, 376, 10, 32, 13,
	32, 14, 32, 377, 3, 32, 3, 32, 5, 32, 382, 10, 32, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 6, 34, 389, 10, 34, 13, 34, 14, 34, 390, 3, 34, 3, 34, 5,
	34, 395, 10, 34, 3, 35, 3, 35, 5, 35, 399, 10, 35, 3, 35, 3, 35, 3, 35,
	5, 35, 404, 10, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 411, 10,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35,
	422, 10, 35, 3, 36, 3, 36, 3, 36, 6, 36, 427, 10, 36, 13, 36, 14, 36, 428,
	3, 36, 3, 36, 5, 36, 433, 10, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 6,
	38, 440, 10, 38, 13, 38, 14, 38, 441, 3, 38, 3, 38, 5, 38, 446, 10, 38,
	3, 39, 3, 39, 5, 39, 450, 10, 39, 3, 39, 3, 39, 3, 39, 5, 39, 455, 10,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 462, 10, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 473, 10, 39, 3,
	40, 3, 40, 3, 40, 6, 40, 478, 10, 40, 13, 40, 14, 40, 479, 3, 40, 3, 40,
	5, 40, 484, 10, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 6, 42, 491, 10,
	42, 13, 42, 14, 42, 492, 3, 42, 3, 42, 5, 42, 497, 10, 42, 3, 43, 3, 43,
	5, 43, 501, 10, 43, 3, 43, 3, 43, 3, 43, 5, 43, 506, 10, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 5, 43, 513, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 524, 10, 43, 3, 44, 3, 44, 3,
	44, 6, 44, 529, 10, 44, 13, 44, 14, 44, 530, 3, 44, 3, 44, 5, 44, 535,
	10, 44, 3, 45, 5, 45, 538, 10, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 6,
	46, 545, 10, 46, 13, 46, 14, 46, 546, 3, 46, 3, 46, 5, 46, 551, 10, 46,
	3, 47, 3, 47, 5, 47, 555, 10, 47, 3, 47, 3, 47, 3, 47, 5, 47, 560, 10,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 567, 10, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 578, 10, 47, 3,
	48, 3, 48, 3, 48, 6, 48, 583, 10, 48, 13, 48, 14, 48, 584, 3, 48, 3, 48,
	5, 48, 589, 10, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 6, 50, 596, 10,
	50, 13, 50, 14, 50, 597, 3, 50, 3, 50, 5, 50, 602, 10, 50, 3, 51, 3, 51,
	3, 51, 2, 2, 52, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 2,
	8, 3, 2, 54, 55, 3, 2, 17, 18, 4, 2, 40, 40, 42, 42, 4, 2, 41, 41, 43,
	43, 3, 2, 36, 37, 3, 2, 12, 15, 2, 665, 2, 108, 3, 2, 2, 2, 4, 113, 3,
	2, 2, 2, 6, 119, 3, 2, 2, 2, 8, 123, 3, 2, 2, 2, 10, 126, 3, 2, 2, 2, 12,
	145, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2, 16, 153, 3, 2, 2, 2, 18, 157, 3,
	2, 2, 2, 20, 161, 3, 2, 2, 2, 22, 173, 3, 2, 2, 2, 24, 180, 3, 2, 2, 2,
	26, 185, 3, 2, 2, 2, 28, 202, 3, 2, 2, 2, 30, 214, 3, 2, 2, 2, 32, 226,
	3, 2, 2, 2, 34, 234, 3, 2, 2, 2, 36, 236, 3, 2, 2, 2, 38, 238, 3, 2, 2,
	2, 40, 250, 3, 2, 2, 2, 42, 254, 3, 2, 2, 2, 44, 290, 3, 2, 2, 2, 46, 292,
	3, 2, 2, 2, 48, 304, 3, 2, 2, 2, 50, 308, 3, 2, 2, 2, 52, 344, 3, 2, 2,
	2, 54, 346, 3, 2, 2, 2, 56, 357, 3, 2, 2, 2, 58, 359, 3, 2, 2, 2, 60, 370,
	3, 2, 2, 2, 62, 372, 3, 2, 2, 2, 64, 383, 3, 2, 2, 2, 66, 385, 3, 2, 2,
	2, 68, 421, 3, 2, 2, 2, 70, 423, 3, 2, 2, 2, 72, 434, 3, 2, 2, 2, 74, 436,
	3, 2, 2, 2, 76, 472, 3, 2, 2, 2, 78, 474, 3, 2, 2, 2, 80, 485, 3, 2, 2,
	2, 82, 487, 3, 2, 2, 2, 84, 523, 3, 2, 2, 2, 86, 525, 3, 2, 2, 2, 88, 537,
	3, 2, 2, 2, 90, 541, 3, 2, 2, 2, 92, 577, 3, 2, 2, 2, 94, 579, 3, 2, 2,
	2, 96, 590, 3, 2, 2, 2, 98, 592, 3, 2, 2, 2, 100, 603, 3, 2, 2, 2, 102,
	104, 5, 4, 3, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 109, 5, 10,
	6, 2, 108, 103, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2,
	110, 112, 7, 2, 2, 3, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	3, 3, 2, 2, 2, 113, 114, 5, 6, 4, 2, 114, 115, 7, 8, 2, 2, 115, 117, 5,
	8, 5, 2, 116, 118, 7, 47, 2, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2,
	2, 118, 5, 3, 2, 2, 2, 119, 120, 9, 2, 2, 2, 120, 7, 3, 2, 2, 2, 121, 124,
	5, 10, 6, 2, 122, 124, 5, 18, 10, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3,
	2, 2, 2, 124, 9, 3, 2, 2, 2, 125, 127, 5, 12, 7, 2, 126, 125, 3, 2, 2,
	2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 141, 7, 15, 2, 2, 129,
	142, 5, 28, 15, 2, 130, 132, 5, 4, 3, 2, 131, 130, 3, 2, 2, 2, 132, 133,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 142, 3, 2,
	2, 2, 135, 137, 5, 14, 8, 2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2,
	138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140,
	142, 7, 6, 2, 2, 141, 129, 3, 2, 2, 2, 141, 131, 3, 2, 2, 2, 141, 136,
	3, 2, 2, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 3, 2,
	2, 2, 143, 144, 7, 14, 2, 2, 144, 11, 3, 2, 2, 2, 145, 146, 7, 48, 2, 2,
	146, 147, 5, 26, 14, 2, 147, 148, 7, 49, 2, 2, 148, 13, 3, 2, 2, 2, 149,
	150, 5, 16, 9, 2, 150, 151, 7, 8, 2, 2, 151, 152, 5, 8, 5, 2, 152, 15,
	3, 2, 2, 2, 153, 154, 7, 50, 2, 2, 154, 155, 5, 30, 16, 2, 155, 156, 7,
	51, 2, 2, 156, 17, 3, 2, 2, 2, 157, 158, 7, 15, 2, 2, 158, 159, 5, 20,
	11, 2, 159, 160, 7, 14, 2, 2, 160, 19, 3, 2, 2, 2, 161, 170, 5, 22, 12,
	2, 162, 163, 7, 11, 2, 2, 163, 171, 7, 7, 2, 2, 164, 165, 7, 11, 2, 2,
	165, 167, 5, 22, 12, 2, 166, 164, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 162,
	3, 2, 2, 2, 170, 166, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 21, 3, 2,
	2, 2, 172, 174, 5, 16, 9, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 176, 3, 2, 2, 2, 175, 177, 5, 24, 13, 2, 176, 175, 3, 2, 2, 2, 177,
	178, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 23, 3,
	2, 2, 2, 180, 181, 7, 9, 2, 2, 181, 183, 7, 55, 2, 2, 182, 184, 5, 16,
	9, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 25, 3, 2, 2, 2,
	185, 197, 7, 54, 2, 2, 186, 187, 7, 15, 2, 2, 187, 192, 5, 26, 14, 2, 188,
	189, 7, 11, 2, 2, 189, 191, 5, 26, 14, 2, 190, 188, 3, 2, 2, 2, 191, 194,
	3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 195, 3, 2,
	2, 2, 194, 192, 3, 2, 2, 2, 195, 196, 7, 14, 2, 2, 196, 198, 3, 2, 2, 2,
	197, 186, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 27, 3, 2, 2, 2, 199, 203,
	5, 30, 16, 2, 200, 203, 5, 32, 17, 2, 201, 203, 5, 34, 18, 2, 202, 199,
	3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 201, 3, 2, 2, 2, 203, 29, 3, 2,
	2, 2, 204, 215, 5, 36, 19, 2, 205, 215, 5, 40, 21, 2, 206, 215, 5, 48,
	25, 2, 207, 215, 5, 56, 29, 2, 208, 215, 5, 60, 31, 2, 209, 215, 5, 96,
	49, 2, 210, 215, 5, 64, 33, 2, 211, 215, 5, 72, 37, 2, 212, 215, 5, 80,
	41, 2, 213, 215, 5, 88, 45, 2, 214, 204, 3, 2, 2, 2, 214, 205, 3, 2, 2,
	2, 214, 206, 3, 2, 2, 2, 214, 207, 3, 2, 2, 2, 214, 208, 3, 2, 2, 2, 214,
	209, 3, 2, 2, 2, 214, 210, 3, 2, 2, 2, 214, 211, 3, 2, 2, 2, 214, 212,
	3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215, 31, 3, 2, 2, 2, 216, 227, 5, 38,
	20, 2, 217, 227, 5, 42, 22, 2, 218, 227, 5, 50, 26, 2, 219, 227, 5, 58,
	30, 2, 220, 227, 5, 62, 32, 2, 221, 227, 5, 98, 50, 2, 222, 227, 5, 66,
	34, 2, 223, 227, 5, 74, 38, 2, 224, 227, 5, 82, 42, 2, 225, 227, 5, 90,
	46, 2, 226, 216, 3, 2, 2, 2, 226, 217, 3, 2, 2, 2, 226, 218, 3, 2, 2, 2,
	226, 219, 3, 2, 2, 2, 226, 220, 3, 2, 2, 2, 226, 221, 3, 2, 2, 2, 226,
	222, 3, 2, 2, 2, 226, 223, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 225,
	3, 2, 2, 2, 227, 33, 3, 2, 2, 2, 228, 235, 5, 44, 23, 2, 229, 235, 5, 52,
	27, 2, 230, 235, 5, 68, 35, 2, 231, 235, 5, 76, 39, 2, 232, 235, 5, 84,
	43, 2, 233, 235, 5, 92, 47, 2, 234, 228, 3, 2, 2, 2, 234, 229, 3, 2, 2,
	2, 234, 230, 3, 2, 2, 2, 234, 231, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234,
	233, 3, 2, 2, 2, 235, 35, 3, 2, 2, 2, 236, 237, 7, 44, 2, 2, 237, 37, 3,
	2, 2, 2, 238, 247, 5, 36, 19, 2, 239, 240, 7, 11, 2, 2, 240, 242, 5, 36,
	19, 2, 241, 239, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2,
	243, 244, 3, 2, 2, 2, 244, 248, 3, 2, 2, 2, 245, 246, 7, 11, 2, 2, 246,
	248, 7, 7, 2, 2, 247, 241, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 39, 3,
	2, 2, 2, 249, 251, 9, 3, 2, 2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 2, 2,
	2, 251, 252, 3, 2, 2, 2, 252, 253, 9, 4, 2, 2, 253, 41, 3, 2, 2, 2, 254,
	263, 5, 40, 21, 2, 255, 256, 7, 11, 2, 2, 256, 258, 5, 40, 21, 2, 257,
	255, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 264, 3, 2, 2, 2, 261, 262, 7, 11, 2, 2, 262, 264, 7, 7,
	2, 2, 263, 257, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264, 43, 3, 2, 2, 2,
	265, 267, 7, 21, 2, 2, 266, 268, 7, 14, 2, 2, 267, 266, 3, 2, 2, 2, 267,
	268, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 5, 40, 21, 2, 270, 272,
	7, 22, 2, 2, 271, 273, 7, 15, 2, 2, 272, 271, 3, 2, 2, 2, 272, 273, 3,
	2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 5, 40, 21, 2, 275, 276, 7, 21,
	2, 2, 276, 291, 3, 2, 2, 2, 277, 279, 7, 21, 2, 2, 278, 280, 5, 100, 51,
	2, 279, 278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281,
	282, 5, 40, 21, 2, 282, 283, 7, 21, 2, 2, 283, 291, 3, 2, 2, 2, 284, 285,
	7, 21, 2, 2, 285, 286, 5, 40, 21, 2, 286, 287, 7, 16, 2, 2, 287, 288, 5,
	40, 21, 2, 288, 289, 7, 21, 2, 2, 289, 291, 3, 2, 2, 2, 290, 265, 3, 2,
	2, 2, 290, 277, 3, 2, 2, 2, 290, 284, 3, 2, 2, 2, 291, 45, 3, 2, 2, 2,
	292, 301, 5, 44, 23, 2, 293, 294, 7, 11, 2, 2, 294, 296, 5, 44, 23, 2,
	295, 293, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297,
	298, 3, 2, 2, 2, 298, 302, 3, 2, 2, 2, 299, 300, 7, 11, 2, 2, 300, 302,
	7, 7, 2, 2, 301, 295, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 302, 47, 3, 2,
	2, 2, 303, 305, 9, 3, 2, 2, 304, 303, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2,
	305, 306, 3, 2, 2, 2, 306, 307, 9, 5, 2, 2, 307, 49, 3, 2, 2, 2, 308, 317,
	5, 48, 25, 2, 309, 310, 7, 11, 2, 2, 310, 312, 5, 48, 25, 2, 311, 309,
	3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2,
	2, 2, 314, 318, 3, 2, 2, 2, 315, 316, 7, 11, 2, 2, 316, 318, 7, 7, 2, 2,
	317, 311, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 318, 51, 3, 2, 2, 2, 319, 321,
	7, 21, 2, 2, 320, 322, 7, 14, 2, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3,
	2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324, 5, 48, 25, 2, 324, 326, 7, 22,
	2, 2, 325, 327, 7, 15, 2, 2, 326, 325, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2,
	327, 328, 3, 2, 2, 2, 328, 329, 5, 48, 25, 2, 329, 330, 7, 21, 2, 2, 330,
	345, 3, 2, 2, 2, 331, 333, 7, 21, 2, 2, 332, 334, 5, 100, 51, 2, 333, 332,
	3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 5, 48,
	25, 2, 336, 337, 7, 21, 2, 2, 337, 345, 3, 2, 2, 2, 338, 339, 7, 21, 2,
	2, 339, 340, 5, 48, 25, 2, 340, 341, 7, 16, 2, 2, 341, 342, 5, 48, 25,
	2, 342, 343, 7, 21, 2, 2, 343, 345, 3, 2, 2, 2, 344, 319, 3, 2, 2, 2, 344,
	331, 3, 2, 2, 2, 344, 338, 3, 2, 2, 2, 345, 53, 3, 2, 2, 2, 346, 355, 5,
	52, 27, 2, 347, 348, 7, 11, 2, 2, 348, 350, 5, 52, 27, 2, 349, 347, 3,
	2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2,
	2, 352, 356, 3, 2, 2, 2, 353, 354, 7, 11, 2, 2, 354, 356, 7, 7, 2, 2, 355,
	349, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 356, 55, 3, 2, 2, 2, 357, 358, 9,
	6, 2, 2, 358, 57, 3, 2, 2, 2, 359, 368, 5, 56, 29, 2, 360, 361, 7, 11,
	2, 2, 361, 363, 5, 56, 29, 2, 362, 360, 3, 2, 2, 2, 363, 364, 3, 2, 2,
	2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 369, 3, 2, 2, 2, 366,
	367, 7, 11, 2, 2, 367, 369, 7, 7, 2, 2, 368, 362, 3, 2, 2, 2, 368, 366,
	3, 2, 2, 2, 369, 59, 3, 2, 2, 2, 370, 371, 7, 45, 2, 2, 371, 61, 3, 2,
	2, 2, 372, 381, 5, 60, 31, 2, 373, 374, 7, 11, 2, 2, 374, 376, 5, 60, 31,
	2, 375, 373, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377,
	378, 3, 2, 2, 2, 378, 382, 3, 2, 2, 2, 379, 380, 7, 11, 2, 2, 380, 382,
	7, 7, 2, 2, 381, 375, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 382, 63, 3, 2,
	2, 2, 383, 384, 7, 32, 2, 2, 384, 65, 3, 2, 2, 2, 385, 394, 5, 64, 33,
	2, 386, 387, 7, 11, 2, 2, 387, 389, 5, 64, 33, 2, 388, 386, 3, 2, 2, 2,
	389, 390, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391,
	395, 3, 2, 2, 2, 392, 393, 7, 11, 2, 2, 393, 395, 7, 7, 2, 2, 394, 388,
	3, 2, 2, 2, 394, 392, 3, 2, 2, 2, 395, 67, 3, 2, 2, 2, 396, 398, 7, 21,
	2, 2, 397, 399, 7, 14, 2, 2, 398, 397, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2,
	399, 400, 3, 2, 2, 2, 400, 401, 5, 64, 33, 2, 401, 403, 7, 22, 2, 2, 402,
	404, 7, 15, 2, 2, 403, 402, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405,
	3, 2, 2, 2, 405, 406, 5, 64, 33, 2, 406, 407, 7, 21, 2, 2, 407, 422, 3,
	2, 2, 2, 408, 410, 7, 21, 2, 2, 409, 411, 5, 100, 51, 2, 410, 409, 3, 2,
	2, 2, 410, 411, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 5, 64, 33,
	2, 413, 414, 7, 21, 2, 2, 414, 422, 3, 2, 2, 2, 415, 416, 7, 21, 2, 2,
	416, 417, 5, 64, 33, 2, 417, 418, 7, 16, 2, 2, 418, 419, 5, 88, 45, 2,
	419, 420, 7, 21, 2, 2, 420, 422, 3, 2, 2, 2, 421, 396, 3, 2, 2, 2, 421,
	408, 3, 2, 2, 2, 421, 415, 3, 2, 2, 2, 422, 69, 3, 2, 2, 2, 423, 432, 5,
	68, 35, 2, 424, 425, 7, 11, 2, 2, 425, 427, 5, 68, 35, 2, 426, 424, 3,
	2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2,
	2, 429, 433, 3, 2, 2, 2, 430, 431, 7, 11, 2, 2, 431, 433, 7, 7, 2, 2, 432,
	426, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 433, 71, 3, 2, 2, 2, 434, 435, 7,
	33, 2, 2, 435, 73, 3, 2, 2, 2, 436, 445, 5, 72, 37, 2, 437, 438, 7, 11,
	2, 2, 438, 440, 5, 72, 37, 2, 439, 437, 3, 2, 2, 2, 440, 441, 3, 2, 2,
	2, 441, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 446, 3, 2, 2, 2, 443,
	444, 7, 11, 2, 2, 444, 446, 7, 7, 2, 2, 445, 439, 3, 2, 2, 2, 445, 443,
	3, 2, 2, 2, 446, 75, 3, 2, 2, 2, 447, 449, 7, 21, 2, 2, 448, 450, 7, 14,
	2, 2, 449, 448, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 451, 3, 2, 2, 2,
	451, 452, 5, 72, 37, 2, 452, 454, 7, 22, 2, 2, 453, 455, 7, 15, 2, 2, 454,
	453, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 457,
	5, 72, 37, 2, 457, 458, 7, 21, 2, 2, 458, 473, 3, 2, 2, 2, 459, 461, 7,
	21, 2, 2, 460, 462, 5, 100, 51, 2, 461, 460, 3, 2, 2, 2, 461, 462, 3, 2,
	2, 2, 462, 463, 3, 2, 2, 2, 463, 464, 5, 72, 37, 2, 464, 465, 7, 21, 2,
	2, 465, 473, 3, 2, 2, 2, 466, 467, 7, 21, 2, 2, 467, 468, 5, 72, 37, 2,
	468, 469, 7, 16, 2, 2, 469, 470, 5, 88, 45, 2, 470, 471, 7, 21, 2, 2, 471,
	473, 3, 2, 2, 2, 472, 447, 3, 2, 2, 2, 472, 459, 3, 2, 2, 2, 472, 466,
	3, 2, 2, 2, 473, 77, 3, 2, 2, 2, 474, 483, 5, 76, 39, 2, 475, 476, 7, 11,
	2, 2, 476, 478, 5, 76, 39, 2, 477, 475, 3, 2, 2, 2, 478, 479, 3, 2, 2,
	2, 479, 477, 3, 2, 2, 2, 479, 480, 3, 2, 2, 2, 480, 484, 3, 2, 2, 2, 481,
	482, 7, 11, 2, 2, 482, 484, 7, 7, 2, 2, 483, 477, 3, 2, 2, 2, 483, 481,
	3, 2, 2, 2, 484, 79, 3, 2, 2, 2, 485, 486, 7, 34, 2, 2, 486, 81, 3, 2,
	2, 2, 487, 496, 5, 80, 41, 2, 488, 489, 7, 11, 2, 2, 489, 491, 5, 80, 41,
	2, 490, 488, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 490, 3, 2, 2, 2, 492,
	493, 3, 2, 2, 2, 493, 497, 3, 2, 2, 2, 494, 495, 7, 11, 2, 2, 495, 497,
	7, 7, 2, 2, 496, 490, 3, 2, 2, 2, 496, 494, 3, 2, 2, 2, 497, 83, 3, 2,
	2, 2, 498, 500, 7, 21, 2, 2, 499, 501, 7, 14, 2, 2, 500, 499, 3, 2, 2,
	2, 500, 501, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 503, 5, 80, 41, 2,
	503, 505, 7, 22, 2, 2, 504, 506, 7, 15, 2, 2, 505, 504, 3, 2, 2, 2, 505,
	506, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 508, 5, 80, 41, 2, 508, 509,
	7, 21, 2, 2, 509, 524, 3, 2, 2, 2, 510, 512, 7, 21, 2, 2, 511, 513, 5,
	100, 51, 2, 512, 511, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 514, 3, 2,
	2, 2, 514, 515, 5, 80, 41, 2, 515, 516, 7, 21, 2, 2, 516, 524, 3, 2, 2,
	2, 517, 518, 7, 21, 2, 2, 518, 519, 5, 80, 41, 2, 519, 520, 7, 16, 2, 2,
	520, 521, 5, 88, 45, 2, 521, 522, 7, 21, 2, 2, 522, 524, 3, 2, 2, 2, 523,
	498, 3, 2, 2, 2, 523, 510, 3, 2, 2, 2, 523, 517, 3, 2, 2, 2, 524, 85, 3,
	2, 2, 2, 525, 534, 5, 84, 43, 2, 526, 527, 7, 11, 2, 2, 527, 529, 5, 84,
	43, 2, 528, 526, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 528, 3, 2, 2, 2,
	530, 531, 3, 2, 2, 2, 531, 535, 3, 2, 2, 2, 532, 533, 7, 11, 2, 2, 533,
	535, 7, 7, 2, 2, 534, 528, 3, 2, 2, 2, 534, 532, 3, 2, 2, 2, 535, 87, 3,
	2, 2, 2, 536, 538, 9, 3, 2, 2, 537, 536, 3, 2, 2, 2, 537, 538, 3, 2, 2,
	2, 538, 539, 3, 2, 2, 2, 539, 540, 7, 35, 2, 2, 540, 89, 3, 2, 2, 2, 541,
	550, 5, 88, 45, 2, 542, 543, 7, 11, 2, 2, 543, 545, 5, 88, 45, 2, 544,
	542, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 544, 3, 2, 2, 2, 546, 547,
	3, 2, 2, 2, 547, 551, 3, 2, 2, 2, 548, 549, 7, 11, 2, 2, 549, 551, 7, 7,
	2, 2, 550, 544, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 551, 91, 3, 2, 2, 2,
	552, 554, 7, 21, 2, 2, 553, 555, 7, 14, 2, 2, 554, 553, 3, 2, 2, 2, 554,
	555, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556, 557, 5, 88, 45, 2, 557, 559,
	7, 22, 2, 2, 558, 560, 7, 15, 2, 2, 559, 558, 3, 2, 2, 2, 559, 560, 3,
	2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 562, 5, 88, 45, 2, 562, 563, 7, 21,
	2, 2, 563, 578, 3, 2, 2, 2, 564, 566, 7, 21, 2, 2, 565, 567, 5, 100, 51,
	2, 566, 565, 3, 2, 2, 2, 566, 567, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568,
	569, 5, 88, 45, 2, 569, 570, 7, 21, 2, 2, 570, 578, 3, 2, 2, 2, 571, 572,
	7, 21, 2, 2, 572, 573, 5, 88, 45, 2, 573, 574, 7, 16, 2, 2, 574, 575, 5,
	88, 45, 2, 575, 576, 7, 21, 2, 2, 576, 578, 3, 2, 2, 2, 577, 552, 3, 2,
	2, 2, 577, 564, 3, 2, 2, 2, 577, 571, 3, 2, 2, 2, 578, 93, 3, 2, 2, 2,
	579, 588, 5, 92, 47, 2, 580, 581, 7, 11, 2, 2, 581, 583, 5, 92, 47, 2,
	582, 580, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 584,
	585, 3, 2, 2, 2, 585, 589, 3, 2, 2, 2, 586, 587, 7, 11, 2, 2, 587, 589,
	7, 7, 2, 2, 588, 582, 3, 2, 2, 2, 588, 586, 3, 2, 2, 2, 589, 95, 3, 2,
	2, 2, 590, 591, 7, 56, 2, 2, 591, 97, 3, 2, 2, 2, 592, 601, 5, 96, 49,
	2, 593, 594, 7, 11, 2, 2, 594, 596, 5, 96, 49, 2, 595, 593, 3, 2, 2, 2,
	596, 597, 3, 2, 2, 2, 597, 595, 3, 2, 2, 2, 597, 598, 3, 2, 2, 2, 598,
	602, 3, 2, 2, 2, 599, 600, 7, 11, 2, 2, 600, 602, 7, 7, 2, 2, 601, 595,
	3, 2, 2, 2, 601, 599, 3, 2, 2, 2, 602, 99, 3, 2, 2, 2, 603, 604, 9, 7,
	2, 2, 604, 101, 3, 2, 2, 2, 81, 105, 108, 111, 117, 123, 126, 133, 138,
	141, 168, 170, 173, 178, 183, 192, 197, 202, 214, 226, 234, 243, 247, 250,
	259, 263, 267, 272, 279, 290, 297, 301, 304, 313, 317, 321, 326, 333, 344,
	351, 355, 364, 368, 377, 381, 390, 394, 398, 403, 410, 421, 428, 432, 441,
	445, 449, 454, 461, 472, 479, 483, 492, 496, 500, 505, 512, 523, 530, 534,
	537, 546, 550, 554, 559, 566, 577, 584, 588, 597, 601,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'...'", "'='", "'/'", "'::'", "','", "", "", "'>'",
	"'<'", "", "'+'", "'-'", "'%'", "'^'", "'|'", "'..'", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"'.'", "';'", "'('", "')'", "'['", "']'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "CMT_LINE", "EOL", "WS", "ODIN_URI", "SYM_LIST_CONTINUE", "SYM_EQ",
	"SYM_SLASH", "SYM_NAMESPACE_SEP", "SYM_COMMA", "SYM_LE", "SYM_GE", "SYM_GT",
	"SYM_LT", "SYM_PLUS_OR_MINUS", "SYM_PLUS", "SYM_MINUS", "SYM_PERCENT",
	"SYM_CARET", "SYM_IVL_DELIM", "SYM_IVL_SEP", "OBJECT_VERSION_ID", "ARCHETYPE_HRID",
	"ARCHETYPE_REF", "VERSION_ID", "TERM_CODE_REF", "ROOT_ID_CODE", "ID_CODE",
	"AT_CODE", "AC_CODE", "ISO8601_DATE_AUGMENTED", "ISO8601_TIME_AUGMENTED",
	"ISO8601_DATE_TIME_AUGMENTED", "ISO8601_DURATION", "SYM_TRUE", "SYM_FALSE",
	"GUID", "UUID", "INTEGER", "REAL", "SCI_INTEGER", "SCI_REAL", "STRING",
	"CHARACTER", "SYM_DOT", "SYM_SEMI_COLON", "SYM_LPAREN", "SYM_RPAREN", "SYM_LBRACKET",
	"SYM_RBRACKET", "SYM_LCURLY", "SYM_RCURLY", "UC_ID", "LC_ID", "QUALIFIED_TERM_CODE_REF",
}

var ruleNames = []string{
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
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type OdinParser struct {
	*antlr.BaseParser
}

func NewOdinParser(input antlr.TokenStream) *OdinParser {
	this := new(OdinParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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

	// IsOdinObjectContext differentiates from other interfaces.
	IsOdinObjectContext()
}

type OdinObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectContext() *OdinObjectContext {
	var p = new(OdinObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinObject
	return p
}

func (*OdinObjectContext) IsOdinObjectContext() {}

func NewOdinObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectContext {
	var p = new(OdinObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObject

	return p
}

func (s *OdinObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectContext) OdinObjectValueBlock() IOdinObjectValueBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinObjectValueBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinObjectValueBlockContext)
}

func (s *OdinObjectContext) EOF() antlr.TerminalNode {
	return s.GetToken(OdinParserEOF, 0)
}

func (s *OdinObjectContext) AllOdinAttrVal() []IOdinAttrValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOdinAttrValContext)(nil)).Elem())
	var tst = make([]IOdinAttrValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOdinAttrValContext)
		}
	}

	return tst
}

func (s *OdinObjectContext) OdinAttrVal(i int) IOdinAttrValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinAttrValContext)(nil)).Elem(), i)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OdinParserUC_ID, OdinParserLC_ID:
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserUC_ID || _la == OdinParserLC_ID {
			{
				p.SetState(100)
				p.OdinAttrVal()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserSYM_LT, OdinParserSYM_LPAREN:
		{
			p.SetState(105)
			p.OdinObjectValueBlock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(108)
			p.Match(OdinParserEOF)
		}

	}

	return localctx
}

// IOdinAttrValContext is an interface to support dynamic dispatch.
type IOdinAttrValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinAttrValContext differentiates from other interfaces.
	IsOdinAttrValContext()
}

type OdinAttrValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinAttrValContext() *OdinAttrValContext {
	var p = new(OdinAttrValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrVal
	return p
}

func (*OdinAttrValContext) IsOdinAttrValContext() {}

func NewOdinAttrValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinAttrValContext {
	var p = new(OdinAttrValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinAttrVal

	return p
}

func (s *OdinAttrValContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinAttrValContext) OdinAttrName() IOdinAttrNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinAttrNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinAttrNameContext)
}

func (s *OdinAttrValContext) OdinObjectBlock() IOdinObjectBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinObjectBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinObjectBlockContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.OdinAttrName()
	}
	{
		p.SetState(112)
		p.Match(OdinParserSYM_EQ)
	}
	{
		p.SetState(113)
		p.OdinObjectBlock()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_SEMI_COLON {
		{
			p.SetState(114)
			p.Match(OdinParserSYM_SEMI_COLON)
		}

	}

	return localctx
}

// IOdinAttrNameContext is an interface to support dynamic dispatch.
type IOdinAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinAttrNameContext differentiates from other interfaces.
	IsOdinAttrNameContext()
}

type OdinAttrNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinAttrNameContext() *OdinAttrNameContext {
	var p = new(OdinAttrNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinAttrName
	return p
}

func (*OdinAttrNameContext) IsOdinAttrNameContext() {}

func NewOdinAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinAttrNameContext {
	var p = new(OdinAttrNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	_la = p.GetTokenStream().LA(1)

	if !(_la == OdinParserUC_ID || _la == OdinParserLC_ID) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IOdinObjectBlockContext is an interface to support dynamic dispatch.
type IOdinObjectBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinObjectBlockContext differentiates from other interfaces.
	IsOdinObjectBlockContext()
}

type OdinObjectBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectBlockContext() *OdinObjectBlockContext {
	var p = new(OdinObjectBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectBlock
	return p
}

func (*OdinObjectBlockContext) IsOdinObjectBlockContext() {}

func NewOdinObjectBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectBlockContext {
	var p = new(OdinObjectBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectBlock

	return p
}

func (s *OdinObjectBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectBlockContext) OdinObjectValueBlock() IOdinObjectValueBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinObjectValueBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinObjectValueBlockContext)
}

func (s *OdinObjectBlockContext) OdinObjectReferenceBlock() IOdinObjectReferenceBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinObjectReferenceBlockContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
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

	}

	return localctx
}

// IOdinObjectValueBlockContext is an interface to support dynamic dispatch.
type IOdinObjectValueBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinObjectValueBlockContext differentiates from other interfaces.
	IsOdinObjectValueBlockContext()
}

type OdinObjectValueBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectValueBlockContext() *OdinObjectValueBlockContext {
	var p = new(OdinObjectValueBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectValueBlock
	return p
}

func (*OdinObjectValueBlockContext) IsOdinObjectValueBlockContext() {}

func NewOdinObjectValueBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectValueBlockContext {
	var p = new(OdinObjectValueBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectValueBlock

	return p
}

func (s *OdinObjectValueBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectValueBlockContext) RmTypeSpec() IRmTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRmTypeSpecContext)
}

func (s *OdinObjectValueBlockContext) PrimitiveObject() IPrimitiveObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveObjectContext)
}

func (s *OdinObjectValueBlockContext) ODIN_URI() antlr.TerminalNode {
	return s.GetToken(OdinParserODIN_URI, 0)
}

func (s *OdinObjectValueBlockContext) AllOdinAttrVal() []IOdinAttrValContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOdinAttrValContext)(nil)).Elem())
	var tst = make([]IOdinAttrValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOdinAttrValContext)
		}
	}

	return tst
}

func (s *OdinObjectValueBlockContext) OdinAttrVal(i int) IOdinAttrValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinAttrValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOdinAttrValContext)
}

func (s *OdinObjectValueBlockContext) AllOdinKeyedObject() []IOdinKeyedObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOdinKeyedObjectContext)(nil)).Elem())
	var tst = make([]IOdinKeyedObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOdinKeyedObjectContext)
		}
	}

	return tst
}

func (s *OdinObjectValueBlockContext) OdinKeyedObject(i int) IOdinKeyedObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinKeyedObjectContext)(nil)).Elem(), i)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
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
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OdinParserSYM_PLUS, OdinParserSYM_MINUS, OdinParserSYM_IVL_DELIM, OdinParserISO8601_DATE_AUGMENTED, OdinParserISO8601_TIME_AUGMENTED, OdinParserISO8601_DATE_TIME_AUGMENTED, OdinParserISO8601_DURATION, OdinParserSYM_TRUE, OdinParserSYM_FALSE, OdinParserINTEGER, OdinParserREAL, OdinParserSCI_INTEGER, OdinParserSCI_REAL, OdinParserSTRING, OdinParserCHARACTER, OdinParserQUALIFIED_TERM_CODE_REF:
		{
			p.SetState(127)
			p.PrimitiveObject()
		}

	case OdinParserUC_ID, OdinParserLC_ID:
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserUC_ID || _la == OdinParserLC_ID {
			{
				p.SetState(128)
				p.OdinAttrVal()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserSYM_LBRACKET:
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_LBRACKET {
			{
				p.SetState(133)
				p.OdinKeyedObject()
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case OdinParserODIN_URI:
		{
			p.SetState(138)
			p.Match(OdinParserODIN_URI)
		}

	case OdinParserSYM_GT:

	default:
	}
	{
		p.SetState(141)
		p.Match(OdinParserSYM_GT)
	}

	return localctx
}

// IRmTypeSpecContext is an interface to support dynamic dispatch.
type IRmTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRmTypeSpecContext differentiates from other interfaces.
	IsRmTypeSpecContext()
}

type RmTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmTypeSpecContext() *RmTypeSpecContext {
	var p = new(RmTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeSpec
	return p
}

func (*RmTypeSpecContext) IsRmTypeSpecContext() {}

func NewRmTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmTypeSpecContext {
	var p = new(RmTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_rmTypeSpec

	return p
}

func (s *RmTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *RmTypeSpecContext) RmTypeId() IRmTypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmTypeIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRmTypeIdContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(OdinParserSYM_LPAREN)
	}
	{
		p.SetState(144)
		p.RmTypeId()
	}
	{
		p.SetState(145)
		p.Match(OdinParserSYM_RPAREN)
	}

	return localctx
}

// IOdinKeyedObjectContext is an interface to support dynamic dispatch.
type IOdinKeyedObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinKeyedObjectContext differentiates from other interfaces.
	IsOdinKeyedObjectContext()
}

type OdinKeyedObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinKeyedObjectContext() *OdinKeyedObjectContext {
	var p = new(OdinKeyedObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeyedObject
	return p
}

func (*OdinKeyedObjectContext) IsOdinKeyedObjectContext() {}

func NewOdinKeyedObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinKeyedObjectContext {
	var p = new(OdinKeyedObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinKeyedObject

	return p
}

func (s *OdinKeyedObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinKeyedObjectContext) OdinKeySpec() IOdinKeySpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinKeySpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinKeySpecContext)
}

func (s *OdinKeyedObjectContext) OdinObjectBlock() IOdinObjectBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinObjectBlockContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.OdinKeySpec()
	}
	{
		p.SetState(148)
		p.Match(OdinParserSYM_EQ)
	}
	{
		p.SetState(149)
		p.OdinObjectBlock()
	}

	return localctx
}

// IOdinKeySpecContext is an interface to support dynamic dispatch.
type IOdinKeySpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinKeySpecContext differentiates from other interfaces.
	IsOdinKeySpecContext()
}

type OdinKeySpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinKeySpecContext() *OdinKeySpecContext {
	var p = new(OdinKeySpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinKeySpec
	return p
}

func (*OdinKeySpecContext) IsOdinKeySpecContext() {}

func NewOdinKeySpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinKeySpecContext {
	var p = new(OdinKeySpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinKeySpec

	return p
}

func (s *OdinKeySpecContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinKeySpecContext) PrimitiveValue() IPrimitiveValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(OdinParserSYM_LBRACKET)
	}
	{
		p.SetState(152)
		p.PrimitiveValue()
	}
	{
		p.SetState(153)
		p.Match(OdinParserSYM_RBRACKET)
	}

	return localctx
}

// IOdinObjectReferenceBlockContext is an interface to support dynamic dispatch.
type IOdinObjectReferenceBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinObjectReferenceBlockContext differentiates from other interfaces.
	IsOdinObjectReferenceBlockContext()
}

type OdinObjectReferenceBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinObjectReferenceBlockContext() *OdinObjectReferenceBlockContext {
	var p = new(OdinObjectReferenceBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinObjectReferenceBlock
	return p
}

func (*OdinObjectReferenceBlockContext) IsOdinObjectReferenceBlockContext() {}

func NewOdinObjectReferenceBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinObjectReferenceBlockContext {
	var p = new(OdinObjectReferenceBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinObjectReferenceBlock

	return p
}

func (s *OdinObjectReferenceBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinObjectReferenceBlockContext) OdinPathList() IOdinPathListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinPathListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinPathListContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(OdinParserSYM_LT)
	}
	{
		p.SetState(156)
		p.OdinPathList()
	}
	{
		p.SetState(157)
		p.Match(OdinParserSYM_GT)
	}

	return localctx
}

// IOdinPathListContext is an interface to support dynamic dispatch.
type IOdinPathListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinPathListContext differentiates from other interfaces.
	IsOdinPathListContext()
}

type OdinPathListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathListContext() *OdinPathListContext {
	var p = new(OdinPathListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathList
	return p
}

func (*OdinPathListContext) IsOdinPathListContext() {}

func NewOdinPathListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathListContext {
	var p = new(OdinPathListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPathList

	return p
}

func (s *OdinPathListContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathListContext) AllOdinPath() []IOdinPathContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOdinPathContext)(nil)).Elem())
	var tst = make([]IOdinPathContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOdinPathContext)
		}
	}

	return tst
}

func (s *OdinPathListContext) OdinPath(i int) IOdinPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinPathContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOdinPathContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.OdinPath()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(161)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 2 {
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(162)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(163)
				p.OdinPath()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IOdinPathContext is an interface to support dynamic dispatch.
type IOdinPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinPathContext differentiates from other interfaces.
	IsOdinPathContext()
}

type OdinPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathContext() *OdinPathContext {
	var p = new(OdinPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinPath
	return p
}

func (*OdinPathContext) IsOdinPathContext() {}

func NewOdinPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathContext {
	var p = new(OdinPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPath

	return p
}

func (s *OdinPathContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathContext) OdinKeySpec() IOdinKeySpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinKeySpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOdinKeySpecContext)
}

func (s *OdinPathContext) AllOdinPathSegment() []IOdinPathSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOdinPathSegmentContext)(nil)).Elem())
	var tst = make([]IOdinPathSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOdinPathSegmentContext)
		}
	}

	return tst
}

func (s *OdinPathContext) OdinPathSegment(i int) IOdinPathSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinPathSegmentContext)(nil)).Elem(), i)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LBRACKET {
		{
			p.SetState(170)
			p.OdinKeySpec()
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OdinParserSYM_SLASH {
		{
			p.SetState(173)
			p.OdinPathSegment()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOdinPathSegmentContext is an interface to support dynamic dispatch.
type IOdinPathSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdinPathSegmentContext differentiates from other interfaces.
	IsOdinPathSegmentContext()
}

type OdinPathSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdinPathSegmentContext() *OdinPathSegmentContext {
	var p = new(OdinPathSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_odinPathSegment
	return p
}

func (*OdinPathSegmentContext) IsOdinPathSegmentContext() {}

func NewOdinPathSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdinPathSegmentContext {
	var p = new(OdinPathSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_odinPathSegment

	return p
}

func (s *OdinPathSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *OdinPathSegmentContext) LC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserLC_ID, 0)
}

func (s *OdinPathSegmentContext) OdinKeySpec() IOdinKeySpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOdinKeySpecContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(OdinParserSYM_SLASH)
	}
	{
		p.SetState(179)
		p.Match(OdinParserLC_ID)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LBRACKET {
		{
			p.SetState(180)
			p.OdinKeySpec()
		}

	}

	return localctx
}

// IRmTypeIdContext is an interface to support dynamic dispatch.
type IRmTypeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRmTypeIdContext differentiates from other interfaces.
	IsRmTypeIdContext()
}

type RmTypeIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmTypeIdContext() *RmTypeIdContext {
	var p = new(RmTypeIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_rmTypeId
	return p
}

func (*RmTypeIdContext) IsRmTypeIdContext() {}

func NewRmTypeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmTypeIdContext {
	var p = new(RmTypeIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_rmTypeId

	return p
}

func (s *RmTypeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *RmTypeIdContext) UC_ID() antlr.TerminalNode {
	return s.GetToken(OdinParserUC_ID, 0)
}

func (s *RmTypeIdContext) AllRmTypeId() []IRmTypeIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRmTypeIdContext)(nil)).Elem())
	var tst = make([]IRmTypeIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRmTypeIdContext)
		}
	}

	return tst
}

func (s *RmTypeIdContext) RmTypeId(i int) IRmTypeIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmTypeIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRmTypeIdContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(OdinParserUC_ID)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_LT {
		{
			p.SetState(184)
			p.Match(OdinParserSYM_LT)
		}
		{
			p.SetState(185)
			p.RmTypeId()
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == OdinParserSYM_COMMA {
			{
				p.SetState(186)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(187)
				p.RmTypeId()
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(193)
			p.Match(OdinParserSYM_GT)
		}

	}

	return localctx
}

// IPrimitiveObjectContext is an interface to support dynamic dispatch.
type IPrimitiveObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveObjectContext differentiates from other interfaces.
	IsPrimitiveObjectContext()
}

type PrimitiveObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveObjectContext() *PrimitiveObjectContext {
	var p = new(PrimitiveObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveObject
	return p
}

func (*PrimitiveObjectContext) IsPrimitiveObjectContext() {}

func NewPrimitiveObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveObjectContext {
	var p = new(PrimitiveObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveObject

	return p
}

func (s *PrimitiveObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveObjectContext) PrimitiveValue() IPrimitiveValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveValueContext)
}

func (s *PrimitiveObjectContext) PrimitiveListValue() IPrimitiveListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveListValueContext)
}

func (s *PrimitiveObjectContext) PrimitiveIntervalValue() IPrimitiveIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveIntervalValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
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

	}

	return localctx
}

// IPrimitiveValueContext is an interface to support dynamic dispatch.
type IPrimitiveValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveValueContext differentiates from other interfaces.
	IsPrimitiveValueContext()
}

type PrimitiveValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveValueContext() *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveValue
	return p
}

func (*PrimitiveValueContext) IsPrimitiveValueContext() {}

func NewPrimitiveValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveValueContext {
	var p = new(PrimitiveValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveValue

	return p
}

func (s *PrimitiveValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *PrimitiveValueContext) IntegerValue() IIntegerValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *PrimitiveValueContext) RealValue() IRealValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
}

func (s *PrimitiveValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *PrimitiveValueContext) CharacterValue() ICharacterValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterValueContext)
}

func (s *PrimitiveValueContext) TermCodeValue() ITermCodeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermCodeValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermCodeValueContext)
}

func (s *PrimitiveValueContext) DateValue() IDateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *PrimitiveValueContext) TimeValue() ITimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *PrimitiveValueContext) DateTimeValue() IDateTimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
}

func (s *PrimitiveValueContext) DurationValue() IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
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

	}

	return localctx
}

// IPrimitiveListValueContext is an interface to support dynamic dispatch.
type IPrimitiveListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveListValueContext differentiates from other interfaces.
	IsPrimitiveListValueContext()
}

type PrimitiveListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveListValueContext() *PrimitiveListValueContext {
	var p = new(PrimitiveListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveListValue
	return p
}

func (*PrimitiveListValueContext) IsPrimitiveListValueContext() {}

func NewPrimitiveListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveListValueContext {
	var p = new(PrimitiveListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveListValue

	return p
}

func (s *PrimitiveListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveListValueContext) StringListValue() IStringListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringListValueContext)
}

func (s *PrimitiveListValueContext) IntegerListValue() IIntegerListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerListValueContext)
}

func (s *PrimitiveListValueContext) RealListValue() IRealListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealListValueContext)
}

func (s *PrimitiveListValueContext) BooleanListValue() IBooleanListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanListValueContext)
}

func (s *PrimitiveListValueContext) CharacterListValue() ICharacterListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterListValueContext)
}

func (s *PrimitiveListValueContext) TermCodeListValue() ITermCodeListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermCodeListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermCodeListValueContext)
}

func (s *PrimitiveListValueContext) DateListValue() IDateListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateListValueContext)
}

func (s *PrimitiveListValueContext) TimeListValue() ITimeListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeListValueContext)
}

func (s *PrimitiveListValueContext) DateTimeListValue() IDateTimeListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateTimeListValueContext)
}

func (s *PrimitiveListValueContext) DurationListValue() IDurationListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationListValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
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

	}

	return localctx
}

// IPrimitiveIntervalValueContext is an interface to support dynamic dispatch.
type IPrimitiveIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveIntervalValueContext differentiates from other interfaces.
	IsPrimitiveIntervalValueContext()
}

type PrimitiveIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveIntervalValueContext() *PrimitiveIntervalValueContext {
	var p = new(PrimitiveIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_primitiveIntervalValue
	return p
}

func (*PrimitiveIntervalValueContext) IsPrimitiveIntervalValueContext() {}

func NewPrimitiveIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveIntervalValueContext {
	var p = new(PrimitiveIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_primitiveIntervalValue

	return p
}

func (s *PrimitiveIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveIntervalValueContext) IntegerIntervalValue() IIntegerIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerIntervalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) RealIntervalValue() IRealIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealIntervalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DateIntervalValue() IDateIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateIntervalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) TimeIntervalValue() ITimeIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DateTimeIntervalValue() IDateTimeIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeIntervalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateTimeIntervalValueContext)
}

func (s *PrimitiveIntervalValueContext) DurationIntervalValue() IDurationIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationIntervalValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
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

	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(OdinParserSTRING)
	}

	return localctx
}

// IStringListValueContext is an interface to support dynamic dispatch.
type IStringListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringListValueContext differentiates from other interfaces.
	IsStringListValueContext()
}

type StringListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListValueContext() *StringListValueContext {
	var p = new(StringListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_stringListValue
	return p
}

func (*StringListValueContext) IsStringListValueContext() {}

func NewStringListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListValueContext {
	var p = new(StringListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_stringListValue

	return p
}

func (s *StringListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListValueContext) AllStringValue() []IStringValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringValueContext)(nil)).Elem())
	var tst = make([]IStringValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringValueContext)
		}
	}

	return tst
}

func (s *StringListValueContext) StringValue(i int) IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.StringValue()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(237)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(238)
				p.StringValue()
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(243)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(244)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_integerValue
	return p
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		p.SetState(247)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	p.SetState(250)
	_la = p.GetTokenStream().LA(1)

	if !(_la == OdinParserINTEGER || _la == OdinParserSCI_INTEGER) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIntegerListValueContext is an interface to support dynamic dispatch.
type IIntegerListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerListValueContext differentiates from other interfaces.
	IsIntegerListValueContext()
}

type IntegerListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerListValueContext() *IntegerListValueContext {
	var p = new(IntegerListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_integerListValue
	return p
}

func (*IntegerListValueContext) IsIntegerListValueContext() {}

func NewIntegerListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerListValueContext {
	var p = new(IntegerListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerListValue

	return p
}

func (s *IntegerListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerListValueContext) AllIntegerValue() []IIntegerValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem())
	var tst = make([]IIntegerValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerValueContext)
		}
	}

	return tst
}

func (s *IntegerListValueContext) IntegerValue(i int) IIntegerValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.IntegerValue()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(253)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(254)
				p.IntegerValue()
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(259)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(260)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IIntegerIntervalValueContext is an interface to support dynamic dispatch.
type IIntegerIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerIntervalValueContext differentiates from other interfaces.
	IsIntegerIntervalValueContext()
}

type IntegerIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerIntervalValueContext() *IntegerIntervalValueContext {
	var p = new(IntegerIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalValue
	return p
}

func (*IntegerIntervalValueContext) IsIntegerIntervalValueContext() {}

func NewIntegerIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerIntervalValueContext {
	var p = new(IntegerIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerIntervalValue

	return p
}

func (s *IntegerIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerIntervalValueContext) AllIntegerValue() []IIntegerValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem())
	var tst = make([]IIntegerValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerValueContext)
		}
	}

	return tst
}

func (s *IntegerIntervalValueContext) IntegerValue(i int) IIntegerValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *IntegerIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *IntegerIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *IntegerIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(264)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(267)
			p.IntegerValue()
		}
		{
			p.SetState(268)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(269)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(272)
			p.IntegerValue()
		}
		{
			p.SetState(273)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(283)
			p.IntegerValue()
		}
		{
			p.SetState(284)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(285)
			p.IntegerValue()
		}
		{
			p.SetState(286)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// IIntegerIntervalListValueContext is an interface to support dynamic dispatch.
type IIntegerIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerIntervalListValueContext differentiates from other interfaces.
	IsIntegerIntervalListValueContext()
}

type IntegerIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerIntervalListValueContext() *IntegerIntervalListValueContext {
	var p = new(IntegerIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_integerIntervalListValue
	return p
}

func (*IntegerIntervalListValueContext) IsIntegerIntervalListValueContext() {}

func NewIntegerIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerIntervalListValueContext {
	var p = new(IntegerIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_integerIntervalListValue

	return p
}

func (s *IntegerIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerIntervalListValueContext) AllIntegerIntervalValue() []IIntegerIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerIntervalValueContext)(nil)).Elem())
	var tst = make([]IIntegerIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerIntervalValueContext)
		}
	}

	return tst
}

func (s *IntegerIntervalListValueContext) IntegerIntervalValue(i int) IIntegerIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.IntegerIntervalValue()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(291)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(292)
				p.IntegerIntervalValue()
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(297)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(298)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IRealValueContext is an interface to support dynamic dispatch.
type IRealValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealValueContext differentiates from other interfaces.
	IsRealValueContext()
}

type RealValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealValueContext() *RealValueContext {
	var p = new(RealValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_realValue
	return p
}

func (*RealValueContext) IsRealValueContext() {}

func NewRealValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealValueContext {
	var p = new(RealValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		p.SetState(301)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	p.SetState(304)
	_la = p.GetTokenStream().LA(1)

	if !(_la == OdinParserREAL || _la == OdinParserSCI_REAL) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IRealListValueContext is an interface to support dynamic dispatch.
type IRealListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealListValueContext differentiates from other interfaces.
	IsRealListValueContext()
}

type RealListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealListValueContext() *RealListValueContext {
	var p = new(RealListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_realListValue
	return p
}

func (*RealListValueContext) IsRealListValueContext() {}

func NewRealListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealListValueContext {
	var p = new(RealListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realListValue

	return p
}

func (s *RealListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealListValueContext) AllRealValue() []IRealValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealValueContext)(nil)).Elem())
	var tst = make([]IRealValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealValueContext)
		}
	}

	return tst
}

func (s *RealListValueContext) RealValue(i int) IRealValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.RealValue()
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(307)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(308)
				p.RealValue()
			}

			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(313)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(314)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IRealIntervalValueContext is an interface to support dynamic dispatch.
type IRealIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealIntervalValueContext differentiates from other interfaces.
	IsRealIntervalValueContext()
}

type RealIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealIntervalValueContext() *RealIntervalValueContext {
	var p = new(RealIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalValue
	return p
}

func (*RealIntervalValueContext) IsRealIntervalValueContext() {}

func NewRealIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealIntervalValueContext {
	var p = new(RealIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realIntervalValue

	return p
}

func (s *RealIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealIntervalValueContext) AllRealValue() []IRealValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealValueContext)(nil)).Elem())
	var tst = make([]IRealValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealValueContext)
		}
	}

	return tst
}

func (s *RealIntervalValueContext) RealValue(i int) IRealValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealValueContext)
}

func (s *RealIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *RealIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *RealIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(318)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(321)
			p.RealValue()
		}
		{
			p.SetState(322)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(323)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(326)
			p.RealValue()
		}
		{
			p.SetState(327)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(337)
			p.RealValue()
		}
		{
			p.SetState(338)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(339)
			p.RealValue()
		}
		{
			p.SetState(340)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// IRealIntervalListValueContext is an interface to support dynamic dispatch.
type IRealIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealIntervalListValueContext differentiates from other interfaces.
	IsRealIntervalListValueContext()
}

type RealIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealIntervalListValueContext() *RealIntervalListValueContext {
	var p = new(RealIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_realIntervalListValue
	return p
}

func (*RealIntervalListValueContext) IsRealIntervalListValueContext() {}

func NewRealIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealIntervalListValueContext {
	var p = new(RealIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_realIntervalListValue

	return p
}

func (s *RealIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RealIntervalListValueContext) AllRealIntervalValue() []IRealIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealIntervalValueContext)(nil)).Elem())
	var tst = make([]IRealIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealIntervalValueContext)
		}
	}

	return tst
}

func (s *RealIntervalListValueContext) RealIntervalValue(i int) IRealIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.RealIntervalValue()
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(345)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(346)
				p.RealIntervalValue()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(351)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(352)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(355)
	_la = p.GetTokenStream().LA(1)

	if !(_la == OdinParserSYM_TRUE || _la == OdinParserSYM_FALSE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBooleanListValueContext is an interface to support dynamic dispatch.
type IBooleanListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanListValueContext differentiates from other interfaces.
	IsBooleanListValueContext()
}

type BooleanListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanListValueContext() *BooleanListValueContext {
	var p = new(BooleanListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_booleanListValue
	return p
}

func (*BooleanListValueContext) IsBooleanListValueContext() {}

func NewBooleanListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanListValueContext {
	var p = new(BooleanListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_booleanListValue

	return p
}

func (s *BooleanListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanListValueContext) AllBooleanValue() []IBooleanValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem())
	var tst = make([]IBooleanValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBooleanValueContext)
		}
	}

	return tst
}

func (s *BooleanListValueContext) BooleanValue(i int) IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.BooleanValue()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(358)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(359)
				p.BooleanValue()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(364)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(365)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// ICharacterValueContext is an interface to support dynamic dispatch.
type ICharacterValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterValueContext differentiates from other interfaces.
	IsCharacterValueContext()
}

type CharacterValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterValueContext() *CharacterValueContext {
	var p = new(CharacterValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_characterValue
	return p
}

func (*CharacterValueContext) IsCharacterValueContext() {}

func NewCharacterValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterValueContext {
	var p = new(CharacterValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(OdinParserCHARACTER)
	}

	return localctx
}

// ICharacterListValueContext is an interface to support dynamic dispatch.
type ICharacterListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterListValueContext differentiates from other interfaces.
	IsCharacterListValueContext()
}

type CharacterListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterListValueContext() *CharacterListValueContext {
	var p = new(CharacterListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_characterListValue
	return p
}

func (*CharacterListValueContext) IsCharacterListValueContext() {}

func NewCharacterListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterListValueContext {
	var p = new(CharacterListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_characterListValue

	return p
}

func (s *CharacterListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterListValueContext) AllCharacterValue() []ICharacterValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterValueContext)(nil)).Elem())
	var tst = make([]ICharacterValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterValueContext)
		}
	}

	return tst
}

func (s *CharacterListValueContext) CharacterValue(i int) ICharacterValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.CharacterValue()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(371)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(372)
				p.CharacterValue()
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(377)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(378)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDateValueContext is an interface to support dynamic dispatch.
type IDateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateValueContext differentiates from other interfaces.
	IsDateValueContext()
}

type DateValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateValueContext() *DateValueContext {
	var p = new(DateValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateValue
	return p
}

func (*DateValueContext) IsDateValueContext() {}

func NewDateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateValueContext {
	var p = new(DateValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(OdinParserISO8601_DATE_AUGMENTED)
	}

	return localctx
}

// IDateListValueContext is an interface to support dynamic dispatch.
type IDateListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateListValueContext differentiates from other interfaces.
	IsDateListValueContext()
}

type DateListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateListValueContext() *DateListValueContext {
	var p = new(DateListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateListValue
	return p
}

func (*DateListValueContext) IsDateListValueContext() {}

func NewDateListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateListValueContext {
	var p = new(DateListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateListValue

	return p
}

func (s *DateListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateListValueContext) AllDateValue() []IDateValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateValueContext)(nil)).Elem())
	var tst = make([]IDateValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateValueContext)
		}
	}

	return tst
}

func (s *DateListValueContext) DateValue(i int) IDateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.DateValue()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(384)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(385)
				p.DateValue()
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(390)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(391)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDateIntervalValueContext is an interface to support dynamic dispatch.
type IDateIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateIntervalValueContext differentiates from other interfaces.
	IsDateIntervalValueContext()
}

type DateIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateIntervalValueContext() *DateIntervalValueContext {
	var p = new(DateIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalValue
	return p
}

func (*DateIntervalValueContext) IsDateIntervalValueContext() {}

func NewDateIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateIntervalValueContext {
	var p = new(DateIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateIntervalValue

	return p
}

func (s *DateIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateIntervalValueContext) AllDateValue() []IDateValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateValueContext)(nil)).Elem())
	var tst = make([]IDateValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateValueContext)
		}
	}

	return tst
}

func (s *DateIntervalValueContext) DateValue(i int) IDateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *DateIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DateIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DateIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *DateIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *DateIntervalValueContext) DurationValue() IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(395)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(398)
			p.DateValue()
		}
		{
			p.SetState(399)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(400)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(403)
			p.DateValue()
		}
		{
			p.SetState(404)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(414)
			p.DateValue()
		}
		{
			p.SetState(415)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(416)
			p.DurationValue()
		}
		{
			p.SetState(417)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// IDateIntervalListValueContext is an interface to support dynamic dispatch.
type IDateIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateIntervalListValueContext differentiates from other interfaces.
	IsDateIntervalListValueContext()
}

type DateIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateIntervalListValueContext() *DateIntervalListValueContext {
	var p = new(DateIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateIntervalListValue
	return p
}

func (*DateIntervalListValueContext) IsDateIntervalListValueContext() {}

func NewDateIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateIntervalListValueContext {
	var p = new(DateIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateIntervalListValue

	return p
}

func (s *DateIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateIntervalListValueContext) AllDateIntervalValue() []IDateIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateIntervalValueContext)(nil)).Elem())
	var tst = make([]IDateIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateIntervalValueContext)
		}
	}

	return tst
}

func (s *DateIntervalListValueContext) DateIntervalValue(i int) IDateIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.DateIntervalValue()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(422)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(423)
				p.DateIntervalValue()
			}

			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(428)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(429)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// ITimeValueContext is an interface to support dynamic dispatch.
type ITimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeValueContext differentiates from other interfaces.
	IsTimeValueContext()
}

type TimeValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeValueContext() *TimeValueContext {
	var p = new(TimeValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_timeValue
	return p
}

func (*TimeValueContext) IsTimeValueContext() {}

func NewTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeValueContext {
	var p = new(TimeValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(OdinParserISO8601_TIME_AUGMENTED)
	}

	return localctx
}

// ITimeListValueContext is an interface to support dynamic dispatch.
type ITimeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeListValueContext differentiates from other interfaces.
	IsTimeListValueContext()
}

type TimeListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeListValueContext() *TimeListValueContext {
	var p = new(TimeListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_timeListValue
	return p
}

func (*TimeListValueContext) IsTimeListValueContext() {}

func NewTimeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeListValueContext {
	var p = new(TimeListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeListValue

	return p
}

func (s *TimeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeListValueContext) AllTimeValue() []ITimeValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeValueContext)(nil)).Elem())
	var tst = make([]ITimeValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeValueContext)
		}
	}

	return tst
}

func (s *TimeListValueContext) TimeValue(i int) ITimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.TimeValue()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(435)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(436)
				p.TimeValue()
			}

			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(441)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(442)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// ITimeIntervalValueContext is an interface to support dynamic dispatch.
type ITimeIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIntervalValueContext differentiates from other interfaces.
	IsTimeIntervalValueContext()
}

type TimeIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalValueContext() *TimeIntervalValueContext {
	var p = new(TimeIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalValue
	return p
}

func (*TimeIntervalValueContext) IsTimeIntervalValueContext() {}

func NewTimeIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalValueContext {
	var p = new(TimeIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeIntervalValue

	return p
}

func (s *TimeIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalValueContext) AllTimeValue() []ITimeValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeValueContext)(nil)).Elem())
	var tst = make([]ITimeValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeValueContext)
		}
	}

	return tst
}

func (s *TimeIntervalValueContext) TimeValue(i int) ITimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *TimeIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *TimeIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *TimeIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *TimeIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *TimeIntervalValueContext) DurationValue() IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(446)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(449)
			p.TimeValue()
		}
		{
			p.SetState(450)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(451)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(454)
			p.TimeValue()
		}
		{
			p.SetState(455)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(464)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(465)
			p.TimeValue()
		}
		{
			p.SetState(466)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(467)
			p.DurationValue()
		}
		{
			p.SetState(468)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// ITimeIntervalListValueContext is an interface to support dynamic dispatch.
type ITimeIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIntervalListValueContext differentiates from other interfaces.
	IsTimeIntervalListValueContext()
}

type TimeIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalListValueContext() *TimeIntervalListValueContext {
	var p = new(TimeIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_timeIntervalListValue
	return p
}

func (*TimeIntervalListValueContext) IsTimeIntervalListValueContext() {}

func NewTimeIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalListValueContext {
	var p = new(TimeIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_timeIntervalListValue

	return p
}

func (s *TimeIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalListValueContext) AllTimeIntervalValue() []ITimeIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeIntervalValueContext)(nil)).Elem())
	var tst = make([]ITimeIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeIntervalValueContext)
		}
	}

	return tst
}

func (s *TimeIntervalListValueContext) TimeIntervalValue(i int) ITimeIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.TimeIntervalValue()
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(473)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(474)
				p.TimeIntervalValue()
			}

			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(479)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(480)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDateTimeValueContext is an interface to support dynamic dispatch.
type IDateTimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimeValueContext differentiates from other interfaces.
	IsDateTimeValueContext()
}

type DateTimeValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeValueContext() *DateTimeValueContext {
	var p = new(DateTimeValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeValue
	return p
}

func (*DateTimeValueContext) IsDateTimeValueContext() {}

func NewDateTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeValueContext {
	var p = new(DateTimeValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(OdinParserISO8601_DATE_TIME_AUGMENTED)
	}

	return localctx
}

// IDateTimeListValueContext is an interface to support dynamic dispatch.
type IDateTimeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimeListValueContext differentiates from other interfaces.
	IsDateTimeListValueContext()
}

type DateTimeListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeListValueContext() *DateTimeListValueContext {
	var p = new(DateTimeListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeListValue
	return p
}

func (*DateTimeListValueContext) IsDateTimeListValueContext() {}

func NewDateTimeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeListValueContext {
	var p = new(DateTimeListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeListValue

	return p
}

func (s *DateTimeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeListValueContext) AllDateTimeValue() []IDateTimeValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateTimeValueContext)(nil)).Elem())
	var tst = make([]IDateTimeValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateTimeValueContext)
		}
	}

	return tst
}

func (s *DateTimeListValueContext) DateTimeValue(i int) IDateTimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.DateTimeValue()
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(486)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(487)
				p.DateTimeValue()
			}

			p.SetState(490)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(492)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(493)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDateTimeIntervalValueContext is an interface to support dynamic dispatch.
type IDateTimeIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimeIntervalValueContext differentiates from other interfaces.
	IsDateTimeIntervalValueContext()
}

type DateTimeIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeIntervalValueContext() *DateTimeIntervalValueContext {
	var p = new(DateTimeIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalValue
	return p
}

func (*DateTimeIntervalValueContext) IsDateTimeIntervalValueContext() {}

func NewDateTimeIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeIntervalValueContext {
	var p = new(DateTimeIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeIntervalValue

	return p
}

func (s *DateTimeIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeIntervalValueContext) AllDateTimeValue() []IDateTimeValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateTimeValueContext)(nil)).Elem())
	var tst = make([]IDateTimeValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateTimeValueContext)
		}
	}

	return tst
}

func (s *DateTimeIntervalValueContext) DateTimeValue(i int) IDateTimeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateTimeValueContext)
}

func (s *DateTimeIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DateTimeIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DateTimeIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *DateTimeIntervalValueContext) SYM_PLUS_OR_MINUS() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_PLUS_OR_MINUS, 0)
}

func (s *DateTimeIntervalValueContext) DurationValue() IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(497)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(500)
			p.DateTimeValue()
		}
		{
			p.SetState(501)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(502)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(505)
			p.DateTimeValue()
		}
		{
			p.SetState(506)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(515)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(516)
			p.DateTimeValue()
		}
		{
			p.SetState(517)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(518)
			p.DurationValue()
		}
		{
			p.SetState(519)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// IDateTimeIntervalListValueContext is an interface to support dynamic dispatch.
type IDateTimeIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimeIntervalListValueContext differentiates from other interfaces.
	IsDateTimeIntervalListValueContext()
}

type DateTimeIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeIntervalListValueContext() *DateTimeIntervalListValueContext {
	var p = new(DateTimeIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_dateTimeIntervalListValue
	return p
}

func (*DateTimeIntervalListValueContext) IsDateTimeIntervalListValueContext() {}

func NewDateTimeIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeIntervalListValueContext {
	var p = new(DateTimeIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_dateTimeIntervalListValue

	return p
}

func (s *DateTimeIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeIntervalListValueContext) AllDateTimeIntervalValue() []IDateTimeIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateTimeIntervalValueContext)(nil)).Elem())
	var tst = make([]IDateTimeIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateTimeIntervalValueContext)
		}
	}

	return tst
}

func (s *DateTimeIntervalListValueContext) DateTimeIntervalValue(i int) IDateTimeIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimeIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateTimeIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.DateTimeIntervalValue()
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(524)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(525)
				p.DateTimeIntervalValue()
			}

			p.SetState(528)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(530)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(531)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDurationValueContext is an interface to support dynamic dispatch.
type IDurationValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationValueContext differentiates from other interfaces.
	IsDurationValueContext()
}

type DurationValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationValueContext() *DurationValueContext {
	var p = new(DurationValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_durationValue
	return p
}

func (*DurationValueContext) IsDurationValueContext() {}

func NewDurationValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationValueContext {
	var p = new(DurationValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS {
		p.SetState(534)
		_la = p.GetTokenStream().LA(1)

		if !(_la == OdinParserSYM_PLUS || _la == OdinParserSYM_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(537)
		p.Match(OdinParserISO8601_DURATION)
	}

	return localctx
}

// IDurationListValueContext is an interface to support dynamic dispatch.
type IDurationListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationListValueContext differentiates from other interfaces.
	IsDurationListValueContext()
}

type DurationListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationListValueContext() *DurationListValueContext {
	var p = new(DurationListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_durationListValue
	return p
}

func (*DurationListValueContext) IsDurationListValueContext() {}

func NewDurationListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationListValueContext {
	var p = new(DurationListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationListValue

	return p
}

func (s *DurationListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationListValueContext) AllDurationValue() []IDurationValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDurationValueContext)(nil)).Elem())
	var tst = make([]IDurationValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDurationValueContext)
		}
	}

	return tst
}

func (s *DurationListValueContext) DurationValue(i int) IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.DurationValue()
	}
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 70, p.GetParserRuleContext()) {
	case 1:
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(540)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(541)
				p.DurationValue()
			}

			p.SetState(544)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(546)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(547)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IDurationIntervalValueContext is an interface to support dynamic dispatch.
type IDurationIntervalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationIntervalValueContext differentiates from other interfaces.
	IsDurationIntervalValueContext()
}

type DurationIntervalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationIntervalValueContext() *DurationIntervalValueContext {
	var p = new(DurationIntervalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalValue
	return p
}

func (*DurationIntervalValueContext) IsDurationIntervalValueContext() {}

func NewDurationIntervalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationIntervalValueContext {
	var p = new(DurationIntervalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationIntervalValue

	return p
}

func (s *DurationIntervalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationIntervalValueContext) AllDurationValue() []IDurationValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDurationValueContext)(nil)).Elem())
	var tst = make([]IDurationValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDurationValueContext)
		}
	}

	return tst
}

func (s *DurationIntervalValueContext) DurationValue(i int) IDurationValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDurationValueContext)
}

func (s *DurationIntervalValueContext) SYM_GT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_GT, 0)
}

func (s *DurationIntervalValueContext) SYM_LT() antlr.TerminalNode {
	return s.GetToken(OdinParserSYM_LT, 0)
}

func (s *DurationIntervalValueContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 74, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_GT {
			{
				p.SetState(551)
				p.Match(OdinParserSYM_GT)
			}

		}
		{
			p.SetState(554)
			p.DurationValue()
		}
		{
			p.SetState(555)
			p.Match(OdinParserSYM_IVL_SEP)
		}
		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == OdinParserSYM_LT {
			{
				p.SetState(556)
				p.Match(OdinParserSYM_LT)
			}

		}
		{
			p.SetState(559)
			p.DurationValue()
		}
		{
			p.SetState(560)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(562)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0 {
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
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(569)
			p.Match(OdinParserSYM_IVL_DELIM)
		}
		{
			p.SetState(570)
			p.DurationValue()
		}
		{
			p.SetState(571)
			p.Match(OdinParserSYM_PLUS_OR_MINUS)
		}
		{
			p.SetState(572)
			p.DurationValue()
		}
		{
			p.SetState(573)
			p.Match(OdinParserSYM_IVL_DELIM)
		}

	}

	return localctx
}

// IDurationIntervalListValueContext is an interface to support dynamic dispatch.
type IDurationIntervalListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationIntervalListValueContext differentiates from other interfaces.
	IsDurationIntervalListValueContext()
}

type DurationIntervalListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationIntervalListValueContext() *DurationIntervalListValueContext {
	var p = new(DurationIntervalListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_durationIntervalListValue
	return p
}

func (*DurationIntervalListValueContext) IsDurationIntervalListValueContext() {}

func NewDurationIntervalListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationIntervalListValueContext {
	var p = new(DurationIntervalListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_durationIntervalListValue

	return p
}

func (s *DurationIntervalListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationIntervalListValueContext) AllDurationIntervalValue() []IDurationIntervalValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDurationIntervalValueContext)(nil)).Elem())
	var tst = make([]IDurationIntervalValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDurationIntervalValueContext)
		}
	}

	return tst
}

func (s *DurationIntervalListValueContext) DurationIntervalValue(i int) IDurationIntervalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationIntervalValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDurationIntervalValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(577)
		p.DurationIntervalValue()
	}
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext()) {
	case 1:
		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(578)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(579)
				p.DurationIntervalValue()
			}

			p.SetState(582)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(584)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(585)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// ITermCodeValueContext is an interface to support dynamic dispatch.
type ITermCodeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermCodeValueContext differentiates from other interfaces.
	IsTermCodeValueContext()
}

type TermCodeValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermCodeValueContext() *TermCodeValueContext {
	var p = new(TermCodeValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeValue
	return p
}

func (*TermCodeValueContext) IsTermCodeValueContext() {}

func NewTermCodeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermCodeValueContext {
	var p = new(TermCodeValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(588)
		p.Match(OdinParserQUALIFIED_TERM_CODE_REF)
	}

	return localctx
}

// ITermCodeListValueContext is an interface to support dynamic dispatch.
type ITermCodeListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermCodeListValueContext differentiates from other interfaces.
	IsTermCodeListValueContext()
}

type TermCodeListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermCodeListValueContext() *TermCodeListValueContext {
	var p = new(TermCodeListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_termCodeListValue
	return p
}

func (*TermCodeListValueContext) IsTermCodeListValueContext() {}

func NewTermCodeListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermCodeListValueContext {
	var p = new(TermCodeListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdinParserRULE_termCodeListValue

	return p
}

func (s *TermCodeListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TermCodeListValueContext) AllTermCodeValue() []ITermCodeValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermCodeValueContext)(nil)).Elem())
	var tst = make([]ITermCodeValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermCodeValueContext)
		}
	}

	return tst
}

func (s *TermCodeListValueContext) TermCodeValue(i int) ITermCodeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermCodeValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermCodeValueContext)
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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.TermCodeValue()
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OdinParserSYM_COMMA {
			{
				p.SetState(591)
				p.Match(OdinParserSYM_COMMA)
			}
			{
				p.SetState(592)
				p.TermCodeValue()
			}

			p.SetState(595)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(597)
			p.Match(OdinParserSYM_COMMA)
		}
		{
			p.SetState(598)
			p.Match(OdinParserSYM_LIST_CONTINUE)
		}

	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdinParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(601)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OdinParserSYM_LE)|(1<<OdinParserSYM_GE)|(1<<OdinParserSYM_GT)|(1<<OdinParserSYM_LT))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
