// Code generated from GraphQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GraphQL

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

type GraphQLParser struct {
	*antlr.BaseParser
}

var GraphQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func graphqlParserInit() {
	staticData := &GraphQLParserStaticData
	staticData.LiteralNames = []string{
		"", "'query'", "'mutation'", "'subscription'", "'{'", "'}'", "'('",
		"')'", "':'", "'...'", "'fragment'", "'on'", "'true'", "'false'", "'null'",
		"'['", "']'", "'$'", "'='", "'!'", "'@'", "'schema'", "'extend'", "'scalar'",
		"'type'", "'implements'", "'&'", "'interface'", "'union'", "'|'", "'enum'",
		"'input'", "'directive'", "'repeatable'", "'QUERY'", "'MUTATION'", "'SUBSCRIPTION'",
		"'FIELD'", "'FRAGMENT_DEFINITION'", "'FRAGMENT_SPREAD'", "'INLINE_FRAGMENT'",
		"'VARIABLE_DEFINITION'", "'SCHEMA'", "'SCALAR'", "'OBJECT'", "'FIELD_DEFINITION'",
		"'ARGUMENT_DEFINITION'", "'INTERFACE'", "'UNION'", "'ENUM'", "'ENUM_VALUE'",
		"'INPUT_OBJECT'", "'INPUT_FIELD_DEFINITION'", "", "", "", "", "", "",
		"", "", "','", "", "", "'\\uEFBBBF'", "'\\uFEFF'", "'\\u0000FEFF'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "NAME", "STRING", "BLOCK_STRING", "ID", "FLOAT", "INT", "PUNCTUATOR",
		"WS", "COMMA", "LineComment", "UNICODE_BOM", "UTF8_BOM", "UTF16_BOM",
		"UTF32_BOM",
	}
	staticData.RuleNames = []string{
		"document", "definition", "executableDocument", "executableDefinition",
		"operationDefinition", "operationType", "selectionSet", "selection",
		"field", "arguments", "argument", "alias", "fragmentSpread", "fragmentDefinition",
		"fragmentName", "typeCondition", "inlineFragment", "value", "intValue",
		"floatValue", "booleanValue", "stringValue", "nullValue", "enumValue",
		"listValue", "objectValue", "objectField", "variable", "variableDefinitions",
		"variableDefinition", "defaultValue", "type_", "namedType", "listType",
		"directives", "directive", "typeSystemDocument", "typeSystemDefinition",
		"typeSystemExtensionDocument", "typeSystemExtension", "schemaDefinition",
		"rootOperationTypeDefinition", "schemaExtension", "description", "typeDefinition",
		"typeExtension", "scalarTypeDefinition", "scalarTypeExtension", "objectTypeDefinition",
		"implementsInterfaces", "fieldsDefinition", "fieldDefinition", "argumentsDefinition",
		"inputValueDefinition", "objectTypeExtension", "interfaceTypeDefinition",
		"interfaceTypeExtension", "unionTypeDefinition", "unionMemberTypes",
		"unionTypeExtension", "enumTypeDefinition", "enumValuesDefinition",
		"enumValueDefinition", "enumTypeExtension", "inputObjectTypeDefinition",
		"inputFieldsDefinition", "inputObjectTypeExtension", "directiveDefinition",
		"directiveLocations", "directiveLocation", "executableDirectiveLocation",
		"typeSystemDirectiveLocation", "name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 707, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 1, 0, 4,
		0, 148, 8, 0, 11, 0, 12, 0, 149, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 157,
		8, 1, 1, 2, 4, 2, 160, 8, 2, 11, 2, 12, 2, 161, 1, 3, 1, 3, 3, 3, 166,
		8, 3, 1, 4, 1, 4, 3, 4, 170, 8, 4, 1, 4, 3, 4, 173, 8, 4, 1, 4, 3, 4, 176,
		8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 181, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 4, 6,
		187, 8, 6, 11, 6, 12, 6, 188, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 196,
		8, 7, 1, 8, 3, 8, 199, 8, 8, 1, 8, 1, 8, 3, 8, 203, 8, 8, 1, 8, 3, 8, 206,
		8, 8, 1, 8, 3, 8, 209, 8, 8, 1, 9, 1, 9, 4, 9, 213, 8, 9, 11, 9, 12, 9,
		214, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 3, 12, 229, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 235,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3,
		16, 246, 8, 16, 1, 16, 3, 16, 249, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 262, 8, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 4, 24, 280, 8, 24, 11, 24, 12, 24, 281,
		1, 24, 1, 24, 3, 24, 286, 8, 24, 1, 25, 1, 25, 5, 25, 290, 8, 25, 10, 25,
		12, 25, 293, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 4, 28, 306, 8, 28, 11, 28, 12, 28, 307, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 316, 8, 29, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 3, 31, 323, 8, 31, 1, 31, 1, 31, 3, 31, 327, 8, 31, 3,
		31, 329, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 4, 34,
		338, 8, 34, 11, 34, 12, 34, 339, 1, 35, 1, 35, 1, 35, 3, 35, 345, 8, 35,
		1, 36, 4, 36, 348, 8, 36, 11, 36, 12, 36, 349, 1, 37, 1, 37, 1, 37, 3,
		37, 355, 8, 37, 1, 38, 4, 38, 358, 8, 38, 11, 38, 12, 38, 359, 1, 39, 1,
		39, 3, 39, 364, 8, 39, 1, 40, 1, 40, 3, 40, 368, 8, 40, 1, 40, 1, 40, 4,
		40, 372, 8, 40, 11, 40, 12, 40, 373, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 3, 42, 385, 8, 42, 1, 42, 1, 42, 4, 42, 389,
		8, 42, 11, 42, 12, 42, 390, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 398,
		8, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 408,
		8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 416, 8, 45, 1,
		46, 3, 46, 419, 8, 46, 1, 46, 1, 46, 1, 46, 3, 46, 424, 8, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 48, 3, 48, 432, 8, 48, 1, 48, 1, 48, 1, 48,
		3, 48, 437, 8, 48, 1, 48, 3, 48, 440, 8, 48, 1, 48, 3, 48, 443, 8, 48,
		1, 49, 1, 49, 1, 49, 3, 49, 448, 8, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 5, 49, 455, 8, 49, 10, 49, 12, 49, 458, 9, 49, 1, 50, 1, 50, 4, 50,
		462, 8, 50, 11, 50, 12, 50, 463, 1, 50, 1, 50, 1, 51, 3, 51, 469, 8, 51,
		1, 51, 1, 51, 3, 51, 473, 8, 51, 1, 51, 1, 51, 1, 51, 3, 51, 478, 8, 51,
		1, 52, 1, 52, 4, 52, 482, 8, 52, 11, 52, 12, 52, 483, 1, 52, 1, 52, 1,
		53, 3, 53, 489, 8, 53, 1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 495, 8, 53, 1,
		53, 3, 53, 498, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 504, 8, 54, 1,
		54, 3, 54, 507, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54,
		515, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 524,
		8, 54, 1, 55, 3, 55, 527, 8, 55, 1, 55, 1, 55, 1, 55, 3, 55, 532, 8, 55,
		1, 55, 3, 55, 535, 8, 55, 1, 55, 3, 55, 538, 8, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 3, 56, 544, 8, 56, 1, 56, 3, 56, 547, 8, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 3, 56, 555, 8, 56, 1, 56, 1, 56, 3, 56, 559, 8, 56,
		1, 57, 3, 57, 562, 8, 57, 1, 57, 1, 57, 1, 57, 3, 57, 567, 8, 57, 1, 57,
		3, 57, 570, 8, 57, 1, 58, 1, 58, 3, 58, 574, 8, 58, 1, 58, 1, 58, 1, 58,
		5, 58, 579, 8, 58, 10, 58, 12, 58, 582, 9, 58, 1, 59, 1, 59, 1, 59, 1,
		59, 3, 59, 588, 8, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59,
		3, 59, 597, 8, 59, 1, 60, 3, 60, 600, 8, 60, 1, 60, 1, 60, 1, 60, 3, 60,
		605, 8, 60, 1, 60, 3, 60, 608, 8, 60, 1, 61, 1, 61, 4, 61, 612, 8, 61,
		11, 61, 12, 61, 613, 1, 61, 1, 61, 1, 62, 3, 62, 619, 8, 62, 1, 62, 1,
		62, 3, 62, 623, 8, 62, 1, 63, 1, 63, 1, 63, 1, 63, 3, 63, 629, 8, 63, 1,
		63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 3, 63, 638, 8, 63, 1, 64,
		3, 64, 641, 8, 64, 1, 64, 1, 64, 1, 64, 3, 64, 646, 8, 64, 1, 64, 3, 64,
		649, 8, 64, 1, 65, 1, 65, 4, 65, 653, 8, 65, 11, 65, 12, 65, 654, 1, 65,
		1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 3, 66, 663, 8, 66, 1, 66, 1, 66, 1,
		66, 1, 66, 1, 66, 1, 66, 1, 66, 3, 66, 672, 8, 66, 1, 67, 3, 67, 675, 8,
		67, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 681, 8, 67, 1, 67, 3, 67, 684, 8,
		67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 5, 68, 692, 8, 68, 10, 68,
		12, 68, 695, 9, 68, 1, 69, 1, 69, 3, 69, 699, 8, 69, 1, 70, 1, 70, 1, 71,
		1, 71, 1, 72, 1, 72, 1, 72, 0, 1, 98, 73, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120,
		122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 0, 5, 1, 0,
		1, 3, 1, 0, 12, 13, 1, 0, 54, 55, 1, 0, 34, 41, 1, 0, 42, 52, 747, 0, 147,
		1, 0, 0, 0, 2, 156, 1, 0, 0, 0, 4, 159, 1, 0, 0, 0, 6, 165, 1, 0, 0, 0,
		8, 180, 1, 0, 0, 0, 10, 182, 1, 0, 0, 0, 12, 184, 1, 0, 0, 0, 14, 195,
		1, 0, 0, 0, 16, 198, 1, 0, 0, 0, 18, 210, 1, 0, 0, 0, 20, 218, 1, 0, 0,
		0, 22, 222, 1, 0, 0, 0, 24, 225, 1, 0, 0, 0, 26, 230, 1, 0, 0, 0, 28, 238,
		1, 0, 0, 0, 30, 240, 1, 0, 0, 0, 32, 243, 1, 0, 0, 0, 34, 261, 1, 0, 0,
		0, 36, 263, 1, 0, 0, 0, 38, 265, 1, 0, 0, 0, 40, 267, 1, 0, 0, 0, 42, 269,
		1, 0, 0, 0, 44, 271, 1, 0, 0, 0, 46, 273, 1, 0, 0, 0, 48, 285, 1, 0, 0,
		0, 50, 287, 1, 0, 0, 0, 52, 296, 1, 0, 0, 0, 54, 300, 1, 0, 0, 0, 56, 303,
		1, 0, 0, 0, 58, 311, 1, 0, 0, 0, 60, 317, 1, 0, 0, 0, 62, 328, 1, 0, 0,
		0, 64, 330, 1, 0, 0, 0, 66, 332, 1, 0, 0, 0, 68, 337, 1, 0, 0, 0, 70, 341,
		1, 0, 0, 0, 72, 347, 1, 0, 0, 0, 74, 354, 1, 0, 0, 0, 76, 357, 1, 0, 0,
		0, 78, 363, 1, 0, 0, 0, 80, 365, 1, 0, 0, 0, 82, 377, 1, 0, 0, 0, 84, 397,
		1, 0, 0, 0, 86, 399, 1, 0, 0, 0, 88, 407, 1, 0, 0, 0, 90, 415, 1, 0, 0,
		0, 92, 418, 1, 0, 0, 0, 94, 425, 1, 0, 0, 0, 96, 431, 1, 0, 0, 0, 98, 444,
		1, 0, 0, 0, 100, 459, 1, 0, 0, 0, 102, 468, 1, 0, 0, 0, 104, 479, 1, 0,
		0, 0, 106, 488, 1, 0, 0, 0, 108, 523, 1, 0, 0, 0, 110, 526, 1, 0, 0, 0,
		112, 558, 1, 0, 0, 0, 114, 561, 1, 0, 0, 0, 116, 571, 1, 0, 0, 0, 118,
		596, 1, 0, 0, 0, 120, 599, 1, 0, 0, 0, 122, 609, 1, 0, 0, 0, 124, 618,
		1, 0, 0, 0, 126, 637, 1, 0, 0, 0, 128, 640, 1, 0, 0, 0, 130, 650, 1, 0,
		0, 0, 132, 671, 1, 0, 0, 0, 134, 674, 1, 0, 0, 0, 136, 688, 1, 0, 0, 0,
		138, 698, 1, 0, 0, 0, 140, 700, 1, 0, 0, 0, 142, 702, 1, 0, 0, 0, 144,
		704, 1, 0, 0, 0, 146, 148, 3, 2, 1, 0, 147, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 152, 5, 0, 0, 1, 152, 1, 1, 0, 0, 0, 153, 157, 3, 4, 2, 0, 154,
		157, 3, 72, 36, 0, 155, 157, 3, 76, 38, 0, 156, 153, 1, 0, 0, 0, 156, 154,
		1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 3, 1, 0, 0, 0, 158, 160, 3, 6, 3,
		0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 5, 1, 0, 0, 0, 163, 166, 3, 8, 4, 0, 164, 166, 3,
		26, 13, 0, 165, 163, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166, 7, 1, 0, 0,
		0, 167, 169, 3, 10, 5, 0, 168, 170, 3, 144, 72, 0, 169, 168, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 173, 3, 56, 28, 0, 172,
		171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174, 176,
		3, 68, 34, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 1,
		0, 0, 0, 177, 178, 3, 12, 6, 0, 178, 181, 1, 0, 0, 0, 179, 181, 3, 12,
		6, 0, 180, 167, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 9, 1, 0, 0, 0, 182,
		183, 7, 0, 0, 0, 183, 11, 1, 0, 0, 0, 184, 186, 5, 4, 0, 0, 185, 187, 3,
		14, 7, 0, 186, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 5, 5, 0, 0, 191,
		13, 1, 0, 0, 0, 192, 196, 3, 16, 8, 0, 193, 196, 3, 24, 12, 0, 194, 196,
		3, 32, 16, 0, 195, 192, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 194, 1,
		0, 0, 0, 196, 15, 1, 0, 0, 0, 197, 199, 3, 22, 11, 0, 198, 197, 1, 0, 0,
		0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 3, 144, 72, 0,
		201, 203, 3, 18, 9, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		205, 1, 0, 0, 0, 204, 206, 3, 68, 34, 0, 205, 204, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 208, 1, 0, 0, 0, 207, 209, 3, 12, 6, 0, 208, 207, 1, 0,
		0, 0, 208, 209, 1, 0, 0, 0, 209, 17, 1, 0, 0, 0, 210, 212, 5, 6, 0, 0,
		211, 213, 3, 20, 10, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217,
		5, 7, 0, 0, 217, 19, 1, 0, 0, 0, 218, 219, 3, 144, 72, 0, 219, 220, 5,
		8, 0, 0, 220, 221, 3, 34, 17, 0, 221, 21, 1, 0, 0, 0, 222, 223, 3, 144,
		72, 0, 223, 224, 5, 8, 0, 0, 224, 23, 1, 0, 0, 0, 225, 226, 5, 9, 0, 0,
		226, 228, 3, 28, 14, 0, 227, 229, 3, 68, 34, 0, 228, 227, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 25, 1, 0, 0, 0, 230, 231, 5, 10, 0, 0, 231, 232,
		3, 28, 14, 0, 232, 234, 3, 30, 15, 0, 233, 235, 3, 68, 34, 0, 234, 233,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 3, 12,
		6, 0, 237, 27, 1, 0, 0, 0, 238, 239, 3, 144, 72, 0, 239, 29, 1, 0, 0, 0,
		240, 241, 5, 11, 0, 0, 241, 242, 3, 64, 32, 0, 242, 31, 1, 0, 0, 0, 243,
		245, 5, 9, 0, 0, 244, 246, 3, 30, 15, 0, 245, 244, 1, 0, 0, 0, 245, 246,
		1, 0, 0, 0, 246, 248, 1, 0, 0, 0, 247, 249, 3, 68, 34, 0, 248, 247, 1,
		0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 3, 12, 6,
		0, 251, 33, 1, 0, 0, 0, 252, 262, 3, 54, 27, 0, 253, 262, 3, 36, 18, 0,
		254, 262, 3, 38, 19, 0, 255, 262, 3, 42, 21, 0, 256, 262, 3, 40, 20, 0,
		257, 262, 3, 44, 22, 0, 258, 262, 3, 46, 23, 0, 259, 262, 3, 48, 24, 0,
		260, 262, 3, 50, 25, 0, 261, 252, 1, 0, 0, 0, 261, 253, 1, 0, 0, 0, 261,
		254, 1, 0, 0, 0, 261, 255, 1, 0, 0, 0, 261, 256, 1, 0, 0, 0, 261, 257,
		1, 0, 0, 0, 261, 258, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 260, 1, 0,
		0, 0, 262, 35, 1, 0, 0, 0, 263, 264, 5, 58, 0, 0, 264, 37, 1, 0, 0, 0,
		265, 266, 5, 57, 0, 0, 266, 39, 1, 0, 0, 0, 267, 268, 7, 1, 0, 0, 268,
		41, 1, 0, 0, 0, 269, 270, 7, 2, 0, 0, 270, 43, 1, 0, 0, 0, 271, 272, 5,
		14, 0, 0, 272, 45, 1, 0, 0, 0, 273, 274, 3, 144, 72, 0, 274, 47, 1, 0,
		0, 0, 275, 276, 5, 15, 0, 0, 276, 286, 5, 16, 0, 0, 277, 279, 5, 15, 0,
		0, 278, 280, 3, 34, 17, 0, 279, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0,
		281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283,
		284, 5, 16, 0, 0, 284, 286, 1, 0, 0, 0, 285, 275, 1, 0, 0, 0, 285, 277,
		1, 0, 0, 0, 286, 49, 1, 0, 0, 0, 287, 291, 5, 4, 0, 0, 288, 290, 3, 52,
		26, 0, 289, 288, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0,
		291, 292, 1, 0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294,
		295, 5, 5, 0, 0, 295, 51, 1, 0, 0, 0, 296, 297, 3, 144, 72, 0, 297, 298,
		5, 8, 0, 0, 298, 299, 3, 34, 17, 0, 299, 53, 1, 0, 0, 0, 300, 301, 5, 17,
		0, 0, 301, 302, 3, 144, 72, 0, 302, 55, 1, 0, 0, 0, 303, 305, 5, 6, 0,
		0, 304, 306, 3, 58, 29, 0, 305, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0,
		307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309,
		310, 5, 7, 0, 0, 310, 57, 1, 0, 0, 0, 311, 312, 3, 54, 27, 0, 312, 313,
		5, 8, 0, 0, 313, 315, 3, 62, 31, 0, 314, 316, 3, 60, 30, 0, 315, 314, 1,
		0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 59, 1, 0, 0, 0, 317, 318, 5, 18, 0,
		0, 318, 319, 3, 34, 17, 0, 319, 61, 1, 0, 0, 0, 320, 322, 3, 64, 32, 0,
		321, 323, 5, 19, 0, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323,
		329, 1, 0, 0, 0, 324, 326, 3, 66, 33, 0, 325, 327, 5, 19, 0, 0, 326, 325,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 320, 1, 0,
		0, 0, 328, 324, 1, 0, 0, 0, 329, 63, 1, 0, 0, 0, 330, 331, 3, 144, 72,
		0, 331, 65, 1, 0, 0, 0, 332, 333, 5, 15, 0, 0, 333, 334, 3, 62, 31, 0,
		334, 335, 5, 16, 0, 0, 335, 67, 1, 0, 0, 0, 336, 338, 3, 70, 35, 0, 337,
		336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 340,
		1, 0, 0, 0, 340, 69, 1, 0, 0, 0, 341, 342, 5, 20, 0, 0, 342, 344, 3, 144,
		72, 0, 343, 345, 3, 18, 9, 0, 344, 343, 1, 0, 0, 0, 344, 345, 1, 0, 0,
		0, 345, 71, 1, 0, 0, 0, 346, 348, 3, 74, 37, 0, 347, 346, 1, 0, 0, 0, 348,
		349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 73, 1,
		0, 0, 0, 351, 355, 3, 80, 40, 0, 352, 355, 3, 88, 44, 0, 353, 355, 3, 134,
		67, 0, 354, 351, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 353, 1, 0, 0, 0,
		355, 75, 1, 0, 0, 0, 356, 358, 3, 78, 39, 0, 357, 356, 1, 0, 0, 0, 358,
		359, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 77, 1,
		0, 0, 0, 361, 364, 3, 84, 42, 0, 362, 364, 3, 90, 45, 0, 363, 361, 1, 0,
		0, 0, 363, 362, 1, 0, 0, 0, 364, 79, 1, 0, 0, 0, 365, 367, 5, 21, 0, 0,
		366, 368, 3, 68, 34, 0, 367, 366, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368,
		369, 1, 0, 0, 0, 369, 371, 5, 4, 0, 0, 370, 372, 3, 82, 41, 0, 371, 370,
		1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 5, 5, 0, 0, 376, 81, 1, 0, 0, 0,
		377, 378, 3, 10, 5, 0, 378, 379, 5, 8, 0, 0, 379, 380, 3, 64, 32, 0, 380,
		83, 1, 0, 0, 0, 381, 382, 5, 22, 0, 0, 382, 384, 5, 21, 0, 0, 383, 385,
		3, 68, 34, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386, 1,
		0, 0, 0, 386, 388, 5, 4, 0, 0, 387, 389, 3, 82, 41, 0, 388, 387, 1, 0,
		0, 0, 389, 390, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0,
		391, 392, 1, 0, 0, 0, 392, 393, 5, 5, 0, 0, 393, 398, 1, 0, 0, 0, 394,
		395, 5, 22, 0, 0, 395, 396, 5, 21, 0, 0, 396, 398, 3, 68, 34, 0, 397, 381,
		1, 0, 0, 0, 397, 394, 1, 0, 0, 0, 398, 85, 1, 0, 0, 0, 399, 400, 3, 42,
		21, 0, 400, 87, 1, 0, 0, 0, 401, 408, 3, 92, 46, 0, 402, 408, 3, 96, 48,
		0, 403, 408, 3, 110, 55, 0, 404, 408, 3, 114, 57, 0, 405, 408, 3, 120,
		60, 0, 406, 408, 3, 128, 64, 0, 407, 401, 1, 0, 0, 0, 407, 402, 1, 0, 0,
		0, 407, 403, 1, 0, 0, 0, 407, 404, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407,
		406, 1, 0, 0, 0, 408, 89, 1, 0, 0, 0, 409, 416, 3, 94, 47, 0, 410, 416,
		3, 108, 54, 0, 411, 416, 3, 112, 56, 0, 412, 416, 3, 118, 59, 0, 413, 416,
		3, 126, 63, 0, 414, 416, 3, 132, 66, 0, 415, 409, 1, 0, 0, 0, 415, 410,
		1, 0, 0, 0, 415, 411, 1, 0, 0, 0, 415, 412, 1, 0, 0, 0, 415, 413, 1, 0,
		0, 0, 415, 414, 1, 0, 0, 0, 416, 91, 1, 0, 0, 0, 417, 419, 3, 86, 43, 0,
		418, 417, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420,
		421, 5, 23, 0, 0, 421, 423, 3, 144, 72, 0, 422, 424, 3, 68, 34, 0, 423,
		422, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 93, 1, 0, 0, 0, 425, 426, 5,
		22, 0, 0, 426, 427, 5, 23, 0, 0, 427, 428, 3, 144, 72, 0, 428, 429, 3,
		68, 34, 0, 429, 95, 1, 0, 0, 0, 430, 432, 3, 86, 43, 0, 431, 430, 1, 0,
		0, 0, 431, 432, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 5, 24, 0, 0,
		434, 436, 3, 144, 72, 0, 435, 437, 3, 98, 49, 0, 436, 435, 1, 0, 0, 0,
		436, 437, 1, 0, 0, 0, 437, 439, 1, 0, 0, 0, 438, 440, 3, 68, 34, 0, 439,
		438, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440, 442, 1, 0, 0, 0, 441, 443,
		3, 100, 50, 0, 442, 441, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 97, 1,
		0, 0, 0, 444, 445, 6, 49, -1, 0, 445, 447, 5, 25, 0, 0, 446, 448, 5, 26,
		0, 0, 447, 446, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0,
		449, 450, 3, 64, 32, 0, 450, 456, 1, 0, 0, 0, 451, 452, 10, 1, 0, 0, 452,
		453, 5, 26, 0, 0, 453, 455, 3, 64, 32, 0, 454, 451, 1, 0, 0, 0, 455, 458,
		1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 99, 1, 0,
		0, 0, 458, 456, 1, 0, 0, 0, 459, 461, 5, 4, 0, 0, 460, 462, 3, 102, 51,
		0, 461, 460, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 463,
		464, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 466, 5, 5, 0, 0, 466, 101,
		1, 0, 0, 0, 467, 469, 3, 86, 43, 0, 468, 467, 1, 0, 0, 0, 468, 469, 1,
		0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 472, 3, 144, 72, 0, 471, 473, 3, 104,
		52, 0, 472, 471, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0,
		474, 475, 5, 8, 0, 0, 475, 477, 3, 62, 31, 0, 476, 478, 3, 68, 34, 0, 477,
		476, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 103, 1, 0, 0, 0, 479, 481,
		5, 6, 0, 0, 480, 482, 3, 106, 53, 0, 481, 480, 1, 0, 0, 0, 482, 483, 1,
		0, 0, 0, 483, 481, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485, 1, 0, 0,
		0, 485, 486, 5, 7, 0, 0, 486, 105, 1, 0, 0, 0, 487, 489, 3, 86, 43, 0,
		488, 487, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490,
		491, 3, 144, 72, 0, 491, 492, 5, 8, 0, 0, 492, 494, 3, 62, 31, 0, 493,
		495, 3, 60, 30, 0, 494, 493, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 497,
		1, 0, 0, 0, 496, 498, 3, 68, 34, 0, 497, 496, 1, 0, 0, 0, 497, 498, 1,
		0, 0, 0, 498, 107, 1, 0, 0, 0, 499, 500, 5, 22, 0, 0, 500, 501, 5, 24,
		0, 0, 501, 503, 3, 144, 72, 0, 502, 504, 3, 98, 49, 0, 503, 502, 1, 0,
		0, 0, 503, 504, 1, 0, 0, 0, 504, 506, 1, 0, 0, 0, 505, 507, 3, 68, 34,
		0, 506, 505, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508,
		509, 3, 100, 50, 0, 509, 524, 1, 0, 0, 0, 510, 511, 5, 22, 0, 0, 511, 512,
		5, 24, 0, 0, 512, 514, 3, 144, 72, 0, 513, 515, 3, 98, 49, 0, 514, 513,
		1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 517, 3, 68,
		34, 0, 517, 524, 1, 0, 0, 0, 518, 519, 5, 22, 0, 0, 519, 520, 5, 24, 0,
		0, 520, 521, 3, 144, 72, 0, 521, 522, 3, 98, 49, 0, 522, 524, 1, 0, 0,
		0, 523, 499, 1, 0, 0, 0, 523, 510, 1, 0, 0, 0, 523, 518, 1, 0, 0, 0, 524,
		109, 1, 0, 0, 0, 525, 527, 3, 86, 43, 0, 526, 525, 1, 0, 0, 0, 526, 527,
		1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 529, 5, 27, 0, 0, 529, 531, 3, 144,
		72, 0, 530, 532, 3, 98, 49, 0, 531, 530, 1, 0, 0, 0, 531, 532, 1, 0, 0,
		0, 532, 534, 1, 0, 0, 0, 533, 535, 3, 68, 34, 0, 534, 533, 1, 0, 0, 0,
		534, 535, 1, 0, 0, 0, 535, 537, 1, 0, 0, 0, 536, 538, 3, 100, 50, 0, 537,
		536, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538, 111, 1, 0, 0, 0, 539, 540,
		5, 22, 0, 0, 540, 541, 5, 27, 0, 0, 541, 543, 3, 144, 72, 0, 542, 544,
		3, 98, 49, 0, 543, 542, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 546, 1,
		0, 0, 0, 545, 547, 3, 68, 34, 0, 546, 545, 1, 0, 0, 0, 546, 547, 1, 0,
		0, 0, 547, 548, 1, 0, 0, 0, 548, 549, 3, 100, 50, 0, 549, 559, 1, 0, 0,
		0, 550, 551, 5, 22, 0, 0, 551, 552, 5, 27, 0, 0, 552, 554, 3, 144, 72,
		0, 553, 555, 3, 98, 49, 0, 554, 553, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0,
		555, 556, 1, 0, 0, 0, 556, 557, 3, 68, 34, 0, 557, 559, 1, 0, 0, 0, 558,
		539, 1, 0, 0, 0, 558, 550, 1, 0, 0, 0, 559, 113, 1, 0, 0, 0, 560, 562,
		3, 86, 43, 0, 561, 560, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 563, 1,
		0, 0, 0, 563, 564, 5, 28, 0, 0, 564, 566, 3, 144, 72, 0, 565, 567, 3, 68,
		34, 0, 566, 565, 1, 0, 0, 0, 566, 567, 1, 0, 0, 0, 567, 569, 1, 0, 0, 0,
		568, 570, 3, 116, 58, 0, 569, 568, 1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570,
		115, 1, 0, 0, 0, 571, 573, 5, 18, 0, 0, 572, 574, 5, 29, 0, 0, 573, 572,
		1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 575, 1, 0, 0, 0, 575, 580, 3, 64,
		32, 0, 576, 577, 5, 29, 0, 0, 577, 579, 3, 64, 32, 0, 578, 576, 1, 0, 0,
		0, 579, 582, 1, 0, 0, 0, 580, 578, 1, 0, 0, 0, 580, 581, 1, 0, 0, 0, 581,
		117, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 583, 584, 5, 22, 0, 0, 584, 585,
		5, 28, 0, 0, 585, 587, 3, 144, 72, 0, 586, 588, 3, 68, 34, 0, 587, 586,
		1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588, 589, 1, 0, 0, 0, 589, 590, 3, 116,
		58, 0, 590, 597, 1, 0, 0, 0, 591, 592, 5, 22, 0, 0, 592, 593, 5, 28, 0,
		0, 593, 594, 3, 144, 72, 0, 594, 595, 3, 68, 34, 0, 595, 597, 1, 0, 0,
		0, 596, 583, 1, 0, 0, 0, 596, 591, 1, 0, 0, 0, 597, 119, 1, 0, 0, 0, 598,
		600, 3, 86, 43, 0, 599, 598, 1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600, 601,
		1, 0, 0, 0, 601, 602, 5, 30, 0, 0, 602, 604, 3, 144, 72, 0, 603, 605, 3,
		68, 34, 0, 604, 603, 1, 0, 0, 0, 604, 605, 1, 0, 0, 0, 605, 607, 1, 0,
		0, 0, 606, 608, 3, 122, 61, 0, 607, 606, 1, 0, 0, 0, 607, 608, 1, 0, 0,
		0, 608, 121, 1, 0, 0, 0, 609, 611, 5, 4, 0, 0, 610, 612, 3, 124, 62, 0,
		611, 610, 1, 0, 0, 0, 612, 613, 1, 0, 0, 0, 613, 611, 1, 0, 0, 0, 613,
		614, 1, 0, 0, 0, 614, 615, 1, 0, 0, 0, 615, 616, 5, 5, 0, 0, 616, 123,
		1, 0, 0, 0, 617, 619, 3, 86, 43, 0, 618, 617, 1, 0, 0, 0, 618, 619, 1,
		0, 0, 0, 619, 620, 1, 0, 0, 0, 620, 622, 3, 46, 23, 0, 621, 623, 3, 68,
		34, 0, 622, 621, 1, 0, 0, 0, 622, 623, 1, 0, 0, 0, 623, 125, 1, 0, 0, 0,
		624, 625, 5, 22, 0, 0, 625, 626, 5, 30, 0, 0, 626, 628, 3, 144, 72, 0,
		627, 629, 3, 68, 34, 0, 628, 627, 1, 0, 0, 0, 628, 629, 1, 0, 0, 0, 629,
		630, 1, 0, 0, 0, 630, 631, 3, 122, 61, 0, 631, 638, 1, 0, 0, 0, 632, 633,
		5, 22, 0, 0, 633, 634, 5, 30, 0, 0, 634, 635, 3, 144, 72, 0, 635, 636,
		3, 68, 34, 0, 636, 638, 1, 0, 0, 0, 637, 624, 1, 0, 0, 0, 637, 632, 1,
		0, 0, 0, 638, 127, 1, 0, 0, 0, 639, 641, 3, 86, 43, 0, 640, 639, 1, 0,
		0, 0, 640, 641, 1, 0, 0, 0, 641, 642, 1, 0, 0, 0, 642, 643, 5, 31, 0, 0,
		643, 645, 3, 144, 72, 0, 644, 646, 3, 68, 34, 0, 645, 644, 1, 0, 0, 0,
		645, 646, 1, 0, 0, 0, 646, 648, 1, 0, 0, 0, 647, 649, 3, 130, 65, 0, 648,
		647, 1, 0, 0, 0, 648, 649, 1, 0, 0, 0, 649, 129, 1, 0, 0, 0, 650, 652,
		5, 4, 0, 0, 651, 653, 3, 106, 53, 0, 652, 651, 1, 0, 0, 0, 653, 654, 1,
		0, 0, 0, 654, 652, 1, 0, 0, 0, 654, 655, 1, 0, 0, 0, 655, 656, 1, 0, 0,
		0, 656, 657, 5, 5, 0, 0, 657, 131, 1, 0, 0, 0, 658, 659, 5, 22, 0, 0, 659,
		660, 5, 31, 0, 0, 660, 662, 3, 144, 72, 0, 661, 663, 3, 68, 34, 0, 662,
		661, 1, 0, 0, 0, 662, 663, 1, 0, 0, 0, 663, 664, 1, 0, 0, 0, 664, 665,
		3, 130, 65, 0, 665, 672, 1, 0, 0, 0, 666, 667, 5, 22, 0, 0, 667, 668, 5,
		31, 0, 0, 668, 669, 3, 144, 72, 0, 669, 670, 3, 68, 34, 0, 670, 672, 1,
		0, 0, 0, 671, 658, 1, 0, 0, 0, 671, 666, 1, 0, 0, 0, 672, 133, 1, 0, 0,
		0, 673, 675, 3, 86, 43, 0, 674, 673, 1, 0, 0, 0, 674, 675, 1, 0, 0, 0,
		675, 676, 1, 0, 0, 0, 676, 677, 5, 32, 0, 0, 677, 678, 5, 20, 0, 0, 678,
		680, 3, 144, 72, 0, 679, 681, 3, 104, 52, 0, 680, 679, 1, 0, 0, 0, 680,
		681, 1, 0, 0, 0, 681, 683, 1, 0, 0, 0, 682, 684, 5, 33, 0, 0, 683, 682,
		1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 685, 1, 0, 0, 0, 685, 686, 5, 11,
		0, 0, 686, 687, 3, 136, 68, 0, 687, 135, 1, 0, 0, 0, 688, 693, 3, 138,
		69, 0, 689, 690, 5, 29, 0, 0, 690, 692, 3, 138, 69, 0, 691, 689, 1, 0,
		0, 0, 692, 695, 1, 0, 0, 0, 693, 691, 1, 0, 0, 0, 693, 694, 1, 0, 0, 0,
		694, 137, 1, 0, 0, 0, 695, 693, 1, 0, 0, 0, 696, 699, 3, 140, 70, 0, 697,
		699, 3, 142, 71, 0, 698, 696, 1, 0, 0, 0, 698, 697, 1, 0, 0, 0, 699, 139,
		1, 0, 0, 0, 700, 701, 7, 3, 0, 0, 701, 141, 1, 0, 0, 0, 702, 703, 7, 4,
		0, 0, 703, 143, 1, 0, 0, 0, 704, 705, 5, 53, 0, 0, 705, 145, 1, 0, 0, 0,
		95, 149, 156, 161, 165, 169, 172, 175, 180, 188, 195, 198, 202, 205, 208,
		214, 228, 234, 245, 248, 261, 281, 285, 291, 307, 315, 322, 326, 328, 339,
		344, 349, 354, 359, 363, 367, 373, 384, 390, 397, 407, 415, 418, 423, 431,
		436, 439, 442, 447, 456, 463, 468, 472, 477, 483, 488, 494, 497, 503, 506,
		514, 523, 526, 531, 534, 537, 543, 546, 554, 558, 561, 566, 569, 573, 580,
		587, 596, 599, 604, 607, 613, 618, 622, 628, 637, 640, 645, 648, 654, 662,
		671, 674, 680, 683, 693, 698,
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

// GraphQLParserInit initializes any static state used to implement GraphQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGraphQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GraphQLParserInit() {
	staticData := &GraphQLParserStaticData
	staticData.once.Do(graphqlParserInit)
}

// NewGraphQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGraphQLParser(input antlr.TokenStream) *GraphQLParser {
	GraphQLParserInit()
	this := new(GraphQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GraphQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GraphQL.g4"

	return this
}

// GraphQLParser tokens.
const (
	GraphQLParserEOF          = antlr.TokenEOF
	GraphQLParserT__0         = 1
	GraphQLParserT__1         = 2
	GraphQLParserT__2         = 3
	GraphQLParserT__3         = 4
	GraphQLParserT__4         = 5
	GraphQLParserT__5         = 6
	GraphQLParserT__6         = 7
	GraphQLParserT__7         = 8
	GraphQLParserT__8         = 9
	GraphQLParserT__9         = 10
	GraphQLParserT__10        = 11
	GraphQLParserT__11        = 12
	GraphQLParserT__12        = 13
	GraphQLParserT__13        = 14
	GraphQLParserT__14        = 15
	GraphQLParserT__15        = 16
	GraphQLParserT__16        = 17
	GraphQLParserT__17        = 18
	GraphQLParserT__18        = 19
	GraphQLParserT__19        = 20
	GraphQLParserT__20        = 21
	GraphQLParserT__21        = 22
	GraphQLParserT__22        = 23
	GraphQLParserT__23        = 24
	GraphQLParserT__24        = 25
	GraphQLParserT__25        = 26
	GraphQLParserT__26        = 27
	GraphQLParserT__27        = 28
	GraphQLParserT__28        = 29
	GraphQLParserT__29        = 30
	GraphQLParserT__30        = 31
	GraphQLParserT__31        = 32
	GraphQLParserT__32        = 33
	GraphQLParserT__33        = 34
	GraphQLParserT__34        = 35
	GraphQLParserT__35        = 36
	GraphQLParserT__36        = 37
	GraphQLParserT__37        = 38
	GraphQLParserT__38        = 39
	GraphQLParserT__39        = 40
	GraphQLParserT__40        = 41
	GraphQLParserT__41        = 42
	GraphQLParserT__42        = 43
	GraphQLParserT__43        = 44
	GraphQLParserT__44        = 45
	GraphQLParserT__45        = 46
	GraphQLParserT__46        = 47
	GraphQLParserT__47        = 48
	GraphQLParserT__48        = 49
	GraphQLParserT__49        = 50
	GraphQLParserT__50        = 51
	GraphQLParserT__51        = 52
	GraphQLParserNAME         = 53
	GraphQLParserSTRING       = 54
	GraphQLParserBLOCK_STRING = 55
	GraphQLParserID           = 56
	GraphQLParserFLOAT        = 57
	GraphQLParserINT          = 58
	GraphQLParserPUNCTUATOR   = 59
	GraphQLParserWS           = 60
	GraphQLParserCOMMA        = 61
	GraphQLParserLineComment  = 62
	GraphQLParserUNICODE_BOM  = 63
	GraphQLParserUTF8_BOM     = 64
	GraphQLParserUTF16_BOM    = 65
	GraphQLParserUTF32_BOM    = 66
)

// GraphQLParser rules.
const (
	GraphQLParserRULE_document                    = 0
	GraphQLParserRULE_definition                  = 1
	GraphQLParserRULE_executableDocument          = 2
	GraphQLParserRULE_executableDefinition        = 3
	GraphQLParserRULE_operationDefinition         = 4
	GraphQLParserRULE_operationType               = 5
	GraphQLParserRULE_selectionSet                = 6
	GraphQLParserRULE_selection                   = 7
	GraphQLParserRULE_field                       = 8
	GraphQLParserRULE_arguments                   = 9
	GraphQLParserRULE_argument                    = 10
	GraphQLParserRULE_alias                       = 11
	GraphQLParserRULE_fragmentSpread              = 12
	GraphQLParserRULE_fragmentDefinition          = 13
	GraphQLParserRULE_fragmentName                = 14
	GraphQLParserRULE_typeCondition               = 15
	GraphQLParserRULE_inlineFragment              = 16
	GraphQLParserRULE_value                       = 17
	GraphQLParserRULE_intValue                    = 18
	GraphQLParserRULE_floatValue                  = 19
	GraphQLParserRULE_booleanValue                = 20
	GraphQLParserRULE_stringValue                 = 21
	GraphQLParserRULE_nullValue                   = 22
	GraphQLParserRULE_enumValue                   = 23
	GraphQLParserRULE_listValue                   = 24
	GraphQLParserRULE_objectValue                 = 25
	GraphQLParserRULE_objectField                 = 26
	GraphQLParserRULE_variable                    = 27
	GraphQLParserRULE_variableDefinitions         = 28
	GraphQLParserRULE_variableDefinition          = 29
	GraphQLParserRULE_defaultValue                = 30
	GraphQLParserRULE_type_                       = 31
	GraphQLParserRULE_namedType                   = 32
	GraphQLParserRULE_listType                    = 33
	GraphQLParserRULE_directives                  = 34
	GraphQLParserRULE_directive                   = 35
	GraphQLParserRULE_typeSystemDocument          = 36
	GraphQLParserRULE_typeSystemDefinition        = 37
	GraphQLParserRULE_typeSystemExtensionDocument = 38
	GraphQLParserRULE_typeSystemExtension         = 39
	GraphQLParserRULE_schemaDefinition            = 40
	GraphQLParserRULE_rootOperationTypeDefinition = 41
	GraphQLParserRULE_schemaExtension             = 42
	GraphQLParserRULE_description                 = 43
	GraphQLParserRULE_typeDefinition              = 44
	GraphQLParserRULE_typeExtension               = 45
	GraphQLParserRULE_scalarTypeDefinition        = 46
	GraphQLParserRULE_scalarTypeExtension         = 47
	GraphQLParserRULE_objectTypeDefinition        = 48
	GraphQLParserRULE_implementsInterfaces        = 49
	GraphQLParserRULE_fieldsDefinition            = 50
	GraphQLParserRULE_fieldDefinition             = 51
	GraphQLParserRULE_argumentsDefinition         = 52
	GraphQLParserRULE_inputValueDefinition        = 53
	GraphQLParserRULE_objectTypeExtension         = 54
	GraphQLParserRULE_interfaceTypeDefinition     = 55
	GraphQLParserRULE_interfaceTypeExtension      = 56
	GraphQLParserRULE_unionTypeDefinition         = 57
	GraphQLParserRULE_unionMemberTypes            = 58
	GraphQLParserRULE_unionTypeExtension          = 59
	GraphQLParserRULE_enumTypeDefinition          = 60
	GraphQLParserRULE_enumValuesDefinition        = 61
	GraphQLParserRULE_enumValueDefinition         = 62
	GraphQLParserRULE_enumTypeExtension           = 63
	GraphQLParserRULE_inputObjectTypeDefinition   = 64
	GraphQLParserRULE_inputFieldsDefinition       = 65
	GraphQLParserRULE_inputObjectTypeExtension    = 66
	GraphQLParserRULE_directiveDefinition         = 67
	GraphQLParserRULE_directiveLocations          = 68
	GraphQLParserRULE_directiveLocation           = 69
	GraphQLParserRULE_executableDirectiveLocation = 70
	GraphQLParserRULE_typeSystemDirectiveLocation = 71
	GraphQLParserRULE_name                        = 72
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDefinition() []IDefinitionContext
	Definition(i int) IDefinitionContext

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEOF, 0)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
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

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54043203478750238) != 0) {
		{
			p.SetState(146)
			p.Definition()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.Match(GraphQLParserEOF)
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

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExecutableDocument() IExecutableDocumentContext
	TypeSystemDocument() ITypeSystemDocumentContext
	TypeSystemExtensionDocument() ITypeSystemExtensionDocumentContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) ExecutableDocument() IExecutableDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecutableDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecutableDocumentContext)
}

