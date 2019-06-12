--
-- @lc app=leetcode.cn id=182 lang=mysql
--
-- [182] 查找重复的电子邮箱
--
# Write your MySQL query statement below

select Email from Person group by Email having count(1)>1

