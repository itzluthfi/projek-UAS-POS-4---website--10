package Handler

import (
	"THR/Controller"
	"THR/Model"
	"THR/Node"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var (
    tmpl       = template.Must(template.ParseFiles("View/manageMember/manageMember.html"))
    insertTmpl = template.Must(template.ParseFiles("View/manageMember/insert.html"))
    updateTmpl = template.Must(template.ParseFiles("View/manageMember/update.html"))
    //loginForm  = template.Must(template.ParseFiles("View/login/login.html"))
)


func ViewHandlerMember(w http.ResponseWriter, r *http.Request) {
    members := Controller.ValidasiMembersView()
    if members == nil {
        members = []Node.MemberNode{}
    }

    var username, role string
    if cookie, err := r.Cookie("username"); err == nil {
        username = cookie.Value
        user := Controller.GetUserByUsername(username)
        role = user.Role
    }

    data := struct {
        Members  []Node.MemberNode
        Username string
        Role     string
    }{
        Members:  members,
        Username: username,
        Role:     role,
    }

    if err := tmpl.ExecuteTemplate(w, "manageMember.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func InsertMemberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		noTelp := r.FormValue("noTelp")
		pointStr := r.FormValue("point")
		point, err := strconv.Atoi(pointStr)
        if err != nil {
			http.Error(w, "Point harus berupa angka", http.StatusBadRequest)
			return
		}

        
		if Controller.ValidasiInsertMember(username, noTelp, point) {
			http.Redirect(w, r, "/manageMember", http.StatusSeeOther)
		} else {
			http.Error(w, "Data tidak valid atau pengguna sudah ada", http.StatusBadRequest)
		}
		return
	}
	insertTmpl.Execute(w, nil)
}

func UpdateMemberHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL parameters
	r.ParseForm()
	id := r.Form.Get("id")
	idInt, _ := strconv.Atoi(id)
	user := Model.SearchMemberWeb(idInt)

	if r.Method == "GET" {
		// Display the update form with user data
		if err := updateTmpl.Execute(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		username := r.Form.Get("username")
		noTelp := r.Form.Get("noTelp")

		// Call the controller to update data
		message, success := Controller.ValidasiUpdateMember(idInt, username, noTelp)
		if !success {
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/manageMember", http.StatusSeeOther)
	}
}



func DeleteMemberHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if msg, ok := Controller.ValidasiDeleteMember(id); ok {
		http.Redirect(w, r, "/manageMember", http.StatusSeeOther)
		
	} else {
		http.Error(w, msg, http.StatusBadRequest)
	}
}



//WEB
func GetMemberDetailsHandler(w http.ResponseWriter, r *http.Request) {
    memberIdStr := r.URL.Query().Get("id")
    memberId, err := strconv.Atoi(memberIdStr)
    if err != nil {
        http.Error(w, "Invalid member ID", http.StatusBadRequest)
        return
    }

    message, member := Controller.ValidasiSearchMember(memberId)
    if member == nil {
        http.Error(w, message, http.StatusNotFound)
        return
    }

    jsonResponse, err := json.Marshal(member.Member)
    if err != nil {
        http.Error(w, "Error converting member data to JSON", http.StatusInternalServerError)
        return
    }

    fmt.Println("JSON Response:", string(jsonResponse))  // Logging data JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func TambahMemberPointsHandler(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan ID anggota dari parameter URL
    memberIdStr := r.URL.Query().Get("id")
    memberId, err := strconv.Atoi(memberIdStr)
    if err != nil {
        http.Error(w, "ID anggota tidak valid", http.StatusBadRequest)
        return
    }

    // Mendapatkan jumlah poin dari parameter URL
    poinStr := r.URL.Query().Get("poin")
    pointReward, err := strconv.Atoi(poinStr)
    if err != nil {
        http.Error(w, "Nilai poin tidak valid", http.StatusBadRequest)
        return
    }

    // Validasi dan tambahkan poin anggota
    message, success := Controller.ValidasiTambahMemberPoints(memberId, pointReward)
    if !success {
        http.Error(w, message, http.StatusBadRequest)
        return
    }

    // Menyiapkan respons JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    response := map[string]string{"message": "Poin berhasil diperbarui"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error saat mengkodekan respons JSON", http.StatusInternalServerError)
        return
    }
    w.Write(jsonResponse)
}

func KurangiMemberPointsHandler(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan ID anggota dari parameter URL
    memberIdStr := r.URL.Query().Get("id")
    memberId, err := strconv.Atoi(memberIdStr)
    if err != nil {
        http.Error(w, "ID anggota tidak valid", http.StatusBadRequest)
        return
    }

    // Mendapatkan jumlah poin yang akan dikurangi dari parameter URL
    poinStr := r.URL.Query().Get("poin")
    pointUsed, err := strconv.Atoi(poinStr)
    if err != nil {
        http.Error(w, "Nilai poin tidak valid", http.StatusBadRequest)
        return
    }

    // Validasi dan kurangi poin anggota
    message, success := Controller.ValidasiKurangiMemberPoints(memberId, pointUsed)
    if !success {
        http.Error(w, message, http.StatusBadRequest)
        return
    }

    // Menyiapkan respons JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    response := map[string]string{"message": "Poin berhasil dikurangi"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error saat mengkodekan respons JSON", http.StatusInternalServerError)
        return
    }
    w.Write(jsonResponse)
}



// func TambahMemberPointsHandler(w http.ResponseWriter, r *http.Request) {
//     memberIdStr := r.URL.Query().Get("id")
//     memberId, err := strconv.Atoi(memberIdStr)
//     if err != nil {
//         http.Error(w, "Invalid member ID", http.StatusBadRequest)
//         return
//     }

//     poinStr := r.URL.Query().Get("poin")
//     pointReward, err := strconv.Atoi(poinStr)
//     if err != nil {
//         http.Error(w, "Invalid points value", http.StatusBadRequest)
//         return
//     }

//     message, success := Controller.ValidasiTambahMemberPoints(memberId, pointReward)
//     if !success {
//         http.Error(w, message, http.StatusBadRequest)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     response := map[string]string{"message": "Points updated successfully"}
//     jsonResponse, err := json.Marshal(response)
//     if err != nil {
//         http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
//         return
//     }
//     w.Write(jsonResponse)
// }

// func KurangiMemberPointsHandler(w http.ResponseWriter, r *http.Request) {
//     memberIdStr := r.URL.Query().Get("id")
//     memberId, err := strconv.Atoi(memberIdStr)
//     if err != nil {
//         http.Error(w, "Invalid member ID", http.StatusBadRequest)
//         return
//     }

//     poinStr := r.URL.Query().Get("poin")
//     pointUsed, err := strconv.Atoi(poinStr)
//     if err != nil {
//         http.Error(w, "Invalid points value", http.StatusBadRequest)
//         return
//     }

//     message, success := Controller.ValidasiTambahMemberPoints(memberId, pointUsed)
//     if !success {
//         http.Error(w, message, http.StatusBadRequest)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)
//     response := map[string]string{"message": "Points updated successfully"}
//     jsonResponse, err := json.Marshal(response)
//     if err != nil {
//         http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
//         return
//     }
//     w.Write(jsonResponse)
// }