func (s *DefinitionContext) TypeSystemDocument() ITypeSystemDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSystemDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDocumentContext)
}

func (s *DefinitionContext) TypeSystemExtensionDocument() ITypeSystemExtensionDocumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSystemExtensionDocumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSystemExtensionDocumentContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLParserRULE_definition)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2, GraphQLParserT__3, GraphQLParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.ExecutableDocument()
		}

	case GraphQLParserT__20, GraphQLParserT__22, GraphQLParserT__23, GraphQLParserT__26, GraphQLParserT__27, GraphQLParserT__29, GraphQLParserT__30, GraphQLParserT__31, GraphQLParserSTRING, GraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.TypeSystemDocument()
		}

	case GraphQLParserT__21:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.TypeSystemExtensionDocument()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IExecutableDocumentContext is an interface to support dynamic dispatch.
type IExecutableDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExecutableDefinition() []IExecutableDefinitionContext
	ExecutableDefinition(i int) IExecutableDefinitionContext

	// IsExecutableDocumentContext differentiates from other interfaces.
	IsExecutableDocumentContext()
}

type ExecutableDocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDocumentContext() *ExecutableDocumentContext {
	var p = new(ExecutableDocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDocument
	return p
}

func InitEmptyExecutableDocumentContext(p *ExecutableDocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDocument
}

func (*ExecutableDocumentContext) IsExecutableDocumentContext() {}

func NewExecutableDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDocumentContext {
	var p = new(ExecutableDocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDocument

	return p
}

func (s *ExecutableDocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutableDocumentContext) AllExecutableDefinition() []IExecutableDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExecutableDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IExecutableDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExecutableDefinitionContext); ok {
			tst[i] = t.(IExecutableDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ExecutableDocumentContext) ExecutableDefinition(i int) IExecutableDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecutableDefinitionContext); ok {
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

	return t.(IExecutableDefinitionContext)
}

func (s *ExecutableDocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDocument(s)
	}
}

