package handlers_test

import (
	"go-http-txt-message/server/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test-messages")
	if err != nil {
		t.Fatalf("임시 디렉토리 생성 실패: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	os.Setenv("APP_MSG_DIR", tmpDir)
	defer os.Unsetenv("APP_MSG_DIR")
	
	testCases := []struct {
		name           string
		username       string
		title          string
		setupContent   string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "사용자와 메시지 유효",
			username:       "testuser1",
			title:          "message1",
			setupContent:   "Hello, this is message 1",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, this is message 1",
		},
		{
			name:           "메시지만 존재하지 않음",
			username:       "testuser1",
			title:          "nonexist",
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "사용자가 없음",
			username:       "nonexistuser",
			title:          "message1",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty username",
			username:       "",
			title:          "message1",
			expectedStatus: http.StatusBadGateway,
		},
		{
			name:           "Empty message title",
			username:       "testuser1",
			title:          "",
			expectedStatus: http.StatusBadGateway,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupContent != "" {
				userDir := filepath.Join(tmpDir, tc.username)
				os.MkdirAll(userDir, os.ModePerm)
				filename := tc.title 
				if filepath.Ext(filename) != ".txt" {
					filename += ".txt"
				}
				err := os.WriteFile(filepath.Join(userDir, filename), []byte(tc.setupContent), 0644)
				if err != nil {
					t.Fatalf("테스트 파일 생성 실패 : %v", err)
				}
			}

			req, err := http.NewRequest("GET", "/message?username="+tc.username+"&title="+tc.title, nil)
			if err != nil {
				t.Fatalf("요청 생성 실패 : %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetMessage)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("핸들러가 잘못된 상태 코드를 반환: 받은 코드(%v) 받아야하는 코드(%v)", status, tc.expectedStatus)
			}

			if rr.Code == http.StatusOK {
				if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tc.expectedBody) {
					t.Errorf("핸들러가 잘못된 본문 반환 : 받은 본문(%v) 받아야 하는 본문(%v)", rr.Body.String(), tc.expectedBody)
				}
			}
		})
	}
}