// Code generated from ../grammar/ImagemonkeyQueryLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package imagemonkeyquerylang

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 374,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 5, 21, 210, 10, 21, 3, 22, 6, 22, 213, 10, 22, 13,
	22, 14, 22, 214, 3, 22, 3, 22, 6, 22, 219, 10, 22, 13, 22, 14, 22, 220,
	5, 22, 223, 10, 22, 3, 22, 7, 22, 226, 10, 22, 12, 22, 14, 22, 229, 11,
	22, 3, 22, 3, 22, 7, 22, 233, 10, 22, 12, 22, 14, 22, 236, 11, 22, 3, 22,
	3, 22, 6, 22, 240, 10, 22, 13, 22, 14, 22, 241, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 6, 23, 252, 10, 23, 13, 23, 14, 23, 253,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 6, 24, 261, 10, 24, 13, 24, 14, 24,
	262, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 6, 24, 277, 10, 24, 13, 24, 14, 24, 278, 3, 24, 3, 24,
	3, 25, 3, 25, 6, 25, 285, 10, 25, 13, 25, 14, 25, 286, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 6, 25,
	301, 10, 25, 13, 25, 14, 25, 302, 3, 25, 3, 25, 3, 26, 3, 26, 6, 26, 309,
	10, 26, 13, 26, 14, 26, 310, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 7, 27, 331, 10, 27, 12, 27, 14, 27, 334, 11, 27, 3, 27, 3, 27,
	5, 27, 338, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 6, 29, 354, 10, 29, 13, 29,
	14, 29, 355, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 35, 6, 35, 369, 10, 35, 13, 35, 14, 35, 370, 3, 35, 3, 35,
	2, 2, 36, 3, 2, 5, 2, 7, 2, 9, 2, 11, 2, 13, 2, 15, 2, 17, 2, 19, 2, 21,
	2, 23, 2, 25, 3, 27, 4, 29, 5, 31, 6, 33, 7, 35, 8, 37, 9, 39, 10, 41,
	11, 43, 12, 45, 13, 47, 14, 49, 15, 51, 16, 53, 17, 55, 18, 57, 19, 59,
	20, 61, 21, 63, 22, 65, 23, 67, 24, 69, 25, 3, 2, 19, 3, 2, 99, 124, 3,
	2, 67, 92, 4, 2, 67, 92, 99, 124, 5, 2, 34, 34, 67, 92, 99, 124, 5, 2,
	50, 59, 67, 92, 99, 124, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103,
	4, 2, 85, 85, 117, 117, 4, 2, 69, 69, 101, 101, 4, 2, 67, 67, 99, 99, 4,
	2, 62, 62, 64, 64, 4, 2, 81, 81, 113, 113, 4, 2, 84, 84, 116, 116, 4, 2,
	68, 68, 100, 100, 4, 2, 91, 91, 123, 123, 3, 2, 50, 59, 5, 2, 11, 12, 15,
	15, 34, 34, 2, 385, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2,
	2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2,
	2, 2, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7, 75, 3, 2,
	2, 2, 9, 77, 3, 2, 2, 2, 11, 79, 3, 2, 2, 2, 13, 84, 3, 2, 2, 2, 15, 86,
	3, 2, 2, 2, 17, 88, 3, 2, 2, 2, 19, 90, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2,
	23, 98, 3, 2, 2, 2, 25, 102, 3, 2, 2, 2, 27, 104, 3, 2, 2, 2, 29, 124,
	3, 2, 2, 2, 31, 136, 3, 2, 2, 2, 33, 149, 3, 2, 2, 2, 35, 166, 3, 2, 2,
	2, 37, 198, 3, 2, 2, 2, 39, 200, 3, 2, 2, 2, 41, 209, 3, 2, 2, 2, 43, 212,
	3, 2, 2, 2, 45, 245, 3, 2, 2, 2, 47, 258, 3, 2, 2, 2, 49, 282, 3, 2, 2,
	2, 51, 306, 3, 2, 2, 2, 53, 337, 3, 2, 2, 2, 55, 339, 3, 2, 2, 2, 57, 353,
	3, 2, 2, 2, 59, 357, 3, 2, 2, 2, 61, 359, 3, 2, 2, 2, 63, 361, 3, 2, 2,
	2, 65, 363, 3, 2, 2, 2, 67, 365, 3, 2, 2, 2, 69, 368, 3, 2, 2, 2, 71, 72,
	9, 2, 2, 2, 72, 4, 3, 2, 2, 2, 73, 74, 9, 3, 2, 2, 74, 6, 3, 2, 2, 2, 75,
	76, 9, 4, 2, 2, 76, 8, 3, 2, 2, 2, 77, 78, 9, 5, 2, 2, 78, 10, 3, 2, 2,
	2, 79, 80, 9, 6, 2, 2, 80, 81, 9, 6, 2, 2, 81, 82, 9, 6, 2, 2, 82, 83,
	9, 6, 2, 2, 83, 12, 3, 2, 2, 2, 84, 85, 7, 34, 2, 2, 85, 14, 3, 2, 2, 2,
	86, 87, 7, 97, 2, 2, 87, 16, 3, 2, 2, 2, 88, 89, 7, 49, 2, 2, 89, 18, 3,
	2, 2, 2, 90, 91, 7, 47, 2, 2, 91, 92, 7, 64, 2, 2, 92, 20, 3, 2, 2, 2,
	93, 94, 9, 7, 2, 2, 94, 95, 9, 8, 2, 2, 95, 96, 9, 9, 2, 2, 96, 97, 9,
	10, 2, 2, 97, 22, 3, 2, 2, 2, 98, 99, 9, 11, 2, 2, 99, 100, 9, 9, 2, 2,
	100, 101, 9, 10, 2, 2, 101, 24, 3, 2, 2, 2, 102, 103, 7, 35, 2, 2, 103,
	26, 3, 2, 2, 2, 104, 105, 7, 99, 2, 2, 105, 106, 7, 112, 2, 2, 106, 107,
	7, 112, 2, 2, 107, 108, 7, 113, 2, 2, 108, 109, 7, 118, 2, 2, 109, 110,
	7, 99, 2, 2, 110, 111, 7, 118, 2, 2, 111, 112, 7, 107, 2, 2, 112, 113,
	7, 113, 2, 2, 113, 114, 7, 112, 2, 2, 114, 115, 7, 48, 2, 2, 115, 116,
	7, 101, 2, 2, 116, 117, 7, 113, 2, 2, 117, 118, 7, 120, 2, 2, 118, 119,
	7, 103, 2, 2, 119, 120, 7, 116, 2, 2, 120, 121, 7, 99, 2, 2, 121, 122,
	7, 105, 2, 2, 122, 123, 7, 103, 2, 2, 123, 28, 3, 2, 2, 2, 124, 125, 7,
	107, 2, 2, 125, 126, 7, 111, 2, 2, 126, 127, 7, 99, 2, 2, 127, 128, 7,
	105, 2, 2, 128, 129, 7, 103, 2, 2, 129, 130, 7, 48, 2, 2, 130, 131, 7,
	121, 2, 2, 131, 132, 7, 107, 2, 2, 132, 133, 7, 102, 2, 2, 133, 134, 7,
	118, 2, 2, 134, 135, 7, 106, 2, 2, 135, 30, 3, 2, 2, 2, 136, 137, 7, 107,
	2, 2, 137, 138, 7, 111, 2, 2, 138, 139, 7, 99, 2, 2, 139, 140, 7, 105,
	2, 2, 140, 141, 7, 103, 2, 2, 141, 142, 7, 48, 2, 2, 142, 143, 7, 106,
	2, 2, 143, 144, 7, 103, 2, 2, 144, 145, 7, 107, 2, 2, 145, 146, 7, 105,
	2, 2, 146, 147, 7, 106, 2, 2, 147, 148, 7, 118, 2, 2, 148, 32, 3, 2, 2,
	2, 149, 150, 7, 107, 2, 2, 150, 151, 7, 111, 2, 2, 151, 152, 7, 99, 2,
	2, 152, 153, 7, 105, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 48, 2,
	2, 155, 156, 7, 112, 2, 2, 156, 157, 7, 119, 2, 2, 157, 158, 7, 111, 2,
	2, 158, 159, 7, 97, 2, 2, 159, 160, 7, 110, 2, 2, 160, 161, 7, 99, 2, 2,
	161, 162, 7, 100, 2, 2, 162, 163, 7, 103, 2, 2, 163, 164, 7, 110, 2, 2,
	164, 165, 7, 117, 2, 2, 165, 34, 3, 2, 2, 2, 166, 167, 7, 107, 2, 2, 167,
	168, 7, 111, 2, 2, 168, 169, 7, 99, 2, 2, 169, 170, 7, 105, 2, 2, 170,
	171, 7, 103, 2, 2, 171, 172, 7, 48, 2, 2, 172, 173, 7, 112, 2, 2, 173,
	174, 7, 119, 2, 2, 174, 175, 7, 111, 2, 2, 175, 176, 7, 97, 2, 2, 176,
	177, 7, 113, 2, 2, 177, 178, 7, 114, 2, 2, 178, 179, 7, 103, 2, 2, 179,
	180, 7, 112, 2, 2, 180, 181, 7, 97, 2, 2, 181, 182, 7, 99, 2, 2, 182, 183,
	7, 112, 2, 2, 183, 184, 7, 112, 2, 2, 184, 185, 7, 113, 2, 2, 185, 186,
	7, 118, 2, 2, 186, 187, 7, 99, 2, 2, 187, 188, 7, 118, 2, 2, 188, 189,
	7, 107, 2, 2, 189, 190, 7, 113, 2, 2, 190, 191, 7, 112, 2, 2, 191, 192,
	7, 97, 2, 2, 192, 193, 7, 118, 2, 2, 193, 194, 7, 99, 2, 2, 194, 195, 7,
	117, 2, 2, 195, 196, 7, 109, 2, 2, 196, 197, 7, 117, 2, 2, 197, 36, 3,
	2, 2, 2, 198, 199, 7, 39, 2, 2, 199, 38, 3, 2, 2, 2, 200, 201, 7, 114,
	2, 2, 201, 202, 7, 122, 2, 2, 202, 40, 3, 2, 2, 2, 203, 210, 9, 12, 2,
	2, 204, 205, 7, 64, 2, 2, 205, 210, 7, 63, 2, 2, 206, 210, 7, 63, 2, 2,
	207, 208, 7, 62, 2, 2, 208, 210, 7, 63, 2, 2, 209, 203, 3, 2, 2, 2, 209,
	204, 3, 2, 2, 2, 209, 206, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 42, 3,
	2, 2, 2, 211, 213, 5, 7, 4, 2, 212, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2,
	2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 222, 3, 2, 2, 2, 216,
	218, 7, 48, 2, 2, 217, 219, 5, 7, 4, 2, 218, 217, 3, 2, 2, 2, 219, 220,
	3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 223, 3, 2,
	2, 2, 222, 216, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 227, 3, 2, 2, 2,
	224, 226, 5, 13, 7, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227,
	225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229, 227,
	3, 2, 2, 2, 230, 234, 7, 63, 2, 2, 231, 233, 5, 13, 7, 2, 232, 231, 3,
	2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2,
	2, 235, 237, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 239, 7, 41, 2, 2, 238,
	240, 5, 9, 5, 2, 239, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 239,
	3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 244, 7, 41,
	2, 2, 244, 44, 3, 2, 2, 2, 245, 246, 9, 13, 2, 2, 246, 247, 9, 14, 2, 2,
	247, 248, 9, 7, 2, 2, 248, 249, 9, 8, 2, 2, 249, 251, 9, 14, 2, 2, 250,
	252, 5, 13, 7, 2, 251, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 251,
	3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 9, 15,
	2, 2, 256, 257, 9, 16, 2, 2, 257, 46, 3, 2, 2, 2, 258, 260, 5, 45, 23,
	2, 259, 261, 5, 13, 7, 2, 260, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262,
	260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265,
	7, 120, 2, 2, 265, 266, 7, 99, 2, 2, 266, 267, 7, 110, 2, 2, 267, 268,
	7, 107, 2, 2, 268, 269, 7, 102, 2, 2, 269, 270, 7, 99, 2, 2, 270, 271,
	7, 118, 2, 2, 271, 272, 7, 107, 2, 2, 272, 273, 7, 113, 2, 2, 273, 274,
	7, 112, 2, 2, 274, 276, 3, 2, 2, 2, 275, 277, 5, 13, 7, 2, 276, 275, 3,
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2,
	2, 279, 280, 3, 2, 2, 2, 280, 281, 5, 21, 11, 2, 281, 48, 3, 2, 2, 2, 282,
	284, 5, 45, 23, 2, 283, 285, 5, 13, 7, 2, 284, 283, 3, 2, 2, 2, 285, 286,
	3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 3, 2,
	2, 2, 288, 289, 7, 120, 2, 2, 289, 290, 7, 99, 2, 2, 290, 291, 7, 110,
	2, 2, 291, 292, 7, 107, 2, 2, 292, 293, 7, 102, 2, 2, 293, 294, 7, 99,
	2, 2, 294, 295, 7, 118, 2, 2, 295, 296, 7, 107, 2, 2, 296, 297, 7, 113,
	2, 2, 297, 298, 7, 112, 2, 2, 298, 300, 3, 2, 2, 2, 299, 301, 5, 13, 7,
	2, 300, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302,
	303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305, 5, 23, 12, 2, 305, 50,
	3, 2, 2, 2, 306, 308, 5, 45, 23, 2, 307, 309, 5, 13, 7, 2, 308, 307, 3,
	2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2,
	2, 311, 312, 3, 2, 2, 2, 312, 313, 7, 120, 2, 2, 313, 314, 7, 99, 2, 2,
	314, 315, 7, 110, 2, 2, 315, 316, 7, 107, 2, 2, 316, 317, 7, 102, 2, 2,
	317, 318, 7, 99, 2, 2, 318, 319, 7, 118, 2, 2, 319, 320, 7, 107, 2, 2,
	320, 321, 7, 113, 2, 2, 321, 322, 7, 112, 2, 2, 322, 52, 3, 2, 2, 2, 323,
	338, 5, 7, 4, 2, 324, 332, 5, 7, 4, 2, 325, 331, 5, 13, 7, 2, 326, 331,
	5, 7, 4, 2, 327, 331, 5, 15, 8, 2, 328, 331, 5, 17, 9, 2, 329, 331, 5,
	19, 10, 2, 330, 325, 3, 2, 2, 2, 330, 326, 3, 2, 2, 2, 330, 327, 3, 2,
	2, 2, 330, 328, 3, 2, 2, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2,
	332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334,
	332, 3, 2, 2, 2, 335, 336, 5, 7, 4, 2, 336, 338, 3, 2, 2, 2, 337, 323,
	3, 2, 2, 2, 337, 324, 3, 2, 2, 2, 338, 54, 3, 2, 2, 2, 339, 340, 5, 11,
	6, 2, 340, 341, 5, 11, 6, 2, 341, 342, 7, 47, 2, 2, 342, 343, 5, 11, 6,
	2, 343, 344, 7, 47, 2, 2, 344, 345, 5, 11, 6, 2, 345, 346, 7, 47, 2, 2,
	346, 347, 5, 11, 6, 2, 347, 348, 7, 47, 2, 2, 348, 349, 5, 11, 6, 2, 349,
	350, 5, 11, 6, 2, 350, 351, 5, 11, 6, 2, 351, 56, 3, 2, 2, 2, 352, 354,
	9, 17, 2, 2, 353, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 353, 3, 2,
	2, 2, 355, 356, 3, 2, 2, 2, 356, 58, 3, 2, 2, 2, 357, 358, 7, 40, 2, 2,
	358, 60, 3, 2, 2, 2, 359, 360, 7, 126, 2, 2, 360, 62, 3, 2, 2, 2, 361,
	362, 7, 128, 2, 2, 362, 64, 3, 2, 2, 2, 363, 364, 7, 42, 2, 2, 364, 66,
	3, 2, 2, 2, 365, 366, 7, 43, 2, 2, 366, 68, 3, 2, 2, 2, 367, 369, 9, 18,
	2, 2, 368, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2,
	370, 371, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 8, 35, 2, 2, 373,
	70, 3, 2, 2, 2, 21, 2, 209, 214, 220, 222, 227, 234, 241, 253, 262, 278,
	286, 302, 310, 330, 332, 337, 355, 370, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'!'", "'annotation.coverage'", "'image.width'", "'image.height'",
	"'image.num_labels'", "'image.num_open_annotation_tasks'", "'%'", "'px'",
	"", "", "", "", "", "", "", "", "", "'&'", "'|'", "'~'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "SEP", "ANNOTATION_COVERAGE_PREFIX", "IMAGE_WIDTH_PREFIX", "IMAGE_HEIGHT_PREFIX",
	"IMAGE_NUM_LABELS_PREFIX", "IMAGE_NUM_OPEN_ANNOTATION_TASKS_PREFIX", "PERCENT",
	"PIXEL", "OPERATOR", "ASSIGNMENT", "ORDER_BY", "ORDER_BY_VALIDATION_DESC",
	"ORDER_BY_VALIDATION_ASC", "ORDER_BY_VALIDATION", "LABEL", "UUID", "VAL",
	"AND", "OR", "NOT", "LPAR", "RPAR", "SKIPPED_TOKENS",
}

