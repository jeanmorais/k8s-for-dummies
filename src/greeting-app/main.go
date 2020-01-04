package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var greetings = map[string]string{
	"PT": "Olá",
	"EN": "Hello",
	"ES": "Hola",
	"FR": "Bonjour",
	"IT": "Ciao",
}

const (
	envAppLanguage = "APP_LANGUAGE"
)

func main() {
	checkRequired(envAppLanguage)
	handler := http.NewServeMux()
	handler.HandleFunc("/", Greet)
	http.ListenAndServe(":8080", handler)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	if v, ok := greetings[getAppLanguage()]; ok {
		log.Println(time.Now(), "Message received")
		fmt.Fprintf(w, v+"! =)")
		return
	}
	http.Error(w, "wtf?", http.StatusNotFound)
	return
}

func checkRequired(envVarArgs ...string) {
	for _, envVar := range envVarArgs {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Environment variable [%s] has not been set", envVar)
		}
	}
}

func getAppLanguage() string {
	return strings.ToUpper(os.Getenv(envAppLanguage))
}
