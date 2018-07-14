package main

import(
	"unicode"
	"errors"
	"strconv"
    "github.com/satori/go.uuid"
	"regexp"
    "strings"
)

type Token int

const (
	LABEL = iota
	AND
	OR
	LPAR
	RPAR 
    NOT
	UNKNOWN
)

func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsLetterOrSpace(s string) bool {
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsSpace(r) {
            continue
        }

        return false
    }
    return true
}

func IsUuid(s string) bool {
    _, err := uuid.FromString(s)
    if err == nil {
        return true
    }
    return false
}

func IsAssignment(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]*\\.[a-zA-Z]*='[a-zA-Z\\s]*'$", s)
	return match
}

func StrToToken(str string) Token {
	switch str {
		case "&":
			return AND
		case "|":
			return OR
		case "(":
			return LPAR
		case ")":
			return RPAR
        case "~":
            return NOT
		default:
			if IsAssignment(str) || IsLetterOrSpace(str) || IsUuid(str) {
				return LABEL
			}
	}

	return UNKNOWN
}

func IsTokenValid(currentToken Token, lastToken Token) bool {
	switch currentToken {
		case AND:
			if (lastToken == LABEL) || (lastToken == RPAR)  {
				return true
			}
			return false
		case OR:
			if (lastToken == LABEL) || (lastToken == RPAR) {
				return true
			}
			return false
		case LABEL:
			if (lastToken == OR) || (lastToken == AND) || (lastToken == LPAR) || (lastToken == NOT) {
				return true
			}
			return false
		case RPAR:
			if lastToken == LABEL {
				return true
			}
			return false
		case LPAR:
			if (lastToken == LABEL) || (lastToken == AND) || (lastToken == OR) || (lastToken == NOT) {
				return true
			}
			return false
        case NOT:
            if (lastToken == LABEL) || (lastToken == AND) || (lastToken == OR) || (lastToken == LPAR) {
                return true
            }
            return false
	}

	return false
} 

func IsLastTokenValid(lastToken Token) bool {
    if lastToken == AND {
        return false
    } else if lastToken == OR {
        return false
    }
    return true
}

func Tokenize(input string) []string {
    var output []string
    label := ""
    in := StripWhitespaces(input)
    for _, ch := range in {
        if ch == '&' || ch == '|' || ch == '(' || ch == ')' || ch == '~' {
            if label != "" {
                output = append(output, label)
                label = ""
            }
            output = append(output, string(ch))
            
        } else {
            label += string(ch)
        }
    }

    if label != "" {
        output = append(output, label)
    }

    return output
} 

func GetBaseLabel(s string) string {
	for i, r := range s {
        if r == '.' {
            return s[:i]
        }
    }
    return s
}

func StripWhitespacesExceptWithinQuotes(s string) string {
    output := ""

    insideAssignment := false
    for _, r := range s {
        if r == '\'' {
            insideAssignment = !insideAssignment
        }

        if !insideAssignment && r != ' ' {
            output += string(r)
        } else if (insideAssignment) {
            output += string(r)
        }
    }

    return output
}

func StripWhitespaces(s string) string {
    output := ""
    temp := ""
    potentialMultiWordString := false
    for _, r := range s {

        if unicode.IsLetter(r) || unicode.IsSpace(r) || r == '\'' || r == '.' || r == '=' {
            potentialMultiWordString = true
        } else {
            potentialMultiWordString = false

            if temp != "" {
                if strings.Contains(temp, ".") && strings.Contains(temp, "=") { //is it a label assignment?(e.q dog.has='mouth')
                    output += StripWhitespacesExceptWithinQuotes(temp) //remove whitespaces except within the quotes
                } else {
                    output += temp
                }
            }

            output += string(r)

            temp = ""
        }

        if potentialMultiWordString {
            temp += string(r)
        }
    }

    if temp != "" {
        if strings.Contains(temp, ".") && strings.Contains(temp, "=") { //is it a label assignment?(e.q dog.has='mouth')
            output += StripWhitespacesExceptWithinQuotes(temp) //remove whitespaces except within the quotes
        } else {
            output += temp
        }
    }

    return output 
}

