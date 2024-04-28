package config

const (
	ALLOW_UPDATE_PERMISSION = false
)

const (
	ADD_NEW_ROLE      = "add new role"
	EDIT_ANY_ROLE     = "edit any role"
	VIEW_ALL_ROLE     = "view all role"
	ADD_NEW_USER      = "add new user"
	EDIT_ANY_USER     = "edit any user"
	DELETE_ANY_USER   = "delete any user"
	VIEW_ALL_USER     = "view all user"
	ADD_NEW_CATEGORY  = "add new category"
	EDIT_ANY_CATEGORY = "edit any category"
	VIEW_ALL_CATEGORY = "view all category"
	ADD_NEW_POST      = "add new post"
	EDIT_ANY_POST     = "edit any post"
	VIEW_ALL_POST     = "view all post"
	DELETE_ANY_POST   = "delete any post"
)

type PermissionSeeder struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ModuleSeeder struct {
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Permissisons []PermissionSeeder `json:"permissions"`
}

func GetPermissionsConfig() []ModuleSeeder {
	return []ModuleSeeder{
		{
			Name:        "Users",
			Description: "Tài khoản",
			Permissisons: []PermissionSeeder{
				{
					Name:        ADD_NEW_USER,
					Description: "Thêm mới tài khoản",
				},
				{
					Name:        EDIT_ANY_USER,
					Description: "Sửa tài khoản bất kỳ",
				},
				{
					Name:        VIEW_ALL_USER,
					Description: "Xem tất cả tài khoản",
				},
				{
					Name:        DELETE_ANY_USER,
					Description: "Xoá tài khoản bất kỳ",
				},
			},
		},
		{
			Name:        "Roles",
			Description: "Vai trò",
			Permissisons: []PermissionSeeder{
				{
					Name:        ADD_NEW_ROLE,
					Description: "Thêm mới vai trò",
				},
				{
					Name:        EDIT_ANY_ROLE,
					Description: "Sửa vai trò bất kỳ",
				},
				{
					Name:        VIEW_ALL_ROLE,
					Description: "Xem tất cả vai trò",
				},
			},
		},
		{
			Name:        "Category",
			Description: "Danh mục",
			Permissisons: []PermissionSeeder{
				{
					Name:        ADD_NEW_CATEGORY,
					Description: "Thêm mới danh mục",
				},
				{
					Name:        EDIT_ANY_CATEGORY,
					Description: "Sửa danh mục bất kỳ",
				},
				{
					Name:        VIEW_ALL_CATEGORY,
					Description: "Xem tất cả danh mục",
				},
			},
		},
		{
			Name:        "Post",
			Description: "Bài viết",
			Permissisons: []PermissionSeeder{
				{
					Name:        ADD_NEW_POST,
					Description: "Thêm mới bài viết",
				},
				{
					Name:        EDIT_ANY_POST,
					Description: "Sửa bài viết bất kỳ",
				},
				{
					Name:        VIEW_ALL_POST,
					Description: "Xem tất cả bài viết",
				},
				{
					Name:        DELETE_ANY_POST,
					Description: "Xoá bài viết bất kỳ",
				},
			},
		},
	}
}
