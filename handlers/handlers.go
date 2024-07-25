package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	asciiart "ascii-art-web/ascii-art"
)

// Index handles requests to "/" and "/Home"
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/Home", "/home", "/HOME":
		if r.URL.Path == "/" || r.URL.Path == "/Home" || r.URL.Path == "/HOME" {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/index.html")
		}
	case "/aboutus", "/Aboutus", "/ABOUTUS", "/ABOUT US", "/about us":
		if r.URL.Path == "/ABOUTUS" || r.URL.Path == "/Aboutus" || r.URL.Path == "about us" || r.URL.Path == "/ABOUT US" {
			http.Redirect(w, r, "/aboutus", http.StatusMovedPermanently)
		}
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/aboutus.html")
		}
	case "/aboutascii", "/ascii":

		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/aboutascii.html")
		}
	// case "/login", "/Login":
	// 	if r.Method == http.MethodGet {
	// 		serveTemplate(w, "templates/login.html")
	// 	}
	default:
		if r.Method == http.MethodGet {
			http.NotFound(w, r)
		}
	}
}

// HandleASCIIArt handles requests to "/ascii-art"
func HandleASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	bannerFile := r.FormValue("banner")
	input := r.FormValue("input")

	if len(input) == 0 || len(bannerFile) == 0 {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if containsIllegalCharacters(input) {
		http.Error(w, "Input contains illegal characters: Status 400: Bad Request", http.StatusBadRequest)
		return
	}

	bannerFilePath := asciiart.FileName(bannerFile)
	v := asciiart.CheckStatus(bannerFilePath)
	if v != nil {
		http.NotFound(w, r)
		return
	}
	_, errr := os.Stat("templates/result.html")
	if os.IsNotExist(errr) {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	ascii := asciiart.Art(input, asciiart.BannerArt(bannerFilePath))
	e := tmpl.Execute(w, ascii.String())
	if e != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("500 Internal Server Error")
		return
	}
}

// serveTemplate loads and executes a template file
func serveTemplate(w http.ResponseWriter, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		http.Error(w, "404 Page not Found", http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles(filename))
	errr := tmpl.Execute(w, nil)
	if errr != nil {
		log.Println("500 Internal Server Error")
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Function to check for illegal characters in input
func containsIllegalCharacters(input string) bool {
	// Regular expression to match non-printable ASCII characters
	illegalCharRegex := regexp.MustCompile(`[^\x00-\x7F]`)
	return illegalCharRegex.MatchString(input)
}
