// Code generated from OdinParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // OdinParser

import "github.com/antlr4-go/antlr/v4"

// BaseOdinParserListener is a complete listener for a parse tree produced by OdinParser.
type BaseOdinParserListener struct{}

var _ OdinParserListener = &BaseOdinParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOdinParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOdinParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOdinParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOdinParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterOdinObject is called when production odinObject is entered.
func (s *BaseOdinParserListener) EnterOdinObject(ctx *OdinObjectContext) {}

// ExitOdinObject is called when production odinObject is exited.
func (s *BaseOdinParserListener) ExitOdinObject(ctx *OdinObjectContext) {}

// EnterOdinAttrVal is called when production odinAttrVal is entered.
func (s *BaseOdinParserListener) EnterOdinAttrVal(ctx *OdinAttrValContext) {}

// ExitOdinAttrVal is called when production odinAttrVal is exited.
func (s *BaseOdinParserListener) ExitOdinAttrVal(ctx *OdinAttrValContext) {}

// EnterOdinAttrName is called when production odinAttrName is entered.
func (s *BaseOdinParserListener) EnterOdinAttrName(ctx *OdinAttrNameContext) {}

// ExitOdinAttrName is called when production odinAttrName is exited.
func (s *BaseOdinParserListener) ExitOdinAttrName(ctx *OdinAttrNameContext) {}

// EnterOdinObjectBlock is called when production odinObjectBlock is entered.
func (s *BaseOdinParserListener) EnterOdinObjectBlock(ctx *OdinObjectBlockContext) {}

// ExitOdinObjectBlock is called when production odinObjectBlock is exited.
func (s *BaseOdinParserListener) ExitOdinObjectBlock(ctx *OdinObjectBlockContext) {}

// EnterOdinObjectValueBlock is called when production odinObjectValueBlock is entered.
func (s *BaseOdinParserListener) EnterOdinObjectValueBlock(ctx *OdinObjectValueBlockContext) {}

// ExitOdinObjectValueBlock is called when production odinObjectValueBlock is exited.
func (s *BaseOdinParserListener) ExitOdinObjectValueBlock(ctx *OdinObjectValueBlockContext) {}

// EnterRmTypeSpec is called when production rmTypeSpec is entered.
func (s *BaseOdinParserListener) EnterRmTypeSpec(ctx *RmTypeSpecContext) {}

// ExitRmTypeSpec is called when production rmTypeSpec is exited.
func (s *BaseOdinParserListener) ExitRmTypeSpec(ctx *RmTypeSpecContext) {}

// EnterOdinKeyedObject is called when production odinKeyedObject is entered.
func (s *BaseOdinParserListener) EnterOdinKeyedObject(ctx *OdinKeyedObjectContext) {}

// ExitOdinKeyedObject is called when production odinKeyedObject is exited.
func (s *BaseOdinParserListener) ExitOdinKeyedObject(ctx *OdinKeyedObjectContext) {}

// EnterOdinKeySpec is called when production odinKeySpec is entered.
func (s *BaseOdinParserListener) EnterOdinKeySpec(ctx *OdinKeySpecContext) {}

// ExitOdinKeySpec is called when production odinKeySpec is exited.
func (s *BaseOdinParserListener) ExitOdinKeySpec(ctx *OdinKeySpecContext) {}

// EnterOdinObjectReferenceBlock is called when production odinObjectReferenceBlock is entered.
func (s *BaseOdinParserListener) EnterOdinObjectReferenceBlock(ctx *OdinObjectReferenceBlockContext) {
}

// ExitOdinObjectReferenceBlock is called when production odinObjectReferenceBlock is exited.
func (s *BaseOdinParserListener) ExitOdinObjectReferenceBlock(ctx *OdinObjectReferenceBlockContext) {}

// EnterOdinPathList is called when production odinPathList is entered.
func (s *BaseOdinParserListener) EnterOdinPathList(ctx *OdinPathListContext) {}

// ExitOdinPathList is called when production odinPathList is exited.
func (s *BaseOdinParserListener) ExitOdinPathList(ctx *OdinPathListContext) {}

