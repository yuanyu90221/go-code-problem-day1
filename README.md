# go-code-problem-day1

## 題目描述

給定一個由小至大排序過的陣列 nums，以及一個整數 target，找出此 target 適當的插入點 (index)，使陣列維持由小至大的排序，並返回其 index。

若有多個可能的 index，需 return 最小的 index
時間複雜度要求為 O(logn)
Example 1:

Input: nums = [1, 3, 5, 7, 9], target = 6
Output: 3
Explanation: 6 的插入點為 5 和 7 之間，而此插入點的 index 為 3。
Example 2:

Input: nums = [1, 3, 3, 3, 3], target = 3
Output: 1
Explanation: 3 的插入點為 1 和 3 之間，而此插入點的 index  為 1

## 題目分析

給定一個 排序過的array nums 以及 一個整數 target

找出此 target 插入的index

由於array nums是排序過的 因此可以使用 BinarySearch 找出對應的 index 為 O(logN)