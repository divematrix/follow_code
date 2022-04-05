package ch15.ex1502;

public class LambdaEx4 {
	public static void main(String[] args) {
		LambdaInterface4 f4 = (x, y) -> {
			return x * y;
		};
		System.out.println("두 수의 곱 : " + f4.cal(3, 2));

		f4 = (x, y) -> x + y;
		System.out.println("두 수의 합 : " + f4.cal(3, 2));

		f4 = (x, y) -> {
			return x / y;
		};
		System.out.println("두 수의 몫 : " + f4.cal(5, 2));

		f4 = (x, y) -> x % y;
		System.out.println("두 수의 나머지 : " + f4.cal(5, 2));

		f4 = (x, y) -> sum(x, y);
		System.out.println("두 수의 합(sum()) : " + f4.cal(3, 2));
	}

	static int sum(int x, int y) {
		return x + y;
	}
}

// 매개변수도 있고 리턴타입도 있는 경우의 예
@FunctionalInterface
interface LambdaInterface4 {
	int cal(int x, int y);
}