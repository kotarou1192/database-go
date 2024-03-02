package database

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	table := CreateTable()
	if table == nil {
		t.Error("failed to create table")
	}
}

func TestDropTable(t *testing.T) {
	table := CreateTable()
	table.DropTable()
	if table == nil {
		t.Error("failed to drop table")
	}
}

func TestAddColumn(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	if len(table.Columns) != 2 {
		t.Error("failed to add column")
		return
	}
	if table.Columns[0] != "name" && table.Columns[1] != "age" {
		t.Error("failed to add column")
	}
}

func TestAddRow(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	name := Item[any]{"John"}
	age := Item[any]{int64(30)}
	email := Item[any]{"email@example.com"}
	row := Row{Values: []*Item[any]{&name, &age, &email}}
	err := table.AddRow(row)
	if err != nil {
		t.Error("failed to add row")
	}
}

func TestEditColumnName(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	table.EditColumnName(2, "hoge")
	if table.Columns[2] != "hoge" {
		t.Error("failed to edit column name")
	}
}

func TestEditColumnType(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	err := table.EditColumnType(2, Int64)
	if err != nil {
		t.Error("failed to edit column type")
	}
}

func TestEditColumnTypeFailedThenDataRollbacks(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	name := Item[any]{"John"}
	age := Item[any]{int64(30)}
	email := Item[any]{"email@example.com"}
	row := Row{Values: []*Item[any]{&name, &age, &email}}
	err := table.AddRow(row)
	if err != nil {
		t.Error("failed to add row")
		t.Error(err)
		return
	}
	err = table.EditColumnType(2, Int64)
	if err == nil {
		t.Error("should be failed. email is not convertible to int64.")
	}
	if table.Rows[1].Values[2].Value != email.Value {
		t.Error("failed to rollback")
	}
}

func TestEditColumnTypeSuccessThenDataUpdated(t *testing.T) {
	table := CreateTable()
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	name := Item[any]{"John"}
	age := Item[any]{int64(30)}
	email := Item[any]{"email@example.com"}
	row := Row{Values: []*Item[any]{&name, &age, &email}}
	err := table.AddRow(row)
	if err != nil {
		t.Error("failed to add row")
		t.Error(err)
		return
	}
	err = table.EditColumnType(1, String)
	if err == nil {
		t.Error("should be failed. email is not convertible to int64.")
		return
	}
	if table.Rows[1].Values[1].Value == "30" {
		t.Error("failed to update")
	}
}
