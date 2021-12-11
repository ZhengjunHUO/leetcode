# LeetCode in Go

## LeetCode Algorithm/Solution in Golang

1. One folder One problem. Each folder contains a txt file which describe the problem. 
- The solution easy to figure out (for me) is in naive*.go, the name may contain some info about method used. For the rest *.go (eg. solution.go) is the solution i learned after research (especially thank to all who contributed to the “Discuss” section for those brillant ideas !). 
- For each problem, the final solution should reach the best possible time & space complexity, before move on to the next problem. 
- Commented in Chinese (we can explain the same idea with less words).
	
2. About Go, normally I only import fmt to print out the result to check quickly if the algorithm works as expected. For some problems which engage data types, i use my [homemade thread-safe go data type](https://github.com/ZhengjunHUO/godtype).
Other [useful tools](https://github.com/ZhengjunHUO/gtoolkit) built during this adventure and will surely come into handy later.

3. I didn’t specifically note down the difficulty of the problems, because it’s quite subjective and depends on your experience on some type of problem: you can easily solve some “hard” problems if you’ve figured out some sort of “template”  of this type in 10 minutes, but cat get stuck for hours by some “easy” problems you are not quite familiar or you've never seen before (like 初見殺し boss in game :) ?)

| # | Title | Tags | 描述 |
|---| ----- | -------- | -------- |
|0001| [Two Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0001_Two_Sum) | Array  HashTable | 找出和为目标值的两个元素 |
|0002| [Add Two Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0002_Add_Two_Numbers) | LinkedList  Math  Recursion | 数组形式表示的两数求和 |
|0003| [Longest Substring Without Repeating Characters](https://github.com/ZhengjunHUO/leetcode/tree/main/0003_Longest_Substring_Without_Repeating_Characters) | HashTable  SlidingWindow  String  TwoPointers | 求不含重复字母的子串的长度 |
|0004| [Median of Two Sorted Arrays](https://github.com/ZhengjunHUO/leetcode/tree/main/0004_Median_of_Two_Sorted_Arrays) | Array  BinarySearch  DivideAndConquer | 求两个有序数列的中位数 |
|0005| [Longest Palindromic Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/0005_Longest_Palindromic_Substring) | String  DynamicProgramming | 求最长回文子串 |
|0007| [Reverse Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0007_Reverse_Integer) | Math | 整数反向 |
|0008| [String to Integer (atoi)](https://github.com/ZhengjunHUO/leetcode/tree/main/0008_String_to_Integer) | Math  String | 字符串转整数实现 |
|0009| [Palindrome Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0009_Palindrome_Number) | Math | 判断整数是否为回文 |
|0010| [Regular Expression Matching](https://github.com/ZhengjunHUO/leetcode/tree/main/0010_Regular_Expression_Matching) | Backtracking  DynamicProgramming  String | 正则表达匹配实现 | 
|0011| [Container With Most Water](https://github.com/ZhengjunHUO/leetcode/tree/main/0011_Container_With_Most_Water) | Array  TwoPointers | 确定容器的两个边界使其含水量最大化 |
|0012| [Integer to Roman](https://github.com/ZhengjunHUO/leetcode/tree/main/0012_Integer_to_Roman) | Math  String | 整型转罗马数字 |
|0013| [Roman to Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0013_Roman_to_Integer) | Math  String | 罗马数字转整型 |
|0014| [Longest Common Prefix](https://github.com/ZhengjunHUO/leetcode/tree/main/0014_Longest_Common_Prefix) | String | 最长公共前缀 |
|0015| [Three Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0015_Three_Sum) | Array  TwoPointers | 求所有和为0的三元组 |
|0016| [3Sum Closest](https://github.com/ZhengjunHUO/leetcode/tree/main/0016_3Sum_Closest) | Array  TwoPointers | 求和最接近目标值的三元组 |
|0017| [Letter Combinations of a Phone Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0017_Letter_Combinations_of_a_Phone_Number) | Backtracking  HashTable  String | 根据电话号码找到对应的字母组合 |
|0018| [Four Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0018_Four_Sum) | Array  HashTable  TwoPointers | 找出所有和为目标值的四元组 |
|0019| [Remove Nth Node From End of List](https://github.com/ZhengjunHUO/leetcode/tree/main/0019_Remove_Nth_Node_From_End_of_List) | LinkedList  TwoPointers | 从列表中移除倒数第N个元素 |
|0020| [Valid Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0020_Valid_Parentheses) | Stack  String | 检查括号是否正确闭合 |
|0021| [Merge Two Sorted Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0021_Merge_Two_Sorted_Lists) | LinkedList  Recursion | 融合两个排好序的链表 |
|0022| [Generate Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0022_Generate_Parentheses) | Backtracking  DynamicProgramming  String | 生成括号对 |
|0023| [Merge k Sorted Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0023_Merge_k_Sorted_Lists) | DivideAndConquer  Heap  LinkedList  MergeSort | 融合k个排好序的列表 |
|0024| [Swap Nodes in Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/0024_Swap_Nodes_in_Pairs) | LinkedList  Recursion | 链表相邻两个结点对调 |
|0025| [Reverse Nodes in k-Group](https://github.com/ZhengjunHUO/leetcode/tree/main/0025_Reverse_Nodes_in_kGroup) | LinkedList  Recursion | 链表中k个元素一组反转组内元素 |
|0026| [Remove Duplicates from Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0026_Remove_Duplicates_from_Sorted_Array) | Array  TwoPointers | 移除数组中多余的元素(in place) |
|0027| [Remove Element](https://github.com/ZhengjunHUO/leetcode/tree/main/0027_Remove_Element) | Array  TwoPointers | 移除数组中指定元素(in place) |
|0028| [Implement strStr()](https://github.com/ZhengjunHUO/leetcode/tree/main/0028_Implement_strStr) | String  TwoPointers | 找出pattern在字符串中第一次出现的位置 |
|0030| [Substring with Concatenation of All Words](https://github.com/ZhengjunHUO/leetcode/tree/main/0030_Substring_with_Concatenation_of_All_Words) | HashTable  String  TwoPointers | 找出内容为"以任意顺序拼接的所有给定单词"的子串 |
|0031| [Next Permutation](https://github.com/ZhengjunHUO/leetcode/tree/main/0031_Next_Permutation) | Array  TwoPointers | 按字典顺序寻找当前数组的后一个排列 |
|0032| [Longest Valid Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0032_Longest_Valid_Parentheses) | DynamicProgramming  Stack  String | 最长的合法括号子串 |
|0033| [Search in Rotated Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0033_Search_in_Rotated_Sorted_Array) | Array  BinarySearch | 在一个循环升序数列中寻找某数 |
|0034| [Find First and Last Position of Element in Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0034_Find_First_and_Last_Position_of_Element_in_Sorted_Array) | Array  BinarySearch | 在有序数列中找到某元素的起始和中止index |
|0035| [Search Insert Position](https://github.com/ZhengjunHUO/leetcode/tree/main/0035_Search_Insert_Position) | Array  BinarySearch | 为某数在一个升序数列中寻找插入位置 |
|0037| [Sudoku Solver](https://github.com/ZhengjunHUO/leetcode/tree/main/0035_Search_Insert_Position) | Array  Backtracking  Matrix | 数独 |
|0039| [Combination Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0039_Combination_Sum) | Array  Backtracking | 找出所有和为目标值的元素（元素可重复使用）组成的列表 |
|0041| [First Missing Positive](https://github.com/ZhengjunHUO/leetcode/tree/main/0041_First_Missing_Positive) | Array  HashTable | 找到数列缺失的最小正整数 |
|0040| [Combination Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0040_Combination_Sum_II) | Array  Backtracking | 找出所有和为目标值的元素（元素不可重复使用）组成的列表 |
|0042| [Trapping Rain Water](https://github.com/ZhengjunHUO/leetcode/tree/main/0042_Trapping_Rain_Water) | Array  Stack  TwoPointers  DynamicProgramming | 计算二维地形的积水量 |
|0043| [Multiply Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0043_Multiply_Strings) | Math  String | 计算以字符串描述的两数的乘积 |
|0045| [Jump Game II](https://github.com/ZhengjunHUO/leetcode/tree/main/0045_Jump_Game_II) | Array  Greedy | 最小跳数到达终点 |
|0046| [Permutations](https://github.com/ZhengjunHUO/leetcode/tree/main/0046_Permutations) | Array  Backtracking | 求数列中所有元素排列组合的集合 |
|0047| [Permutations II](https://github.com/ZhengjunHUO/leetcode/tree/main/0047_Permutations_II) | Array  Backtracking | 求含重复元素数列中所有元素排列组合的集合 |
|0048| [Rotate Image](https://github.com/ZhengjunHUO/leetcode/tree/main/0048_Rotate_Image) | Array  Math  Matrix | 顺时针旋转矩阵 |
|0049| [Group Anagrams](https://github.com/ZhengjunHUO/leetcode/tree/main/0049_Group_Anagrams) | HashTable  String | 将相关联的易位词组合起来 |
|0050| [Pow(x, n)](https://github.com/ZhengjunHUO/leetcode/tree/main/0050_Pow) | Math  Recursion | 实现次方运算 |
|0051| [N-Queens](https://github.com/ZhengjunHUO/leetcode/tree/main/0051_N-Queens) | Array  Backtracking | 计算N*N的棋盘上部署N皇后的所有可能性 |
|0052| [N-Queens II](https://github.com/ZhengjunHUO/leetcode/tree/main/0052_N-Queens_II) | Backtracking | 计算N*N的棋盘上部署N皇后的所有可能性的数量 |
|0053| [Maximum Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0053_Maximum_Subarray) | Array  DivideAndConquer  DynamicProgramming | 找到和最大的子数列 |
|0054| [Spiral Matrix](https://github.com/ZhengjunHUO/leetcode/tree/main/0054_Spiral_Matrix) | Array  Matrix  Simulation | 按顺时针螺旋输出矩阵所有元素 |
|0055| [Jump Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0055_Jump_Game) | Array  Greedy | 求是否可以跳到终点 |
|0056| [Merge Intervals](https://github.com/ZhengjunHUO/leetcode/tree/main/0056_Merge_Intervals) | Array  Sorting | 合并重叠的区间 |
|0057| [Insert Interval](https://github.com/ZhengjunHUO/leetcode/tree/main/0057_Insert_Interval) | Array | 插入区间（如有重叠需融合） |
|0058| [Length of Last Word](https://github.com/ZhengjunHUO/leetcode/tree/main/0058_Length_of_Last_Word) | String | 最后一个词的长度 |
|0059| [Spiral Matrix II](https://github.com/ZhengjunHUO/leetcode/tree/main/0059_Spiral_Matrix_II) | Array  Matrix  Simulation | 按顺时针螺旋向矩阵中插入元素 |
|0060| [Permutation Sequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0060_Permutation_Sequence) | Math  Recursion | 在数列的所有排列组合中按字典顺序找出第k个组合 |
|0061| [Rotate List](https://github.com/ZhengjunHUO/leetcode/tree/main/0061_Rotate_List) | LinkedList  TwoPointers | 将链表的末尾移到链首 |
|0062| [Unique Paths](https://github.com/ZhengjunHUO/leetcode/tree/main/0062_Unique_Paths) | Combinatorics  DynamicProgramming  Math | 计算从矩阵左上角到右下角的路线数 |
|0063| [Unique Paths II](https://github.com/ZhengjunHUO/leetcode/tree/main/0063_Unique_Paths_II) | Array  DynamicProgramming  Matrix | 计算从含障碍物的矩阵左上角到右下角的路线数 |
|0064| [Minimum Path Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0064_Minimum_Path_Sum) | Array  DynamicProgramming  Matrix | 计算从矩阵左上角到右下角的格子值之和最小的路线 |
|0065| [Valid Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0065_Valid_Number) | Math  String | 检查字符串是否为合法的数 |
|0066| [Plus One](https://github.com/ZhengjunHUO/leetcode/tree/main/0066_Plus_One) | Array | 以数组描述的数加一 |
|0067| [Add Binary](https://github.com/ZhengjunHUO/leetcode/tree/main/0067_Add_Binary) | Math  String | 字符串形式的二进制数相加 |
|0069| [Sqrt(x)](https://github.com/ZhengjunHUO/leetcode/tree/main/0069_Sqrt_x) | Math  BinarySearch | 求平方根 |
|0070| [Climbing Stairs](https://github.com/ZhengjunHUO/leetcode/tree/main/0070_Climbing_Stairs) | DynamicProgramming  Math  Memoization | 求共有几种方法到达n |
|0071| [Simplify Path](https://github.com/ZhengjunHUO/leetcode/tree/main/0071_Simplify_Path) | Stack  String | 简化文件路径 |
|0072| [Edit Distance](https://github.com/ZhengjunHUO/leetcode/tree/main/0072_Edit_Distance) | String  DynamicProgramming | 求两个字符串的距离（需要增删改操作的次数）|
|0073| [Set Matrix Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0073_Set_Matrix_Zeroes) | Array  HashTable  Matrix | 将矩阵中值为0的格子的行和列都置0 |
|0074| [Search a 2D Matrix](https://github.com/ZhengjunHUO/leetcode/tree/main/0074_Search_a_2D_Matrix) | Array  BinarySearch  Matrix | 在排好序的矩阵中找元素 |
|0075| [Sort Colors](https://github.com/ZhengjunHUO/leetcode/tree/main/0075_Sort_Colors) | Array  Sorting  TwoPointers | 三种颜色分类排序 |
|0076| [Minimum Window Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/0076_Minimum_Window_Substring) | HashTable  SlidingWindow  String  TwoPointers | 在A字符串中找到包含所有B字符串字母的最小子串 |
|0077| [Combinations](https://github.com/ZhengjunHUO/leetcode/tree/main/0077_Combinations) | Array  Backtracking | 在1到n的范围内找出所有长度为k的组合 |
|0078| [Subsets](https://github.com/ZhengjunHUO/leetcode/tree/main/0078_Subsets) | Array  Backtracking  BitManipulation | 找出数列的所有可能子集 |
|0079| [Word Search](https://github.com/ZhengjunHUO/leetcode/tree/main/0079_Word_Search) | Array  Backtracking  Matrix | 矩阵中寻找一个单词 |
|0080| [Remove Duplicates from Sorted Array II](https://github.com/ZhengjunHUO/leetcode/tree/main/0080_Remove_Duplicates_from_Sorted_Array_II) | Array  TwoPointers | 移除数组中多余的元素(in place) |
|0083| [Remove Duplicates from Sorted List](https://github.com/ZhengjunHUO/leetcode/tree/main/0083_Remove_Duplicates_from_Sorted_List) | LinkedList | 有序链表中删除重复结点 |
|0084| [Largest Rectangle in Histogram](https://github.com/ZhengjunHUO/leetcode/tree/main/0084_Largest_Rectangle_in_Histogram) | Array  MonotonicStack  Stack | 直方图找到最到矩形 |
|0088| [Merge Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0088_Merge_Sorted_Array) | Array  Sorting  TwoPointers | 将两个排好序的字符串in-place融合 |
|0090| [Subsets II](https://github.com/ZhengjunHUO/leetcode/tree/main/0090_Subsets_II) | Array  Backtracking  BitManipulation | 找出含重复元素数列的所有可能的不重复子集 |
|0092| [Reverse Linked List II](https://github.com/ZhengjunHUO/leetcode/tree/main/0092_Reverse_Linked_List_II) | LinkedList | 反转部分链表 |
|0093| [Restore IP Addresses](https://github.com/ZhengjunHUO/leetcode/tree/main/0093_Restore_IP_Addresses) | Backtracking  String | 从字符串还原IP地址 |
|0094| [Binary Tree Inorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0094_Binary_Tree_Inorder_Traversal) | BinaryTree  DepthFirstSearch  Stack  Tree | 中序遍历输出 |
|0095| [Unique Binary Search Trees II](https://github.com/ZhengjunHUO/leetcode/tree/main/0095_Unique_Binary_Search_Trees_II) | Backtracking  BinarySearchTree  BinaryTree  DynamicProgramming  Tree | 返回从n个结点能构造的所有BST |
|0096| [Unique Binary Search Trees](https://github.com/ZhengjunHUO/leetcode/tree/main/0096_Unique_Binary_Search_Trees) | BinarySearchTree  BinaryTree  DynamicProgramming  Math  Tree | 计算n个结点可以构造多少个BST |
|0098| [Validate Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0098_Validate_Binary_Search_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 判断树是否为BST |
|0099| [Recover Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0099_Recover_Binary_Search_Tree) | BinaryTree  BinarySearchTree  DepthFirstSearch  Tree | swap一对错位的结点使树成为BST |
|0100| [Same Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0100_Same_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 判断两棵树是否相同 |
|0101| [Symmetric Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0101_Symmetric_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 判断树是否对称 |
|0105| [Construct Binary Tree from Preorder and Inorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal) | Array  BinaryTree  DivideAndConquer  HashTable  Tree | 从前序遍历和中序遍历列表构建树 |
|0106| [Construct Binary Tree from Inorder and Postorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal) | Array  BinaryTree  DivideAndConquer  HashTable  Tree | 从中序遍历和后序遍历列表构建树 |
|0108| [Convert Sorted Array to Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0108_Convert_Sorted_Array_to_Binary_Search_Tree) | Array  BinaryTree  BinarySearchTree  DivideAndConquer  Tree | 从升序数列构建BST |
|0109| [Convert Sorted List to Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0109_Convert_Sorted_List_to_Binary_Search_Tree) | BinaryTree  BinarySearchTree  DivideAndConquer  LinkedList  Tree | 从升序链表构建BST |
|0111| [Minimum Depth of Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0111_Minimum_Depth_of_Binary_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 求二叉树的最短深度 |
|0112| [Path Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0112_Path_Sum) | BinaryTree  DepthFirstSearch  Tree | 是否存在从根到叶节点值的和等于某值的路径 |
|0114| [Flatten Binary Tree to Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0114_Flatten_Binary_Tree_to_Linked_List) | BinaryTree  DepthFirstSearch  LinkedList  Stack  Tree | 把树压成链表 |
|0116| [Populating Next Right Pointers in Each Node](https://github.com/ZhengjunHUO/leetcode/tree/main/0116_Populating_Next_Right_Pointers_in_Each_Node) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 树的同一层结点横向连接起来 |
|0121| [Best Time to Buy and Sell Stock](https://github.com/ZhengjunHUO/leetcode/tree/main/0121_Best_Time_to_Buy_and_Sell_Stock) | Array  DynamicProgramming | 一次买卖交易获得的最大值 |
|0122| [Best Time to Buy and Sell Stock II](https://github.com/ZhengjunHUO/leetcode/tree/main/0122_Best_Time_to_Buy_and_Sell_Stock_II) | Array  Greedy | 不限次数交易获得的最大值 |
|0123| [Best Time to Buy and Sell Stock III](https://github.com/ZhengjunHUO/leetcode/tree/main/0123_Best_Time_to_Buy_and_Sell_Stock_III) | Array  DynamicProgramming | 至多两次交易获得的最大值 |
|0130| [Surrounded Regions](https://github.com/ZhengjunHUO/leetcode/tree/main/0130_Surrounded_Regions) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix  UnionFind | 把被X包围的O换成X |
|0131| [Palindrome Partitioning](https://github.com/ZhengjunHUO/leetcode/tree/main/0131_Palindrome_Partitioning) | Backtracking  DynamicProgramming  String | 求字符串拆分成回文字符串数组的所有可能性 |
|0136| [Single Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0136_Single_Number) | Array  BitManipulation | 找出数列中只出现一次的元素 |
|0141| [Linked List Cycle](https://github.com/ZhengjunHUO/leetcode/tree/main/0141_Linked_List_Cycle) | HashTable  LinkedList  TwoPointers | 判断链表是否有环 |
|0142| [Linked List Cycle II](https://github.com/ZhengjunHUO/leetcode/tree/main/0142_Linked_List_Cycle_II) | HashTable  LinkedList  TwoPointers | 判断链表是否有环并找出环开始的地方 |
|0144| [Binary Tree Preorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0144_Binary_Tree_Preorder_Traversal) | BinaryTree  DepthFirstSearch  Stack  Tree | 前序遍历输出 |
|0145| [Binary Tree Postorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0145_Binary_Tree_Postorder_Traversal) | BinaryTree  DepthFirstSearch  Stack  Tree | 后序遍历输出 |
|0146| [LRU Cache](https://github.com/ZhengjunHUO/leetcode/tree/main/0146_LRU_Cache) | Design  DoublyLinkedList  HashTable  LinkedList | 实现一个LRU缓存 |
|0152| [Maximum Product Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0152_Maximum_Product_Subarray) | Array  DynamicProgramming | 找出所有元素乘积最大的子串 |
|0155| [Min Stack](https://github.com/ZhengjunHUO/leetcode/tree/main/0155_Min_Stack) | Design  Stack | 构建能快速返回最小值的栈 |
|0167| [Two Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0167_Two_Sum_II) | Array  BinarySearch  TwoPointers | 递增数列中找出和为目标值的两个元素 |
|0168| [Excel Sheet Column Title](https://github.com/ZhengjunHUO/leetcode/tree/main/0168_Excel_Sheet_Column_Title) | Math  String | 把数字转化为excel列名 |
|0169| [Majority Element](https://github.com/ZhengjunHUO/leetcode/tree/main/0169_Majority_Element) | Array  Counting  DivideAndConquer  HashTable  Sorting | 找到多数元素 |
|0170| [Two Sum III Data Structure Design](https://github.com/ZhengjunHUO/leetcode/tree/main/0170_Two_Sum_III_Data_Structure_Design) | Design  HastTable | 设计可以插入数据并判断存储中是否有和为给定值的数据对的数据结构 |
|0171| [Excel Sheet Column Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0171_Excel_Sheet_Column_Number) | Math  String | 把excel列名转化为数字 |
|0172| [Factorial Trailing Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0172_Factorial_Trailing_Zeroes) | Math | 求n!的末尾0的个数 |
|0173| [Binary Search Tree Iterator](https://github.com/ZhengjunHUO/leetcode/tree/main/0173_Binary_Search_Tree_Iterator) | BinarySearchTree  BinaryTree  Design  Iterator  Stack  Tree | 设计一个返回BST中序遍历元素的生成器 |
|0174| [Dungeon Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0174_Dungeon_Game) | Array  DynamicProgramming  Matrix | 地牢游戏求能到达终点所需的最少初始生命值 |
|0187| [Repeated DNA Sequences](https://github.com/ZhengjunHUO/leetcode/tree/main/0187_Repeated_DNA_Sequences) | BitManipulation  HashFunction  HashTable  RollingHash  SlidingWindow  String | 找到至少出现两次的基因段 |
|0188| [Best Time to Buy and Sell Stock IV](https://github.com/ZhengjunHUO/leetcode/tree/main/0188_Best_Time_to_Buy_and_Sell_Stock_IV) | DynamicProgramming | 至多K次交易获得的最大值 |
|0190| [Reverse Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0190_Reverse_Bits) | BitManipulation | 字符串表示的二进制数取反 |
|0191| [Number of 1 Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0191_Number_of_1_Bits) | BitManipulation | 字符串表示的二进制数中1的个数 |
|0192| [Word Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/0192_Word_Frequency) | Shell | 统计词出现个数 |
|0193| [Valid Phone Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0193_Valid_Phone_Numbers) | Shell | 从文本中找出合法的电话号码 |
|0195| [Tenth Line](https://github.com/ZhengjunHUO/leetcode/tree/main/0195_Tenth_Line) | Shell | 打印出文本的第10行 |
|0198| [House Robber](https://github.com/ZhengjunHUO/leetcode/tree/main/0198_House_Robber) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0200| [Number of Islands](https://github.com/ZhengjunHUO/leetcode/tree/main/0200_Number_of_Islands) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix  UnionFind | 找到有几个岛屿 |
|0203| [Remove Linked List Elements](https://github.com/ZhengjunHUO/leetcode/tree/main/0203_Remove_Linked_List_Elements) | LinkedList  Recursion | 链表中删除特定值的元素 |
|0204| [Count Primes](https://github.com/ZhengjunHUO/leetcode/tree/main/0204_Count_Primes) | Array  Enumeration  Math  NumberTheory | 统计小于n的质数个数 |
|0206| [Reverse Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0206_Reverse_Linked_List) | LinkedList  Recursion | 反转链表 |
|0207| [Course Schedule](https://github.com/ZhengjunHUO/leetcode/tree/main/0207_Course_Schedule) | BreadthFirstSearch  DepthFirstSearch  Graph  TopologicalSort | 是否能顺利完成所有课程 |
|0209| [Minimum Size Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0209_Minimum_Size_Subarray_Sum) | Array  BinarySearch  TwoPointers | 求和大于等于目标值的最短子串 |
|0210| [Course Schedule II](https://github.com/ZhengjunHUO/leetcode/tree/main/0210_Course_Schedule_II) | BreadthFirstSearch  DepthFirstSearch  Graph  TopologicalSort | 如能完成所有课程给出完成顺序 |
|0213| [House Robber II](https://github.com/ZhengjunHUO/leetcode/tree/main/0213_House_Robber_II) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0214| [Shortest Palindrome](https://github.com/ZhengjunHUO/leetcode/tree/main/0214_Shortest_Palindrome) | HashFunction  RollingHash  String  StringMatching  | 找到从字符串头开始最长的回文组 |
|0215| [Kth Largest Element in an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0215_Kth_Largest_Element_in_an_Array) | Array  DivideAndConquer  Heap  Quickselect  Sorting | 数列中寻找第k大元素 |
|0222| [Count Complete Tree Nodes](https://github.com/ZhengjunHUO/leetcode/tree/main/0222_Count_Complete_Tree_Nodes) | BinarySearch  BinaryTree  DepthFirstSearch  Tree | 数一颗Complete二叉树的结点个数 |
|0224| [Basic Calculator](https://github.com/ZhengjunHUO/leetcode/tree/main/0224_Basic_Calculator) | Math  Recursion  Stack  String | 实现一个基础的加减计算器 |
|0225| [Implement Stack using Queues](https://github.com/ZhengjunHUO/leetcode/tree/main/0225_Implement_Stack_using_Queues) | Design  Queue  Stack | 用queue来实现stack |
|0226| [Invert Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0226_Invert_Binary_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 二叉树元素镜像转置 |
|0227| [Basic Calculator II](https://github.com/ZhengjunHUO/leetcode/tree/main/0227_Basic_Calculator_II) | Math  Stack  String | 实现一个基础的加减乘除计算器 |
|0228| [Summary Ranges](https://github.com/ZhengjunHUO/leetcode/tree/main/0228_Summary_Ranges) | Array | 把递增数列中连续的数字用范围表示 |
|0230| [Kth Smallest Element in a BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0230_Kth_Smallest_Element_in_a_BST) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 在BST中找到第k小的元素 |
|0231| [Power of Two](https://github.com/ZhengjunHUO/leetcode/tree/main/0231_Power_of_Two) | BitManipulation  Math  Recursion | 判断数是否为2的次方 |
|0232| [Implement Queue using Stacks](https://github.com/ZhengjunHUO/leetcode/tree/main/0232_Implement_Queue_using_Stacks) | Design  Queue  Stack | 用stack来实现queue |
|0234| [Palindrome Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0234_Palindrome_Linked_List) | LinkedList  Recursion  Stack  TwoPointers | 判断链表是否是回文 |
|0235| [Lowest Common Ancestor of a Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0235_Lowest_Common_Ancestor_of_a_Binary_Search_Tree) | BinarySearchTree  BinaryTree  DepthFirstSearch  Tree | 找到BST中两个元素的LCA |
|0236| [Lowest Common Ancestor of a Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0236_Lowest_Common_Ancestor_of_a_Binary_Tree) | BinaryTree  DepthFirstSearch  Tree | 找到二叉树中两个元素的LCA |
|0238| [Product of Array Except Self](https://github.com/ZhengjunHUO/leetcode/tree/main/0238_Product_of_Array_Except_Self) | Array | 返回数列所有元素乘积和除以各元素的数列 |
|0239| [Sliding Window Maximum](https://github.com/ZhengjunHUO/leetcode/tree/main/0239_Sliding_Window_Maximum) | Dequeue  Heap  SlidingWindow | 返回定长的向右滑动的窗口每一步的局部最大值组成的数组 |
|0241| [Different Ways to Add Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0241_Different_Ways_to_Add_Parentheses) | DynamicProgramming  Math  Memoization  Recursion  String | 给算式任意位置添加括号得到的所有可能解 |
|0242| [Valid Anagram](https://github.com/ZhengjunHUO/leetcode/tree/main/0242_Valid_Anagram) | HastTable  Sort | 判断一个字符串是否为另一个的易位词 |
|0252| [Meeting Rooms](https://github.com/ZhengjunHUO/leetcode/tree/main/0252_Meeting_Rooms) | Array  Sorting | 会议时间是否重合 |
|0257| [Binary Tree Paths](https://github.com/ZhengjunHUO/leetcode/tree/main/0257_Binary_Tree_Paths) | BinaryTree  DepthFirstSearch  String  Tree | 打印出从根到叶子的所有路径 |
|0259| [3Sum Smaller](https://github.com/ZhengjunHUO/leetcode/tree/main/0259_3Sum_Smaller) | Array  TwoPointers | 求数组中三元组和小于目标值的个数 |
|0263| [Ugly Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0263_Ugly_Number) | Math | 判断一个数的质因数是否限定在2,3,5之中 |
|0264| [Ugly Number II](https://github.com/ZhengjunHUO/leetcode/tree/main/0264_Ugly_Number_II) | DynamicProgramming  HashTable  Heap  Math | 求从1开始第n个题0263中描述的数字 |
|0268| [Missing Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0268_Missing_Number) | Array  BitManipulation  HashTable  Math  Sorting | 找出缺失元素 |
|0278| [First Bad Version](https://github.com/ZhengjunHUO/leetcode/tree/main/0278_First_Bad_Version) | BinarySearch  Interactive | 找到第一个坏版本 |
|0279| [Perfect Squares](https://github.com/ZhengjunHUO/leetcode/tree/main/0279_Perfect_Squares) | BreadthFirstSearch  DynamicProgramming  Math | 把数分解为最短的平方和 |
|0283| [Move Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0283_Move_Zeroes) | Array  TwoPointers | 将数列中的0移到最后 |
|0292| [Nim Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0292_Nim_Game) | Brainteaser  GameTheory  Math | 拿走最后一个石子获胜的游戏 |
|0295| [Find Median from Data Stream](https://github.com/ZhengjunHUO/leetcode/tree/main/0295_Find_Median_from_Data_Stream) | DataStream  Design  Heap  Sorting  TwoPointers | 数据流中求中位数 |
|0297| [Serialize and Deserialize Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0297_Serialize_and_Deserialize_Binary_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Design  String  Tree | 序列化及反序列化一棵树 | 
|0300| [Longest Increasing Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0300_Longest_Increasing_Subsequence) | Array  BinarySearch  DynamicProgramming | 求最长的递增子数列长度 |
|0305| [Number of Islands II](https://github.com/ZhengjunHUO/leetcode/tree/main/0305_Number_of_Islands_II) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix  UnionFind | 填海造岛过程中岛屿个数 |
|0309| [Best Time to Buy and Sell Stock with Cooldown](https://github.com/ZhengjunHUO/leetcode/tree/main/0309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown) | DynamicProgramming | 不限次数交易获得的最大值，卖出后有一天冷却期 |
|0310| [Minimum Height Trees](https://github.com/ZhengjunHUO/leetcode/tree/main/0310_Minimum_Height_Trees) | BreadthFirstSearch  DepthFirstSearch  Graph  TopologicalSort | 找到最矮的树 |
|0313| [Super Ugly Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0313_Super_Ugly_Number) | Array  DynamicProgramming  HashTable  Heap  Math | |
|0316| [Remove Duplicate Letters](https://github.com/ZhengjunHUO/leetcode/tree/main/0316_Remove_Duplicate_Letters) | Greedy  MonotonicStack  Stack  String | 保持相对顺序的前提下移除重复字母，并且字典序最小 |
|0319| [Bulb Switcher](https://github.com/ZhengjunHUO/leetcode/tree/main/0319_Bulb_Switcher) | Math  Brainteaser | 开关灯游戏 |
|0322| [Coin Change](https://github.com/ZhengjunHUO/leetcode/tree/main/0322_Coin_Change) | Array  BreadthFirstSearch  DynamicProgramming | 最少能用几枚硬币凑成目标数额 |
|0325| [Maximum Size Subarray Sum Equals k](https://github.com/ZhengjunHUO/leetcode/tree/main/0325_Maximum_Size_Subarray_Sum_Equals_k) | ??? | 求和等于目标值的最长子串 |
|0326| [Power of Three](https://github.com/ZhengjunHUO/leetcode/tree/main/0326_Power_of_Three) | Math  Recursion | 判断某数是否是3的次方 |
|0337| [House Robber III](https://github.com/ZhengjunHUO/leetcode/tree/main/0337_House_Robber_III) | DepthFirstSearch  DynamicProgramming  Tree | 抢劫树形排列不相邻房屋获得的最大收益 |
|0341| [Flatten Nested List Iterator](https://github.com/ZhengjunHUO/leetcode/tree/main/0341_Flatten_Nested_List_Iterator) | DepthFirstSearch  Design  Iterator  Queue  Stack  Tree | 实现一个iterator来扁平化输出一个嵌套的列表 |
|0342| [Power of Four](https://github.com/ZhengjunHUO/leetcode/tree/main/0342_Power_of_Four) | BitManipulation  Math  Recursion | 判断某数是否是4的次方 |
|0344| [Reverse String](https://github.com/ZhengjunHUO/leetcode/tree/main/0344_Reverse_String) | Recursion  String  TwoPointers | 将字符数组反转 |
|0347| [Top K Frequent Elements](https://github.com/ZhengjunHUO/leetcode/tree/main/0347_Top_K_Frequent_Elements) | Array  BucketSort  Counting  DivideAndConquer  HashTable  Heap  Quickselect  Sorting | 返回前k个最常见的元素 |
|0354| [Russian Doll Envelopes](https://github.com/ZhengjunHUO/leetcode/tree/main/0354_Russian_Doll_Envelopes) | Array  BinarySearch  DynamicProgramming  Sorting | 最多有多少小信封能塞进大信封 |
|0355| [Design Twitter](https://github.com/ZhengjunHUO/leetcode/tree/main/0355_Design_Twitter) | Design  HashTable  Heap  LinkedList | 实现简单的推特应用 |
|0371| [Sum of Two Integers](https://github.com/ZhengjunHUO/leetcode/tree/main/0371_Sum_of_Two_Integers) | BitManipulation | 位操作实现两数相加 |
|0372| [Super Pow](https://github.com/ZhengjunHUO/leetcode/tree/main/0371_Sum_of_Two_Integers) | DivideAndConquer  Math | 求大数的高次方 |
|0373| [Find K Pairs with Smallest Sums](https://github.com/ZhengjunHUO/leetcode/tree/main/0373_Find_K_Pairs_with_Smallest_Sums) | Array  Heap | 两个数列中找到和最小的k对元素 |
|0378| [Kth Smallest Element in a Sorted Matrix](https://github.com/ZhengjunHUO/leetcode/tree/main/0378_Kth_Smallest_Element_in_a_Sorted_Matrix) | Array  BinarySearch  Heap  Matrix  Sorting | 在行和列都是升序的矩阵中找到第k小的元素 |
|0380| [Insert Delete GetRandom O(1)](https://github.com/ZhengjunHUO/leetcode/tree/main/0380_Insert_Delete_GetRandom_O1) | Array  Design  HashTable  Math  Randomized | 实现增删随机读的复杂度都为O(1)的数据结构 |
|0382| [Linked List Random Node](https://github.com/ZhengjunHUO/leetcode/tree/main/0382_Linked_List_Random_Node) | LinkedList  Math  Randomized  ReservoirSampling | 随机均匀选取链表中的某结点 | 
|0384| [Shuffle an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0384_Shuffle_an_Array) | Array  Math  Randomized | 随机均匀地洗牌 |
|0387| [First Unique Character in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0387_First_Unique_Character_in_a_String) | Counting  HashTable  Queue  String | 找到第一个只出现一次的字符 |
|0389| [Find the Difference](https://github.com/ZhengjunHUO/leetcode/tree/main/0389_Find_the_Difference) | BitManipulation  HastTable  Sorting  String | 找出多余的一个字符 |
|0392| [Is Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0392_Is_Subsequence) | DynamicProgramming  String  TwoPointers | 判断是否是子串 |
|0407| [Trapping Rain Water II](https://github.com/ZhengjunHUO/leetcode/tree/main/0407_Trapping_Rain_Water_II) | BreadthFirstSearch  Heap | 计算三维地形的积水量 |
|0410| [Split Array Largest Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0410_Split_Array_Largest_Sum) | Array  BinarySearch  DynamicProgramming  Greedy | 把数列切分成N个子数列，求令子数列和最大值最小化的方案 |
|0414| [Third Maximum Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0414_Third_Maximum_Number) | Array  Sorting | 第三大的数字 |
|0415| [Add Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0415_Add_Strings) | String | 以字符串形式表示的两数之和 |
|0416| [Partition Equal Subset Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0416_Partition_Equal_Subset_Sum) | Array  DynamicProgramming | 求数组能否分成两个和相等的子集 |
|0435| [Non-overlapping Intervals](https://github.com/ZhengjunHUO/leetcode/tree/main/0435_Non-overlapping_Intervals) | Array  DynamicProgramming  Greedy  Sorting | 移除最少的区间使列表中的区间没有重叠 |
|0438| [Find All Anagrams in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0438_Find_All_Anagrams_in_a_String) | HastTable | 在字符串A中找到所有字符串B的易位词 | 
|0441| [Arranging Coins](https://github.com/ZhengjunHUO/leetcode/tree/main/0441_Arranging_Coins) | BinarySearchTree  Math | 用硬币填满阶梯 |
|0445| [Add Two Numbers II](https://github.com/ZhengjunHUO/leetcode/tree/main/0445_Add_Two_Numbers_II) | LinkedList | 以链表形式表示的两数之和 |
|0448| [Find All Numbers Disappeared in an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0448_Find_All_Numbers_Disappeared_in_an_Array) | Array  HashTable | 找到所有缺失的数 |
|0449| [Serialize and Deserialize BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0449_Serialize_and_Deserialize_BST) | BinarySearchTree  BinaryTree  BreadthFirstSearch  DepthFirstSearch  Design  String  Tree | 序列化反序列化BST |
|0450| [Delete Node in a BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0450_Delete_Node_in_a_BST) | BinarySearchTree  BinaryTree  Tree | 在BST中删除一个结点 |
|0451| [Sort Characters By Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/0451_Sort_Characters_By_Frequency) | BucketSort  Counting  HashTable  Heap  Sorting  String | 按字符出现的频率降序排列字符串 |
|0452| [Minimum Number of Arrows to Burst Balloons](https://github.com/ZhengjunHUO/leetcode/tree/main/0452_Minimum_Number_of_Arrows_to_Burst_Balloons) | Array  Greedy  Sorting | 射出最少的箭穿破所有气球 |
|0454| [4Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0454_4Sum_II) | BinarySearch  HastTable | 四个数列中分别取一个数q之和等于0的四元组个数 |
|0455| [Assign Cookies](https://github.com/ZhengjunHUO/leetcode/tree/main/0455_Assign_Cookies) | Array  Greedy  Sorting | 分配饼干 |
|0459| [Repeated Substring Pattern](https://github.com/ZhengjunHUO/leetcode/tree/main/0459_Repeated_Substring_Pattern) | String | 检查字符串是否由某个子串重复n次构成 |
|0460| [LFU Cache](https://github.com/ZhengjunHUO/leetcode/tree/main/0460_LFU_Cache) | Design  DoublyLinkedList  HashTable  LinkedList | 实现一个LFU缓存 |
|0463| [Island Perimeter](https://github.com/ZhengjunHUO/leetcode/tree/main/0463_Island_Perimeter) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix | 计算岛屿的周长 |
|0474| [Ones and Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0474_Ones_and_Zeroes) | DynamicProgramming | 字符串表示的二进制数组中，找出最大的包含至多m个0和n个1的子集合 |
|0476| [Number Complement](https://github.com/ZhengjunHUO/leetcode/tree/main/0476_Number_Complement) | BitManipulation | 求数的二进制反转 |
|0494| [Target Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0494_Target_Sum) | Array  Backtracking  DynamicProgramming | 数列中的数字可以取正负，求其和可以等于目标值的组合数 |
|0495| [Teemo Attacking](https://github.com/ZhengjunHUO/leetcode/tree/main/0495_Teemo_Attacking) | Array  Simulation | 计算进入异常状态的时间 |
|0496| [Next Greater Element I](https://github.com/ZhengjunHUO/leetcode/tree/main/0496_Next_Greater_Element_I) | Array  HashTable  MonotonicStack  Stack | 找到下一个大于自己的元素 |
|0502| [IPO](https://github.com/ZhengjunHUO/leetcode/tree/main/0502_IPO) | Array  Greedy  Heap  Sorting | 完成k个项目可以获得的最大纯利润 |
|0503| [Next Greater Element II](https://github.com/ZhengjunHUO/leetcode/tree/main/0503_Next_Greater_Element_II) | Array  MonotonicStack  Stack | 首尾循环的列表寻找下一个大于自己的元素 |
|0504| [Base 7](https://github.com/ZhengjunHUO/leetcode/tree/main/0504_Base_7) | Math | 十进制转七进制 |
|0506| [Relative Ranks](https://github.com/ZhengjunHUO/leetcode/tree/main/0506_Relative_Ranks) | Array  Heap  Sorting | 把分数转换为名次 |
|0509| [Fibonacci Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0509_Fibonacci_Number) | DynamicProgramming  Math  Memoization | 斐波那契数列 |
|0516| [Longest Palindromic Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0516_Longest_Palindromic_Subsequence) | DynamicProgramming  String | 求最长回文子序列 |
|0518| [Coin Change 2](https://github.com/ZhengjunHUO/leetcode/tree/main/0518_Coin_Change_2) | Array  DynamicProgramming | 求用不同面值硬币能凑出给定总额的方法数 |
|0523| [Continuous Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0523_Continuous_Subarray_Sum) | DynamicProgramming  Math | 求是否有一个长度至少为2的子串其和为给定值的倍数 |
|0528| [Random Pick with Weight](https://github.com/ZhengjunHUO/leetcode/tree/main/0528_Random_Pick_with_Weight) | BinarySearch  Math  PrefixSum  Randomized | 按照权重随机选取元素 |
|0530| [Minimum Absolute Difference in BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0530_Minimum_Absolute_Difference_in_BST) | BinarySearchTree  BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | BST中搜索两个结点差的最小值 |
|0538| [Convert BST to Greater Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0538_Convert_BST_to_Greater_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 结点的值更新为整棵树所有大于等于该结点的值的和 |
|0543| [Diameter of Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0543_Diameter_of_Binary_Tree) | BinaryTree  DepthFirstSearch  Tree | 计算相距最远的两个结点的距离 |
|0547| [Number of Provinces](https://github.com/ZhengjunHUO/leetcode/tree/main/0547_Number_of_Provinces) | BreadthFirstSearch  DepthFirstSearch  Graph  UnionFind | 统计有几个城市群 |
|0560| [Subarray Sum Equals K](https://github.com/ZhengjunHUO/leetcode/tree/main/0560_Subarray_Sum_Equals_K) | Array  HashTable | 求和为目标值的子串个数 |
|0567| [Permutation in String](https://github.com/ZhengjunHUO/leetcode/tree/main/0567_Permutation_in_String) | TwoPointer  SlidingWindow | 求字符串A是否包含字符串B的变种 |
|0572| [Subtree of Another Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0572_Subtree_of_Another_Tree) | BinaryTree  DepthFirstSearch  HashFunction  StringMatching  Tree | 判断一棵树是否为另一棵的子树 |
|0583| [Delete Operation for Two Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0583_Delete_Operation_for_Two_Strings) | String  DynamicProgramming | 求令两个字符串相等需要的删除操作数 |
|0600| [Non-negative Integers without Consecutive Ones](https://github.com/ZhengjunHUO/leetcode/tree/main/0600_Non-negative_Integers_without_Consecutive_Ones) | DynamicProgramming | 从0到给定值之间所有数的二进制表示中没有连续1的数的个数 |
|0628| [Maximum Product of Three Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0628_Maximum_Product_of_Three_Numbers) | Array  Math | 数列中挑选三个数使其乘积最大 |
|0632| [Smallest Range Covering Elements from K Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0632_Smallest_Range_Covering_Elements_from_K_Lists) | HashTable  String  TwoPointers | 找到一个最小范围覆盖所有K个数列中至少一个元素 |
|0645| [Set Mismatch](https://github.com/ZhengjunHUO/leetcode/tree/main/0645_Set_Mismatch) | Array  BitManipulation  HashTable  Sorting | 集合中找出出现两次的元素和缺少的元素 |
|0652| [Find Duplicate Subtrees](https://github.com/ZhengjunHUO/leetcode/tree/main/0652_Find_Duplicate_Subtrees) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 寻找重复的子树 |
|0653| [Two Sum IV Input Is A BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0653_Two_Sum_IV_Input_Is_A_BST) | Tree | 在BST中找到等于目标值的两个数 |
|0654| [Maximum Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0654_Maximum_Binary_Tree) | Array  BinaryTree  DivideAndConquer  MonotonicStack  Stack  Tree | 按要求构建二叉树 |
|0657| [Robot Return to Origin](https://github.com/ZhengjunHUO/leetcode/tree/main/0657_Robot_Return_to_Origin) | Simulation  String | 判断机器人经过一系列指令后能否回到原点 |
|0659| [Split Array into Consecutive Subsequences](https://github.com/ZhengjunHUO/leetcode/tree/main/0659_Split_Array_into_Consecutive_Subsequences) | Array  Greedy  HashTable  Heap | 能否将数组分成数个递增的子数组 |
|0684| [Redundant Connection](https://github.com/ZhengjunHUO/leetcode/tree/main/0684_Redundant_Connection) | BreadthFirstSearch  DepthFirstSearch  Graph  UnionFind | 移除多余的边使带环的图成为树 |
|0686| [Repeated String Match](https://github.com/ZhengjunHUO/leetcode/tree/main/0686_Repeated_String_Match) | String | a字符串重复多少次可以包含b字符串 |
|0689| [Maximum Sum of 3 Non-Overlapping Subarrays](https://github.com/ZhengjunHUO/leetcode/tree/main/0689_Maximum_Sum_of_3_Non-Overlapping_Subarrays) | Array  DynamicProgramming | 找到三个指定长度的和最大的独立子串 |
|0692| [Top K Frequent Words](https://github.com/ZhengjunHUO/leetcode/tree/main/0692_Top_K_Frequent_Words) | BucketSort  Counting  HashTable  Heap  Sorting  String  Trie | 出现频率最高的k个词 |
|0693| [Binary Number with Alternating Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0693_Binary_Number_with_Alternating_Bits) | BitManipulation | 判断数的二进制表示是否为0和1相间 |
|0695| [Max Area of Island](https://github.com/ZhengjunHUO/leetcode/tree/main/0695_Max_Area_of_Island) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix  UnionFind | 找到最大的岛 |
|0697| [Degree of an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0697_Degree_of_an_Array) | Array | 找到和字符串有相同度的子串 |
|0698| [Partition to K Equal Sum Subsets](https://github.com/ZhengjunHUO/leetcode/tree/main/0698_Partition_to_K_Equal_Sum_Subsets) | Array  Backtracking  BitManipulation  Bitmask  DynamicProgramming  Memoization | 能否把集合分成k个和相等的子集 |
|0700| [Search in a Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0700_Search_in_a_Binary_Search_Tree) | BinarySearchTree  BinaryTree  Tree | 在BST中搜索一个元素 |
|0701| [Insert into a Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0701_Insert_into_a_Binary_Search_Tree) | BinarySearchTree  BinaryTree  Tree | 在BST中插入一个元素 |
|0703| [Kth Largest Element in a Stream](https://github.com/ZhengjunHUO/leetcode/tree/main/0703_Kth_Largest_Element_in_a_Stream) | BinarySearchTree  BinaryTree  DataStream  Design  Heap  Tree | 不断插入新元素同时找到第k大的元素 |
|0704| [Binary Search](https://github.com/ZhengjunHUO/leetcode/tree/main/0704_Binary_Search) | Array  BinarySearch | 二分搜索元素 |
|0710| [Random Pick with Blacklist](https://github.com/ZhengjunHUO/leetcode/tree/main/0710_Random_Pick_with_Blacklist) | BinarySearch  HashTable  Math  Randomized  Sorting | 随机返回数列中除黑名单以外的元素 |
|0712| [Minimum ASCII Delete Sum for Two Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0712_Minimum_ASCII_Delete_Sum_for_Two_Strings) | String  DynamicProgramming | 删除字母使字符串相等，求删除的字母ASCII值最小 |
|0713| [Subarray Product Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/0713_Subarray_Product_Less_Than_K) | Array  TwoPointers | 乘积小于目标值的子串的个数 |
|0714| [Best Time to Buy and Sell Stock with Transaction Fee](https://github.com/ZhengjunHUO/leetcode/tree/main/0714_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee) | Array  DynamicProgramming  Greedy | 不限次数交易获得的最大值，每笔交易有手续费 |
|0718| [Maximum Length of Repeated Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0718_Maximum_Length_of_Repeated_Subarray) | Array  BinarySearch  DynamicProgramming  HashTable | 寻找最大公共子串 | 
|0724| [Find Pivot Index](https://github.com/ZhengjunHUO/leetcode/tree/main/0724_Find_Pivot_Index) | Array | 找到左边所有元素和等于右边所有元素和的index |
|0727| [Minimum Window Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0727_Minimum_Window_Subsequence) | DynamicProgramming  HashTable  SlidingWindow  String  TwoPointers | 字符串A中找到最短的可以包含字符串B的子串
|0739| [Daily Temperatures](https://github.com/ZhengjunHUO/leetcode/tree/main/0739_Daily_Temperatures) | Array  MonotonicStack  Stack | 求等待几天可以比今天更热 |
|0740| [Delete and Earn](https://github.com/ZhengjunHUO/leetcode/tree/main/0740_Delete_and_Earn) | DynamicProgramming | 在数组中取一个数可得分，但是要删除它和如果存在的它+-1的值，求最大得分 |
|0743| [Network Delay Time](https://github.com/ZhengjunHUO/leetcode/tree/main/0743_Network_Delay_Time) | BreadthFirstSearch  DepthFirstSearch  Graph  Heap  ShortestPath | 从一个点出发遍历所有网络需要花费的时间 |
|0752| [Open the Lock](https://github.com/ZhengjunHUO/leetcode/tree/main/0752_Open_the_Lock) | Array  BreadthFirstSearch  HashTable  String | 密码锁至少需要转动几下才能打开 |
|0763| [Partition Labels](https://github.com/ZhengjunHUO/leetcode/tree/main/0763_Partition_Labels) | Greedy  HashTable  String  TwoPointers | 尽可能将字符串分段，且每个字母仅出现在某段中 |
|0772| [Basic Calculator III](https://github.com/ZhengjunHUO/leetcode/tree/main/0772_Basic_Calculator_III) | Math  Stack  String | 实现一个基础的加减乘除计算器 |
|0773| [Sliding Puzzle](https://github.com/ZhengjunHUO/leetcode/tree/main/0773_Sliding_Puzzle) | Array  BreadthFirstSearch  Matrix | 滑块游戏 |
|0774| [Minimize Max Distance to Gas Station](https://github.com/ZhengjunHUO/leetcode/tree/main/0774_Minimize_Max_Distance_to_Gas_Station) | Array  BinarySearch | 插入K个加油站，找到两个加油站之间最大距离的最小值 |
|0783| [Minimum Distance Between BST Nodes](https://github.com/ZhengjunHUO/leetcode/tree/main/0783_Minimum_Distance_Between_BST_Nodes) | BinarySearchTree  BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | BST中搜索两个结点差的最小值 |
|0784| [Letter Case Permutation](https://github.com/ZhengjunHUO/leetcode/tree/main/0784_Letter_Case_Permutation) | Backtracking  BitManipulation  String | alnum字符串中字母大小写变换的排列组合 |
|0785| [Is Graph Bipartite?](https://github.com/ZhengjunHUO/leetcode/tree/main/0785_Is_Graph_Bipartite) | BreadthFirstSearch  DepthFirstSearch  Graph  UnionFind | 能否把结点分为两组，使每条边的两边的结点分属两个组 |
|0787| [Cheapest Flights Within K Stops](https://github.com/ZhengjunHUO/leetcode/tree/main/0787_Cheapest_Flights_Within_K_Stops) | BreadthFirstSearch  DepthFirstSearch  DynamicProgramming  Graph  Heap  ShortestPath | 求转机k次以内两地间最便宜的花费 |
|0793| [Preimage Size of Factorial Zeroes Function](https://github.com/ZhengjunHUO/leetcode/tree/main/0793_Preimage_Size_of_Factorial_Zeroes_Function) | Math  BinarySearch | 求阶乘值的末尾有k个0的数的个数 |
|0796| [Rotate String](https://github.com/ZhengjunHUO/leetcode/tree/main/0796_Rotate_String) | String  StringMatching | 判断一个字符串是否是另一个字符串shift后的形态 |
|0797| [All Paths From Source to Target](https://github.com/ZhengjunHUO/leetcode/tree/main/0797_All_Paths_From_Source_to_Target) | Backtracking  BreadthFirstSearch  DepthFirstSearch  Graph | 找到DAG中起点到终点的所有路径 |
|0827| [Making A Large Island](https://github.com/ZhengjunHUO/leetcode/tree/main/0827_Making_A_Large_Island) | Array  BreadthFirstSearch  DepthFirstSearch  Matrix  UnionFind | 海中最多填一个土能造出的最大岛 |
|0870| [Advantage Shuffle](https://github.com/ZhengjunHUO/leetcode/tree/main/0870_Advantage_Shuffle) | Array  Greedy  Sorting | 交换数列中的元素使其对另一个数列逐元素比较大小占据最大优势 |
|0875| [Koko Eating Bananas](https://github.com/ZhengjunHUO/leetcode/tree/main/0875_Koko_Eating_Bananas) | Array  BinarySearch | 求在规定时间内吃完所有香蕉堆的最小进食速度 |
|0876| [Middle of the Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0876_Middle_of_the_Linked_List) | LinkedList  TwoPointers | 寻找链表的中点 |
|0877| [Stone Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0877_Stone_Game) | Array  DynamicProgramming  GameTheory  Math | 拿石头堆的游戏 |
|0887| [Super Egg Drop](https://github.com/ZhengjunHUO/leetcode/tree/main/0887_Super_Egg_Drop) | BinarySearch  DynamicProgramming  Math | 找到扔下鸡蛋刚好不会碎的楼层数需要的最少尝试次数 |
|0895| [Maximum Frequency Stack](https://github.com/ZhengjunHUO/leetcode/tree/main/0895_Maximum_Frequency_Stack) | Design  HashTable  OrderedSet  Stack | 实现一个Pop最大出现频率的值的栈 |
|0912| [Sort an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0912_Sort_an_Array) | Array  BucketSort  CountingSort  DivideAndConquer  Heap  MergeSort  RadixSort  Sorting | 给数列排序 |
|0921| [Minimum Add to Make Parentheses Valid](https://github.com/ZhengjunHUO/leetcode/tree/main/0921_Minimum_Add_to_Make_Parentheses_Valid) | Greedy  Stack  String | 添加最小的括号数时括号合法成对 |
|0931| [Minimum Falling Path Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0931_Minimum_Falling_Path_Sum) | Array  DynamicProgramming  Matrix | 求坠落路径上数值之和最小的路径 |
|0938| [Range_Sum_of_BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0938_Range_Sum_of_BST) | BinarySearchTree  BinaryTree  DepthFirstSearch  Tree | 范围内BST节点值的和 |
|0969| [Pancake Sorting](https://github.com/ZhengjunHUO/leetcode/tree/main/0969_Pancake_Sorting) | Array  Greedy  Sorting  TwoPointers | 煎饼排序 |
|0974| [Subarray Sums Divisible by K](https://github.com/ZhengjunHUO/leetcode/tree/main/0974_Subarray_Sums_Divisible_by_K) | Array  HashTable | 和可以整除目标值的子串的个数 |
|0977| [Squares of a Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0977_Squares_of_a_Sorted_Array) | Array  Sorting  TwoPointers | 求排好序的数组各元素平方后得到的新升序数组 |
|0978| [Longest Turbulent Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0978_Longest_Turbulent_Subarray) | Array  DynamicProgramming  SlidingWindow | 找到符合峰谷交替规律的最长的子串 |
|0986| [Interval List Intersections](https://github.com/ZhengjunHUO/leetcode/tree/main/0986_Interval_List_Intersections) | Array  TwoPointers | 两个区间组成的列表找交集 |
|0989| [Add to Array-Form of Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0989_Add_to_Array-Form_of_Integer) | Array | 求一个以数组形式表示的数和一个整数之和 |
|0990| [Satisfiability of Equality Equations](https://github.com/ZhengjunHUO/leetcode/tree/main/0990_Satisfiability_of_Equality_Equations) | Array  Graph  String  UnionFind | 判断等式集合是否存在矛盾 |
|1009| [Complement of Base 10 Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/1009_Complement_of_Base_10_Integer) | BitManipulation | 求数的二进制反转 |
|1011| [Capacity To Ship Packages Within D Days](https://github.com/ZhengjunHUO/leetcode/tree/main/1011_Capacity_To_Ship_Packages_Within_D_Days) | Array  BinarySearch  Greedy | 求在规定的天数内能装完所有货物的最小的船的容量 |
|1023| [Camelcase Matching](https://github.com/ZhengjunHUO/leetcode/tree/main/1023_Camelcase_Matching) | String  StringMatching  Trie  TwoPointers | 字符串匹配 |
|1024| [Video Stitching](https://github.com/ZhengjunHUO/leetcode/tree/main/1024_Video_Stitching) | Array  DynamicProgramming  Greedy | 拼接一段完整视屏需要的最少数量的片段 |
|1038| [Binary Search Tree to Greater Sum Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/1038_Binary_Search_Tree_to_Greater_Sum_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 结点的值更新为整棵树所有大于等于该结点的值的和 |
|1081| [Smallest Subsequence of Distinct Characters](https://github.com/ZhengjunHUO/leetcode/tree/main/1081_Smallest_Subsequence_of_Distinct_Characters) | Greedy  MonotonicStack  Stack  String | 保持相对顺序的前提下移除重复字母，并且字典序最小 |
|1099| [Two Sum Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/1099_Two_Sum_Less_Than_K) | Array | 数组中找到小于目标值的最大两数之和 | 
|1109| [Corporate Flight Bookings](https://github.com/ZhengjunHUO/leetcode/tree/main/1109_Corporate_Flight_Bookings) | Array  PrefixSum | 从订单中统计各航班座位预定总数 |
|1143| [Longest Common Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/1143_Longest_Common_Subsequence) | DynamicProgramming  String | 求最长公共子数列 |
|1201| [Ugly Number III](https://github.com/ZhengjunHUO/leetcode/tree/main/1201_Ugly_Number_III) | BinarySearch  Math  NumberTheory | |
|1288| [Remove Covered Intervals](https://github.com/ZhengjunHUO/leetcode/tree/main/1288_Remove_Covered_Intervals) | Array  Sorting | 移除能被另一个区间覆盖的区间 |
|1306| [Jump Game III](https://github.com/ZhengjunHUO/leetcode/tree/main/1306_Jump_Game_III) | BreadthFirstSearch  DepthFirstSearch  Recursion | 从给定的出发点允许向前向后跳跃，判断是否能到终点 |
|1340| [Jump Game V](https://github.com/ZhengjunHUO/leetcode/tree/main/1340_Jump_Game_V) | DynamicProgramming | 从任意点出发允许向前向后最远跳d的距离且只能从大数跳向小数，求最多能覆盖几个index ｜
|1345| [Jump Game IV](https://github.com/ZhengjunHUO/leetcode/tree/main/1345_Jump_Game_IV) | BreadthFirstSearch | 从头开始，每一条可以向前向后一格，或传送到其他值相等的index，求到达终点最少的跳跃次数 |
|1373| [Maximum Sum BST in Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/1373_Maximum_Sum_BST_in_Binary_Tree) | BinarySearchTree  BinaryTree  DepthFirstSearch  DynamicProgramming  Tree | 找到和最大的BST子树 |
|1392| [Longest Happy Prefix](https://github.com/ZhengjunHUO/leetcode/tree/main/1392_Longest_Happy_Prefix) | HashFunction  RollingHash  String  StringMatching | KMP算法的应用 |
|1408| [String Matching in an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/1408_String_Matching_in_an_Array) | String  StringMatching | 在一组字符串中找出所有能成为其他字符串的子串的串 |
|1455| [Check If a Word Occurs As a Prefix of Any Word in a Sentence](https://github.com/ZhengjunHUO/leetcode/tree/main/1455_Check_If_a_Word_Occurs_As_a_Prefix_of_Any_Word_in_a_Sentence) | String  StringMatching | 找到以某个pattern为前缀的词的序号 |
|1514| [Path with Maximum Probability](https://github.com/ZhengjunHUO/leetcode/tree/main/1514_Path_with_Maximum_Probability) | Graph  Heap  ShortestPath | 找到到达终点成功率最高的路径 |
|1541| [Minimum Insertions to Balance a Parentheses String](https://github.com/ZhengjunHUO/leetcode/tree/main/1541_Minimum_Insertions_to_Balance_a_Parentheses_String) | Greedy  Stack  String | 添加最小的括号数时括号合法成对 |
|1590| [Make Sum Divisible by P](https://github.com/ZhengjunHUO/leetcode/tree/main/1590_Make_Sum_Divisible_by_P) | Array  BinarySearch  HashTable  Math | 字符串中删除最小的子串使剩下的元素和可以整除目标值 |
|1615| [Maximal Network Rank](https://github.com/ZhengjunHUO/leetcode/tree/main/1615_Maximal_Network_Rank) | Graph | 图中两个节点出度和的最大值 |
|1631| [Path With Minimum Effort](https://github.com/ZhengjunHUO/leetcode/tree/main/1631_Path_With_Minimum_Effort) | Array  BinarySearch  BreadthFirstSearch  DepthFirstSearch  Heap  Matrix  UnionFind | 寻找一条消耗体力最少的登山路径 |
|1636| [Sort Array by Increasing Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/1636_Sort_Array_by_Increasing_Frequency) | Array  HashTable  Sorting | 按照元素出现的频率对字符串升序排序，如果频率相同则按数字大小降序排列 |
|1658| [Minimum Operations to Reduce X to Zero](https://github.com/ZhengjunHUO/leetcode/tree/main/1658_Minimum_Operations_to_Reduce_X_to_Zero) | BinarySearch  Greedy  SlidingWindow  TwoPointers | 每一步可以移除字符串头或尾的元素，目标值减去该元素，求将目标值减为0的最少步数 |
|1668| [Maximum Repeating Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/1668_Maximum_Repeating_Substring) | String  StringMatching | 保持作为某字符串子串的前提下pattern自身最多能重复多少次 |
|1679| [Max Number of K-Sum Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/1679_Max_Number_of_K-Sum_Pairs) | HashTable | 数组中每一步可以移除和为K的两个元素，求最大可执行次数 |
|1695| [Maximum_Erasure_Value](https://github.com/ZhengjunHUO/leetcode/tree/main/1695_Maximum_Erasure_Value) | TwoPointers | 找到不含重复元素的且和为最大的子串 |
|1696| [Jump Game VI](https://github.com/ZhengjunHUO/leetcode/tree/main/1696_Jump_Game_VI) | Dequeue | 从头出发，每次最远可以前进k，求到终点经过的index对应的值之和的最大值 |
|1711| [Count Good Meals](https://github.com/ZhengjunHUO/leetcode/tree/main/1711_Count_Good_Meals) | Array  HashTable  TwoPointers | 数列中和为2的次方的二元组的数量 |
|1749| [Maximum Absolute Sum of Any Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/1749_Maximum_Absolute_Sum_of_Any_Subarray) | Greedy | 和的绝对值最大的子串 |
|1786| [Number of Restricted Paths From First to Last Node](https://github.com/ZhengjunHUO/leetcode/tree/main/1786_Number_of_Restricted_Paths_From_First_to_Last_Node) | DynamicProgramming  Graph  Heap  ShortestPath  TopologicalSort | 起点到终点之间的有条件的路径数 |
|1791| [Find Center of Star Graph](https://github.com/ZhengjunHUO/leetcode/tree/main/1791_Find_Center_of_Star_Graph) | Graph | 找出星型图的中点 |
|1871| [Jump Game VII](https://github.com/ZhengjunHUO/leetcode/tree/main/1871_Jump_Game_VII) | BreadthFirstSearch  Greedy  LineSweep | 从一个01字符串的头部出发，每次可以前进的距离介于[i,j]之间，且只能落在值为0的index上，求是否可以到达终点 |
|1891| [Cutting Ribbons](https://github.com/ZhengjunHUO/leetcode/tree/main/1891_Cutting_Ribbons) | Array  BinarySearch | 切割一些绳子获得k条相同长度的绳子，求其长度的最大值 |
|1928| [Minimum Cost to Reach Destination in Time](https://github.com/ZhengjunHUO/leetcode/tree/main/1928_Minimum_Cost_to_Reach_Destination_in_Time) | DynamicProgramming  Graph | 在规定时间内到达目的地的最小花费 |
|1971| [Find if Path Exists in Graph](https://github.com/ZhengjunHUO/leetcode/tree/main/1971_Find_if_Path_Exists_in_Graph) | BreadthFirstSearch  DepthFirstSearch  Graph | 判断图中两点是否有通路 |
|1976| [Number of Ways to Arrive at Destination](https://github.com/ZhengjunHUO/leetcode/tree/main/1976_Number_of_Ways_to_Arrive_at_Destination) | DynamicProgramming  Graph  ShortestPath  TopologicalSort | 到达终点最短的距离的路径数 |
