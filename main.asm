"".main STEXT size=138 args=0x0 locals=0x58 funcid=0x0
	0x0000 00000 (main.go:8)	TEXT	"".main(SB), ABIInternal, $88-0
	0x0000 00000 (main.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:8)	PCDATA	$0, $-2
	0x000d 00013 (main.go:8)	JLS	128
	0x000f 00015 (main.go:8)	PCDATA	$0, $-1
	0x000f 00015 (main.go:8)	SUBQ	$88, SP
	0x0013 00019 (main.go:8)	MOVQ	BP, 80(SP)
	0x0018 00024 (main.go:8)	LEAQ	80(SP), BP
	0x001d 00029 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x001d 00029 (main.go:8)	FUNCDATA	$2, "".main.stkobj(SB)
	0x001d 00029 (main.go:9)	XORPS	X0, X0
	0x0020 00032 (main.go:9)	MOVUPS	X0, ""..autotmp_11+64(SP)
	0x0025 00037 (main.go:9)	LEAQ	type.string(SB), AX
	0x002c 00044 (main.go:9)	MOVQ	AX, ""..autotmp_11+64(SP)
	0x0031 00049 (main.go:9)	LEAQ	""..stmp_0(SB), AX
	0x0038 00056 (main.go:9)	MOVQ	AX, ""..autotmp_11+72(SP)
	0x003d 00061 (<unknown line number>)	NOP
	0x003d 00061 ($GOROOT/src/fmt/print.go:274)	MOVQ	os.Stdout(SB), AX
	0x0044 00068 ($GOROOT/src/fmt/print.go:274)	LEAQ	go.itab.*os.File,io.Writer(SB), CX
	0x004b 00075 ($GOROOT/src/fmt/print.go:274)	MOVQ	CX, (SP)
	0x004f 00079 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 8(SP)
	0x0054 00084 ($GOROOT/src/fmt/print.go:274)	LEAQ	""..autotmp_11+64(SP), AX
	0x0059 00089 ($GOROOT/src/fmt/print.go:274)	MOVQ	AX, 16(SP)
	0x005e 00094 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 24(SP)
	0x0067 00103 ($GOROOT/src/fmt/print.go:274)	MOVQ	$1, 32(SP)
	0x0070 00112 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $0
	0x0070 00112 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0075 00117 (main.go:9)	MOVQ	80(SP), BP
	0x007a 00122 (main.go:9)	ADDQ	$88, SP
	0x007e 00126 (main.go:9)	RET
	0x007f 00127 (main.go:9)	NOP
	0x007f 00127 (main.go:8)	PCDATA	$1, $-1
	0x007f 00127 (main.go:8)	PCDATA	$0, $-2
	0x007f 00127 (main.go:8)	NOP
	0x0080 00128 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0085 00133 (main.go:8)	PCDATA	$0, $-1
	0x0085 00133 (main.go:8)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 71 48  dH..%....H;a.vqH
	0x0010 83 ec 58 48 89 6c 24 50 48 8d 6c 24 50 0f 57 c0  ..XH.l$PH.l$P.W.
	0x0020 0f 11 44 24 40 48 8d 05 00 00 00 00 48 89 44 24  ..D$@H......H.D$
	0x0030 40 48 8d 05 00 00 00 00 48 89 44 24 48 48 8b 05  @H......H.D$HH..
	0x0040 00 00 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 48  ....H......H..$H
	0x0050 89 44 24 08 48 8d 44 24 40 48 89 44 24 10 48 c7  .D$.H.D$@H.D$.H.
	0x0060 44 24 18 01 00 00 00 48 c7 44 24 20 01 00 00 00  D$.....H.D$ ....
	0x0070 e8 00 00 00 00 48 8b 6c 24 50 48 83 c4 58 c3 90  .....H.l$PH..X..
	0x0080 e8 00 00 00 00 e9 76 ff ff ff                    ......v...
	rel 2+0 t=25 type.string+0
	rel 2+0 t=25 type.*os.File+0
	rel 5+4 t=17 TLS+0
	rel 40+4 t=16 type.string+0
	rel 52+4 t=16 ""..stmp_0+0
	rel 64+4 t=16 os.Stdout+0
	rel 71+4 t=16 go.itab.*os.File,io.Writer+0
	rel 113+4 t=8 fmt.Fprintln+0
	rel 129+4 t=8 runtime.morestack_noctxt+0
os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (<autogenerated>:1)	MOVQ	""..this+8(SP), AX
	0x0005 00005 (<autogenerated>:1)	MOVQ	(AX), AX
	0x0008 00008 (<autogenerated>:1)	MOVQ	AX, ""..this+8(SP)
	0x000d 00013 (<autogenerated>:1)	XORPS	X0, X0
	0x0010 00016 (<autogenerated>:1)	MOVUPS	X0, "".~r0+16(SP)
	0x0015 00021 (<autogenerated>:1)	JMP	os.(*file).close(SB)
	0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
	0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
	rel 22+4 t=8 os.(*file).close+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=24 type.[]interface {}+0
	rel 0+0 t=24 type.error+0
	rel 0+0 t=24 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
go.string."!... Hello World ...!" SRODATA dupok size=21
	0x0000 21 2e 2e 2e 20 48 65 6c 6c 6f 20 57 6f 72 6c 64  !... Hello World
	0x0010 20 2e 2e 2e 21                                    ...!
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
""..stmp_0 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 15 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."!... Hello World ...!"+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=1 os.(*File).Write+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
"".main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
	rel 16+8 t=1 type.[1]interface {}+0
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........