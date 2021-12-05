package com.demo.design.io;

import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;

public class BIOTestCase {

    public static class NameThreadFactory implements ThreadFactory {
        private final AtomicInteger threadNum = new AtomicInteger(1);

        @Override
        public Thread newThread(Runnable r) {
            Thread t = new Thread(r, "thread-" + threadNum.getAndIncrement());
            return t;
        }
    }

    static class IoHandler extends Thread {
        private Socket socket;
        public IoHandler(Socket socket) {
            this.socket = socket;
        }
        @Override
        public void run() {
            while (!Thread.currentThread().isInterrupted() &&
                    !socket.isClosed()) {
                try {
                    //读数据
                    BufferedReader br = new BufferedReader(
                            new InputStreamReader(socket.getInputStream()));
                    String s = br.readLine();
                    //......处理数据
                    //返回数据
                    BufferedWriter bw = new BufferedWriter(
                            new OutputStreamWriter(socket.getOutputStream()));
                    bw.write(s);
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    public static void main(String[] args) throws IOException {
        ExecutorService executor = new ThreadPoolExecutor(12,
                12,
                0,
                TimeUnit.SECONDS,
                new ArrayBlockingQueue<Runnable>(12),
                new NameThreadFactory());
        ServerSocket serverSocket = new ServerSocket(8080);
        while (!Thread.currentThread().isInterrupted()) {
            Socket socket = serverSocket.accept();
            executor.submit(new IoHandler(socket));
        }
    }

}
