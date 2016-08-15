public class Main {

	static {
		// Load "gojni.dll" on Windows
		// Load "libgojni.so" on Linux
		// Load "libgojni.jnilib" on Mac
		System.loadLibrary("gojni");
	}

	private static native void test();

	private static void test2() {
		System.out.println("Hello, world from Java method!");
	}

	public static void main(String[] args) {
		test();
	}

}
