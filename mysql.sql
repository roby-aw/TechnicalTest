1. Employee
SELECT e.Name as Employee FROM Employee e, Employee m WHERE e.ManagerId is not NULL and e.ManagerId = m.ID and e.Salary > m.Salary

2. SecondHighestSalary
SELECT MAX(Salary) as SecondHighestSalary FROM Employee WHERE SALARY < (SELECT MAX(Salary) FROM Employee);

3. Classes More Than 5 Students
SELECT class FROM class group by class having count(distinct student)>=5;
