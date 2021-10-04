package corutil

// NilPtrToInt returns the dereferenced int or def for nil
func NilPtrToInt(i *int, def int) int {
	if i == nil {
		return def
	}
	return *i
}

// NilPtrToInt32 returns the dereferenced int or def for nil
func NilPtrToInt32(i *int32, def int32) int32 {
	if i == nil {
		return def
	}
	return *i
}

// NilPtrToInt64 returns the dereferenced int or def for nil
func NilPtrToInt64(i *int64, def int64) int64 {
	if i == nil {
		return def
	}
	return *i
}

// NilPtrToFloat32 returns the dereferenced float or def for nil
func NilPtrToFloat32(f *float32, def float32) float32 {
	if f == nil {
		return def
	}
	return *f
}

// NilPtrToFloat64 returns the dereferenced float or def for nil
func NilPtrToFloat64(f *float64, def float64) float64 {
	if f == nil {
		return def
	}
	return *f
}

// NilPtrToStr returns the dereferenced string or def for nil
func NilPtrToStr(s *string, def string) string {
	if s == nil {
		return def
	}
	return *s
}