func (s *ExecutableDocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDocument(s)
	}
}

func (s *ExecutableDocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitExecutableDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ExecutableDocument() (localctx IExecutableDocumentContext) {
	localctx = NewExecutableDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLParserRULE_executableDocument)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(158)
				p.ExecutableDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
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

// IExecutableDefinitionContext is an interface to support dynamic dispatch.
type IExecutableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperationDefinition() IOperationDefinitionContext
	FragmentDefinition() IFragmentDefinitionContext

	// IsExecutableDefinitionContext differentiates from other interfaces.
	IsExecutableDefinitionContext()
}

type ExecutableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDefinitionContext() *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDefinition
	return p
}

func InitEmptyExecutableDefinitionContext(p *ExecutableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDefinition
}

func (*ExecutableDefinitionContext) IsExecutableDefinitionContext() {}

func NewExecutableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDefinition

	return p
}

func (s *ExecutableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutableDefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *ExecutableDefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFragmentDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *ExecutableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitExecutableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ExecutableDefinition() (localctx IExecutableDefinitionContext) {
	localctx = NewExecutableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLParserRULE_executableDefinition)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2, GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.OperationDefinition()
		}

	case GraphQLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.FragmentDefinition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperationType() IOperationTypeContext
	SelectionSet() ISelectionSetContext
	Name() INameContext
	VariableDefinitions() IVariableDefinitionsContext
	Directives() IDirectivesContext

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
	return p
}

