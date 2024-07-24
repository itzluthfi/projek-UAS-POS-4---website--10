package Controller

import (
	"THR/Database"
	"THR/Model"
	"THR/Node"
)

//cek apakah ada data yang sama
func IsMemberSame (username string, noTelp string)bool{
	cur := &Database.HeadMember
	for cur.Next != nil {
		if(cur.Next.Member.Username == username && cur.Next.Member.NoTelp == noTelp){
			return true
		}
		cur = cur.Next
	}
	return false
}

func ValidasiInsertMember(username string, noTelp string,point int) bool {
	//tidak boleh ada data yang sama
	isDataSama := IsMemberSame(username,noTelp)
	if username == "" ||  isDataSama {
		return false
	}
	Model.MemberInsert(username,noTelp,point)
	return true
}

func ValidasiMembersView() []Node.MemberNode {
	AllMember := Model.MemberReadAll()
	if AllMember == nil {
		return nil
	}
	return AllMember
}

func ValidasiDeleteMember(id int)(string,bool){
	result := Model.SearchMember(id)
    
	if(result != nil){
		Model.MemberDelete(result)
		return "Data Member dengan Id tersebut Telah Dihapus!",true
	}
	return "Tidak Ada Data Member Dengan Id Tersebut!",false
}

func ValidasiSearchMember(id int)(string,*Node.MemberLL){
	cur := &Database.HeadMember
	if(cur.Next != nil){
		result := Model.SearchMember(id)
		if(result != nil){
			return "Data Member Ditemukan!",result.Next
		}
		return "Tidak ada Data Member dengan id Tersebut!",nil
	}
	return "Data Member Masih Kosong!",nil
}

func ValidasiUpdateMember(id int,username string,noTelp string)(string,bool){
	result := Model.SearchMember(id)
	if(result != nil){
		Model.MemberUpdate(result.Next,username,noTelp)
		return "Data Member Telah Di Update!",true
	}
	return "Data Member Tersebut Tidak Ditemukan!",false
}

//WEB
func ValidasiTambahMemberPoints(id int, pointReward int) (string, bool) {
    message, member := ValidasiSearchMember(id)
    if member == nil {
        return message, false
    }
    // Update member's points
    Model.TambahMemberPoint(member, pointReward)

    return "Member points updated successfully", true
}

func ValidasiKurangiMemberPoints(id int,  pointUsed int) (string, bool) {
    message, member := ValidasiSearchMember(id)
    if member == nil {
        return message, false
    }
    // Update member's points
    Model.KurangiMemberPoint(member, pointUsed)

    return "Member points updated successfully", true
}
