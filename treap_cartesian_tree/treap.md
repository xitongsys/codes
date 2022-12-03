# treap

树堆（也就是内核中的priority search tree）

每个节点存（key,value），按照value所binary search tree，按照key做priority

插入的时候按照value插入，然后逐层跟parent比较key是否满足heap,如果不满足，则旋转调整

可以通过随机化key,来实现近似平衡的binary search tree。也就是每次插入新数据的时候，随机给个key,插入

删除过程也类似，通过value,查找到要删除的点，然后跟其两个子节点比较priority,选出较大的旋转替换，执导删除的节点成为叶子节点，删除之


[ref](http://wjhsh.net/fusiwei-p-12884254.html)