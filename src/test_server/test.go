package main

import (
	//	"io" //　ログ出力設定用。必要に応じて。
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"time"

	"github.com/Chouette2100/srhandler"
)

//Webサーバーを起動する。
func main() {

	//ログファイルの作成
	logfilename := time.Now().Format("20060102") + ".txt"
	logfile, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic("cannnot open logfile: " + logfilename + err.Error())
	}
	defer logfile.Close()

	//	ログ出力を設定する。
	log.SetOutput(logfile)

	//	環境変数 SCRIPT_NAME を取得しrootPathとする。この環境変数が設定されていればCGIとして起動されている。
	//	SCRIPT_NAME が設定されていなければWebサーバーとして起動されている。
	rootPath := os.Getenv("SCRIPT_NAME")
	log.Printf("rootPath: \"%s\"\n", rootPath)

	//	URLに対するハンドラーを定義する。
	http.HandleFunc(rootPath+"/t009top", srhandler.HandlerT009topForm)

	httpport := "8080"

	if rootPath == "" {
		//	Webサーバーとして起動された時
		http.Handle("/", http.FileServer(http.Dir("public")))
		err = http.ListenAndServe(":"+httpport, nil)
	} else {
		//	cgiとして起動された時
		err = cgi.Serve(nil)
	}

	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}
