Problem Set 2

1. 1-1 分形渲染

科赫曲线是一种分形。其形态似雪花，又称科赫雪花(Koch snowflake)、科赫星、科赫岛或雪花曲线。

(a) [1 point] What is the height of the recursion tree for rendering a snowflake of LoD n?
  1. log n (final)
  4. 4n

(b) [2 points] How many nodes are there in the recursion tree at level i, for 0 ≤ i ≤ n?
  4. 3*4^i

(c) [1 point] What is the asymptotic rendering time (triangle count) for a node in the
recursion tree at level i, for 0 ≤ i < n?
  2. O(1)

(d) [1 point] What is the asymptotic rendering time (triangle count) at each level i of the
recursion tree, for 0 ≤ i < n?
  4. O(4^i)

(e) [2 points] What is the total asymptotic cost for the CPU, when rendering a snowflake
with LoD n using 3D hardware-accelerated rendering?
  4. O(4^n)


2D

(f) [1 point] What is the height of the recursion tree for rendering a snowflake of LoD n
using 2D hardware-accelerated rendering?
  1. log n

(g) [1 point] How many nodes are there in the recursion tree at level i, for 0 ≤ i ≤ n?
  4. 3*4^i

(h) [1 point] What is the asymptotic rendering time (line segment count) for a node in the
recursion tree at level i, for 0 ≤ i < n?

Problem 2-2. [60 points] Digital Circuit Simulation