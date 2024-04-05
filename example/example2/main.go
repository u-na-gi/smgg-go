package example2

type A struct {
	Name    string
	Age     int
	Email   string
	Country string
}

type B struct {
	Name    string
	Age     int
	Email   string
	Country string
}

// MergeUpsert はtargetにsourceをマージし、targetにsourceの値が存在しない場合はsourceの値を追加します。
// targetにsourceの値が存在する場合はsourceの値で上書きします。
// targetに値があり、sourceの値が存在しない場合はtargetの値を保持します。
// targetとsourceは同じ型である必要があります。
func MergeUpsert(a A, b B) A {
	if b.Name != "" {
		a.Name = b.Name
	}
	if b.Age != 0 {
		a.Age = b.Age
	}
	if b.Email != "" {
		a.Email = b.Email
	}
	if b.Country != "" {
		a.Country = b.Country
	}
	return a
}

// MergeUpsertAllowDelete はtargetにsourceをマージし、targetにsourceの値が存在しない場合はsourceの値を追加します。
// targetにsourceの値が存在する場合はsourceの値で上書きします。
// targetとsourceは同じ型である必要があります。
func MergeUpsertAllowDelete(a A, b B) A {
	a.Name = b.Name
	a.Age = b.Age
	a.Email = b.Email
	a.Country = b.Country

	return a
}

func main() {
}