// EnterOdinPath is called when production odinPath is entered.
func (s *BaseOdinParserListener) EnterOdinPath(ctx *OdinPathContext) {}

// ExitOdinPath is called when production odinPath is exited.
func (s *BaseOdinParserListener) ExitOdinPath(ctx *OdinPathContext) {}

// EnterOdinPathSegment is called when production odinPathSegment is entered.
func (s *BaseOdinParserListener) EnterOdinPathSegment(ctx *OdinPathSegmentContext) {}

// ExitOdinPathSegment is called when production odinPathSegment is exited.
func (s *BaseOdinParserListener) ExitOdinPathSegment(ctx *OdinPathSegmentContext) {}

// EnterRmTypeId is called when production rmTypeId is entered.
func (s *BaseOdinParserListener) EnterRmTypeId(ctx *RmTypeIdContext) {}

// ExitRmTypeId is called when production rmTypeId is exited.
func (s *BaseOdinParserListener) ExitRmTypeId(ctx *RmTypeIdContext) {}

// EnterPrimitiveObject is called when production primitiveObject is entered.
func (s *BaseOdinParserListener) EnterPrimitiveObject(ctx *PrimitiveObjectContext) {}

// ExitPrimitiveObject is called when production primitiveObject is exited.
func (s *BaseOdinParserListener) ExitPrimitiveObject(ctx *PrimitiveObjectContext) {}

// EnterPrimitiveValue is called when production primitiveValue is entered.
func (s *BaseOdinParserListener) EnterPrimitiveValue(ctx *PrimitiveValueContext) {}

// ExitPrimitiveValue is called when production primitiveValue is exited.
func (s *BaseOdinParserListener) ExitPrimitiveValue(ctx *PrimitiveValueContext) {}

// EnterPrimitiveListValue is called when production primitiveListValue is entered.
func (s *BaseOdinParserListener) EnterPrimitiveListValue(ctx *PrimitiveListValueContext) {}

// ExitPrimitiveListValue is called when production primitiveListValue is exited.
func (s *BaseOdinParserListener) ExitPrimitiveListValue(ctx *PrimitiveListValueContext) {}

// EnterPrimitiveIntervalValue is called when production primitiveIntervalValue is entered.
func (s *BaseOdinParserListener) EnterPrimitiveIntervalValue(ctx *PrimitiveIntervalValueContext) {}

// ExitPrimitiveIntervalValue is called when production primitiveIntervalValue is exited.
func (s *BaseOdinParserListener) ExitPrimitiveIntervalValue(ctx *PrimitiveIntervalValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseOdinParserListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseOdinParserListener) ExitStringValue(ctx *StringValueContext) {}

// EnterStringListValue is called when production stringListValue is entered.
func (s *BaseOdinParserListener) EnterStringListValue(ctx *StringListValueContext) {}

// ExitStringListValue is called when production stringListValue is exited.
func (s *BaseOdinParserListener) ExitStringListValue(ctx *StringListValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseOdinParserListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseOdinParserListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterIntegerListValue is called when production integerListValue is entered.
func (s *BaseOdinParserListener) EnterIntegerListValue(ctx *IntegerListValueContext) {}

// ExitIntegerListValue is called when production integerListValue is exited.
func (s *BaseOdinParserListener) ExitIntegerListValue(ctx *IntegerListValueContext) {}

// EnterIntegerIntervalValue is called when production integerIntervalValue is entered.
func (s *BaseOdinParserListener) EnterIntegerIntervalValue(ctx *IntegerIntervalValueContext) {}

// ExitIntegerIntervalValue is called when production integerIntervalValue is exited.
func (s *BaseOdinParserListener) ExitIntegerIntervalValue(ctx *IntegerIntervalValueContext) {}

// EnterIntegerIntervalListValue is called when production integerIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterIntegerIntervalListValue(ctx *IntegerIntervalListValueContext) {
}

// ExitIntegerIntervalListValue is called when production integerIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitIntegerIntervalListValue(ctx *IntegerIntervalListValueContext) {}

// EnterRealValue is called when production realValue is entered.
func (s *BaseOdinParserListener) EnterRealValue(ctx *RealValueContext) {}

