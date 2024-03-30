package tests

import (
	"Ozon/Storages"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestDBConnect(t *testing.T) {
	database := Storages.DBConnection(logrus.New())
	database.AutoMigrate(&Storages.Model{})
	_, err := database.DB()
	if err != nil {
		t.Errorf("Error in database connection: %v", err)
	}
	Storages.DBDisconnection(database, logrus.New())
}

func TestDBDisconnect(t *testing.T) {
	database := Storages.DBConnection(logrus.New())
	database.AutoMigrate(&Storages.Model{})
	Storages.DBDisconnection(database, logrus.New())
	err := database.Commit().Error
	if err == nil {
		t.Errorf("Error to disconnect: %v", err)
	}
}

func TestDBConstructor(t *testing.T) {
	database := Storages.DBConnection(logrus.New())
	database.AutoMigrate(&Storages.Model{})
	obj := Storages.DatabaseConstr(database)
	if obj.Database != database {
		t.Error("Error in db constructor")
	}
	Storages.DBDisconnection(database, logrus.New())
}

func TestDBAdd(t *testing.T) {
	database := Storages.DBConnection(logrus.New())
	database.AutoMigrate(&Storages.Model{})
	st := Storages.DatabaseConstr(database)
	err := st.WriteByUrl("abc", "cfhgjkl", logrus.New())
	if err != nil {
		t.Errorf("Error in writing")
	}
	if st.GetByHash("abc", logrus.New()) != "cfhgjkl" {
		t.Error("Error in DB")
	}
	if st.GetByUrl("cfhgjkl", logrus.New()) != "abc" {
		t.Error("Error in DB")
	}
	st.DeleteByUrl("cfhgjkl", logrus.New())
	Storages.DBDisconnection(database, logrus.New())
}

func TestDBContains(t *testing.T) {
	database := Storages.DBConnection(logrus.New())
	database.AutoMigrate(&Storages.Model{})
	st := Storages.DatabaseConstr(database)
	err := st.WriteByUrl("123", "567", logrus.New())
	if err != nil {
		t.Error("Error writing to database")
	}
	b := st.ContainsByHash("123", logrus.New())
	if !b {
		t.Errorf("Error of containing files")
	}
	b = st.ContainsByUrl("567", logrus.New())
	if !b {
		t.Errorf("Error of containing files")
	}
	st.DeleteByUrl("567", logrus.New())
	Storages.DBDisconnection(database, logrus.New())

}

func TestInMemoryContains(t *testing.T) {
	st := Storages.InMemoryStorageConstr()
	err := st.WriteByUrl("123", "567", logrus.New())
	if err != nil {
		t.Error("Error writing to storage")
	}
	b := st.ContainsByHash("123", logrus.New())
	if !b {
		t.Errorf("Error of containing files")
	}
	b = st.ContainsByUrl("567", logrus.New())
	if !b {
		t.Errorf("Error of containing files")
	}
	st.DeleteByUrl("567", logrus.New())

}

func TestInMemoryConstructor(t *testing.T) {
	obj := Storages.InMemoryStorageConstr()
	if len(obj.StorageUrlByHash) != 0 || len(obj.StorageHashByUrl) != 0 {
		t.Error("Error in in memory constructor")
	}
}

func TestInMemoryAdd(t *testing.T) {
	st := Storages.InMemoryStorageConstr()
	err := st.WriteByUrl("abc", "cfhgjkl", logrus.New())
	if err != nil {
		t.Errorf("Error in writing")
	}
	if st.GetByHash("abc", logrus.New()) != "cfhgjkl" {
		t.Error("Error in storage")
	}
	if st.GetByUrl("cfhgjkl", logrus.New()) != "abc" {
		t.Error("Error in storage")
	}
	st.DeleteByUrl("cfhgkl", logrus.New())
}
