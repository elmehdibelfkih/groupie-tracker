package internal

import (
	"groupie-tracker/error"
	"html/template"
	"log"
)

func InitTmpl() {
	tmp, err := template.ParseFiles(INDEX_PATH)
	if err != nil {
		log.Fatal(err)
	}
	INDEX_TMPL = template.Must(tmp, err)

	tmp1, err1 := template.ParseFiles(error.ERROR_TMPL_PATH)
	if err1 != nil {
		log.Fatal(err1)
	}
	error.ERRORTMPL = template.Must(tmp1, err1)

	tmp2, err2 := template.ParseFiles(RELATIONS_PATH)
	if err2 != nil {
		log.Fatal(err2)
	}
	RELATIONS_TMPL = template.Must(tmp2, err2)

	tmp3, err3 := template.ParseFiles(DATES_PATH)
	if err3 != nil {
		log.Fatal(err3)
	}
	DATES_TMPL = template.Must(tmp3, err3)

	tmp4, err4 := template.ParseFiles(LOCATIONS_PATH)
	if err4 != nil {
		log.Fatal(err4)
	}
	LOCATIONS_TMPL = template.Must(tmp4, err4)
}
