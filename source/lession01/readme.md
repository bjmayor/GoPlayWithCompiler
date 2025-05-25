# 查看AST
mac上使用  -ast-dump

1. 直接使用
```shell
clang -Xclang -ast-dump -fsyntax-only hello.c
```

2. 输出为 JSON 格式（更易读）

```shell
clang -Xclang -ast-dump=json -fsyntax-only hello.c
```

javascript ast可视化地址。
https://www.jointjs.com/demos/abstract-syntax-tree
