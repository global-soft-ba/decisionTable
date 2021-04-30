// Code generated from SFeel.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // SFeel

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 366,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 100, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 107, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 127, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 134,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	146, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 153, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 160, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 180, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 187, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 193, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	200, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 207, 10, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 214, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 221, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 229, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 243, 10, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5,
	24, 258, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 274, 10, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 3, 34, 5, 34, 301, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35,
	307, 10, 35, 3, 35, 3, 35, 3, 35, 7, 35, 312, 10, 35, 12, 35, 14, 35, 315,
	11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 321, 10, 36, 3, 36, 3, 36, 3,
	36, 7, 36, 326, 10, 36, 12, 36, 14, 36, 329, 11, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 7, 37, 337, 10, 37, 12, 37, 14, 37, 340, 11, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 5, 38, 346, 10, 38, 3, 38, 3, 38, 3, 38, 7, 38,
	351, 10, 38, 12, 38, 14, 38, 354, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 5, 39, 362, 10, 39, 3, 39, 3, 39, 3, 39, 2, 6, 68, 70, 72, 74,
	40, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 2, 5, 3, 2, 8, 9, 3, 2, 12, 15, 3, 2, 16, 17, 2, 389, 2, 78, 3,
	2, 2, 2, 4, 106, 3, 2, 2, 2, 6, 133, 3, 2, 2, 2, 8, 152, 3, 2, 2, 2, 10,
	159, 3, 2, 2, 2, 12, 186, 3, 2, 2, 2, 14, 192, 3, 2, 2, 2, 16, 199, 3,
	2, 2, 2, 18, 206, 3, 2, 2, 2, 20, 213, 3, 2, 2, 2, 22, 220, 3, 2, 2, 2,
	24, 228, 3, 2, 2, 2, 26, 230, 3, 2, 2, 2, 28, 232, 3, 2, 2, 2, 30, 234,
	3, 2, 2, 2, 32, 236, 3, 2, 2, 2, 34, 242, 3, 2, 2, 2, 36, 244, 3, 2, 2,
	2, 38, 246, 3, 2, 2, 2, 40, 248, 3, 2, 2, 2, 42, 250, 3, 2, 2, 2, 44, 252,
	3, 2, 2, 2, 46, 257, 3, 2, 2, 2, 48, 259, 3, 2, 2, 2, 50, 261, 3, 2, 2,
	2, 52, 264, 3, 2, 2, 2, 54, 267, 3, 2, 2, 2, 56, 273, 3, 2, 2, 2, 58, 275,
	3, 2, 2, 2, 60, 277, 3, 2, 2, 2, 62, 283, 3, 2, 2, 2, 64, 289, 3, 2, 2,
	2, 66, 300, 3, 2, 2, 2, 68, 302, 3, 2, 2, 2, 70, 316, 3, 2, 2, 2, 72, 330,
	3, 2, 2, 2, 74, 341, 3, 2, 2, 2, 76, 355, 3, 2, 2, 2, 78, 79, 5, 24, 13,
	2, 79, 80, 7, 2, 2, 3, 80, 3, 3, 2, 2, 2, 81, 82, 5, 36, 19, 2, 82, 83,
	7, 2, 2, 3, 83, 107, 3, 2, 2, 2, 84, 85, 5, 50, 26, 2, 85, 86, 7, 2, 2,
	3, 86, 107, 3, 2, 2, 2, 87, 88, 5, 60, 31, 2, 88, 89, 7, 2, 2, 3, 89, 107,
	3, 2, 2, 2, 90, 91, 5, 68, 35, 2, 91, 92, 7, 2, 2, 3, 92, 107, 3, 2, 2,
	2, 93, 94, 7, 19, 2, 2, 94, 99, 7, 3, 2, 2, 95, 100, 5, 36, 19, 2, 96,
	100, 5, 50, 26, 2, 97, 100, 5, 60, 31, 2, 98, 100, 5, 68, 35, 2, 99, 95,
	3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2,
	100, 101, 3, 2, 2, 2, 101, 102, 7, 4, 2, 2, 102, 103, 7, 2, 2, 3, 103,
	107, 3, 2, 2, 2, 104, 105, 7, 5, 2, 2, 105, 107, 7, 2, 2, 3, 106, 81, 3,
	2, 2, 2, 106, 84, 3, 2, 2, 2, 106, 87, 3, 2, 2, 2, 106, 90, 3, 2, 2, 2,
	106, 93, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 5, 3, 2, 2, 2, 108, 109,
	5, 38, 20, 2, 109, 110, 7, 2, 2, 3, 110, 134, 3, 2, 2, 2, 111, 112, 5,
	54, 28, 2, 112, 113, 7, 2, 2, 3, 113, 134, 3, 2, 2, 2, 114, 115, 5, 62,
	32, 2, 115, 116, 7, 2, 2, 3, 116, 134, 3, 2, 2, 2, 117, 118, 5, 70, 36,
	2, 118, 119, 7, 2, 2, 3, 119, 134, 3, 2, 2, 2, 120, 121, 7, 19, 2, 2, 121,
	126, 7, 3, 2, 2, 122, 127, 5, 38, 20, 2, 123, 127, 5, 54, 28, 2, 124, 127,
	5, 62, 32, 2, 125, 127, 5, 70, 36, 2, 126, 122, 3, 2, 2, 2, 126, 123, 3,
	2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2,
	2, 128, 129, 7, 4, 2, 2, 129, 130, 7, 2, 2, 3, 130, 134, 3, 2, 2, 2, 131,
	132, 7, 5, 2, 2, 132, 134, 7, 2, 2, 3, 133, 108, 3, 2, 2, 2, 133, 111,
	3, 2, 2, 2, 133, 114, 3, 2, 2, 2, 133, 117, 3, 2, 2, 2, 133, 120, 3, 2,
	2, 2, 133, 131, 3, 2, 2, 2, 134, 7, 3, 2, 2, 2, 135, 136, 5, 42, 22, 2,
	136, 137, 7, 2, 2, 3, 137, 153, 3, 2, 2, 2, 138, 139, 5, 72, 37, 2, 139,
	140, 7, 2, 2, 3, 140, 153, 3, 2, 2, 2, 141, 142, 7, 19, 2, 2, 142, 145,
	7, 3, 2, 2, 143, 146, 5, 42, 22, 2, 144, 146, 5, 72, 37, 2, 145, 143, 3,
	2, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 7, 4, 2,
	2, 148, 149, 7, 2, 2, 3, 149, 153, 3, 2, 2, 2, 150, 151, 7, 5, 2, 2, 151,
	153, 7, 2, 2, 3, 152, 135, 3, 2, 2, 2, 152, 138, 3, 2, 2, 2, 152, 141,
	3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 9, 3, 2, 2, 2, 154, 155, 5, 40,
	21, 2, 155, 156, 7, 2, 2, 3, 156, 160, 3, 2, 2, 2, 157, 158, 7, 5, 2, 2,
	158, 160, 7, 2, 2, 3, 159, 154, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160,
	11, 3, 2, 2, 2, 161, 162, 5, 44, 23, 2, 162, 163, 7, 2, 2, 3, 163, 187,
	3, 2, 2, 2, 164, 165, 5, 52, 27, 2, 165, 166, 7, 2, 2, 3, 166, 187, 3,
	2, 2, 2, 167, 168, 5, 64, 33, 2, 168, 169, 7, 2, 2, 3, 169, 187, 3, 2,
	2, 2, 170, 171, 5, 74, 38, 2, 171, 172, 7, 2, 2, 3, 172, 187, 3, 2, 2,
	2, 173, 174, 7, 19, 2, 2, 174, 179, 7, 3, 2, 2, 175, 180, 5, 44, 23, 2,
	176, 180, 5, 52, 27, 2, 177, 180, 5, 64, 33, 2, 178, 180, 5, 74, 38, 2,
	179, 175, 3, 2, 2, 2, 179, 176, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179,
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 4, 2, 2, 182, 183,
	7, 2, 2, 3, 183, 187, 3, 2, 2, 2, 184, 185, 7, 5, 2, 2, 185, 187, 7, 2,
	2, 3, 186, 161, 3, 2, 2, 2, 186, 164, 3, 2, 2, 2, 186, 167, 3, 2, 2, 2,
	186, 170, 3, 2, 2, 2, 186, 173, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187,
	13, 3, 2, 2, 2, 188, 189, 7, 8, 2, 2, 189, 193, 7, 2, 2, 3, 190, 191, 7,
	5, 2, 2, 191, 193, 7, 2, 2, 3, 192, 188, 3, 2, 2, 2, 192, 190, 3, 2, 2,
	2, 193, 15, 3, 2, 2, 2, 194, 195, 5, 26, 14, 2, 195, 196, 7, 2, 2, 3, 196,
	200, 3, 2, 2, 2, 197, 198, 7, 5, 2, 2, 198, 200, 7, 2, 2, 3, 199, 194,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 17, 3, 2, 2, 2, 201, 202, 5, 28,
	15, 2, 202, 203, 7, 2, 2, 3, 203, 207, 3, 2, 2, 2, 204, 205, 7, 5, 2, 2,
	205, 207, 7, 2, 2, 3, 206, 201, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207,
	19, 3, 2, 2, 2, 208, 209, 5, 30, 16, 2, 209, 210, 7, 2, 2, 3, 210, 214,
	3, 2, 2, 2, 211, 212, 7, 5, 2, 2, 212, 214, 7, 2, 2, 3, 213, 208, 3, 2,
	2, 2, 213, 211, 3, 2, 2, 2, 214, 21, 3, 2, 2, 2, 215, 216, 5, 32, 17, 2,
	216, 217, 7, 2, 2, 3, 217, 221, 3, 2, 2, 2, 218, 219, 7, 5, 2, 2, 219,
	221, 7, 2, 2, 3, 220, 215, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 23, 3,
	2, 2, 2, 222, 229, 5, 34, 18, 2, 223, 229, 5, 46, 24, 2, 224, 229, 5, 56,
	29, 2, 225, 229, 5, 66, 34, 2, 226, 229, 5, 76, 39, 2, 227, 229, 7, 5,
	2, 2, 228, 222, 3, 2, 2, 2, 228, 223, 3, 2, 2, 2, 228, 224, 3, 2, 2, 2,
	228, 225, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229,
	25, 3, 2, 2, 2, 230, 231, 9, 2, 2, 2, 231, 27, 3, 2, 2, 2, 232, 233, 7,
	10, 2, 2, 233, 29, 3, 2, 2, 2, 234, 235, 7, 11, 2, 2, 235, 31, 3, 2, 2,
	2, 236, 237, 7, 20, 2, 2, 237, 33, 3, 2, 2, 2, 238, 243, 5, 38, 20, 2,
	239, 243, 5, 42, 22, 2, 240, 243, 5, 44, 23, 2, 241, 243, 5, 40, 21, 2,
	242, 238, 3, 2, 2, 2, 242, 239, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242,
	241, 3, 2, 2, 2, 243, 35, 3, 2, 2, 2, 244, 245, 7, 8, 2, 2, 245, 37, 3,
	2, 2, 2, 246, 247, 5, 26, 14, 2, 247, 39, 3, 2, 2, 2, 248, 249, 5, 30,
	16, 2, 249, 41, 3, 2, 2, 2, 250, 251, 5, 28, 15, 2, 251, 43, 3, 2, 2, 2,
	252, 253, 5, 32, 17, 2, 253, 45, 3, 2, 2, 2, 254, 258, 5, 52, 27, 2, 255,
	258, 5, 50, 26, 2, 256, 258, 5, 54, 28, 2, 257, 254, 3, 2, 2, 2, 257, 255,
	3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 47, 3, 2, 2, 2, 259, 260, 9, 3,
	2, 2, 260, 49, 3, 2, 2, 2, 261, 262, 5, 48, 25, 2, 262, 263, 7, 8, 2, 2,
	263, 51, 3, 2, 2, 2, 264, 265, 5, 48, 25, 2, 265, 266, 5, 32, 17, 2, 266,
	53, 3, 2, 2, 2, 267, 268, 5, 48, 25, 2, 268, 269, 5, 26, 14, 2, 269, 55,
	3, 2, 2, 2, 270, 274, 5, 60, 31, 2, 271, 274, 5, 64, 33, 2, 272, 274, 5,
	62, 32, 2, 273, 270, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 273, 272, 3, 2,
	2, 2, 274, 57, 3, 2, 2, 2, 275, 276, 9, 4, 2, 2, 276, 59, 3, 2, 2, 2, 277,
	278, 5, 58, 30, 2, 278, 279, 7, 8, 2, 2, 279, 280, 7, 6, 2, 2, 280, 281,
	7, 8, 2, 2, 281, 282, 5, 58, 30, 2, 282, 61, 3, 2, 2, 2, 283, 284, 5, 58,
	30, 2, 284, 285, 5, 26, 14, 2, 285, 286, 7, 6, 2, 2, 286, 287, 5, 26, 14,
	2, 287, 288, 5, 58, 30, 2, 288, 63, 3, 2, 2, 2, 289, 290, 5, 58, 30, 2,
	290, 291, 5, 32, 17, 2, 291, 292, 7, 6, 2, 2, 292, 293, 5, 32, 17, 2, 293,
	294, 3, 2, 2, 2, 294, 295, 5, 58, 30, 2, 295, 65, 3, 2, 2, 2, 296, 301,
	5, 68, 35, 2, 297, 301, 5, 70, 36, 2, 298, 301, 5, 72, 37, 2, 299, 301,
	5, 74, 38, 2, 300, 296, 3, 2, 2, 2, 300, 297, 3, 2, 2, 2, 300, 298, 3,
	2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 67, 3, 2, 2, 2, 302, 306, 8, 35, 1,
	2, 303, 307, 7, 8, 2, 2, 304, 307, 5, 50, 26, 2, 305, 307, 5, 60, 31, 2,
	306, 303, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307,
	313, 3, 2, 2, 2, 308, 309, 12, 4, 2, 2, 309, 310, 7, 18, 2, 2, 310, 312,
	5, 68, 35, 5, 311, 308, 3, 2, 2, 2, 312, 315, 3, 2, 2, 2, 313, 311, 3,
	2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 69, 3, 2, 2, 2, 315, 313, 3, 2, 2,
	2, 316, 320, 8, 36, 1, 2, 317, 321, 5, 26, 14, 2, 318, 321, 5, 54, 28,
	2, 319, 321, 5, 62, 32, 2, 320, 317, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2,
	320, 319, 3, 2, 2, 2, 321, 327, 3, 2, 2, 2, 322, 323, 12, 4, 2, 2, 323,
	324, 7, 18, 2, 2, 324, 326, 5, 70, 36, 5, 325, 322, 3, 2, 2, 2, 326, 329,
	3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 71, 3, 2,
	2, 2, 329, 327, 3, 2, 2, 2, 330, 331, 8, 37, 1, 2, 331, 332, 5, 28, 15,
	2, 332, 338, 3, 2, 2, 2, 333, 334, 12, 4, 2, 2, 334, 335, 7, 18, 2, 2,
	335, 337, 5, 72, 37, 5, 336, 333, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2, 338,
	336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 73, 3, 2, 2, 2, 340, 338, 3,
	2, 2, 2, 341, 345, 8, 38, 1, 2, 342, 346, 5, 32, 17, 2, 343, 346, 5, 52,
	27, 2, 344, 346, 5, 64, 33, 2, 345, 342, 3, 2, 2, 2, 345, 343, 3, 2, 2,
	2, 345, 344, 3, 2, 2, 2, 346, 352, 3, 2, 2, 2, 347, 348, 12, 4, 2, 2, 348,
	349, 7, 18, 2, 2, 349, 351, 5, 74, 38, 5, 350, 347, 3, 2, 2, 2, 351, 354,
	3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 75, 3, 2,
	2, 2, 354, 352, 3, 2, 2, 2, 355, 356, 7, 19, 2, 2, 356, 361, 7, 3, 2, 2,
	357, 362, 5, 34, 18, 2, 358, 362, 5, 46, 24, 2, 359, 362, 5, 56, 29, 2,
	360, 362, 5, 66, 34, 2, 361, 357, 3, 2, 2, 2, 361, 358, 3, 2, 2, 2, 361,
	359, 3, 2, 2, 2, 361, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364,
	7, 4, 2, 2, 364, 77, 3, 2, 2, 2, 29, 99, 106, 126, 133, 145, 152, 159,
	179, 186, 192, 199, 206, 213, 220, 228, 242, 257, 273, 300, 306, 313, 320,
	327, 338, 345, 352, 361,
}
var literalNames = []string{
	"", "'('", "')'", "'-'", "'..'", "", "", "", "", "", "'<'", "'<='", "'>'",
	"'>='", "'['", "']'", "','", "'not'",
}
var symbolicNames = []string{
	"", "", "", "", "", "SIGN", "INTEGER", "FLOAT", "STRING", "BOOL", "LESS",
	"LESSEQ", "GREATER", "GREATEREQ", "RANGEIN", "RANGEOUT", "DISJUNCTION",
	"NEGATION", "DATEANDTIME", "FORMAT",
}

