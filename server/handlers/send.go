package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Message struct {
    Recipient string `json:"Recipient"`
    Sender    string `json:"Sender"`
    Title   string `json:"Title"`
    Content   string `json:"Content"`
}

var specialCharPattern = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>?]+`)

func Send(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)

	if err != nil {
		http.Error(w, "잘못된 JSON 요청입니다.", http.StatusBadRequest)
		return
	}

	// 검증
	validateMsg := validateRequest(msg);

	if validateMsg != "" {
		http.Error(w, validateMsg, http.StatusBadRequest)
	}
	
	msgDir := os.Getenv("APP_MSG_DIR")
	if msgDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, "경로를 가져오는 데 실패했습니다", http.StatusInternalServerError)
            return
		}
		msgDir = filepath.Join(wd, "message", msg.Recipient)
	} else {
		msgDir = filepath.Join(msgDir, msg.Recipient)
	}
	
	// 수신자 디렉토리가 없으면 생성하고 메시지 파일 생성
	err = os.MkdirAll(msgDir, os.ModePerm)
	if err != nil {
		http.Error(w, "수신자 디렉토리 생성 오류", http.StatusInternalServerError)
		return
	}

	// 메시지 파일 생성
	now := time.Now()
	filename := fmt.Sprintf("%s_%s.txt", msg.Title, now.Format("20060102"))
	fullPath := filepath.Join(msgDir, filename)

	contentWithSender := fmt.Sprintf("보낸 사람: %s\n\n%s", msg.Sender, msg.Content)


	err = os.WriteFile(fullPath, []byte(contentWithSender), 0644)
	if err != nil {
		http.Error(w, "메시지 파일 생성 오류", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "메시지 전송 완료")
}

func validateRequest(m Message) string {
	if len(m.Title) > 10 {
		return "제목은 10이하로만 작성할 수 있습니다."
	}

	if m.Recipient == "" {
		return "수신인 이름을 입력하세요."
	}

	if specialCharPattern.MatchString(m.Recipient) {
		return "수신자 이름이 특수문자를 포함하고 있습니다."
	}

	return ""
}