type Parser interface {
	Parse() error
}

type QueryParser struct {
	query string
	lastToken Token
    lastStrToken string
	isFirst bool
	brackets int32
    isUuidQuery bool
    version int32
}

type ParseResult struct {
	input string
	query string
    subquery string
    isUuidQuery bool
	//validationQuery string
	queryValues []interface{}
}

func NewQueryParser(query string) *QueryParser {
    return &QueryParser {
        query: query,
        isFirst: true,
        version: 1,
    } 
}

func NewQueryParserV2(query string) *QueryParser {
    return &QueryParser {
        query: query,
        isFirst: true,
        version: 2,
    } 
}

func (p *QueryParser) Parse(offset int) (ParseResult, error) {
	parseResult := ParseResult{}
	parseResult.query = ""
    parseResult.isUuidQuery = p.isUuidQuery

    parseResult.isUuidQuery = false

    tokens := Tokenize(p.query)

    i := offset
    numOfLabels := 1
    for _, token := range tokens {
        //strip tailing and leading white spaces
        token = strings.TrimSpace(token)

        if token == "" {
            continue
        }

    	t := StrToToken(token)
    	if p.isFirst {
    		if !((t == LABEL) || (t == LPAR) || (t == NOT)) {
    			e := "Error: invalid token\n" + token + "\n^\nExpecting 'LABEL' (e.q dog) or 'ASSIGNMENT' (e.q dog.breed='Labrador') "
    			return parseResult, errors.New(e)
    		}

            //use the first entry to determine whether its a UUID or not. We can't have both labels and UUIDs in the same query, so
            //we use the first entry to determine the type of the query.
            parseResult.isUuidQuery = IsUuid(token)


    		p.isFirst = false
    	} else {
    		if !IsTokenValid(t, p.lastToken) {
    			e := "Error: invalid token\n" + token + "\n^\nExpecting 'LABEL' (e.q dog), 'ASSIGNMENT' (e.q dog.breed='Labrador'), '|', '&' or '~' "
    			return parseResult, errors.New(e)
    		}
    	}

    	if t == LABEL {
            if p.version == 1 {
                if parseResult.isUuidQuery {
                    parseResult.query += ("l.uuid = $" + strconv.Itoa(i))
                } else {
                    parseResult.query += ("a.accessor = $" + strconv.Itoa(i))
                }
            } else {
                if !parseResult.isUuidQuery {
                    parseResult.query += ("q.accessors @> ARRAY[$" + strconv.Itoa(i) + "]::text[]")
                    parseResult.subquery += ("a.accessor = $" + strconv.Itoa(i))
                }
            }
    		parseResult.queryValues = append(parseResult.queryValues, token)
    		i += 1
    		numOfLabels += 1
    	} else if t == AND {
    		parseResult.query += "AND"
            parseResult.subquery += "OR"
    	} else if t == OR {
    		parseResult.query += "OR"
            parseResult.subquery += "OR"
        } else if t == NOT {
            parseResult.query += "NOT"
            parseResult.subquery += "NOT"
    	} else {
    		parseResult.query += token
            parseResult.subquery += token
    	}
    	parseResult.query += " "
        parseResult.subquery += " "


    	if t == LPAR {
    		p.brackets += 1
    	}
    	if t == RPAR {
    		p.brackets -= 1
    	}

    	p.lastToken = t
        p.lastStrToken = token
    }

    if len(tokens) > 0 {
        if !IsLastTokenValid(p.lastToken) {
            e := "Error: invalid token\n" + p.lastStrToken + "\n^\nExpecting 'LABEL' (e.q dog)"
            return parseResult, errors.New(e)
        }
    }

    if numOfLabels > 10 {
    	return parseResult, errors.New("Please limit your query to 10 label expressions")
    }

    if p.brackets != 0 {
    	return parseResult, errors.New("brackets mismatch!")
    }

    return parseResult, nil
}
