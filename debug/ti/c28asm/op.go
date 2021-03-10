package c28asm

type Op uint16

const (
	UNKNOWN Op = iota
	ABORTI
	ABS
	ABSTC
	ADD
	ADDB
	ADDCL
	ADDCU
	ADDL
	ADDU
	ADDUL
	ADRK
	AND
	ANDB
	ASP
	ASR
	ASR64
	ASRL
	B
	BANZ
	BAR
	BF
	C27MAP
	C27OBJ
	C28ADDR
	C28MAP
	C28OBJ
	CLRC
	CMP
	CMP64
	CMPB
	CMPL
	CMPR
	CSB
	DEC
	DINT
	DMAC
	DMOV
	EALLOW
	EDIS
	EINT
	ESTOP0
	ESTOP1
	FFC
	FLIP
	IACK
	IDLE
	IMACL
	IMPYAL
	IMPYL
	IMPYSL
	IMPYXUL
	IN
	INC
	INTR
	IRET
	LB
	LC
	LCR
	LOOPNZ
	LOOPZ
	LPADDR
	LRET
	LRETE
	LRETR
	LSL
	LSL64
	LSLL
	LSR
	LSR64
	LSRL
	MAC
	MAX
	MAXCUL
	MAXL
	MIN
	MINCUL
	MINL
	MOV
	MOVA
	MOVAD
	MOVB
	MOVDL
	MOVH
	MOVL
	MOVP
	MOVS
	MOVU
	MOVW
	MOVX
	MOVZ
	MPY
	MPYA
	MPYB
	MPYS
	MPYU
	MPYXU
	NASP
	NEG
	NEG64
	NEGTC
	NOP
	NORM
	NOT
	OR
	ORB
	OUT
	POP
	PREAD
	PUSH
	PWRITE
	QMACL
	QMPYAL
	QMPYL
	QMPYSL
	QMPYUL
	QMPYXUL
	ROL
	ROR
	RPT
	SAT
	SAT64
	SB
	SBBU
	SBF
	SBRK
	SETC
	SFR
	SPM
	SQRA
	SUB
	SUBB
	SUBBL
	SUBCU
	SUBCUL
	SUBLSUBR
	SUBRL
	SUBU
	SUBUL
	TBIT
	TCLR
	TEST
	TRAP
	TSET
	UOUT
	XB
	XBANZ
	XCALL
	XMAC
	XMACD
	XOR
	XPREAD
	XPWRITE
	XRET
	XRETC
	ZALR
	ZAP
	ZAPA
)
