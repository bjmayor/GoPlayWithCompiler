# 01 | 理解代码：编译器的前端技术

![06b80f8484f4d88c6510213eb27f2093.webp](../../assets/06b80f8484f4d88c6510213eb27f2093.webp)
这里的“前端”指的是编译器对程序代码的分析和理解过程。
“后端”则是生成目标代码，跟目标机器有关。

## 词法分析（Lexical Analysis）
>词法分析是把程序分割成一个个Token的过程，可以通过构造有限自动机来实现。

辅助工具:Lex

## 语法分析
>语法分析是把程序的结构识别出来，并形成一棵便于由计算机处理的抽象语法树。可以用递归下降的算法来实现。

mac上使用下面的命令可以体验下ast树

`clang -cc1 -ast-dump hello.c`

[js ast体验](https://www.jointjs.com/demos/abstract-syntax-tree)

辅助工具：Yacc、Antlr、JavaCC等。更多列表[参考](https://blog.csdn.net/gongwx/article/details/99645305)

[golang ast](https://zupzup.org/go-ast-traversal/)

## 语义分析
>语义分析是消除语义模糊，生成一些属性信息，让计算机能够依据这些信息生成目标代码。



## 查看AST
mac上使用  -ast-dump

1. 直接使用
```shell
clang -Xclang -ast-dump -fsyntax-only hello.c
```

2. 输出为 JSON 格式（更易读）

```shell
clang -Xclang -ast-dump=json -fsyntax-only hello.c
```
