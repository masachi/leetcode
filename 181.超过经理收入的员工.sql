--
-- @lc app=leetcode.cn id=181 lang=mysql
--
-- [181] 超过经理收入的员工
--
# Write your MySQL query statement below

select Name as 'Employee' from Employee a where Salary > (select Salary from Employee where a.ManagerId = Id)

