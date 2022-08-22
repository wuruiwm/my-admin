package response

type AdminRoleAuth struct {
	Menu []*AdminRoleAuthDetail `json:"menu"`
	Api  []*AdminRoleAuthDetail `json:"api"`
}

type AdminRoleAuthDetail struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsChecked int    `json:"is_checked"`
	ParentId  string `json:"parent_id"`
}
