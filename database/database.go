package database

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
	Columns []string
	Rows    []*Row
}

// Rowはテーブル内の行の定義
type Row struct {
	Values []*Item[any]
}

// Itemはテーブル内のデータの定義
type Item[T any] struct {
	Value T
}

// Table全体に対しての操作を行うための関数を定義する
//
// CreateTableは新しいテーブルを作成する
func CreateTable() *Table {
	return &Table{}
}

// DropTableはテーブルを削除する
func (t *Table) DropTable() {
	t = nil // 今はメモリ上で管理しているので、メモリから削除する
}

// EditColumnNameはテーブルのカラム名を編集する
func (t *Table) EditColumnName(columnIndex int, newName string) {
	t.Columns[columnIndex] = newName
}

// EditColumnTypeはテーブルのカラムの型を編集する
func (t *Table) EditColumnType(columnIndex int, newType ColumnType) error {
	tt := t.deepCopy()
	ditm, err := generateDefaultItem(newType)
	if err != nil {
		return err
	}
	tt.Rows[0].Values[columnIndex].Value = ditm
	for i, row := range tt.Rows {
		if i == 0 {
			continue
		}
		tt.Rows[i].Values[columnIndex], err = convertType(row.Values[columnIndex].Value, newType)
		if err != nil {
			return err
		}
	}
	*t = *tt
	return nil
}

func (t *Table) deepCopy() *Table {
	newTable := &Table{}
	newTable.Columns = append(newTable.Columns, t.Columns...)
	for _, row := range t.Rows {
		newRow := &Row{}
		for _, item := range row.Values {
			newItem := &Item[any]{Value: item.Value}
			newRow.Values = append(newRow.Values, newItem)
		}
		newTable.Rows = append(newTable.Rows, newRow)
	}
	return newTable
}

func convertType(value any, newType ColumnType) (*Item[any], error) {
	val := reflect.ValueOf(value)
	switch newType {
	case Int64:
		if val.Kind() == reflect.Int64 {
			return &Item[any]{int64(value.(int64))}, nil
		}
	case String:
		if val.Kind() == reflect.String {
			return &Item[any]{string(value.(string))}, nil
		}
	case Float64:
		if val.Kind() == reflect.Float64 {
			return &Item[any]{float64(value.(float64))}, nil
		}
	case Bool:
		if val.Kind() == reflect.Bool {
			return &Item[any]{bool(value.(bool))}, nil
		}
	case Uint64:
		if val.Kind() == reflect.Uint64 {
			return &Item[any]{uint64(value.(uint64))}, nil
		}
	case Byte:
		if val.Kind() == reflect.Uint8 {
			return &Item[any]{byte(value.(byte))}, nil
		}
	case Rune:
		if val.Kind() == reflect.Int32 {
			return &Item[any]{rune(value.(rune))}, nil
		}
	}
	return nil, errors.New("No conversion available '" + fmt.Sprintf("%v", value) + "' for type " + fmt.Sprintf("%T", value) + " to " + string(newType) + ".")
}

// AddColumnはテーブルに新しいカラムを追加する
func (t *Table) AddColumn(name string, columnType ColumnType) error {
	t.Columns = append(t.Columns, name)
	if len(t.Rows) == 0 {
		t.Rows = append(t.Rows, &Row{})
	}
	for _, row := range t.Rows {
		ditm, err := generateDefaultItem(columnType)
		if err != nil {
			return err
		}
		row.Values = append(row.Values, ditm)
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
	for i, item := range row.Values {
		if len(t.Rows) == 0 {
			break
		}
		if reflect.TypeOf(item.Value) != reflect.TypeOf(t.Rows[0].Values[i].Value) {
			return errors.New("Invalid type of value '" + fmt.Sprintf("%v", item.Value) + "' in row. '" +
				fmt.Sprintf("%v", item.Value) + "' is a '" + fmt.Sprintf("%T", item.Value) + "' but should be a '" +
				fmt.Sprintf("%T", t.Rows[0].Values[i].Value) + "'.")
		}
	}
	t.Rows = append(t.Rows, &row)
	return nil
}

// 特定のRowに対しての操作を行うための関数を定義する
//
// UpdateValueは特定のRowの特定のカラムの値を更新する
func (r *Row) UpdateValue(columnIndex int, value any) {
	r.Values[columnIndex].Value = value
}
