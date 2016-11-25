package andrea

import (
	"fmt"
	"io"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct{}


type zifstatement struct{}


type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}


func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}


	//Erasmo

	if tok, lit := p.scanIgnoreWhitespace(); tok != zIF {
		return nil, fmt.Errorf("encontro %q, esperaba zif", lit)
	}
	return stmt,nil
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != zELSE {
		return nil, fmt.Errorf("encontro %q, esperaba zELSE", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//Suma global.

	if tok, lit := p.scanIgnoreWhitespace(); tok != VAR {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != VAR {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba zfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//suma normal

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba zfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	// CICLO zFOR

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFOR {
		return nil, fmt.Errorf("encontro %q, esperaba zfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba zfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zELSEIF

	if tok, lit := p.scanIgnoreWhitespace(); tok != zELSEIF {
		return nil, fmt.Errorf("encontro %q, esperaba zelseif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zswitch

	if tok, lit := p.scanIgnoreWhitespace(); tok != zSWITCH {
		return nil, fmt.Errorf("encontro %q, esperaba zswitch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zBREAK {
		return nil, fmt.Errorf("encontro %q, esperaba zbreak", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zwhile

	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba zwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zdo ciclo
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zDO {
		return nil, fmt.Errorf("encontro %q, esperaba zDO", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba zwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zTRY
	if tok, lit := p.scanIgnoreWhitespace(); tok != zTRY {
		return nil, fmt.Errorf("encontro %q, esperaba ztry", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zCATCH {
		return nil, fmt.Errorf("encontro %q, esperaba zcatch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zfinally

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFINALLY {
		return nil, fmt.Errorf("encontro %q, esperaba zfinally", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zpoint

	if tok, lit := p.scanIgnoreWhitespace(); tok != zPOINT {
		return nil, fmt.Errorf("encontro %q, esperaba zPOINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zNEW {
		return nil, fmt.Errorf("encontro %q, esperaba zNEW", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zFIXED

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFIXED {
		return nil, fmt.Errorf("encontro %q, esperaba zFIXED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zUNCHECKED

	if tok, lit := p.scanIgnoreWhitespace(); tok != zUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba zUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba zUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	////ARREGLO
	//
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zNEW {
		return nil, fmt.Errorf("encontro %q, esperaba znew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zINT {
		return nil, fmt.Errorf("encontro %q, esperaba zint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//
	////zforeach

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFOREACH {
		return nil, fmt.Errorf("encontro %q, esperaba zFOREACH", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//PAQUETES
	if tok, lit := p.scanIgnoreWhitespace(); tok != zIMPORT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//funciones

	if tok, lit := p.scanIgnoreWhitespace(); tok != zFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba zfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//zmain

	if tok, lit := p.scanIgnoreWhitespace(); tok != zMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba zmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != zEND {
		return nil, fmt.Errorf("encontro %q, esperaba zend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	return stmt, nil

}
func (p *Parser) ParsezIF() (*zifstatement, error) {
	stmt := &zifstatement{}
	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != zIF {
		return nil, fmt.Errorf("encontro %q, esperaba zif", lit)
	}
	return stmt, nil
}


func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

/*
// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Message  string
	Found    string
	Expected []string
}


// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// newParseError returns a new instance of ParseError.
func newParseError(found string, expected []string) *ParseError {
	return &ParseError{Found: found, Expected: expected}
}

// Error returns the string representation of the error.
func (e *ParseError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s at line %d, char %d", e.Message)
	}
	return fmt.Sprintf("found %s, expected %s at line %d, char %d", e.Found)
}
*/
