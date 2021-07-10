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
|0015| [Three Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0015_Three_Sum) | Array  TwoPointers | 求所有和为0的三元组 |
|0016| [3Sum Closest](https://github.com/ZhengjunHUO/leetcode/tree/main/0016_3Sum_Closest) | Array  TwoPointers | 求和最接近目标值的三元组 |
|0017| [Letter Combinations of a Phone Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0017_Letter_Combinations_of_a_Phone_Number) | Backtracking  HashTable  String | 根据电话号码找到对应的字母组合 |
|0018| [Four Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0018_Four_Sum) | Array  HashTable  TwoPointers | 找出所有和为目标值的四元组 |
|0020| [Valid Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0020_Valid_Parentheses) | Stack  String | 检查括号是否正确闭合 |
|0021| [Merge Two Sorted Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0021_Merge_Two_Sorted_Lists) | LinkedList  Recursion | 融合两个排好序的链表 |
|0024| [Swap Nodes in Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/0024_Swap_Nodes_in_Pairs) | LinkedList  Recursion | 链表相邻两个结点对调 |
|0028| [Implement strStr()](https://github.com/ZhengjunHUO/leetcode/tree/main/0028_Implement_strStr) | String  TwoPointers | 找出pattern在字符串中第一次出现的位置 |
|0030| [Substring with Concatenation of All Words](https://github.com/ZhengjunHUO/leetcode/tree/main/0030_Substring_with_Concatenation_of_All_Words) | HashTable  String  TwoPointers | 找出内容为"以任意顺序拼接的所有给定单词"的子串 |
|0031| [Next Permutation](https://github.com/ZhengjunHUO/leetcode/tree/main/0031_Next_Permutation) | Array  TwoPointers | 按字典顺序寻找当前数组的后一个排列 |
|0032| [Longest Valid Parentheses](https://github.com/ZhengjunHUO/leetcode/tree/main/0032_Longest_Valid_Parentheses) | DynamicProgramming  Stack  String | 最长的合法括号子串 |
|0033| [Search in Rotated Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0033_Search_in_Rotated_Sorted_Array) | Array  BinarySearch | 在一个循环升序数列中寻找某数 |
|0034| [Find First and Last Position of Element in Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0034_Find_First_and_Last_Position_of_Element_in_Sorted_Array) | Array  BinarySearch | 在有序数列中找到某元素的起始和中止index |
|0035| [Search Insert Position](https://github.com/ZhengjunHUO/leetcode/tree/main/0035_Search_Insert_Position) | Array  BinarySearch | 为某数在一个升序数列中寻找插入位置 |
|0039| [Combination Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0039_Combination_Sum) | Array  Backtracking | 找出所有和为目标值的元素（元素可重复使用）组成的列表 |
|0040| [Combination Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0040_Combination_Sum_II) | Array  Backtracking | 找出所有和为目标值的元素（元素不可重复使用）组成的列表 |
|0042| [Trapping Rain Water](https://github.com/ZhengjunHUO/leetcode/tree/main/0042_Trapping_Rain_Water) | Array  Stack  TwoPointers  DynamicProgramming | 计算二维地形的积水量 |
|0043| [Multiply Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0043_Multiply_Strings) | Math  String | 计算以字符串描述的两数的乘积 |
|0045| [Jump Game II](https://github.com/ZhengjunHUO/leetcode/tree/main/0045_Jump_Game_II) | Array  Greedy | 最小跳数到达终点 |
|0046| [Permutations](https://github.com/ZhengjunHUO/leetcode/tree/main/0046_Permutations) | Array  Backtracking | 求数列中所有元素排列组合的集合 |
|0047| [Permutations II](https://github.com/ZhengjunHUO/leetcode/tree/main/0047_Permutations_II) | Array  Backtracking | 求含重复元素数列中所有元素排列组合的集合 |
|0049| [Group Anagrams](https://github.com/ZhengjunHUO/leetcode/tree/main/0049_Group_Anagrams) | HashTable  String | 将相关联的易位词组合起来 |
|0053| [Maximum Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0053_Maximum_Subarray) | Array  DivideAndConquer  DynamicProgramming | 找到和最大的子数列 |
|0055| [Jump Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0055_Jump_Game) | Array  Greedy | 求是否可以跳到终点 |
|0060| [Permutation Sequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0060_Permutation_Sequence) | Math  Recursion | 在数列的所有排列组合中按字典顺序找出第k个组合 |
|0065| [Valid Number](https://github.com/ZhengjunHUO/leetcode/tree/main/0065_Valid_Number) | Math  String | 检查字符串是否为合法的数 |
|0066| [Plus One](https://github.com/ZhengjunHUO/leetcode/tree/main/0066_Plus_One) | Array | 以数组描述的数加一 |
|0067| [Add Binary](https://github.com/ZhengjunHUO/leetcode/tree/main/0067_Add_Binary) | Math  String | 字符串形式的二进制数相加 |
|0076| [Minimum Window Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/0076_Minimum_Window_Substring) | HashTable  SlidingWindow  String  TwoPointers | 在A字符串中找到包含所有B字符串字母的最小子串 |
|0077| [Combinations](https://github.com/ZhengjunHUO/leetcode/tree/main/0077_Combinations) | Array  Backtracking | 在1到n的范围内找出所有长度为k的组合 |
|0078| [Subsets](https://github.com/ZhengjunHUO/leetcode/tree/main/0078_Subsets) | Array  Backtracking  BitManipulation | 找出数列的所有可能子集 |
|0088| [Merge Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0088_Merge_Sorted_Array) | Array  Sorting  TwoPointers | 将两个排好序的字符串in-place融合 |
|0090| [Subsets II](https://github.com/ZhengjunHUO/leetcode/tree/main/0090_Subsets_II) | Array  Backtracking  BitManipulation | 找出含重复元素数列的所有可能的不重复子集 |
|0121| [Best Time to Buy and Sell Stock](https://github.com/ZhengjunHUO/leetcode/tree/main/0121_Best_Time_to_Buy_and_Sell_Stock) | Array  DynamicProgramming | 一次买卖交易获得的最大值 |
|0122| [Best Time to Buy and Sell Stock II](https://github.com/ZhengjunHUO/leetcode/tree/main/0122_Best_Time_to_Buy_and_Sell_Stock_II) | Array  Greedy | 不限次数交易获得的最大值 |
|0123| [Best Time to Buy and Sell Stock III](https://github.com/ZhengjunHUO/leetcode/tree/main/0123_Best_Time_to_Buy_and_Sell_Stock_III) | Array  DynamicProgramming | 至多两次交易获得的最大值 |
|0131| [Palindrome Partitioning](https://github.com/ZhengjunHUO/leetcode/tree/main/0131_Palindrome_Partitioning) | Backtracking  DynamicProgramming  String | 求字符串拆分成回文字符串数组的所有可能性 |
|0152| [Maximum Product Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0152_Maximum_Product_Subarray) | Array  DynamicProgramming | 找出所有元素乘积最大的子串 |
|0155| [Min Stack](https://github.com/ZhengjunHUO/leetcode/tree/main/0155_Min_Stack) | Design  Stack | 构建能快速返回最小值的栈 |
|0167| [Two Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0167_Two_Sum_II) | Array  BinarySearch  TwoPointers | 递增数列中找出和为目标值的两个元素 |
|0170| [Two Sum III Data Structure Design](https://github.com/ZhengjunHUO/leetcode/tree/main/0170_Two_Sum_III_Data_Structure_Design) | Design  HastTable | 设计可以插入数据并判断存储中是否有和为给定值的数据对的数据结构 |
|0188| [Best Time to Buy and Sell Stock IV](https://github.com/ZhengjunHUO/leetcode/tree/main/0188_Best_Time_to_Buy_and_Sell_Stock_IV) | DynamicProgramming | 至多K次交易获得的最大值 |
|0190| [Reverse Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0190_Reverse_Bits) | BitManipulation | 字符串表示的二进制数取反 |
|0191| [Number of 1 Bits](https://github.com/ZhengjunHUO/leetcode/tree/main/0191_Number_of_1_Bits) | BitManipulation | 字符串表示的二进制数中1的个数 |
|0198| [House Robber](https://github.com/ZhengjunHUO/leetcode/tree/main/0198_House_Robber) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0209| [Minimum Size Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0209_Minimum_Size_Subarray_Sum) | Array  BinarySearch  TwoPointers | 求和大于等于目标值的最短子串 |
|0213| [House Robber II](https://github.com/ZhengjunHUO/leetcode/tree/main/0213_House_Robber_II) | DynamicProgramming | 抢劫不相邻房屋获得的最大收益 |
|0214| [Shortest Palindrome](https://github.com/ZhengjunHUO/leetcode/tree/main/0214_Shortest_Palindrome) | String | 找到从字符串头开始最长的回文组 |
|0238| [Product of Array Except Self](https://github.com/ZhengjunHUO/leetcode/tree/main/0238_Product_of_Array_Except_Self) | Array | 返回数列所有元素乘积和除以各元素的数列 |
|0239| [Sliding Window Maximum](https://github.com/ZhengjunHUO/leetcode/tree/main/0239_Sliding_Window_Maximum) | Dequeue  Heap  SlidingWindow | 返回定长的向右滑动的窗口每一步的局部最大值组成的数组 |
|0242| [Valid Anagram](https://github.com/ZhengjunHUO/leetcode/tree/main/0242_Valid_Anagram) | HastTable  Sort | 判断一个字符串是否为另一个的易位词 |
|0259| [3Sum Smaller](https://github.com/ZhengjunHUO/leetcode/tree/main/0259_3Sum_Smaller) | Array  TwoPointers | 求数组中三元组和小于目标值的个数 |
|0283| [Move Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0283_Move_Zeroes) | Array  TwoPointers | 将数列中的0移到最后 |
|0309| [Best Time to Buy and Sell Stock with Cooldown](https://github.com/ZhengjunHUO/leetcode/tree/main/0309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown) | DynamicProgramming | 不限次数交易获得的最大值，卖出后有一天冷却期 |
|0325| [Maximum Size Subarray Sum Equals k](https://github.com/ZhengjunHUO/leetcode/tree/main/0325_Maximum_Size_Subarray_Sum_Equals_k) | ??? | 求和等于目标值的最长子串 |
|0337| [House Robber III](https://github.com/ZhengjunHUO/leetcode/tree/main/0337_House_Robber_III) | DepthFirstSearch  DynamicProgramming  Tree | 抢劫树形排列不相邻房屋获得的最大收益 |
|0371| [Sum of Two Integers](https://github.com/ZhengjunHUO/leetcode/tree/main/0371_Sum_of_Two_Integers) | BitManipulation | 位操作实现两数相加 |
|0407| [Trapping Rain Water II](https://github.com/ZhengjunHUO/leetcode/tree/main/0407_Trapping_Rain_Water_II) | BreadthFirstSearch  Heap | 计算三维地形的积水量 |
|0415| [Add Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0415_Add_Strings) | String | 以字符串形式表示的两数之和 |
|0438| [Find All Anagrams in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0438_Find_All_Anagrams_in_a_String) | HastTable | 在字符串A中找到所有字符串B的易位词 | 
|0445| [Add Two Numbers II](https://github.com/ZhengjunHUO/leetcode/tree/main/0445_Add_Two_Numbers_II) | LinkedList | 以链表形式表示的两数之和 |
|0454| [4Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0454_4Sum_II) | BinarySearch  HastTable | 四个数列中分别取一个数q之和等于0的四元组个数 |
|0459| [Repeated Substring Pattern](https://github.com/ZhengjunHUO/leetcode/tree/main/0459_Repeated_Substring_Pattern) | String | 检查字符串是否由某个子串重复n次构成 |
|0474| [Ones and Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0474_Ones_and_Zeroes) | DynamicProgramming | 字符串表示的二进制数组中，找出最大的包含至多m个0和n个1的子集合 |
|0523| [Continuous Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0523_Continuous_Subarray_Sum) | DynamicProgramming  Math | 求是否有一个长度至少为2的子串其和为给定值的倍数 |
|0560| [Subarray Sum Equals K](https://github.com/ZhengjunHUO/leetcode/tree/main/0560_Subarray_Sum_Equals_K) | Array  HashTable | 求和为目标值的子串个数 |
|0567| [Permutation in String](https://github.com/ZhengjunHUO/leetcode/tree/main/0567_Permutation_in_String) | TwoPointer  SlidingWindow | 求字符串A是否包含字符串B的变种 |
|0600| [Non-negative Integers without Consecutive Ones](https://github.com/ZhengjunHUO/leetcode/tree/main/0600_Non-negative_Integers_without_Consecutive_Ones) | DynamicProgramming | 从0到给定值之间所有数的二进制表示中没有连续1的数的个数 |
|0628| [Maximum Product of Three Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0628_Maximum_Product_of_Three_Numbers) | Array  Math | 数列中挑选三个数使其乘积最大 |
|0632| [Smallest Range Covering Elements from K Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0632_Smallest_Range_Covering_Elements_from_K_Lists) | HashTable  String  TwoPointers | 找到一个最小范围覆盖所有K个数列中至少一个元素 |
|0653| [Two Sum IV Input Is A BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0653_Two_Sum_IV_Input_Is_A_BST) | Tree | 在BST中找到等于目标值的两个数 |
|0686| [Repeated String Match](https://github.com/ZhengjunHUO/leetcode/tree/main/0686_Repeated_String_Match) | String | a字符串重复多少次可以包含b字符串 |
|0689| [Maximum Sum of 3 Non-Overlapping Subarrays](https://github.com/ZhengjunHUO/leetcode/tree/main/0689_Maximum_Sum_of_3_Non-Overlapping_Subarrays) | Array  DynamicProgramming | 找到三个指定长度的和最大的独立子串 |
|0697| [Degree of an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0697_Degree_of_an_Array) | Array | 找到和字符串有相同度的子串 |
|0713| [Subarray Product Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/0713_Subarray_Product_Less_Than_K) | Array  TwoPointers | 乘积小于目标值的子串的个数 |
|0714| [Best Time to Buy and Sell Stock with Transaction Fee](https://github.com/ZhengjunHUO/leetcode/tree/main/0714_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee) | Array  DynamicProgramming  Greedy | 不限次数交易获得的最大值，每笔交易有手续费 |
|0718| [Maximum Length of Repeated Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0718_Maximum_Length_of_Repeated_Subarray) | Array  BinarySearch  DynamicProgramming  HashTable | 寻找最大公共子串 | 
|0724| [Find Pivot Index](https://github.com/ZhengjunHUO/leetcode/tree/main/0724_Find_Pivot_Index) | Array | 找到左边所有元素和等于右边所有元素和的index |
|0727| [Minimum Window Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0727_Minimum_Window_Subsequence) | DynamicProgramming  HashTable  SlidingWindow  String  TwoPointers | 字符串A中找到最短的可以包含字符串B的子串
|0740| [Delete and Earn](https://github.com/ZhengjunHUO/leetcode/tree/main/0740_Delete_and_Earn) | DynamicProgramming | 在数组中取一个数可得分，但是要删除它和如果存在的它+-1的值，求最大得分 |
|0784| [Letter Case Permutation](https://github.com/ZhengjunHUO/leetcode/tree/main/0784_Letter_Case_Permutation) | Backtracking  BitManipulation  String | alnum字符串中字母大小写变换的排列组合 |
|0974| [Subarray Sums Divisible by K](https://github.com/ZhengjunHUO/leetcode/tree/main/0974_Subarray_Sums_Divisible_by_K) | Array  HashTable | 和可以整除目标值的子串的个数 |
|0977| [Squares of a Sorted Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0977_Squares_of_a_Sorted_Array) | Array  Sorting  TwoPointers | 求排好序的数组各元素平方后得到的新升序数组 |
|0978| [Longest Turbulent Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0978_Longest_Turbulent_Subarray) | Array  DynamicProgramming  SlidingWindow | 找到符合峰谷交替规律的最长的子串 |
|0989| [Add to Array-Form of Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0989_Add_to_Array-Form_of_Integer) | Array | 求一个以数组形式表示的数和一个整数之和 |
|1099| [Two Sum Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/1099_Two_Sum_Less_Than_K) | Array | 数组中找到小于目标值的最大两数之和 | 
|1306| [Jump Game III](https://github.com/ZhengjunHUO/leetcode/tree/main/1306_Jump_Game_III) | BreadthFirstSearch  DepthFirstSearch  Recursion | 从给定的出发点允许向前向后跳跃，判断是否能到终点 |
|1340| [Jump Game V](https://github.com/ZhengjunHUO/leetcode/tree/main/1340_Jump_Game_V) | DynamicProgramming | 从任意点出发允许向前向后最远跳d的距离且只能从大数跳向小数，求最多能覆盖几个index ｜
|1345| [Jump Game IV](https://github.com/ZhengjunHUO/leetcode/tree/main/1345_Jump_Game_IV) | BreadthFirstSearch | 从头开始，每一条可以向前向后一格，或传送到其他值相等的index，求到达终点最少的跳跃次数 |
|1590| [Make Sum Divisible by P](https://github.com/ZhengjunHUO/leetcode/tree/main/1590_Make_Sum_Divisible_by_P) | Array  BinarySearch  HashTable  Math | 字符串中删除最小的子串使剩下的元素和可以整除目标值 |
|1658| [Minimum Operations to Reduce X to Zero](https://github.com/ZhengjunHUO/leetcode/tree/main/1658_Minimum_Operations_to_Reduce_X_to_Zero) | BinarySearch  Greedy  SlidingWindow  TwoPointers | 每一步可以移除字符串头或尾的元素，目标值减去该元素，求将目标值减为0的最少步数 |
|1679| [Max Number of K-Sum Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/1679_Max_Number_of_K-Sum_Pairs) | HashTable | 数组中每一步可以移除和为K的两个元素，求最大可执行次数 |
|1695| [Maximum_Erasure_Value](https://github.com/ZhengjunHUO/leetcode/tree/main/1695_Maximum_Erasure_Value) | TwoPointers | 找到不含重复元素的且和为最大的子串 |
|1696| [Jump Game VI](https://github.com/ZhengjunHUO/leetcode/tree/main/1696_Jump_Game_VI) | Dequeue | 从头出发，每次最远可以前进k，求到终点经过的index对应的值之和的最大值 |
|1711| [Count Good Meals](https://github.com/ZhengjunHUO/leetcode/tree/main/1711_Count_Good_Meals) | Array  HashTable  TwoPointers | 数列中和为2的次方的二元组的数量 |
|1749| [Maximum Absolute Sum of Any Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/1749_Maximum_Absolute_Sum_of_Any_Subarray) | Greedy | 和的绝对值最大的子串 |
|1871| [Jump Game VII](https://github.com/ZhengjunHUO/leetcode/tree/main/1871_Jump_Game_VII) | BreadthFirstSearch  Greedy  LineSweep | 从一个01字符串的头部出发，每次可以前进的距离介于[i,j]之间，且只能落在值为0的index上，求是否可以到达终点 |