var ruleNames = []string{
	"entry", "validIntegerInput", "validNumberInput", "validStringInput", "validBoolInput",
	"validDateTimeInput", "validIntegerOutput", "validNumberOutput", "validStringOutput",
	"validBoolOutput", "validDateTimeOutput", "expression", "number", "strings",
	"bools", "datetime", "equalcomparison", "equalcomparisonInteger", "equalcomparisonNumber",
	"equalcomparisonBool", "equalcomparisonStrings", "equalcomparisonDateTime",
	"comparison", "comparisonOps", "comparisonInteger", "comparisonDateTime",
	"comparisonNumber", "ranges", "rop", "rangeInteger", "rangeNumber", "rangeDateTime",
	"disjunctions", "disjunctionsInteger", "disjunctionsNumber", "disjunctionsString",
	"disjunctionsDateTime", "negation",
}

type SFeelParser struct {
	*antlr.BaseParser
}

// NewSFeelParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *SFeelParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSFeelParser(input antlr.TokenStream) *SFeelParser {
	this := new(SFeelParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SFeel.g4"

	return this
}

// SFeelParser tokens.
const (
	SFeelParserEOF         = antlr.TokenEOF
	SFeelParserT__0        = 1
	SFeelParserT__1        = 2
	SFeelParserT__2        = 3
	SFeelParserT__3        = 4
	SFeelParserSIGN        = 5
	SFeelParserINTEGER     = 6
	SFeelParserFLOAT       = 7
	SFeelParserSTRING      = 8
	SFeelParserBOOL        = 9
	SFeelParserLESS        = 10
	SFeelParserLESSEQ      = 11
	SFeelParserGREATER     = 12
	SFeelParserGREATEREQ   = 13
	SFeelParserRANGEIN     = 14
	SFeelParserRANGEOUT    = 15
	SFeelParserDISJUNCTION = 16
	SFeelParserNEGATION    = 17
	SFeelParserDATEANDTIME = 18
	SFeelParserFORMAT      = 19
)

