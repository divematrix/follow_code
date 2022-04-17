## ch17 숫자 맞추기 게임 만들기
	1. 랜덤한 숫자 생성
	2. 입력받기 - 숫자값
	3. 입력값과 랜덤값 비교
	4. 비교 메세지 출력
	5. 숫자를 맞추면 종료

### 1. 랜덤한 숫자 생성
	1.1 사전지식 - math/rand 패키지
		1) 초기값 - 유사랜덤값은 초기값이 같으면 매번 똑같은 값이 생성
		2) 초기값 설정 : func Seed (seed int64)
		3) 유사랜덤값 생성 - func intn (n int) int
			0이상 n미만 유사 랜덤 정수 반환
	1.2 사전지식 - time 패키지
		- 날짜, 시각, 시간 등을 다루는 패키지
		- 매번 다른 랜덤값을 얻기위해 랜덤 시드값을 설정해줘야 함
		- 가장 많이 사용되는 방법은 현재 시각 값을 설정 해주는 것
		1) 현재 시각 : func Now() Time
			현재 시각 Time 객체 반환
		2) func (t Time) UnixNano() int64
			UTC 시간 기준인 1970년 1월 1일부터
			Time 객체가 나타내는 시각까지 경과한 시간을
			나노초 단위로 나타낸 값을 반환