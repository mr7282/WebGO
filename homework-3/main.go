package main

import(
	"net/http"
	"log"
	"html/template"
)


// 1. Создайте роут и шаблон для отображения всех постов в блоге.
// 2. Создайте роут и шаблон для просмотра конкретного поста в блоге.
// 3. Создайте роут и шаблон для редактирования и создания материала.
// 4. * Добавьте к роуту редактирования и создания материала работу с Markdown с помощью пакета blackfriday.
// Рекомендуем хранить контент поста в блоге в типе template.HTML, чтобы использовать html-разметку внутри поста (для blackfriday это обязательное условие корректного отображения материала).


type BlogList struct {
	Name string
	Blog []Post
}

type Post struct {
	Id int
	Name string
	Body string
}

var myBlog = BlogList{
	Name: "Мой Блог",
	Blog: []Post{
		Post{1, "Мой первый пост", "Далее создадим глобальную переменную, опишем в ней простой лист, подготовим переменную, в которую будет считываться шаблон при запуске приложения, и создадим роутер, который будет отдавать нам страницу со списком."},
		Post{2, "Мой второй пост", "{{year}} — вызов функций шаблона происходит по названию, без указания точки перед ее именем. Функции необходимо добавлять к структуре *Template перед тем, как производить чтение шаблона. "},
		Post{3, "Мой третий пост", "В шаблон можно встроить общие функции. Например, на сайтах в конце страницы часто указывают копирайт, и в этой строке — текущий год. Создадим функцию, которая будет возвращать в шаблон текущий год."},

	},
}

var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/tmpl.html"))

func main() {
	route := http.NewServeMux()
	route.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	route.HandleFunc("/", viewBlog)

	log.Fatal(http.ListenAndServe(":8080", route))
}

func viewBlog(wr http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}