// SFeelParser rules.
const (
	SFeelParserRULE_entry                   = 0
	SFeelParserRULE_validIntegerInput       = 1
	SFeelParserRULE_validNumberInput        = 2
	SFeelParserRULE_validStringInput        = 3
	SFeelParserRULE_validBoolInput          = 4
	SFeelParserRULE_validDateTimeInput      = 5
	SFeelParserRULE_validIntegerOutput      = 6
	SFeelParserRULE_validNumberOutput       = 7
	SFeelParserRULE_validStringOutput       = 8
	SFeelParserRULE_validBoolOutput         = 9
	SFeelParserRULE_validDateTimeOutput     = 10
	SFeelParserRULE_expression              = 11
	SFeelParserRULE_number                  = 12
	SFeelParserRULE_strings                 = 13
	SFeelParserRULE_bools                   = 14
	SFeelParserRULE_datetime                = 15
	SFeelParserRULE_equalcomparison         = 16
	SFeelParserRULE_equalcomparisonInteger  = 17
	SFeelParserRULE_equalcomparisonNumber   = 18
	SFeelParserRULE_equalcomparisonBool     = 19
	SFeelParserRULE_equalcomparisonStrings  = 20
	SFeelParserRULE_equalcomparisonDateTime = 21
	SFeelParserRULE_comparison              = 22
	SFeelParserRULE_comparisonOps           = 23
	SFeelParserRULE_comparisonInteger       = 24
	SFeelParserRULE_comparisonDateTime      = 25
	SFeelParserRULE_comparisonNumber        = 26
	SFeelParserRULE_ranges                  = 27
	SFeelParserRULE_rop                     = 28
	SFeelParserRULE_rangeInteger            = 29
	SFeelParserRULE_rangeNumber             = 30
	SFeelParserRULE_rangeDateTime           = 31
	SFeelParserRULE_disjunctions            = 32
	SFeelParserRULE_disjunctionsInteger     = 33
	SFeelParserRULE_disjunctionsNumber      = 34
	SFeelParserRULE_disjunctionsString      = 35
	SFeelParserRULE_disjunctionsDateTime    = 36
	SFeelParserRULE_negation                = 37
)

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_entry
	return p
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *SFeelParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SFeelParserRULE_entry)

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
		p.SetState(76)
		p.Expression()
	}
	{
		p.SetState(77)
		p.Match(SFeelParserEOF)
	}

	return localctx
}

// IValidIntegerInputContext is an interface to support dynamic dispatch.
type IValidIntegerInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidIntegerInputContext differentiates from other interfaces.
	IsValidIntegerInputContext()
}

type ValidIntegerInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidIntegerInputContext() *ValidIntegerInputContext {
	var p = new(ValidIntegerInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validIntegerInput
	return p
}

func (*ValidIntegerInputContext) IsValidIntegerInputContext() {}

func NewValidIntegerInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidIntegerInputContext {
	var p = new(ValidIntegerInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validIntegerInput

	return p
}

func (s *ValidIntegerInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidIntegerInputContext) EqualcomparisonInteger() IEqualcomparisonIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonIntegerContext)
}

func (s *ValidIntegerInputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidIntegerInputContext) ComparisonInteger() IComparisonIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonIntegerContext)
}

func (s *ValidIntegerInputContext) RangeInteger() IRangeIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeIntegerContext)
}

func (s *ValidIntegerInputContext) DisjunctionsInteger() IDisjunctionsIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsIntegerContext)
}

func (s *ValidIntegerInputContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *ValidIntegerInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidIntegerInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidIntegerInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidIntegerInput(s)
	}
}

func (s *ValidIntegerInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidIntegerInput(s)
	}
}

func (p *SFeelParser) ValidIntegerInput() (localctx IValidIntegerInputContext) {
	localctx = NewValidIntegerInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SFeelParserRULE_validIntegerInput)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.EqualcomparisonInteger()
		}
		{
			p.SetState(80)
			p.Match(SFeelParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.ComparisonInteger()
		}
		{
			p.SetState(83)
			p.Match(SFeelParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)
			p.RangeInteger()
		}
		{
			p.SetState(86)
			p.Match(SFeelParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.disjunctionsInteger(0)
		}
		{
			p.SetState(89)
			p.Match(SFeelParserEOF)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.Match(SFeelParserNEGATION)
		}
		{
			p.SetState(92)
			p.Match(SFeelParserT__0)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(93)
				p.EqualcomparisonInteger()
			}

		case 2:
			{
				p.SetState(94)
				p.ComparisonInteger()
			}

		case 3:
			{
				p.SetState(95)
				p.RangeInteger()
			}

		case 4:
			{
				p.SetState(96)
				p.disjunctionsInteger(0)
			}

		}
		{
			p.SetState(99)
			p.Match(SFeelParserT__1)
		}
		{
			p.SetState(100)
			p.Match(SFeelParserEOF)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(102)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(103)
			p.Match(SFeelParserEOF)
		}

	}

	return localctx
}

// IValidNumberInputContext is an interface to support dynamic dispatch.
type IValidNumberInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidNumberInputContext differentiates from other interfaces.
	IsValidNumberInputContext()
}

type ValidNumberInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidNumberInputContext() *ValidNumberInputContext {
	var p = new(ValidNumberInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validNumberInput
	return p
}

func (*ValidNumberInputContext) IsValidNumberInputContext() {}

func NewValidNumberInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidNumberInputContext {
	var p = new(ValidNumberInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validNumberInput

	return p
}

func (s *ValidNumberInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidNumberInputContext) EqualcomparisonNumber() IEqualcomparisonNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonNumberContext)
}

func (s *ValidNumberInputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidNumberInputContext) ComparisonNumber() IComparisonNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonNumberContext)
}

func (s *ValidNumberInputContext) RangeNumber() IRangeNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeNumberContext)
}

func (s *ValidNumberInputContext) DisjunctionsNumber() IDisjunctionsNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsNumberContext)
}

func (s *ValidNumberInputContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *ValidNumberInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidNumberInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidNumberInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidNumberInput(s)
	}
}

func (s *ValidNumberInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidNumberInput(s)
	}
}

func (p *SFeelParser) ValidNumberInput() (localctx IValidNumberInputContext) {
	localctx = NewValidNumberInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SFeelParserRULE_validNumberInput)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.EqualcomparisonNumber()
		}
		{
			p.SetState(107)
			p.Match(SFeelParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.ComparisonNumber()
		}
		{
			p.SetState(110)
			p.Match(SFeelParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.RangeNumber()
		}
		{
			p.SetState(113)
			p.Match(SFeelParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)
			p.disjunctionsNumber(0)
		}
		{
			p.SetState(116)
			p.Match(SFeelParserEOF)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.Match(SFeelParserNEGATION)
		}
		{
			p.SetState(119)
			p.Match(SFeelParserT__0)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(120)
				p.EqualcomparisonNumber()
			}

		case 2:
			{
				p.SetState(121)
				p.ComparisonNumber()
			}

		case 3:
			{
				p.SetState(122)
				p.RangeNumber()
			}

		case 4:
			{
				p.SetState(123)
				p.disjunctionsNumber(0)
			}

		}
		{
			p.SetState(126)
			p.Match(SFeelParserT__1)
		}
		{
			p.SetState(127)
			p.Match(SFeelParserEOF)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(130)
			p.Match(SFeelParserEOF)
		}

	}

	return localctx
}

// IValidStringInputContext is an interface to support dynamic dispatch.
type IValidStringInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidStringInputContext differentiates from other interfaces.
	IsValidStringInputContext()
}

type ValidStringInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidStringInputContext() *ValidStringInputContext {
	var p = new(ValidStringInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validStringInput
	return p
}

func (*ValidStringInputContext) IsValidStringInputContext() {}

func NewValidStringInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidStringInputContext {
	var p = new(ValidStringInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validStringInput

	return p
}

func (s *ValidStringInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidStringInputContext) EqualcomparisonStrings() IEqualcomparisonStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonStringsContext)
}

func (s *ValidStringInputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidStringInputContext) DisjunctionsString() IDisjunctionsStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsStringContext)
}

func (s *ValidStringInputContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *ValidStringInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidStringInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidStringInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidStringInput(s)
	}
}

func (s *ValidStringInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidStringInput(s)
	}
}