func InitEmptyOperationDefinitionContext(p *OperationDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitOperationDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLParserRULE_operationDefinition)
	var _la int

	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.OperationType()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserNAME {
			{
				p.SetState(168)
				p.Name()
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__5 {
			{
				p.SetState(171)
				p.VariableDefinitions()
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(174)
				p.Directives()
			}

		}
		{
			p.SetState(177)
			p.SelectionSet()
		}

	case GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.SelectionSet()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
	return p
}

func InitEmptyOperationTypeContext(p *OperationTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (s *OperationTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitOperationType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLParserRULE_operationType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
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

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSelection() []ISelectionContext
	Selection(i int) ISelectionContext

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
	return p
}

func InitEmptySelectionSetContext(p *SelectionSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectionContext); ok {
			len++
		}
	}

	tst := make([]ISelectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectionContext); ok {
			tst[i] = t.(ISelectionContext)
			i++
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionContext); ok {
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

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (s *SelectionSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSelectionSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLParserRULE_selectionSet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__8 || _la == GraphQLParserNAME {
		{
			p.SetState(185)
			p.Selection()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(GraphQLParserT__4)
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

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	FragmentSpread() IFragmentSpreadContext
	InlineFragment() IInlineFragmentContext

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
	return p
}

func InitEmptySelectionContext(p *SelectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFragmentSpreadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineFragmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (s *SelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLParserRULE_selection)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.InlineFragment()
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Alias() IAliasContext
	Arguments() IArgumentsContext
	Directives() IDirectivesContext
	SelectionSet() ISelectionSetContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(197)
			p.Alias()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(200)
		p.Name()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(201)
			p.Arguments()
		}

	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(204)
			p.Directives()
		}

	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__3 {
		{
			p.SetState(207)
			p.SelectionSet()
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
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

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(GraphQLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserNAME {
		{
			p.SetState(211)
			p.Argument()
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(GraphQLParserT__6)
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

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Value() IValueContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Name()
	}
	{
		p.SetState(219)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Value()
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

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
	return p
}

func InitEmptyAliasContext(p *AliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLParserRULE_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Name()
	}
	{
		p.SetState(223)
		p.Match(GraphQLParserT__7)
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

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FragmentName() IFragmentNameContext
	Directives() IDirectivesContext

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
	return p
}

func InitEmptyFragmentSpreadContext(p *FragmentSpreadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFragmentNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLParserRULE_fragmentSpread)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(GraphQLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.FragmentName()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(227)
			p.Directives()
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

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FragmentName() IFragmentNameContext
	TypeCondition() ITypeConditionContext
	SelectionSet() ISelectionSetContext
	Directives() IDirectivesContext

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
	return p
}

func InitEmptyFragmentDefinitionContext(p *FragmentDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFragmentNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLParserRULE_fragmentDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(GraphQLParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.FragmentName()
	}
	{
		p.SetState(232)
		p.TypeCondition()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(233)
			p.Directives()
		}

	}
	{
		p.SetState(236)
		p.SelectionSet()
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

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
	return p
}

func InitEmptyFragmentNameContext(p *FragmentNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (s *FragmentNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLParserRULE_fragmentName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Name()
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

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamedType() INamedTypeContext

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
	return p
}

func InitEmptyTypeConditionContext(p *TypeConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) NamedType() INamedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (s *TypeConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLParserRULE_typeCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(GraphQLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.NamedType()
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

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SelectionSet() ISelectionSetContext
	TypeCondition() ITypeConditionContext
	Directives() IDirectivesContext

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
	return p
}

func InitEmptyInlineFragmentContext(p *InlineFragmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (s *InlineFragmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInlineFragment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraphQLParserRULE_inlineFragment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(GraphQLParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__10 {
		{
			p.SetState(244)
			p.TypeCondition()
		}

	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(247)
			p.Directives()
		}

	}
	{
		p.SetState(250)
		p.SelectionSet()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	IntValue() IIntValueContext
	FloatValue() IFloatValueContext
	StringValue() IStringValueContext
	BooleanValue() IBooleanValueContext
	NullValue() INullValueContext
	EnumValue() IEnumValueContext
	ListValue() IListValueContext
	ObjectValue() IObjectValueContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) IntValue() IIntValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntValueContext)
}

func (s *ValueContext) FloatValue() IFloatValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatValueContext)
}

func (s *ValueContext) StringValue() IStringValueContext {
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

func (s *ValueContext) BooleanValue() IBooleanValueContext {
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

func (s *ValueContext) NullValue() INullValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullValueContext)
}

func (s *ValueContext) EnumValue() IEnumValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) ObjectValue() IObjectValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraphQLParserRULE_value)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Variable()
		}

	case GraphQLParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.IntValue()
		}

	case GraphQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)
			p.FloatValue()
		}

	case GraphQLParserSTRING, GraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(255)
			p.StringValue()
		}

	case GraphQLParserT__11, GraphQLParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(256)
			p.BooleanValue()
		}

	case GraphQLParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(257)
			p.NullValue()
		}

	case GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(258)
			p.EnumValue()
		}

	case GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(259)
			p.ListValue()
		}

	case GraphQLParserT__3:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(260)
			p.ObjectValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IIntValueContext is an interface to support dynamic dispatch.
type IIntValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntValueContext differentiates from other interfaces.
	IsIntValueContext()
}

type IntValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntValueContext() *IntValueContext {
	var p = new(IntValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_intValue
	return p
}

func InitEmptyIntValueContext(p *IntValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_intValue
}

func (*IntValueContext) IsIntValueContext() {}

func NewIntValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntValueContext {
	var p = new(IntValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_intValue

	return p
}

func (s *IntValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntValueContext) INT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINT, 0)
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) IntValue() (localctx IIntValueContext) {
	localctx = NewIntValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraphQLParserRULE_intValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(GraphQLParserINT)
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

// IFloatValueContext is an interface to support dynamic dispatch.
type IFloatValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode

	// IsFloatValueContext differentiates from other interfaces.
	IsFloatValueContext()
}

type FloatValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValueContext() *FloatValueContext {
	var p = new(FloatValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_floatValue
	return p
}

func InitEmptyFloatValueContext(p *FloatValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_floatValue
}

func (*FloatValueContext) IsFloatValueContext() {}

func NewFloatValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValueContext {
	var p = new(FloatValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_floatValue

	return p
}

func (s *FloatValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFLOAT, 0)
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (s *FloatValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFloatValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FloatValue() (localctx IFloatValueContext) {
	localctx = NewFloatValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraphQLParserRULE_floatValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(GraphQLParserFLOAT)
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

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = GraphQLParserRULE_booleanValue
	return p
}

func InitEmptyBooleanValueContext(p *BooleanValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_booleanValue
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (s *BooleanValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitBooleanValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraphQLParserRULE_booleanValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserT__11 || _la == GraphQLParserT__12) {
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

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	BLOCK_STRING() antlr.TerminalNode

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
	p.RuleIndex = GraphQLParserRULE_stringValue
	return p
}

func InitEmptyStringValueContext(p *StringValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_stringValue
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSTRING, 0)
}

func (s *StringValueContext) BLOCK_STRING() antlr.TerminalNode {
	return s.GetToken(GraphQLParserBLOCK_STRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraphQLParserRULE_stringValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING) {
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

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_nullValue
	return p
}

func InitEmptyNullValueContext(p *NullValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_nullValue
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nullValue

	return p
}

func (s *NullValueContext) GetParser() antlr.Parser { return s.parser }
func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) NullValue() (localctx INullValueContext) {
	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraphQLParserRULE_nullValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(GraphQLParserT__13)
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

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValue
	return p
}

func InitEmptyEnumValueContext(p *EnumValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValue
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (s *EnumValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraphQLParserRULE_enumValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Name()
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

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_listValue
	return p
}

func InitEmptyListValueContext(p *ListValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_listValue
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ListValueContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListValue(s)
	}
}

func (s *ListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListValue(s)
	}
}

func (s *ListValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitListValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ListValue() (localctx IListValueContext) {
	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraphQLParserRULE_listValue)
	var _la int

	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(GraphQLParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(GraphQLParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(GraphQLParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&495395959010947088) != 0) {
			{
				p.SetState(278)
				p.Value()
			}

			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(283)
			p.Match(GraphQLParserT__15)
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

// IObjectValueContext is an interface to support dynamic dispatch.
type IObjectValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllObjectField() []IObjectFieldContext
	ObjectField(i int) IObjectFieldContext

	// IsObjectValueContext differentiates from other interfaces.
	IsObjectValueContext()
}

type ObjectValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueContext() *ObjectValueContext {
	var p = new(ObjectValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValue
	return p
}

func InitEmptyObjectValueContext(p *ObjectValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValue
}

func (*ObjectValueContext) IsObjectValueContext() {}

func NewObjectValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueContext {
	var p = new(ObjectValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectValue

	return p
}

func (s *ObjectValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueContext) AllObjectField() []IObjectFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectFieldContext); ok {
			len++
		}
	}

	tst := make([]IObjectFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectFieldContext); ok {
			tst[i] = t.(IObjectFieldContext)
			i++
		}
	}

	return tst
}

func (s *ObjectValueContext) ObjectField(i int) IObjectFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectFieldContext); ok {
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

	return t.(IObjectFieldContext)
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (s *ObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectValue() (localctx IObjectValueContext) {
	localctx = NewObjectValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraphQLParserRULE_objectValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserNAME {
		{
			p.SetState(288)
			p.ObjectField()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(294)
		p.Match(GraphQLParserT__4)
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

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Value() IValueContext

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectField
	return p
}

func InitEmptyObjectFieldContext(p *ObjectFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectField
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (s *ObjectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectField() (localctx IObjectFieldContext) {
	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraphQLParserRULE_objectField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Name()
	}
	{
		p.SetState(297)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Value()
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraphQLParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(GraphQLParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Name()
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

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDefinition() []IVariableDefinitionContext
	VariableDefinition(i int) IVariableDefinitionContext

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
	return p
}

func InitEmptyVariableDefinitionsContext(p *VariableDefinitionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IVariableDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDefinitionContext); ok {
			tst[i] = t.(IVariableDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionContext); ok {
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

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariableDefinitions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraphQLParserRULE_variableDefinitions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(GraphQLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__16 {
		{
			p.SetState(304)
			p.VariableDefinition()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(309)
		p.Match(GraphQLParserT__6)
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

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	Type_() IType_Context
	DefaultValue() IDefaultValueContext

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
	return p
}

func InitEmptyVariableDefinitionContext(p *VariableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraphQLParserRULE_variableDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Variable()
	}
	{
		p.SetState(312)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Type_()
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(314)
			p.DefaultValue()
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

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
	return p
}

func InitEmptyDefaultValueContext(p *DefaultValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (s *DefaultValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDefaultValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraphQLParserRULE_defaultValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(GraphQLParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Value()
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

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamedType() INamedTypeContext
	ListType() IListTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) NamedType() INamedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *Type_Context) ListType() IListTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraphQLParserRULE_type_)
	var _la int

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.NamedType()
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__18 {
			{
				p.SetState(321)
				p.Match(GraphQLParserT__18)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.ListType()
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__18 {
			{
				p.SetState(325)
				p.Match(GraphQLParserT__18)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// INamedTypeContext is an interface to support dynamic dispatch.
type INamedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsNamedTypeContext differentiates from other interfaces.
	IsNamedTypeContext()
}

type NamedTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedTypeContext() *NamedTypeContext {
	var p = new(NamedTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_namedType
	return p
}

func InitEmptyNamedTypeContext(p *NamedTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_namedType
}

func (*NamedTypeContext) IsNamedTypeContext() {}

func NewNamedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeContext {
	var p = new(NamedTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_namedType

	return p
}

func (s *NamedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (s *NamedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitNamedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) NamedType() (localctx INamedTypeContext) {
	localctx = NewNamedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraphQLParserRULE_namedType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Name()
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

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
	return p
}

func InitEmptyListTypeContext(p *ListTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (s *ListTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitListType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraphQLParserRULE_listType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(GraphQLParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Type_()
	}
	{
		p.SetState(334)
		p.Match(GraphQLParserT__15)
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

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDirective() []IDirectiveContext
	Directive(i int) IDirectiveContext

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
	return p
}

func InitEmptyDirectivesContext(p *DirectivesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirectiveContext); ok {
			tst[i] = t.(IDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
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

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (s *DirectivesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectives(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraphQLParserRULE_directives)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__19 {
		{
			p.SetState(336)
			p.Directive()
		}

		p.SetState(339)
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

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Arguments() IArgumentsContext

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
	return p
}

func InitEmptyDirectiveContext(p *DirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (s *DirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraphQLParserRULE_directive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(GraphQLParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Name()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(343)
			p.Arguments()
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

// ITypeSystemDocumentContext is an interface to support dynamic dispatch.
type ITypeSystemDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeSystemDefinition() []ITypeSystemDefinitionContext
	TypeSystemDefinition(i int) ITypeSystemDefinitionContext

	// IsTypeSystemDocumentContext differentiates from other interfaces.
	IsTypeSystemDocumentContext()
}

type TypeSystemDocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDocumentContext() *TypeSystemDocumentContext {
	var p = new(TypeSystemDocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDocument
	return p
}

func InitEmptyTypeSystemDocumentContext(p *TypeSystemDocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDocument
}

func (*TypeSystemDocumentContext) IsTypeSystemDocumentContext() {}

func NewTypeSystemDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDocumentContext {
	var p = new(TypeSystemDocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDocument

	return p
}

func (s *TypeSystemDocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDocumentContext) AllTypeSystemDefinition() []ITypeSystemDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSystemDefinitionContext); ok {
			len++
		}
	}

	tst := make([]ITypeSystemDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSystemDefinitionContext); ok {
			tst[i] = t.(ITypeSystemDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *TypeSystemDocumentContext) TypeSystemDefinition(i int) ITypeSystemDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSystemDefinitionContext); ok {
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

	return t.(ITypeSystemDefinitionContext)
}

func (s *TypeSystemDocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDocument(s)
	}
}

func (s *TypeSystemDocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDocument(s)
	}
}

func (s *TypeSystemDocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemDocument() (localctx ITypeSystemDocumentContext) {
	localctx = NewTypeSystemDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraphQLParserRULE_typeSystemDocument)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(346)
				p.TypeSystemDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
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

// ITypeSystemDefinitionContext is an interface to support dynamic dispatch.
type ITypeSystemDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SchemaDefinition() ISchemaDefinitionContext
	TypeDefinition() ITypeDefinitionContext
	DirectiveDefinition() IDirectiveDefinitionContext

	// IsTypeSystemDefinitionContext differentiates from other interfaces.
	IsTypeSystemDefinitionContext()
}

type TypeSystemDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDefinitionContext() *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition
	return p
}

func InitEmptyTypeSystemDefinitionContext(p *TypeSystemDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition
}

func (*TypeSystemDefinitionContext) IsTypeSystemDefinitionContext() {}

func NewTypeSystemDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition

	return p
}

func (s *TypeSystemDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDefinitionContext) SchemaDefinition() ISchemaDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeSystemDefinitionContext) DirectiveDefinition() IDirectiveDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveDefinitionContext)
}

func (s *TypeSystemDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemDefinition() (localctx ITypeSystemDefinitionContext) {
	localctx = NewTypeSystemDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraphQLParserRULE_typeSystemDefinition)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(353)
			p.DirectiveDefinition()
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

// ITypeSystemExtensionDocumentContext is an interface to support dynamic dispatch.
type ITypeSystemExtensionDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeSystemExtension() []ITypeSystemExtensionContext
	TypeSystemExtension(i int) ITypeSystemExtensionContext

	// IsTypeSystemExtensionDocumentContext differentiates from other interfaces.
	IsTypeSystemExtensionDocumentContext()
}

type TypeSystemExtensionDocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemExtensionDocumentContext() *TypeSystemExtensionDocumentContext {
	var p = new(TypeSystemExtensionDocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemExtensionDocument
	return p
}

func InitEmptyTypeSystemExtensionDocumentContext(p *TypeSystemExtensionDocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemExtensionDocument
}

func (*TypeSystemExtensionDocumentContext) IsTypeSystemExtensionDocumentContext() {}

func NewTypeSystemExtensionDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemExtensionDocumentContext {
	var p = new(TypeSystemExtensionDocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemExtensionDocument

	return p
}

func (s *TypeSystemExtensionDocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemExtensionDocumentContext) AllTypeSystemExtension() []ITypeSystemExtensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSystemExtensionContext); ok {
			len++
		}
	}

	tst := make([]ITypeSystemExtensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSystemExtensionContext); ok {
			tst[i] = t.(ITypeSystemExtensionContext)
			i++
		}
	}

	return tst
}

func (s *TypeSystemExtensionDocumentContext) TypeSystemExtension(i int) ITypeSystemExtensionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSystemExtensionContext); ok {
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

	return t.(ITypeSystemExtensionContext)
}

func (s *TypeSystemExtensionDocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemExtensionDocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemExtensionDocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemExtensionDocument(s)
	}
}

func (s *TypeSystemExtensionDocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemExtensionDocument(s)
	}
}

func (s *TypeSystemExtensionDocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemExtensionDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemExtensionDocument() (localctx ITypeSystemExtensionDocumentContext) {
	localctx = NewTypeSystemExtensionDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraphQLParserRULE_typeSystemExtensionDocument)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(356)
				p.TypeSystemExtension()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
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

// ITypeSystemExtensionContext is an interface to support dynamic dispatch.
type ITypeSystemExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SchemaExtension() ISchemaExtensionContext
	TypeExtension() ITypeExtensionContext

	// IsTypeSystemExtensionContext differentiates from other interfaces.
	IsTypeSystemExtensionContext()
}

type TypeSystemExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemExtensionContext() *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemExtension
	return p
}

func InitEmptyTypeSystemExtensionContext(p *TypeSystemExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemExtension
}

func (*TypeSystemExtensionContext) IsTypeSystemExtensionContext() {}

func NewTypeSystemExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemExtension

	return p
}

func (s *TypeSystemExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemExtensionContext) SchemaExtension() ISchemaExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISchemaExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISchemaExtensionContext)
}

func (s *TypeSystemExtensionContext) TypeExtension() ITypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExtensionContext)
}

func (s *TypeSystemExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemExtension(s)
	}
}

func (s *TypeSystemExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemExtension(s)
	}
}

func (s *TypeSystemExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemExtension() (localctx ITypeSystemExtensionContext) {
	localctx = NewTypeSystemExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraphQLParserRULE_typeSystemExtension)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.SchemaExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.TypeExtension()
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

// ISchemaDefinitionContext is an interface to support dynamic dispatch.
type ISchemaDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directives() IDirectivesContext
	AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext
	RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext

	// IsSchemaDefinitionContext differentiates from other interfaces.
	IsSchemaDefinitionContext()
}

type SchemaDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaDefinitionContext() *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaDefinition
	return p
}

func InitEmptySchemaDefinitionContext(p *SchemaDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaDefinition
}

func (*SchemaDefinitionContext) IsSchemaDefinitionContext() {}

func NewSchemaDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaDefinition

	return p
}

func (s *SchemaDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaDefinitionContext) AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IRootOperationTypeDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
			tst[i] = t.(IRootOperationTypeDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *SchemaDefinitionContext) RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
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

	return t.(IRootOperationTypeDefinitionContext)
}

func (s *SchemaDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSchemaDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) SchemaDefinition() (localctx ISchemaDefinitionContext) {
	localctx = NewSchemaDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraphQLParserRULE_schemaDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(GraphQLParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(366)
			p.Directives()
		}

	}
	{
		p.SetState(369)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
		{
			p.SetState(370)
			p.RootOperationTypeDefinition()
		}

		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(375)
		p.Match(GraphQLParserT__4)
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

// IRootOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IRootOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperationType() IOperationTypeContext
	NamedType() INamedTypeContext

	// IsRootOperationTypeDefinitionContext differentiates from other interfaces.
	IsRootOperationTypeDefinitionContext()
}

type RootOperationTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootOperationTypeDefinitionContext() *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_rootOperationTypeDefinition
	return p
}

func InitEmptyRootOperationTypeDefinitionContext(p *RootOperationTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_rootOperationTypeDefinition
}

func (*RootOperationTypeDefinitionContext) IsRootOperationTypeDefinitionContext() {}

func NewRootOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_rootOperationTypeDefinition

	return p
}

