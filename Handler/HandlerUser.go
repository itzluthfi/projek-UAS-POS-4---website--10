package Handler

import (
	"THR/Controller"
	"THR/Node"
	"html/template"
	"log"
	"net/http"
)

var (
	tmplUser        = template.Must(template.ParseFiles("View/manageUser/manageUser.html"))
	loginTmpl       = template.Must(template.ParseFiles("View/login/login.html"))
	insertUserTmpl  = template.Must(template.ParseFiles("View/manageUser/insert.html"))
	updateUserTmpl  = template.Must(template.ParseFiles("View/manageUser/update.html"))
	
)

func ViewHandlerUser(w http.ResponseWriter, r *http.Request) {
    users := Controller.GetAllUsers()

    var username, role string
    if cookie, err := r.Cookie("username"); err == nil {
        username = cookie.Value
        user := Controller.GetUserByUsername(username)
        role = user.Role
    }

    data := struct {
        Users    []Node.NodeUser
        Username string
        Role     string
        Success  bool
    }{
        Users:    users,
        Username: username,
        Role:     role,
        Success:  r.URL.Query().Get("login") == "success",
    }

    tmplUser.ExecuteTemplate(w, "manageUser.html", data)
}




func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Proses form submission untuk insert user
		username := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")

		// Panggil fungsi controller untuk insert user
		Controller.InsertUser(username, password, role)

		// Redirect ke halaman manageUser setelah berhasil insert
		http.Redirect(w, r, "/manageUser", http.StatusSeeOther)
		return
	}

	// Jika bukan method POST, tampilkan form insert user
	insertUserTmpl.Execute(w, nil)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := r.URL.Query().Get("username")
		user := Controller.GetUserByUsername(username)

		if err := updateUserTmpl.Execute(w, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		// Handle submission form update
		oldUsername := r.FormValue("oldUsername")
		newUsername := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")

		// Debugging: Log values
		log.Printf("Updating user: oldUsername=%s, newUsername=%s, password=%s, role=%s\n", oldUsername, newUsername, password, role)

		// Panggil fungsi controller untuk update data pengguna
		Controller.UpdateUser(oldUsername, newUsername, password, role)

		// Redirect kembali ke halaman manageUser setelah update
		http.Redirect(w, r, "/manageUser", http.StatusSeeOther)
	}
}


func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	Controller.DeleteUser(username)
	http.Redirect(w, r, "/manageUser", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")
        if valid, role := Controller.VerifikasiUser(username, password); valid {
            // Simpan username di cookie
            http.SetCookie(w, &http.Cookie{
                Name:  "username",
                Value: username,
                Path:  "/",
            })
            
            // Redirect based on role
            if role == "admin" {
                http.Redirect(w, r, "/manageUser?login=success", http.StatusSeeOther)
            } else if role == "kasir" {
                http.Redirect(w, r, "/kasirNonMember?login=success", http.StatusSeeOther)
            }
            return
        } else {
            loginTmpl.Execute(w, map[string]string{"Error": "Username atau password salah"})
            return
        }
    }
    loginTmpl.Execute(w, nil)
}




func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Render template atau tampilkan halaman home
	homeTmpl := template.Must(template.ParseFiles("View/login/login.html"))
	homeTmpl.Execute(w, nil)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Logika logout seperti menghapus sesi atau token di sini jika diperlukan
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
