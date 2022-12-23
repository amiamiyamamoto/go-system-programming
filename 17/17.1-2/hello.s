    .global main

    .text
main:                               # Cライブラリのスタートアップコードから呼ばれる
    mov     message(%rip), %rdi     # メッセージのポインタをセット
    cal     puts                    # Cライブラリのputs関数呼び出し
    ret                             # Cライブラリコードに戻る

message:
    .asciz "Hello, world"