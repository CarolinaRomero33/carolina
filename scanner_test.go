package andrea_test

import (
	"github.com/CarolinaRomero33/carolina"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok andrea.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: andrea.EOF},
		{s: `#`, tok: andrea.ILLEGAL, lit: `#`},
		{s: ` `, tok: andrea.WS, lit: " "},
		{s: "\t", tok: andrea.WS, lit: "\t"},
		//{s: "\n", tok: andrea.WS, lit: "\n"},
		{s: "\r", tok: andrea.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: andrea.ASTERISK, lit: "*"},
		{s: `:`, tok: andrea.DosPunt, lit: ":"},
		{s: `(`, tok: andrea.ParDer, lit: "("},
		{s: `)`, tok: andrea.ParIsq, lit: ")"},
		{s: `;`, tok: andrea.PuntCom, lit: ";"},
		{s: `=`, tok: andrea.Asignaicon, lit: "="},
		//{s: `[`, tok: andrea.AgruppadorDER, lit: "]"},
		//{s: `[`, tok: andrea.AgruppadorISQ, lit: "]"},
		//{s: `\n`, tok: andrea.SaltoDeL, lit: "\n"},

		// Identifiers
		{s: `foo`, tok: andrea.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: andrea.IDENT, lit: `Zx12_3U_`},

		//Keywords
		{s: `zIF`, tok: andrea.zIF, lit: "zIF"},
		{s: `zELSEIF`, tok: andrea.zELSEIF, lit: "zELSEIF"},
		{s: `zMAIN`, tok: andrea.zMAIN, lit: "zMAIN"},
		{s: `zELSE`, tok: andrea.zELSE, lit: "zELSE"},
		{s: `zSWITCH`, tok: andrea.zSWITCH, lit: "zSWITCH"},
		{s: `zFUNC`, tok: andrea.zFUNC, lit: "zFUNC"},
		{s: `zINT`, tok: andrea.zINT, lit: "zINT"},
		{s: `zEND`, tok: andrea.zEND, lit: "zEND"},
		{s: `VAR`, tok: andrea.VAR, lit: "VAR"},
		{s: `zWHILE`, tok: andrea.zWHILE, lit: "zWHILE"},
		{s: `zDO`, tok: andrea.zDO, lit: "zDO"},
		{s: `zTRY`, tok: andrea.zTRY, lit: "zTRY"},
		{s: `zCATCH`, tok: andrea.zCATCH, lit: "zCATCH"},
		{s: `zFINALLY`, tok: andrea.zFINALLY, lit: "zFINALLY"},
		{s: `zPOINT`, tok: andrea.zPOINT, lit: "zPOINT"},
		{s: `zFIXED`, tok: andrea.zFIXED, lit: "zFIXED"},
		{s: `zUNCHECKED`, tok: andrea.zUNCHECKED, lit: "zUNCHECKED"},
		{s: `zNEW`, tok: andrea.zNEW, lit: "zNEW"},
		{s: `zFOREACH`, tok: andrea.zFOREACH, lit: "zFOREACH"},
		{s: `zIMPORT`, tok: andrea.zIMPORT, lit: "zIMPORT"},


		//Bonifaz
		//{s: `MCCREATE`, tok: andrea.MCCREATE, lit: "MCCREATE"},
		//
		//{s: `MCUSE`, tok: andrea.MCUSE, lit: "MCUSE"},
		//
		//{s: `MCCREATETABLE`, tok: andrea.MCCREATETABLE, lit: "MCCREATETABLE"},
		//
		//{s: `DATABASE`, tok: andrea.DATABASE, lit: "DATABASE"},




	}

	for i, tt := range tests {
		s := andrea.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q zt=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q zt=%q", i, tt.s, tt.lit, lit)
		}
	}
}
