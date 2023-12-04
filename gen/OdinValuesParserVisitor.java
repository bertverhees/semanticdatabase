// Generated from /home/verhees/Development/semanticdatabase/antlr_source/bmm/OdinValuesParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link OdinValuesParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface OdinValuesParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#primitiveObject}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimitiveObject(OdinValuesParser.PrimitiveObjectContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#primitiveValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimitiveValue(OdinValuesParser.PrimitiveValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#primitiveListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimitiveListValue(OdinValuesParser.PrimitiveListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#primitiveIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimitiveIntervalValue(OdinValuesParser.PrimitiveIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#stringValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStringValue(OdinValuesParser.StringValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#stringListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStringListValue(OdinValuesParser.StringListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#integerValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerValue(OdinValuesParser.IntegerValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#integerListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerListValue(OdinValuesParser.IntegerListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#integerIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerIntervalValue(OdinValuesParser.IntegerIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#integerIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntegerIntervalListValue(OdinValuesParser.IntegerIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#realValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRealValue(OdinValuesParser.RealValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#realListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRealListValue(OdinValuesParser.RealListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#realIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRealIntervalValue(OdinValuesParser.RealIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#realIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRealIntervalListValue(OdinValuesParser.RealIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#booleanValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBooleanValue(OdinValuesParser.BooleanValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#booleanListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBooleanListValue(OdinValuesParser.BooleanListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#characterValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCharacterValue(OdinValuesParser.CharacterValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#characterListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCharacterListValue(OdinValuesParser.CharacterListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateValue(OdinValuesParser.DateValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateListValue(OdinValuesParser.DateListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateIntervalValue(OdinValuesParser.DateIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateIntervalListValue(OdinValuesParser.DateIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#timeValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTimeValue(OdinValuesParser.TimeValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#timeListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTimeListValue(OdinValuesParser.TimeListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#timeIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTimeIntervalValue(OdinValuesParser.TimeIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#timeIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTimeIntervalListValue(OdinValuesParser.TimeIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateTimeValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateTimeValue(OdinValuesParser.DateTimeValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateTimeListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateTimeListValue(OdinValuesParser.DateTimeListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateTimeIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateTimeIntervalValue(OdinValuesParser.DateTimeIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#dateTimeIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDateTimeIntervalListValue(OdinValuesParser.DateTimeIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#durationValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDurationValue(OdinValuesParser.DurationValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#durationListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDurationListValue(OdinValuesParser.DurationListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#durationIntervalValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDurationIntervalValue(OdinValuesParser.DurationIntervalValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#durationIntervalListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDurationIntervalListValue(OdinValuesParser.DurationIntervalListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#termCodeValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTermCodeValue(OdinValuesParser.TermCodeValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#termCodeListValue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTermCodeListValue(OdinValuesParser.TermCodeListValueContext ctx);
	/**
	 * Visit a parse tree produced by {@link OdinValuesParser#relop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRelop(OdinValuesParser.RelopContext ctx);
}