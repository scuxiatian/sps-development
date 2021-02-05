package service

import (
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
)

//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.SysMenu)  {
	var allMenus []model.SysMenu
	treeMap = make(map[string][]model.SysMenu)
	err = global.SdDB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error

func getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error)  {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.SysMenu

func GetMenuTree(authorityId string) (err error, menus []model.SysMenu)  {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error

func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu

func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []model.SysMenu) {
	err = global.SdDB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	return err, menus
}