func (p *SFeelParser) ValidStringInput() (localctx IValidStringInputContext) {
	localctx = NewValidStringInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SFeelParserRULE_validStringInput)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.EqualcomparisonStrings()
		}
		{
			p.SetState(134)
			p.Match(SFeelParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.disjunctionsString(0)
		}
		{
			p.SetState(137)
			p.Match(SFeelParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(SFeelParserNEGATION)
		}
		{
			p.SetState(140)
			p.Match(SFeelParserT__0)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(141)
				p.EqualcomparisonStrings()
			}

		case 2:
			{
				p.SetState(142)
				p.disjunctionsString(0)
			}

		}
		{
			p.SetState(145)
			p.Match(SFeelParserT__1)
		}
		{
			p.SetState(146)
			p.Match(SFeelParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(149)
			p.Match(SFeelParserEOF)
		}

	}

	return localctx
}

// IValidBoolInputContext is an interface to support dynamic dispatch.
type IValidBoolInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidBoolInputContext differentiates from other interfaces.
	IsValidBoolInputContext()
}

type ValidBoolInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidBoolInputContext() *ValidBoolInputContext {
	var p = new(ValidBoolInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validBoolInput
	return p
}

func (*ValidBoolInputContext) IsValidBoolInputContext() {}

func NewValidBoolInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidBoolInputContext {
	var p = new(ValidBoolInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validBoolInput

	return p
}

func (s *ValidBoolInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidBoolInputContext) EqualcomparisonBool() IEqualcomparisonBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonBoolContext)
}

func (s *ValidBoolInputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidBoolInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidBoolInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidBoolInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidBoolInput(s)
	}
}

func (s *ValidBoolInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidBoolInput(s)
	}
}

func (p *SFeelParser) ValidBoolInput() (localctx IValidBoolInputContext) {
	localctx = NewValidBoolInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SFeelParserRULE_validBoolInput)

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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.EqualcomparisonBool()
		}
		{
			p.SetState(153)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(156)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidDateTimeInputContext is an interface to support dynamic dispatch.
type IValidDateTimeInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidDateTimeInputContext differentiates from other interfaces.
	IsValidDateTimeInputContext()
}

type ValidDateTimeInputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidDateTimeInputContext() *ValidDateTimeInputContext {
	var p = new(ValidDateTimeInputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validDateTimeInput
	return p
}

func (*ValidDateTimeInputContext) IsValidDateTimeInputContext() {}

func NewValidDateTimeInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidDateTimeInputContext {
	var p = new(ValidDateTimeInputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validDateTimeInput

	return p
}

func (s *ValidDateTimeInputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidDateTimeInputContext) EqualcomparisonDateTime() IEqualcomparisonDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonDateTimeContext)
}

func (s *ValidDateTimeInputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidDateTimeInputContext) ComparisonDateTime() IComparisonDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonDateTimeContext)
}

func (s *ValidDateTimeInputContext) RangeDateTime() IRangeDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDateTimeContext)
}

func (s *ValidDateTimeInputContext) DisjunctionsDateTime() IDisjunctionsDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsDateTimeContext)
}

func (s *ValidDateTimeInputContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *ValidDateTimeInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidDateTimeInputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidDateTimeInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidDateTimeInput(s)
	}
}

func (s *ValidDateTimeInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidDateTimeInput(s)
	}
}

func (p *SFeelParser) ValidDateTimeInput() (localctx IValidDateTimeInputContext) {
	localctx = NewValidDateTimeInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SFeelParserRULE_validDateTimeInput)

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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.EqualcomparisonDateTime()
		}
		{
			p.SetState(160)
			p.Match(SFeelParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.ComparisonDateTime()
		}
		{
			p.SetState(163)
			p.Match(SFeelParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.RangeDateTime()
		}
		{
			p.SetState(166)
			p.Match(SFeelParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.disjunctionsDateTime(0)
		}
		{
			p.SetState(169)
			p.Match(SFeelParserEOF)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(171)
			p.Match(SFeelParserNEGATION)
		}
		{
			p.SetState(172)
			p.Match(SFeelParserT__0)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(173)
				p.EqualcomparisonDateTime()
			}

		case 2:
			{
				p.SetState(174)
				p.ComparisonDateTime()
			}

		case 3:
			{
				p.SetState(175)
				p.RangeDateTime()
			}

		case 4:
			{
				p.SetState(176)
				p.disjunctionsDateTime(0)
			}

		}
		{
			p.SetState(179)
			p.Match(SFeelParserT__1)
		}
		{
			p.SetState(180)
			p.Match(SFeelParserEOF)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(182)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(183)
			p.Match(SFeelParserEOF)
		}

	}

	return localctx
}

// IValidIntegerOutputContext is an interface to support dynamic dispatch.
type IValidIntegerOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidIntegerOutputContext differentiates from other interfaces.
	IsValidIntegerOutputContext()
}

type ValidIntegerOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidIntegerOutputContext() *ValidIntegerOutputContext {
	var p = new(ValidIntegerOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validIntegerOutput
	return p
}

func (*ValidIntegerOutputContext) IsValidIntegerOutputContext() {}

func NewValidIntegerOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidIntegerOutputContext {
	var p = new(ValidIntegerOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validIntegerOutput

	return p
}

func (s *ValidIntegerOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidIntegerOutputContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *ValidIntegerOutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidIntegerOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidIntegerOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidIntegerOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidIntegerOutput(s)
	}
}

func (s *ValidIntegerOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidIntegerOutput(s)
	}
}

func (p *SFeelParser) ValidIntegerOutput() (localctx IValidIntegerOutputContext) {
	localctx = NewValidIntegerOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SFeelParserRULE_validIntegerOutput)

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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(SFeelParserINTEGER)
		}
		{
			p.SetState(187)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(189)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidNumberOutputContext is an interface to support dynamic dispatch.
type IValidNumberOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidNumberOutputContext differentiates from other interfaces.
	IsValidNumberOutputContext()
}

type ValidNumberOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidNumberOutputContext() *ValidNumberOutputContext {
	var p = new(ValidNumberOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validNumberOutput
	return p
}

func (*ValidNumberOutputContext) IsValidNumberOutputContext() {}

func NewValidNumberOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidNumberOutputContext {
	var p = new(ValidNumberOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validNumberOutput

	return p
}

func (s *ValidNumberOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidNumberOutputContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ValidNumberOutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidNumberOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidNumberOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidNumberOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidNumberOutput(s)
	}
}

func (s *ValidNumberOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidNumberOutput(s)
	}
}

func (p *SFeelParser) ValidNumberOutput() (localctx IValidNumberOutputContext) {
	localctx = NewValidNumberOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SFeelParserRULE_validNumberOutput)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER, SFeelParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Number()
		}
		{
			p.SetState(193)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(196)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidStringOutputContext is an interface to support dynamic dispatch.
type IValidStringOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidStringOutputContext differentiates from other interfaces.
	IsValidStringOutputContext()
}

type ValidStringOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidStringOutputContext() *ValidStringOutputContext {
	var p = new(ValidStringOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validStringOutput
	return p
}

func (*ValidStringOutputContext) IsValidStringOutputContext() {}

func NewValidStringOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidStringOutputContext {
	var p = new(ValidStringOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validStringOutput

	return p
}

func (s *ValidStringOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidStringOutputContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *ValidStringOutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidStringOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidStringOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidStringOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidStringOutput(s)
	}
}

func (s *ValidStringOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidStringOutput(s)
	}
}

func (p *SFeelParser) ValidStringOutput() (localctx IValidStringOutputContext) {
	localctx = NewValidStringOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SFeelParserRULE_validStringOutput)

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Strings()
		}
		{
			p.SetState(200)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(203)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidBoolOutputContext is an interface to support dynamic dispatch.
type IValidBoolOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidBoolOutputContext differentiates from other interfaces.
	IsValidBoolOutputContext()
}

type ValidBoolOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidBoolOutputContext() *ValidBoolOutputContext {
	var p = new(ValidBoolOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validBoolOutput
	return p
}

func (*ValidBoolOutputContext) IsValidBoolOutputContext() {}

func NewValidBoolOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidBoolOutputContext {
	var p = new(ValidBoolOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validBoolOutput

	return p
}

func (s *ValidBoolOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidBoolOutputContext) Bools() IBoolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolsContext)
}

func (s *ValidBoolOutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidBoolOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidBoolOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidBoolOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidBoolOutput(s)
	}
}

func (s *ValidBoolOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidBoolOutput(s)
	}
}

func (p *SFeelParser) ValidBoolOutput() (localctx IValidBoolOutputContext) {
	localctx = NewValidBoolOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SFeelParserRULE_validBoolOutput)

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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Bools()
		}
		{
			p.SetState(207)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(210)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValidDateTimeOutputContext is an interface to support dynamic dispatch.
type IValidDateTimeOutputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValidDateTimeOutputContext differentiates from other interfaces.
	IsValidDateTimeOutputContext()
}

type ValidDateTimeOutputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValidDateTimeOutputContext() *ValidDateTimeOutputContext {
	var p = new(ValidDateTimeOutputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_validDateTimeOutput
	return p
}

func (*ValidDateTimeOutputContext) IsValidDateTimeOutputContext() {}

func NewValidDateTimeOutputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValidDateTimeOutputContext {
	var p = new(ValidDateTimeOutputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_validDateTimeOutput

	return p
}

func (s *ValidDateTimeOutputContext) GetParser() antlr.Parser { return s.parser }

func (s *ValidDateTimeOutputContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *ValidDateTimeOutputContext) EOF() antlr.TerminalNode {
	return s.GetToken(SFeelParserEOF, 0)
}

func (s *ValidDateTimeOutputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValidDateTimeOutputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValidDateTimeOutputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterValidDateTimeOutput(s)
	}
}

func (s *ValidDateTimeOutputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitValidDateTimeOutput(s)
	}
}

func (p *SFeelParser) ValidDateTimeOutput() (localctx IValidDateTimeOutputContext) {
	localctx = NewValidDateTimeOutputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SFeelParserRULE_validDateTimeOutput)

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

	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserDATEANDTIME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Datetime()
		}
		{
			p.SetState(214)
			p.Match(SFeelParserEOF)
		}

	case SFeelParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(SFeelParserT__2)
		}
		{
			p.SetState(217)
			p.Match(SFeelParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DisjunctionRuleContext struct {
	*ExpressionContext
}

func NewDisjunctionRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DisjunctionRuleContext {
	var p = new(DisjunctionRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DisjunctionRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionRuleContext) Disjunctions() IDisjunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsContext)
}

func (s *DisjunctionRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionRule(s)
	}
}

func (s *DisjunctionRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionRule(s)
	}
}

type ComparisionsRuleContext struct {
	*ExpressionContext
}

func NewComparisionsRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisionsRuleContext {
	var p = new(ComparisionsRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisionsRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisionsRuleContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ComparisionsRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisionsRule(s)
	}
}

func (s *ComparisionsRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisionsRule(s)
	}
}

type EmptyInputRuleContext struct {
	*ExpressionContext
}

func NewEmptyInputRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyInputRuleContext {
	var p = new(EmptyInputRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EmptyInputRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyInputRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEmptyInputRule(s)
	}
}

func (s *EmptyInputRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEmptyInputRule(s)
	}
}

type NegationRuleContext struct {
	*ExpressionContext
}

func NewNegationRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationRuleContext {
	var p = new(NegationRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegationRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationRuleContext) Negation() INegationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegationContext)
}

func (s *NegationRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNegationRule(s)
	}
}

func (s *NegationRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNegationRule(s)
	}
}

type EqualcomparisonRuleContext struct {
	*ExpressionContext
}

func NewEqualcomparisonRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualcomparisonRuleContext {
	var p = new(EqualcomparisonRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualcomparisonRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonRuleContext) Equalcomparison() IEqualcomparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonContext)
}

func (s *EqualcomparisonRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonRule(s)
	}
}

func (s *EqualcomparisonRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonRule(s)
	}
}

type RangeRuleContext struct {
	*ExpressionContext
}

func NewRangeRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeRuleContext {
	var p = new(RangeRuleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RangeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeRuleContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *RangeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangeRule(s)
	}
}

func (s *RangeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangeRule(s)
	}
}

func (p *SFeelParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SFeelParserRULE_expression)

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

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualcomparisonRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Equalcomparison()
		}

	case 2:
		localctx = NewComparisionsRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Comparison()
		}

	case 3:
		localctx = NewRangeRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(222)
			p.Ranges()
		}

	case 4:
		localctx = NewDisjunctionRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(223)
			p.Disjunctions()
		}

	case 5:
		localctx = NewNegationRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Negation()
		}

	case 6:
		localctx = NewEmptyInputRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(225)
			p.Match(SFeelParserT__2)
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SFeelParserFLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SFeelParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SFeelParserRULE_number)
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
		p.SetState(228)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserINTEGER || _la == SFeelParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringsContext is an interface to support dynamic dispatch.
type IStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringsContext differentiates from other interfaces.
	IsStringsContext()
}

type StringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringsContext() *StringsContext {
	var p = new(StringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_strings
	return p
}

func (*StringsContext) IsStringsContext() {}

func NewStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringsContext {
	var p = new(StringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_strings

	return p
}

func (s *StringsContext) GetParser() antlr.Parser { return s.parser }

func (s *StringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(SFeelParserSTRING, 0)
}

func (s *StringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterStrings(s)
	}
}

func (s *StringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitStrings(s)
	}
}

func (p *SFeelParser) Strings() (localctx IStringsContext) {
	localctx = NewStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SFeelParserRULE_strings)

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
		p.SetState(230)
		p.Match(SFeelParserSTRING)
	}

	return localctx
}

// IBoolsContext is an interface to support dynamic dispatch.
type IBoolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolsContext differentiates from other interfaces.
	IsBoolsContext()
}

type BoolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolsContext() *BoolsContext {
	var p = new(BoolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_bools
	return p
}

func (*BoolsContext) IsBoolsContext() {}

func NewBoolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolsContext {
	var p = new(BoolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_bools

	return p
}

func (s *BoolsContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SFeelParserBOOL, 0)
}

func (s *BoolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterBools(s)
	}
}

func (s *BoolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitBools(s)
	}
}

func (p *SFeelParser) Bools() (localctx IBoolsContext) {
	localctx = NewBoolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SFeelParserRULE_bools)

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
		p.SetState(232)
		p.Match(SFeelParserBOOL)
	}

	return localctx
}

// IDatetimeContext is an interface to support dynamic dispatch.
type IDatetimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeContext differentiates from other interfaces.
	IsDatetimeContext()
}

type DatetimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeContext() *DatetimeContext {
	var p = new(DatetimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_datetime
	return p
}

func (*DatetimeContext) IsDatetimeContext() {}

func NewDatetimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeContext {
	var p = new(DatetimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_datetime

	return p
}

func (s *DatetimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeContext) DATEANDTIME() antlr.TerminalNode {
	return s.GetToken(SFeelParserDATEANDTIME, 0)
}

func (s *DatetimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDatetime(s)
	}
}

func (s *DatetimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDatetime(s)
	}
}

func (p *SFeelParser) Datetime() (localctx IDatetimeContext) {
	localctx = NewDatetimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SFeelParserRULE_datetime)

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
		p.Match(SFeelParserDATEANDTIME)
	}

	return localctx
}

// IEqualcomparisonContext is an interface to support dynamic dispatch.
type IEqualcomparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonContext differentiates from other interfaces.
	IsEqualcomparisonContext()
}

type EqualcomparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonContext() *EqualcomparisonContext {
	var p = new(EqualcomparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparison
	return p
}

func (*EqualcomparisonContext) IsEqualcomparisonContext() {}

func NewEqualcomparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonContext {
	var p = new(EqualcomparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparison

	return p
}

func (s *EqualcomparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonContext) EqualcomparisonNumber() IEqualcomparisonNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonNumberContext)
}

func (s *EqualcomparisonContext) EqualcomparisonStrings() IEqualcomparisonStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonStringsContext)
}

func (s *EqualcomparisonContext) EqualcomparisonDateTime() IEqualcomparisonDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonDateTimeContext)
}

func (s *EqualcomparisonContext) EqualcomparisonBool() IEqualcomparisonBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonBoolContext)
}

func (s *EqualcomparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparison(s)
	}
}

func (s *EqualcomparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparison(s)
	}
}

func (p *SFeelParser) Equalcomparison() (localctx IEqualcomparisonContext) {
	localctx = NewEqualcomparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SFeelParserRULE_equalcomparison)

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
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER, SFeelParserFLOAT:
		{
			p.SetState(236)
			p.EqualcomparisonNumber()
		}

	case SFeelParserSTRING:
		{
			p.SetState(237)
			p.EqualcomparisonStrings()
		}

	case SFeelParserDATEANDTIME:
		{
			p.SetState(238)
			p.EqualcomparisonDateTime()
		}

	case SFeelParserBOOL:
		{
			p.SetState(239)
			p.EqualcomparisonBool()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqualcomparisonIntegerContext is an interface to support dynamic dispatch.
type IEqualcomparisonIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonIntegerContext differentiates from other interfaces.
	IsEqualcomparisonIntegerContext()
}

type EqualcomparisonIntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonIntegerContext() *EqualcomparisonIntegerContext {
	var p = new(EqualcomparisonIntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparisonInteger
	return p
}

func (*EqualcomparisonIntegerContext) IsEqualcomparisonIntegerContext() {}

func NewEqualcomparisonIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonIntegerContext {
	var p = new(EqualcomparisonIntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparisonInteger

	return p
}

func (s *EqualcomparisonIntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonIntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *EqualcomparisonIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonInteger(s)
	}
}

func (s *EqualcomparisonIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonInteger(s)
	}
}

func (p *SFeelParser) EqualcomparisonInteger() (localctx IEqualcomparisonIntegerContext) {
	localctx = NewEqualcomparisonIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SFeelParserRULE_equalcomparisonInteger)

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
		p.SetState(242)
		p.Match(SFeelParserINTEGER)
	}

	return localctx
}

// IEqualcomparisonNumberContext is an interface to support dynamic dispatch.
type IEqualcomparisonNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonNumberContext differentiates from other interfaces.
	IsEqualcomparisonNumberContext()
}

type EqualcomparisonNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonNumberContext() *EqualcomparisonNumberContext {
	var p = new(EqualcomparisonNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparisonNumber
	return p
}

func (*EqualcomparisonNumberContext) IsEqualcomparisonNumberContext() {}

func NewEqualcomparisonNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonNumberContext {
	var p = new(EqualcomparisonNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparisonNumber

	return p
}

func (s *EqualcomparisonNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonNumberContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *EqualcomparisonNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonNumber(s)
	}
}

func (s *EqualcomparisonNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonNumber(s)
	}
}

func (p *SFeelParser) EqualcomparisonNumber() (localctx IEqualcomparisonNumberContext) {
	localctx = NewEqualcomparisonNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SFeelParserRULE_equalcomparisonNumber)

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
		p.SetState(244)
		p.Number()
	}

	return localctx
}

// IEqualcomparisonBoolContext is an interface to support dynamic dispatch.
type IEqualcomparisonBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonBoolContext differentiates from other interfaces.
	IsEqualcomparisonBoolContext()
}

type EqualcomparisonBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonBoolContext() *EqualcomparisonBoolContext {
	var p = new(EqualcomparisonBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparisonBool
	return p
}

func (*EqualcomparisonBoolContext) IsEqualcomparisonBoolContext() {}

func NewEqualcomparisonBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonBoolContext {
	var p = new(EqualcomparisonBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparisonBool

	return p
}

func (s *EqualcomparisonBoolContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonBoolContext) Bools() IBoolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolsContext)
}

func (s *EqualcomparisonBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonBool(s)
	}
}

func (s *EqualcomparisonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonBool(s)
	}
}

func (p *SFeelParser) EqualcomparisonBool() (localctx IEqualcomparisonBoolContext) {
	localctx = NewEqualcomparisonBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SFeelParserRULE_equalcomparisonBool)

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
		p.SetState(246)
		p.Bools()
	}

	return localctx
}

// IEqualcomparisonStringsContext is an interface to support dynamic dispatch.
type IEqualcomparisonStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonStringsContext differentiates from other interfaces.
	IsEqualcomparisonStringsContext()
}

type EqualcomparisonStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonStringsContext() *EqualcomparisonStringsContext {
	var p = new(EqualcomparisonStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparisonStrings
	return p
}

func (*EqualcomparisonStringsContext) IsEqualcomparisonStringsContext() {}

func NewEqualcomparisonStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonStringsContext {
	var p = new(EqualcomparisonStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparisonStrings

	return p
}

func (s *EqualcomparisonStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonStringsContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *EqualcomparisonStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonStringsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonStrings(s)
	}
}

func (s *EqualcomparisonStringsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonStrings(s)
	}
}

func (p *SFeelParser) EqualcomparisonStrings() (localctx IEqualcomparisonStringsContext) {
	localctx = NewEqualcomparisonStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SFeelParserRULE_equalcomparisonStrings)

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
		p.SetState(248)
		p.Strings()
	}

	return localctx
}

// IEqualcomparisonDateTimeContext is an interface to support dynamic dispatch.
type IEqualcomparisonDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualcomparisonDateTimeContext differentiates from other interfaces.
	IsEqualcomparisonDateTimeContext()
}

type EqualcomparisonDateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualcomparisonDateTimeContext() *EqualcomparisonDateTimeContext {
	var p = new(EqualcomparisonDateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_equalcomparisonDateTime
	return p
}

func (*EqualcomparisonDateTimeContext) IsEqualcomparisonDateTimeContext() {}

func NewEqualcomparisonDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualcomparisonDateTimeContext {
	var p = new(EqualcomparisonDateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_equalcomparisonDateTime

	return p
}

func (s *EqualcomparisonDateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualcomparisonDateTimeContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *EqualcomparisonDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualcomparisonDateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualcomparisonDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterEqualcomparisonDateTime(s)
	}
}

func (s *EqualcomparisonDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitEqualcomparisonDateTime(s)
	}
}

func (p *SFeelParser) EqualcomparisonDateTime() (localctx IEqualcomparisonDateTimeContext) {
	localctx = NewEqualcomparisonDateTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SFeelParserRULE_equalcomparisonDateTime)

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
		p.SetState(250)
		p.Datetime()
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) ComparisonDateTime() IComparisonDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonDateTimeContext)
}

func (s *ComparisonContext) ComparisonInteger() IComparisonIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonIntegerContext)
}

func (s *ComparisonContext) ComparisonNumber() IComparisonNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonNumberContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *SFeelParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SFeelParserRULE_comparison)

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
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(252)
			p.ComparisonDateTime()
		}

	case 2:
		{
			p.SetState(253)
			p.ComparisonInteger()
		}

	case 3:
		{
			p.SetState(254)
			p.ComparisonNumber()
		}

	}

	return localctx
}

// IComparisonOpsContext is an interface to support dynamic dispatch.
type IComparisonOpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOpsContext differentiates from other interfaces.
	IsComparisonOpsContext()
}

type ComparisonOpsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOpsContext() *ComparisonOpsContext {
	var p = new(ComparisonOpsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisonOps
	return p
}

func (*ComparisonOpsContext) IsComparisonOpsContext() {}

func NewComparisonOpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOpsContext {
	var p = new(ComparisonOpsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisonOps

	return p
}

func (s *ComparisonOpsContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOpsContext) LESS() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESS, 0)
}

func (s *ComparisonOpsContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserLESSEQ, 0)
}

func (s *ComparisonOpsContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATER, 0)
}

func (s *ComparisonOpsContext) GREATEREQ() antlr.TerminalNode {
	return s.GetToken(SFeelParserGREATEREQ, 0)
}

func (s *ComparisonOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisonOps(s)
	}
}

func (s *ComparisonOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisonOps(s)
	}
}

func (p *SFeelParser) ComparisonOps() (localctx IComparisonOpsContext) {
	localctx = NewComparisonOpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SFeelParserRULE_comparisonOps)
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
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SFeelParserLESS)|(1<<SFeelParserLESSEQ)|(1<<SFeelParserGREATER)|(1<<SFeelParserGREATEREQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonIntegerContext is an interface to support dynamic dispatch.
type IComparisonIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonIntegerContext differentiates from other interfaces.
	IsComparisonIntegerContext()
}

type ComparisonIntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonIntegerContext() *ComparisonIntegerContext {
	var p = new(ComparisonIntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisonInteger
	return p
}

func (*ComparisonIntegerContext) IsComparisonIntegerContext() {}

func NewComparisonIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonIntegerContext {
	var p = new(ComparisonIntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisonInteger

	return p
}

func (s *ComparisonIntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonIntegerContext) ComparisonOps() IComparisonOpsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOpsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOpsContext)
}

func (s *ComparisonIntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *ComparisonIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisonInteger(s)
	}
}

func (s *ComparisonIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisonInteger(s)
	}
}

func (p *SFeelParser) ComparisonInteger() (localctx IComparisonIntegerContext) {
	localctx = NewComparisonIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SFeelParserRULE_comparisonInteger)

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
		p.SetState(259)
		p.ComparisonOps()
	}
	{
		p.SetState(260)
		p.Match(SFeelParserINTEGER)
	}

	return localctx
}

// IComparisonDateTimeContext is an interface to support dynamic dispatch.
type IComparisonDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonDateTimeContext differentiates from other interfaces.
	IsComparisonDateTimeContext()
}

type ComparisonDateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonDateTimeContext() *ComparisonDateTimeContext {
	var p = new(ComparisonDateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisonDateTime
	return p
}

func (*ComparisonDateTimeContext) IsComparisonDateTimeContext() {}

func NewComparisonDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonDateTimeContext {
	var p = new(ComparisonDateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisonDateTime

	return p
}

func (s *ComparisonDateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonDateTimeContext) ComparisonOps() IComparisonOpsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOpsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOpsContext)
}

func (s *ComparisonDateTimeContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *ComparisonDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonDateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisonDateTime(s)
	}
}

func (s *ComparisonDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisonDateTime(s)
	}
}

func (p *SFeelParser) ComparisonDateTime() (localctx IComparisonDateTimeContext) {
	localctx = NewComparisonDateTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SFeelParserRULE_comparisonDateTime)

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
		p.SetState(262)
		p.ComparisonOps()
	}
	{
		p.SetState(263)
		p.Datetime()
	}

	return localctx
}

// IComparisonNumberContext is an interface to support dynamic dispatch.
type IComparisonNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonNumberContext differentiates from other interfaces.
	IsComparisonNumberContext()
}

type ComparisonNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonNumberContext() *ComparisonNumberContext {
	var p = new(ComparisonNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_comparisonNumber
	return p
}

func (*ComparisonNumberContext) IsComparisonNumberContext() {}

func NewComparisonNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonNumberContext {
	var p = new(ComparisonNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_comparisonNumber

	return p
}

func (s *ComparisonNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonNumberContext) ComparisonOps() IComparisonOpsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOpsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOpsContext)
}

func (s *ComparisonNumberContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ComparisonNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterComparisonNumber(s)
	}
}

func (s *ComparisonNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitComparisonNumber(s)
	}
}

func (p *SFeelParser) ComparisonNumber() (localctx IComparisonNumberContext) {
	localctx = NewComparisonNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SFeelParserRULE_comparisonNumber)

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
		p.SetState(265)
		p.ComparisonOps()
	}
	{
		p.SetState(266)
		p.Number()
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) RangeInteger() IRangeIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeIntegerContext)
}

