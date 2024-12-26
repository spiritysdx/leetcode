# 前后指针

right先走n步，然后left再指向dummy节点，登right.Next为nil时，left到倒数n+1的节点

一般来说如果会修改/删除头节点的话，加个 dummy 会更加方便。
