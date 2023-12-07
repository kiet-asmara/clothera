CREATE TABLE `Address` (
  `AddressID` INT PRIMARY KEY AUTO_INCREMENT,
  `AddressCountry` VARCHAR(255) NOT NULL,
  `AddressCity` VARCHAR(255) NOT NULL,
  `AddressStreet` VARCHAR(255) NOT NULL
);

CREATE TABLE `Customers` (
  `CustomerID` INT PRIMARY KEY AUTO_INCREMENT,
  `AddressID` INT NOT NULL,
  `CustomerName` VARCHAR(100) NOT NULL,
  `CustomerEmail` VARCHAR(100) UNIQUE NOT NULL,
  `CustomerPassword` BLOB NOT NULL,
  `CustomerType` ENUM('admin', 'user') NOT NULL
);

CREATE TABLE `Clothes` (
  `ClothesID` INT PRIMARY KEY AUTO_INCREMENT,
  `ClothesName` VARCHAR(100) NOT NULL,
  `ClothesCategory` ENUM('Kemeja', 'Celana', 'Hoodie', 'Jaket', 'T-Shirt') NOT NULL,
  `ClothesPrice` DECIMAL(10,2) NOT NULL CHECK(`ClothesPrice` >= 0),
  `ClothesStock` INT NOT NULL CHECK(`ClothesStock` >= 0)
);

CREATE TABLE `Costumes` (
  `CostumeID` INT PRIMARY KEY AUTO_INCREMENT,
  `CostumeName` VARCHAR(100) NOT NULL,
  `CostumeCategory` ENUM('Cosplay', 'Formal') NOT NULL,
  `CostumePrice` DECIMAL(10,2) NOT NULL CHECK(`CostumePrice` >= 0),
  CostumeStock INT NOT NULL CHECK(CostumeStock >= 0)
);

CREATE TABLE `Orders` (
  `OrderID` INT PRIMARY KEY AUTO_INCREMENT,
  `CustomerID` INT NOT NULL,
  `OrderDate` DATE NOT NULL,
  TotalPrice INT CHECK(TotalPrice >= 0)
);

CREATE TABLE `Sales` (
  `SaleID` INT PRIMARY KEY AUTO_INCREMENT,
  `OrderID` INT NOT NULL,
  `ClothesID` INT NOT NULL,
  `Quantity` INT NOT NULL CHECK(`Quantity` >= 0)
);

CREATE TABLE `Rents` (
  `RentID` INT PRIMARY KEY AUTO_INCREMENT,
  `OrderID` INT NOT NULL,
  `CostumeID` INT NOT NULL,
  `Quantity` INT NOT NULL CHECK(`Quantity` >= 0),
  `StartDate` DATE NOT NULL,
  `EndDate` DATE NOT NULL,
  RentPrice DEC(10,2) NOT NULL CHECK(RentPrice >= 0)
);

ALTER TABLE `Customers` ADD FOREIGN KEY (`AddressID`) REFERENCES `Address` (`AddressID`);

ALTER TABLE `Orders` ADD FOREIGN KEY (`CustomerID`) REFERENCES `Customers` (`CustomerID`);

ALTER TABLE `Sales` ADD FOREIGN KEY (`OrderID`) REFERENCES `Orders` (`OrderID`);

ALTER TABLE `Sales` ADD FOREIGN KEY (`ClothesID`) REFERENCES `Clothes` (`ClothesID`);

ALTER TABLE `Rents` ADD FOREIGN KEY (`OrderID`) REFERENCES `Orders` (`OrderID`);

ALTER TABLE `Rents` ADD FOREIGN KEY (`CostumeID`) REFERENCES `Costumes` (`CostumeID`);