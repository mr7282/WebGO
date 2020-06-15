package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"html/template"
	"io/ioutil"
	// "log"
	"net/http"
	// "strconv"
	"myBlog/models"
)

func (serv *Server) indexTpl(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/index.html"))

	ctx := context.Background()

	type OnePost struct {
		UserName  string
		CreatedAt string
		Post      models.Post
	}

	showAll, err := models.Posts(qm.Select(models.PostColumns.ID, models.PostColumns.Name, models.PostColumns.Post, models.PostColumns.UserID, models.PostColumns.CreatedAt)).All(ctx, serv.db)
	if err != nil {
		serv.lg.WithError(err).Infoln("can't get all posts")
	}

	var arrPost []OnePost
	for _, d := range showAll {
		nameuser, err := d.User(qm.Select(models.UserColumns.Name)).One(ctx, serv.db)
		if err != nil {
			serv.lg.WithError(err).Warningln("can't show user")
		}

		dOnePost := OnePost{
			UserName:  nameuser.Name,
			CreatedAt: d.CreatedAt.Format("02-01-2006 15:04:05"),
			Post: models.Post{
				ID:   d.ID,
				Name: d.Name,
				Post: d.Post,
			},
		}

		arrPost = append(arrPost, dOnePost)
	}

	var allPostView = struct {
		Title string
		Slice []OnePost
	}{
		Title: "myBlog",
		Slice: arrPost,
	}

	boil.DebugMode = true

	if err := tmpl.ExecuteTemplate(wr, "Blog", allPostView); err != nil {
		serv.lg.WithError(err).Infoln("can't show all posts")
	}

}

func (serv *Server) editView(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/edit.html"))
	IDInput := r.PostFormValue("ID")
	NameInput := r.PostFormValue("Name")
	PostInput := r.PostFormValue("Post")

	type Posts struct {
		ID   string
		Name string
		Post string
	}

	Post := Posts{
		ID: IDInput,
		Name: NameInput,
		Post: PostInput,
	}

	if err := tmpl.ExecuteTemplate(wr, "Blog", Post); err != nil {
		serv.lg.WithError(err).Warningln("can't show edit tepmlate")
	}
}

func (serv *Server) editPost(wr http.ResponseWriter, r *http.Request) {
	readResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serv.lg.WithError(err).Warningln("can't read response from edit.html")
	}
	defer r.Body.Close()

	ctx := context.Background()

	decodingResponse := models.Post{}

	if err := json.Unmarshal(readResponse, &decodingResponse); err != nil {
		serv.lg.WithError(err).Warningln("can't decoding response from edit.html")
	}

	post, err := models.FindPost(ctx, serv.db, decodingResponse.ID, models.PostColumns.ID, models.PostColumns.Name, models.PostColumns.Post)
	if err != nil {
		serv.lg.WithError(err).Warningln("can't find post")
	}
	post.Name = decodingResponse.Name
	post.Post = decodingResponse.Post

	fmt.Println("IDInput^ ", post)
	fmt.Printf("%T\n", post)
	_, err = post.Update(ctx, serv.db, boil.Whitelist(models.PostColumns.Name, models.PostColumns.Post))
	if err != nil {
		serv.lg.WithError(err).Warningln("can't update this Post")
	}
}

// func (serv *Server) findPost(wr http.ResponseWriter, r *http.Request) {

// 	readResponse, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		serv.lg.WithError(err).Warningln("can't read response from the browser")
// 	}

// 	var searchTarget string
// 	if err := json.Unmarshal(readResponse, &searchTarget); err != nil {
// 		serv.lg.WithError(err).Warningln("can't deencoding response from the browser")
// 	}
// }

func (serv *Server) createView(wr http.ResponseWriter, r *http.Request) {
 var tmpl = template.Must(template.New("MyBlog").ParseFiles("./www/templates/create.html"))

 if err := tmpl.ExecuteTemplate(wr, "Blog", nil); err != nil {
	 serv.lg.WithError(err).Warningln("can't show create.html")
 }
}

func (serv *Server) createPost(wr http.ResponseWriter, r *http.Request) {
	readResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serv.lg.WithError(err).Warningln("can't read response from create.html")
	}
	defer r.Body.Close()

	ctx := context.Background()

	decodingResponse := models.Post{}
	decodingResponse.UserID = 3


	if err := json.Unmarshal(readResponse, &decodingResponse); err != nil {
		serv.lg.WithError(err).Warningln("can't unmarshal response from create.html")
	}

	if err = decodingResponse.Insert(ctx, serv.db, boil.Whitelist(models.PostColumns.Name, models.PostColumns.Post, models.PostColumns.UserID)); err!= nil {
		serv.lg.WithError(err).Warningln("can't insert new row in database")
	}
}