// ExitRealValue is called when production realValue is exited.
func (s *BaseOdinParserListener) ExitRealValue(ctx *RealValueContext) {}

// EnterRealListValue is called when production realListValue is entered.
func (s *BaseOdinParserListener) EnterRealListValue(ctx *RealListValueContext) {}

// ExitRealListValue is called when production realListValue is exited.
func (s *BaseOdinParserListener) ExitRealListValue(ctx *RealListValueContext) {}

// EnterRealIntervalValue is called when production realIntervalValue is entered.
func (s *BaseOdinParserListener) EnterRealIntervalValue(ctx *RealIntervalValueContext) {}

// ExitRealIntervalValue is called when production realIntervalValue is exited.
func (s *BaseOdinParserListener) ExitRealIntervalValue(ctx *RealIntervalValueContext) {}

// EnterRealIntervalListValue is called when production realIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterRealIntervalListValue(ctx *RealIntervalListValueContext) {}

// ExitRealIntervalListValue is called when production realIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitRealIntervalListValue(ctx *RealIntervalListValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseOdinParserListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseOdinParserListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterBooleanListValue is called when production booleanListValue is entered.
func (s *BaseOdinParserListener) EnterBooleanListValue(ctx *BooleanListValueContext) {}

// ExitBooleanListValue is called when production booleanListValue is exited.
func (s *BaseOdinParserListener) ExitBooleanListValue(ctx *BooleanListValueContext) {}

// EnterCharacterValue is called when production characterValue is entered.
func (s *BaseOdinParserListener) EnterCharacterValue(ctx *CharacterValueContext) {}

// ExitCharacterValue is called when production characterValue is exited.
func (s *BaseOdinParserListener) ExitCharacterValue(ctx *CharacterValueContext) {}

// EnterCharacterListValue is called when production characterListValue is entered.
func (s *BaseOdinParserListener) EnterCharacterListValue(ctx *CharacterListValueContext) {}

// ExitCharacterListValue is called when production characterListValue is exited.
func (s *BaseOdinParserListener) ExitCharacterListValue(ctx *CharacterListValueContext) {}

// EnterDateValue is called when production dateValue is entered.
func (s *BaseOdinParserListener) EnterDateValue(ctx *DateValueContext) {}

// ExitDateValue is called when production dateValue is exited.
func (s *BaseOdinParserListener) ExitDateValue(ctx *DateValueContext) {}

// EnterDateListValue is called when production dateListValue is entered.
func (s *BaseOdinParserListener) EnterDateListValue(ctx *DateListValueContext) {}

// ExitDateListValue is called when production dateListValue is exited.
func (s *BaseOdinParserListener) ExitDateListValue(ctx *DateListValueContext) {}

// EnterDateIntervalValue is called when production dateIntervalValue is entered.
func (s *BaseOdinParserListener) EnterDateIntervalValue(ctx *DateIntervalValueContext) {}

// ExitDateIntervalValue is called when production dateIntervalValue is exited.
func (s *BaseOdinParserListener) ExitDateIntervalValue(ctx *DateIntervalValueContext) {}

// EnterDateIntervalListValue is called when production dateIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterDateIntervalListValue(ctx *DateIntervalListValueContext) {}

// ExitDateIntervalListValue is called when production dateIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitDateIntervalListValue(ctx *DateIntervalListValueContext) {}

// EnterTimeValue is called when production timeValue is entered.
func (s *BaseOdinParserListener) EnterTimeValue(ctx *TimeValueContext) {}

// ExitTimeValue is called when production timeValue is exited.
func (s *BaseOdinParserListener) ExitTimeValue(ctx *TimeValueContext) {}

// EnterTimeListValue is called when production timeListValue is entered.
func (s *BaseOdinParserListener) EnterTimeListValue(ctx *TimeListValueContext) {}

// ExitTimeListValue is called when production timeListValue is exited.
func (s *BaseOdinParserListener) ExitTimeListValue(ctx *TimeListValueContext) {}

// EnterTimeIntervalValue is called when production timeIntervalValue is entered.
func (s *BaseOdinParserListener) EnterTimeIntervalValue(ctx *TimeIntervalValueContext) {}

