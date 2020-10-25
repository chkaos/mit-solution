Problem Set 2

1. 1-1 分形渲染

科赫曲线是一种分形。其形态似雪花，又称科赫雪花(Koch snowflake)、科赫星、科赫岛或雪花曲线。

(a) [1 point] What is the height of the recursion tree for rendering a snowflake of LoD n?
  1. log n (final)
  4. 4n 
  (X) n

(b) [2 points] How many nodes are there in the recursion tree at level i, for 0 ≤ i ≤ n?
  4. 3*4^i

(c) [1 point] What is the asymptotic rendering time (triangle count) for a node in the
recursion tree at level i, for 0 ≤ i < n?
  2. O(1)
  4. O((1/3)^i) (final)
  (X) O(1)

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
  (X) n

(g) [1 point] How many nodes are there in the recursion tree at level i, for 0 ≤ i ≤ n?
  4. 3*4^i

(h) [1 point] What is the asymptotic rendering time (line segment count) for a node in the
recursion tree at level i, for 0 ≤ i < n?
  4. O((1/3)^i)
  (X) 0
  Solution: Line segments are only rendered at the last level, so the cost is 0 for levels
0 ≤ i < n.

(i) [1 point] What is the asymptotic rendering time (line segment count) for a node in the
last level n of the recursion tree?
  4. O(3^i)
  Solution: At the last level, each side is turned into one straight line, so the cost is
Θ(1) / node.

(j) [1 point] What is the asymptotic rendering time (line segment count) at each level i
of the recursion tree, for 0 ≤ i < n?
  4. O(3^i)

  Solution: 0, because all the line segments are rendered at the last level.

(k) [1 point] What is the asymptotic rendering time (line segment count) at the last level
n in the recursion tree?
  4. O(4^n)

(l) [1 point] What is the total asymptotic cost for the CPU, when rendering a snowflake
with LoD n using 2D hardware-accelerated rendering?
  4. O(4^n)

(m) [1 point] What is the height of the recursion tree for rendering a snowflake of LoD n?
  1. log n
  Solution: n, because the recursion tree is the same as in the previous part.

(n) [1 point] How many nodes are there in the recursion tree at level i, for 0 ≤ i ≤ n?
  4. 3*4^i

(o) [1 point] What is the asymptotic rendering time (line segment length) for a node in
the recursion tree at level i, for 0 ≤ i < n? Assume that the sides of the initial triangle
have length 1.
  4. O(3^i)

  Solution: Line segments are only rendered at the last level, so the cost is 0 for levels
0 ≤ i < n

(p) [1 point] What is the asymptotic rendering time (line segment length) for a node in
the last level n of the recursion tree?
  4. O((1/3)^i)

(q) [1 point] What is the asymptotic rendering time (line segment length) at each level i
of the recursion tree, for 0 ≤ i < n?
  4. O(4^n)

  Solution: 0, because all the line segments are rendered at the last level.

(r) [1 point] What is the asymptotic rendering time (line segment length) at the last level
n in the recursion tree?
  3. O((4/3)^i)

(s) [1 point] What is the total asymptotic cost for the CPU, when rendering a snowflake
with LoD n using 2D software (not hardware-accelerated) rendering?
  4. O(4^n)
  (X)  O((4/3)^i)

(t) [4 points] What is the total asymptotic cost of rendering a snowflake with LoD n?
Assume that initial triangle’s side length is 1.
  4. O(4^n)

  Solution: Θ(1). See the proof below. Also, intuitively, the snowflake does not overflow the rectangle in the visualizer (or the phone’s screen), so its area must be bounded by a constant. The triangles are adjacent and do not overlap, so the sum of their areas equals the snowflake’s area.

(u) [15 points] Write a succinct proof for your answer using the recursion-tree method
  --

Problem 2-2. [60 points] Digital Circuit Simulation

(a) [8 points] Run the code under the python profiler with the command below, and
identify the method that takes up most of the CPU time. If two methods have similar
CPU usage times, ignore the simpler one.
What is the name of the method with the highest CPU usage?
  _find_min

(b) [6 points] How many times is the method called?
  259964

(c) [8 points] The class containing the troublesome method is implementing a familiar
data structure. What is the tightest asymptotic bound for the worst-case running time
of the method that contains the bottleneck? Express your answer in terms of n, the
number of elements in the data structure.
  3. O(n)

(d) [8 points] If the data structure were implemented using the most efficient method
we learned in class, what would be the tightest asymptotic bound for the worst-case
running time of the method discussed in the questions above?
  4. O(nlog n) X
  O(1).
  Solution: A priority queue implementation based on a binary min-heap takes Θ(1)
time to find the minimum element, since it’s at the top of the heap.

(e) [30 points] Rewrite the data structure class using the most efficient method we learned
in class. Please note that you are not allowed to import any additional Python
libraries and our test will check this.


49分...