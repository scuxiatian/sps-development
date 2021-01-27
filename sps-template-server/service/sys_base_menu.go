package service

import (
	"errors"
	"gorm.io/gorm"
	"sps-template-server/global"
	"sps-template-server/model"
	"strconv"
)

//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func GetInfoList() (err error, list interface{}, total int64)  {
	var menuList []model.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu

func getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = global.SdDB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu

func GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: err error

func AddBaseMenu(menu model.SysBaseMenu) (err error) {
	if !errors.Is(global.SdDB.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		err = errors.New("存在重复name，请修改name")
		return err
	}
	err = global.SdDB.Create(&menu).Error
	return err
}

//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return:err error

func UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	updateMap := make(map[string]interface{})
	updateMap["keep_alive"] = menu.KeepAlive
	updateMap["default_menu"] = menu.DefaultMenu
	updateMap["parent_id"] = menu.ParentId
	updateMap["path"] = menu.Path
	updateMap["name"] = menu.Name
	updateMap["hidden"] = menu.Hidden
	updateMap["component"] = menu.Component
	updateMap["title"] = menu.Title
	updateMap["icon"] = menu.Icon
	updateMap["sort"] = menu.Sort

	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.SdLog.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.SdLog.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k, _ := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.SdLog.Debug(txErr.Error())
				return txErr
			}
		}
		txErr = db.Updates(updateMap).Error
		if txErr != nil {
			global.SdLog.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func DeleteBaseMenu(id float64) (err error) {
	err = global.SdDB.Preload("Parameters").Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := global.SdDB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.SdDB.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if len(menu.SysAuthoritys) > 0 {
			err = global.SdDB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: err error, menu model.SysBaseMenu

func GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = global.SdDB.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}