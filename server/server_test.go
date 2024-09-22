package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go-http-txt-message/server/handlers"
)

func TestServer(t *testing.T) {
	// 테스트용 임시 디렉토리 생성
	tmpDir, err := os.MkdirTemp("", "test-messages")
	if err != nil {
		t.Fatalf("임시 디렉토리 생성 실패: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 환경 변수 설정
	os.Setenv("APP_MSG_DIR", tmpDir)
	defer os.Unsetenv("APP_MSG_DIR")

	// 서버 생성
	mux := http.NewServeMux()
	mux.HandleFunc("/send", handlers.Send)
	mux.HandleFunc("/list/", handlers.List)
	mux.HandleFunc("/message/", handlers.GetMessage)

	server := httptest.NewServer(mux)
	defer server.Close()

	now := time.Now()

	// 테스트 케이스
	t.Run("Send Message", func(t *testing.T) {
		message := handlers.Message{
			Recipient: "testuser",
			Sender:    "sender",
			Title:     "TestTitle",
			Content:   "This is a test message",
		}
		body, _ := json.Marshal(message)
		resp, err := http.Post(server.URL+"/send", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("요청 실패: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("예상치 못한 상태 코드: 받음 %v, 원함 %v", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("List Messages", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/list/testuser")
		if err != nil {
			t.Fatalf("요청 실패: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("예상치 못한 상태 코드: 받음 %v, 원함 %v", resp.StatusCode, http.StatusOK)
		}

		var messages []string
		json.NewDecoder(resp.Body).Decode(&messages)

		mname := fmt.Sprintf("TestTitle_%s", now.Format("20060102"))		
		if len(messages) != 1 || messages[0] != mname {
			t.Errorf("예상치 못한 메시지 리스트: %v", messages)
		}
	})

	t.Run("Get Message", func(t *testing.T) {
		filename := fmt.Sprintf("TestTitle_%s.txt", now.Format("20060102"))
		resp, err := http.Get(server.URL + "/message?username=testuser&title="+filename)
		
		if err != nil {
			t.Fatalf("요청 실패: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("예상치 못한 상태 코드: 받음 %v, 원함 %v", resp.StatusCode, http.StatusOK)
		}
		content, _ := os.ReadFile(filepath.Join(tmpDir, "testuser", filename))
		if string(content) != "보낸 사람: sender\n\nThis is a test message" {
			t.Errorf("예상치 못한 메시지 내용: %s", content)
		}
	})
}