var lexerRuleNames = []string{
	"LOWERCASE", "UPPERCASE", "UPPERLOWERCASE", "UPPERLOWERCASEWS", "UUIDBLOCK",
	"WS", "UNDERSCORE", "SLASH", "GRAPHARROW", "DESC", "ASC", "SEP", "ANNOTATION_COVERAGE_PREFIX",
	"IMAGE_WIDTH_PREFIX", "IMAGE_HEIGHT_PREFIX", "IMAGE_NUM_LABELS_PREFIX",
	"IMAGE_NUM_OPEN_ANNOTATION_TASKS_PREFIX", "PERCENT", "PIXEL", "OPERATOR",
	"ASSIGNMENT", "ORDER_BY", "ORDER_BY_VALIDATION_DESC", "ORDER_BY_VALIDATION_ASC",
	"ORDER_BY_VALIDATION", "LABEL", "UUID", "VAL", "AND", "OR", "NOT", "LPAR",
	"RPAR", "SKIPPED_TOKENS",
}

type ImagemonkeyQueryLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewImagemonkeyQueryLangLexer(input antlr.CharStream) *ImagemonkeyQueryLangLexer {

	l := new(ImagemonkeyQueryLangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ImagemonkeyQueryLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ImagemonkeyQueryLangLexer tokens.
const (
	ImagemonkeyQueryLangLexerSEP                                    = 1
	ImagemonkeyQueryLangLexerANNOTATION_COVERAGE_PREFIX             = 2
	ImagemonkeyQueryLangLexerIMAGE_WIDTH_PREFIX                     = 3
	ImagemonkeyQueryLangLexerIMAGE_HEIGHT_PREFIX                    = 4
	ImagemonkeyQueryLangLexerIMAGE_NUM_LABELS_PREFIX                = 5
	ImagemonkeyQueryLangLexerIMAGE_NUM_OPEN_ANNOTATION_TASKS_PREFIX = 6
	ImagemonkeyQueryLangLexerPERCENT                                = 7
	ImagemonkeyQueryLangLexerPIXEL                                  = 8
	ImagemonkeyQueryLangLexerOPERATOR                               = 9
	ImagemonkeyQueryLangLexerASSIGNMENT                             = 10
	ImagemonkeyQueryLangLexerORDER_BY                               = 11
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION_DESC               = 12
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION_ASC                = 13
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION                    = 14
	ImagemonkeyQueryLangLexerLABEL                                  = 15
	ImagemonkeyQueryLangLexerUUID                                   = 16
	ImagemonkeyQueryLangLexerVAL                                    = 17
	ImagemonkeyQueryLangLexerAND                                    = 18
	ImagemonkeyQueryLangLexerOR                                     = 19
	ImagemonkeyQueryLangLexerNOT                                    = 20
	ImagemonkeyQueryLangLexerLPAR                                   = 21
	ImagemonkeyQueryLangLexerRPAR                                   = 22
	ImagemonkeyQueryLangLexerSKIPPED_TOKENS                         = 23
)
