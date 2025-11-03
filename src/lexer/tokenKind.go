package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND

	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	PLUS
	MINUS
	SLASH
	ASTERISK
	PERCENT

	PLUS_PLUS
	PLUS_EQUALS
	MINUS_MINUS
	MINUS_EQUALS
	SLASH_EQUALS
	ASTERISK_EQUALS
	PERCENT_EQUALS

	//reserved keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	EXPORT
	FROM
	FN
	IF
	ELSE
	FOR
	WHILE
	FOREACH
    TYPEOF
	IN
)

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT_EQUALS:
		return "not_equals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case SLASH_EQUALS:
		return "slash_equals"
	case ASTERISK_EQUALS:
		return "asterisk_equals"
	case PERCENT_EQUALS:
		return "percent_equals"
	case MINUS:
		return "minus"
	case PLUS:
		return "plus"
	case SLASH:
		return "slash"
	case ASTERISK:
		return "star"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:	
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

var reserved_lookup = map[string]TokenKind{
	"let":		LET,
	"const":	CONST,
	"class":	CLASS,
	"new":		NEW,
	"import":	IMPORT,
	"export":	EXPORT,
	"from":		FROM,
	"fn":		FN,
	"if":		IF,
	"else":		ELSE,
	"for":		FOR,
	"while":	WHILE,
	"foreach":	FOREACH,
	"typeof":	TYPEOF,
	"in":		IN,
}