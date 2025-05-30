## 语法分析（一）：纯手工打造公式计算器
主要是 语法分析的原理和递归下降算法，上下文无关文法。
本课稍有单独，需要看原文。
[作者源码](https://github.com/RichardGong/PlayWithCompiler/blob/master/lab/craft/SimpleCalculator.java)

自己手动实现一遍有助于理解。
这节引入了 几个概念：
- 上下文无关文法

正则文法无法表达 算数表达式 这种需要递归表达的形式。
所以引入了 上下文无关文法。它可以递归调用自己。

additiveExpression
    :   multiplicativeExpression
    |   additiveExpression Plus multiplicativeExpression
    ;

multiplicativeExpression
    :   IntLiteral
    |   multiplicativeExpression Star IntLiteral
    ;

加法表达式 或者是乘法表达式， 或者是 加法表达式 + 乘法表达式。
而 乘法表达式 或者是 整形字面量 或者是 乘法表达式 * 整形字面量。

递归也是也有问题：
additiveExpression
    :   IntLiteral
    |   additiveExpression Plus IntLiteral
    ;

这个就是左递归。
加法表达式 或者 是整形字面量 或者是 加法表达式 + 整形字面量。

 2+3 解释的时候
 1. 发现 2+3 不是 整形字面量。
 2. 继续判断 2+3 是否是加法表达式。
 3. 回到 1. 继续判断 2+3 是否是 加法表达式。死循环了。

**左递归是递归下降算法无法处理的，这是递归下降算法最大的问题。**

additiveExpression
    :   IntLiteral
    |   IntLiteral Plus additiveExpression
    ;

换个顺序：
 2+3 解释的时候
 1. 发现 2+3 不是 整形字面量。
 2. 继续判断 2+3 是否是加法表达式。
 3. 回到 1. 继续判断 2+3 是否是 加法表达式。死循环了。

作者给出了一个思路：

additiveExpression
    :   multiplicativeExpression
    |   multiplicativeExpression Plus additiveExpression
    ;

multiplicativeExpression
    :   IntLiteral
    |   IntLiteral Star multiplicativeExpression
    ;

2+3*5 解释结果：
additiveExpression
     / | \
    /  |  \
   /   |   \
additiveExp  Plus  multiplicativeExp
|                /   |   \
multiplicativeExp     /    |    \
|           Int(3) Star Int(5)
Int(2)

他的问题是 结合顺序出问题了。
2+3+5 会解释成
additiveExpression
     / | \
    /  |  \
   /   |   \
additiveExp  Plus  additiveExp
|                /   |   \
multiplicativeExp     /    |    \
|           Int(3) Plus Int(5)
Int(2)

2+（3+5）

第二个概念是表达式求值。就是遍历树的过程。