// ExitTimeIntervalValue is called when production timeIntervalValue is exited.
func (s *BaseOdinParserListener) ExitTimeIntervalValue(ctx *TimeIntervalValueContext) {}

// EnterTimeIntervalListValue is called when production timeIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterTimeIntervalListValue(ctx *TimeIntervalListValueContext) {}

// ExitTimeIntervalListValue is called when production timeIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitTimeIntervalListValue(ctx *TimeIntervalListValueContext) {}

// EnterDateTimeValue is called when production dateTimeValue is entered.
func (s *BaseOdinParserListener) EnterDateTimeValue(ctx *DateTimeValueContext) {}

// ExitDateTimeValue is called when production dateTimeValue is exited.
func (s *BaseOdinParserListener) ExitDateTimeValue(ctx *DateTimeValueContext) {}

// EnterDateTimeListValue is called when production dateTimeListValue is entered.
func (s *BaseOdinParserListener) EnterDateTimeListValue(ctx *DateTimeListValueContext) {}

// ExitDateTimeListValue is called when production dateTimeListValue is exited.
func (s *BaseOdinParserListener) ExitDateTimeListValue(ctx *DateTimeListValueContext) {}

// EnterDateTimeIntervalValue is called when production dateTimeIntervalValue is entered.
func (s *BaseOdinParserListener) EnterDateTimeIntervalValue(ctx *DateTimeIntervalValueContext) {}

// ExitDateTimeIntervalValue is called when production dateTimeIntervalValue is exited.
func (s *BaseOdinParserListener) ExitDateTimeIntervalValue(ctx *DateTimeIntervalValueContext) {}

// EnterDateTimeIntervalListValue is called when production dateTimeIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterDateTimeIntervalListValue(ctx *DateTimeIntervalListValueContext) {
}

// ExitDateTimeIntervalListValue is called when production dateTimeIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitDateTimeIntervalListValue(ctx *DateTimeIntervalListValueContext) {
}

// EnterDurationValue is called when production durationValue is entered.
func (s *BaseOdinParserListener) EnterDurationValue(ctx *DurationValueContext) {}

// ExitDurationValue is called when production durationValue is exited.
func (s *BaseOdinParserListener) ExitDurationValue(ctx *DurationValueContext) {}

// EnterDurationListValue is called when production durationListValue is entered.
func (s *BaseOdinParserListener) EnterDurationListValue(ctx *DurationListValueContext) {}

// ExitDurationListValue is called when production durationListValue is exited.
func (s *BaseOdinParserListener) ExitDurationListValue(ctx *DurationListValueContext) {}

// EnterDurationIntervalValue is called when production durationIntervalValue is entered.
func (s *BaseOdinParserListener) EnterDurationIntervalValue(ctx *DurationIntervalValueContext) {}

// ExitDurationIntervalValue is called when production durationIntervalValue is exited.
func (s *BaseOdinParserListener) ExitDurationIntervalValue(ctx *DurationIntervalValueContext) {}

// EnterDurationIntervalListValue is called when production durationIntervalListValue is entered.
func (s *BaseOdinParserListener) EnterDurationIntervalListValue(ctx *DurationIntervalListValueContext) {
}

// ExitDurationIntervalListValue is called when production durationIntervalListValue is exited.
func (s *BaseOdinParserListener) ExitDurationIntervalListValue(ctx *DurationIntervalListValueContext) {
}

// EnterTermCodeValue is called when production termCodeValue is entered.
func (s *BaseOdinParserListener) EnterTermCodeValue(ctx *TermCodeValueContext) {}

// ExitTermCodeValue is called when production termCodeValue is exited.
func (s *BaseOdinParserListener) ExitTermCodeValue(ctx *TermCodeValueContext) {}

// EnterTermCodeListValue is called when production termCodeListValue is entered.
func (s *BaseOdinParserListener) EnterTermCodeListValue(ctx *TermCodeListValueContext) {}

// ExitTermCodeListValue is called when production termCodeListValue is exited.
func (s *BaseOdinParserListener) ExitTermCodeListValue(ctx *TermCodeListValueContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaseOdinParserListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaseOdinParserListener) ExitRelop(ctx *RelopContext) {}
