CREATE TABlE Customers (
	CustomerID INT PRIMARY KEY AUTO_INCREMENT,
    CustomerName VARCHAR(100) NOT NULL,
    CustomerEmail VARCHAR(100) UNIQUE NOT NULL,
    CustomerPassword BLOB NOT NULL,
    CustomerType ENUM ('Admin', 'Customer'),
    AddressID INT NOT NULL,
    FOREIGN KEY(AddressID) REFERENCES Address(AddressID)
);

CREATE TABLE Clothes (
    ClothesID INT PRIMARY KEY AUTO_INCREMENT,
    ClothesName VARCHAR(100) NOT NULL,
    ClothesCategory ENUM ('Kemeja', 'Celana', 'Hoodie', 'Jaket', 'T-Shirt') NOT NULL,
    ClothesPrice DECIMAL(10,2) NOT NULL,
    ClothesStock INT CHECK(ClothesStock > 0)
);

CREATE TABLE Costumes (
    CostumeID INT PRIMARY KEY AUTO_INCREMENT,
    CostumeName VARCHAR(100) NOT NULL,
    CostumeCategory ENUM ('Cosplay', 'Formal') NOT NULL,
    CostumePrice DECIMAL(10,2) CHECK(CostumePrice > 0) NOT NULL
);

CREATE TABLE Orders (
    OrderID INT PRIMARY KEY AUTO_INCREMENT,
    CustomerID INT,
    OrderDate DATE NOT NULL,
    FOREIGN KEY(CustomerID) REFERENCES Customer(CustomerID)
);

CREATE TABLE Sales (
    SaleID INT PRIMARY KEY AUTO_INCREMENT,
    OrderID INT,
    ClothesID INT,
    Quantity INT CHECK(Quantity > 0) NOT NULL,
    FOREIGN KEY(OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY(ClothesID) REFERENCES Clothes(ClothesID)
);

CREATE TABLE Rents (
    RentID INT PRIMARY KEY AUTO_INCREMENT,
    OrderID INT,
    CostumeID INT,
    Quantity INT CHECK(Quantity > 0) NOT NULL,
    StartDate DATE NOT NULL,
    EndDate DATE NOT NULL,
    FOREIGN KEY(OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY(CostumeID) REFERENCES Costumes(CostumeID)
);

CREATE TABLE Address (
  AddressID INT PRIMARY KEY AUTO_INCREMENT,
  AddressCountry VARCHAR(255) NOT NULL,
  AddressCity VARCHAR(255) NOT NULL,
  AddressStreet VARCHAR(255) NOT NULL
);

