package Handler

import (
	"THR/Controller"
	"THR/Model"
	"THR/Node"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

var (
	tmplItem       = template.Must(template.ParseFiles("View/manageItem/manageItem.html"))
	insertItemTmpl = template.Must(template.ParseFiles("View/manageItem/insert.html"))
	updateItemTmpl = template.Must(template.ParseFiles("View/manageItem/update.html"))
)

func GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	items := Controller.ValidasiItemView()
	if items == nil {
		items = []Node.NodeItem{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}


func ViewHandlerItem(w http.ResponseWriter, r *http.Request) {
	items := Controller.ValidasiItemView()
	if items == nil {
		items = []Node.NodeItem{}
	}

	var username, role string
	if cookie, err := r.Cookie("username"); err == nil {
		username = cookie.Value
		user := Controller.GetUserByUsername(username)
		role = user.Role
	}

	data := struct {
		Items    []Node.NodeItem
		Username string
		Role     string
	}{
		Items:    items,
		Username: username,
		Role:     role,
	}

	tmplItem.ExecuteTemplate(w, "manageItem.html", data)
}


func InsertItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		jmlStock, _ := strconv.Atoi(r.FormValue("jmlStock"))
		harga, _ := strconv.Atoi(r.FormValue("harga"))
		diskon, _ := strconv.Atoi(r.FormValue("diskon"))

		if Controller.ValidasiInsertItem(nama, jmlStock, harga, diskon) {
			http.Redirect(w, r, "/manageItem", http.StatusSeeOther)
		} else {
			http.Error(w, "Data tidak valid atau item sudah ada", http.StatusBadRequest)
		}
		return
	}
	insertItemTmpl.Execute(w, nil)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL parameters
	r.ParseForm()
	id := r.Form.Get("id")
	idInt, _ := strconv.Atoi(id)
	item := Model.SearchItemWeb(idInt)

	if r.Method == "GET" {
		// Display the update form with item data
		if err := updateItemTmpl.Execute(w, item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle form submission
		r.ParseForm()
		nama := r.Form.Get("nama")
		jmlStock, _ := strconv.Atoi(r.Form.Get("jmlStock"))
		harga, _ := strconv.Atoi(r.Form.Get("harga"))
		diskon, _ := strconv.Atoi(r.Form.Get("diskon"))

		// Call the controller to update data
		message, success := Controller.ValidasiUpdateItem(nama, jmlStock, harga, diskon, idInt)
		if !success {
			http.Error(w, message, http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/manageItem", http.StatusSeeOther)
	}
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if msg, ok := Controller.ValidasiDeleteItem(id); ok {
		http.Redirect(w, r, "/manageItem", http.StatusSeeOther)
	} else {
		http.Error(w, msg, http.StatusBadRequest)
	}
}



