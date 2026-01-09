# Day 9

In this doc I want to compare different approaches to solve AoC 2025 day 9 part 2.


## Context
**Problem:** https://adventofcode.com/2025/day/9

Input provides a set of points on a grid which form a closed polygon. Problem asks to find the biggest rectangle formed by provided points which is fully inside the polygon.


## Even-odd rule
Even-odd rule allows to check if a  point is inside the polygon by casting a beam horizontally from the point to some outside point and counting number of edges crossed by the beam. If the number is even - the point is outside else the point is inside. 

### The issue
Since this is a grid problem some beams lay on an edge. This create controvercy.
Due to this issue it is inposible to use `Even-odd` rule only to solve the problem.


#### Example 1
```

0 +--1---+ 2
  |      |
  |      |
  |      |
  |      |
  +------+

```

#### Example 2

```
         +------+
         |      |
         |      |
0 +---1--+  2   + 3
  |             |
  |             |
  |             |
  +-------------+
```



## Flood-fill 
Since the polygon is closed running BFS from outside with polygon edges as walls gives all outside points. Afterwards it's easy to check that there is no outside points in a rectangle.

**TC:** O(max(X) * max(Y)) * O(CheckPointOnEdge)


## Compression

Coordinates of Polygon in the problem contains big number like 10^5. Which makes grid as big as 10^10. 
But there are only 100 nodes in the polygon. So grid can be comperessed into a size of 100x100.
Which is much smaller.

If we apply compression to `Flood-fill` algorithm it changes time complexity dependancy from grid size to number of nodes in the polygon.

**TC:** O(N^2) * O(CheckPointOnEdge)

### TC analysis
Time complexity in this case can be analysed furser. 
CheckPointOnEdge is using linear search with hash map optimisation. TC worst case is O(N) when all points grouped on one line.

For BFS TC is N^2 if all nodes have unique X and Y.

So we can see that worst case for CheckPointOnEdge is a bast for BFS. 

So actual TC is O(N^2)

## Rectangle in Polygon rule
To check that one polygon contains another we shoud check that all its nodes are inside the other and their edges neter crosses.

**TC:** O(NumberOfRectangeles * 4 * NumberOfEdges) = O(N^3)




## Recommendation

Flood-fill + compression was enough to solve the problem.



