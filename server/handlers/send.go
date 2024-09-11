package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Message struct {
    Recipient string `json:"Recipient"`
    Sender    string `json:"Sender"`
    Title   string `json:"Title"`
    Content   string `json:"Content"`
}

var specialCharPattern = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>?]+`)

func Send(w http.ResponseWriter, r *http.Request) {
	
	
	fmt.Fprintf(w, "Send Handler")

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

	// 수신자 폴더 찾기 (임시 경로)
	_, findErr := findOrCreateDir(msg.Recipient, "go-http-txt-message\\message")
	
	if findErr != nil {
		// msg.Content_date

	}
}

func findOrCreateDir(recipient string, root string) (bool, error) {
	// 임시 경로
	cwd, err := os.Getwd()
	if err != nil {
		return false, fmt.Errorf("현재 작업 디렉터리 가져오기 실패: %v", err)
	}
	fullPath := filepath.Join(cwd, "..", "..", root, recipient)

	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) { // 디렉터리가 없는 경우
		log.Printf("생성할 경로 : %s", fullPath)
		err = os.Mkdir(fullPath, 0775)
		if err != nil {
			return false, fmt.Errorf("디렉토리 생성 중 오류 발생 : %v", err)
		}
		return true, nil // 디렉터리 생성 완료
	} else if err != nil {
		return false, fmt.Errorf("디렉토리 상태 확인 오류 : %v", err)
	}

	return true, nil // 디렉터리 이미 존재
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