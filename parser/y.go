
//line parser/glop.y:2
// Package parser contains the logic to extract the structure of the code. This
// file contains the YACC grammar. See README.md to the corresponding y.go from
// it.
package parser
import __yyfmt__ "fmt"
//line parser/glop.y:5
		
import (
  "github.com/palats/glop/nodes"
)


//line parser/glop.y:13
type yySymType struct {
	yys int
  token tokenInfo
  node nodes.Node
}

const tokOpen = 57346
const tokClose = 57347
const tokIdentifier = 57348
const tokInteger = 57349

var yyToknames = []string{
	"tokOpen",
	"tokClose",
	"tokIdentifier",
	"tokInteger",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser/glop.y:71



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 9
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 15

var yyAct = []int{

	3, 8, 5, 6, 7, 3, 4, 5, 6, 10,
	9, 2, 1, 0, 11,
}
var yyPact = []int{

	1, -1000, -1000, -4, -1000, -1000, -1000, 4, -1000, 1,
	-1000, -1000,
}
var yyPgo = []int{

	0, 12, 10, 6, 4,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 4, 4, 3, 3,
}
var yyR2 = []int{

	0, 1, 3, 2, 1, 1, 2, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, 6, 7, -4, 5, -2,
	5, -4,
}
var yyDef = []int{

	0, -2, 1, 0, 4, 7, 8, 0, 3, 5,
	2, 6,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser/glop.y:28
		{
	    yylex.(*lexer).program = yyS[yypt-0].node
	  }
	case 2:
		//line parser/glop.y:35
		{
	    yyVAL.node = yyS[yypt-1].node
	  }
	case 3:
		//line parser/glop.y:40
		{
	    yyVAL.node = nodes.NewExpr(nil, nil)
	  }
	case 4:
		//line parser/glop.y:45
		{
	    yyVAL.node = yyS[yypt-0].node
	  }
	case 5:
		//line parser/glop.y:52
		{
	    yyVAL.node = nodes.NewExpr(yyS[yypt-0].node, nil)
	  }
	case 6:
		//line parser/glop.y:56
		{
	    yyVAL.node = nodes.NewExpr(yyS[yypt-1].node, yyS[yypt-0].node)
	  }
	case 7:
		//line parser/glop.y:63
		{
	    yyVAL.node = nodes.NewIdentifier(yyS[yypt-0].token)
	  }
	case 8:
		//line parser/glop.y:67
		{
	    yyVAL.node = nodes.NewInteger(yyS[yypt-0].token)
	  }
	}
	goto yystack /* stack new state and value */
}
