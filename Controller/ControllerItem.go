package Controller

import (
	"THR/Database"
	"THR/Model"
	"THR/Node"
)

//SERVICE
func ValidasiTambahStokItem(id int,addJml int)(string,bool){
	hasil := Model.SearchItem(id)
	if(hasil != nil){
        Model.TambahStokItem(hasil.Next,addJml)
		return "Stok Item Berhasil Di Tambahkan!",true
	}
	return "Data Item Dengan Id Tersebut Tidak Di Temukan!",false
}

func ValidasiKurangiStokItem(id int,addJml int)(string,bool){
	hasil := Model.SearchItem(id)
	if(hasil != nil){
        Model.KurangiStokItem(hasil.Next,addJml)
		return "Stok Item Berhasil Di Tambahkan!",true
	}
	return "Data Item Dengan Id Tersebut Tidak Di Temukan!",false
}

func ValidasiInsertItem(nama string,jmlStock int, Harga int,diskon int) bool {

	//jika input tidak valid
    if(nama == "" || Harga == 0 || jmlStock == 0){
       return false
	}	
	
	Model.InsertItem(nama, jmlStock,Harga,diskon)
	return true
}


func ValidasiItemView()[]Node.NodeItem{
	AllItem := Model.ReadAllItem()
    if AllItem == nil {
		return nil
	}
    return AllItem
}

func ValidasiSearchItem(id int) (string, *Node.ItemLL) {
   cur := &Database.HeadItem
   
   //jika data tidak kosong
   if(cur.Next != nil){
	    result := Model.SearchItem(id)
		if(result != nil){
			return "Searching Berhasil, Data Item Berhasil Ditemukan!",result.Next
		}else{
			return "Searching GAGAL, Data Item Dengan Id Tersebut Tidak Ditemukan",nil
		}
   }

	//jika data kosong
   return "Searching GAGAL, Data Item Masih Kosong!",nil
}


func ValidasiDeleteItem(id int)(string,bool){
	cur := &Database.HeadItem
	
	//jika data tidak kosong
	if(cur.Next != nil){
		result := Model.SearchItem(id)
		if result == nil {
			return "Delete Gagal, Data Item Dengan Id Tersebut Tidak Ditemukan",false
		}else{
		    Model.DeleteItem(result)
		    return "Delete Berhasil, Data Mahasiwa Berhasil Dihapus",true
		}
    }

	return "Delete GAGAL, Data Item Masih Kosong",false
}


func ValidasiUpdateItem(nama string,jmlStock int,Harga int,diskon int,id int)(string,bool){
	cur := &Database.HeadItem

	if(cur.Next != nil){
		result := Model.SearchItem(id)
		//jika data tidak ketemu
		if(result == nil){
			return "Update GAGAL, Data Item Dengan Id Tersebut Tidak Ditemukan",false
		}else{
			//jika data ketemu
			Model.UpdateItem(result.Next,nama,jmlStock,Harga,diskon)
			return "Update Berhasil, Data Item Berhasil di Update!",true
		}
    }

	return "Update GAGAL, Data Item Masih Kosong",false
	
}