func (s *RangesContext) RangeDateTime() IRangeDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDateTimeContext)
}

func (s *RangesContext) RangeNumber() IRangeNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeNumberContext)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *SFeelParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SFeelParserRULE_ranges)

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
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(268)
			p.RangeInteger()
		}

	case 2:
		{
			p.SetState(269)
			p.RangeDateTime()
		}

	case 3:
		{
			p.SetState(270)
			p.RangeNumber()
		}

	}

	return localctx
}

// IRopContext is an interface to support dynamic dispatch.
type IRopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRopContext differentiates from other interfaces.
	IsRopContext()
}

type RopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRopContext() *RopContext {
	var p = new(RopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rop
	return p
}

func (*RopContext) IsRopContext() {}

func NewRopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RopContext {
	var p = new(RopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rop

	return p
}

func (s *RopContext) GetParser() antlr.Parser { return s.parser }

func (s *RopContext) RANGEIN() antlr.TerminalNode {
	return s.GetToken(SFeelParserRANGEIN, 0)
}

func (s *RopContext) RANGEOUT() antlr.TerminalNode {
	return s.GetToken(SFeelParserRANGEOUT, 0)
}

func (s *RopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRop(s)
	}
}

func (s *RopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRop(s)
	}
}

func (p *SFeelParser) Rop() (localctx IRopContext) {
	localctx = NewRopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SFeelParserRULE_rop)
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
		p.SetState(273)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SFeelParserRANGEIN || _la == SFeelParserRANGEOUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRangeIntegerContext is an interface to support dynamic dispatch.
type IRangeIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeIntegerContext differentiates from other interfaces.
	IsRangeIntegerContext()
}

type RangeIntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeIntegerContext() *RangeIntegerContext {
	var p = new(RangeIntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rangeInteger
	return p
}

func (*RangeIntegerContext) IsRangeIntegerContext() {}

func NewRangeIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeIntegerContext {
	var p = new(RangeIntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rangeInteger

	return p
}

func (s *RangeIntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeIntegerContext) AllRop() []IRopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRopContext)(nil)).Elem())
	var tst = make([]IRopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRopContext)
		}
	}

	return tst
}

func (s *RangeIntegerContext) Rop(i int) IRopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRopContext)
}

func (s *RangeIntegerContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SFeelParserINTEGER)
}

func (s *RangeIntegerContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, i)
}

func (s *RangeIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangeInteger(s)
	}
}

func (s *RangeIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangeInteger(s)
	}
}

func (p *SFeelParser) RangeInteger() (localctx IRangeIntegerContext) {
	localctx = NewRangeIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SFeelParserRULE_rangeInteger)

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
		p.SetState(275)
		p.Rop()
	}
	{
		p.SetState(276)
		p.Match(SFeelParserINTEGER)
	}
	{
		p.SetState(277)
		p.Match(SFeelParserT__3)
	}
	{
		p.SetState(278)
		p.Match(SFeelParserINTEGER)
	}
	{
		p.SetState(279)
		p.Rop()
	}

	return localctx
}

// IRangeNumberContext is an interface to support dynamic dispatch.
type IRangeNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeNumberContext differentiates from other interfaces.
	IsRangeNumberContext()
}

type RangeNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeNumberContext() *RangeNumberContext {
	var p = new(RangeNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rangeNumber
	return p
}

func (*RangeNumberContext) IsRangeNumberContext() {}

func NewRangeNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeNumberContext {
	var p = new(RangeNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rangeNumber

	return p
}

func (s *RangeNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeNumberContext) AllRop() []IRopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRopContext)(nil)).Elem())
	var tst = make([]IRopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRopContext)
		}
	}

	return tst
}

func (s *RangeNumberContext) Rop(i int) IRopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRopContext)
}

func (s *RangeNumberContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *RangeNumberContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RangeNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangeNumber(s)
	}
}

func (s *RangeNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangeNumber(s)
	}
}

func (p *SFeelParser) RangeNumber() (localctx IRangeNumberContext) {
	localctx = NewRangeNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SFeelParserRULE_rangeNumber)

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
		p.SetState(281)
		p.Rop()
	}
	{
		p.SetState(282)
		p.Number()
	}
	{
		p.SetState(283)
		p.Match(SFeelParserT__3)
	}
	{
		p.SetState(284)
		p.Number()
	}
	{
		p.SetState(285)
		p.Rop()
	}

	return localctx
}

// IRangeDateTimeContext is an interface to support dynamic dispatch.
type IRangeDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeDateTimeContext differentiates from other interfaces.
	IsRangeDateTimeContext()
}

type RangeDateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeDateTimeContext() *RangeDateTimeContext {
	var p = new(RangeDateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_rangeDateTime
	return p
}

func (*RangeDateTimeContext) IsRangeDateTimeContext() {}

func NewRangeDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeDateTimeContext {
	var p = new(RangeDateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_rangeDateTime

	return p
}

func (s *RangeDateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeDateTimeContext) AllRop() []IRopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRopContext)(nil)).Elem())
	var tst = make([]IRopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRopContext)
		}
	}

	return tst
}

func (s *RangeDateTimeContext) Rop(i int) IRopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRopContext)
}

func (s *RangeDateTimeContext) AllDatetime() []IDatetimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatetimeContext)(nil)).Elem())
	var tst = make([]IDatetimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatetimeContext)
		}
	}

	return tst
}

func (s *RangeDateTimeContext) Datetime(i int) IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *RangeDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeDateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterRangeDateTime(s)
	}
}

func (s *RangeDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitRangeDateTime(s)
	}
}

func (p *SFeelParser) RangeDateTime() (localctx IRangeDateTimeContext) {
	localctx = NewRangeDateTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SFeelParserRULE_rangeDateTime)

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
		p.SetState(287)
		p.Rop()
	}

	{
		p.SetState(288)
		p.Datetime()
	}
	{
		p.SetState(289)
		p.Match(SFeelParserT__3)
	}
	{
		p.SetState(290)
		p.Datetime()
	}

	{
		p.SetState(292)
		p.Rop()
	}

	return localctx
}

// IDisjunctionsContext is an interface to support dynamic dispatch.
type IDisjunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsContext differentiates from other interfaces.
	IsDisjunctionsContext()
}

type DisjunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsContext() *DisjunctionsContext {
	var p = new(DisjunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctions
	return p
}

func (*DisjunctionsContext) IsDisjunctionsContext() {}

func NewDisjunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsContext {
	var p = new(DisjunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctions

	return p
}

func (s *DisjunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsContext) DisjunctionsInteger() IDisjunctionsIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsIntegerContext)
}

func (s *DisjunctionsContext) DisjunctionsNumber() IDisjunctionsNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsNumberContext)
}

func (s *DisjunctionsContext) DisjunctionsString() IDisjunctionsStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsStringContext)
}

func (s *DisjunctionsContext) DisjunctionsDateTime() IDisjunctionsDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsDateTimeContext)
}

func (s *DisjunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctions(s)
	}
}

func (s *DisjunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctions(s)
	}
}

func (p *SFeelParser) Disjunctions() (localctx IDisjunctionsContext) {
	localctx = NewDisjunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SFeelParserRULE_disjunctions)

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
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(294)
			p.disjunctionsInteger(0)
		}

	case 2:
		{
			p.SetState(295)
			p.disjunctionsNumber(0)
		}

	case 3:
		{
			p.SetState(296)
			p.disjunctionsString(0)
		}

	case 4:
		{
			p.SetState(297)
			p.disjunctionsDateTime(0)
		}

	}

	return localctx
}

// IDisjunctionsIntegerContext is an interface to support dynamic dispatch.
type IDisjunctionsIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsIntegerContext differentiates from other interfaces.
	IsDisjunctionsIntegerContext()
}

type DisjunctionsIntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsIntegerContext() *DisjunctionsIntegerContext {
	var p = new(DisjunctionsIntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsInteger
	return p
}

func (*DisjunctionsIntegerContext) IsDisjunctionsIntegerContext() {}

func NewDisjunctionsIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsIntegerContext {
	var p = new(DisjunctionsIntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsInteger

	return p
}

func (s *DisjunctionsIntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsIntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SFeelParserINTEGER, 0)
}

func (s *DisjunctionsIntegerContext) ComparisonInteger() IComparisonIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonIntegerContext)
}

func (s *DisjunctionsIntegerContext) RangeInteger() IRangeIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeIntegerContext)
}

func (s *DisjunctionsIntegerContext) AllDisjunctionsInteger() []IDisjunctionsIntegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsIntegerContext)(nil)).Elem())
	var tst = make([]IDisjunctionsIntegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsIntegerContext)
		}
	}

	return tst
}

func (s *DisjunctionsIntegerContext) DisjunctionsInteger(i int) IDisjunctionsIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsIntegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsIntegerContext)
}

func (s *DisjunctionsIntegerContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsIntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsInteger(s)
	}
}

func (s *DisjunctionsIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsInteger(s)
	}
}

func (p *SFeelParser) DisjunctionsInteger() (localctx IDisjunctionsIntegerContext) {
	return p.disjunctionsInteger(0)
}

