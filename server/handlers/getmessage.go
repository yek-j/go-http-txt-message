package handlers

import (
	"go-http-txt-message/server/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	// URL에서 message_title 추출
	username := strings.TrimSpace(r.URL.Query().Get("username"))
	title := strings.TrimSpace(r.URL.Query().Get("title")) 

	if username == "" || title == "" {
		http.Error(w, "메시지 검색을 위해 사용자 이름과 제목이 필요합니다.", http.StatusBadGateway)
		return 
	}

	// title에 .txt 확장자 추가
	if filepath.Ext(title) != ".txt" {
		title += ".txt"
	}

	filePath, err := utils.GetFilePath(username, title)
	if err != nil {
		http.Error(w, "경로를 가져오는데 실패했습니다.", http.StatusInternalServerError)
		return 
	}


	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "메시지 내용을 읽는데 실패했습니다.", http.StatusInternalServerError)
		return 
	}

	// 응답 설정
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(content)
	log.Printf("%s의 %s 메시지 내용 : %s", username, title, content)
}