/*
package k8s implements a Kubernetes client.
*/
package k8s

func Bool(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

func Int32(s *int32) int32 {
	if s == nil {
		return 0
	}
	return *s
}

func Int64(s *int64) int64 {
	if s == nil {
		return 0
	}
	return *s
}

func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func BoolP(s bool) *bool       { x := s; return &x }
func Int32P(s int32) *int32    { x := s; return &x }
func Int64P(s int64) *int64    { x := s; return &x }
func StringP(s string) *string { x := s; return &x }