func (s *RootOperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RootOperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *RootOperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *RootOperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootOperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootOperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterRootOperationTypeDefinition(s)
	}
}

func (s *RootOperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitRootOperationTypeDefinition(s)
	}
}

func (s *RootOperationTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitRootOperationTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) RootOperationTypeDefinition() (localctx IRootOperationTypeDefinitionContext) {
	localctx = NewRootOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraphQLParserRULE_rootOperationTypeDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.OperationType()
	}
	{
		p.SetState(378)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.NamedType()
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

// ISchemaExtensionContext is an interface to support dynamic dispatch.
type ISchemaExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directives() IDirectivesContext
	AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext
	RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext

	// IsSchemaExtensionContext differentiates from other interfaces.
	IsSchemaExtensionContext()
}

type SchemaExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaExtensionContext() *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaExtension
	return p
}

func InitEmptySchemaExtensionContext(p *SchemaExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaExtension
}

func (*SchemaExtensionContext) IsSchemaExtensionContext() {}

func NewSchemaExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaExtension

	return p
}

func (s *SchemaExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaExtensionContext) AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IRootOperationTypeDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
			tst[i] = t.(IRootOperationTypeDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *SchemaExtensionContext) RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootOperationTypeDefinitionContext); ok {
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

	return t.(IRootOperationTypeDefinitionContext)
}

func (s *SchemaExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaExtension(s)
	}
}

func (s *SchemaExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaExtension(s)
	}
}

func (s *SchemaExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSchemaExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) SchemaExtension() (localctx ISchemaExtensionContext) {
	localctx = NewSchemaExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraphQLParserRULE_schemaExtension)
	var _la int

	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.Match(GraphQLParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(383)
				p.Directives()
			}

		}
		{
			p.SetState(386)
			p.Match(GraphQLParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
			{
				p.SetState(387)
				p.RootOperationTypeDefinition()
			}

			p.SetState(390)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(392)
			p.Match(GraphQLParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(GraphQLParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Directives()
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

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringValue() IStringValueContext

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_description
	return p
}

func InitEmptyDescriptionContext(p *DescriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_description
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) StringValue() IStringValueContext {
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

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraphQLParserRULE_description)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.StringValue()
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

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarTypeDefinition() IScalarTypeDefinitionContext
	ObjectTypeDefinition() IObjectTypeDefinitionContext
	InterfaceTypeDefinition() IInterfaceTypeDefinitionContext
	UnionTypeDefinition() IUnionTypeDefinitionContext
	EnumTypeDefinition() IEnumTypeDefinitionContext
	InputObjectTypeDefinition() IInputObjectTypeDefinitionContext

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeDefinition
	return p
}

func InitEmptyTypeDefinitionContext(p *TypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeDefinition
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) ScalarTypeDefinition() IScalarTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarTypeDefinitionContext)
}

func (s *TypeDefinitionContext) ObjectTypeDefinition() IObjectTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InterfaceTypeDefinition() IInterfaceTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeDefinitionContext)
}

func (s *TypeDefinitionContext) UnionTypeDefinition() IUnionTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionTypeDefinitionContext)
}

func (s *TypeDefinitionContext) EnumTypeDefinition() IEnumTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InputObjectTypeDefinition() IInputObjectTypeDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputObjectTypeDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraphQLParserRULE_typeDefinition)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.ObjectTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
			p.InterfaceTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(404)
			p.UnionTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(405)
			p.EnumTypeDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(406)
			p.InputObjectTypeDefinition()
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

// ITypeExtensionContext is an interface to support dynamic dispatch.
type ITypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ScalarTypeExtension() IScalarTypeExtensionContext
	ObjectTypeExtension() IObjectTypeExtensionContext
	InterfaceTypeExtension() IInterfaceTypeExtensionContext
	UnionTypeExtension() IUnionTypeExtensionContext
	EnumTypeExtension() IEnumTypeExtensionContext
	InputObjectTypeExtension() IInputObjectTypeExtensionContext

	// IsTypeExtensionContext differentiates from other interfaces.
	IsTypeExtensionContext()
}

type TypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExtensionContext() *TypeExtensionContext {
	var p = new(TypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeExtension
	return p
}

func InitEmptyTypeExtensionContext(p *TypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeExtension
}

func (*TypeExtensionContext) IsTypeExtensionContext() {}

func NewTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExtensionContext {
	var p = new(TypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeExtension

	return p
}

func (s *TypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExtensionContext) ScalarTypeExtension() IScalarTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarTypeExtensionContext)
}

func (s *TypeExtensionContext) ObjectTypeExtension() IObjectTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) InterfaceTypeExtension() IInterfaceTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeExtensionContext)
}

func (s *TypeExtensionContext) UnionTypeExtension() IUnionTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionTypeExtensionContext)
}

func (s *TypeExtensionContext) EnumTypeExtension() IEnumTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeExtensionContext)
}

func (s *TypeExtensionContext) InputObjectTypeExtension() IInputObjectTypeExtensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputObjectTypeExtensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeExtension(s)
	}
}

func (s *TypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeExtension(s)
	}
}

func (s *TypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeExtension() (localctx ITypeExtensionContext) {
	localctx = NewTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraphQLParserRULE_typeExtension)
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.ScalarTypeExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.ObjectTypeExtension()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.InterfaceTypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.UnionTypeExtension()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.EnumTypeExtension()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(414)
			p.InputObjectTypeExtension()
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

// IScalarTypeDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	Directives() IDirectivesContext

	// IsScalarTypeDefinitionContext differentiates from other interfaces.
	IsScalarTypeDefinitionContext()
}

type ScalarTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeDefinitionContext() *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition
	return p
}

func InitEmptyScalarTypeDefinitionContext(p *ScalarTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition
}

func (*ScalarTypeDefinitionContext) IsScalarTypeDefinitionContext() {}

func NewScalarTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition

	return p
}

func (s *ScalarTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ScalarTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitScalarTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ScalarTypeDefinition() (localctx IScalarTypeDefinitionContext) {
	localctx = NewScalarTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GraphQLParserRULE_scalarTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(417)
			p.Description()
		}

	}
	{
		p.SetState(420)
		p.Match(GraphQLParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.Name()
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(422)
			p.Directives()
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

// IScalarTypeExtensionContext is an interface to support dynamic dispatch.
type IScalarTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Directives() IDirectivesContext

	// IsScalarTypeExtensionContext differentiates from other interfaces.
	IsScalarTypeExtensionContext()
}

type ScalarTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeExtensionContext() *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension
	return p
}

func InitEmptyScalarTypeExtensionContext(p *ScalarTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension
}

func (*ScalarTypeExtensionContext) IsScalarTypeExtensionContext() {}

func NewScalarTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension

	return p
}

func (s *ScalarTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitScalarTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ScalarTypeExtension() (localctx IScalarTypeExtensionContext) {
	localctx = NewScalarTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GraphQLParserRULE_scalarTypeExtension)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(GraphQLParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(GraphQLParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)
		p.Name()
	}
	{
		p.SetState(428)
		p.Directives()
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

// IObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	ImplementsInterfaces() IImplementsInterfacesContext
	Directives() IDirectivesContext
	FieldsDefinition() IFieldsDefinitionContext

	// IsObjectTypeDefinitionContext differentiates from other interfaces.
	IsObjectTypeDefinitionContext()
}

type ObjectTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeDefinitionContext() *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition
	return p
}

func InitEmptyObjectTypeDefinitionContext(p *ObjectTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition
}

func (*ObjectTypeDefinitionContext) IsObjectTypeDefinitionContext() {}

func NewObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition

	return p
}

func (s *ObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ObjectTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementsInterfacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectTypeDefinition() (localctx IObjectTypeDefinitionContext) {
	localctx = NewObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GraphQLParserRULE_objectTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(430)
			p.Description()
		}

	}
	{
		p.SetState(433)
		p.Match(GraphQLParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.Name()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__24 {
		{
			p.SetState(435)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(438)
			p.Directives()
		}

	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(441)
			p.FieldsDefinition()
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

// IImplementsInterfacesContext is an interface to support dynamic dispatch.
type IImplementsInterfacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NamedType() INamedTypeContext
	ImplementsInterfaces() IImplementsInterfacesContext

	// IsImplementsInterfacesContext differentiates from other interfaces.
	IsImplementsInterfacesContext()
}

type ImplementsInterfacesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementsInterfacesContext() *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces
	return p
}

func InitEmptyImplementsInterfacesContext(p *ImplementsInterfacesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces
}

func (*ImplementsInterfacesContext) IsImplementsInterfacesContext() {}

func NewImplementsInterfacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces

	return p
}

func (s *ImplementsInterfacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementsInterfacesContext) NamedType() INamedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *ImplementsInterfacesContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementsInterfacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ImplementsInterfacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementsInterfacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementsInterfacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitImplementsInterfaces(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ImplementsInterfaces() (localctx IImplementsInterfacesContext) {
	return p.implementsInterfaces(0)
}

func (p *GraphQLParser) implementsInterfaces(_p int) (localctx IImplementsInterfacesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewImplementsInterfacesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImplementsInterfacesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, GraphQLParserRULE_implementsInterfaces, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Match(GraphQLParserT__24)
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

	if _la == GraphQLParserT__25 {
		{
			p.SetState(446)
			p.Match(GraphQLParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(449)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImplementsInterfacesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_implementsInterfaces)
			p.SetState(451)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(452)
				p.Match(GraphQLParserT__25)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(453)
				p.NamedType()
			}

		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldsDefinitionContext is an interface to support dynamic dispatch.
type IFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldDefinition() []IFieldDefinitionContext
	FieldDefinition(i int) IFieldDefinitionContext

	// IsFieldsDefinitionContext differentiates from other interfaces.
	IsFieldsDefinitionContext()
}

type FieldsDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDefinitionContext() *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition
	return p
}

func InitEmptyFieldsDefinitionContext(p *FieldsDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition
}

func (*FieldsDefinitionContext) IsFieldsDefinitionContext() {}

func NewFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition

	return p
}

