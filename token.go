package andrea

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	AgruppadorDER // [
	AgruppadorISQ // ]

	// Keywords
	SELECT
	FROM
	zIF
	zSWITCH
	zEND
	zINT // Tipodedato
	zFUNC
	zMAIN
	zELSE
	zFOR
	zELSEIF
	VAR // variables
	zBREAK
	zWHILE
	zDO
	zTRY
	zCATCH
	zFINALLY
	zPOINT
	zNEW
	zFIXED
	zFOREACH
	zUNCHECKED
	zIMPORT

	//Bonifaz

	MCCREATE
	DATABASE
	MCUSE
	MCCREATETABLE

)
