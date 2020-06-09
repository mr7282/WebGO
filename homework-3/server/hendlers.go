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
	// conv, err := strconv.Atoi(r.PostFormValue("ID"))
	// if err != nil {
	// 	serv.lg.WithError(err).Warningln("can't conversion IDInput")
	// }

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
		serv.lg.WithError(err).Warningln("can't read response from browser")
	}
	defer r.Body.Close()

	ctx := context.Background()

	encodingResponse := models.Post{}

	if err := json.Unmarshal(readResponse, &encodingResponse); err != nil {
		serv.lg.WithError(err).Warningln("can't decoding response from browser")
	}

	post, err := models.FindPost(ctx, serv.db, encodingResponse.ID, models.PostColumns.ID, models.PostColumns.Name, models.PostColumns.Post)
	if err != nil {
		serv.lg.WithError(err).Warningln("can't find post")
	}
	post.Name = encodingResponse.Name
	post.Post = encodingResponse.Post

	fmt.Println("IDInput^ ", post)
	fmt.Printf("%T\n", post)
	_, err = post.Update(ctx, serv.db, boil.Whitelist(models.PostColumns.Name, models.PostColumns.Post))
	if err != nil {
		serv.lg.WithError(err).Warningln("can't update this Post")
	}
}

func (serv *Server) findPost(wr http.ResponseWriter, r *http.Request) {

	readResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serv.lg.WithError(err).Warningln("can't read response from the browser")
	}

	var searchTarget string
	if err := json.Unmarshal(readResponse, &searchTarget); err != nil {
		serv.lg.WithError(err).Warningln("can't deencoding response from the browser")
	}

	// serv.lg.Infoln(readResponse)
	// serv.lg.Infoln(searchTarget)
}
