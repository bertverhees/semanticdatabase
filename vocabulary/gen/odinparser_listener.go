// Generated from OdinParser.g4 by ANTLR 4.7.

package parser // OdinParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OdinParserListener is a complete listener for a parse tree produced by OdinParser.
type OdinParserListener interface {
	antlr.ParseTreeListener

	// EnterOdinObject is called when entering the odinObject production.
	EnterOdinObject(c *OdinObjectContext)

	// EnterOdinAttrVal is called when entering the odinAttrVal production.
	EnterOdinAttrVal(c *OdinAttrValContext)

	// EnterOdinAttrName is called when entering the odinAttrName production.
	EnterOdinAttrName(c *OdinAttrNameContext)

	// EnterOdinObjectBlock is called when entering the odinObjectBlock production.
	EnterOdinObjectBlock(c *OdinObjectBlockContext)

	// EnterOdinObjectValueBlock is called when entering the odinObjectValueBlock production.
	EnterOdinObjectValueBlock(c *OdinObjectValueBlockContext)

	// EnterRmTypeSpec is called when entering the rmTypeSpec production.
	EnterRmTypeSpec(c *RmTypeSpecContext)

	// EnterOdinKeyedObject is called when entering the odinKeyedObject production.
	EnterOdinKeyedObject(c *OdinKeyedObjectContext)

	// EnterOdinKeySpec is called when entering the odinKeySpec production.
	EnterOdinKeySpec(c *OdinKeySpecContext)

	// EnterOdinObjectReferenceBlock is called when entering the odinObjectReferenceBlock production.
	EnterOdinObjectReferenceBlock(c *OdinObjectReferenceBlockContext)

	// EnterOdinPathList is called when entering the odinPathList production.
	EnterOdinPathList(c *OdinPathListContext)

	// EnterOdinPath is called when entering the odinPath production.
	EnterOdinPath(c *OdinPathContext)

	// EnterOdinPathSegment is called when entering the odinPathSegment production.
	EnterOdinPathSegment(c *OdinPathSegmentContext)

	// EnterRmTypeId is called when entering the rmTypeId production.
	EnterRmTypeId(c *RmTypeIdContext)

	// EnterPrimitiveObject is called when entering the primitiveObject production.
	EnterPrimitiveObject(c *PrimitiveObjectContext)

	// EnterPrimitiveValue is called when entering the primitiveValue production.
	EnterPrimitiveValue(c *PrimitiveValueContext)

	// EnterPrimitiveListValue is called when entering the primitiveListValue production.
	EnterPrimitiveListValue(c *PrimitiveListValueContext)

	// EnterPrimitiveIntervalValue is called when entering the primitiveIntervalValue production.
	EnterPrimitiveIntervalValue(c *PrimitiveIntervalValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterStringListValue is called when entering the stringListValue production.
	EnterStringListValue(c *StringListValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterIntegerListValue is called when entering the integerListValue production.
	EnterIntegerListValue(c *IntegerListValueContext)

	// EnterIntegerIntervalValue is called when entering the integerIntervalValue production.
	EnterIntegerIntervalValue(c *IntegerIntervalValueContext)

	// EnterIntegerIntervalListValue is called when entering the integerIntervalListValue production.
	EnterIntegerIntervalListValue(c *IntegerIntervalListValueContext)

	// EnterRealValue is called when entering the realValue production.
	EnterRealValue(c *RealValueContext)

	// EnterRealListValue is called when entering the realListValue production.
	EnterRealListValue(c *RealListValueContext)

	// EnterRealIntervalValue is called when entering the realIntervalValue production.
	EnterRealIntervalValue(c *RealIntervalValueContext)

	// EnterRealIntervalListValue is called when entering the realIntervalListValue production.
	EnterRealIntervalListValue(c *RealIntervalListValueContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterBooleanListValue is called when entering the booleanListValue production.
	EnterBooleanListValue(c *BooleanListValueContext)

	// EnterCharacterValue is called when entering the characterValue production.
	EnterCharacterValue(c *CharacterValueContext)

	// EnterCharacterListValue is called when entering the characterListValue production.
	EnterCharacterListValue(c *CharacterListValueContext)

	// EnterDateValue is called when entering the dateValue production.
	EnterDateValue(c *DateValueContext)

	// EnterDateListValue is called when entering the dateListValue production.
	EnterDateListValue(c *DateListValueContext)

	// EnterDateIntervalValue is called when entering the dateIntervalValue production.
	EnterDateIntervalValue(c *DateIntervalValueContext)

	// EnterDateIntervalListValue is called when entering the dateIntervalListValue production.
	EnterDateIntervalListValue(c *DateIntervalListValueContext)

	// EnterTimeValue is called when entering the timeValue production.
	EnterTimeValue(c *TimeValueContext)

	// EnterTimeListValue is called when entering the timeListValue production.
	EnterTimeListValue(c *TimeListValueContext)

	// EnterTimeIntervalValue is called when entering the timeIntervalValue production.
	EnterTimeIntervalValue(c *TimeIntervalValueContext)

	// EnterTimeIntervalListValue is called when entering the timeIntervalListValue production.
	EnterTimeIntervalListValue(c *TimeIntervalListValueContext)

	// EnterDateTimeValue is called when entering the dateTimeValue production.
	EnterDateTimeValue(c *DateTimeValueContext)

	// EnterDateTimeListValue is called when entering the dateTimeListValue production.
	EnterDateTimeListValue(c *DateTimeListValueContext)

	// EnterDateTimeIntervalValue is called when entering the dateTimeIntervalValue production.
	EnterDateTimeIntervalValue(c *DateTimeIntervalValueContext)

	// EnterDateTimeIntervalListValue is called when entering the dateTimeIntervalListValue production.
	EnterDateTimeIntervalListValue(c *DateTimeIntervalListValueContext)

	// EnterDurationValue is called when entering the durationValue production.
	EnterDurationValue(c *DurationValueContext)

	// EnterDurationListValue is called when entering the durationListValue production.
	EnterDurationListValue(c *DurationListValueContext)

	// EnterDurationIntervalValue is called when entering the durationIntervalValue production.
	EnterDurationIntervalValue(c *DurationIntervalValueContext)

	// EnterDurationIntervalListValue is called when entering the durationIntervalListValue production.
	EnterDurationIntervalListValue(c *DurationIntervalListValueContext)

	// EnterTermCodeValue is called when entering the termCodeValue production.
	EnterTermCodeValue(c *TermCodeValueContext)

	// EnterTermCodeListValue is called when entering the termCodeListValue production.
	EnterTermCodeListValue(c *TermCodeListValueContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitOdinObject is called when exiting the odinObject production.
	ExitOdinObject(c *OdinObjectContext)

	// ExitOdinAttrVal is called when exiting the odinAttrVal production.
	ExitOdinAttrVal(c *OdinAttrValContext)

	// ExitOdinAttrName is called when exiting the odinAttrName production.
	ExitOdinAttrName(c *OdinAttrNameContext)

	// ExitOdinObjectBlock is called when exiting the odinObjectBlock production.
	ExitOdinObjectBlock(c *OdinObjectBlockContext)

	// ExitOdinObjectValueBlock is called when exiting the odinObjectValueBlock production.
	ExitOdinObjectValueBlock(c *OdinObjectValueBlockContext)

	// ExitRmTypeSpec is called when exiting the rmTypeSpec production.
	ExitRmTypeSpec(c *RmTypeSpecContext)

	// ExitOdinKeyedObject is called when exiting the odinKeyedObject production.
	ExitOdinKeyedObject(c *OdinKeyedObjectContext)

	// ExitOdinKeySpec is called when exiting the odinKeySpec production.
	ExitOdinKeySpec(c *OdinKeySpecContext)

	// ExitOdinObjectReferenceBlock is called when exiting the odinObjectReferenceBlock production.
	ExitOdinObjectReferenceBlock(c *OdinObjectReferenceBlockContext)

	// ExitOdinPathList is called when exiting the odinPathList production.
	ExitOdinPathList(c *OdinPathListContext)

	// ExitOdinPath is called when exiting the odinPath production.
	ExitOdinPath(c *OdinPathContext)

	// ExitOdinPathSegment is called when exiting the odinPathSegment production.
	ExitOdinPathSegment(c *OdinPathSegmentContext)

	// ExitRmTypeId is called when exiting the rmTypeId production.
	ExitRmTypeId(c *RmTypeIdContext)

	// ExitPrimitiveObject is called when exiting the primitiveObject production.
	ExitPrimitiveObject(c *PrimitiveObjectContext)

	// ExitPrimitiveValue is called when exiting the primitiveValue production.
	ExitPrimitiveValue(c *PrimitiveValueContext)

	// ExitPrimitiveListValue is called when exiting the primitiveListValue production.
	ExitPrimitiveListValue(c *PrimitiveListValueContext)

	// ExitPrimitiveIntervalValue is called when exiting the primitiveIntervalValue production.
	ExitPrimitiveIntervalValue(c *PrimitiveIntervalValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitStringListValue is called when exiting the stringListValue production.
	ExitStringListValue(c *StringListValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitIntegerListValue is called when exiting the integerListValue production.
	ExitIntegerListValue(c *IntegerListValueContext)

	// ExitIntegerIntervalValue is called when exiting the integerIntervalValue production.
	ExitIntegerIntervalValue(c *IntegerIntervalValueContext)

	// ExitIntegerIntervalListValue is called when exiting the integerIntervalListValue production.
	ExitIntegerIntervalListValue(c *IntegerIntervalListValueContext)

	// ExitRealValue is called when exiting the realValue production.
	ExitRealValue(c *RealValueContext)

	// ExitRealListValue is called when exiting the realListValue production.
	ExitRealListValue(c *RealListValueContext)

	// ExitRealIntervalValue is called when exiting the realIntervalValue production.
	ExitRealIntervalValue(c *RealIntervalValueContext)

	// ExitRealIntervalListValue is called when exiting the realIntervalListValue production.
	ExitRealIntervalListValue(c *RealIntervalListValueContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitBooleanListValue is called when exiting the booleanListValue production.
	ExitBooleanListValue(c *BooleanListValueContext)

	// ExitCharacterValue is called when exiting the characterValue production.
	ExitCharacterValue(c *CharacterValueContext)

	// ExitCharacterListValue is called when exiting the characterListValue production.
	ExitCharacterListValue(c *CharacterListValueContext)

	// ExitDateValue is called when exiting the dateValue production.
	ExitDateValue(c *DateValueContext)

	// ExitDateListValue is called when exiting the dateListValue production.
	ExitDateListValue(c *DateListValueContext)

	// ExitDateIntervalValue is called when exiting the dateIntervalValue production.
	ExitDateIntervalValue(c *DateIntervalValueContext)

	// ExitDateIntervalListValue is called when exiting the dateIntervalListValue production.
	ExitDateIntervalListValue(c *DateIntervalListValueContext)

	// ExitTimeValue is called when exiting the timeValue production.
	ExitTimeValue(c *TimeValueContext)

	// ExitTimeListValue is called when exiting the timeListValue production.
	ExitTimeListValue(c *TimeListValueContext)

	// ExitTimeIntervalValue is called when exiting the timeIntervalValue production.
	ExitTimeIntervalValue(c *TimeIntervalValueContext)

	// ExitTimeIntervalListValue is called when exiting the timeIntervalListValue production.
	ExitTimeIntervalListValue(c *TimeIntervalListValueContext)

	// ExitDateTimeValue is called when exiting the dateTimeValue production.
	ExitDateTimeValue(c *DateTimeValueContext)

	// ExitDateTimeListValue is called when exiting the dateTimeListValue production.
	ExitDateTimeListValue(c *DateTimeListValueContext)

	// ExitDateTimeIntervalValue is called when exiting the dateTimeIntervalValue production.
	ExitDateTimeIntervalValue(c *DateTimeIntervalValueContext)

	// ExitDateTimeIntervalListValue is called when exiting the dateTimeIntervalListValue production.
	ExitDateTimeIntervalListValue(c *DateTimeIntervalListValueContext)

	// ExitDurationValue is called when exiting the durationValue production.
	ExitDurationValue(c *DurationValueContext)

	// ExitDurationListValue is called when exiting the durationListValue production.
	ExitDurationListValue(c *DurationListValueContext)

	// ExitDurationIntervalValue is called when exiting the durationIntervalValue production.
	ExitDurationIntervalValue(c *DurationIntervalValueContext)

	// ExitDurationIntervalListValue is called when exiting the durationIntervalListValue production.
	ExitDurationIntervalListValue(c *DurationIntervalListValueContext)

	// ExitTermCodeValue is called when exiting the termCodeValue production.
	ExitTermCodeValue(c *TermCodeValueContext)

	// ExitTermCodeListValue is called when exiting the termCodeListValue production.
	ExitTermCodeListValue(c *TermCodeListValueContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
