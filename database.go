package main

import (
	"errors"
	"fmt"
	"reflect"
)

// ColumnTypeはカラムの型を定義
type ColumnType string

const (
	Int64   ColumnType = "int64"
	String  ColumnType = "string"
	Float64 ColumnType = "float64"
	Bool    ColumnType = "bool"
	Uint64  ColumnType = "uint64"
	Byte    ColumnType = "byte"
	Rune    ColumnType = "rune"
)

// StringはColumnTypeの文字列表現を返す
func (c ColumnType) String() string {
	return string(c)
}

// Tableはテーブルの定義
type Table struct {
	columns []string
	rows    []*Row
}

// Rowはテーブル内の行の定義
type Row struct {
	values []*Item[any]
}

// Itemはテーブル内のデータの定義
type Item[T any] struct {
	value T
}

// Table全体に対しての操作を行うための関数を定義する
//
// CreateTableは新しいテーブルを作成する
func CreateTable() *Table {
	return &Table{}
}

// AddColumnはテーブルに新しいカラムを追加する
func (t *Table) AddColumn(name string, columnType ColumnType) error {
	t.columns = append(t.columns, name)
	if len(t.rows) == 0 {
		t.rows = append(t.rows, &Row{})
	}
	for _, row := range t.rows {
		ditm, err := generateDefaultItem(columnType)
		if err != nil {
			return err
		}
		row.values = append(row.values, ditm)
	}
	return nil
}

// generateDefaultItemはカラムの型に応じたデフォルトのItemを生成する
func generateDefaultItem(columnType ColumnType) (*Item[any], error) {
	switch columnType {
	case Int64:
		return &Item[any]{int64(0)}, nil
	case String:
		return &Item[any]{"a"}, nil
	case Float64:
		return &Item[any]{float64(0.0)}, nil
	case Bool:
		return &Item[any]{false}, nil
	case Uint64:
		return &Item[any]{uint64(0)}, nil
	case Byte:
		return &Item[any]{byte(0)}, nil
	case Rune:
		return &Item[any]{rune(0)}, nil
	default:
		return nil, errors.New("Invalid type for column " + string(columnType))
	}
}

// AddRowはテーブルに新しい行を追加する
func (t *Table) AddRow(row Row) error {
	for i, item := range row.values {
		if len(t.rows) == 0 {
			break
		}
		if reflect.TypeOf(item.value) != reflect.TypeOf(t.rows[0].values[i].value) {
			return errors.New("Invalid type of value '" + fmt.Sprintf("%v", item.value) + "' in row. '" +
				fmt.Sprintf("%v", item.value) + "' is a '" + fmt.Sprintf("%T", item.value) + "' but should be a '" +
				fmt.Sprintf("%T", t.rows[0].values[i].value) + "'.")
		}
	}
	t.rows = append(t.rows, &row)
	return nil
}

// 特定のRowに対しての操作を行うための関数を定義する
//
// UpdateValueは特定のRowの特定のカラムの値を更新する
func (r *Row) UpdateValue(columnIndex int, value any) {
	r.values[columnIndex].value = value
}
