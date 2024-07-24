package Node

type MemberNode struct {
	Id       int
	Username string
	NoTelp   string
	Point    int
	CreateAt string
}

type MemberLL struct {
	Member MemberNode
	Next   *MemberLL
}
