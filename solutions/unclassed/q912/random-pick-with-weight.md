

# Random Pick with Weight
refer: https://leetcode.com/problems/random-pick-with-weight

Given an array w of positive integers, where w[i] describes the weight of index i, write a function pickIndex which randomly picks an index in proportion to its weight.

Note:


	1 &lt;= w.length &lt;= 10000
	1 &lt;= w[i] &lt;= 10^5
	pickIndex will be called at most 10000 times.


Example 1:

Input: 
[&#34;Solution&#34;,&#34;pickIndex&#34;]
[[[1]],[]]
Output: [null,0]



Example 2:

Input: 
[&#34;Solution&#34;,&#34;pickIndex&#34;,&#34;pickIndex&#34;,&#34;pickIndex&#34;,&#34;pickIndex&#34;,&#34;pickIndex&#34;]
[[[1,3]],[],[],[],[],[]]
Output: [null,0,1,1,1,0]


Explanation of Input Syntax:

The input is two lists: the subroutines called and their arguments. Solution&#39;s constructor has one argument, the array w. pickIndex has no arguments. Arguments are always wrapped with a list, even if there aren&#39;t any.



#### Tags

- Binary Search
- Random



