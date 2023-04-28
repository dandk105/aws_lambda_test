package schema

type User struct {
	name        string `form:"name"`
	family_name string `form:"family_name"`
	id          int64  `form:"id"`
}
