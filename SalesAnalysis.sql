/*
Анализ продаж
Условия
В базе данных есть таблицы Invoice, Customer, Employee
Напишите запрос, который будет искать трех продавцов на маркетплейсе, совершивших больше всего продаж, начиная с 2010
года. На выходе в первой колонке должны быть имя и фамилия продавца, а во второй количество их продаж, отсортированное
в порядке убывания.
Примечание
Для решения задачи используется база данных Chinook Database в формате Sqlite - см. файл  Chinook_Sqlite.sqlite в
разделе Данные под описанием задачи.
 */
SELECT FirstName || ' ' || LastName, CountBuys FROM Employee JOIN (SELECT SupportRepId, COUNT(SupportRepId) as CountBuys
FROM Customer JOIN Invoice ON Invoice.CustomerID==Customer.CustomerId WHERE InvoiceDate >= "2010-01-01"
GROUP BY SupportRepId) as Res ON Res.SupportRepId==EmployeeId ORDER BY CountBuys DESC LIMIT 3;
