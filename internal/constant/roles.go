package constant

type Role struct {
	ID          uint
	Name        string
	DisplayName string
}

var (
	SystemRoleSuperuser = Role{1000, "superuser", "超级管理员"}
	SystemRoleAdmin     = Role{1001, "admin", "管理员"}
)

var (
	TeamRoleLeader              = Role{2000, "leader", "组长"}
	TeamRoleViceLeader          = Role{2001, "vice_leader", "副组长"}
	TeamRoleTranslator          = Role{2002, "translator", "翻译"}
	TeamRoleTimelineMaker       = Role{2003, "timeline_maker", "时轴"}
	TeamRoleProofreader         = Role{2004, "proofreader", "校对"}
	TeamRolePostProductionStaff = Role{2005, "post_production_staff", "后期"}
	TeamRoleCompressor          = Role{2006, "compressor", "压制"}
	TeamRoleSupervisor          = Role{2007, "supervisor", "监制"}
)
