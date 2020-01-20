package main

import (
	"fmt"
	"net/http"
	"src/pkg/forms"
	"src/pkg/models"
	"time"

	"github.com/gorilla/mux"
)

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "home.page.html", &templateData{})
}
func (app *application) offertsPage(w http.ResponseWriter, r *http.Request) {
	// panic("oops! Something went wrong!")
	// var posts []*models.Post

	posts, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "offerts.page.html", &templateData{
		Posts: posts,
	})
}

func (app *application) showPostPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println(vars["id"])
	// var userid = r.URL.Query().Get(":id")
	// fmt.Println(userid)

	var userid = vars["id"]
	post, err := app.posts.Get(userid)
	// fmt.Println(post.Title)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	// var User *models.User
	user, err := app.users.Get(post.UserId)
	if err != nil {

		app.serverError(w, err)
		return
	}
	// fmt.Println(app.authenicatedUser(r))
	app.render(w, r, "show.page.html", &templateData{
		Post: post,
		User: user,
	})
}

func (app *application) createPostPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "ok")
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content")
	form.MaxLength("title", 100)

	if !form.Valid() {
		app.render(w, r, "create.page.html", &templateData{Form: form})
		return
	}

	var post models.Post

	post.Title = form.Get("title")
	post.Content = form.Get("content")
	post.Created = time.Now()
	// app.session.
	// fmt.Println(app.session.GetString(r, "userID"))
	post.UserId = app.session.GetString(r, "userID")
	post.Type = form.Get("type")
	// post.Tags = form.Get("tags")
	// post.Status = false

	err = app.posts.Insert(&post)
	if err == models.ErrBadRequest {
		// fmt.Println("I'm here!")
		form.Errors.Add("generic", "Bad request!")
		app.render(w, r, "create.page.html", &templateData{Form: form})
		return
	}
	if err != nil {
		app.serverError(w, err)
	}

	app.session.Put(r, "flash", "Dodano twoje ogłoszenie!")
	http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
}

func (app *application) createPostPageForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", &templateData{
		Form: forms.New(nil),
	})

}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.html", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	form := forms.New(r.PostForm)
	form.Required("nick", "email", "password")
	form.MatchesPattern("email", forms.EmailRx)
	form.MinLength("password", 10)
	// fmt.Println("ok")
	existsEmail := app.users.Exists(form.Get("email"))
	if existsEmail != nil {
		form.ExistsEmail("email")
	}
	existsNick := app.users.Exists(form.Get("nick"))
	if existsNick != nil {
		form.ExistsNick("nick")
	}

	if !form.Valid() {
		app.render(w, r, "signup.page.html", &templateData{
			Form: form,
		})
		return
	}

	var user models.User

	// user.Name = form.Get("name")
	// user.Surname = form.Get("surname")
	user.Nick = form.Get("nick")
	user.Email = form.Get("email")
	password := form.Get("password")
	user.SoftDelete = false

	err = app.users.Insert(&user, password)
	if err != nil {
		app.serverError(w, err)
	}

	app.session.Put(r, "flash", "Twoje konto zostało założone.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.html", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err == models.ErrInvalidCredentials {
		// fmt.Println("I'm here!")
		form.Errors.Add("generic", "Email lub hasło jest nieprawidłowe")
		app.render(w, r, "login.page.html", &templateData{Form: form})
		return
	}

	app.session.Put(r, "userID", id)

	// fmt.Println(app.session.GetString(r, "userID"))

	http.Redirect(w, r, "/", http.StatusSeeOther)

	// fmt.Fprintln(w, "This is loginUser")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "OK")
	// fmt.Println("Nothing!")

	app.session.Remove(r, "userID")

	app.session.Put(r, "flash", "Zostałeś wylogowany!")
	http.Redirect(w, r, "/", 303)
	// fmt.Fprintln(w, "This is logoutUser")
}