func (s *FieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefinitionContext); ok {
			tst[i] = t.(IFieldDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *FieldsDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefinitionContext); ok {
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

	return t.(IFieldDefinitionContext)
}

func (s *FieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFieldsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FieldsDefinition() (localctx IFieldsDefinitionContext) {
	localctx = NewFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, GraphQLParserRULE_fieldsDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63050394783186944) != 0) {
		{
			p.SetState(460)
			p.FieldDefinition()
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(465)
		p.Match(GraphQLParserT__4)
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

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Type_() IType_Context
	Description() IDescriptionContext
	ArgumentsDefinition() IArgumentsDefinitionContext
	Directives() IDirectivesContext

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldDefinition
	return p
}

func InitEmptyFieldDefinitionContext(p *FieldDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldDefinition
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldDefinitionContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *FieldDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, GraphQLParserRULE_fieldDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(467)
			p.Description()
		}

	}
	{
		p.SetState(470)
		p.Name()
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(471)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(474)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(475)
		p.Type_()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(476)
			p.Directives()
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

// IArgumentsDefinitionContext is an interface to support dynamic dispatch.
type IArgumentsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInputValueDefinition() []IInputValueDefinitionContext
	InputValueDefinition(i int) IInputValueDefinitionContext

	// IsArgumentsDefinitionContext differentiates from other interfaces.
	IsArgumentsDefinitionContext()
}

type ArgumentsDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsDefinitionContext() *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition
	return p
}

func InitEmptyArgumentsDefinitionContext(p *ArgumentsDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition
}

func (*ArgumentsDefinitionContext) IsArgumentsDefinitionContext() {}

func NewArgumentsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition

	return p
}

func (s *ArgumentsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputValueDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IInputValueDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputValueDefinitionContext); ok {
			tst[i] = t.(IInputValueDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputValueDefinitionContext); ok {
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

	return t.(IInputValueDefinitionContext)
}

func (s *ArgumentsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArgumentsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ArgumentsDefinition() (localctx IArgumentsDefinitionContext) {
	localctx = NewArgumentsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, GraphQLParserRULE_argumentsDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.Match(GraphQLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63050394783186944) != 0) {
		{
			p.SetState(480)
			p.InputValueDefinition()
		}

		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(485)
		p.Match(GraphQLParserT__6)
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

// IInputValueDefinitionContext is an interface to support dynamic dispatch.
type IInputValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Type_() IType_Context
	Description() IDescriptionContext
	DefaultValue() IDefaultValueContext
	Directives() IDirectivesContext

	// IsInputValueDefinitionContext differentiates from other interfaces.
	IsInputValueDefinitionContext()
}

type InputValueDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputValueDefinitionContext() *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition
	return p
}

func InitEmptyInputValueDefinitionContext(p *InputValueDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition
}

func (*InputValueDefinitionContext) IsInputValueDefinitionContext() {}

func NewInputValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition

	return p
}

func (s *InputValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputValueDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputValueDefinitionContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InputValueDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputValueDefinitionContext) DefaultValue() IDefaultValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *InputValueDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputValueDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputValueDefinition() (localctx IInputValueDefinitionContext) {
	localctx = NewInputValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, GraphQLParserRULE_inputValueDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(487)
			p.Description()
		}

	}
	{
		p.SetState(490)
		p.Name()
	}
	{
		p.SetState(491)
		p.Match(GraphQLParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)
		p.Type_()
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(493)
			p.DefaultValue()
		}

	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(496)
			p.Directives()
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

// IObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	FieldsDefinition() IFieldsDefinitionContext
	ImplementsInterfaces() IImplementsInterfacesContext
	Directives() IDirectivesContext

	// IsObjectTypeExtensionContext differentiates from other interfaces.
	IsObjectTypeExtensionContext()
}

type ObjectTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeExtensionContext() *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension
	return p
}

func InitEmptyObjectTypeExtensionContext(p *ObjectTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension
}

func (*ObjectTypeExtensionContext) IsObjectTypeExtensionContext() {}

func NewObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension

	return p
}

func (s *ObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeExtensionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementsInterfacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectTypeExtension() (localctx IObjectTypeExtensionContext) {
	localctx = NewObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, GraphQLParserRULE_objectTypeExtension)
	var _la int

	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.Match(GraphQLParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)
			p.Name()
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(502)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(505)
				p.Directives()
			}

		}
		{
			p.SetState(508)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.Match(GraphQLParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.Name()
		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(513)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(516)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(518)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Match(GraphQLParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.Name()
		}
		{
			p.SetState(521)
			p.implementsInterfaces(0)
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

// IInterfaceTypeDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	ImplementsInterfaces() IImplementsInterfacesContext
	Directives() IDirectivesContext
	FieldsDefinition() IFieldsDefinitionContext

	// IsInterfaceTypeDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeDefinitionContext()
}

type InterfaceTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeDefinitionContext() *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition
	return p
}

func InitEmptyInterfaceTypeDefinitionContext(p *InterfaceTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition
}

func (*InterfaceTypeDefinitionContext) IsInterfaceTypeDefinitionContext() {}

func NewInterfaceTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition

	return p
}

func (s *InterfaceTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InterfaceTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementsInterfacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *InterfaceTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInterfaceTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InterfaceTypeDefinition() (localctx IInterfaceTypeDefinitionContext) {
	localctx = NewInterfaceTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, GraphQLParserRULE_interfaceTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(525)
			p.Description()
		}

	}
	{
		p.SetState(528)
		p.Match(GraphQLParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
		p.Name()
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__24 {
		{
			p.SetState(530)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(533)
			p.Directives()
		}

	}
	p.SetState(537)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(536)
			p.FieldsDefinition()
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

// IInterfaceTypeExtensionContext is an interface to support dynamic dispatch.
type IInterfaceTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	FieldsDefinition() IFieldsDefinitionContext
	ImplementsInterfaces() IImplementsInterfacesContext
	Directives() IDirectivesContext

	// IsInterfaceTypeExtensionContext differentiates from other interfaces.
	IsInterfaceTypeExtensionContext()
}

type InterfaceTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeExtensionContext() *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension
	return p
}

func InitEmptyInterfaceTypeExtensionContext(p *InterfaceTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension
}

func (*InterfaceTypeExtensionContext) IsInterfaceTypeExtensionContext() {}

func NewInterfaceTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension

	return p
}

func (s *InterfaceTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeExtensionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImplementsInterfacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *InterfaceTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInterfaceTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InterfaceTypeExtension() (localctx IInterfaceTypeExtensionContext) {
	localctx = NewInterfaceTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, GraphQLParserRULE_interfaceTypeExtension)
	var _la int

	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Match(GraphQLParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)
			p.Name()
		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(542)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(545)
				p.Directives()
			}

		}
		{
			p.SetState(548)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(550)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(GraphQLParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Name()
		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__24 {
			{
				p.SetState(553)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(556)
			p.Directives()
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

// IUnionTypeDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	Directives() IDirectivesContext
	UnionMemberTypes() IUnionMemberTypesContext

	// IsUnionTypeDefinitionContext differentiates from other interfaces.
	IsUnionTypeDefinitionContext()
}

type UnionTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeDefinitionContext() *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition
	return p
}

func InitEmptyUnionTypeDefinitionContext(p *UnionTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition
}

func (*UnionTypeDefinitionContext) IsUnionTypeDefinitionContext() {}

func NewUnionTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition

	return p
}

func (s *UnionTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *UnionTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeDefinitionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionMemberTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionTypeDefinition() (localctx IUnionTypeDefinitionContext) {
	localctx = NewUnionTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GraphQLParserRULE_unionTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(560)
			p.Description()
		}

	}
	{
		p.SetState(563)
		p.Match(GraphQLParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(564)
		p.Name()
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(565)
			p.Directives()
		}

	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__17 {
		{
			p.SetState(568)
			p.UnionMemberTypes()
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

// IUnionMemberTypesContext is an interface to support dynamic dispatch.
type IUnionMemberTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNamedType() []INamedTypeContext
	NamedType(i int) INamedTypeContext

	// IsUnionMemberTypesContext differentiates from other interfaces.
	IsUnionMemberTypesContext()
}

type UnionMemberTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMemberTypesContext() *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes
	return p
}

func InitEmptyUnionMemberTypesContext(p *UnionMemberTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes
}

func (*UnionMemberTypesContext) IsUnionMemberTypesContext() {}

func NewUnionMemberTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes

	return p
}

func (s *UnionMemberTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMemberTypesContext) AllNamedType() []INamedTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamedTypeContext); ok {
			len++
		}
	}

	tst := make([]INamedTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamedTypeContext); ok {
			tst[i] = t.(INamedTypeContext)
			i++
		}
	}

	return tst
}

func (s *UnionMemberTypesContext) NamedType(i int) INamedTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedTypeContext); ok {
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

	return t.(INamedTypeContext)
}

func (s *UnionMemberTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMemberTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMemberTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionMemberTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionMemberTypes() (localctx IUnionMemberTypesContext) {
	localctx = NewUnionMemberTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, GraphQLParserRULE_unionMemberTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		p.Match(GraphQLParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__28 {
		{
			p.SetState(572)
			p.Match(GraphQLParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(575)
		p.NamedType()
	}
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__28 {
		{
			p.SetState(576)
			p.Match(GraphQLParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(577)
			p.NamedType()
		}

		p.SetState(582)
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

// IUnionTypeExtensionContext is an interface to support dynamic dispatch.
type IUnionTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	UnionMemberTypes() IUnionMemberTypesContext
	Directives() IDirectivesContext

	// IsUnionTypeExtensionContext differentiates from other interfaces.
	IsUnionTypeExtensionContext()
}

type UnionTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeExtensionContext() *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension
	return p
}

func InitEmptyUnionTypeExtensionContext(p *UnionTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension
}

func (*UnionTypeExtensionContext) IsUnionTypeExtensionContext() {}

func NewUnionTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension

	return p
}

func (s *UnionTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeExtensionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionMemberTypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionTypeExtension() (localctx IUnionTypeExtensionContext) {
	localctx = NewUnionTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, GraphQLParserRULE_unionTypeExtension)
	var _la int

	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(583)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(584)
			p.Match(GraphQLParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Name()
		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(586)
				p.Directives()
			}

		}
		{
			p.SetState(589)
			p.UnionMemberTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(591)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(592)
			p.Match(GraphQLParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)
			p.Name()
		}
		{
			p.SetState(594)
			p.Directives()
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

// IEnumTypeDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	Directives() IDirectivesContext
	EnumValuesDefinition() IEnumValuesDefinitionContext

	// IsEnumTypeDefinitionContext differentiates from other interfaces.
	IsEnumTypeDefinitionContext()
}

type EnumTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeDefinitionContext() *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition
	return p
}

func InitEmptyEnumTypeDefinitionContext(p *EnumTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition
}

func (*EnumTypeDefinitionContext) IsEnumTypeDefinitionContext() {}

func NewEnumTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition

	return p
}

func (s *EnumTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeDefinitionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValuesDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumTypeDefinition() (localctx IEnumTypeDefinitionContext) {
	localctx = NewEnumTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, GraphQLParserRULE_enumTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(598)
			p.Description()
		}

	}
	{
		p.SetState(601)
		p.Match(GraphQLParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(602)
		p.Name()
	}
	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(603)
			p.Directives()
		}

	}
	p.SetState(607)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(606)
			p.EnumValuesDefinition()
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

// IEnumValuesDefinitionContext is an interface to support dynamic dispatch.
type IEnumValuesDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEnumValueDefinition() []IEnumValueDefinitionContext
	EnumValueDefinition(i int) IEnumValueDefinitionContext

	// IsEnumValuesDefinitionContext differentiates from other interfaces.
	IsEnumValuesDefinitionContext()
}

type EnumValuesDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValuesDefinitionContext() *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition
	return p
}

func InitEmptyEnumValuesDefinitionContext(p *EnumValuesDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition
}

func (*EnumValuesDefinitionContext) IsEnumValuesDefinitionContext() {}

func NewEnumValuesDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition

	return p
}

func (s *EnumValuesDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValuesDefinitionContext) AllEnumValueDefinition() []IEnumValueDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueDefinitionContext); ok {
			tst[i] = t.(IEnumValueDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *EnumValuesDefinitionContext) EnumValueDefinition(i int) IEnumValueDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueDefinitionContext); ok {
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

	return t.(IEnumValueDefinitionContext)
}

func (s *EnumValuesDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValuesDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValuesDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValuesDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValuesDefinition() (localctx IEnumValuesDefinitionContext) {
	localctx = NewEnumValuesDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, GraphQLParserRULE_enumValuesDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63050394783186944) != 0) {
		{
			p.SetState(610)
			p.EnumValueDefinition()
		}

		p.SetState(613)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(615)
		p.Match(GraphQLParserT__4)
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

// IEnumValueDefinitionContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumValue() IEnumValueContext
	Description() IDescriptionContext
	Directives() IDirectivesContext

	// IsEnumValueDefinitionContext differentiates from other interfaces.
	IsEnumValueDefinitionContext()
}

type EnumValueDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionContext() *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition
	return p
}

func InitEmptyEnumValueDefinitionContext(p *EnumValueDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition
}

func (*EnumValueDefinitionContext) IsEnumValueDefinitionContext() {}

func NewEnumValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition

	return p
}

func (s *EnumValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionContext) EnumValue() IEnumValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValueDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumValueDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValueDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValueDefinition() (localctx IEnumValueDefinitionContext) {
	localctx = NewEnumValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, GraphQLParserRULE_enumValueDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(618)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(617)
			p.Description()
		}

	}
	{
		p.SetState(620)
		p.EnumValue()
	}
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(621)
			p.Directives()
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

// IEnumTypeExtensionContext is an interface to support dynamic dispatch.
type IEnumTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	EnumValuesDefinition() IEnumValuesDefinitionContext
	Directives() IDirectivesContext

	// IsEnumTypeExtensionContext differentiates from other interfaces.
	IsEnumTypeExtensionContext()
}

type EnumTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeExtensionContext() *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension
	return p
}

func InitEmptyEnumTypeExtensionContext(p *EnumTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension
}

func (*EnumTypeExtensionContext) IsEnumTypeExtensionContext() {}

func NewEnumTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension

	return p
}

func (s *EnumTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeExtensionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValuesDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumTypeExtension() (localctx IEnumTypeExtensionContext) {
	localctx = NewEnumTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, GraphQLParserRULE_enumTypeExtension)
	var _la int

	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 83, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(624)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.Match(GraphQLParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Name()
		}
		p.SetState(628)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(627)
				p.Directives()
			}

		}
		{
			p.SetState(630)
			p.EnumValuesDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(632)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(633)
			p.Match(GraphQLParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)
			p.Name()
		}
		{
			p.SetState(635)
			p.Directives()
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

// IInputObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Description() IDescriptionContext
	Directives() IDirectivesContext
	InputFieldsDefinition() IInputFieldsDefinitionContext

	// IsInputObjectTypeDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeDefinitionContext()
}

type InputObjectTypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeDefinitionContext() *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition
	return p
}

func InitEmptyInputObjectTypeDefinitionContext(p *InputObjectTypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition
}

func (*InputObjectTypeDefinitionContext) IsInputObjectTypeDefinitionContext() {}

func NewInputObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition

	return p
}

func (s *InputObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeDefinitionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputObjectTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputObjectTypeDefinition() (localctx IInputObjectTypeDefinitionContext) {
	localctx = NewInputObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, GraphQLParserRULE_inputObjectTypeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(640)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(639)
			p.Description()
		}

	}
	{
		p.SetState(642)
		p.Match(GraphQLParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(643)
		p.Name()
	}
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__19 {
		{
			p.SetState(644)
			p.Directives()
		}

	}
	p.SetState(648)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 86, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(647)
			p.InputFieldsDefinition()
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

// IInputFieldsDefinitionContext is an interface to support dynamic dispatch.
type IInputFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInputValueDefinition() []IInputValueDefinitionContext
	InputValueDefinition(i int) IInputValueDefinitionContext

	// IsInputFieldsDefinitionContext differentiates from other interfaces.
	IsInputFieldsDefinitionContext()
}

type InputFieldsDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldsDefinitionContext() *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition
	return p
}

func InitEmptyInputFieldsDefinitionContext(p *InputFieldsDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition
}

func (*InputFieldsDefinitionContext) IsInputFieldsDefinitionContext() {}

func NewInputFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition

	return p
}

func (s *InputFieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputValueDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IInputValueDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputValueDefinitionContext); ok {
			tst[i] = t.(IInputValueDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *InputFieldsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputValueDefinitionContext); ok {
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

	return t.(IInputValueDefinitionContext)
}

func (s *InputFieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputFieldsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputFieldsDefinition() (localctx IInputFieldsDefinitionContext) {
	localctx = NewInputFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, GraphQLParserRULE_inputFieldsDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(650)
		p.Match(GraphQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63050394783186944) != 0) {
		{
			p.SetState(651)
			p.InputValueDefinition()
		}

		p.SetState(654)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(656)
		p.Match(GraphQLParserT__4)
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

// IInputObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IInputObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	InputFieldsDefinition() IInputFieldsDefinitionContext
	Directives() IDirectivesContext

	// IsInputObjectTypeExtensionContext differentiates from other interfaces.
	IsInputObjectTypeExtensionContext()
}

type InputObjectTypeExtensionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeExtensionContext() *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension
	return p
}

func InitEmptyInputObjectTypeExtensionContext(p *InputObjectTypeExtensionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension
}

func (*InputObjectTypeExtensionContext) IsInputObjectTypeExtensionContext() {}

func NewInputObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension

	return p
}

func (s *InputObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeExtensionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeExtensionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputFieldsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectivesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputObjectTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputObjectTypeExtension() (localctx IInputObjectTypeExtensionContext) {
	localctx = NewInputObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, GraphQLParserRULE_inputObjectTypeExtension)
	var _la int

	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 89, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(658)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(659)
			p.Match(GraphQLParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(660)
			p.Name()
		}
		p.SetState(662)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__19 {
			{
				p.SetState(661)
				p.Directives()
			}

		}
		{
			p.SetState(664)
			p.InputFieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(666)
			p.Match(GraphQLParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(667)
			p.Match(GraphQLParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(668)
			p.Name()
		}
		{
			p.SetState(669)
			p.Directives()
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

// IDirectiveDefinitionContext is an interface to support dynamic dispatch.
type IDirectiveDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	DirectiveLocations() IDirectiveLocationsContext
	Description() IDescriptionContext
	ArgumentsDefinition() IArgumentsDefinitionContext

	// IsDirectiveDefinitionContext differentiates from other interfaces.
	IsDirectiveDefinitionContext()
}

type DirectiveDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveDefinitionContext() *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveDefinition
	return p
}

func InitEmptyDirectiveDefinitionContext(p *DirectiveDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveDefinition
}

func (*DirectiveDefinitionContext) IsDirectiveDefinitionContext() {}

func NewDirectiveDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveDefinition

	return p
}

func (s *DirectiveDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveDefinitionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveDefinitionContext) DirectiveLocations() IDirectiveLocationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveLocationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveDefinitionContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DirectiveDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *DirectiveDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectiveDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DirectiveDefinition() (localctx IDirectiveDefinitionContext) {
	localctx = NewDirectiveDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, GraphQLParserRULE_directiveDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserSTRING || _la == GraphQLParserBLOCK_STRING {
		{
			p.SetState(673)
			p.Description()
		}

	}
	{
		p.SetState(676)
		p.Match(GraphQLParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(677)
		p.Match(GraphQLParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(678)
		p.Name()
	}
	p.SetState(680)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(679)
			p.ArgumentsDefinition()
		}

	}
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__32 {
		{
			p.SetState(682)
			p.Match(GraphQLParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(685)
		p.Match(GraphQLParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(686)
		p.DirectiveLocations()
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

// IDirectiveLocationsContext is an interface to support dynamic dispatch.
type IDirectiveLocationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDirectiveLocation() []IDirectiveLocationContext
	DirectiveLocation(i int) IDirectiveLocationContext

	// IsDirectiveLocationsContext differentiates from other interfaces.
	IsDirectiveLocationsContext()
}

type DirectiveLocationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationsContext() *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocations
	return p
}

func InitEmptyDirectiveLocationsContext(p *DirectiveLocationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocations
}

func (*DirectiveLocationsContext) IsDirectiveLocationsContext() {}

func NewDirectiveLocationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocations

	return p
}

func (s *DirectiveLocationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationsContext) AllDirectiveLocation() []IDirectiveLocationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirectiveLocationContext); ok {
			len++
		}
	}

	tst := make([]IDirectiveLocationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirectiveLocationContext); ok {
			tst[i] = t.(IDirectiveLocationContext)
			i++
		}
	}

	return tst
}

func (s *DirectiveLocationsContext) DirectiveLocation(i int) IDirectiveLocationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveLocationContext); ok {
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

	return t.(IDirectiveLocationContext)
}

func (s *DirectiveLocationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectiveLocations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DirectiveLocations() (localctx IDirectiveLocationsContext) {
	localctx = NewDirectiveLocationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, GraphQLParserRULE_directiveLocations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(688)
		p.DirectiveLocation()
	}
	p.SetState(693)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GraphQLParserT__28 {
		{
			p.SetState(689)
			p.Match(GraphQLParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(690)
			p.DirectiveLocation()
		}

		p.SetState(695)
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

// IDirectiveLocationContext is an interface to support dynamic dispatch.
type IDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExecutableDirectiveLocation() IExecutableDirectiveLocationContext
	TypeSystemDirectiveLocation() ITypeSystemDirectiveLocationContext

	// IsDirectiveLocationContext differentiates from other interfaces.
	IsDirectiveLocationContext()
}

type DirectiveLocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationContext() *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocation
	return p
}

func InitEmptyDirectiveLocationContext(p *DirectiveLocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocation
}

func (*DirectiveLocationContext) IsDirectiveLocationContext() {}

func NewDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocation

	return p
}

func (s *DirectiveLocationContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationContext) ExecutableDirectiveLocation() IExecutableDirectiveLocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecutableDirectiveLocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecutableDirectiveLocationContext)
}

func (s *DirectiveLocationContext) TypeSystemDirectiveLocation() ITypeSystemDirectiveLocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSystemDirectiveLocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDirectiveLocationContext)
}

func (s *DirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocation(s)
	}
}

func (s *DirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocation(s)
	}
}

func (s *DirectiveLocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectiveLocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DirectiveLocation() (localctx IDirectiveLocationContext) {
	localctx = NewDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, GraphQLParserRULE_directiveLocation)
	p.SetState(698)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__33, GraphQLParserT__34, GraphQLParserT__35, GraphQLParserT__36, GraphQLParserT__37, GraphQLParserT__38, GraphQLParserT__39, GraphQLParserT__40:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(696)
			p.ExecutableDirectiveLocation()
		}

	case GraphQLParserT__41, GraphQLParserT__42, GraphQLParserT__43, GraphQLParserT__44, GraphQLParserT__45, GraphQLParserT__46, GraphQLParserT__47, GraphQLParserT__48, GraphQLParserT__49, GraphQLParserT__50, GraphQLParserT__51:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(697)
			p.TypeSystemDirectiveLocation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IExecutableDirectiveLocationContext is an interface to support dynamic dispatch.
type IExecutableDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExecutableDirectiveLocationContext differentiates from other interfaces.
	IsExecutableDirectiveLocationContext()
}

type ExecutableDirectiveLocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDirectiveLocationContext() *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDirectiveLocation
	return p
}

func InitEmptyExecutableDirectiveLocationContext(p *ExecutableDirectiveLocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDirectiveLocation
}

func (*ExecutableDirectiveLocationContext) IsExecutableDirectiveLocationContext() {}

func NewExecutableDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDirectiveLocation

	return p
}

func (s *ExecutableDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *ExecutableDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDirectiveLocation(s)
	}
}

func (s *ExecutableDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDirectiveLocation(s)
	}
}

func (s *ExecutableDirectiveLocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitExecutableDirectiveLocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ExecutableDirectiveLocation() (localctx IExecutableDirectiveLocationContext) {
	localctx = NewExecutableDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, GraphQLParserRULE_executableDirectiveLocation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(700)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4380866641920) != 0) {
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

// ITypeSystemDirectiveLocationContext is an interface to support dynamic dispatch.
type ITypeSystemDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeSystemDirectiveLocationContext differentiates from other interfaces.
	IsTypeSystemDirectiveLocationContext()
}

type TypeSystemDirectiveLocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDirectiveLocationContext() *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDirectiveLocation
	return p
}

func InitEmptyTypeSystemDirectiveLocationContext(p *TypeSystemDirectiveLocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDirectiveLocation
}

func (*TypeSystemDirectiveLocationContext) IsTypeSystemDirectiveLocationContext() {}

func NewTypeSystemDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDirectiveLocation

	return p
}

func (s *TypeSystemDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeSystemDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDirectiveLocation(s)
	}
}

func (s *TypeSystemDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDirectiveLocation(s)
	}
}

func (s *TypeSystemDirectiveLocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemDirectiveLocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemDirectiveLocation() (localctx ITypeSystemDirectiveLocationContext) {
	localctx = NewTypeSystemDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, GraphQLParserRULE_typeSystemDirectiveLocation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(702)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9002801208229888) != 0) {
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

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GraphQLParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, GraphQLParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(704)
		p.Match(GraphQLParserNAME)
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

func (p *GraphQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 49:
		var t *ImplementsInterfacesContext = nil
		if localctx != nil {
			t = localctx.(*ImplementsInterfacesContext)
		}
		return p.ImplementsInterfaces_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GraphQLParser) ImplementsInterfaces_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
