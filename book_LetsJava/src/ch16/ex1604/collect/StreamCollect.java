package ch16.ex1604.collect;

import ch16.ex1604.Circle;
import ch16.ex1604.Rectangle;
import ch16.ex1604.Shape;

import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class StreamCollect {
	public static void main (String[] args) {

		Shape s1 = new Rectangle(10, 3);
		Shape s2 = new Circle(10);
		Shape s3 = new Rectangle(20, 2);
		Shape s4 = new Circle(11);

		List<Shape> list = Arrays.asList(s1, s2, s3, s4);

		// collect() 메서드를 이용하면, 스트림에서 요소들을 필터링 하거나 매핑한 후 새로운 객체를 생성 할 수 있다.

		// 요소가 Rectangle 객체인 경우 collect 메서드로 List 으로 변환
		List<Shape> rectList = list.stream().filter(s -> s instanceof Rectangle).collect(Collectors.toList());
		rectList.stream().forEach(System.out::println);
		System.out.println();

		// 요소가 Rectangle 객체인 경우 collect 메서드로 Set 으로 변환
		Set<Shape> rectSet = list.stream().filter(s -> s instanceof Rectangle).collect(Collectors.toSet());
		rectSet.stream().forEach(System.out::println);
	}
}
