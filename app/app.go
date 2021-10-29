package app

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"os"
	"strings"

	"goweb/web22-1/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New() //전역변수 render.New() 초기화

type AppHandler struct {
	http.Handler //handler http.Handler인데 handler를 생략, 암시적으로 인터페이스를 포함한 멤버 변수를 포함한 상태
	db           model.DBHandler
	db2          model.DBHandler2
	db3          model.DBHandler3
	db4          model.DBHandler4
	db5          model.DBHandler5
}

func getSesssionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "login/login.html", http.StatusTemporaryRedirect)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) getMemberListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetMembers() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) getWorkOutHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db2.GetWorkOutlog() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) getMyDataHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db3.GetMyData() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) getCommunityHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db4.GetCommunity() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) getFileHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db5.GetFile() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addMemberHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	id := r.FormValue("id")
	pswd := r.FormValue("pswd")
	name := r.FormValue("name")
	birth := r.FormValue("birth")
	gender := r.FormValue("gender")
	email := r.FormValue("email")
	area := r.FormValue("area") // js에서 보낸 input value를 name에 추가
	bike_info := r.FormValue("bike_info")
	career := r.FormValue("career")
	club := r.FormValue("club")
	hash := sha512.New()
	hash.Write([]byte(pswd))
	ps := hash.Sum(nil)
	pswd = hex.EncodeToString(ps)
	member := a.db.AddMember(id, pswd, name, birth, gender, email, area, bike_info, career, club) //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusCreated, member)                                                        // JSON으로 member 값을 반환
}

func (a *AppHandler) addCommunityHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	board_id := r.FormValue("board_id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	id := r.FormValue("id")
	date := r.FormValue("date")
	file_id := r.FormValue("file_id")
	good := r.FormValue("good")
	community := a.db4.AddCommunity(board_id, title, content, id, date, file_id, good)
	rd.JSON(w, http.StatusCreated, community)
}

func (a *AppHandler) addFileHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	file_id := r.FormValue("file_id")
	name := r.FormValue("name")
	location := r.FormValue("location")
	size := r.FormValue("size")
	file := a.db5.AddFile(file_id, name, location, size)
	rd.JSON(w, http.StatusCreated, file)
}

type Success struct { //(클라이언트) 응답 결과를 알려주기 위한 구조체
	Success bool `json:"success"`
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) removeMemberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ok := a.db.RemoveMember(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) removeCommunityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	board_id := vars["board_id"]
	ok := a.db4.RemoveCommunity(board_id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) removeFileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file_id := vars["file_id"]
	ok := a.db5.RemoveFile(file_id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) Close() { //새롭게 Close()를 외부에서 만들어 준 것.
	a.db.Close() //model -> a.db로 바꾼다
	a.db2.Close2()
	a.db3.Close3()
	a.db4.Close4()
	a.db5.Close5()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// if request URL is /signin.html, then next()
	if strings.Contains(r.URL.Path, "/signin") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	// if user already signed in
	sessionID := getSesssionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}

	// if not user sign in
	// redirect singin.html
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
		db2:     model.NewDBHandler2(filepath),
		db3:     model.NewDBHandler3(filepath),
		db4:     model.NewDBHandler4(filepath),
		db5:     model.NewDBHandler5(filepath),
	}

	//mebers 핸들러
	r.HandleFunc("/members", a.getMemberListHandler).Methods("GET")
	r.HandleFunc("/members", a.addMemberHandler).Methods("POST")
	r.HandleFunc("/members/{id:[0-9]+}", a.removeMemberHandler).Methods("DELETE")
	//workout 핸들러
	r.HandleFunc("/workouttime", a.getWorkOutHandler).Methods("GET")
	//mydata 핸들러
	r.HandleFunc("/mydata", a.getMyDataHandler).Methods("GET")
	//community 핸들러
	r.HandleFunc("/community", a.getCommunityHandler).Methods("GET")
	r.HandleFunc("/community/file", a.getFileHandler).Methods("GET")
	r.HandleFunc("/community", a.addCommunityHandler).Methods("POST")
	r.HandleFunc("/community/file", a.addFileHandler).Methods("POST")
	r.HandleFunc("/community/{board_id:[0-9]+}", a.removeCommunityHandler).Methods("DELETE")
	r.HandleFunc("/community/{file_id:[0-9]+}", a.removeFileHandler).Methods("DELETE")
	//인덱스 페이지 핸들러
	r.HandleFunc("/", a.indexHandler)

	return a
}
