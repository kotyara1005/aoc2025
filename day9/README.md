# Day 9

In this doc I want to compare different approaches to solve AoC 2025 day 9 part 2.

**Problem:** https://adventofcode.com/2025/day/9

## Even-odd rule
This is a simple rule. To check if point is inside or outside we send a beam horizontally to outside point and count number of edges it crosses.
If this number is even point is outside else point is inside.

### The issue
Since this is a grid problem it is likely that beam lays on an edge. 

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
Use BFS from an outside corner to find all outside space and then check if there is any inside of a rectangle.


## Compression

Polygon in the problem contains big number like 10^5. Which makes grid as big as 10^10. 
But there are only 100 nodes in the polygon. That means that grid can be comperessed into a size of 100x100.
Which is much smaller.


## Rectangle in Polygon rule
To check that one polygon contains another we shoud check that all its nodes are inside the other and their edges neter crosses.








