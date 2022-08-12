// 既存の関数を呼び出し
Function()

// 無名関数をその場で作って実行
func() {
	// ここに処理を書く
}()

// 別のgoroutineを作って、既存の関数を呼び出し
go Function()

// 別のgoroutineを作って、無名関数をその場で作って実行
go func() {
	// 別のgoroutineから呼びたい処理を書く
}()