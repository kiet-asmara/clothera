package handler

import (
	"database/sql"
	"log"
	"pair-project/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_AddClothes(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal("error init mock:", err)
	}
	defer db.Close()

	// succes test case
	clothesSuccess := entity.Clothes{
		ClothesName:     "Sweater",
		ClothesCategory: "Jaket",
		ClothesPrice:    50,
		ClothesStock:    50,
	}
	mock.ExpectExec("INSERT INTO Clothes (ClothesName, ClothesCategory, ClothesPrice, ClothesStock) VALUES (?, ?, ?, ?)").
		WithArgs(clothesSuccess.ClothesName, clothesSuccess.ClothesCategory, clothesSuccess.ClothesPrice, clothesSuccess.ClothesStock).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = InsertProductIntoDatabase(db, clothesSuccess)
	if err != nil {
		t.Errorf("Test_AddClothes success assertion failed got err: %v", err)
	}
}

// read
func TestRentPrice(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal("error init mock:", err)
	}
	defer db.Close()

	costume := entity.Costume{
		CostumeID:       1,
		CostumeName:     "Naruto",
		CostumeCategory: "Cosplay",
		CostumePrice:    50,
		CostumeStock:    50,
	}

	rent := entity.Rent{
		CostumeID: 1,
		StartDate: "2023-11-20",
		EndDate:   "2023-11-25",
	}

	// 	query := `SELECT Costumes.CostumePrice FROM Costumes
	// 	// JOIN Rents ON Costumes.CostumeID = Rents.CostumeID
	// 	// WHERE Costumes.CostumeID = \\?`

	// 	rows := sqlmock.NewRows([]string{"CostumePrice"}).AddRow(c.CostumePrice)

	// 	mock.ExpectQuery(query).WithArgs(c.CostumeID).WillReturnRows(rows)

	// 	user, err := repo.FindByID(u.ID)
	// 	assert.NotNil(t, user)
	// 	assert.NoError(t, err)
	// }

	mock.ExpectQuery(`SELECT Costumes.CostumePrice FROM Costumes
	JOIN Rents ON Costumes.CostumeID = Rents.CostumeID
	WHERE Costumes.CostumeID = ?`).WithArgs(costume.CostumeID).WillReturnRows(
		sqlmock.NewRows([]string{"CostumePrice"}).AddRow(costume.CostumePrice))

	_, err = RentPrice(db, rent)
	if err != nil {
		t.Errorf("Test_AddClothes success assertion failed got err: %v", err)
	}
}
