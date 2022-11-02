-- 3 entidades Products - ProductPriceHistory - Clients - Bills

-- Products
--  id_product
--  product_name
--  product_stock
--  id_price

-- ProductPriceHistory
--  id_price
--  price
--  badge
--  date

-- Clients
--  id_client
--  client_first_name
--  client_middle_name
--  client_surname
--  client_second_surname    
--  client_birthday

-- Bills
--  id_bill
--  id_product
--  id_client
--  buy_date
--  

-- 1 Product has 1 to many ProductPriceHistory
-- 1 Bill has 1 to many Product
-- 1 Bill has 1 Client