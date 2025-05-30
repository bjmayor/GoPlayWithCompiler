## 正则文法和有限自动机：纯手工打造词法分析器
![77f7c8384906864229808eb82f06b158.png](evernotecid://9FAD049A-8342-475F-9BFD-2B663B1AF63C/appyinxiangcom/2980418/ENResource/p8476)

示例程序分析以下语句
```
age >= 45
int age = 40
2 + 3*5
intA = 10
```

作者的[实现代码](https://github.com/RichardGong/PlayWithCompiler/blob/master/lab/craft/SimpleLexer.java)

>要实现一个词法分析器，首先需要写出每个词法的正则表达式，并画出有限自动机，之后，只要用代码表示这种状态迁移过程就可以了。

golang可以参考json的实现方式。
main.go 是用AI实现的类似版本。[main.go](https://github.com/bjmayor/GoPlayWithCompiler/blob/master/source/lession02/main.go)
