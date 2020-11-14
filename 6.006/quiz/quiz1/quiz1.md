假装每页都写了名字 得2分

Problem 2. Asymptotics & Recurrences [20 points] (3 parts)

(a) 3 5 1 8 4 6 7 2   X 1 5 3 8 7 6 4 2
(b) T(n) = T(n/3) + T(2n/3) + Θ(n)
    T(n/3) = T(n/9) + T(2n/9) + Θ(n/3)
    T(2n/3) = T(2n/9) + T(4n/9) + Θ(2n/3)
    T(3) = T(1) + T(2) + O(3)

    T(n) = O(3n-1) => O(n) 等比求和?

    X
    Draw recursion tree. At each level, do Θ(n) work. Number of levels is
    log3/2 n = Θ(lg n), so guess T(n) = Θ(n lg n) and use the substitution method to
    verify guess.

(c) T(n) = log n + T(√n)  skip
    T(√n) = log (√n) + log (√√n)

    T(n) = Θ(log n).
    没有 base case
    画一颗递归树,假设高度为h 第一层是 节点值是n, cost 是 logN, 有一个子节点(值是√n, cost以此类推)
    最后一层(n to the 1/2^h 约等于1)

Problem 3. True/False [18 points] (9 parts)

Circle (T)rue or (F)alse. You don’t need to justify your choice.
(a) F [2 points] Binary insertion sorting (insertion sort that uses binary search to find
each insertion point) requires O(n log n) total operations.
(b) T [2 points] In the merge-sort execution tree, roughly the same amount of work is
done at each level of the tree.
(c) T [2 points] In a BST, we can find the next smallest element to a given element in
O(1) time.  X false O(h)

(d) T [2 points] In an AVL tree, during the insert operation there are at most two
rotations needed.
(e) F [2 points] Counting sort is a stable, in-place sorting algorithm.
(f) T [2 points] In a min-heap, the next largest element of any element can be found
in O(log n) time.
X false O(n)

(g) F [2 points] The multiplication method satisfies the simple uniform hashing assumption.
(h) T [2 points] Double hashing satisfies the uniform hashing assumption.
X false 
The notes state that double hashing ‘comes close.’ Double
hashing only provides n2 permutations, not n!.

(i) T [2 points] Python generators can be used to iterate over potentially infinite countable sets with O(1) memory

12 分...

Problem 4. Peak Finding (again!) [20 points] (2 parts)
(a) 时间复杂度为O(nlogm)  X O(n)
(b) skip ..

Problem 5. Who Let The Zombies Out? [20 points] (2 parts)
有点意思的考试题, 外星人污染了地球水源, 人喝了变僵尸 C是城市列表, n个城市可能被污染, k个城市确定被污染, 最少的志愿者喝被污染的水(混合)
最小代价解决固定问题.

决策树推导.

Problem 6. Shopping Madness [20 points] (3 parts)
Ben Bitdiddle was peer-pressured into signing up for the tryouts in a shopping reality TV show,
and he needs your help to make it past the first round. In order to qualify, Ben must browse a
store’s inventory, which has N items with different positive prices P[1], P[2], . . . , P[N], and the
challenge is to spend exactly S dollars on exactly K items, where K is a small even integer. Ben
can buy the same item multiple times. For example, 3 brooms and 2 wizard hats add up to 5 items.
In your solutions below, you may use a subroutine MULTISETS(k, T) which iterates over all the
k-element multisets (like subsets, except the same elements can show up multiple times) of a set
T, in time O(k · |T|^k), using O(k) total space. Note that if your code holds onto the results of
MULTISETS, it may end up using more than O(k) space.

Problem 7. When I Was Your Age. . . [20 points] (2 parts)