# 05 | 语法分析（三）：实现一门简单的脚本语言
## 增加所需要的语法规则

我们用扩展巴科斯范式（EBNF）写出下面的语法规则：

```
programm: statement+;

statement
: intDeclaration
| expressionStatement
| assignmentStatement
;
```

**变量声明语句**以 int 开头，后面跟标识符，然后有可选的初始化部分，也就是一个等号和一个表达式，最后再加分号：

```
intDeclaration : 'int' Id ( '=' additiveExpression)? ';';
```

**表达式语句** 目前只支持加法表达式，未来可以加其他的表达式，比如条件表达式，它后面同样加分号：

```
expressionStatement : additiveExpression ';';
```

**赋值语句**是标识符后面跟着等号和一个表达式，再加分号：

```
assignmentStatement : Identifier '=' additiveExpression ';';
```

为了在表达式中可以使用变量，我们还需要把 primaryExpression 改写，除了包含整型字面量以外，还要包含标识符和用括号括起来的表达式：

```
primaryExpression : Identifier| IntLiteral | '(' additiveExpression ')';
```

## 让脚本语言支持变量
可以考虑用一个hashmap 存下变量 和 他的值。
