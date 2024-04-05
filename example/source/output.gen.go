package source

// MergeUpsertPerson merges the source into the target. If an entry exists only in the source, it's added to the target.
// If an entry exists in both, the source's value overwrites the target's.
// Entries present in the target but not in the source are preserved.
// Both target and source must be of the same type.
func MergeUpsertPerson(target Person, source Person) Person {

	if source.Name != "" {
		target.Name = source.Name
	}

	if source.Age != 0 {
		target.Age = source.Age
	}

	if source.Shark != nil {
		target.Shark = source.Shark
	}

	if source.B != false {
		target.B = source.B
	}

	if source.I8 != 0 {
		target.I8 = source.I8
	}

	if source.I16 != 0 {
		target.I16 = source.I16
	}

	if source.I32 != 0 {
		target.I32 = source.I32
	}

	if source.Bytet != 0 {
		target.Bytet = source.Bytet
	}

	if source.Float32 != 0.0 {
		target.Float32 = source.Float32
	}

	if source.Com != 0 + 0i {
		target.Com = source.Com
	}

	if source.Age != 0 {
		target.Age = source.Age
	}

	if source.Shark != nil {
		target.Shark = source.Shark
	}

	if source.B != false {
		target.B = source.B
	}

	if source.Name != "" {
		target.Name = source.Name
	}

	if source.I16 != 0 {
		target.I16 = source.I16
	}

	if source.I32 != 0 {
		target.I32 = source.I32
	}

	if source.Bytet != 0 {
		target.Bytet = source.Bytet
	}

	if source.Float32 != 0.0 {
		target.Float32 = source.Float32
	}

	if source.Com != 0 + 0i {
		target.Com = source.Com
	}

	if source.I8 != 0 {
		target.I8 = source.I8
	}

	return target
}

// MergeUpsertAllowDeletePerson merges the source object into the target object. If a value from the source object
// does not exist in the target object, it will be added. If the target object already contains
// a value from the source object, it will be overwritten with the value from the source object.
// Both the target and source objects must be of the same type.
func MergeUpsertAllowDeletePerson(target Person, source Person) Person {

	target.I8 = source.I8

	target.I16 = source.I16

	target.I32 = source.I32

	target.Bytet = source.Bytet

	target.Float32 = source.Float32

	target.Com = source.Com

	target.Name = source.Name

	target.Age = source.Age

	target.Shark = source.Shark

	target.B = source.B

	target.I8 = source.I8

	target.I16 = source.I16

	target.I32 = source.I32

	target.Bytet = source.Bytet

	target.Float32 = source.Float32

	target.Com = source.Com

	target.Name = source.Name

	target.Age = source.Age

	target.Shark = source.Shark

	target.B = source.B

	return target
}

// MergeUpsertPerson2 merges the source into the target. If an entry exists only in the source, it's added to the target.
// If an entry exists in both, the source's value overwrites the target's.
// Entries present in the target but not in the source are preserved.
// Both target and source must be of the same type.
func MergeUpsertPerson2(target Person2, source Person2) Person2 {

	if source.Name != "" {
		target.Name = source.Name
	}

	if source.Age != 0 {
		target.Age = source.Age
	}

	if source.Shark != nil {
		target.Shark = source.Shark
	}

	if source.Name != "" {
		target.Name = source.Name
	}

	if source.Age != 0 {
		target.Age = source.Age
	}

	if source.Shark != nil {
		target.Shark = source.Shark
	}

	return target
}

// MergeUpsertAllowDeletePerson2 merges the source object into the target object. If a value from the source object
// does not exist in the target object, it will be added. If the target object already contains
// a value from the source object, it will be overwritten with the value from the source object.
// Both the target and source objects must be of the same type.
func MergeUpsertAllowDeletePerson2(target Person2, source Person2) Person2 {

	target.Name = source.Name

	target.Age = source.Age

	target.Shark = source.Shark

	target.Name = source.Name

	target.Age = source.Age

	target.Shark = source.Shark

	return target
}
