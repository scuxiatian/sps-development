package datas

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"sps-template-server/global"
	"sps-template-server/model"
	"time"
)

var BaseMenus = []model.SysBaseMenu{
	{SdModel: global.SdModel{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "HomeOutlined"}},
	{SdModel: global.SdModel{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "about/index.vue", Sort: 7, Meta: model.Meta{Title: "关于我们", Icon: "InfoCircleOutlined"}},
	{SdModel: global.SdModel{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "systemManage/index.vue", Sort: 3, Meta: model.Meta{Title: "系统管理", Icon: "SettingOutlined"}},
	{SdModel: global.SdModel{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "systemManage/authority/index.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "TeamOutlined"}},
	{SdModel: global.SdModel{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "systemManage/menu/index.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "MenuOutlined", KeepAlive: true}},
	{SdModel: global.SdModel{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "systemManage/api/index.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "LaptopOutlined", KeepAlive: true}},
	{SdModel: global.SdModel{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "systemManage/user/index.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "UserOutlined"}},
	{SdModel: global.SdModel{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "person/index.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "UserOutlined"}},
	{SdModel: global.SdModel{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "fileManage/index.vue", Sort: 6, Meta: model.Meta{Title: "文件管理", Icon: "FileOutlined"}},
	{SdModel: global.SdModel{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "fileManage/excel/index.vue", Sort: 4, Meta: model.Meta{Title: "excel导入导出", Icon: "FileExcelOutlined"}},
	{SdModel: global.SdModel{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "fileManage/upload/index.vue", Sort: 5, Meta: model.Meta{Title: "上传下载", Icon: "UploadOutlined"}},
	{SdModel: global.SdModel{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "fileManage/breakpoint/index.vue", Sort: 6, Meta: model.Meta{Title: "断点续传", Icon: "CloudUploadOutlined"}},
	{SdModel: global.SdModel{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "fileManage/customer/index.vue", Sort: 7, Meta: model.Meta{Title: "客户列表（资源示例）", Icon: "ContactsOutlined"}},
	{SdModel: global.SdModel{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "ToolOutlined"}},
	{SdModel: global.SdModel{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "CodeOutlined", KeepAlive: true}},
	{SdModel: global.SdModel{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "FormOutlined", KeepAlive: true}},
	{SdModel: global.SdModel{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "systemTools/system/index.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
	{SdModel: global.SdModel{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "systemManage/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "BookOutlined"}},
	{SdModel: global.SdModel{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "systemManage/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "BookOutlined"}},
	{SdModel: global.SdModel{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "systemManage/operation/index.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "HistoryOutlined"}},
	{SdModel: global.SdModel{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "fileManage/simpleUploader/index", Sort: 6, Meta: model.Meta{Title: "断点续传（插件版）", Icon: "CloudUploadOutlined"}},
	{SdModel: global.SdModel{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "CloudOutlined"}},
	{SdModel: global.SdModel{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "workflow", Name: "workflow", Hidden: false, Component: "workflow/index.vue", Sort: 5, Meta: model.Meta{Title: "流程功能", Icon: "ApartmentOutlined"}},
	{SdModel: global.SdModel{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowCreate", Name: "workflowCreate", Hidden: false, Component: "workflow/workflowCreate/index.vue", Sort: 0, Meta: model.Meta{Title: "流程绘制", Icon: "PlusCircleOutlined"}},
	{SdModel: global.SdModel{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowProcess", Name: "workflowProcess", Hidden: false, Component: "workflow/workflowProcess/index.vue", Sort: 0, Meta: model.Meta{Title: "流程列表", Icon: "UnorderedListOutlined"}},
	{SdModel: global.SdModel{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowUse", Name: "workflowUse", Hidden: true, Component: "workflow/workflowUse/index.vue", Sort: 0, Meta: model.Meta{Title: "使用工作流", Icon: "ApartmentOutlined"}},
	{SdModel: global.SdModel{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "started", Name: "started", Hidden: false, Component: "workflow/userList/started.vue", Sort: 0, Meta: model.Meta{Title: "我发起的", Icon: "AppstoreAddOutlined"}},
	{SdModel: global.SdModel{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "need", Name: "need", Hidden: false, Component: "workflow/userList/need.vue", Sort: 0, Meta: model.Meta{Title: "我的待办", Icon: "BellOutlined"}},
}

func InitSysBaseMenus(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			global.SdLog.Warn("sys_base_menus表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&BaseMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		global.SdLog.Error("[Mysql]--> sys_base_menus 表的初始数据失败,err: %v\n", zap.Any("err", err))
		os.Exit(0)
	}
}
