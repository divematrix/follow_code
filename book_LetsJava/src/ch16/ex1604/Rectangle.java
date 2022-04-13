package ch16.ex1604;

public class Rectangle extends Shape {
	int w, h;

	public Rectangle () {
		this(1, 1);
	}

	public Rectangle (int w, int h) {
		this.w = w;
		this.h = h;
	}

	@Override
	public double area () {
		return w * h;
	}

	@Override
	public double length () {
		return (w + h) * 2;
	}

	@Override
	public String toString () {
		return "넓이 : " + this.area();
	}
}
