package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/erwinhermanto31/crud-image-with-auth/action"
	"github.com/erwinhermanto31/crud-image-with-auth/entity"
	"github.com/erwinhermanto31/crud-image-with-auth/repo/mysql"
	"github.com/erwinhermanto31/crud-image-with-auth/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.InitCon()
	mysql.InitMigration()

	http.HandleFunc("/login", login)
	http.HandleFunc("/upload_image", uploadImage)
	http.HandleFunc("/image/", getImage)
	http.HandleFunc("/images", listImage)
	fmt.Printf("server started at localhost:%v \n", os.Getenv("portApp"))
	http.ListenAndServe(os.Getenv("portApp"), nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	request := entity.Users{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := action.NewLogin().Handler(context.Background(), request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := make(map[string]string)
	resp["token"] = token

	jsonInBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// req.Header.Set("Authorization", "application/json")
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	checkToken, err := utils.ParsingToken(token)
	if err != nil {
		resp["messege"] = "Token not match"

		http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		return
	}
	fmt.Println(checkToken)

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get email from FormFile
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get file nmae
	filename := handler.Filename
	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request := entity.Images{
		UserId:     int(checkToken.ID),
		ImageURL:   filename,
		UploadTime: time.Now(),
	}
	err = action.NewUploadImage().Handler(context.Background(), request)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 403)
		return
	}

	log.Println("filename")
	log.Println(filename)

	resp["messege"] = "Upload Success"

	jsonInBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func getImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)

	urlSplit := strings.Split(r.URL.Path, "/")

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileLocation := filepath.Join(dir, "files", urlSplit[2])
	buf, err := ioutil.ReadFile(fileLocation)

	log.Println(err)
	if err != nil {

		resp := make(map[string]string)
		resp["message"] = "Image not found"

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)

	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", `attachment;filename=`+urlSplit[2])

	w.Write(buf)
}

func listImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// req.Header.Set("Authorization", "application/json")
	token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
	checkToken, err := utils.ParsingToken(token)
	if err != nil {
		resp := make(map[string]string)
		resp["message"] = "Token invalid"

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)

	}

	fullUrl := "http://" + r.Host + "/image/"

	request := entity.Images{
		UserId: int(checkToken.ID),
	}
	listImage, err := action.NewListImage().Handler(context.Background(), request, fullUrl)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 403)
		return
	}
	resp := make(map[string]interface{})
	resp["message"] = "Success Get Data"
	resp["data"] = listImage

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
