import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class ConcurrentExample {
    public static void main(String[] args) {
        ExecutorService executorService = Executors.newFixedThreadPool(2);
        executorService.submit(new MessagePrinter("Hello from Task 1"));
        executorService.submit(new MessagePrinter("Hello from Task 2"));
        executorService.shutdown();
    }
}

// A class that implements Runnable to define the task
class MessagePrinter implements Runnable {
    private String message;
    public MessagePrinter(String message) {
        this.message = message;
    }

    @Override
    public void run() {
        for (int i = 0; i < 5; i++) {
            System.out.println(message);
            try {
                Thread.sleep(1000); // Pause for 1 second
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
