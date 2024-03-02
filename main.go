package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	// テーブルを作成
	table := CreateTable()
	// カラムを追加
	table.AddColumn("name", String)
	table.AddColumn("age", Int64)
	table.AddColumn("email", String)
	// データを追加
	name := Item[any]{"John"}
	age := Item[any]{int64(30)}
	email := Item[any]{"john@example.com"}
	row := Row{values: []*Item[any]{&name, &age, &email}}
	err := table.AddRow(row)
	// エラーが発生した場合はエラーを表示して終了
	if err != nil {
		fmt.Println(err)
		return
	}
	// カラム名を表示
	fmt.Println(table.columns)
	// テーブル内のデータを表示
	for _, row := range table.rows {
		for _, value := range row.values {
			fmt.Println(value)
		}
	}
	// emailの型が違うデータを追加
	name2 := Item[any]{"John2"}
	age2 := Item[any]{41}
	email2 := Item[any]{2}
	row2 := Row{values: []*Item[any]{&name2, &age2, &email2}}
	err2 := table.AddRow(row2)
	// エラーが発生するので確認
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(table.columns)
	for _, row := range table.rows {
		for _, value := range row.values {
			fmt.Println(value)
		}
	}
	fmt.Println("Complete!")
}
