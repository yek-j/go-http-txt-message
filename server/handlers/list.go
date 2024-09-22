package handlers

import (
	"encoding/json"
	"go-http-txt-message/server/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func List(w http.ResponseWriter, r *http.Request) {
	// URL에서 user_name 추출 
	path := strings.TrimPrefix(r.URL.Path, "/list/")
	username := strings.TrimSpace(path) // 공백 제거

	if username == "" {
		http.Error(w, "메시지 리스트를 검색할 사용자 이름이 유효하지 않습니다.", http.StatusBadRequest)
		return 
	}

	// 사용자 디렉토리가 있는지 확인
	userExist, _ := utils.IsDirExisting(username)
	if  !userExist {
		http.Error(w, "해당 사용자의 메시지는 존재하지 않습니다.", http.StatusNotFound)
		return 
	}

	msgDir, err := utils.GetDirPath(username)
	if err != nil {
		http.Error(w, "경로를 가져오는 데 실패했습니다.", http.StatusInternalServerError)
		return 
	}


	// 디렉토리 읽기
	files, err := os.ReadDir(msgDir)
	if err != nil {
		http.Error(w, "메시지 목록을 읽는 데 실패했습니다.", http.StatusInternalServerError)
		return 
	}

	var messages []string 
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			messages = append(messages, title)
		}
 	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(messages);
    if err != nil {
        log.Printf("JSON 인코딩 에러: %v", err)
        http.Error(w, "내부 서버 오류", http.StatusInternalServerError)
        return
    }

    // 디버깅을 위한 로그
   	log.Printf("전송된 메시지: %+v", messages)
}