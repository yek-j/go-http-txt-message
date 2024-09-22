package handlers_test

import (
	"bytes"
	"encoding/json"
	"go-http-txt-message/server/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSend(t *testing.T) {
	// 테스트용 임시 디렉토리 생성
	tmpDir, err := os.MkdirTemp("", "test-messages")
	if err != nil {
		t.Fatalf("임시 디렉토리 생성 실패: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 환경 변수 설정
	os.Setenv("APP_MSG_DIR", tmpDir)
	defer os.Unsetenv("APP_MSG_DIR")

	// 테스트 케이스
	testCases := []struct {
		name           string
		message        handlers.Message
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "유효한 메시지",
			message: handlers.Message{
				Recipient: "user1",
				Sender:    "sender1",
				Title:     "Test",
				Content:   "This is a test message",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "메시지 전송 완료",
		},
		{
			name: "긴 제목",
			message: handlers.Message{
				Recipient: "user2",
				Sender:    "sender2",
				Title:     "ThisTitleIsTooLong",
				Content:   "This is a test message",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "제목은 10이하로만 작성할 수 있습니다.",
		},
		{
			name: "수신자 없음",
			message: handlers.Message{
				Recipient: "",
				Sender:    "sender3",
				Title:     "Test",
				Content:   "This is a test message",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "수신인 이름을 입력하세요.",
		},
		{
			name: "특수문자가 있는 수신자",
			message: handlers.Message{
				Recipient: "user@!4",
				Sender:    "sender4",
				Title:     "Test",
				Content:   "This is a test message",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "수신자 이름이 특수문자를 포함하고 있습니다.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 요청 본문 생성
			body, err := json.Marshal(tc.message)
			if err != nil {
				t.Fatalf("JSON 마샬링 실패: %v", err)
			}	

			// HTTP 요청 생성
			req, err := http.NewRequest("POST", "/send", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("요청 생성 실패: %v", err)
			}

			// 응답 레코더 생성
			rr := httptest.NewRecorder()

			// 핸들러 호출 
			handler := http.HandlerFunc(handlers.Send)
			handler.ServeHTTP(rr, req)

			// 상태 코드 확인 
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("핸들러가 잘못된 상태 코드를 반환: 받은 코드(%v) 받아야하는 코드(%v)", status, tc.expectedStatus)
			}

			// 응답 본문 확인 
			gotBody := strings.TrimSpace(rr.Body.String())
			wantBody := strings.TrimSpace(tc.expectedBody)

			if gotBody != wantBody {
				t.Errorf("핸들러가 잘못된 본문 반환 : 받은 본문(%v) 받아야 하는 본문(%v)", gotBody, wantBody)
			}

			// 파일 생성 확인 (성공 케이스의 경우)
			if tc.expectedStatus == http.StatusOK {
				testFilePath := filepath.Join(tmpDir, tc.message.Recipient, tc.message.Title+"_*.txt")
				matches, err := filepath.Glob(testFilePath)
				if err != nil {
					t.Fatalf("파일 검색 실패: %v", err)
				}
				if len(matches) == 0 {
					t.Errorf("메시지 파일이 생성되지 않았습니다.: %s", testFilePath)
				}
			}
		})
	}
}