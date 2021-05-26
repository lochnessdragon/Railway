package main

import (
  "html/template"
  "net/http"
  "log"
  "io/ioutil"
  "fmt"
)

const template_dir = "tmpl/"
const data_dir = "data/"

var templates *template.Template

type User struct {
  Name string
  Id uint32 
  AccountType uint8
}

func read_file(filename string) ([]byte, error) {
  //log.Printf("Attempting to read: %s", filename)
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return content, nil
}

func render_resource(w http.ResponseWriter, resource string) {
  file, err := read_file(data_dir + resource)
  if err != nil {
    //log.Warn("Could not find resource file")
    http.Error(w, err.Error(), http.StatusNotFound)
    return 
  }
  fmt.Fprintf(w, "%s", file)
}

func render_template(w http.ResponseWriter, tmpl string, data interface{}) {
  err := templates.ExecuteTemplate(w, tmpl+".html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func default_handler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[1:] // grab url
  switch url {
    case "":
      u := User{Name: "Lochnessdragon", Id: 1, AccountType: 1}
      render_template(w, "main", u)
    default:
      //log.Printf("Loading resource: %s", url)
      render_resource(w, url)
  }

}

func login_handler(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path[1:]

  switch r.Method {
  case http.MethodPost:
    //log.Print("Method post recieved")
    //http.Redirect(w, r, "/", http.StatusFound)
    
  case http.MethodGet:
    if url != "login/" {
      log.Print("Could not find url: " + url)
      http.Error(w, "Could not find requested resource", http.StatusNotFound)
      return 
    } else {
      render_resource(w, "login.html")
    }
  default:
    http.Error(w, "Unsupported method", http.StatusForbidden)
    return
  }
}

func setup_template_cache(dir string) {
  var files = get_file_list(dir)
  templates = template.Must(template.ParseFiles(files...))
}

func main() {
  setup_template_cache(template_dir)

  http.HandleFunc("/", default_handler)
  http.HandleFunc("/login/", login_handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}