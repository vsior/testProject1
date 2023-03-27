%include "stud_io.inc"
global _start

section .text
_start: mov eax, 0
again: PRINT "Hello"
        PUTCHAR 10
        inc eax
        cmp eax, 5
        jl again
        FINISH

section .bss
welwsg resb 16
welwsg2 resb 16
welwsg3 resb 16000
