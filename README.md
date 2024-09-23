# Text Message Server & Client

## 프로젝트 개요
이 프로젝트는 HTTP 서버 및 클라이언트 구현을 학습하면서 텍스트 메시지를 생성하고 저장하는 간단한 서버를 구축하는 것을 목표로 합니다. Go 언어를 사용하여 개발되었으며, 기본적인 HTTP 요청 처리, 파일 시스템 조작, 그리고 JSON 데이터 처리 등의 개념을 실습합니다.

## 주요 기능
1. 텍스트 메시지 생성 및 저장
2. HTTP 요청 처리
3. JSON 데이터 파싱
4. 텍스트 파일 읽기
5. 디렉토리의 파일 리스트 읽기

## API 엔드포인트

### 1. Send Message API

#### 엔드포인트
`POST /send`

#### 설명
이 API는 클라이언트로부터 메시지 정보를 받아 텍스트 파일로 저장합니다.

#### 요청 형식
```json
{
  "Recipient": "수신자이름",
  "Sender": "발신자이름",
  "Title": "메시지제목",
  "Content": "메시지내용"
}
```

#### 응답
- 성공: 200 OK와 함께 "메시지 전송 완료" 메시지 반환
- 실패: 적절한 오류 상태 코드와 오류 메시지 반환

### 2. List Message API

#### 엔드포인트
`GET /list/{user_name}`

#### 설명
이 API는 특정 사용자의 모든 메시지 목록을 반환한다. 

#### 응답
- 성공: 200 OK와 함께 메시지 목록을 배열로 반환
- 실패: 적절한 오류 상태 코드와 오류 메시지 반환

### 3. Get Message API

#### 엔드포인트
`GET /message?username={username}&title={message_title}`

#### 설명
이 API는 특정 사용자의 특정 메시지 내용을 반환한다.

#### 응답
- 성공: 200 OK와 함께 메시지 내용을 반환
- 실패: 적절한 오류 상태 코드와 오류 메시지 반환

## 주요 학습 포인트
1. HTTP 서버 구현: `net/http` 패키지를 사용한 기본적인 HTTP 서버 설정
2. 요청 처리: GET 및 POST 요청 데이터 파싱 및 처리
3. JSON 처리: `encoding/json` 패키지를 이용한 JSON 데이터 인코딩 및 디코딩
4. 파일 시스템 조작: `os` 패키지를 사용한 디렉토리 생성, 파일 읽기 및 쓰기
5. 에러 처리: HTTP 응답 코드를 통한 다양한 에러 상황 처리
6. 환경 변수 사용: `os.Getenv`를 통한 설정 관리
7. URL 파라미터 및 쿼리 문자열 처리

## 코드 구조
- `server.go`: 메인 서버 설정 및 라우팅
- `handlers/send.go`: 메시지 전송 핸들러 및 관련 로직
- `handlers/list.go`: 특정 사용자 메시지 목록 조회 핸들러
- `handlers/getmessage.go`: 특정 사용자의 특정 메시지를 조회하는 핸들러
- `utils/paths.go`: 디렉토리 경로와 파일 경로관련 로직

## 설정
- `APP_MSG_DIR` 환경 변수: 메시지 저장 디렉토리 지정 (설정되지 않은 경우 기본 경로 사용)

## 클라이언트
클라이언트는 명령줄 인터페이스(CLI)를 통해 서버와 상호 작용 한다. 다음과 같은 명령어를 지원한다.
- `send` : 새 메시지 전송
- `list` : 특정 사용자의 메시지 목록 조회
- `message` : 특정 사용자의 특정 메시지 내용 조회

### 클라이언트 전송 예시
1. 메시지 전송:
   ```
   ./client send
   ```
   프롬프트에 따라 수신자, 발신자, 제목, 내용을 입력합니다.

2. 메시지 목록 조회:
   ```
   ./client list -u username
   ```

3. 특정 메시지 조회:
   ```
   ./client message -u username -t "message title"
   ```