func (p *SFeelParser) disjunctionsInteger(_p int) (localctx IDisjunctionsIntegerContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsIntegerContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsIntegerContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, SFeelParserRULE_disjunctionsInteger, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER:
		{
			p.SetState(301)
			p.Match(SFeelParserINTEGER)
		}

	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		{
			p.SetState(302)
			p.ComparisonInteger()
		}

	case SFeelParserRANGEIN, SFeelParserRANGEOUT:
		{
			p.SetState(303)
			p.RangeInteger()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsIntegerContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsInteger)
			p.SetState(306)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(307)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(308)
				p.disjunctionsInteger(3)
			}

		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IDisjunctionsNumberContext is an interface to support dynamic dispatch.
type IDisjunctionsNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsNumberContext differentiates from other interfaces.
	IsDisjunctionsNumberContext()
}

type DisjunctionsNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsNumberContext() *DisjunctionsNumberContext {
	var p = new(DisjunctionsNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsNumber
	return p
}

func (*DisjunctionsNumberContext) IsDisjunctionsNumberContext() {}

func NewDisjunctionsNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsNumberContext {
	var p = new(DisjunctionsNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsNumber

	return p
}

func (s *DisjunctionsNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsNumberContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *DisjunctionsNumberContext) ComparisonNumber() IComparisonNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonNumberContext)
}

func (s *DisjunctionsNumberContext) RangeNumber() IRangeNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeNumberContext)
}

func (s *DisjunctionsNumberContext) AllDisjunctionsNumber() []IDisjunctionsNumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem())
	var tst = make([]IDisjunctionsNumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsNumberContext)
		}
	}

	return tst
}

func (s *DisjunctionsNumberContext) DisjunctionsNumber(i int) IDisjunctionsNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsNumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsNumberContext)
}

func (s *DisjunctionsNumberContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsNumber(s)
	}
}

func (s *DisjunctionsNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsNumber(s)
	}
}

func (p *SFeelParser) DisjunctionsNumber() (localctx IDisjunctionsNumberContext) {
	return p.disjunctionsNumber(0)
}

func (p *SFeelParser) disjunctionsNumber(_p int) (localctx IDisjunctionsNumberContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsNumberContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsNumberContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, SFeelParserRULE_disjunctionsNumber, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserINTEGER, SFeelParserFLOAT:
		{
			p.SetState(315)
			p.Number()
		}

	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		{
			p.SetState(316)
			p.ComparisonNumber()
		}

	case SFeelParserRANGEIN, SFeelParserRANGEOUT:
		{
			p.SetState(317)
			p.RangeNumber()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsNumberContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsNumber)
			p.SetState(320)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(321)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(322)
				p.disjunctionsNumber(3)
			}

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IDisjunctionsStringContext is an interface to support dynamic dispatch.
type IDisjunctionsStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsStringContext differentiates from other interfaces.
	IsDisjunctionsStringContext()
}

type DisjunctionsStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsStringContext() *DisjunctionsStringContext {
	var p = new(DisjunctionsStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsString
	return p
}

func (*DisjunctionsStringContext) IsDisjunctionsStringContext() {}

func NewDisjunctionsStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsStringContext {
	var p = new(DisjunctionsStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsString

	return p
}

func (s *DisjunctionsStringContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsStringContext) Strings() IStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsContext)
}

func (s *DisjunctionsStringContext) AllDisjunctionsString() []IDisjunctionsStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem())
	var tst = make([]IDisjunctionsStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsStringContext)
		}
	}

	return tst
}

func (s *DisjunctionsStringContext) DisjunctionsString(i int) IDisjunctionsStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsStringContext)
}

func (s *DisjunctionsStringContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsString(s)
	}
}

func (s *DisjunctionsStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsString(s)
	}
}

func (p *SFeelParser) DisjunctionsString() (localctx IDisjunctionsStringContext) {
	return p.disjunctionsString(0)
}

func (p *SFeelParser) disjunctionsString(_p int) (localctx IDisjunctionsStringContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsStringContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsStringContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, SFeelParserRULE_disjunctionsString, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Strings()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsStringContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsString)
			p.SetState(331)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(332)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(333)
				p.disjunctionsString(3)
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IDisjunctionsDateTimeContext is an interface to support dynamic dispatch.
type IDisjunctionsDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionsDateTimeContext differentiates from other interfaces.
	IsDisjunctionsDateTimeContext()
}

type DisjunctionsDateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionsDateTimeContext() *DisjunctionsDateTimeContext {
	var p = new(DisjunctionsDateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_disjunctionsDateTime
	return p
}

func (*DisjunctionsDateTimeContext) IsDisjunctionsDateTimeContext() {}

func NewDisjunctionsDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionsDateTimeContext {
	var p = new(DisjunctionsDateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_disjunctionsDateTime

	return p
}

func (s *DisjunctionsDateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionsDateTimeContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *DisjunctionsDateTimeContext) ComparisonDateTime() IComparisonDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonDateTimeContext)
}

func (s *DisjunctionsDateTimeContext) RangeDateTime() IRangeDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDateTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDateTimeContext)
}

func (s *DisjunctionsDateTimeContext) AllDisjunctionsDateTime() []IDisjunctionsDateTimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem())
	var tst = make([]IDisjunctionsDateTimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDisjunctionsDateTimeContext)
		}
	}

	return tst
}

func (s *DisjunctionsDateTimeContext) DisjunctionsDateTime(i int) IDisjunctionsDateTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsDateTimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsDateTimeContext)
}

func (s *DisjunctionsDateTimeContext) DISJUNCTION() antlr.TerminalNode {
	return s.GetToken(SFeelParserDISJUNCTION, 0)
}

func (s *DisjunctionsDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionsDateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionsDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterDisjunctionsDateTime(s)
	}
}

func (s *DisjunctionsDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitDisjunctionsDateTime(s)
	}
}

func (p *SFeelParser) DisjunctionsDateTime() (localctx IDisjunctionsDateTimeContext) {
	return p.disjunctionsDateTime(0)
}

func (p *SFeelParser) disjunctionsDateTime(_p int) (localctx IDisjunctionsDateTimeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDisjunctionsDateTimeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionsDateTimeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 72
	p.EnterRecursionRule(localctx, 72, SFeelParserRULE_disjunctionsDateTime, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(343)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SFeelParserDATEANDTIME:
		{
			p.SetState(340)
			p.Datetime()
		}

	case SFeelParserLESS, SFeelParserLESSEQ, SFeelParserGREATER, SFeelParserGREATEREQ:
		{
			p.SetState(341)
			p.ComparisonDateTime()
		}

	case SFeelParserRANGEIN, SFeelParserRANGEOUT:
		{
			p.SetState(342)
			p.RangeDateTime()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionsDateTimeContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SFeelParserRULE_disjunctionsDateTime)
			p.SetState(345)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(346)
				p.Match(SFeelParserDISJUNCTION)
			}
			{
				p.SetState(347)
				p.disjunctionsDateTime(3)
			}

		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// INegationContext is an interface to support dynamic dispatch.
type INegationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegationContext differentiates from other interfaces.
	IsNegationContext()
}

type NegationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegationContext() *NegationContext {
	var p = new(NegationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SFeelParserRULE_negation
	return p
}

func (*NegationContext) IsNegationContext() {}

func NewNegationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegationContext {
	var p = new(NegationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SFeelParserRULE_negation

	return p
}

func (s *NegationContext) GetParser() antlr.Parser { return s.parser }

func (s *NegationContext) NEGATION() antlr.TerminalNode {
	return s.GetToken(SFeelParserNEGATION, 0)
}

func (s *NegationContext) Equalcomparison() IEqualcomparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualcomparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualcomparisonContext)
}

func (s *NegationContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *NegationContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *NegationContext) Disjunctions() IDisjunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionsContext)
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SFeelListener); ok {
		listenerT.ExitNegation(s)
	}
}

func (p *SFeelParser) Negation() (localctx INegationContext) {
	localctx = NewNegationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SFeelParserRULE_negation)

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
		p.SetState(353)
		p.Match(SFeelParserNEGATION)
	}
	{
		p.SetState(354)
		p.Match(SFeelParserT__0)
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(355)
			p.Equalcomparison()
		}

	case 2:
		{
			p.SetState(356)
			p.Comparison()
		}

	case 3:
		{
			p.SetState(357)
			p.Ranges()
		}

	case 4:
		{
			p.SetState(358)
			p.Disjunctions()
		}

	}
	{
		p.SetState(361)
		p.Match(SFeelParserT__1)
	}

	return localctx
}

func (p *SFeelParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 33:
		var t *DisjunctionsIntegerContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsIntegerContext)
		}
		return p.DisjunctionsInteger_Sempred(t, predIndex)

	case 34:
		var t *DisjunctionsNumberContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsNumberContext)
		}
		return p.DisjunctionsNumber_Sempred(t, predIndex)

	case 35:
		var t *DisjunctionsStringContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsStringContext)
		}
		return p.DisjunctionsString_Sempred(t, predIndex)

	case 36:
		var t *DisjunctionsDateTimeContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionsDateTimeContext)
		}
		return p.DisjunctionsDateTime_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SFeelParser) DisjunctionsInteger_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) DisjunctionsNumber_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) DisjunctionsString_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SFeelParser) DisjunctionsDateTime_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
