Problem Set 3

3-1. [45 points] Range Queries
(a) 3. Binary Search Tree (BST)  4. AVL Trees(选这个) 5. B-Trees(这个没学过好像, 这是构建数据库索引的数据结构, 本来想选这个..?)
(b) 3.O(log N)
(c) 3.O(log N)
(d) 3.O(log N)

(e) 3. RANK(h) − RANK(l) + 1
(f) 5. RANK(h) − RANK(l)
(g) 3. RANK(h) − RANK(l) + 1
(h) 8. RANK(h) + RANK(l)    X RANK(h) - RANK(l)  -1

(i) 4. the number of nodes in the subtree rooted at node
(j) 1. O(1)   X   -1
Solution: The number of nodes in a subtree is at most N, so the γ field in each node needs to be O(log N)-bits wide to be able to store numbers between 0 and N.
(k) 2. 1
(l) 3. 3
(m) 4. 6
(n) 4. 10
(o) 1-6 TRUE X ROTATE-LEFT, ROTATE-RIGHT, and REBALANCE  -3
(p) 3. O(log N)

(q) 2. lowest common ancestor
(r) 3. O(log N)
(s) 8. O(1)  X O(log N) + O(L)  -1
(t) 8. O(log N) + O(L)
(u) .... 论证题  -20

Problem 3-2. [55 points] Digital Circuit Layout
(a) intersects
(b) 187590314
(c) 1-4 true 5 false  X 34 false 125 true  -2
(d) 1. the wire is added to the range index
(e) 3. a range index query is performed   X 2. the wire is removed from the range index
(f) 3. a range index query is performed   X 4. nothing happens
(g) 3. a range index query is performed

...做不下去了,下面的不计分
(h) 2. the range index holds all the horizontal wires “stabbed” by the sweep line
(i) 2. the y coordinate of the wire’s midpoint

31..