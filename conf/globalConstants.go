/**
 * @Descripttion: 全局常量列表
 */
package conf

const (
	// 登录用户的Session名
	LoginSessionName = "LoginName"

	// 验证码session名
	CaptchaSessionName = "__captcha__"
)

// PageSize 默认分页条数.
const PageSize = 10

// 用户权限
const (
	// 超级管理员.
	MemberSuperRole = 0
	// 普通管理员.
	MemberAdminRole = 1
	// 读者.
	MemberGeneralRole = 2
	// 作者（可以创建作品）
	MemberEditorRole = 3
)

const (
	// 创始人.
	BookFounder = 0
	//管理者
	BookAdmin = 1
	//编辑者.
	BookEditor = 2
	//观察者
	BookObserver = 3
)
