# LeetCode in Go

## LeetCode Algorithm/Solution in Golang

1. One folder One problem. Each folder contains a txt file which describe the problem. 
- The solution easy to figure out (for me) is in naive*.go, the name may contain some info about method used. For the rest *.go (eg. solution.go) is the solution i learned after research (especially thank to all who contributed to the “Discuss” section for those brillant ideas !). 
- For each problem, the final solution should reach the best possible time & space complexity, before move on to the next problem. 
- Commented in Chinese (we can explain the same idea with less words).
	
2. About Go, normally I only import fmt to print out the result to check quickly if the algorithm works as expected. For some problems which engage data types, i use my [homemade thread-safe go data type](https://github.com/ZhengjunHUO/godtype)

3. I didn’t specifically note down the difficulty of the problems, because it’s quite subjective and depends on your experience on some type of problem: you can easily solve some “hard” problems if you’ve figured out some sort of “template”  of this type in 10 minutes, but cat get stuck for hours by some “easy” problems you are not quite familiar or you've never seen before (like 初見殺し boss in game :) ?)

| # | Title | Tags |
|---| ----- | -------- |
|0001| [Two Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0001_Two_Sum) | Array  HashTable |
|0002| [Add Two Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0002_Add_Two_Numbers) | LinkedList  Math  Recursion |
|0003| [Longest Substring Without Repeating Characters](https://github.com/ZhengjunHUO/leetcode/tree/main/0003_Longest_Substring_Without_Repeating_Characters) | HashTable  SlidingWindow  String  TwoPointers |
|0011| [Container With Most Water](https://github.com/ZhengjunHUO/leetcode/tree/main/0011_Container_With_Most_Water) | Array  TwoPointers |
|0015| [Three Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0015_Three_Sum) | Array  TwoPointers |
|0016| [3Sum Closest](https://github.com/ZhengjunHUO/leetcode/tree/main/0016_3Sum_Closest) | Array  TwoPointers |
|0018| [Four Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0018_Four_Sum) | Array  HashTable  TwoPointers |
|0030| [Substring with Concatenation of All Words](https://github.com/ZhengjunHUO/leetcode/tree/main/0030_Substring_with_Concatenation_of_All_Words) | HashTable  String  TwoPointers |
|0042| [Trapping Rain Water](https://github.com/ZhengjunHUO/leetcode/tree/main/0042_Trapping_Rain_Water) | Array  Stack  TwoPointers  DynamicProgramming |
|0043| [Multiply Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0043_Multiply_Strings) | Math  String |
|0045| [Jump Game II](https://github.com/ZhengjunHUO/leetcode/tree/main/0045_Jump_Game_II) | Array  Greedy |
|0049| [Group Anagrams](https://github.com/ZhengjunHUO/leetcode/tree/main/0049_Group_Anagrams) | HashTable  String |
|0053| [Maximum Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0053_Maximum_Subarray) | Array  DivideAndConquer  DynamicProgramming |
|0055| [Jump Game](https://github.com/ZhengjunHUO/leetcode/tree/main/0055_Jump_Game) | Array  Greedy |
|0066| [Plus One](https://github.com/ZhengjunHUO/leetcode/tree/main/0066_Plus_One) | Array |
|0067| [Add Binary](https://github.com/ZhengjunHUO/leetcode/tree/main/0067_Add_Binary) | Math  String |
|0076| [Minimum Window Substring](https://github.com/ZhengjunHUO/leetcode/tree/main/0076_Minimum_Window_Substring) | HashTable  SlidingWindow  String  TwoPointers |
|0121| [Best Time to Buy and Sell Stock](https://github.com/ZhengjunHUO/leetcode/tree/main/0121_Best_Time_to_Buy_and_Sell_Stock) | Array  DynamicProgramming |
|0122| [Best Time to Buy and Sell Stock II](https://github.com/ZhengjunHUO/leetcode/tree/main/0122_Best_Time_to_Buy_and_Sell_Stock_II) | Array  Greedy |
|0123| [Best Time to Buy and Sell Stock III](https://github.com/ZhengjunHUO/leetcode/tree/main/0123_Best_Time_to_Buy_and_Sell_Stock_III) | Array  DynamicProgramming |
|0152| [Maximum Product Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0152_Maximum_Product_Subarray) | Array  DynamicProgramming |
|0155| [Min Stack](https://github.com/ZhengjunHUO/leetcode/tree/main/0155_Min_Stack) | Design  Stack |
|0167| [Two Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0167_Two_Sum_II) | Array  BinarySearch  TwoPointers |
|0170| [Two Sum III Data Structure Design](https://github.com/ZhengjunHUO/leetcode/tree/main/0170_Two_Sum_III_Data_Structure_Design) | Design  HastTable |
|0188| [Best Time to Buy and Sell Stock IV](https://github.com/ZhengjunHUO/leetcode/tree/main/0188_Best_Time_to_Buy_and_Sell_Stock_IV) | DynamicProgramming |
|0198| [House Robber](https://github.com/ZhengjunHUO/leetcode/tree/main/0198_House_Robber) | DynamicProgramming |
|0209| [Minimum Size Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0209_Minimum_Size_Subarray_Sum) | Array  BinarySearch  TwoPointers |
|0213| [House Robber II](https://github.com/ZhengjunHUO/leetcode/tree/main/0213_House_Robber_II) | DynamicProgramming |
|0238| [Product of Array Except Self](https://github.com/ZhengjunHUO/leetcode/tree/main/0238_Product_of_Array_Except_Self) | Array |
|0239| [Sliding Window Maximum](https://github.com/ZhengjunHUO/leetcode/tree/main/0239_Sliding_Window_Maximum) | Dequeue  Heap  SlidingWindow |
|0242| [Valid Anagram](https://github.com/ZhengjunHUO/leetcode/tree/main/0242_Valid_Anagram) | HastTable  Sort |
|0259| [3Sum Smaller](https://github.com/ZhengjunHUO/leetcode/tree/main/0259_3Sum_Smaller) | Array  TwoPointers |
|0283| [Move Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0283_Move_Zeroes) | Array  TwoPointers |
|0309| [Best Time to Buy and Sell Stock with Cooldown](https://github.com/ZhengjunHUO/leetcode/tree/main/0309_Best_Time_to_Buy_and_Sell_Stock_with_Cooldown) | DynamicProgramming |
|0325| [Maximum Size Subarray Sum Equals k](https://github.com/ZhengjunHUO/leetcode/tree/main/0325_Maximum_Size_Subarray_Sum_Equals_k) | ??? |
|0337| [House Robber III](https://github.com/ZhengjunHUO/leetcode/tree/main/0337_House_Robber_III) | DepthFirstSearch  DynamicProgramming  Tree |
|0371| [Sum of Two Integers](https://github.com/ZhengjunHUO/leetcode/tree/main/0371_Sum_of_Two_Integers) | BitManipulation |
|0407| [Trapping Rain Water II](https://github.com/ZhengjunHUO/leetcode/tree/main/0407_Trapping_Rain_Water_II) | BreadthFirstSearch  Heap |
|0415| [Add Strings](https://github.com/ZhengjunHUO/leetcode/tree/main/0415_Add_Strings) | String |
|0438| [Find All Anagrams in a String](https://github.com/ZhengjunHUO/leetcode/tree/main/0438_Find_All_Anagrams_in_a_String) | HastTable |
|0445| [Add Two Numbers II](https://github.com/ZhengjunHUO/leetcode/tree/main/0445_Add_Two_Numbers_II) | LinkedList |
|0454| [4Sum II](https://github.com/ZhengjunHUO/leetcode/tree/main/0454_4Sum_II) | BinarySearch  HastTable |
|0474| [Ones and Zeroes](https://github.com/ZhengjunHUO/leetcode/tree/main/0474_Ones_and_Zeroes) | DynamicProgramming |
|0523| [Continuous Subarray Sum](https://github.com/ZhengjunHUO/leetcode/tree/main/0523_Continuous_Subarray_Sum) | DynamicProgramming  Math |
|0560| [Subarray Sum Equals K](https://github.com/ZhengjunHUO/leetcode/tree/main/0560_Subarray_Sum_Equals_K) | Array  HashTable |
|0567| [Permutation in String](https://github.com/ZhengjunHUO/leetcode/tree/main/0567_Permutation_in_String) | TwoPointer  SlidingWindow |
|0600| [Non-negative Integers without Consecutive Ones](https://github.com/ZhengjunHUO/leetcode/tree/main/0600_Non-negative_Integers_without_Consecutive_Ones) | DynamicProgramming |
|0628| [Maximum Product of Three Numbers](https://github.com/ZhengjunHUO/leetcode/tree/main/0628_Maximum_Product_of_Three_Numbers) | Array  Math |
|0632| [Smallest Range Covering Elements from K Lists](https://github.com/ZhengjunHUO/leetcode/tree/main/0632_Smallest_Range_Covering_Elements_from_K_Lists) | HashTable  String  TwoPointers |
|0653| [Two Sum IV Input Is A BST](https://github.com/ZhengjunHUO/leetcode/tree/main/0653_Two_Sum_IV_Input_Is_A_BST) | Tree |
|0689| [Maximum Sum of 3 Non-Overlapping Subarrays](https://github.com/ZhengjunHUO/leetcode/tree/main/0689_Maximum_Sum_of_3_Non-Overlapping_Subarrays) | Array  DynamicProgramming |
|0697| [Degree of an Array](https://github.com/ZhengjunHUO/leetcode/tree/main/0697_Degree_of_an_Array) | Array |
|0713| [Subarray Product Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/0713_Subarray_Product_Less_Than_K) | Array  TwoPointers |
|0714| [Best Time to Buy and Sell Stock with Transaction Fee](https://github.com/ZhengjunHUO/leetcode/tree/main/0714_Best_Time_to_Buy_and_Sell_Stock_with_Transaction_Fee) | Array  DynamicProgramming  Greedy |
|0718| [Maximum Length of Repeated Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0718_Maximum_Length_of_Repeated_Subarray) | Array  BinarySearch  DynamicProgramming  HashTable |
|0724| [Find Pivot Index](https://github.com/ZhengjunHUO/leetcode/tree/main/0724_Find_Pivot_Index) | Array |
|0727| [Minimum Window Subsequence](https://github.com/ZhengjunHUO/leetcode/tree/main/0727_Minimum_Window_Subsequence) | DynamicProgramming  HashTable  SlidingWindow  String  TwoPointers |
|0740| [Delete and Earn](https://github.com/ZhengjunHUO/leetcode/tree/main/0740_Delete_and_Earn) | DynamicProgramming |
|0974| [Subarray Sums Divisible by K](https://github.com/ZhengjunHUO/leetcode/tree/main/0974_Subarray_Sums_Divisible_by_K) | Array  HashTable |
|0978| [Longest Turbulent Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/0978_Longest_Turbulent_Subarray) | Array  DynamicProgramming  SlidingWindow |
|0989| [Add to Array-Form of Integer](https://github.com/ZhengjunHUO/leetcode/tree/main/0989_Add_to_Array-Form_of_Integer) | Array |
|1099| [Two Sum Less Than K](https://github.com/ZhengjunHUO/leetcode/tree/main/1099_Two_Sum_Less_Than_K) | Array |
|1306| [Jump Game III](https://github.com/ZhengjunHUO/leetcode/tree/main/1306_Jump_Game_III) | BreadthFirstSearch  DepthFirstSearch  Recursion |
|1340| [Jump Game V](https://github.com/ZhengjunHUO/leetcode/tree/main/1340_Jump_Game_V) | DynamicProgramming |
|1345| [Jump Game IV](https://github.com/ZhengjunHUO/leetcode/tree/main/1345_Jump_Game_IV) | BreadthFirstSearch |
|1590| [Make Sum Divisible by P](https://github.com/ZhengjunHUO/leetcode/tree/main/1590_Make_Sum_Divisible_by_P) | Array  BinarySearch  HashTable  Math |
|1658| [Minimum Operations to Reduce X to Zero](https://github.com/ZhengjunHUO/leetcode/tree/main/1658_Minimum_Operations_to_Reduce_X_to_Zero) | BinarySearch  Greedy  SlidingWindow  TwoPointers |
|1679| [Max Number of K-Sum Pairs](https://github.com/ZhengjunHUO/leetcode/tree/main/1679_Max_Number_of_K-Sum_Pairs) | HashTable |
|1696| [Jump Game VI](https://github.com/ZhengjunHUO/leetcode/tree/main/1696_Jump_Game_VI) | Dequeue |
|1711| [Count Good Meals](https://github.com/ZhengjunHUO/leetcode/tree/main/1711_Count_Good_Meals) | Array  HashTable  TwoPointers |
|1749| [Maximum Absolute Sum of Any Subarray](https://github.com/ZhengjunHUO/leetcode/tree/main/1749_Maximum_Absolute_Sum_of_Any_Subarray) | Greedy |
|1871| [Jump Game VII](https://github.com/ZhengjunHUO/leetcode/tree/main/1871_Jump_Game_VII) | BreadthFirstSearch  Greedy  LineSweep |
