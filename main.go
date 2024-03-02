package main

import (
	"fmt"

	"kotarou1192/databasego/database"
)

func main() {
	fmt.Println("Hello, playground")
	// テーブルを作成
	table := database.CreateTable()
	// カラムを追加
	table.AddColumn("name", database.String)
	table.AddColumn("age", database.Int64)
	table.AddColumn("email", database.String)
	// データを追加
	name := database.Item[any]{"John"}
	age := database.Item[any]{int64(30)}
	email := database.Item[any]{"john@example.com"}
	row := database.Row{Values: []*database.Item[any]{&name, &age, &email}}
	err := table.AddRow(row)
	// エラーが発生した場合はエラーを表示して終了
	if err != nil {
		fmt.Println(err)
		return
	}
	// カラム名を表示
	fmt.Println(table.Columns)
	// テーブル内のデータを表示
	for _, row := range table.Rows {
		for _, value := range row.Values {
			fmt.Println(value)
		}
	}
	// emailの型が違うデータを追加
	name2 := database.Item[any]{"John2"}
	age2 := database.Item[any]{int64(41)}
	email2 := database.Item[any]{2}
	row2 := database.Row{Values: []*database.Item[any]{&name2, &age2, &email2}}
	err2 := table.AddRow(row2)
	// エラーが発生するので確認
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(table.Columns)
	for _, row := range table.Rows {
		for _, value := range row.Values {
			fmt.Println(value)
		}
	}
	// カラム名を変更
	table.EditColumnName(2, "hoge")
	fmt.Println(table.Columns)
	// カラムの型を変更
	err3 := table.EditColumnType(2, database.Int64)
	if err3 != nil {
		fmt.Println(err3)
	}
	// テーブル内のデータを表示
	for _, row := range table.Rows {
		for _, value := range row.Values {
			fmt.Println(value)
		}
	}
	// 今度は成功するので確認
	err2 = table.AddRow(row2)
	if err2 != nil {
		fmt.Println(err2)
	}
	// テーブル内のデータを表示
	for _, row := range table.Rows {
		for _, value := range row.Values {
			fmt.Println(value)
		}
	}
	fmt.Println("Complete!")
}
