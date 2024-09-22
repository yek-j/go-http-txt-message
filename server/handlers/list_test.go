package handlers_test

import (
	"encoding/json"
	"go-http-txt-message/server/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestList(t *testing.T) {
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
		setupFiles     []string
		expectedStatus int
		expectedFiles  []string
	}{
		{
			name:           "메시지가 있는 사용자",
			username:       "testuser1",
			setupFiles:     []string{"message1_20240101.txt", "message2_20240102.txt"},
			expectedStatus: http.StatusOK,
			expectedFiles:  []string{"message1_20240101", "message2_20240102"},
		},
		{
			name:           "메시지가 없는 사용자",
			username:       "testuser2",
			setupFiles:     []string{},
			expectedStatus: http.StatusOK,
			expectedFiles:  []string{},
		},
		{
			name:           "존재하지 않는 사용자",
			username:       "nonexistent",
			setupFiles:     []string{},
			expectedStatus: http.StatusNotFound,
			expectedFiles:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedFiles != nil {
				userDir := filepath.Join(tmpDir, tc.username)
				os.MkdirAll(userDir, os.ModePerm)
				for _, file := range tc.setupFiles {
					os.WriteFile(filepath.Join(userDir, file), []byte("test"), 0644)
				}
			}
			
			req, err := http.NewRequest("GET", "/list/"+tc.username, nil)
			if err != nil {
				t.Fatalf("요청 생성 실패: %v", err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(handlers.List)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("핸들러가 잘못된 상태 코드를 반환: 받은 코드(%v) 받아야하는 코드(%v)", status, tc.expectedStatus)
			}

			if tc.expectedStatus == http.StatusOK {
				var files []string
				err = json.NewDecoder(rr.Body).Decode(&files)
				if err != nil {
					t.Fatalf("Body 디코딩 실패 : %v", err)
				}

				if len(files) != len(tc.expectedFiles) {
					t.Errorf("핸들러가 잘못된 파일 갯수를 반환 : 받은 갯수(%v) 받아야 하는 갯수(%v)", len(files), len(tc.expectedFiles))
				}

				for i, file := range files {
					if file != tc.expectedFiles[i] {
						t.Errorf("핸들러가 잘못된 파일 이름 반환 : 받은 파일(%v) 받아야 하는 파일(%v)", file, tc.expectedFiles[i])
					}
				}
			}
		})
	}

}