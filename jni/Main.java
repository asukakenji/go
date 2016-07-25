public class Main {

	static {
		// Load "test.dll" on Windows
		// Load "libtest.so" on Linux
		// Load "libtest.jnilib" on Mac
		System.loadLibrary("test");
	}

	private static native void test();

	private static void test2() {
		System.out.println("Hello, world from Java method!");
	}

	public static void main(String[] args) {
		test();
	}

}
