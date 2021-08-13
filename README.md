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
|0073| [Set Matrix Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0073_Set_Matrix_Zeroes) | Array  HashTable  Matrix | 将矩阵中值为0的格子的行和列都置0 |
|0074| [Search a 2D Matrix](https://github.com/ZhengjunHUO/leetcode/tree/main/0074_Search_a_2D_Matrix) | Array  BinarySearch  Matrix | 在排好序的矩阵中找元素 |
|0075| [Sort Colors](https://github.com/ZhengjunHUO/leetcode/tree/main/0075_Sort_Colors) | Array  Sorting  TwoPointers | 三种颜色分类排序 |
|0076| [Minimum Window Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/0076_Minimum_Window_Substring) | HashTable  SlidingWindow  String  TwoPointers | 在A字符串中找到包含所有B字符串字母的最小子串 |
|0077| [Combinations](https://github.com/ZhengjunHUO/leetcode/tree/main/0077_Combinations) | Array  Backtracking | 在1到n的范围内找出所有长度为k的组合 |
|0078| [Subsets](https://github.com/ZhengjunHUO/leetcode/tree/main/0078_Subsets) | Array  Backtracking  BitManipulation | 找出数列的所有可能子集 |
|0079| [Word Search](https://github.com/ZhengjunHUO/leetcode/tree/main/0079_Word_Search) | Array  Backtracking  Matrix | 矩阵中寻找一个单词 |
|0080| [Remove Duplicates from Sorted Array II](https://github.com/ZhengjunHUO/leetcode/tree/main/0080_Remove_Duplicates_from_Sorted_Array_II) | Array  TwoPointers | 移除数组中多余的元素(in place) |
|0084| [Largest Rectangle in Histogram](https://github.com/ZhengjunHUO/leetcode/tree/main/0084_Largest_Rectangle_in_Histogram) | Array  MonotonicStack  Stack | 直方图找到最到矩形 |
|0088| [Merge Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0088_Merge_Sorted_Array) | Array  Sorting  TwoPointers | 将两个排好序的字符串in-place融合 |
|0090| [Subsets II](https://github.com/ZhengjunHUO/leetcode/tree/main/0090_Subsets_II) | Array  Backtracking  BitManipulation | 找出含重复元素数列的所有可能的不重复子集 |
|0092| [Reverse Linked List II](https://github.com/ZhengjunHUO/leetcode/tree/main/0092_Reverse_Linked_List_II) | LinkedList | 反转部分链表 |
|0093| [Restore IP Addresses](https://github.com/ZhengjunHUO/leetcode/tree/main/0093_Restore_IP_Addresses) | Backtracking  String | 从字符串还原IP地址 |
|0095| [Unique Binary Search Trees II](https://github.com/ZhengjunHUO/leetcode/tree/main/0095_Unique_Binary_Search_Trees_II) | Backtracking  BinarySearchTree  BinaryTree  DynamicProgramming  Tree | 返回从n个结点能构造的所有BST |
|0096| [Unique Binary Search Trees](https://github.com/ZhengjunHUO/leetcode/tree/main/0096_Unique_Binary_Search_Trees) | BinarySearchTree  BinaryTree  DynamicProgramming  Math  Tree | 计算n个结点可以构造多少个BST |
|0098| [Validate Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0098_Validate_Binary_Search_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 判断树是否为BST |
|0105| [Construct Binary Tree from Preorder and Inorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal) | Array  BinaryTree  DivideAndConquer  HashTable  Tree | 从前序遍历和中序遍历列表构建树 |
|0106| [Construct Binary Tree from Inorder and Postorder Traversal](https://github.com/ZhengjunHUO/leetcode/tree/main/0106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal) | Array  BinaryTree  DivideAndConquer  HashTable  Tree | 从中序遍历和后序遍历列表构建树 |
|0111| [Minimum Depth of Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0111_Minimum_Depth_of_Binary_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 求二叉树的最短深度 |
|0114| [Flatten Binary Tree to Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0114_Flatten_Binary_Tree_to_Linked_List) | BinaryTree  DepthFirstSearch  LinkedList  Stack  Tree | 把树压成链表 |
|0116| [Populating Next Right Pointers in Each Node](https://github.com/ZhengjunHUO/leetcode/tree/main/0116_Populating_Next_Right_Pointers_in_Each_Node) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 树的同一层结点横向连接起来 |
|0121| [Best Time to Buy and Sell Stock](https://github.com/ZhengjunHUO/leetcode/tree/main/0121_Best_Time_to_Buy_and_Sell_Stock) | Array  DynamicProgramming | 一次买卖交易获得的最大值 |
|0122| [Best Time to Buy and Sell Stock II](https://github.com/ZhengjunHUO/leetcode/tree/main/0122_Best_Time_to_Buy_and_Sell_Stock_II) | Array  Greedy | 不限次数交易获得的最大值 |
|0123| [Best Time to Buy and Sell Stock III](https://github.com/ZhengjunHUO/leetcode/tree/main/0123_Best_Time_to_Buy_and_Sell_Stock_III) | Array  DynamicProgramming | 至多两次交易获得的最大值 |
|0131| [Palindrome Partitioning](https://github.com/ZhengjunHUO/leetcode/tree/main/0131_Palindrome_Partitioning) | Backtracking  DynamicProgramming  String | 求字符串拆分成回文字符串数组的所有可能性 |
|0141| [Linked List Cycle](https://github.com/ZhengjunHUO/leetcode/tree/main/0141_Linked_List_Cycle) | HashTable  LinkedList  TwoPointers | 判断链表是否有环 |
|0142| [Linked List Cycle II](https://github.com/ZhengjunHUO/leetcode/tree/main/0142_Linked_List_Cycle_II) | HashTable  LinkedList  TwoPointers | 判断链表是否有环并找出环开始的地方 |
|0152| [Maximum Product Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0152_Maximum_Product_Subarray) | Array  DynamicProgramming | 找出所有元素乘积最大的子串 |
|0155| [Min Stack](https://github.com/ZhengjunHUO/leetcode/tree/main/0155_Min_Stack) | Design  Stack | 构建能快速返回最小值的栈 |
|0167| [Two Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0167_Two_Sum_II) | Array  BinarySearch  TwoPointers | 递增数列中找出和为目标值的两个元素 |
|0170| [Two Sum III Data Structure Design](https://github.com/ZhengjunHUO/leetcode/tree/main/0170_Two_Sum_III_Data_Structure_Design) | Design  HastTable | 设计可以插入数据并判断存储中是否有和为给定值的数据对的数据结构 |
|0187| [Repeated DNA Sequences](https://github.com/ZhengjunHUO/leetcode/tree/main/0187_Repeated_DNA_Sequences) | BitManipulation  HashFunction  HashTable  RollingHash  SlidingWindow  String | 找到至少出现两次的基因段 |
|0188| [Best Time to Buy and Sell Stock IV](https://github.com/ZhengjunHUO/leetcode/tree/main/0188_Best_Time_to_Buy_and_Sell_Stock_IV) | DynamicProgramming | 至多K次交易获得的最大值 |
|0190| [Reverse Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0190_Reverse_Bits) | BitManipulation | 字符串表示的二进制数取反 |
|0191| [Number of 1 Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0191_Number_of_1_Bits) | BitManipulation | 字符串表示的二进制数中1的个数 |
|0192| [Word Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/0192_Word_Frequency) | Shell | 统计词出现个数 |
|0193| [Valid Phone Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0193_Valid_Phone_Numbers) | Shell | 从文本中找出合法的电话号码 |
|0195| [Tenth Line](https://github.com/ZhengjunHUO/leetcode/tree/main/0195_Tenth_Line) | Shell | 打印出文本的第10行 |
|0198| [House Robber](https://github.com/ZhengjunHUO/leetcode/tree/main/0198_House_Robber) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0204| [Count Primes](https://github.com/ZhengjunHUO/leetcode/tree/main/0204_Count_Primes) | Array  Enumeration  Math  NumberTheory | 统计小于n的质数个数 |
|0206| [Reverse Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0206_Reverse_Linked_List) | LinkedList  Recursion | 反转链表 |
|0209| [Minimum Size Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0209_Minimum_Size_Subarray_Sum) | Array  BinarySearch  TwoPointers | 求和大于等于目标值的最短子串 |
|0213| [House Robber II](https://github.com/ZhengjunHUO/leetcode/tree/main/0213_House_Robber_II) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0214| [Shortest Palindrome](https://github.com/ZhengjunHUO/leetcode/tree/main/0214_Shortest_Palindrome) | HashFunction  RollingHash  String  StringMatching  | 找到从字符串头开始最长的回文组 |
|0215| [Kth Largest Element in an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0215_Kth_Largest_Element_in_an_Array) | Array  DivideAndConquer  Heap  Quickselect  Sorting | 数列中寻找第k大元素 |
|0226| [Invert Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0226_Invert_Binary_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 二叉树元素镜像转置 |
|0230| [Kth Smallest Element in a BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0230_Kth_Smallest_Element_in_a_BST) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 在BST中找到第k小的元素 |
|0234| [Palindrome Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0234_Palindrome_Linked_List) | LinkedList  Recursion  Stack  TwoPointers | 判断链表是否是回文 |
|0238| [Product of Array Except Self](https://github.com/ZhengjunHUO/leetcode/tree/main/0238_Product_of_Array_Except_Self) | Array | 返回数列所有元素乘积和除以各元素的数列 |
|0239| [Sliding Window Maximum](https://github.com/ZhengjunHUO/leetcode/tree/main/0239_Sliding_Window_Maximum) | Dequeue  Heap  SlidingWindow | 返回定长的向右滑动的窗口每一步的局部最大值组成的数组 |
|0242| [Valid Anagram](https://github.com/ZhengjunHUO/leetcode/tree/main/0242_Valid_Anagram) | HastTable  Sort | 判断一个字符串是否为另一个的易位词 |
|0252| [Meeting Rooms](https://github.com/ZhengjunHUO/leetcode/tree/main/0252_Meeting_Rooms) | Array  Sorting | 会议时间是否重合 |
|0259| [3Sum Smaller](https://github.com/ZhengjunHUO/leetcode/tree/main/0259_3Sum_Smaller) | Array  TwoPointers | 求数组中三元组和小于目标值的个数 |
|0263| [Ugly Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0263_Ugly_Number) | Math | 判断一个数的质因数是否限定在2,3,5之中 |
|0264| [Ugly Number II](https://github.com/ZhengjunHUO/leetcode/tree/main/0264_Ugly_Number_II) | DynamicProgramming  HashTable  Heap  Math | 求从1开始第n个题0263中描述的数字 |
|0279| [Perfect Squares](https://github.com/ZhengjunHUO/leetcode/tree/main/0279_Perfect_Squares) | BreadthFirstSearch  DynamicProgramming  Math | 把数分解为最短的平方和 |
|0283| [Move Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0283_Move_Zeroes) | Array  TwoPointers | 将数列中的0移到最后 |
|0295| [Find Median from Data Stream](https://github.com/ZhengjunHUO/leetcode/tree/main/0295_Find_Median_from_Data_Stream) | DataStream  Design  Heap  Sorting  TwoPointers | 数据流中求中位数 |
|0309| [Best Time to Buy and Sell Stock with Cooldown](https://github.com/ZhengjunHUO/leetcode/tree/main/0309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown) | DynamicProgramming | 不限次数交易获得的最大值，卖出后有一天冷却期 |
|0313| [Super Ugly Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0313_Super_Ugly_Number) | Array  DynamicProgramming  HashTable  Heap  Math | |
|0322| [Coin Change](https://github.com/ZhengjunHUO/leetcode/tree/main/0322_Coin_Change) | Array  BreadthFirstSearch  DynamicProgramming | 最少能用几枚硬币凑成目标数额 |
|0325| [Maximum Size Subarray Sum Equals k](https://github.com/ZhengjunHUO/leetcode/tree/main/0325_Maximum_Size_Subarray_Sum_Equals_k) | ??? | 求和等于目标值的最长子串 |
|0337| [House Robber III](https://github.com/ZhengjunHUO/leetcode/tree/main/0337_House_Robber_III) | DepthFirstSearch  DynamicProgramming  Tree | 抢劫树形排列不相邻房屋获得的最大收益 |
|0344| [Reverse String](https://github.com/ZhengjunHUO/leetcode/tree/main/0344_Reverse_String) | Recursion  String  TwoPointers | 将字符数组反转 |
|0347| [Top K Frequent Elements](https://github.com/ZhengjunHUO/leetcode/tree/main/0347_Top_K_Frequent_Elements) | Array  BucketSort  Counting  DivideAndConquer  HashTable  Heap  Quickselect  Sorting | 返回前k个最常见的元素 |
|0355| [Design Twitter](https://github.com/ZhengjunHUO/leetcode/tree/main/0355_Design_Twitter) | Design  HashTable  Heap  LinkedList | 实现简单的推特应用 |
|0371| [Sum of Two Integers](https://github.com/ZhengjunHUO/leetcode/tree/main/0371_Sum_of_Two_Integers) | BitManipulation | 位操作实现两数相加 |
|0373| [Find K Pairs with Smallest Sums](https://github.com/ZhengjunHUO/leetcode/tree/main/0373_Find_K_Pairs_with_Smallest_Sums) | Array  Heap | 两个数列中找到和最小的k对元素 |
|0378| [Kth Smallest Element in a Sorted Matrix](https://github.com/ZhengjunHUO/leetcode/tree/main/0378_Kth_Smallest_Element_in_a_Sorted_Matrix) | Array  BinarySearch  Heap  Matrix  Sorting | 在行和列都是升序的矩阵中找到第k小的元素 |
|0387| [First Unique Character in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0387_First_Unique_Character_in_a_String) | Counting  HashTable  Queue  String | 找到第一个只出现一次的字符 |
|0407| [Trapping Rain Water II](https://github.com/ZhengjunHUO/leetcode/tree/main/0407_Trapping_Rain_Water_II) | BreadthFirstSearch  Heap | 计算三维地形的积水量 |
|0414| [Third Maximum Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0414_Third_Maximum_Number) | Array  Sorting | 第三大的数字 |
|0415| [Add Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0415_Add_Strings) | String | 以字符串形式表示的两数之和 |
|0435| [Non-overlapping Intervals](https://github.com/ZhengjunHUO/leetcode/tree/main/0435_Non-overlapping_Intervals) | Array  DynamicProgramming  Greedy  Sorting | 移除最少的区间使列表中的区间没有重叠 |
|0438| [Find All Anagrams in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0438_Find_All_Anagrams_in_a_String) | HastTable | 在字符串A中找到所有字符串B的易位词 | 
|0445| [Add Two Numbers II](https://github.com/ZhengjunHUO/leetcode/tree/main/0445_Add_Two_Numbers_II) | LinkedList | 以链表形式表示的两数之和 |
|0450| [Delete Node in a BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0450_Delete_Node_in_a_BST) | BinarySearchTree  BinaryTree  Tree | 在BST中删除一个结点 |
|0451| [Sort Characters By Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/0451_Sort_Characters_By_Frequency) | BucketSort  Counting  HashTable  Heap  Sorting  String | 按字符出现的频率降序排列字符串 |
|0454| [4Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0454_4Sum_II) | BinarySearch  HastTable | 四个数列中分别取一个数q之和等于0的四元组个数 |
|0459| [Repeated Substring Pattern](https://github.com/ZhengjunHUO/leetcode/tree/main/0459_Repeated_Substring_Pattern) | String | 检查字符串是否由某个子串重复n次构成 |
|0474| [Ones and Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0474_Ones_and_Zeroes) | DynamicProgramming | 字符串表示的二进制数组中，找出最大的包含至多m个0和n个1的子集合 |
|0495| [Teemo Attacking](https://github.com/ZhengjunHUO/leetcode/tree/main/0495_Teemo_Attacking) | Array  Simulation | 计算进入异常状态的时间 |
|0496| [Next Greater Element I](https://github.com/ZhengjunHUO/leetcode/tree/main/0496_Next_Greater_Element_I) | Array  HashTable  MonotonicStack  Stack | 找到下一个大于自己的元素 |
|0502| [IPO](https://github.com/ZhengjunHUO/leetcode/tree/main/0502_IPO) | Array  Greedy  Heap  Sorting | 完成k个项目可以获得的最大纯利润 |
|0503| [Next Greater Element II](https://github.com/ZhengjunHUO/leetcode/tree/main/0503_Next_Greater_Element_II) | Array  MonotonicStack  Stack | 首尾循环的列表寻找下一个大于自己的元素 |
|0506| [Relative Ranks](https://github.com/ZhengjunHUO/leetcode/tree/main/0506_Relative_Ranks) | Array  Heap  Sorting | 把分数转换为名次 |
|0509| [Fibonacci Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0509_Fibonacci_Number) | DynamicProgramming  Math  Memoization | 斐波那契数列 |
|0523| [Continuous Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0523_Continuous_Subarray_Sum) | DynamicProgramming  Math | 求是否有一个长度至少为2的子串其和为给定值的倍数 |
|0538| [Convert BST to Greater Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0538_Convert_BST_to_Greater_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 结点的值更新为整棵树所有大于等于该结点的值的和 |
|0560| [Subarray Sum Equals K](https://github.com/ZhengjunHUO/leetcode/tree/main/0560_Subarray_Sum_Equals_K) | Array  HashTable | 求和为目标值的子串个数 |
|0567| [Permutation in String](https://github.com/ZhengjunHUO/leetcode/tree/main/0567_Permutation_in_String) | TwoPointer  SlidingWindow | 求字符串A是否包含字符串B的变种 |
|0600| [Non-negative Integers without Consecutive Ones](https://github.com/ZhengjunHUO/leetcode/tree/main/0600_Non-negative_Integers_without_Consecutive_Ones) | DynamicProgramming | 从0到给定值之间所有数的二进制表示中没有连续1的数的个数 |
|0628| [Maximum Product of Three Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0628_Maximum_Product_of_Three_Numbers) | Array  Math | 数列中挑选三个数使其乘积最大 |
|0632| [Smallest Range Covering Elements from K Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0632_Smallest_Range_Covering_Elements_from_K_Lists) | HashTable  String  TwoPointers | 找到一个最小范围覆盖所有K个数列中至少一个元素 |
|0652| [Find Duplicate Subtrees](https://github.com/ZhengjunHUO/leetcode/tree/main/0652_Find_Duplicate_Subtrees) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 寻找重复的子树 |
|0653| [Two Sum IV Input Is A BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0653_Two_Sum_IV_Input_Is_A_BST) | Tree | 在BST中找到等于目标值的两个数 |
|0654| [Maximum Binary Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0654_Maximum_Binary_Tree) | Array  BinaryTree  DivideAndConquer  MonotonicStack  Stack  Tree | 按要求构建二叉树 |
|0686| [Repeated String Match](https://github.com/ZhengjunHUO/leetcode/tree/main/0686_Repeated_String_Match) | String | a字符串重复多少次可以包含b字符串 |
|0689| [Maximum Sum of 3 Non-Overlapping Subarrays](https://github.com/ZhengjunHUO/leetcode/tree/main/0689_Maximum_Sum_of_3_Non-Overlapping_Subarrays) | Array  DynamicProgramming | 找到三个指定长度的和最大的独立子串 |
|0692| [Top K Frequent Words](https://github.com/ZhengjunHUO/leetcode/tree/main/0692_Top_K_Frequent_Words) | BucketSort  Counting  HashTable  Heap  Sorting  String  Trie | 出现频率最高的k个词 |
|0697| [Degree of an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0697_Degree_of_an_Array) | Array | 找到和字符串有相同度的子串 |
|0700| [Search in a Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0700_Search_in_a_Binary_Search_Tree) | BinarySearchTree  BinaryTree  Tree | 在BST中搜索一个元素 |
|0701| [Insert into a Binary Search Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/0701_Insert_into_a_Binary_Search_Tree) | BinarySearchTree  BinaryTree  Tree | 在BST中插入一个元素 |
|0703| [Kth Largest Element in a Stream](https://github.com/ZhengjunHUO/leetcode/tree/main/0703_Kth_Largest_Element_in_a_Stream) | BinarySearchTree  BinaryTree  DataStream  Design  Heap  Tree | 不断插入新元素同时找到第k大的元素 |
|0704| [Binary Search](https://github.com/ZhengjunHUO/leetcode/tree/main/0704_Binary_Search) | Array  BinarySearch | 二分搜索元素 |
|0713| [Subarray Product Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/0713_Subarray_Product_Less_Than_K) | Array  TwoPointers | 乘积小于目标值的子串的个数 |
|0714| [Best Time to Buy and Sell Stock with Transaction Fee](https://github.com/ZhengjunHUO/leetcode/tree/main/0714_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee) | Array  DynamicProgramming  Greedy | 不限次数交易获得的最大值，每笔交易有手续费 |
|0718| [Maximum Length of Repeated Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0718_Maximum_Length_of_Repeated_Subarray) | Array  BinarySearch  DynamicProgramming  HashTable | 寻找最大公共子串 | 
|0724| [Find Pivot Index](https://github.com/ZhengjunHUO/leetcode/tree/main/0724_Find_Pivot_Index) | Array | 找到左边所有元素和等于右边所有元素和的index |
|0727| [Minimum Window Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0727_Minimum_Window_Subsequence) | DynamicProgramming  HashTable  SlidingWindow  String  TwoPointers | 字符串A中找到最短的可以包含字符串B的子串
|0739| [Daily Temperatures](https://github.com/ZhengjunHUO/leetcode/tree/main/0739_Daily_Temperatures) | Array  MonotonicStack  Stack | 求等待几天可以比今天更热 |
|0740| [Delete and Earn](https://github.com/ZhengjunHUO/leetcode/tree/main/0740_Delete_and_Earn) | DynamicProgramming | 在数组中取一个数可得分，但是要删除它和如果存在的它+-1的值，求最大得分 |
|0752| [Open the Lock](https://github.com/ZhengjunHUO/leetcode/tree/main/0752_Open_the_Lock) | Array  BreadthFirstSearch  HashTable  String | 密码锁至少需要转动几下才能打开 |
|0763| [Partition Labels](https://github.com/ZhengjunHUO/leetcode/tree/main/0763_Partition_Labels) | Greedy  HashTable  String  TwoPointers | 尽可能将字符串分段，且每个字母仅出现在某段中 |
|0784| [Letter Case Permutation](https://github.com/ZhengjunHUO/leetcode/tree/main/0784_Letter_Case_Permutation) | Backtracking  BitManipulation  String | alnum字符串中字母大小写变换的排列组合 |
|0876| [Middle of the Linked List](https://github.com/ZhengjunHUO/leetcode/tree/main/0876_Middle_of_the_Linked_List) | LinkedList  TwoPointers | 寻找链表的中点 |
|0912| [Sort an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0912_Sort_an_Array) | Array  BucketSort  CountingSort  DivideAndConquer  Heap  MergeSort  RadixSort  Sorting | 给数列排序 |
|0974| [Subarray Sums Divisible by K](https://github.com/ZhengjunHUO/leetcode/tree/main/0974_Subarray_Sums_Divisible_by_K) | Array  HashTable | 和可以整除目标值的子串的个数 |
|0977| [Squares of a Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0977_Squares_of_a_Sorted_Array) | Array  Sorting  TwoPointers | 求排好序的数组各元素平方后得到的新升序数组 |
|0978| [Longest Turbulent Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0978_Longest_Turbulent_Subarray) | Array  DynamicProgramming  SlidingWindow | 找到符合峰谷交替规律的最长的子串 |
|0986| [Interval List Intersections](https://github.com/ZhengjunHUO/leetcode/tree/main/0986_Interval_List_Intersections) | Array  TwoPointers | 两个区间组成的列表找交集 |
|0989| [Add to Array-Form of Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0989_Add_to_Array-Form_of_Integer) | Array | 求一个以数组形式表示的数和一个整数之和 |
|1038| [Binary Search Tree to Greater Sum Tree](https://github.com/ZhengjunHUO/leetcode/tree/main/1038_Binary_Search_Tree_to_Greater_Sum_Tree) | BinaryTree  BreadthFirstSearch  DepthFirstSearch  Tree | 结点的值更新为整棵树所有大于等于该结点的值的和 |
|1099| [Two Sum Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/1099_Two_Sum_Less_Than_K) | Array | 数组中找到小于目标值的最大两数之和 | 
|1201| [Ugly Number III](https://github.com/ZhengjunHUO/leetcode/tree/main/1201_Ugly_Number_III) | BinarySearch  Math  NumberTheory | |
|1288| [Remove Covered Intervals](https://github.com/ZhengjunHUO/leetcode/tree/main/1288_Remove_Covered_Intervals) | Array  Sorting | 移除能被另一个区间覆盖的区间 |
|1306| [Jump Game III](https://github.com/ZhengjunHUO/leetcode/tree/main/1306_Jump_Game_III) | BreadthFirstSearch  DepthFirstSearch  Recursion | 从给定的出发点允许向前向后跳跃，判断是否能到终点 |
|1340| [Jump Game V](https://github.com/ZhengjunHUO/leetcode/tree/main/1340_Jump_Game_V) | DynamicProgramming | 从任意点出发允许向前向后最远跳d的距离且只能从大数跳向小数，求最多能覆盖几个index ｜
|1345| [Jump Game IV](https://github.com/ZhengjunHUO/leetcode/tree/main/1345_Jump_Game_IV) | BreadthFirstSearch | 从头开始，每一条可以向前向后一格，或传送到其他值相等的index，求到达终点最少的跳跃次数 |
|1590| [Make Sum Divisible by P](https://github.com/ZhengjunHUO/leetcode/tree/main/1590_Make_Sum_Divisible_by_P) | Array  BinarySearch  HashTable  Math | 字符串中删除最小的子串使剩下的元素和可以整除目标值 |
|1636| [Sort Array by Increasing Frequency](https://github.com/ZhengjunHUO/leetcode/tree/main/1636_Sort_Array_by_Increasing_Frequency) | Array  HashTable  Sorting | 按照元素出现的频率对字符串升序排序，如果频率相同则按数字大小降序排列 |
|1658| [Minimum Operations to Reduce X to Zero](https://github.com/ZhengjunHUO/leetcode/tree/main/1658_Minimum_Operations_to_Reduce_X_to_Zero) | BinarySearch  Greedy  SlidingWindow  TwoPointers | 每一步可以移除字符串头或尾的元素，目标值减去该元素，求将目标值减为0的最少步数 |
|1679| [Max Number of K-Sum Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/1679_Max_Number_of_K-Sum_Pairs) | HashTable | 数组中每一步可以移除和为K的两个元素，求最大可执行次数 |
|1695| [Maximum_Erasure_Value](https://github.com/ZhengjunHUO/leetcode/tree/main/1695_Maximum_Erasure_Value) | TwoPointers | 找到不含重复元素的且和为最大的子串 |
|1696| [Jump Game VI](https://github.com/ZhengjunHUO/leetcode/tree/main/1696_Jump_Game_VI) | Dequeue | 从头出发，每次最远可以前进k，求到终点经过的index对应的值之和的最大值 |
|1711| [Count Good Meals](https://github.com/ZhengjunHUO/leetcode/tree/main/1711_Count_Good_Meals) | Array  HashTable  TwoPointers | 数列中和为2的次方的二元组的数量 |
|1749| [Maximum Absolute Sum of Any Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/1749_Maximum_Absolute_Sum_of_Any_Subarray) | Greedy | 和的绝对值最大的子串 |
|1871| [Jump Game VII](https://github.com/ZhengjunHUO/leetcode/tree/main/1871_Jump_Game_VII) | BreadthFirstSearch  Greedy  LineSweep | 从一个01字符串的头部出发，每次可以前进的距离介于[i,j]之间，且只能落在值为0的index上，求是否可以到达终点 |
