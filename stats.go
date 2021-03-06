package main

// reference form https://leetcode.com/discuss/general-discussion/522705/1000-LeetCode-Problems-within-a-Year
// data: 2020/3/8
// ,,45.00%,38.20%,30.80%,,,,30.80%,
// ,Title,Acc.,Diff.,Norm. Acc.,Avg. Acc.,Rel. Acc.,Points,Score,Premium

const stats =
`771,Jewels and Stones,84.60%,Easy,73.29%,45.10%,-51.35%,1,0.486,,
1119,Remove Vowels from a String,88.90%,Easy,72.49%,45.10%,-49.89%,1,0.501,,1
1350,Students With Invalid Departments,91.40%,Easy,71.60%,45.10%,-48.27%,1,0.517,Sql,1
511,Game Play Analysis I,79.00%,Easy,71.51%,45.10%,-48.10%,1,0.519,Sql,1
1068,Product Sales Analysis I,86.50%,Easy,70.84%,45.10%,-46.88%,1,0.531,Sql,1
1108,Defanging an IP Address,86.10%,Easy,69.85%,45.10%,-45.08%,1,0.549,,
1069,Product Sales Analysis II,85.40%,Easy,69.72%,45.10%,-44.85%,1,0.552,Sql,1
760,Find Anagram Mappings,80.50%,Easy,69.35%,45.10%,-44.18%,1,0.558,,1
709,To Lower Case,78.50%,Easy,68.10%,45.10%,-41.90%,1,0.581,,
595,Big Countries,76.20%,Easy,67.47%,45.10%,-40.75%,1,0.592,Sql,
1165,Single-Row Keyboard,84.50%,Easy,67.41%,45.10%,-40.64%,1,0.594,,1
1365,How Many Numbers Are Smaller Than the Current Number,87.30%,Easy,67.28%,45.10%,-40.40%,1,0.596,,
1342,Number of Steps to Reduce a Number to Zero,86.80%,Easy,67.12%,45.10%,-40.10%,1,0.599,,
1303,Find the Team Size,86.10%,Easy,66.99%,45.10%,-39.87%,1,0.601,Sql,1
613,Shortest Distance in a Line,75.80%,Easy,66.81%,45.10%,-39.54%,1,0.605,Sql,1
339,Nested List Weight Sum,71.70%,Easy,66.73%,45.10%,-39.40%,1,0.606,,1
1281,Subtract the Product and Sum of Digits of an Integer,84.60%,Easy,65.81%,45.10%,-37.73%,1,0.623,,
938,Range Sum of BST,79.30%,Easy,65.54%,45.10%,-37.24%,1,0.628,,
1295,Find Numbers with Even Number of Digits,84.50%,Easy,65.51%,45.10%,-37.17%,1,0.628,,
1313,Decompress Run-Length Encoded List,84.80%,Easy,65.54%,45.10%,-37.24%,1,0.628,,
461,Hamming Distance,71.10%,Easy,64.34%,45.10%,-35.04%,1,0.650,,
804,Unique Morse Code Words,76.00%,Easy,64.21%,45.10%,-34.81%,1,0.652,,
346,Moving Average from Data Stream,69.10%,Easy,64.03%,45.10%,-34.47%,1,0.655,,1
627,Swap Salary,73.10%,Easy,63.90%,45.10%,-34.25%,1,0.657,Sql,
617,Merge Two Binary Trees,72.70%,Easy,63.65%,45.10%,-33.79%,1,0.662,,
1221,Split a String in Balanced Strings,81.50%,Easy,63.59%,45.10%,-33.68%,1,0.663,,
359,Logger Rate Limiter,68.70%,Easy,63.43%,45.10%,-33.40%,1,0.666,,1
657,Robot Return to Origin,72.90%,Easy,63.26%,45.10%,-33.09%,1,0.669,,
561,Array Partition I,71.00%,Easy,62.77%,45.10%,-32.19%,1,0.678,,
104,Maximum Depth of Binary Tree,64.10%,Easy,62.57%,45.10%,-31.83%,1,0.682,,
832,Flipping an Image,74.70%,Easy,62.50%,45.10%,-31.69%,1,0.683,,
728,Self Dividing Numbers,73.00%,Easy,62.32%,45.10%,-31.37%,1,0.686,,
586,Customer Placing the Largest Number of Orders,70.40%,Easy,61.81%,45.10%,-30.43%,1,0.696,Sql,1
590,N-ary Tree Postorder Traversal,70.40%,Easy,61.75%,45.10%,-30.32%,1,0.697,,
589,N-ary Tree Preorder Traversal,70.30%,Easy,61.66%,45.10%,-30.17%,1,0.698,,
584,Find Customer Referee,70.20%,Easy,61.63%,45.10%,-30.12%,1,0.699,Sql,1
1021,Remove Outermost Parentheses,76.60%,Easy,61.63%,45.10%,-30.10%,1,0.699,,
1179,Reformat Department Table,78.80%,Easy,61.51%,45.10%,-29.89%,1,0.701,Sql,
1086,High Five,77.40%,Easy,61.47%,45.10%,-29.82%,1,0.702,,1
1134,Armstrong Number,78.10%,Easy,61.47%,45.10%,-29.81%,1,0.702,,1
1173,Immediate Food Delivery I,78.40%,Easy,61.20%,45.10%,-29.32%,1,0.707,Sql,1
136,Single Number,63.00%,Easy,61.01%,45.10%,-28.97%,1,0.710,,
1290,Convert Binary Number in a Linked List to Integer,79.90%,Easy,60.98%,45.10%,-28.93%,1,0.711,,
1266,Minimum Time Visiting All Points,79.30%,Easy,60.73%,45.10%,-28.47%,1,0.715,,
344,Reverse String,65.70%,Easy,60.65%,45.10%,-28.33%,1,0.717,,
905,Sort Array By Parity,73.70%,Easy,60.43%,45.10%,-27.92%,1,0.721,,
1251,Average Selling Price,78.70%,Easy,60.35%,45.10%,-27.78%,1,0.722,Sql,1
700,Search in a Binary Search Tree,70.40%,Easy,60.13%,45.10%,-27.38%,1,0.726,,
1213,Intersection of Three Sorted Arrays,77.70%,Easy,59.91%,45.10%,-26.98%,1,0.730,,1
1351,Count Negative Numbers in a Sorted Matrix,79.50%,Easy,59.69%,45.10%,-26.57%,1,0.734,,
557,Reverse Words in a String III,67.70%,Easy,59.53%,45.10%,-26.29%,1,0.737,,
1252,Cells with Odd Values in a Matrix,77.80%,Easy,59.44%,45.10%,-26.12%,1,0.739,,
509,Fibonacci Number,66.80%,Easy,59.33%,45.10%,-25.93%,1,0.741,,
1323,Maximum 69 Number,78.70%,Easy,59.30%,45.10%,-25.86%,1,0.741,,
559,Maximum Depth of N-ary Tree,67.40%,Easy,59.20%,45.10%,-25.69%,1,0.743,,
1180,Count Substrings with Only One Distinct Letter,76.50%,Easy,59.19%,45.10%,-25.67%,1,0.743,,1
961,N-Repeated Element in Size 2N Array,73.00%,Easy,58.91%,45.10%,-25.15%,1,0.749,,
1279,Traffic Light Controlled Intersection,77.40%,Easy,58.64%,45.10%,-24.67%,1,0.753,Concurrency,1
1085,Sum of Digits in the Minimum Number,74.50%,Easy,58.59%,45.10%,-24.57%,1,0.754,,1
1299,Replace Elements with Greatest Element on Right Side,77.60%,Easy,58.55%,45.10%,-24.50%,1,0.755,,
226,Invert Binary Tree,61.80%,Easy,58.49%,45.10%,-24.38%,1,0.756,,
852,Peak Index in a Mountain Array,70.90%,Easy,58.40%,45.10%,-24.23%,1,0.758,,
1148,Article Views I,75.00%,Easy,58.16%,45.10%,-23.79%,1,0.762,Sql,1
977,Squares of a Sorted Array,72.20%,Easy,57.87%,45.10%,-23.26%,1,0.767,,
1265,Print Immutable Linked List in Reverse,94.80%,Medium,76.25%,38.20%,-61.56%,2,0.769,,1
1304,Find N Unique Integers Sum up to Zero,76.60%,Easy,57.47%,45.10%,-22.54%,1,0.775,,
1309,Decrypt String from Alphabet to Integer Mapping,76.60%,Easy,57.40%,45.10%,-22.41%,1,0.776,,
266,Palindrome Permutation,61.20%,Easy,57.30%,45.10%,-22.22%,1,0.778,,1
942,DI String Match,71.10%,Easy,57.28%,45.10%,-22.19%,1,0.778,,
182,Duplicate Emails,59.50%,Easy,56.83%,45.10%,-21.37%,1,0.786,Sql,
206,Reverse Linked List,59.80%,Easy,56.78%,45.10%,-21.27%,1,0.787,,
933,Number of Recent Calls,70.50%,Easy,56.82%,45.10%,-21.34%,1,0.787,,
1050,Actors and Directors Who Cooperated At Least Three Times,71.90%,Easy,56.50%,45.10%,-20.77%,1,0.792,Sql,1
500,Keyboard Row,63.80%,Easy,56.47%,45.10%,-20.70%,1,0.793,,
1082,Sales Analysis I,72.30%,Easy,56.43%,45.10%,-20.64%,1,0.794,Sql,1
243,Shortest Word Distance,59.90%,Easy,56.34%,45.10%,-20.47%,1,0.795,,1
577,Employee Bonus,64.80%,Easy,56.34%,45.10%,-20.47%,1,0.795,Sql,1
476,Number Complement,63.20%,Easy,56.22%,45.10%,-20.25%,1,0.797,,
620,Not Boring Movies,65.20%,Easy,56.11%,45.10%,-20.05%,1,0.800,Sql,
811,Subdomain Visit Count,68.00%,Easy,56.11%,45.10%,-20.05%,1,0.800,,
944,Delete Columns to Make Sorted,69.90%,Easy,56.05%,45.10%,-19.95%,1,0.800,,
463,Island Perimeter,62.80%,Easy,56.01%,45.10%,-19.87%,1,0.801,,
293,Flip Game,60.00%,Easy,55.70%,45.10%,-19.31%,1,0.807,,1
175,Combine Two Tables,57.70%,Easy,55.13%,45.10%,-18.28%,1,0.817,Sql,
412,Fizz Buzz,61.20%,Easy,55.16%,45.10%,-18.32%,1,0.817,,
610,Triangle Judgement,64.00%,Easy,55.05%,45.10%,-18.13%,1,0.819,Sql,1
496,Next Greater Element I,62.20%,Easy,54.93%,45.10%,-17.90%,1,0.821,,
237,Delete Node in a Linked List,58.30%,Easy,54.82%,45.10%,-17.71%,1,0.823,,
922,Sort Array By Parity II,68.00%,Easy,54.48%,45.10%,-17.08%,1,0.829,,
13,Roman to Integer,54.60%,Easy,54.41%,45.10%,-16.96%,1,0.830,,
349,Intersection of Two Arrays,59.50%,Easy,54.38%,45.10%,-16.91%,1,0.831,,
897,Increasing Order Search Tree,67.40%,Easy,54.24%,45.10%,-16.66%,1,0.833,,
929,Unique Email Addresses,67.60%,Easy,53.97%,45.10%,-16.17%,1,0.838,,
883,Projection Area of 3D Shapes,66.90%,Easy,53.95%,45.10%,-16.12%,1,0.839,,
108,Convert Sorted Array to Binary Search Tree,55.20%,Easy,53.62%,45.10%,-15.51%,1,0.845,,
821,Shortest Distance to a Character,65.60%,Easy,53.56%,45.10%,-15.41%,1,0.846,,
1207,Unique Number of Occurrences,71.20%,Easy,53.50%,45.10%,-15.30%,1,0.847,,
169,Majority Element,55.90%,Easy,53.42%,45.10%,-15.16%,1,0.848,,
876,Middle of the Linked List,66.20%,Easy,53.35%,45.10%,-15.03%,1,0.850,,
965,Univalued Binary Tree,67.30%,Easy,53.15%,45.10%,-14.66%,1,0.853,,
1064,Fixed Point,68.70%,Easy,53.09%,45.10%,-14.56%,1,0.854,,1
1051,Height Checker,68.50%,Easy,53.09%,45.10%,-14.55%,1,0.855,,
603,Consecutive Available Seats,61.80%,Easy,52.96%,45.10%,-14.31%,1,0.857,Sql,1
122,Best Time to Buy and Sell Stock II,54.40%,Easy,52.61%,45.10%,-13.68%,1,0.863,,
1327,List the Products Ordered in a Period,72.10%,Easy,52.64%,45.10%,-13.73%,1,0.863,Sql,1
682,Baseball Game,62.50%,Easy,52.50%,45.10%,-13.47%,1,0.865,,
575,Distribute Candies,60.80%,Easy,52.37%,45.10%,-13.24%,1,0.868,,
766,Toeplitz Matrix,63.60%,Easy,52.37%,45.10%,-13.23%,1,0.868,,
1241,Number of Comments per Post,70.50%,Easy,52.30%,45.10%,-13.11%,1,0.869,Sql,1
806,Number of Lines To Write String,64.00%,Easy,52.18%,45.10%,-12.89%,1,0.871,,
283,Move Zeroes,56.30%,Easy,52.15%,45.10%,-12.84%,1,0.872,,
669,Trim a Binary Search Tree,61.90%,Easy,52.09%,45.10%,-12.73%,1,0.873,,
607,Sales Person,60.90%,Easy,52.00%,45.10%,-12.56%,1,0.874,Sql,1
872,Leaf-Similar Trees,64.80%,Easy,52.01%,45.10%,-12.59%,1,0.874,,
1002,Find Common Characters,66.70%,Easy,52.00%,45.10%,-12.58%,1,0.874,,
637,Average of Levels in Binary Tree,61.30%,Easy,51.96%,45.10%,-12.49%,1,0.875,,
893,Groups of Special-Equivalent Strings,65.00%,Easy,51.90%,45.10%,-12.39%,1,0.876,,
908,Smallest Range I,65.20%,Easy,51.88%,45.10%,-12.35%,1,0.876,,
258,Add Digits,55.60%,Easy,51.82%,45.10%,-12.23%,1,0.878,,
217,Contains Duplicate,54.90%,Easy,51.72%,45.10%,-12.05%,1,0.879,,
292,Nim Game,56.00%,Easy,51.72%,45.10%,-12.05%,1,0.879,,
242,Valid Anagram,55.20%,Easy,51.65%,45.10%,-11.93%,1,0.881,,
566,Reshape the Matrix,59.90%,Easy,51.60%,45.10%,-11.84%,1,0.882,,
999,Available Captures for Rook,66.00%,Easy,51.35%,45.10%,-11.38%,1,0.886,,
1047,Remove All Adjacent Duplicates In String,66.50%,Easy,51.14%,45.10%,-11.01%,1,0.890,,
21,Merge Two Sorted Lists,51.40%,Easy,51.09%,45.10%,-10.91%,1,0.891,,
181,Employees Earning More Than Their Managers,53.70%,Easy,51.05%,45.10%,-10.83%,1,0.892,Sql,
751,IP to CIDR,61.90%,Easy,50.89%,45.10%,-10.54%,1,0.895,,1
1122,Relative Sort Array,67.30%,Easy,50.84%,45.10%,-10.46%,1,0.895,,
171,Excel Sheet Column Number,53.30%,Easy,50.79%,45.10%,-10.37%,1,0.896,,
1280,Students and Examinations,69.50%,Easy,50.73%,45.10%,-10.25%,1,0.898,Sql,1
867,Transpose Matrix,63.30%,Easy,50.58%,45.10%,-9.99%,1,0.900,,
762,Prime Number of Set Bits in Binary Representation,61.70%,Easy,50.52%,45.10%,-9.88%,1,0.901,,
1025,Divisor Game,65.50%,Easy,50.47%,45.10%,-9.78%,1,0.902,,
100,Same Tree,51.80%,Easy,50.33%,45.10%,-9.53%,1,0.905,,
784,Letter Case Permutation,61.80%,Easy,50.30%,45.10%,-9.47%,1,0.905,,
1211,Queries Quality and Percentage,68.10%,Easy,50.34%,45.10%,-9.54%,1,0.905,Sql,1
167,Two Sum II - Input array is sorted,52.70%,Easy,50.25%,45.10%,-9.38%,1,0.906,,
1133,Largest Unique Number,66.90%,Easy,50.28%,45.10%,-9.44%,1,0.906,,1
1196,How Many Apples Can You Put into the Basket,67.80%,Easy,50.26%,45.10%,-9.40%,1,0.906,,1
252,Meeting Rooms,53.80%,Easy,50.10%,45.10%,-9.11%,1,0.909,,1
1160,Find Words That Can Be Formed by Characters,67.10%,Easy,50.09%,45.10%,-9.08%,1,0.909,,
1030,Matrix Cells in Distance Order,64.90%,Easy,49.79%,45.10%,-8.55%,1,0.915,,
1237,Find Positive Integer Solution for a Given Equation,67.90%,Easy,49.76%,45.10%,-8.48%,1,0.915,,
1022,Sum of Root To Leaf Binary Numbers,64.70%,Easy,49.71%,45.10%,-8.40%,1,0.916,,
1356,Sort Integers by The Number of 1 Bits,69.60%,Easy,49.71%,45.10%,-8.40%,1,0.916,,
521,Longest Uncommon Subsequence I,57.10%,Easy,49.46%,45.10%,-7.94%,1,0.921,,
800,Similar RGB Color,61.10%,Easy,49.37%,45.10%,-7.77%,1,0.922,,1
1337,The K Weakest Rows in a Matrix,68.80%,Easy,49.19%,45.10%,-7.45%,1,0.925,,
824,Goat Latin,61.20%,Easy,49.11%,45.10%,-7.31%,1,0.927,,
485,Max Consecutive Ones,56.10%,Easy,48.99%,45.10%,-7.08%,1,0.929,,
884,Uncommon Words from Two Sentences,61.90%,Easy,48.93%,45.10%,-6.98%,1,0.930,,
1078,Occurrences After Bigram,64.60%,Easy,48.79%,45.10%,-6.72%,1,0.933,,
693,Binary Number with Alternating Bits,58.90%,Easy,48.74%,45.10%,-6.62%,1,0.934,,
706,Design HashMap,59.00%,Easy,48.65%,45.10%,-6.46%,1,0.935,,
389,Find the Difference,54.30%,Easy,48.59%,45.10%,-6.37%,1,0.936,,
807,Max Increase to Keep City Skyline,82.90%,Medium,71.06%,38.20%,-53.18%,2,0.936,,
1200,Minimum Absolute Difference,66.20%,Easy,48.60%,45.10%,-6.38%,1,0.936,,
448,Find All Numbers Disappeared in an Array,55.10%,Easy,48.53%,45.10%,-6.25%,1,0.938,,
107,Binary Tree Level Order Traversal II,50.00%,Easy,48.43%,45.10%,-6.07%,1,0.939,,
118,Pascal's Triangle,50.20%,Easy,48.47%,45.10%,-6.14%,1,0.939,,
535,Encode and Decode TinyURL,78.70%,Medium,70.85%,38.20%,-52.84%,2,0.943,,
1075,Project Employees I,64.00%,Easy,48.23%,45.10%,-5.71%,1,0.943,Sql,1
705,Design HashSet,58.40%,Easy,48.06%,45.10%,-5.39%,1,0.946,,
512,Game Play Analysis II,55.20%,Easy,47.69%,45.10%,-4.72%,1,0.953,Sql,1
121,Best Time to Buy and Sell Stock,49.40%,Easy,47.63%,45.10%,-4.60%,1,0.954,,
985,Sum of Even Numbers After Queries,62.00%,Easy,47.55%,45.10%,-4.47%,1,0.955,,
256,Paint House,51.20%,Easy,47.45%,45.10%,-4.27%,1,0.957,,1
868,Binary Gap,60.20%,Easy,47.47%,45.10%,-4.32%,1,0.957,,
183,Customers Who Never Order,50.10%,Easy,47.42%,45.10%,-4.22%,1,0.958,Sql,
1046,Last Stone Weight,62.50%,Easy,47.16%,45.10%,-3.75%,1,0.963,,
1114,Print in Order,63.40%,Easy,47.06%,45.10%,-3.57%,1,0.964,Concurrency,
1185,Day of the Week,64.30%,Easy,46.92%,45.10%,-3.32%,1,0.967,,
9,Palindrome Number,46.70%,Easy,46.57%,45.10%,-2.67%,1,0.973,,
27,Remove Element,46.90%,Easy,46.50%,45.10%,-2.56%,1,0.974,,
268,Missing Number,50.30%,Easy,46.37%,45.10%,-2.31%,1,0.977,,
1217,Play with Chips,64.00%,Easy,46.15%,45.10%,-1.91%,1,0.981,,
1113,Reported Posts,62.40%,Easy,46.08%,45.10%,-1.78%,1,0.982,Sql,1
383,Ransom Note,51.50%,Easy,45.88%,45.10%,-1.43%,1,0.986,,
387,First Unique Character in a String,51.50%,Easy,45.82%,45.10%,-1.32%,1,0.987,,
538,Convert BST to Greater Tree,53.60%,Easy,45.71%,45.10%,-1.11%,1,0.989,,
690,Employee Importance,55.80%,Easy,45.68%,45.10%,-1.06%,1,0.989,,
202,Happy Number,48.60%,Easy,45.64%,45.10%,-0.98%,1,0.990,,
812,Largest Triangle Area,57.50%,Easy,45.59%,45.10%,-0.89%,1,0.991,,
1294,Weather Type in Each Country,64.60%,Easy,45.62%,45.10%,-0.95%,1,0.991,Sql,1
520,Detect Capital,53.00%,Easy,45.37%,45.10%,-0.50%,1,0.995,,
257,Binary Tree Paths,49.10%,Easy,45.33%,45.10%,-0.42%,1,0.996,,
119,Pascal's Triangle II,47.00%,Easy,45.25%,45.10%,-0.28%,1,0.997,,
70,Climbing Stairs,46.20%,Easy,45.17%,45.10%,-0.13%,1,0.999,,
350,Intersection of Two Arrays II,50.30%,Easy,45.17%,45.10%,-0.12%,1,0.999,,
371,Sum of Two Integers,50.60%,Easy,45.16%,45.10%,-0.11%,1,0.999,,
1,Two Sum,45.10%,Easy,45.09%,45.10%,0.03%,1,1.000,,
53,Maximum Subarray,45.70%,Easy,44.92%,45.10%,0.32%,1,1.003,,
748,Shortest Completing Word,55.90%,Easy,44.93%,45.10%,0.31%,1,1.003,,
788,Rotated Digits,56.50%,Easy,44.94%,45.10%,0.29%,1,1.003,,
696,Count Binary Substrings,55.10%,Easy,44.89%,45.10%,0.38%,1,1.004,,
892,Surface Area of 3D Shapes,57.90%,Easy,44.82%,45.10%,0.51%,1,1.005,,
653,Two Sum IV - Input is a BST,54.30%,Easy,44.72%,45.10%,0.69%,1,1.007,,
1009,Complement of Base 10 Integer,59.50%,Easy,44.70%,45.10%,0.73%,1,1.007,,
235,Lowest Common Ancestor of a Binary Search Tree,48.10%,Easy,44.65%,45.10%,0.81%,1,1.008,,
447,Number of Boomerangs,51.20%,Easy,44.64%,45.10%,0.83%,1,1.008,,
530,Minimum Absolute Difference in BST,52.40%,Easy,44.63%,45.10%,0.86%,1,1.009,,
654,Maximum Binary Tree,78.30%,Medium,68.71%,38.20%,-49.37%,2,1.013,,
1099,Two Sum Less Than K,60.50%,Easy,44.38%,45.10%,1.31%,1,1.013,,1
404,Sum of Left Leaves,50.20%,Easy,44.27%,45.10%,1.50%,1,1.015,,
101,Symmetric Tree,45.70%,Easy,44.22%,45.10%,1.61%,1,1.016,,
191,Number of 1 Bits,47.00%,Easy,44.20%,45.10%,1.64%,1,1.016,,
606,Construct String from Binary Tree,53.10%,Easy,44.21%,45.10%,1.62%,1,1.016,,
888,Fair Candy Swap,57.20%,Easy,44.18%,45.10%,1.68%,1,1.017,,
701,Insert into a Binary Search Tree,78.80%,Medium,68.52%,38.20%,-49.06%,2,1.019,,
1103,Distribute Candies to People,60.10%,Easy,43.92%,45.10%,2.14%,1,1.021,,
917,Reverse Only Letters,57.20%,Easy,43.75%,45.10%,2.46%,1,1.025,,
232,Implement Queue using Stacks,47.10%,Easy,43.70%,45.10%,2.55%,1,1.026,,
896,Monotonic Array,56.80%,Easy,43.66%,45.10%,2.63%,1,1.026,,
1065,Index Pairs of a String,59.20%,Easy,43.58%,45.10%,2.77%,1,1.028,,1
409,Longest Palindrome,49.50%,Easy,43.50%,45.10%,2.91%,1,1.029,,
1189,Maximum Number of Balloons,60.90%,Easy,43.46%,45.10%,2.98%,1,1.030,,
1150,Check If a Number Is Majority Element in a Sorted Array,60.20%,Easy,43.33%,45.10%,3.22%,1,1.032,,1
453,Minimum Moves to Equal Array Elements,49.80%,Easy,43.16%,45.10%,3.54%,1,1.035,,
83,Remove Duplicates from Sorted List,44.30%,Easy,43.08%,45.10%,3.67%,1,1.037,,
26,Remove Duplicates from Sorted Array,43.40%,Easy,43.02%,45.10%,3.79%,1,1.038,,
976,Largest Perimeter Triangle,57.10%,Easy,42.79%,45.10%,4.22%,1,1.042,,
1089,Duplicate Zeros,58.70%,Easy,42.73%,45.10%,4.32%,1,1.043,,
697,Degree of an Array,52.90%,Easy,42.68%,45.10%,4.41%,1,1.044,,
38,Count and Say,43.10%,Easy,42.54%,45.10%,4.66%,1,1.047,,
455,Assign Cookies,49.20%,Easy,42.53%,45.10%,4.69%,1,1.047,,
733,Flood Fill,53.10%,Easy,42.35%,45.10%,5.01%,1,1.050,,
270,Closest Binary Search Tree Value,46.20%,Easy,42.24%,45.10%,5.21%,1,1.052,,1
1013,Partition Array Into Three Parts With Equal Sum,57.10%,Easy,42.24%,45.10%,5.20%,1,1.052,,
506,Relative Ranks,49.60%,Easy,42.18%,45.10%,5.32%,1,1.053,,
392,Is Subsequence,47.90%,Easy,42.15%,45.10%,5.37%,1,1.054,,
1260,Shift 2D Grid,60.60%,Easy,42.12%,45.10%,5.43%,1,1.054,,
492,Construct the Rectangle,49.20%,Easy,41.98%,45.10%,5.68%,1,1.057,,
67,Add Binary,42.30%,Easy,41.32%,45.10%,6.89%,1,1.069,,
1287,Element Appearing More Than 25% In Sorted Array,60.20%,Easy,41.32%,45.10%,6.88%,1,1.069,,
1170,Compare Strings by Frequency of the Smallest Character,58.30%,Easy,41.14%,45.10%,7.21%,1,1.072,,
599,Minimum Index Sum of Two Lists,49.80%,Easy,41.01%,45.10%,7.44%,1,1.074,,
953,Verifying an Alien Dictionary,55.00%,Easy,41.02%,45.10%,7.43%,1,1.074,,
35,Search Insert Position,41.50%,Easy,40.99%,45.10%,7.49%,1,1.075,,
110,Balanced Binary Tree,42.60%,Easy,40.99%,45.10%,7.49%,1,1.075,,
66,Plus One,41.90%,Easy,40.93%,45.10%,7.59%,1,1.076,,
661,Image Smoother,50.60%,Easy,40.91%,45.10%,7.64%,1,1.076,,
544,Output Contest Matches,74.60%,Medium,66.62%,38.20%,-45.99%,2,1.080,,1
401,Binary Watch,46.50%,Easy,40.62%,45.10%,8.16%,1,1.082,,
1118,Number of Days in a Month,57.00%,Easy,40.60%,45.10%,8.19%,1,1.082,,1
246,Strobogrammatic Number,44.10%,Easy,40.49%,45.10%,8.39%,1,1.084,,1
1029,Two City Scheduling,55.60%,Easy,40.51%,45.10%,8.36%,1,1.084,,
598,Range Addition II,49.20%,Easy,40.43%,45.10%,8.51%,1,1.085,,
704,Binary Search,50.60%,Easy,40.27%,45.10%,8.79%,1,1.088,,
543,Diameter of Binary Tree,48.20%,Easy,40.24%,45.10%,8.86%,1,1.089,,
1137,N-th Tribonacci Number,56.90%,Easy,40.22%,45.10%,8.88%,1,1.089,,
415,Add Strings,46.00%,Easy,39.91%,45.10%,9.45%,1,1.094,,
783,Minimum Distance Between BST Nodes,51.40%,Easy,39.92%,45.10%,9.44%,1,1.094,,
937,Reorder Data in Log Files,53.70%,Easy,39.96%,45.10%,9.37%,1,1.094,,
1331,Rank Transform of an Array,59.40%,Easy,39.88%,45.10%,9.51%,1,1.095,,
225,Implement Stack using Queues,42.90%,Easy,39.60%,45.10%,10.02%,1,1.100,,
1322,Ads Performance,59.00%,Easy,39.61%,45.10%,10.00%,1,1.100,Sql,1
541,Reverse String II,47.40%,Easy,39.47%,45.10%,10.26%,1,1.103,,
563,Binary Tree Tilt,47.70%,Easy,39.44%,45.10%,10.30%,1,1.103,,
231,Power of Two,42.80%,Easy,39.41%,45.10%,10.36%,1,1.104,,
1084,Sales Analysis III,55.30%,Easy,39.40%,45.10%,10.38%,1,1.104,Sql,1
155,Min Stack,41.00%,Easy,38.73%,45.10%,11.61%,1,1.116,,
198,House Robber,41.60%,Easy,38.70%,45.10%,11.66%,1,1.117,,
1270,All People Report to the Given Manager,84.10%,Medium,65.47%,38.20%,-44.13%,2,1.117,Sql,1
1317,Convert Integer to the Sum of Two No-Zero Integers,57.90%,Easy,38.58%,45.10%,11.87%,1,1.119,,
437,Path Sum III,44.90%,Easy,38.49%,45.10%,12.04%,1,1.120,,
860,Lemonade Change,51.10%,Easy,38.49%,45.10%,12.05%,1,1.120,,
534,Game Play Analysis III,73.20%,Medium,65.37%,38.20%,-43.96%,2,1.121,Sql,1
717,1-bit and 2-bit Characters,49.00%,Easy,38.48%,45.10%,12.05%,1,1.121,,
551,Student Attendance Record I,46.50%,Easy,38.42%,45.10%,12.17%,1,1.122,,
112,Path Sum,39.90%,Easy,38.26%,45.10%,12.46%,1,1.125,,
504,Base 7,45.60%,Easy,38.21%,45.10%,12.55%,1,1.126,,
746,Min Cost Climbing Stairs,49.10%,Easy,38.16%,45.10%,12.64%,1,1.126,,
345,Reverse Vowels of a String,43.10%,Easy,38.04%,45.10%,12.86%,1,1.129,,
703,Kth Largest Element in a Stream,48.20%,Easy,37.89%,45.10%,13.13%,1,1.131,,
1184,Distance Between Bus Stops,55.30%,Easy,37.93%,45.10%,13.05%,1,1.131,,
20,Valid Parentheses,38.10%,Easy,37.81%,45.10%,13.28%,1,1.133,,
796,Rotate String,49.50%,Easy,37.83%,45.10%,13.25%,1,1.133,,
1071,Greatest Common Divisor of Strings,53.40%,Easy,37.69%,45.10%,13.49%,1,1.135,,
141,Linked List Cycle,39.70%,Easy,37.63%,45.10%,13.60%,1,1.136,,
993,Cousins in Binary Tree,52.20%,Easy,37.64%,45.10%,13.60%,1,1.136,,
1076,Project Employees II,53.40%,Easy,37.62%,45.10%,13.63%,1,1.136,Sql,1
303,Range Sum Query - Immutable,42.00%,Easy,37.56%,45.10%,13.74%,1,1.137,,
628,Maximum Product of Three Numbers,46.80%,Easy,37.59%,45.10%,13.68%,1,1.137,,
263,Ugly Number,41.20%,Easy,37.34%,45.10%,14.13%,1,1.141,,
1282,Group the People Given the Group Size They Belong To,83.50%,Medium,64.70%,38.20%,-42.88%,2,1.142,,
405,Convert a Number to Hexadecimal,43.20%,Easy,37.26%,45.10%,14.28%,1,1.143,,
1141,User Activity for the Past 30 Days I,54.00%,Easy,37.27%,45.10%,14.27%,1,1.143,Sql,1
326,Power of Three,41.90%,Easy,37.12%,45.10%,14.54%,1,1.145,,
299,Bulls and Cows,41.40%,Easy,37.01%,45.10%,14.73%,1,1.147,,
88,Merge Sorted Array,38.10%,Easy,36.81%,45.10%,15.10%,1,1.151,,
830,Positions of Large Groups,48.80%,Easy,36.63%,45.10%,15.43%,1,1.154,,
1332,Remove Palindromic Subsequences,56.20%,Easy,36.66%,45.10%,15.37%,1,1.154,,
594,Longest Harmonious Subsequence,45.30%,Easy,36.59%,45.10%,15.50%,1,1.155,,
1302,Deepest Leaves Sum,83.40%,Medium,64.30%,38.20%,-42.24%,2,1.155,,
1315,Sum of Nodes with Even-Valued Grandparent,83.60%,Medium,64.31%,38.20%,-42.25%,2,1.155,,
720,Longest Word in Dictionary,47.10%,Easy,36.54%,45.10%,15.59%,1,1.156,,
836,Rectangle Overlap,48.50%,Easy,36.24%,45.10%,16.14%,1,1.161,,
1038,Binary Search Tree to Greater Sum Tree,79.30%,Medium,64.08%,38.20%,-41.87%,2,1.163,,
342,Power of Four,41.10%,Easy,36.08%,45.10%,16.42%,1,1.164,,
1083,Sales Analysis II,52.00%,Easy,36.12%,45.10%,16.36%,1,1.164,Sql,1
374,Guess Number Higher or Lower,41.50%,Easy,36.01%,45.10%,16.55%,1,1.165,,
1005,Maximize Sum Of Array After K Negations,50.80%,Easy,36.06%,45.10%,16.47%,1,1.165,,
205,Isomorphic Strings,39.00%,Easy,35.99%,45.10%,16.59%,1,1.166,,
366,Find Leaves of Binary Tree,69.30%,Medium,63.93%,38.20%,-41.64%,2,1.167,,1
1308,Running Total for Different Genders,83.00%,Medium,63.82%,38.20%,-41.45%,2,1.171,Sql,1
160,Intersection of Two Linked Lists,37.90%,Easy,35.55%,45.10%,17.39%,1,1.174,,
482,License Key Formatting,42.60%,Easy,35.53%,45.10%,17.43%,1,1.174,,
367,Valid Perfect Square,40.90%,Easy,35.52%,45.10%,17.45%,1,1.175,,
674,Longest Continuous Increasing Subsequence,45.30%,Easy,35.41%,45.10%,17.64%,1,1.176,,
572,Subtree of Another Tree,43.70%,Easy,35.31%,45.10%,17.83%,1,1.178,,
1176,Diet Plan Performance,52.60%,Easy,35.35%,45.10%,17.76%,1,1.178,,1
172,Factorial Trailing Zeroes,37.70%,Easy,35.18%,45.10%,18.07%,1,1.181,,
997,Find the Town Judge,49.80%,Easy,35.18%,45.10%,18.07%,1,1.181,,
196,Delete Duplicate Emails,38.00%,Easy,35.13%,45.10%,18.17%,1,1.182,Sql,
1228,Missing Number In Arithmetic Progression,53.10%,Easy,35.09%,45.10%,18.23%,1,1.182,,1
111,Minimum Depth of Binary Tree,36.60%,Easy,34.97%,45.10%,18.45%,1,1.184,,
844,Backspace String Compare,47.20%,Easy,34.82%,45.10%,18.72%,1,1.187,,
1056,Confusing Number,50.30%,Easy,34.81%,45.10%,18.74%,1,1.187,,1
459,Repeated Substring Pattern,41.50%,Easy,34.77%,45.10%,18.82%,1,1.188,,
1275,Find Winner on a Tic Tac Toe Game,53.30%,Easy,34.60%,45.10%,19.13%,1,1.191,,
234,Palindrome Linked List,38.00%,Easy,34.57%,45.10%,19.18%,1,1.192,,
14,Longest Common Prefix,34.70%,Easy,34.49%,45.10%,19.32%,1,1.193,,
197,Rising Temperature,37.20%,Easy,34.31%,45.10%,19.65%,1,1.197,Sql,
1271,Hexspeak,52.80%,Easy,34.16%,45.10%,19.93%,1,1.199,,1
203,Remove Linked List Elements,37.10%,Easy,34.12%,45.10%,20.00%,1,1.200,,
1243,Array Transformation,52.30%,Easy,34.07%,45.10%,20.09%,1,1.201,,1
744,Find Smallest Letter Greater Than Target,44.90%,Easy,33.99%,45.10%,20.24%,1,1.202,,
501,Find Mode in Binary Search Tree,41.10%,Easy,33.75%,45.10%,20.67%,1,1.207,,
219,Contains Duplicate II,36.80%,Easy,33.59%,45.10%,20.97%,1,1.210,,
276,Paint Fence,37.60%,Easy,33.55%,45.10%,21.03%,1,1.210,,1
1175,Prime Arrangements,50.70%,Easy,33.47%,45.10%,21.19%,1,1.212,,
28,Implement strStr(),33.80%,Easy,33.39%,45.10%,21.33%,1,1.213,,
443,String Compression,39.80%,Easy,33.30%,45.10%,21.49%,1,1.215,,
758,Bold Words in String,44.30%,Easy,33.18%,45.10%,21.71%,1,1.217,,1
763,Partition Labels,73.60%,Medium,62.41%,38.20%,-39.17%,2,1.217,,
671,Second Minimum Node In a Binary Tree,43.00%,Easy,33.16%,45.10%,21.75%,1,1.218,,
1042,Flower Planting With No Adjacent,48.20%,Easy,32.92%,45.10%,22.19%,1,1.222,,
441,Arranging Coins,39.30%,Easy,32.83%,45.10%,22.35%,1,1.223,,
994,Rotting Oranges,47.20%,Easy,32.62%,45.10%,22.73%,1,1.227,,
125,Valid Palindrome,34.20%,Easy,32.37%,45.10%,23.19%,1,1.232,,
190,Reverse Bits,35.10%,Easy,32.31%,45.10%,23.29%,1,1.233,,
619,Biggest Single Number,41.40%,Easy,32.32%,45.10%,23.28%,1,1.233,Sql,1
724,Find Pivot Index,42.90%,Easy,32.28%,45.10%,23.35%,1,1.233,,
1010,Pairs of Songs With Total Durations Divisible by 60,47.10%,Easy,32.29%,45.10%,23.34%,1,1.233,,
419,Battleships in a Board,68.00%,Medium,61.85%,38.20%,-38.28%,2,1.234,,
338,Counting Bits,66.80%,Medium,61.84%,38.20%,-38.26%,2,1.235,,
1018,Binary Prefix Divisible By 5,47.00%,Easy,32.07%,45.10%,23.74%,1,1.237,,
1154,Day of the Year,49.00%,Easy,32.07%,45.10%,23.73%,1,1.237,,
290,Word Pattern,36.30%,Easy,32.05%,45.10%,23.78%,1,1.238,,
645,Set Mismatch,41.50%,Easy,32.04%,45.10%,23.79%,1,1.238,,
797,All Paths From Source to Target,73.40%,Medium,61.71%,38.20%,-38.04%,2,1.239,,
597,Friend Requests I: Overall Acceptance Rate,40.70%,Easy,31.94%,45.10%,23.96%,1,1.240,Sql,1
69,Sqrt(x),32.90%,Easy,31.89%,45.10%,24.07%,1,1.241,,
819,Most Common Word,43.70%,Easy,31.69%,45.10%,24.43%,1,1.244,,
58,Length of Last Word,32.40%,Easy,31.55%,45.10%,24.68%,1,1.247,,
643,Maximum Average Subarray I,40.90%,Easy,31.47%,45.10%,24.83%,1,1.248,,
716,Max Stack,42.00%,Easy,31.50%,45.10%,24.77%,1,1.248,,1
925,Long Pressed Name,44.90%,Easy,31.33%,45.10%,25.08%,1,1.251,,
434,Number of Segments in a String,37.50%,Easy,31.13%,45.10%,25.44%,1,1.254,,
814,Binary Tree Pruning,73.20%,Medium,61.26%,38.20%,-37.32%,2,1.254,,
422,Valid Word Square,37.10%,Easy,30.91%,45.10%,25.85%,1,1.258,,1
734,Sentence Similarity,41.60%,Easy,30.83%,45.10%,25.98%,1,1.260,,1
195,Tenth Line,33.50%,Easy,30.64%,45.10%,26.34%,1,1.263,Shell,
1128,Number of Equivalent Domino Pairs,47.00%,Easy,30.46%,45.10%,26.67%,1,1.267,,
170,Two Sum III - Data structure design,32.70%,Easy,30.21%,45.10%,27.13%,1,1.271,,1
189,Rotate Array,33.00%,Easy,30.23%,45.10%,27.09%,1,1.271,,
747,Largest Number At Least Twice of Others,41.10%,Easy,30.14%,45.10%,27.24%,1,1.272,,
157,Read N Characters Given Read4,32.20%,Easy,29.90%,45.10%,27.69%,1,1.277,,1
849,Maximize Distance to Closest Person,42.20%,Easy,29.75%,45.10%,27.96%,1,1.280,,
1285,Find the Start and End Number of Continuous Ranges,79.30%,Medium,60.45%,38.20%,-36.01%,2,1.280,Sql,1
989,Add to Array-Form of Integer,44.10%,Easy,29.59%,45.10%,28.24%,1,1.282,,
1008,Construct Binary Search Tree from Preorder Traversal,75.10%,Medium,60.32%,38.20%,-35.79%,2,1.284,,
624,Maximum Distance in Arrays,38.50%,Easy,29.35%,45.10%,28.69%,1,1.287,,1
278,First Bad Version,33.00%,Easy,28.92%,45.10%,29.47%,1,1.295,,
1232,Check If It Is a Straight Line,47.00%,Easy,28.93%,45.10%,29.45%,1,1.295,,
596,Classes More Than 5 Students,37.30%,Easy,28.56%,45.10%,30.13%,1,1.301,Sql,
894,All Possible Full Binary Trees,72.90%,Medium,59.79%,38.20%,-34.93%,2,1.301,,
46,Permutations,60.40%,Medium,59.73%,38.20%,-34.83%,2,1.303,,
890,Find and Replace Pattern,72.80%,Medium,59.75%,38.20%,-34.87%,2,1.303,,
1360,Number of Days Between Two Dates,48.20%,Easy,28.25%,45.10%,30.69%,1,1.307,,
22,Generate Parentheses,59.90%,Medium,59.58%,38.20%,-34.59%,2,1.308,,
94,Binary Tree Inorder Traversal,60.90%,Medium,59.52%,38.20%,-34.50%,2,1.310,,
507,Perfect Number,35.30%,Easy,27.86%,45.10%,31.40%,1,1.314,,
950,Reveal Cards In Increasing Order,73.30%,Medium,59.37%,38.20%,-34.25%,2,1.315,,
168,Excel Sheet Column Title,30.20%,Easy,27.74%,45.10%,31.63%,1,1.316,,
204,Count Primes,30.60%,Easy,27.61%,45.10%,31.86%,1,1.319,,
176,Second Highest Salary,30.10%,Easy,27.52%,45.10%,32.02%,1,1.320,Sql,
604,Design Compressed String Iterator,36.30%,Easy,27.44%,45.10%,32.17%,1,1.322,,1
537,Complex Number Multiplication,66.70%,Medium,58.82%,38.20%,-33.37%,2,1.333,,
861,Score After Flipping Matrix,71.40%,Medium,58.77%,38.20%,-33.29%,2,1.334,,
1079,Letter Tile Possibilities,74.60%,Medium,58.77%,38.20%,-33.29%,2,1.334,,
1329,Sort the Matrix Diagonally,78.20%,Medium,58.71%,38.20%,-33.18%,2,1.336,,
280,Wiggle Sort,62.60%,Medium,58.49%,38.20%,-32.84%,2,1.343,,1
921,Minimum Add to Make Parentheses Valid,71.80%,Medium,58.29%,38.20%,-32.51%,2,1.350,,
680,Valid Palindrome II,35.80%,Easy,25.83%,45.10%,35.11%,1,1.351,,
475,Heaters,32.60%,Easy,25.63%,45.10%,35.46%,1,1.355,,
7,Reverse Integer,25.60%,Easy,25.50%,45.10%,35.71%,1,1.357,,
970,Powerful Integers,39.70%,Easy,25.47%,45.10%,35.75%,1,1.357,,
687,Longest Univalue Path,35.20%,Easy,25.12%,45.10%,36.39%,1,1.364,,
1077,Project Employees III,73.60%,Medium,57.80%,38.20%,-31.72%,2,1.366,Sql,1
1033,Moving Stones Until Consecutive,39.80%,Easy,24.65%,45.10%,37.25%,1,1.373,,
442,Find All Duplicates in an Array,64.00%,Medium,57.52%,38.20%,-31.26%,2,1.375,,
840,Magic Squares In Grid,36.70%,Easy,24.38%,45.10%,37.74%,1,1.377,,
408,Valid Word Abbreviation,30.10%,Easy,24.12%,45.10%,38.22%,1,1.382,,1
414,Third Maximum Number,30.00%,Easy,23.93%,45.10%,38.56%,1,1.386,,
362,Design Hit Counter,62.20%,Medium,56.89%,38.20%,-30.24%,2,1.395,,1
78,Subsets,58.00%,Medium,56.86%,38.20%,-30.19%,2,1.396,,
570,Managers with at Least 5 Direct Reports,65.10%,Medium,56.74%,38.20%,-30.00%,2,1.400,Sql,1
633,Sum of Square Numbers,32.40%,Easy,23.12%,45.10%,40.04%,1,1.400,,
1198,Find Smallest Common Element in All Rows,74.30%,Medium,56.73%,38.20%,-29.98%,2,1.400,,1
532,K-diff Pairs in an Array,30.90%,Easy,23.10%,45.10%,40.08%,1,1.401,,
406,Queue Reconstruction by Height,62.60%,Medium,56.65%,38.20%,-29.85%,2,1.403,,
370,Range Addition,62.00%,Medium,56.57%,38.20%,-29.73%,2,1.405,,1
605,Can Place Flowers,31.60%,Easy,22.73%,45.10%,40.75%,1,1.408,,
1347,Minimum Number of Steps to Make Two Strings Anagram,76.20%,Medium,56.44%,38.20%,-29.52%,2,1.410,,
193,Valid Phone Numbers,25.20%,Easy,22.37%,45.10%,41.40%,1,1.414,Shell,
429,N-ary Tree Level Order Traversal,62.60%,Medium,56.31%,38.20%,-29.30%,2,1.414,,
723,Candy Crush,66.90%,Medium,56.30%,38.20%,-29.28%,2,1.414,,1
1037,Valid Boomerang,37.60%,Easy,22.39%,45.10%,41.36%,1,1.414,,
581,Shortest Unsorted Continuous Subarray,30.70%,Easy,22.18%,45.10%,41.75%,1,1.418,,
686,Repeated String Match,32.00%,Easy,21.94%,45.10%,42.19%,1,1.422,,
949,Largest Time for Given Digits,35.70%,Easy,21.78%,45.10%,42.47%,1,1.425,,
941,Valid Mountain Array,35.50%,Easy,21.70%,45.10%,42.63%,1,1.426,,
364,Nested List Weight Sum II,61.20%,Medium,55.86%,38.20%,-28.58%,2,1.428,,1
1346,Check If N and Its Double Exist,41.10%,Easy,21.36%,45.10%,43.24%,1,1.432,,
260,Single Number III,59.50%,Medium,55.69%,38.20%,-28.30%,2,1.434,,
874,Walking Robot Simulation,34.10%,Easy,21.28%,45.10%,43.39%,1,1.434,,
1100,Find K-Length Substrings With No Repeated Characters,71.70%,Medium,55.57%,38.20%,-28.10%,2,1.438,,1
1305,All Elements in Two Binary Search Trees,74.70%,Medium,55.56%,38.20%,-28.09%,2,1.438,,
1364,Number of Trusted Contacts of a Customer  New,75.40%,Medium,55.39%,38.20%,-27.82%,2,1.444,Sql,1
914,X of a Kind in a Deck of Cards,34.00%,Easy,20.59%,45.10%,44.64%,1,1.446,,
750,Number Of Corner Rectangles,66.20%,Medium,55.20%,38.20%,-27.51%,2,1.450,,1
238,Product of Array Except Self,58.60%,Medium,55.11%,38.20%,-27.36%,2,1.453,,
885,Spiral Matrix III,67.70%,Medium,54.72%,38.20%,-26.73%,2,1.465,,
1104,Path In Zigzag Labelled Binary Tree,70.90%,Medium,54.71%,38.20%,-26.71%,2,1.466,,
1126,Active Businesses,71.20%,Medium,54.69%,38.20%,-26.68%,2,1.466,Sql,1
311,Sparse Matrix Multiplication,59.20%,Medium,54.64%,38.20%,-26.60%,2,1.468,,1
608,Tree Node,63.50%,Medium,54.58%,38.20%,-26.51%,2,1.470,Sql,1
1261,Find Elements in a Contaminated Binary Tree,72.90%,Medium,54.41%,38.20%,-26.22%,2,1.476,,
1161,Maximum Level Sum of a Binary Tree,71.40%,Medium,54.37%,38.20%,-26.17%,2,1.477,,
427,Construct Quad Tree,60.50%,Medium,54.24%,38.20%,-25.95%,2,1.481,,
1142,User Activity for the Past 30 Days II,34.90%,Easy,18.15%,45.10%,49.09%,1,1.491,Sql,1
347,Top K Frequent Elements,59.00%,Medium,53.91%,38.20%,-25.42%,2,1.492,,
1325,Delete Leaves With a Given Value,73.30%,Medium,53.87%,38.20%,-25.35%,2,1.493,,
979,Distribute Coins in Binary Tree,68.10%,Medium,53.74%,38.20%,-25.15%,2,1.497,,
702,Search in a Sorted Array of Unknown Size,63.90%,Medium,53.60%,38.20%,-24.93%,2,1.501,,1
12,Integer to Roman,53.60%,Medium,53.42%,38.20%,-24.63%,2,1.507,,
281,Zigzag Iterator,57.50%,Medium,53.38%,38.20%,-24.56%,2,1.509,,1
1045,Customers Who Bought All Products,68.40%,Medium,53.07%,38.20%,-24.07%,2,1.519,Sql,1
1314,Matrix Block Sum,72.30%,Medium,53.03%,38.20%,-23.99%,2,1.520,,
39,Combination Sum,53.40%,Medium,52.83%,38.20%,-23.67%,2,1.527,,
48,Rotate Image,53.50%,Medium,52.80%,38.20%,-23.62%,2,1.528,,
513,Find Bottom Left Tree Value,60.30%,Medium,52.78%,38.20%,-23.59%,2,1.528,,
791,Custom Sort String,64.40%,Medium,52.80%,38.20%,-23.62%,2,1.528,,
1111,Maximum Nesting Depth of Two Valid Parentheses Strings,69.00%,Medium,52.71%,38.20%,-23.47%,2,1.531,,
230,Kth Smallest Element in a BST,56.00%,Medium,52.63%,38.20%,-23.34%,2,1.533,,
484,Find Permutation,59.60%,Medium,52.50%,38.20%,-23.14%,2,1.537,,1
369,Plus One Linked List,57.70%,Medium,52.29%,38.20%,-22.80%,2,1.544,,1
451,Sort Characters By Frequency,58.90%,Medium,52.29%,38.20%,-22.79%,2,1.544,,
859,Buddy Strings,27.70%,Easy,15.10%,45.10%,54.64%,1,1.546,,
49,Group Anagrams,52.80%,Medium,52.08%,38.20%,-22.46%,2,1.551,,
515,Find Largest Value in Each Tree Row,59.50%,Medium,51.95%,38.20%,-22.24%,2,1.555,,
144,Binary Tree Preorder Traversal,54.00%,Medium,51.89%,38.20%,-22.15%,2,1.557,,
156,Binary Tree Upside Down,53.90%,Medium,51.61%,38.20%,-21.70%,2,1.566,,1
951,Flip Equivalent Binary Trees,65.50%,Medium,51.55%,38.20%,-21.61%,2,1.568,,
889,Construct Binary Tree from Preorder and Postorder Traversal,64.50%,Medium,51.46%,38.20%,-21.46%,2,1.571,,
216,Combination Sum III,54.60%,Medium,51.43%,38.20%,-21.41%,2,1.572,,
245,Shortest Word Distance III,54.90%,Medium,51.31%,38.20%,-21.21%,2,1.576,,1
739,Daily Temperatures,62.00%,Medium,51.16%,38.20%,-20.97%,2,1.581,,
1188,Design Bounded Blocking Queue,68.50%,Medium,51.08%,38.20%,-20.83%,2,1.583,Concurrency,1
413,Arithmetic Slices,57.10%,Medium,51.04%,38.20%,-20.78%,2,1.584,,
1204,Last Person to Fit in the Elevator,68.70%,Medium,51.04%,38.20%,-20.78%,2,1.584,Sql,1
77,Combinations,52.10%,Medium,50.97%,38.20%,-20.66%,2,1.587,,
959,Regions Cut By Slashes,65.00%,Medium,50.93%,38.20%,-20.61%,2,1.588,,
986,Interval List Intersections,65.30%,Medium,50.84%,38.20%,-20.45%,2,1.591,,
173,Binary Search Tree Iterator,53.30%,Medium,50.76%,38.20%,-20.33%,2,1.593,,
969,Pancake Sorting,65.00%,Medium,50.79%,38.20%,-20.37%,2,1.593,,
102,Binary Tree Level Order Traversal,52.20%,Medium,50.70%,38.20%,-20.23%,2,1.595,,
431,Encode N-ary Tree to Binary Tree,69.50%,Hard,63.18%,30.80%,-46.79%,3,1.596,,1
426,Convert Binary Search Tree to Sorted Doubly Linked List,56.90%,Medium,50.65%,38.20%,-20.15%,2,1.597,,1
531,Lonely Pixel I,58.40%,Medium,50.61%,38.20%,-20.08%,2,1.598,,1
62,Unique Paths,51.40%,Medium,50.49%,38.20%,-19.89%,2,1.602,,
1112,Highest Grade For Each Student,66.80%,Medium,50.49%,38.20%,-19.89%,2,1.602,Sql,1
582,Kill Process,59.00%,Medium,50.46%,38.20%,-19.84%,2,1.603,,1
59,Spiral Matrix II,51.30%,Medium,50.43%,38.20%,-19.80%,2,1.604,,
841,Keys and Rooms,62.70%,Medium,50.37%,38.20%,-19.69%,2,1.606,,
695,Max Area of Island,60.50%,Medium,50.31%,38.20%,-19.59%,2,1.608,,
877,Stone Game,63.10%,Medium,50.24%,38.20%,-19.48%,2,1.610,,
1193,Monthly Transactions I,67.60%,Medium,50.10%,38.20%,-19.26%,2,1.615,Sql,1
64,Minimum Path Sum,51.00%,Medium,50.06%,38.20%,-19.19%,2,1.616,,
1361,Validate Binary Tree Nodes,69.90%,Medium,49.94%,38.20%,-18.99%,2,1.620,,
626,Exchange Seats,59.00%,Medium,49.82%,38.20%,-18.80%,2,1.624,Sql,
1222,Queens That Can Attack the King,67.70%,Medium,49.78%,38.20%,-18.73%,2,1.625,,
241,Different Ways to Add Parentheses,53.30%,Medium,49.77%,38.20%,-18.71%,2,1.626,,
647,Palindromic Substrings,59.20%,Medium,49.71%,38.20%,-18.63%,2,1.627,,
540,Single Element in a Sorted Array,57.60%,Medium,49.68%,38.20%,-18.58%,2,1.628,,
1277,Count Square Submatrices with All Ones,68.40%,Medium,49.67%,38.20%,-18.56%,2,1.629,,
323,Number of Connected Components in an Undirected Graph,54.30%,Medium,49.56%,38.20%,-18.39%,2,1.632,,1
215,Kth Largest Element in an Array,52.50%,Medium,49.35%,38.20%,-18.04%,2,1.639,,
510,Inorder Successor in BST II,56.70%,Medium,49.22%,38.20%,-17.83%,2,1.643,,1
529,Minesweeper,57.00%,Medium,49.24%,38.20%,-17.87%,2,1.643,,
912,Sort an Array,62.60%,Medium,49.22%,38.20%,-17.84%,2,1.643,,
1123,Lowest Common Ancestor of Deepest Leaves,65.70%,Medium,49.23%,38.20%,-17.85%,2,1.643,,
508,Most Frequent Subtree Sum,56.60%,Medium,49.15%,38.20%,-17.72%,2,1.646,,
665,Non-decreasing Array,19.40%,Easy,9.65%,45.10%,64.58%,1,1.646,,
609,Find Duplicate File in System,58.00%,Medium,49.07%,38.20%,-17.59%,2,1.648,,
1195,Fizz Buzz Multithreaded,66.60%,Medium,49.07%,38.20%,-17.59%,2,1.648,Concurrency,
526,Beautiful Arrangement,56.70%,Medium,48.99%,38.20%,-17.45%,2,1.651,,
1061,Lexicographically Smallest Equivalent String,64.50%,Medium,48.94%,38.20%,-17.38%,2,1.652,,1
11,Container With Most Water,49.00%,Medium,48.84%,38.20%,-17.21%,2,1.656,,
1110,Delete Nodes And Return Forest,65.10%,Medium,48.82%,38.20%,-17.18%,2,1.656,,
199,Binary Tree Right Side View,51.70%,Medium,48.78%,38.20%,-17.12%,2,1.658,,
547,Friend Circles,56.80%,Medium,48.78%,38.20%,-17.12%,2,1.658,,
1101,The Earliest Moment When Everyone Become Friends,64.90%,Medium,48.75%,38.20%,-17.07%,2,1.659,,1
439,Ternary Expression Parser,55.10%,Medium,48.66%,38.20%,-16.93%,2,1.661,,1
612,Shortest Distance in a Plane,57.60%,Medium,48.62%,38.20%,-16.87%,2,1.663,Sql,1
287,Find the Duplicate Number,52.70%,Medium,48.49%,38.20%,-16.65%,2,1.667,,
96,Unique Binary Search Trees,49.70%,Medium,48.29%,38.20%,-16.33%,2,1.673,,
1043,Partition Array for Maximum Sum,63.60%,Medium,48.30%,38.20%,-16.35%,2,1.673,,
1130,Minimum Cost Tree From Leaf Values,64.80%,Medium,48.23%,38.20%,-16.22%,2,1.676,,
286,Walls and Gates,52.30%,Medium,48.11%,38.20%,-16.03%,2,1.679,,1
553,Optimal Division,56.20%,Medium,48.09%,38.20%,-16.00%,2,1.680,,
1214,Two Sum BSTs,65.90%,Medium,48.09%,38.20%,-16.01%,2,1.680,,1
24,Swap Nodes in Pairs,48.30%,Medium,47.95%,38.20%,-15.77%,2,1.685,,
635,Design Log Storage System,57.20%,Medium,47.89%,38.20%,-15.67%,2,1.687,,1
1026,Maximum Difference Between Node and Ancestor,62.90%,Medium,47.85%,38.20%,-15.62%,2,1.688,,
249,Group Shifted Strings,51.40%,Medium,47.75%,38.20%,-15.45%,2,1.691,,1
244,Shortest Word Distance II,51.20%,Medium,47.62%,38.20%,-15.24%,2,1.695,,1
348,Design Tic-Tac-Toe,52.70%,Medium,47.60%,38.20%,-15.20%,2,1.696,,1
998,Maximum Binary Tree II,62.10%,Medium,47.46%,38.20%,-14.99%,2,1.700,,
973,K Closest Points to Origin,61.70%,Medium,47.43%,38.20%,-14.93%,2,1.701,,
398,Random Pick Index,53.20%,Medium,47.36%,38.20%,-14.83%,2,1.703,,
250,Count Univalue Subtrees,51.00%,Medium,47.33%,38.20%,-14.78%,2,1.704,,1
328,Odd Even Linked List,52.10%,Medium,47.29%,38.20%,-14.71%,2,1.706,,
931,Minimum Falling Path Sum,60.90%,Medium,47.25%,38.20%,-14.64%,2,1.707,,
1264,Page Recommendations,65.80%,Medium,47.26%,38.20%,-14.66%,2,1.707,Sql,1
1310,XOR Queries of a Subarray,66.40%,Medium,47.19%,38.20%,-14.54%,2,1.709,,
289,Game of Life,51.40%,Medium,47.16%,38.20%,-14.50%,2,1.710,,
489,Robot Room Cleaner,67.70%,Hard,60.53%,30.80%,-42.96%,3,1.711,,1
573,Squirrel Simulation,55.30%,Medium,46.90%,38.20%,-14.07%,2,1.719,,1
1338,Reduce Array Size to The Half,66.50%,Medium,46.88%,38.20%,-14.04%,2,1.719,,
378,Kth Smallest Element in a Sorted Matrix,52.30%,Medium,46.76%,38.20%,-13.84%,2,1.723,,
865,Smallest Subtree with all the Deepest Nodes,59.40%,Medium,46.71%,38.20%,-13.78%,2,1.724,,
503,Next Greater Element II,54.00%,Medium,46.62%,38.20%,-13.63%,2,1.727,,
89,Gray Code,47.90%,Medium,46.59%,38.20%,-13.58%,2,1.728,,
421,Maximum XOR of Two Numbers in an Array,52.70%,Medium,46.53%,38.20%,-13.47%,2,1.731,,
137,Single Number II,48.50%,Medium,46.49%,38.20%,-13.42%,2,1.732,,
712,Minimum ASCII Delete Sum for Two Strings,56.90%,Medium,46.46%,38.20%,-13.36%,2,1.733,,
462,Minimum Moves to Equal Array Elements II,53.20%,Medium,46.42%,38.20%,-13.31%,2,1.734,,
856,Score of Parentheses,58.90%,Medium,46.35%,38.20%,-13.18%,2,1.736,,
565,Array Nesting,54.60%,Medium,46.31%,38.20%,-13.13%,2,1.737,,
159,Longest Substring with At Most Two Distinct Characters,48.60%,Medium,46.27%,38.20%,-13.06%,2,1.739,,1
384,Shuffle an Array,51.90%,Medium,46.27%,38.20%,-13.06%,2,1.739,,
445,Add Two Numbers II,52.80%,Medium,46.27%,38.20%,-13.06%,2,1.739,,
1286,Iterator for Combination,65.10%,Medium,46.24%,38.20%,-13.01%,2,1.740,,
341,Flatten Nested List Iterator,51.20%,Medium,46.20%,38.20%,-12.94%,2,1.741,,
36,Valid Sudoku,46.70%,Medium,46.17%,38.20%,-12.90%,2,1.742,,
1256,Encode Number,64.50%,Medium,46.08%,38.20%,-12.75%,2,1.745,,1
320,Generalized Abbreviation,50.70%,Medium,46.01%,38.20%,-12.63%,2,1.747,,1
1087,Brace Expansion,61.90%,Medium,45.96%,38.20%,-12.55%,2,1.749,,1
495,Teemo Attacking,53.20%,Medium,45.94%,38.20%,-12.52%,2,1.750,,
454,4Sum II,52.30%,Medium,45.64%,38.20%,-12.04%,2,1.759,,
1219,Path with Maximum Gold,63.50%,Medium,45.62%,38.20%,-12.01%,2,1.760,,
1120,Maximum Average Subtree,62.00%,Medium,45.57%,38.20%,-11.93%,2,1.761,,1
318,Maximum Product of Word Lengths,50.20%,Medium,45.54%,38.20%,-11.87%,2,1.763,,
16,3Sum Closest,45.70%,Medium,45.47%,38.20%,-11.76%,2,1.765,,
382,Linked List Random Node,51.00%,Medium,45.40%,38.20%,-11.65%,2,1.767,,
648,Replace Words,54.90%,Medium,45.40%,38.20%,-11.64%,2,1.767,,
684,Redundant Connection,55.30%,Medium,45.27%,38.20%,-11.44%,2,1.771,,
1236,Web Crawler,63.40%,Medium,45.27%,38.20%,-11.44%,2,1.771,,1
40,Combination Sum II,45.70%,Medium,45.11%,38.20%,-11.19%,2,1.776,,
789,Escape The Ghosts,56.70%,Medium,45.13%,38.20%,-11.21%,2,1.776,,
294,Flip Game II,49.40%,Medium,45.09%,38.20%,-11.15%,2,1.777,,1
946,Validate Stack Sequences,58.80%,Medium,44.93%,38.20%,-10.88%,2,1.782,,
114,Flatten Binary Tree to Linked List,46.50%,Medium,44.83%,38.20%,-10.72%,2,1.786,,
337,House Robber III,49.70%,Medium,44.76%,38.20%,-10.61%,2,1.788,,
17,Letter Combinations of a Phone Number,44.90%,Medium,44.65%,38.20%,-10.44%,2,1.791,,
666,Path Sum IV,54.30%,Medium,44.53%,38.20%,-10.25%,2,1.795,,1
1164,Product Price at a Given Date,61.60%,Medium,44.53%,38.20%,-10.24%,2,1.795,Sql,1
1167,Minimum Cost to Connect Sticks,61.60%,Medium,44.48%,38.20%,-10.17%,2,1.797,,1
817,Linked List Components,56.40%,Medium,44.42%,38.20%,-10.06%,2,1.799,,
1072,Flip Columns For Maximum Number of Equal Rows,60.10%,Medium,44.38%,38.20%,-10.00%,2,1.800,,
1016,Binary String With Substrings Representing 1 To N,59.20%,Medium,44.30%,38.20%,-9.87%,2,1.803,,
105,Construct Binary Tree from Preorder and Inorder Traversal,45.80%,Medium,44.26%,38.20%,-9.81%,2,1.804,,
343,Integer Break,49.30%,Medium,44.27%,38.20%,-9.82%,2,1.804,,
694,Number of Distinct Islands,54.40%,Medium,44.22%,38.20%,-9.74%,2,1.805,,1
835,Image Overlap,56.40%,Medium,44.15%,38.20%,-9.63%,2,1.807,,
399,Evaluate Division,50.00%,Medium,44.15%,38.20%,-9.62%,2,1.808,,
655,Print Binary Tree,53.70%,Medium,44.09%,38.20%,-9.54%,2,1.809,,
1238,Circular Permutation in Binary Representation,62.20%,Medium,44.04%,38.20%,-9.45%,2,1.811,,
1343,Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold,63.70%,Medium,44.00%,38.20%,-9.39%,2,1.812,,
90,Subsets II,45.30%,Medium,43.98%,38.20%,-9.35%,2,1.813,,
1140,Stone Game II,60.70%,Medium,43.98%,38.20%,-9.35%,2,1.813,,
983,Minimum Cost For Tickets,58.30%,Medium,43.88%,38.20%,-9.20%,2,1.816,,
386,Lexicographical Numbers,49.50%,Medium,43.84%,38.20%,-9.12%,2,1.818,,
449,Serialize and Deserialize BST,50.40%,Medium,43.81%,38.20%,-9.09%,2,1.818,,
103,Binary Tree Zigzag Level Order Traversal,45.30%,Medium,43.79%,38.20%,-9.04%,2,1.819,,
980,Unique Paths III,72.40%,Hard,58.03%,30.80%,-39.34%,3,1.820,,
1258,Synonymous Sentences,62.20%,Medium,43.75%,38.20%,-8.98%,2,1.820,,1
1273,Delete Tree Nodes,62.30%,Medium,43.63%,38.20%,-8.79%,2,1.824,,1
47,Permutations II,44.30%,Medium,43.61%,38.20%,-8.76%,2,1.825,,
75,Sort Colors,44.70%,Medium,43.60%,38.20%,-8.74%,2,1.825,,
129,Sum Root to Leaf Numbers,45.50%,Medium,43.61%,38.20%,-8.75%,2,1.825,,
667,Beautiful Arrangement II,53.40%,Medium,43.62%,38.20%,-8.77%,2,1.825,,
109,Convert Sorted List to Binary Search Tree,45.00%,Medium,43.40%,38.20%,-8.42%,2,1.832,,
585,Investments in 2016,51.90%,Medium,43.32%,38.20%,-8.28%,2,1.834,Sql,1
1321,Restaurant Growth,62.70%,Medium,43.33%,38.20%,-8.29%,2,1.834,Sql,1
477,Total Hamming Distance,50.20%,Medium,43.20%,38.20%,-8.10%,2,1.838,,
676,Implement Magic Dictionary,53.00%,Medium,43.09%,38.20%,-7.91%,2,1.842,,
1017,Convert to Base -2,58.00%,Medium,43.08%,38.20%,-7.90%,2,1.842,,
516,Longest Palindromic Subsequence,50.60%,Medium,43.03%,38.20%,-7.82%,2,1.844,,
932,Beautiful Array,56.70%,Medium,43.03%,38.20%,-7.82%,2,1.844,,
490,The Maze,50.10%,Medium,42.91%,38.20%,-7.63%,2,1.847,,1
430,Flatten a Multilevel Doubly Linked List,49.20%,Medium,42.89%,38.20%,-7.59%,2,1.848,,
131,Palindrome Partitioning,44.80%,Medium,42.88%,38.20%,-7.57%,2,1.849,,
247,Strobogrammatic Number II,46.50%,Medium,42.88%,38.20%,-7.57%,2,1.849,,1
677,Map Sum Pairs,52.80%,Medium,42.87%,38.20%,-7.56%,2,1.849,,
259,3Sum Smaller,46.60%,Medium,42.80%,38.20%,-7.45%,2,1.851,,1
1057,Campus Bikes,58.30%,Medium,42.80%,38.20%,-7.44%,2,1.851,,1
113,Path Sum II,44.40%,Medium,42.74%,38.20%,-7.35%,2,1.853,,
360,Sort Transformed Array,48.00%,Medium,42.72%,38.20%,-7.31%,2,1.854,,1
776,Split BST,54.10%,Medium,42.72%,38.20%,-7.31%,2,1.854,,1
919,Complete Binary Tree Inserter,56.20%,Medium,42.72%,38.20%,-7.32%,2,1.854,,
539,Minimum Time Difference,50.60%,Medium,42.69%,38.20%,-7.27%,2,1.855,,
357,Count Numbers with Unique Digits,47.90%,Medium,42.66%,38.20%,-7.22%,2,1.856,,
714,Best Time to Buy and Sell Stock with Transaction Fee,53.10%,Medium,42.63%,38.20%,-7.17%,2,1.857,,
756,Pyramid Transition Matrix,53.60%,Medium,42.51%,38.20%,-6.98%,2,1.860,,
769,Max Chunks To Make Sorted,53.80%,Medium,42.52%,38.20%,-6.99%,2,1.860,,
1288,Remove Covered Intervals,61.40%,Medium,42.51%,38.20%,-6.97%,2,1.861,,
1090,Largest Values From Labels,58.40%,Medium,42.41%,38.20%,-6.82%,2,1.864,,
351,Android Unlock Patterns,47.50%,Medium,42.35%,38.20%,-6.72%,2,1.866,,1
394,Decode String,48.10%,Medium,42.32%,38.20%,-6.67%,2,1.867,,
254,Factor Combinations,46.00%,Medium,42.27%,38.20%,-6.59%,2,1.868,,1
651,4 Keys Keyboard,51.70%,Medium,42.15%,38.20%,-6.39%,2,1.872,,1
1190,Reverse Substrings Between Each Pair of Parentheses,59.60%,Medium,42.15%,38.20%,-6.39%,2,1.872,,
153,Find Minimum in Rotated Sorted Array,44.20%,Medium,41.96%,38.20%,-6.08%,2,1.878,,
1004,Max Consecutive Ones III,56.70%,Medium,41.97%,38.20%,-6.11%,2,1.878,,
1306,Jump Game III,61.10%,Medium,41.95%,38.20%,-6.06%,2,1.879,,
200,Number of Islands,44.80%,Medium,41.87%,38.20%,-5.93%,2,1.881,,
1318,Minimum Flips to Make a OR b Equal to c,61.20%,Medium,41.87%,38.20%,-5.94%,2,1.881,,
781,Rabbits in Forest,53.30%,Medium,41.85%,38.20%,-5.90%,2,1.882,,
1011,Capacity To Ship Packages Within D Days,56.60%,Medium,41.77%,38.20%,-5.78%,2,1.884,,
298,Binary Tree Longest Consecutive Sequence,46.10%,Medium,41.73%,38.20%,-5.71%,2,1.886,,1
1249,Minimum Remove to Make Valid Parentheses,60.00%,Medium,41.68%,38.20%,-5.63%,2,1.887,,
1019,Next Greater Node In Linked List,56.60%,Medium,41.65%,38.20%,-5.59%,2,1.888,,
1151,Minimum Swaps to Group All 1's Together,58.50%,Medium,41.62%,38.20%,-5.53%,2,1.889,,1
641,Design Circular Deque,51.00%,Medium,41.60%,38.20%,-5.50%,2,1.890,,
452,Minimum Number of Arrows to Burst Balloons,48.20%,Medium,41.57%,38.20%,-5.45%,2,1.891,,
106,Construct Binary Tree from Inorder and Postorder Traversal,43.10%,Medium,41.55%,38.20%,-5.41%,2,1.892,,
255,Verify Preorder Sequence in Binary Search Tree,45.20%,Medium,41.46%,38.20%,-5.28%,2,1.894,,1
1344,Angle Between Hands of a Clock,61.20%,Medium,41.49%,38.20%,-5.32%,2,1.894,,
325,Maximum Size Subarray Sum Equals k,46.20%,Medium,41.43%,38.20%,-5.23%,2,1.895,,1
1055,Shortest Way to Form String,56.90%,Medium,41.43%,38.20%,-5.22%,2,1.896,,1
80,Remove Duplicates from Sorted Array II,42.50%,Medium,41.33%,38.20%,-5.06%,2,1.899,,
251,Flatten 2D Vector,45.00%,Medium,41.32%,38.20%,-5.05%,2,1.899,,1
1105,Filling Bookcase Shelves,57.50%,Medium,41.29%,38.20%,-5.01%,2,1.900,,
636,Exclusive Time of Functions,50.60%,Medium,41.27%,38.20%,-4.97%,2,1.901,,
1115,Print FooBar Alternately,57.60%,Medium,41.25%,38.20%,-4.93%,2,1.901,Concurrency,
947,Most Stones Removed with Same Row or Column,55.10%,Medium,41.21%,38.20%,-4.87%,2,1.903,,
1268,Search Suggestions System,59.80%,Medium,41.20%,38.20%,-4.86%,2,1.903,,
487,Max Consecutive Ones II,48.30%,Medium,41.16%,38.20%,-4.79%,2,1.904,,1
1254,Number of Closed Islands,59.50%,Medium,41.11%,38.20%,-4.71%,2,1.906,,
309,Best Time to Buy and Sell Stock with Cooldown,45.60%,Medium,41.07%,38.20%,-4.64%,2,1.907,,
602,Friend Requests II: Who Has the Most Friends,49.90%,Medium,41.07%,38.20%,-4.65%,2,1.907,Sql,1
1031,Maximum Sum of Two Non-Overlapping Subarrays,56.20%,Medium,41.08%,38.20%,-4.66%,2,1.907,,
638,Shopping Offers,50.40%,Medium,41.04%,38.20%,-4.60%,2,1.908,,
1143,Longest Common Subsequence,57.80%,Medium,41.04%,38.20%,-4.59%,2,1.908,,
646,Maximum Length of Pair Chain,50.50%,Medium,41.03%,38.20%,-4.57%,2,1.909,,
1227,Airplane Seat Assignment Probability,59.00%,Medium,41.00%,38.20%,-4.54%,2,1.909,,
253,Meeting Rooms II,44.70%,Medium,40.99%,38.20%,-4.51%,2,1.910,,1
554,Brick Wall,49.10%,Medium,40.97%,38.20%,-4.49%,2,1.910,,
73,Set Matrix Zeroes,42.00%,Medium,40.93%,38.20%,-4.42%,2,1.912,,
1158,Market Analysis I,57.90%,Medium,40.92%,38.20%,-4.39%,2,1.912,Sql,1
208,Implement Trie (Prefix Tree),43.90%,Medium,40.85%,38.20%,-4.29%,2,1.914,,
1357,Apply Discount Every n Orders,60.70%,Medium,40.80%,38.20%,-4.20%,2,1.916,,
1094,Car Pooling,56.80%,Medium,40.75%,38.20%,-4.13%,2,1.917,,
1020,Number of Enclaves,55.70%,Medium,40.74%,38.20%,-4.11%,2,1.918,,
1355,Activity Participants,60.60%,Medium,40.73%,38.20%,-4.09%,2,1.918,Sql,1
120,Triangle,42.40%,Medium,40.64%,38.20%,-3.95%,2,1.921,,
672,Bulb Switcher II,50.50%,Medium,40.64%,38.20%,-3.95%,2,1.921,,
1023,Camelcase Matching,55.60%,Medium,40.60%,38.20%,-3.88%,2,1.922,,
279,Perfect Squares,44.60%,Medium,40.51%,38.20%,-3.73%,2,1.925,,
858,Mirror Reflection,53.10%,Medium,40.52%,38.20%,-3.75%,2,1.925,,
319,Bulb Switcher,45.00%,Medium,40.32%,38.20%,-3.43%,2,1.931,,
116,Populating Next Right Pointers in Each Node,42.00%,Medium,40.30%,38.20%,-3.40%,2,1.932,,
423,Reconstruct Original Digits from English,46.50%,Medium,40.30%,38.20%,-3.39%,2,1.932,,
486,Predict the Winner,47.40%,Medium,40.27%,38.20%,-3.35%,2,1.933,,
162,Find Peak Element,42.60%,Medium,40.22%,38.20%,-3.28%,2,1.934,,
863,All Nodes Distance K in Binary Tree,52.90%,Medium,40.24%,38.20%,-3.31%,2,1.934,,
1247,Minimum Swaps to Make Strings Equal,58.50%,Medium,40.21%,38.20%,-3.25%,2,1.935,,
725,Split Linked List in Parts,50.80%,Medium,40.17%,38.20%,-3.18%,2,1.936,,
1174,Immediate Food Delivery II,57.30%,Medium,40.08%,38.20%,-3.04%,2,1.939,Sql,1
178,Rank Scores,42.60%,Medium,39.99%,38.20%,-2.90%,2,1.942,Sql,
481,Magical String,47.00%,Medium,39.95%,38.20%,-2.82%,2,1.944,,
361,Bomb Enemy,45.20%,Medium,39.91%,38.20%,-2.76%,2,1.945,,1
524,Longest Word in Dictionary through Deleting,47.60%,Medium,39.91%,38.20%,-2.77%,2,1.945,,
901,Online Stock Span,53.10%,Medium,39.89%,38.20%,-2.73%,2,1.945,,
1135,Connecting Cities With Minimum Cost,56.50%,Medium,39.85%,38.20%,-2.68%,2,1.946,,1
379,Design Phone Directory,45.40%,Medium,39.84%,38.20%,-2.66%,2,1.947,,1
380,Insert Delete GetRandom O(1),45.40%,Medium,39.83%,38.20%,-2.63%,2,1.947,,
1166,Design File System,56.80%,Medium,39.70%,38.20%,-2.43%,2,1.951,,1
1267,Count Servers that Communicate,58.30%,Medium,39.72%,38.20%,-2.46%,2,1.951,,
533,Lonely Pixel II,47.50%,Medium,39.68%,38.20%,-2.40%,2,1.952,,1
869,Reordered Power of 2,52.30%,Medium,39.55%,38.20%,-2.19%,2,1.956,,
52,N-Queens II,55.60%,Hard,54.84%,30.80%,-34.74%,3,1.958,,
729,My Calendar I,50.20%,Medium,39.51%,38.20%,-2.12%,2,1.958,,
623,Add One Row to Tree,48.60%,Medium,39.46%,38.20%,-2.04%,2,1.959,,
1324,Print Words Vertically,58.90%,Medium,39.48%,38.20%,-2.07%,2,1.959,,
498,Diagonal Traverse,46.70%,Medium,39.40%,38.20%,-1.94%,2,1.961,,
592,Fraction Addition and Subtraction,48.10%,Medium,39.42%,38.20%,-1.97%,2,1.961,,
900,RLE Iterator,52.50%,Medium,39.30%,38.20%,-1.78%,2,1.964,,
846,Hand of Straights,51.70%,Medium,39.29%,38.20%,-1.77%,2,1.965,,
1003,Check If Word Is Valid After Substitutions,54.00%,Medium,39.29%,38.20%,-1.76%,2,1.965,,
536,Construct Binary Tree from String,47.10%,Medium,39.24%,38.20%,-1.68%,2,1.966,,1
652,Find Duplicate Subtrees,48.80%,Medium,39.24%,38.20%,-1.68%,2,1.966,,
284,Peeking Iterator,43.40%,Medium,39.23%,38.20%,-1.67%,2,1.967,,
424,Longest Repeating Character Replacement,45.40%,Medium,39.18%,38.20%,-1.59%,2,1.968,,
494,Target Sum,46.40%,Medium,39.15%,38.20%,-1.54%,2,1.969,,
692,Top K Frequent Words,49.30%,Medium,39.15%,38.20%,-1.54%,2,1.969,,
236,Lowest Common Ancestor of a Binary Tree,42.60%,Medium,39.14%,38.20%,-1.52%,2,1.970,,
313,Super Ugly Number,43.70%,Medium,39.11%,38.20%,-1.47%,2,1.971,,
470,Implement Rand10() Using Rand7(),46.00%,Medium,39.11%,38.20%,-1.47%,2,1.971,,
505,The Maze II,46.50%,Medium,39.09%,38.20%,-1.45%,2,1.971,,1
377,Combination Sum IV,44.50%,Medium,38.97%,38.20%,-1.25%,2,1.975,,
1052,Grumpy Bookstore Owner,54.40%,Medium,38.97%,38.20%,-1.25%,2,1.975,,
1209,Remove All Adjacent Duplicates in String II,56.70%,Medium,38.97%,38.20%,-1.24%,2,1.975,,
1060,Missing Element in Sorted Array,54.50%,Medium,38.95%,38.20%,-1.22%,2,1.976,,1
1062,Longest Repeating Substring,54.50%,Medium,38.92%,38.20%,-1.17%,2,1.977,,1
240,Search a 2D Matrix II,42.40%,Medium,38.88%,38.20%,-1.10%,2,1.978,,
1006,Clumsy Factorial,53.60%,Medium,38.85%,38.20%,-1.04%,2,1.979,,
518,Coin Change 2,46.40%,Medium,38.80%,38.20%,-0.98%,2,1.980,,
1028,Recover a Tree From Preorder Traversal,69.40%,Hard,54.32%,30.80%,-33.99%,3,1.980,,
1063,Number of Valid Subarrays,69.90%,Hard,54.31%,30.80%,-33.97%,3,1.981,,1
314,Binary Tree Vertical Order Traversal,43.30%,Medium,38.69%,38.20%,-0.80%,2,1.984,,1
186,Reverse Words in a String II,41.40%,Medium,38.67%,38.20%,-0.76%,2,1.985,,1
390,Elimination Game,44.30%,Medium,38.58%,38.20%,-0.61%,2,1.988,,
650,2 Keys Keyboard,48.10%,Medium,38.57%,38.20%,-0.59%,2,1.988,,
86,Partition List,39.80%,Medium,38.54%,38.20%,-0.55%,2,1.989,,
611,Valid Triangle Number,47.50%,Medium,38.54%,38.20%,-0.55%,2,1.989,,
621,Task Scheduler,47.60%,Medium,38.49%,38.20%,-0.47%,2,1.991,,
583,Delete Operation for Two Strings,46.80%,Medium,38.25%,38.20%,-0.08%,2,1.998,,
752,Open the Lock,49.30%,Medium,38.27%,38.20%,-0.11%,2,1.998,,
939,Minimum Area Rectangle,52.00%,Medium,38.23%,38.20%,-0.05%,2,1.999,,
1027,Longest Arithmetic Sequence,53.30%,Medium,38.24%,38.20%,-0.06%,2,1.999,,
1245,Tree Diameter,56.50%,Medium,38.24%,38.20%,-0.06%,2,1.999,,1
1233,Remove Sub-Folders from the Filesystem,56.20%,Medium,38.12%,38.20%,0.14%,2,2.003,,
1066,Campus Bikes II,53.70%,Medium,38.07%,38.20%,0.22%,2,2.004,,1
549,Binary Tree Longest Consecutive Sequence II,46.10%,Medium,38.05%,38.20%,0.25%,2,2.005,,1
550,Game Play Analysis IV,46.10%,Medium,38.03%,38.20%,0.27%,2,2.005,Sql,1
1035,Uncrossed Lines,53.20%,Medium,38.02%,38.20%,0.29%,2,2.006,,
718,Maximum Length of Repeated Subarray,48.50%,Medium,37.97%,38.20%,0.37%,2,2.007,,
580,Count Student Number in Departments,46.40%,Medium,37.89%,38.20%,0.50%,2,2.010,Sql,1
1257,Smallest Common Region,56.30%,Medium,37.86%,38.20%,0.54%,2,2.011,,1
851,Loud and Rich,50.30%,Medium,37.82%,38.20%,0.62%,2,2.012,,
926,Flip String to Monotone Increasing,51.40%,Medium,37.82%,38.20%,0.62%,2,2.012,,
436,Find Right Interval,44.20%,Medium,37.81%,38.20%,0.64%,2,2.013,,
207,Course Schedule,40.80%,Medium,37.76%,38.20%,0.71%,2,2.014,,
222,Count Complete Tree Nodes,40.90%,Medium,37.64%,38.20%,0.90%,2,2.018,,
1048,Longest String Chain,53.00%,Medium,37.63%,38.20%,0.92%,2,2.018,,
1272,Remove Interval,56.30%,Medium,37.64%,38.20%,0.90%,2,2.018,,1
981,Time Based Key-Value Store,51.90%,Medium,37.51%,38.20%,1.11%,2,2.022,,
1116,Print Zero Even Odd,53.90%,Medium,37.53%,38.20%,1.08%,2,2.022,Concurrency,
300,Longest Increasing Subsequence,41.90%,Medium,37.50%,38.20%,1.13%,2,2.023,,
147,Insertion Sort List,39.60%,Medium,37.44%,38.20%,1.22%,2,2.024,,
261,Graph Valid Tree,41.20%,Medium,37.37%,38.20%,1.34%,2,2.027,,1
820,Short Encoding of Words,49.40%,Medium,37.37%,38.20%,1.34%,2,2.027,,
1244,Design A Leaderboard,55.60%,Medium,37.35%,38.20%,1.37%,2,2.027,,1
148,Sort List,39.50%,Medium,37.33%,38.20%,1.41%,2,2.028,,
759,Employee Free Time,64.30%,Hard,53.17%,30.80%,-32.32%,3,2.030,,1
95,Unique Binary Search Trees II,38.60%,Medium,37.21%,38.20%,1.61%,2,2.032,,
688,Knight Probability in Chessboard,47.30%,Medium,37.21%,38.20%,1.60%,2,2.032,,
56,Merge Intervals,38.00%,Medium,37.18%,38.20%,1.65%,2,2.033,,
562,Longest Line of Consecutive One in Matrix,45.40%,Medium,37.16%,38.20%,1.69%,2,2.034,,1
958,Check Completeness of a Binary Tree,51.20%,Medium,37.15%,38.20%,1.70%,2,2.034,,
525,Contiguous Array,44.80%,Medium,37.10%,38.20%,1.78%,2,2.036,,
740,Delete and Earn,47.90%,Medium,37.05%,38.20%,1.87%,2,2.037,,
833,Find And Replace in String,49.20%,Medium,36.98%,38.20%,1.97%,2,2.039,,
731,My Calendar II,47.70%,Medium,36.98%,38.20%,1.98%,2,2.040,,
1014,Best Sightseeing Pair,51.70%,Medium,36.83%,38.20%,2.22%,2,2.044,,
491,Increasing Subsequences,44.00%,Medium,36.80%,38.20%,2.27%,2,2.045,,
875,Koko Eating Bananas,49.60%,Medium,36.77%,38.20%,2.32%,2,2.046,,
1040,Moving Stones Until Consecutive II,51.80%,Medium,36.55%,38.20%,2.68%,2,2.054,,
296,Best Meeting Point,56.90%,Hard,52.56%,30.80%,-31.44%,3,2.057,,1
1131,Maximum of Absolute Value Expression,53.00%,Medium,36.41%,38.20%,2.89%,2,2.058,,
548,Split Array with Equal Sum,44.40%,Medium,36.36%,38.20%,2.97%,2,2.059,,1
813,Largest Sum of Averages,48.30%,Medium,36.38%,38.20%,2.95%,2,2.059,,
1212,Team Scores in Football Tournament,54.10%,Medium,36.32%,38.20%,3.04%,2,2.061,Sql,1
139,Word Break,38.30%,Medium,36.26%,38.20%,3.14%,2,2.063,,
416,Partition Equal Subset Sum,42.30%,Medium,36.20%,38.20%,3.24%,2,2.065,,
277,Find the Celebrity,40.10%,Medium,36.04%,38.20%,3.50%,2,2.070,,1
1248,Count Number of Nice Subarrays,54.30%,Medium,36.00%,38.20%,3.57%,2,2.071,,
92,Reverse Linked List II,37.30%,Medium,35.95%,38.20%,3.64%,2,2.073,,
911,Online Election,49.30%,Medium,35.94%,38.20%,3.66%,2,2.073,,
435,Non-overlapping Intervals,42.20%,Medium,35.82%,38.20%,3.85%,2,2.077,,
528,Random Pick with Weight,43.40%,Medium,35.66%,38.20%,4.12%,2,2.082,,
558,Logical OR of Two Binary Grids Represented as Quad-Trees,43.60%,Medium,35.42%,38.20%,4.50%,2,2.090,,
117,Populating Next Right Pointers in Each Node II,37.10%,Medium,35.38%,38.20%,4.56%,2,2.091,,
802,Find Eventual Safe States,47.10%,Medium,35.34%,38.20%,4.63%,2,2.093,,
388,Longest Absolute File Path,40.90%,Medium,35.21%,38.20%,4.84%,2,2.097,,
450,Delete Node in a BST,41.80%,Medium,35.20%,38.20%,4.85%,2,2.097,,
560,Subarray Sum Equals K,43.40%,Medium,35.19%,38.20%,4.88%,2,2.098,,
767,Reorganize String,46.40%,Medium,35.15%,38.20%,4.93%,2,2.099,,
1007,Minimum Domino Rotations For Equal Row,49.90%,Medium,35.13%,38.20%,4.97%,2,2.099,,
1182,Shortest Distance to Target Color,52.40%,Medium,35.06%,38.20%,5.07%,2,2.101,,1
187,Repeated DNA Sequences,37.80%,Medium,35.06%,38.20%,5.09%,2,2.102,,
210,Course Schedule II,38.10%,Medium,35.02%,38.20%,5.15%,2,2.103,,
721,Accounts Merge,45.60%,Medium,35.03%,38.20%,5.14%,2,2.103,,
334,Increasing Triplet Subsequence,39.80%,Medium,34.90%,38.20%,5.34%,2,2.107,,
792,Number of Matching Subsequences,46.50%,Medium,34.88%,38.20%,5.37%,2,2.107,,
180,Consecutive Numbers,37.50%,Medium,34.86%,38.20%,5.40%,2,2.108,Sql,
228,Summary Ranges,38.20%,Medium,34.86%,38.20%,5.41%,2,2.108,,
474,Ones and Zeroes,41.80%,Medium,34.85%,38.20%,5.42%,2,2.108,,
264,Ugly Number II,38.70%,Medium,34.83%,38.20%,5.46%,2,2.109,,
331,Verify Preorder Serialization of a Binary Tree,39.70%,Medium,34.85%,38.20%,5.43%,2,2.109,,
74,Search a 2D Matrix,35.90%,Medium,34.81%,38.20%,5.48%,2,2.110,,
873,Length of Longest Fibonacci Subsequence,47.60%,Medium,34.80%,38.20%,5.51%,2,2.110,,
1109,Corporate Flight Bookings,51.00%,Medium,34.73%,38.20%,5.61%,2,2.112,,
6,ZigZag Conversion,34.80%,Medium,34.71%,38.20%,5.64%,2,2.113,,
963,Minimum Area Rectangle II,48.80%,Medium,34.68%,38.20%,5.70%,2,2.114,,
134,Gas Station,36.60%,Medium,34.63%,38.20%,5.77%,2,2.115,,
428,Serialize and Deserialize N-ary Tree,57.50%,Hard,51.22%,30.80%,-29.51%,3,2.115,,1
34,Find First and Last Position of Element in Sorted Array,35.10%,Medium,34.60%,38.20%,5.82%,2,2.116,,
1117,Building H2O,51.00%,Medium,34.62%,38.20%,5.80%,2,2.116,Concurrency,
19,Remove Nth Node From End of List,34.80%,Medium,34.52%,38.20%,5.95%,2,2.119,,
809,Expressive Words,46.40%,Medium,34.53%,38.20%,5.93%,2,2.119,,
60,Permutation Sequence,35.30%,Medium,34.42%,38.20%,6.12%,2,2.122,,
395,Longest Substring with At Least K Repeating Characters,40.20%,Medium,34.41%,38.20%,6.14%,2,2.123,,
838,Push Dominoes,46.70%,Medium,34.41%,38.20%,6.13%,2,2.123,,
743,Network Delay Time,45.20%,Medium,34.30%,38.20%,6.31%,2,2.126,,
698,Partition to K Equal Sum Subsets,44.50%,Medium,34.26%,38.20%,6.37%,2,2.127,,
785,Is Graph Bipartite?,45.80%,Medium,34.29%,38.20%,6.33%,2,2.127,,
816,Ambiguous Coordinates,46.20%,Medium,34.23%,38.20%,6.42%,2,2.128,,
737,Sentence Similarity II,45.00%,Medium,34.19%,38.20%,6.49%,2,2.130,,1
764,Largest Plus Sign,45.40%,Medium,34.19%,38.20%,6.48%,2,2.130,,
201,Bitwise AND of Numbers Range,37.10%,Medium,34.15%,38.20%,6.55%,2,2.131,,
82,Remove Duplicates from Sorted List II,35.30%,Medium,34.10%,38.20%,6.64%,2,2.133,,
574,Winning Candidate,42.50%,Medium,34.08%,38.20%,6.66%,2,2.133,Sql,1
285,Inorder Successor in BST,38.20%,Medium,34.02%,38.20%,6.76%,2,2.135,,1
375,Guess Number Higher or Lower II,39.50%,Medium,34.00%,38.20%,6.80%,2,2.136,,
681,Next Closest Time,44.00%,Medium,34.01%,38.20%,6.78%,2,2.136,,1
1070,Product Sales Analysis III,49.70%,Medium,34.01%,38.20%,6.79%,2,2.136,Sql,1
438,Find All Anagrams in a String,40.40%,Medium,33.98%,38.20%,6.83%,2,2.137,,
433,Minimum Genetic Mutation,40.30%,Medium,33.95%,38.20%,6.88%,2,2.138,,
223,Rectangle Area,37.10%,Medium,33.83%,38.20%,7.07%,2,2.141,,
1145,Binary Tree Coloring Game,50.60%,Medium,33.81%,38.20%,7.11%,2,2.142,,
1255,Maximum Score Words Formed by Letters,69.00%,Hard,50.59%,30.80%,-28.60%,3,2.142,,
1362,Closest Divisors,53.70%,Medium,33.72%,38.20%,7.24%,2,2.145,,
1284,Minimum Number of Flips to Convert Binary Matrix to Zero Matrix,69.30%,Hard,50.47%,30.80%,-28.42%,3,2.147,,
209,Minimum Size Subarray Sum,36.70%,Medium,33.63%,38.20%,7.39%,2,2.148,,
145,Binary Tree Postorder Traversal,52.50%,Hard,50.37%,30.80%,-28.29%,3,2.151,,
795,Number of Subarrays with Bounded Maximum,45.20%,Medium,33.54%,38.20%,7.54%,2,2.151,,
555,Split Concatenated Strings,41.60%,Medium,33.46%,38.20%,7.67%,2,2.153,,1
376,Wiggle Subsequence,38.90%,Medium,33.39%,38.20%,7.79%,2,2.156,,
417,Pacific Atlantic Water Flow,39.50%,Medium,33.38%,38.20%,7.79%,2,2.156,,
1024,Video Stitching,48.40%,Medium,33.38%,38.20%,7.80%,2,2.156,,
974,Subarray Sums Divisible by K,47.60%,Medium,33.31%,38.20%,7.91%,2,2.158,,
916,Word Subsets,46.70%,Medium,33.27%,38.20%,7.98%,2,2.160,,
1333,"Filter Restaurants by Vegan-Friendly, Price and Distance",52.80%,Medium,33.25%,38.20%,8.01%,2,2.160,,
142,Linked List Cycle II,35.30%,Medium,33.22%,38.20%,8.06%,2,2.161,,
593,Valid Square,41.90%,Medium,33.20%,38.20%,8.09%,2,2.162,,
33,Search in Rotated Sorted Array,33.60%,Medium,33.12%,38.20%,8.23%,2,2.165,,
622,Design Circular Queue,42.20%,Medium,33.08%,38.20%,8.29%,2,2.166,,
659,Split Array into Consecutive Subsequences,42.70%,Medium,33.03%,38.20%,8.36%,2,2.167,,
1341,Movie Rating,52.70%,Medium,33.03%,38.20%,8.36%,2,2.167,Sql,1
63,Unique Paths II,33.90%,Medium,32.98%,38.20%,8.45%,2,2.169,,
1102,Path With Maximum Minimum Value,49.10%,Medium,32.94%,38.20%,8.52%,2,2.170,,1
1291,Sequential Digits,51.80%,Medium,32.87%,38.20%,8.63%,2,2.173,,
2,Add Two Numbers,32.80%,Medium,32.77%,38.20%,8.79%,2,2.176,,
213,House Robber II,35.90%,Medium,32.78%,38.20%,8.78%,2,2.176,,
150,Evaluate Reverse Polish Notation,34.80%,Medium,32.60%,38.20%,9.06%,2,2.181,,
738,Monotone Increasing Digits,43.40%,Medium,32.58%,38.20%,9.10%,2,2.182,,
79,Word Search,33.70%,Medium,32.54%,38.20%,9.16%,2,2.183,,
881,Boats to Save People,45.40%,Medium,32.48%,38.20%,9.26%,2,2.185,,
1053,Previous Permutation With One Swap,47.90%,Medium,32.46%,38.20%,9.29%,2,2.186,,
1155,Number of Dice Rolls With Target Sum,49.40%,Medium,32.46%,38.20%,9.29%,2,2.186,,
55,Jump Game,33.20%,Medium,32.39%,38.20%,9.40%,2,2.188,,
934,Shortest Bridge,46.10%,Medium,32.40%,38.20%,9.38%,2,2.188,,
227,Basic Calculator II,35.70%,Medium,32.37%,38.20%,9.43%,2,2.189,,
18,4Sum,32.60%,Medium,32.34%,38.20%,9.49%,2,2.190,,
93,Restore IP Addresses,33.70%,Medium,32.34%,38.20%,9.49%,2,2.190,,
1041,Robot Bounded In Circle,47.60%,Medium,32.33%,38.20%,9.50%,2,2.190,,
143,Reorder List,34.40%,Medium,32.30%,38.20%,9.54%,2,2.191,,
616,Add Bold Tag in String,41.20%,Medium,32.17%,38.20%,9.76%,2,2.195,,1
221,Maximal Square,35.40%,Medium,32.16%,38.20%,9.78%,2,2.196,,
43,Multiply Strings,32.70%,Medium,32.07%,38.20%,9.92%,2,2.198,,
275,H-Index II,36.10%,Medium,32.07%,38.20%,9.92%,2,2.198,,
54,Spiral Matrix,32.80%,Medium,32.01%,38.20%,10.02%,2,2.200,,
966,Vowel Spellchecker,46.20%,Medium,32.03%,38.20%,9.98%,2,2.200,,
640,Solve the Equation,41.30%,Medium,31.91%,38.20%,10.17%,2,2.203,,
755,Pour Water,43.00%,Medium,31.93%,38.20%,10.15%,2,2.203,,1
567,Permutation in String,40.20%,Medium,31.88%,38.20%,10.22%,2,2.204,,
670,Maximum Swap,41.70%,Medium,31.87%,38.20%,10.24%,2,2.205,,
304,Range Sum Query 2D - Immutable,36.20%,Medium,31.74%,38.20%,10.45%,2,2.209,,
81,Search in Rotated Sorted Array II,32.90%,Medium,31.71%,38.20%,10.50%,2,2.210,,
978,Longest Turbulent Subarray,46.00%,Medium,31.66%,38.20%,10.59%,2,2.212,,
1039,Minimum Score Triangulation of Polygon,46.80%,Medium,31.56%,38.20%,10.74%,2,2.215,,
870,Advantage Shuffle,44.30%,Medium,31.54%,38.20%,10.78%,2,2.216,,
267,Palindrome Permutation II,35.40%,Medium,31.48%,38.20%,10.87%,2,2.217,,1
478,Generate Random Point in a Circle,38.40%,Medium,31.39%,38.20%,11.02%,2,2.220,,
1080,Insufficient Nodes in Root to Leaf Paths,47.20%,Medium,31.36%,38.20%,11.07%,2,2.221,,
1358,Number of Substrings Containing All Three Characters,51.30%,Medium,31.38%,38.20%,11.03%,2,2.221,,
31,Next Permutation,31.80%,Medium,31.35%,38.20%,11.09%,2,2.222,,
393,UTF-8 Validation,37.00%,Medium,31.24%,38.20%,11.27%,2,2.225,,
945,Minimum Increment to Make Array Unique,45.10%,Medium,31.24%,38.20%,11.26%,2,2.225,,
1149,Article Views II,48.10%,Medium,31.25%,38.20%,11.25%,2,2.225,Sql,1
274,H-Index,35.20%,Medium,31.18%,38.20%,11.36%,2,2.227,,
915,Partition Array into Disjoint Intervals,44.60%,Medium,31.18%,38.20%,11.36%,2,2.227,,
1319,Number of Operations to Make Network Connected,50.50%,Medium,31.15%,38.20%,11.40%,2,2.228,,
742,Closest Leaf in a Binary Tree,42.00%,Medium,31.12%,38.20%,11.46%,2,2.229,,1
184,Department Highest Salary,33.80%,Medium,31.10%,38.20%,11.49%,2,2.230,Sql,
211,Add and Search Word - Data structure design,34.10%,Medium,31.01%,38.20%,11.64%,2,2.233,,
848,Shifting Letters,43.40%,Medium,30.96%,38.20%,11.71%,2,2.234,,
831,Masking Personal Information,43.10%,Medium,30.91%,38.20%,11.79%,2,2.236,,
229,Majority Element II,34.20%,Medium,30.84%,38.20%,11.91%,2,2.238,,
1081,Smallest Subsequence of Distinct Characters,46.70%,Medium,30.85%,38.20%,11.90%,2,2.238,,
1156,Swap For Longest Repeated Character Substring,47.80%,Medium,30.85%,38.20%,11.90%,2,2.238,,
1276,Number of Burgers with No Waste of Ingredients,49.50%,Medium,30.79%,38.20%,12.00%,2,2.240,,
372,Super Pow,36.20%,Medium,30.74%,38.20%,12.06%,2,2.241,,
138,Copy List with Random Pointer,32.70%,Medium,30.68%,38.20%,12.17%,2,2.243,,
368,Largest Divisible Subset,35.90%,Medium,30.50%,38.20%,12.46%,2,2.249,,
988,Smallest String Starting From Leaf,45.00%,Medium,30.51%,38.20%,12.44%,2,2.249,,
1093,Statistics from a Large Sample,46.50%,Medium,30.47%,38.20%,12.51%,2,2.250,,
542,01 Matrix,38.40%,Medium,30.45%,38.20%,12.54%,2,2.251,,
634,Find the Derangement of An Array,39.70%,Medium,30.40%,38.20%,12.62%,2,2.252,,1
971,Flip Binary Tree To Match Preorder Traversal,44.50%,Medium,30.26%,38.20%,12.85%,2,2.257,,
396,Rotate Function,35.90%,Medium,30.09%,38.20%,13.12%,2,2.262,,
71,Simplify Path,31.10%,Medium,30.06%,38.20%,13.17%,2,2.263,,
373,Find K Pairs with Smallest Sums,35.50%,Medium,30.03%,38.20%,13.22%,2,2.264,,
991,Broken Calculator,44.50%,Medium,29.97%,38.20%,13.32%,2,2.266,,
473,Matchsticks to Square,36.90%,Medium,29.96%,38.20%,13.33%,2,2.267,,
658,Find K Closest Elements,39.60%,Medium,29.95%,38.20%,13.35%,2,2.267,,
822,Card Flipping Game,42.00%,Medium,29.94%,38.20%,13.36%,2,2.267,,
497,Random Point in Non-overlapping Rectangles,37.20%,Medium,29.91%,38.20%,13.41%,2,2.268,,
545,Boundary of Binary Tree,37.90%,Medium,29.91%,38.20%,13.42%,2,2.268,,1
662,Maximum Width of Binary Tree,39.60%,Medium,29.89%,38.20%,13.45%,2,2.269,,
732,My Calendar III,58.40%,Hard,47.66%,30.80%,-24.37%,3,2.269,,
1366,Rank Teams by Votes,49.90%,Medium,29.87%,38.20%,13.49%,2,2.270,,
161,One Edit Distance,32.20%,Medium,29.84%,38.20%,13.53%,2,2.271,,1
333,Largest BST Subtree,34.70%,Medium,29.82%,38.20%,13.57%,2,2.271,,1
962,Maximum Width Ramp,43.90%,Medium,29.79%,38.20%,13.61%,2,2.272,,
1059,All Paths from Source Lead to Destination,45.20%,Medium,29.67%,38.20%,13.81%,2,2.276,,1
133,Clone Graph,31.60%,Medium,29.65%,38.20%,13.84%,2,2.277,,
775,Global and Local Inversions,41.00%,Medium,29.63%,38.20%,13.86%,2,2.277,,
469,Convex Polygon,36.50%,Medium,29.62%,38.20%,13.88%,2,2.278,,1
935,Knight Dialer,43.30%,Medium,29.59%,38.20%,13.94%,2,2.279,,
3,Longest Substring Without Repeating Characters,29.60%,Medium,29.56%,38.20%,13.99%,2,2.280,,
1138,Alphabet Board Path,46.20%,Medium,29.51%,38.20%,14.06%,2,2.281,,
853,Car Fleet,41.90%,Medium,29.39%,38.20%,14.26%,2,2.285,,
1296,Divide Array in Sets of K Consecutive Numbers,48.30%,Medium,29.29%,38.20%,14.41%,2,2.288,,
735,Asteroid Collision,40.00%,Medium,29.22%,38.20%,14.53%,2,2.291,,
886,Possible Bipartition,42.20%,Medium,29.21%,38.20%,14.55%,2,2.291,,
663,Equal Tree Partition,38.90%,Medium,29.18%,38.20%,14.60%,2,2.292,,1
1107,New Users Daily Count,45.40%,Medium,29.16%,38.20%,14.62%,2,2.292,Sql,1
332,Reconstruct Itinerary,34.00%,Medium,29.13%,38.20%,14.68%,2,2.294,,
1226,The Dining Philosophers,47.10%,Medium,29.12%,38.20%,14.69%,2,2.294,Concurrency,
649,Dota2 Senate,38.60%,Medium,29.08%,38.20%,14.76%,2,2.295,,
1139,Largest 1-Bordered Square,45.80%,Medium,29.09%,38.20%,14.73%,2,2.295,,
825,Friends Of Appropriate Ages,41.10%,Medium,29.00%,38.20%,14.89%,2,2.298,,
322,Coin Change,33.60%,Medium,28.88%,38.20%,15.09%,2,2.302,,
855,Exam Room,41.40%,Medium,28.86%,38.20%,15.11%,2,2.302,,
904,Fruit Into Baskets,42.10%,Medium,28.84%,38.20%,15.14%,2,2.303,,
1034,Coloring A Border,43.90%,Medium,28.73%,38.20%,15.32%,2,2.306,,
5,Longest Palindromic Substring,28.80%,Medium,28.73%,38.20%,15.33%,2,2.307,,
152,Maximum Product Subarray,30.80%,Medium,28.57%,38.20%,15.58%,2,2.312,,
578,Get Highest Answer Rate Question,36.90%,Medium,28.42%,38.20%,15.82%,2,2.316,Sql,1
895,Maximum Frequency Stack,59.70%,Hard,46.57%,30.80%,-22.79%,3,2.316,,
50,"Pow(x, n)",29.10%,Medium,28.37%,38.20%,15.91%,2,2.318,,
1230,Toss Strange Coins,46.30%,Medium,28.26%,38.20%,16.08%,2,2.322,,1
302,Smallest Rectangle Enclosing Black Pixels,50.80%,Hard,46.37%,30.80%,-22.50%,3,2.325,,1
467,Unique Substrings in Wraparound String,34.90%,Medium,28.05%,38.20%,16.42%,2,2.328,,
990,Satisfiability of Equality Equations,42.60%,Medium,28.08%,38.20%,16.38%,2,2.328,,
42,Trapping Rain Water,46.90%,Hard,46.28%,30.80%,-22.38%,3,2.329,,
61,Rotate List,28.90%,Medium,28.01%,38.20%,16.50%,2,2.330,,
713,Subarray Product Less Than K,38.40%,Medium,27.94%,38.20%,16.60%,2,2.332,,
146,LRU Cache,30.00%,Medium,27.86%,38.20%,16.73%,2,2.335,,
1225,Report Contiguous Dates,64.10%,Hard,46.13%,30.80%,-22.16%,3,2.335,Sql,1
1144,Decrease Elements To Make Array Zigzag,44.60%,Medium,27.82%,38.20%,16.79%,2,2.336,,
307,Range Sum Query - Mutable,32.30%,Medium,27.80%,38.20%,16.83%,2,2.337,,
519,Random Flip Matrix,35.40%,Medium,27.79%,38.20%,16.85%,2,2.337,,
1274,Number of Ships in a Rectangle,64.70%,Hard,46.01%,30.80%,-21.99%,3,2.340,,1
1229,Meeting Scheduler,45.70%,Medium,27.67%,38.20%,17.03%,2,2.341,,1
353,Design Snake Game,32.80%,Medium,27.62%,38.20%,17.12%,2,2.342,,1
1049,Last Stone Weight II,42.90%,Medium,27.51%,38.20%,17.29%,2,2.346,,
385,Mini Parser,33.00%,Medium,27.35%,38.20%,17.55%,2,2.351,,
773,Sliding Puzzle,57.10%,Hard,45.76%,30.80%,-21.62%,3,2.351,,
930,Binary Subarrays With Sum,40.90%,Medium,27.26%,38.20%,17.70%,2,2.354,,
1239,Maximum Length of a Concatenated String with Unique Characters,45.30%,Medium,27.13%,38.20%,17.92%,2,2.358,,
177,Nth Highest Salary,29.70%,Medium,27.10%,38.20%,17.95%,2,2.359,Sql,
569,Median Employee Salary,53.90%,Hard,45.55%,30.80%,-21.32%,3,2.360,Sql,1
808,Soup Servings,38.90%,Medium,27.05%,38.20%,18.04%,2,2.361,,
1096,Brace Expansion II,61.60%,Hard,45.53%,30.80%,-21.28%,3,2.362,,
312,Burst Balloons,50.00%,Hard,45.42%,30.80%,-21.13%,3,2.366,,
1098,Unpopular Books,43.00%,Medium,26.90%,38.20%,18.29%,2,2.366,Sql,1
310,Minimum Height Trees,31.40%,Medium,26.85%,38.20%,18.36%,2,2.367,,
1181,Before and After Puzzle,44.10%,Medium,26.78%,38.20%,18.48%,2,2.370,,1
1300,Sum of Mutated Array Closest to Target,45.80%,Medium,26.73%,38.20%,18.55%,2,2.371,,
1205,Monthly Transactions II,44.30%,Medium,26.63%,38.20%,18.73%,2,2.375,Sql,1
527,Word Abbreviation,52.90%,Hard,45.17%,30.80%,-20.77%,3,2.377,,1
397,Integer Replacement,32.30%,Medium,26.48%,38.20%,18.97%,2,2.379,,
1292,Maximum Side Length of a Square with Sum Less than or Equal to Threshold,45.40%,Medium,26.45%,38.20%,19.01%,2,2.380,,
790,Domino and Tromino Tiling,37.90%,Medium,26.31%,38.20%,19.23%,2,2.385,,
948,Bag of Tokens,40.20%,Medium,26.30%,38.20%,19.26%,2,2.385,,
1202,Smallest String With Swaps,43.90%,Medium,26.27%,38.20%,19.30%,2,2.386,,
801,Minimum Swaps To Make Sequences Increasing,38.00%,Medium,26.25%,38.20%,19.33%,2,2.387,,
356,Line Reflection,31.40%,Medium,26.18%,38.20%,19.45%,2,2.389,,1
271,Encode and Decode Strings,30.10%,Medium,26.13%,38.20%,19.54%,2,2.391,,1
1058,Minimize Rounding Error to Meet Target,41.60%,Medium,26.08%,38.20%,19.61%,2,2.392,,1
272,Closest Binary Search Tree Value II,48.80%,Hard,44.81%,30.80%,-20.25%,3,2.393,,1
1074,Number of Submatrices That Sum to Target,60.50%,Hard,44.75%,30.80%,-20.16%,3,2.395,,
418,Sentence Screen Fitting,32.10%,Medium,25.97%,38.20%,19.79%,2,2.396,,1
522,Longest Uncommon Subsequence II,33.50%,Medium,25.84%,38.20%,19.99%,2,2.400,,
779,K-th Symbol in Grammar,37.20%,Medium,25.77%,38.20%,20.11%,2,2.402,,
127,Word Ladder,27.60%,Medium,25.74%,38.20%,20.17%,2,2.403,,
1223,Dice Roll Simulation,43.70%,Medium,25.76%,38.20%,20.13%,2,2.403,,
98,Validate Binary Search Tree,27.10%,Medium,25.66%,38.20%,20.29%,2,2.406,,
787,Cheapest Flights Within K Stops,37.20%,Medium,25.66%,38.20%,20.30%,2,2.406,,
1152,Analyze User Website Visit Pattern,42.50%,Medium,25.60%,38.20%,20.38%,2,2.408,,1
1242,Web Crawler Multithreaded,43.80%,Medium,25.58%,38.20%,20.41%,2,2.408,Concurrency,1
826,Most Profit Assigning Work,37.60%,Medium,25.49%,38.20%,20.57%,2,2.411,,
15,3Sum,25.70%,Medium,25.48%,38.20%,20.58%,2,2.412,,
576,Out of Boundary Paths,33.80%,Medium,25.35%,38.20%,20.79%,2,2.416,,
1297,Maximum Number of Occurrences of a Substring,44.30%,Medium,25.28%,38.20%,20.91%,2,2.418,,
1262,Greatest Sum Divisible by Three,43.70%,Medium,25.19%,38.20%,21.05%,2,2.421,,
957,Prison Cells After N Days,39.20%,Medium,25.16%,38.20%,21.09%,2,2.422,,
400,Nth Digit,31.00%,Medium,25.13%,38.20%,21.14%,2,2.423,,
1054,Distant Barcodes,40.60%,Medium,25.14%,38.20%,21.13%,2,2.423,,
673,Number of Longest Increasing Subsequence,34.90%,Medium,25.03%,38.20%,21.31%,2,2.426,,
179,Largest Number,27.50%,Medium,24.87%,38.20%,21.56%,2,2.431,,
1162,As Far from Land as Possible,41.70%,Medium,24.66%,38.20%,21.91%,2,2.438,,
365,Water and Jug Problem,30.00%,Medium,24.65%,38.20%,21.93%,2,2.439,,
306,Additive Number,29.00%,Medium,24.51%,38.20%,22.15%,2,2.443,,
1171,Remove Zero Sum Consecutive Nodes from Linked List,41.70%,Medium,24.53%,38.20%,22.13%,2,2.443,,
324,Wiggle Sort II,29.20%,Medium,24.45%,38.20%,22.25%,2,2.445,,
967,Numbers With Same Consecutive Differences,38.60%,Medium,24.42%,38.20%,22.30%,2,2.446,,
660,Remove 9,53.20%,Hard,43.52%,30.80%,-18.38%,3,2.449,,1
1218,Longest Arithmetic Subsequence of Given Difference,42.20%,Medium,24.34%,38.20%,22.43%,2,2.449,,
1136,Parallel Courses,60.10%,Hard,43.44%,30.80%,-18.26%,3,2.452,,1
355,Design Twitter,29.10%,Medium,23.89%,38.20%,23.15%,2,2.463,,
678,Valid Parenthesis String,33.80%,Medium,23.86%,38.20%,23.21%,2,2.464,,
909,Snakes and Ladders,37.10%,Medium,23.77%,38.20%,23.35%,2,2.467,,
51,N-Queens,43.80%,Hard,43.05%,30.80%,-17.71%,3,2.469,,
192,Word Frequency,26.40%,Medium,23.58%,38.20%,23.65%,2,2.473,Shell,
130,Surrounded Regions,25.40%,Medium,23.49%,38.20%,23.80%,2,2.476,,
165,Compare Version Numbers,25.90%,Medium,23.48%,38.20%,23.82%,2,2.476,,
845,Longest Mountain in Array,35.90%,Medium,23.51%,38.20%,23.78%,2,2.476,,
625,Minimum Factorization,32.60%,Medium,23.43%,38.20%,23.89%,2,2.478,,1
761,Special Binary String,54.00%,Hard,42.84%,30.80%,-17.40%,3,2.478,,
842,Split Array into Fibonacci Sequence,35.70%,Medium,23.35%,38.20%,24.03%,2,2.481,,
898,Bitwise ORs of Subarrays,36.50%,Medium,23.33%,38.20%,24.06%,2,2.481,,
777,Swap Adjacent in LR String,34.70%,Medium,23.30%,38.20%,24.10%,2,2.482,,
799,Champagne Tower,35.00%,Medium,23.28%,38.20%,24.14%,2,2.483,,
754,Reach a Number,34.10%,Medium,23.04%,38.20%,24.53%,2,2.491,,
1311,Get Watched Videos by Your Friends,42.20%,Medium,22.97%,38.20%,24.64%,2,2.493,,
1168,Optimize Water Distribution in a Village,59.60%,Hard,42.47%,30.80%,-16.86%,3,2.494,,1
556,Next Greater Element III,31.00%,Medium,22.85%,38.20%,24.85%,2,2.497,,
823,Binary Trees With Factors,34.80%,Medium,22.73%,38.20%,25.03%,2,2.501,,
765,Couples Holding Hands,53.50%,Hard,42.28%,30.80%,-16.59%,3,2.502,,
722,Remove Comments,33.20%,Medium,22.61%,38.20%,25.23%,2,2.505,,
91,Decode Ways,23.70%,Medium,22.37%,38.20%,25.62%,2,2.512,,
1283,Find the Smallest Divisor Given a Threshold,41.20%,Medium,22.38%,38.20%,25.59%,2,2.512,,
128,Longest Consecutive Sequence,43.90%,Hard,42.02%,30.80%,-16.22%,3,2.513,,
1106,Parsing A Boolean Expression,58.10%,Hard,41.88%,30.80%,-16.01%,3,2.520,,
456,132 Pattern,28.70%,Medium,22.01%,38.20%,26.19%,2,2.524,,
1334,Find the City With the Smallest Number of Neighbors at a Threshold Distance,41.50%,Medium,21.93%,38.20%,26.32%,2,2.526,,
457,Circular Array Loop,28.50%,Medium,21.80%,38.20%,26.54%,2,2.531,,
464,Can I Win,28.50%,Medium,21.69%,38.20%,26.71%,2,2.534,,
923,3Sum With Multiplicity,35.20%,Medium,21.66%,38.20%,26.76%,2,2.535,,
402,Remove K Digits,27.50%,Medium,21.60%,38.20%,26.85%,2,2.537,,
954,Array of Doubled Pairs,35.60%,Medium,21.61%,38.20%,26.85%,2,2.537,,
984,String Without AAA or BBB,36.00%,Medium,21.57%,38.20%,26.91%,2,2.538,,
632,Smallest Range Covering Elements from K Lists,50.70%,Hard,41.43%,30.80%,-15.36%,3,2.539,,
163,Missing Ranges,23.90%,Medium,21.51%,38.20%,27.01%,2,2.540,,1
1328,Break a Palindrome,41.00%,Medium,21.52%,38.20%,26.99%,2,2.540,,
1129,Shortest Path with Alternating Colors,38.00%,Medium,21.44%,38.20%,27.12%,2,2.542,,
837,New 21 Game,33.70%,Medium,21.42%,38.20%,27.15%,2,2.543,,
1091,Shortest Path in Binary Matrix,37.30%,Medium,21.30%,38.20%,27.35%,2,2.547,,
1147,Longest Chunked Palindrome Decomposition,58.00%,Hard,41.18%,30.80%,-15.00%,3,2.550,,
982,Triples with Bitwise AND Equal To Zero,55.50%,Hard,41.10%,30.80%,-14.88%,3,2.554,,
194,Transpose File,23.90%,Medium,21.05%,38.20%,27.74%,2,2.555,Shell,
1208,Get Equal Substrings Within Budget,38.60%,Medium,20.88%,38.20%,28.02%,2,2.560,,
297,Serialize and Deserialize Binary Tree,45.20%,Hard,40.84%,30.80%,-14.51%,3,2.565,,
708,Insert into a Sorted Circular Linked List,31.10%,Medium,20.72%,38.20%,28.29%,2,2.566,,1
1253,Reconstruct a 2-Row Binary Matrix,39.10%,Medium,20.72%,38.20%,28.28%,2,2.566,,
794,Valid Tic-Tac-Toe State,32.20%,Medium,20.55%,38.20%,28.55%,2,2.571,,
425,Word Squares,46.90%,Hard,40.67%,30.80%,-14.26%,3,2.572,,1
918,Maximum Sum Circular Subarray,33.90%,Medium,20.44%,38.20%,28.74%,2,2.575,,
37,Sudoku Solver,41.00%,Hard,40.46%,30.80%,-13.96%,3,2.581,,
1215,Stepping Numbers,38.00%,Medium,20.18%,38.20%,29.16%,2,2.583,,1
72,Edit Distance,41.30%,Hard,40.24%,30.80%,-13.65%,3,2.591,,
352,Data Stream as Disjoint Intervals,45.40%,Hard,40.24%,30.80%,-13.64%,3,2.591,,
987,Vertical Order Traversal of a Binary Tree,34.10%,Medium,19.62%,38.20%,30.06%,2,2.601,,
458,Poor Pigs,46.70%,Hard,39.98%,30.80%,-13.27%,3,2.602,,
265,Paint House II,43.80%,Hard,39.91%,30.80%,-13.17%,3,2.605,,1
1289,Minimum Falling Path Sum II,58.80%,Hard,39.89%,30.80%,-13.14%,3,2.606,,
618,Students Report By Geography,48.70%,Hard,39.64%,30.80%,-12.77%,3,2.617,Sql,1
1146,Snapshot Array,35.90%,Medium,19.09%,38.20%,30.92%,2,2.618,,
1352,Product of the Last K Numbers,38.90%,Medium,19.07%,38.20%,30.95%,2,2.619,,
960,Delete Columns to Make Sorted III,53.60%,Hard,39.52%,30.80%,-12.60%,3,2.622,,
778,Swim in Rising Water,50.80%,Hard,39.39%,30.80%,-12.41%,3,2.628,,
955,Delete Columns to Make Sorted II,32.80%,Medium,18.79%,38.20%,31.40%,2,2.628,,
1206,Design Skiplist,57.00%,Hard,39.31%,30.80%,-12.30%,3,2.631,,
471,Encode String with Shortest Length,46.20%,Hard,39.29%,30.80%,-12.27%,3,2.632,,1
1121,Divide Array Into Increasing Sequences,55.70%,Hard,39.26%,30.80%,-12.22%,3,2.633,,1
25,Reverse Nodes in k-Group,39.60%,Hard,39.23%,30.80%,-12.19%,3,2.634,,
614,Second Degree Follower,27.60%,Medium,18.59%,38.20%,31.72%,2,2.634,Sql,1
166,Fraction to Recurring Decimal,20.80%,Medium,18.37%,38.20%,32.09%,2,2.642,,
1278,Palindrome Partitioning III,57.80%,Hard,39.06%,30.80%,-11.93%,3,2.642,,
1298,Maximum Candies You Can Get from Boxes,58.00%,Hard,38.96%,30.80%,-11.80%,3,2.646,,
1367,Linked List in Binary Tree,38.20%,Medium,18.15%,38.20%,32.44%,2,2.649,,
1186,Maximum Subarray Sum with One Deletion,35.50%,Medium,18.11%,38.20%,32.52%,2,2.650,,
1097,Game Play Analysis V,54.80%,Hard,38.71%,30.80%,-11.43%,3,2.657,Sql,1
465,Optimal Account Balancing,45.50%,Hard,38.68%,30.80%,-11.39%,3,2.658,,1
1132,Reported Posts II,34.30%,Medium,17.70%,38.20%,33.18%,2,2.664,Sql,1
1073,Adding Two Negabinary Numbers,33.40%,Medium,17.66%,38.20%,33.23%,2,2.665,,
1359,Count All Valid Pickup and Delivery Options,58.40%,Hard,38.47%,30.80%,-11.08%,3,2.668,,
291,Word Pattern II,42.70%,Hard,38.43%,30.80%,-11.03%,3,2.669,,1
753,Cracking the Safe,49.40%,Hard,38.36%,30.80%,-10.92%,3,2.672,,
907,Sum of Subarray Minimums,30.60%,Medium,17.30%,38.20%,33.82%,2,2.676,,
220,Contains Duplicate III,20.50%,Medium,17.27%,38.20%,33.86%,2,2.677,,
151,Reverse Words in a String,19.40%,Medium,17.19%,38.20%,34.00%,2,2.680,,
847,Shortest Path Visiting All Nodes,50.60%,Hard,38.18%,30.80%,-10.66%,3,2.680,,
154,Find Minimum in Rotated Sorted Array II,40.30%,Hard,38.04%,30.80%,-10.46%,3,2.686,,
288,Unique Word Abbreviation,21.20%,Medium,16.98%,38.20%,34.34%,2,2.687,,1
23,Merge k Sorted Lists,38.30%,Hard,37.96%,30.80%,-10.35%,3,2.689,,
711,Number of Distinct Islands II,48.30%,Hard,37.87%,30.80%,-10.22%,3,2.693,,1
523,Continuous Subarray Sum,24.40%,Medium,16.73%,38.20%,34.74%,2,2.695,,
340,Longest Substring with At Most K Distinct Characters,42.70%,Hard,37.71%,30.80%,-9.99%,3,2.700,,1
410,Split Array Largest Sum,43.60%,Hard,37.59%,30.80%,-9.81%,3,2.706,,
1015,Smallest Integer Divisible by K,31.10%,Medium,16.21%,38.20%,35.58%,2,2.712,,
899,Orderly Queue,50.60%,Hard,37.41%,30.80%,-9.56%,3,2.713,,
329,Longest Increasing Path in a Matrix,42.20%,Hard,37.37%,30.80%,-9.50%,3,2.715,,
239,Sliding Window Maximum,40.80%,Hard,37.29%,30.80%,-9.39%,3,2.718,,
301,Remove Invalid Parentheses,41.70%,Hard,37.29%,30.80%,-9.37%,3,2.719,,
1197,Minimum Knight Moves,33.50%,Medium,15.94%,38.20%,36.01%,2,2.720,,1
29,Divide Two Integers,16.20%,Medium,15.77%,38.20%,36.29%,2,2.726,,
1177,Can Make Palindrome from Substring,33.00%,Medium,15.74%,38.20%,36.35%,2,2.727,,
1124,Longest Well-Performing Interval,32.20%,Medium,15.71%,38.20%,36.38%,2,2.728,,
1339,Maximum Product of Splitted Binary Tree,35.30%,Medium,15.66%,38.20%,36.47%,2,2.729,,
295,Find Median from Data Stream,41.30%,Hard,36.97%,30.80%,-8.92%,3,2.732,,
1320,Minimum Distance to Type a Word Using Two Fingers,56.10%,Hard,36.74%,30.80%,-8.58%,3,2.742,,
468,Validate IP Address,22.10%,Medium,15.24%,38.20%,37.16%,2,2.743,,
768,Max Chunks To Make Sorted II,47.80%,Hard,36.54%,30.80%,-8.29%,3,2.751,,
444,Sequence Reconstruction,21.40%,Medium,14.89%,38.20%,37.72%,2,2.754,,1
8,String to Integer (atoi),15.00%,Medium,14.88%,38.20%,37.73%,2,2.755,,
726,Number of Atoms,47.10%,Hard,36.45%,30.80%,-8.17%,3,2.755,,
857,Minimum Cost to Hire K Workers,49.00%,Hard,36.43%,30.80%,-8.14%,3,2.756,,
571,Find Median Given Frequency of Numbers,44.60%,Hard,36.23%,30.80%,-7.84%,3,2.765,Sql,1
315,Count of Smaller Numbers After Self,40.70%,Hard,36.08%,30.80%,-7.63%,3,2.771,,
99,Recover Binary Search Tree,37.50%,Hard,36.05%,30.80%,-7.58%,3,2.772,,
1340,Jump Game V,55.70%,Hard,36.05%,30.80%,-7.58%,3,2.773,,
770,Basic Calculator IV,47.30%,Hard,36.01%,30.80%,-7.52%,3,2.774,,
1312,Minimum Insertion Steps to Make a String Palindrome,55.20%,Hard,35.96%,30.80%,-7.45%,3,2.776,,
305,Number of Islands II,40.40%,Hard,35.93%,30.80%,-7.41%,3,2.778,,1
1194,Tournament Winners,53.20%,Hard,35.69%,30.80%,-7.06%,3,2.788,Sql,1
317,Shortest Distance from All Buildings,40.10%,Hard,35.45%,30.80%,-6.72%,3,2.798,,1
736,Parse Lisp Expression,46.10%,Hard,35.31%,30.80%,-6.51%,3,2.805,,
1335,Minimum Difficulty of a Job Schedule,54.80%,Hard,35.22%,30.80%,-6.39%,3,2.808,,
1234,Replace the Substring for Balanced String,31.30%,Medium,13.20%,38.20%,40.45%,2,2.809,,
115,Distinct Subsequences,36.80%,Hard,35.11%,30.80%,-6.23%,3,2.813,,
810,Chalkboard XOR Game,47.00%,Hard,35.12%,30.80%,-6.24%,3,2.813,,
679,24 Game,45.00%,Hard,35.04%,30.80%,-6.13%,3,2.816,,
1183,Maximum Number of Ones,52.40%,Hard,35.05%,30.80%,-6.14%,3,2.816,,1
689,Maximum Sum of 3 Non-Overlapping Subarrays,45.10%,Hard,34.99%,30.80%,-6.06%,3,2.818,,
248,Strobogrammatic Number III,38.60%,Hard,34.96%,30.80%,-6.02%,3,2.820,,1
903,Valid Permutations for DI Sequence,48.20%,Hard,34.96%,30.80%,-6.01%,3,2.820,,
1169,Invalid Transactions,30.00%,Medium,12.85%,38.20%,41.01%,2,2.820,,
85,Maximal Rectangle,36.00%,Hard,34.75%,30.80%,-5.71%,3,2.829,,
407,Trapping Rain Water II,40.70%,Hard,34.73%,30.80%,-5.68%,3,2.830,,
1220,Count Vowels Permutation,52.60%,Hard,34.71%,30.80%,-5.65%,3,2.831,,
588,Design In-Memory File System,43.30%,Hard,34.68%,30.80%,-5.60%,3,2.832,,1
910,Smallest Range II,25.80%,Medium,12.45%,38.20%,41.66%,2,2.833,,
668,Kth Smallest Number in Multiplication Table,44.40%,Hard,34.60%,30.80%,-5.50%,3,2.835,,
514,Freedom Trail,41.90%,Hard,34.36%,30.80%,-5.15%,3,2.846,,
1092,Shortest Common Supersequence,50.30%,Hard,34.28%,30.80%,-5.03%,3,2.849,,
1259,Handshakes That Don't Cross,52.70%,Hard,34.23%,30.80%,-4.96%,3,2.851,,
123,Best Time to Buy and Sell Stock III,36.00%,Hard,34.20%,30.80%,-4.91%,3,2.853,,
707,Design Linked List,22.20%,Medium,11.83%,38.20%,42.67%,2,2.853,,
1353,Maximum Number of Events That Can Be Attended,31.50%,Medium,11.66%,38.20%,42.95%,2,2.859,,
774,Minimize Max Distance to Gas Station,45.30%,Hard,33.95%,30.80%,-4.55%,3,2.864,,1
850,Rectangle Area II,46.40%,Hard,33.93%,30.80%,-4.53%,3,2.864,,
615,Average Salary: Departments VS Company,42.90%,Hard,33.88%,30.80%,-4.45%,3,2.866,Sql,1
1250,Check If It Is a Good Array,52.10%,Hard,33.77%,30.80%,-4.29%,3,2.871,,
866,Prime Palindrome,23.90%,Medium,11.20%,38.20%,43.69%,2,2.874,,
1159,Market Analysis II,50.60%,Hard,33.60%,30.80%,-4.05%,3,2.879,Sql,1
472,Concatenated Words,40.50%,Hard,33.58%,30.80%,-4.01%,3,2.880,,
880,Decoded String at Index,23.80%,Medium,10.89%,38.20%,44.19%,2,2.884,,
488,Zuma Game,40.30%,Hard,33.14%,30.80%,-3.39%,3,2.898,,
642,Design Search Autocomplete System,42.50%,Hard,33.08%,30.80%,-3.30%,3,2.901,,1
843,Guess the Word,45.40%,Hard,33.04%,30.80%,-3.23%,3,2.903,,
996,Number of Squareful Arrays,47.60%,Hard,32.99%,30.80%,-3.17%,3,2.905,,
1231,Divide Chocolate,50.70%,Hard,32.65%,30.80%,-2.67%,3,2.920,,1
403,Frog Jump,38.40%,Hard,32.49%,30.80%,-2.44%,3,2.927,,
499,The Maze III,39.80%,Hard,32.48%,30.80%,-2.43%,3,2.927,,1
546,Remove Boxes,40.50%,Hard,32.49%,30.80%,-2.45%,3,2.927,,
827,Making A Large Island,44.60%,Hard,32.47%,30.80%,-2.41%,3,2.928,,
76,Minimum Window Substring,33.40%,Hard,32.29%,30.80%,-2.15%,3,2.936,,
84,Largest Rectangle in Histogram,33.50%,Hard,32.27%,30.80%,-2.12%,3,2.936,,
224,Basic Calculator,35.50%,Hard,32.21%,30.80%,-2.04%,3,2.939,,
749,Contain Virus,43.10%,Hard,32.11%,30.80%,-1.90%,3,2.943,,
164,Maximum Gap,34.40%,Hard,31.99%,30.80%,-1.73%,3,2.948,,
920,Number of Music Playlists,45.40%,Hard,31.91%,30.80%,-1.60%,3,2.952,,
992,Subarrays with K Different Integers,46.40%,Hard,31.85%,30.80%,-1.52%,3,2.954,,
57,Insert Interval,32.60%,Hard,31.76%,30.80%,-1.39%,3,2.958,,
995,Minimum Number of K Consecutive Bit Flips,46.30%,Hard,31.71%,30.80%,-1.31%,3,2.961,,
502,IPO,39.00%,Hard,31.64%,30.80%,-1.21%,3,2.964,,
1191,K-Concatenation Maximum Sum,25.80%,Medium,8.33%,38.20%,48.33%,2,2.967,,
87,Scramble String,32.80%,Hard,31.52%,30.80%,-1.05%,3,2.969,,
798,Smallest Rotation with Highest Score,43.10%,Hard,31.40%,30.80%,-0.86%,3,2.974,,
691,Stickers to Spell Word,41.50%,Hard,31.37%,30.80%,-0.82%,3,2.975,,
568,Maximum Vacation Days,39.50%,Hard,31.17%,30.80%,-0.53%,3,2.984,,1
828,Count Unique Characters of All Substrings of a Given String,43.30%,Hard,31.16%,30.80%,-0.51%,3,2.985,,
1192,Critical Connections in a Network,48.50%,Hard,31.02%,30.80%,-0.31%,3,2.991,,
363,Max Sum of Rectangle No Larger Than K,36.30%,Hard,30.98%,30.80%,-0.25%,3,2.992,,
601,Human Traffic of Stadium,39.80%,Hard,30.99%,30.80%,-0.27%,3,2.992,Sql,
1032,Stream of Characters,46.10%,Hard,30.96%,30.80%,-0.24%,3,2.993,,
1127,User Purchase Platform,47.50%,Hard,30.97%,30.80%,-0.25%,3,2.993,Sql,1
1240,Tiling a Rectangle with the Fewest Squares,49.10%,Hard,30.91%,30.80%,-0.16%,3,2.995,,
1201,Ugly Number III,25.00%,Medium,7.39%,38.20%,49.86%,2,2.997,,
699,Falling Squares,41.10%,Hard,30.85%,30.80%,-0.07%,3,2.998,,
282,Expression Add Operators,34.70%,Hard,30.56%,30.80%,0.34%,3,3.010,,
730,Count Different Palindromic Subsequences,41.20%,Hard,30.49%,30.80%,0.44%,3,3.013,,
124,Binary Tree Maximum Path Sum,32.30%,Hard,30.48%,30.80%,0.46%,3,3.014,,
786,K-th Smallest Prime Fraction,41.80%,Hard,30.27%,30.80%,0.76%,3,3.023,,
517,Super Washing Machines,37.80%,Hard,30.22%,30.80%,0.84%,3,3.025,,
41,First Missing Positive,30.80%,Hard,30.20%,30.80%,0.87%,3,3.026,,
218,The Skyline Problem,33.40%,Hard,30.20%,30.80%,0.86%,3,3.026,,
269,Alien Dictionary,34.00%,Hard,30.05%,30.80%,1.08%,3,3.032,,1
727,Minimum Window Subsequence,40.70%,Hard,30.04%,30.80%,1.10%,3,3.033,,1
834,Sum of Distances in Tree,42.20%,Hard,29.97%,30.80%,1.20%,3,3.036,,
1348,Tweet Counts Per Frequency,25.80%,Medium,6.03%,38.20%,52.06%,2,3.041,,
308,Range Sum Query 2D - Mutable,34.30%,Hard,29.78%,30.80%,1.47%,3,3.044,,1
815,Bus Routes,41.70%,Hard,29.75%,30.80%,1.52%,3,3.046,,
316,Remove Duplicate Letters,34.30%,Hard,29.67%,30.80%,1.64%,3,3.049,,
411,Minimum Unique Word Abbreviation,35.70%,Hard,29.67%,30.80%,1.63%,3,3.049,,1
327,Count of Range Sum,34.40%,Hard,29.60%,30.80%,1.73%,3,3.052,,
354,Russian Doll Envelopes,34.80%,Hard,29.61%,30.80%,1.72%,3,3.052,,
782,Transform to Chessboard,41.00%,Hard,29.53%,30.80%,1.83%,3,3.055,,
1368,Minimum Cost to Make at Least One Valid Path in a Grid,49.60%,Hard,29.54%,30.80%,1.83%,3,3.055,,
772,Basic Calculator III,40.80%,Hard,29.48%,30.80%,1.91%,3,3.057,,1
975,Odd Even Jump,43.70%,Hard,29.40%,30.80%,2.02%,3,3.061,,
330,Patching Array,34.10%,Hard,29.26%,30.80%,2.23%,3,3.067,,
664,Strange Printer,39.00%,Hard,29.26%,30.80%,2.22%,3,3.067,,
964,Least Operators to Express Number,43.40%,Hard,29.26%,30.80%,2.22%,3,3.067,,
45,Jump Game II,29.70%,Hard,29.04%,30.80%,2.54%,3,3.076,,
212,Word Search II,32.10%,Hard,28.99%,30.80%,2.61%,3,3.078,,
1125,Smallest Sufficient Team,45.50%,Hard,29.00%,30.80%,2.60%,3,3.078,,
358,Rearrange String k Distance Apart,34.20%,Hard,28.95%,30.80%,2.67%,3,3.080,,1
97,Interleaving String,30.10%,Hard,28.68%,30.80%,3.07%,3,3.092,,
185,Department Top Three Salaries,31.40%,Hard,28.69%,30.80%,3.05%,3,3.092,Sql,
1203,Sort Items by Groups Respecting Dependencies,46.10%,Hard,28.46%,30.80%,3.39%,3,3.102,,
158,Read N Characters Given Read4 II - Call multiple times,30.70%,Hard,28.38%,30.80%,3.49%,3,3.105,,1
4,Median of Two Sorted Arrays,28.40%,Hard,28.34%,30.80%,3.55%,3,3.107,,
135,Candy,30.30%,Hard,28.32%,30.80%,3.58%,3,3.108,,
480,Sliding Window Median,35.30%,Hard,28.26%,30.80%,3.67%,3,3.110,,
793,Preimage Size of Factorial Zeroes Function,39.90%,Hard,28.27%,30.80%,3.66%,3,3.110,,
1216,Valid Palindrome III,46.00%,Hard,28.17%,30.80%,3.81%,3,3.114,,1
483,Smallest Good Base,35.20%,Hard,28.12%,30.80%,3.88%,3,3.116,,
924,Minimize Malware Spread,41.60%,Hard,28.05%,30.80%,3.98%,3,3.119,,
336,Palindrome Pairs,32.80%,Hard,27.87%,30.80%,4.23%,3,3.127,,
943,Find the Shortest Superstring,41.70%,Hard,27.87%,30.80%,4.24%,3,3.127,,
381,Insert Delete GetRandom O(1) - Duplicates allowed,33.40%,Hard,27.81%,30.80%,4.32%,3,3.130,,
140,Word Break II,29.80%,Hard,27.75%,30.80%,4.41%,3,3.132,,
233,Number of Digit One,31.00%,Hard,27.58%,30.80%,4.65%,3,3.139,,
940,Distinct Subsequences II,41.30%,Hard,27.51%,30.80%,4.75%,3,3.142,,
757,Set Intersection Size At Least Two,38.50%,Hard,27.40%,30.80%,4.92%,3,3.148,,
132,Palindrome Partitioning II,29.30%,Hard,27.36%,30.80%,4.97%,3,3.149,,
715,Range Module,37.80%,Hard,27.31%,30.80%,5.04%,3,3.151,,
552,Student Attendance Record II,35.30%,Hard,27.20%,30.80%,5.20%,3,3.156,,
579,Find Cumulative Salary of an Employee,35.70%,Hard,27.21%,30.80%,5.19%,3,3.156,Sql,1
972,Equal Rational Numbers,41.40%,Hard,27.14%,30.80%,5.28%,3,3.158,,
882,Reachable Nodes In Subdivided Graph,40.00%,Hard,27.06%,30.80%,5.40%,3,3.162,,
32,Longest Valid Parentheses,27.40%,Hard,26.93%,30.80%,5.59%,3,3.168,,
864,Shortest Path to Get All Keys,39.30%,Hard,26.63%,30.80%,6.03%,3,3.181,,
1316,Distinct Echo Substrings,45.90%,Hard,26.60%,30.80%,6.07%,3,3.182,,
1210,Minimum Moves to Reach Target with Rotations,44.30%,Hard,26.55%,30.80%,6.14%,3,3.184,,
1246,Palindrome Removal,44.80%,Hard,26.53%,30.80%,6.18%,3,3.185,,1
587,Erect the Fence,35.10%,Hard,26.49%,30.80%,6.23%,3,3.187,,
928,Minimize Malware Spread II,40.10%,Hard,26.49%,30.80%,6.23%,3,3.187,,
174,Dungeon Game,29.00%,Hard,26.45%,30.80%,6.29%,3,3.189,,
1336,Number of Transactions per Visit,45.70%,Hard,26.11%,30.80%,6.78%,3,3.204,Sql,1
10,Regular Expression Matching,26.20%,Hard,26.05%,30.80%,6.86%,3,3.206,,
214,Shortest Palindrome,29.00%,Hard,25.86%,30.80%,7.14%,3,3.214,,
1235,Maximum Profit in Job Scheduling,43.80%,Hard,25.69%,30.80%,7.39%,3,3.222,,
262,Trips and Users,29.50%,Hard,25.66%,30.80%,7.43%,3,3.223,Sql,
460,LFU Cache,32.40%,Hard,25.65%,30.80%,7.44%,3,3.223,,
879,Profitable Schemes,38.50%,Hard,25.61%,30.80%,7.50%,3,3.225,,
818,Race Car,37.50%,Hard,25.50%,30.80%,7.66%,3,3.230,,
446,Arithmetic Slices II - Subsequence,31.90%,Hard,25.36%,30.80%,7.86%,3,3.236,,
683,K Empty Slots,35.20%,Hard,25.18%,30.80%,8.12%,3,3.244,,1
591,Tag Validator,33.80%,Hard,25.13%,30.80%,8.19%,3,3.246,,
68,Text Justification,26.10%,Hard,25.10%,30.80%,8.23%,3,3.247,,
956,Tallest Billboard,39.10%,Hard,25.08%,30.80%,8.27%,3,3.248,,
432,All O'one Data Structure,31.40%,Hard,25.06%,30.80%,8.29%,3,3.249,,
854,K-Similar Strings,37.40%,Hard,24.87%,30.80%,8.56%,3,3.257,,
188,Best Time to Buy and Sell Stock IV,27.50%,Hard,24.74%,30.80%,8.75%,3,3.263,,
1088,Confusing Number II,40.70%,Hard,24.74%,30.80%,8.75%,3,3.263,,1
600,Non-negative Integers without Consecutive Ones,33.50%,Hard,24.70%,30.80%,8.82%,3,3.264,,
839,Similar String Groups,36.90%,Hard,24.59%,30.80%,8.97%,3,3.269,,
30,Substring with Concatenation of All Words,24.80%,Hard,24.36%,30.80%,9.31%,3,3.279,,
1000,Minimum Cost to Merge Stones,38.70%,Hard,24.03%,30.80%,9.78%,3,3.293,,
829,Consecutive Numbers Sum,36.10%,Hard,23.94%,30.80%,9.91%,3,3.297,,
391,Perfect Rectangle,29.60%,Hard,23.87%,30.80%,10.02%,3,3.301,,
936,Stamping The Sequence,37.50%,Hard,23.77%,30.80%,10.16%,3,3.305,,
630,Course Schedule III,32.90%,Hard,23.66%,30.80%,10.32%,3,3.310,,
675,Cut Off Trees for Golf Event,33.50%,Hard,23.60%,30.80%,10.40%,3,3.312,,
44,Wildcard Matching,24.10%,Hard,23.45%,30.80%,10.61%,3,3.318,,
1187,Make Array Strictly Increasing,40.60%,Hard,23.19%,30.80%,11.00%,3,3.330,,
335,Self Crossing,27.50%,Hard,22.59%,30.80%,11.87%,3,3.356,,
1269,Number of Ways to Stay in the Same Place After Some Steps,41.20%,Hard,22.59%,30.80%,11.87%,3,3.356,,
1067,Digit Count in Range,38.20%,Hard,22.55%,30.80%,11.92%,3,3.358,,1
1293,Shortest Path in a Grid with Obstacles Elimination,41.40%,Hard,22.44%,30.80%,12.09%,3,3.363,,
968,Binary Tree Cameras,36.60%,Hard,22.40%,30.80%,12.13%,3,3.364,,
1172,Dinner Plate Stacks,39.60%,Hard,22.41%,30.80%,12.12%,3,3.364,,
741,Cherry Pickup,33.10%,Hard,22.23%,30.80%,12.38%,3,3.371,,
710,Random Pick with Blacklist,32.50%,Hard,22.09%,30.80%,12.59%,3,3.378,,
745,Prefix and Suffix Search,33.00%,Hard,22.07%,30.80%,12.61%,3,3.378,,
631,Design Excel Sum Formula,31.20%,Hard,21.95%,30.80%,12.80%,3,3.384,,1
685,Redundant Connection II,32.00%,Hard,21.95%,30.80%,12.78%,3,3.384,,
273,Integer to English Words,25.90%,Hard,21.90%,30.80%,12.87%,3,3.386,,
1012,Numbers With Repeated Digits,36.60%,Hard,21.76%,30.80%,13.07%,3,3.392,,
321,Create Maximum Number,26.30%,Hard,21.59%,30.80%,13.31%,3,3.399,,
1326,Minimum Number of Taps to Open to Water a Garden,41.00%,Hard,21.55%,30.80%,13.36%,3,3.401,,
440,K-th Smallest in Lexicographical Order,27.90%,Hard,21.45%,30.80%,13.52%,3,3.405,,
479,Largest Palindrome Product,28.40%,Hard,21.37%,30.80%,13.62%,3,3.409,,
644,Maximum Average Subarray II,30.80%,Hard,21.35%,30.80%,13.65%,3,3.409,,1
629,K Inverse Pairs Array,30.50%,Hard,21.27%,30.80%,13.76%,3,3.413,,
466,Count The Repetitions,27.90%,Hard,21.07%,30.80%,14.07%,3,3.422,,
1263,Minimum Moves to Move a Box to Their Target Location,39.50%,Hard,20.98%,30.80%,14.20%,3,3.426,,
1001,Grid Illumination,35.10%,Hard,20.42%,30.80%,15.00%,3,3.450,,
1036,Escape a Large Maze,35.60%,Hard,20.41%,30.80%,15.02%,3,3.451,,
1157,Online Majority Element In Subarray,37.20%,Hard,20.23%,30.80%,15.27%,3,3.458,,
719,Find K-th Smallest Pair Distance,30.60%,Hard,20.05%,30.80%,15.53%,3,3.466,,
1178,Number of Valid Words for Each Puzzle,37.20%,Hard,19.92%,30.80%,15.72%,3,3.472,,
1095,Find in Mountain Array,35.50%,Hard,19.44%,30.80%,16.42%,3,3.492,,
927,Three Equal Parts,32.60%,Hard,19.00%,30.80%,17.05%,3,3.511,,
906,Super Palindromes,32.00%,Hard,18.71%,30.80%,17.47%,3,3.524,,
126,Word Ladder II,20.50%,Hard,18.65%,30.80%,17.55%,3,3.527,,
656,Coin Path,28.10%,Hard,18.48%,30.80%,17.81%,3,3.534,,1
1307,Verbal Arithmetic Puzzle,37.40%,Hard,18.23%,30.80%,18.16%,3,3.545,,
871,Minimum Number of Refueling Stops,30.90%,Hard,18.13%,30.80%,18.32%,3,3.549,,
803,Bricks Falling When Hit,29.90%,Hard,18.12%,30.80%,18.32%,3,3.550,,
1199,Minimum Time to Build Blocks,35.70%,Hard,18.11%,30.80%,18.33%,3,3.550,,1
1153,String Transforms Into Another String,34.90%,Hard,17.99%,30.80%,18.51%,3,3.555,,1
891,Sum of Subsequence Widths,31.00%,Hard,17.93%,30.80%,18.60%,3,3.558,,
780,Reaching Points,28.90%,Hard,17.46%,30.80%,19.28%,3,3.578,,
902,Numbers At Most N Given Digit Set,30.70%,Hard,17.47%,30.80%,19.26%,3,3.578,,
1301,Number of Paths with Max Score,36.50%,Hard,17.42%,30.80%,19.34%,3,3.580,,
913,Cat and Mouse,30.70%,Hard,17.31%,30.80%,19.50%,3,3.585,,
493,Reverse Pairs,24.40%,Hard,17.17%,30.80%,19.70%,3,3.591,,
1349,Maximum Students Taking Exam,36.70%,Hard,16.91%,30.80%,20.07%,3,3.602,,
639,Decode Ways II,25.90%,Hard,16.53%,30.80%,20.62%,3,3.619,,
1163,Last Substring in Lexicographical Order,32.60%,Hard,15.54%,30.80%,22.05%,3,3.661,,
1354,Construct Target Array With Multiple Sums,35.10%,Hard,15.24%,30.80%,22.48%,3,3.675,,
952,Largest Component Size by Common Factor,29.10%,Hard,15.14%,30.80%,22.63%,3,3.679,,
1224,Maximum Equal Frequency,33.00%,Hard,15.05%,30.80%,22.76%,3,3.683,,
878,Nth Magical Number,27.50%,Hard,14.62%,30.80%,23.38%,3,3.701,,
1345,Jump Game IV,34.30%,Hard,14.57%,30.80%,23.45%,3,3.703,,
149,Max Points on a Line,16.60%,Hard,14.41%,30.80%,23.68%,3,3.710,,
805,Split Array With Same Average,26.00%,Hard,14.19%,30.80%,24.00%,3,3.720,,
65,Valid Number,14.80%,Hard,13.85%,30.80%,24.50%,3,3.735,,
887,Super Egg Drop,26.20%,Hard,13.19%,30.80%,25.45%,3,3.763,,
1363,Largest Multiple of Three,32.20%,Hard,12.21%,30.80%,26.87%,3,3.806,,
1330,Reverse Subarray To Maximize Array Value,31.10%,Hard,11.59%,30.80%,27.76%,3,3.833,,
564,Find the Closest Palindrome,19.30%,Hard,11.03%,30.80%,28.57%,3,3.857,,
862,Shortest Subarray with Sum at Least K,23.20%,Hard,10.56%,30.80%,29.25%,3,3.878,,
1044,Longest Duplicate Substring,25.30%,Hard,9.99%,30.80%,30.08%,3,3.902,,
420,Strong Password Checker,15.50%,Hard,9.34%,30.80%,31.01%,3,3.930,,`
