package com.demo.design.design.singleton;

/**
 * 懒汉，线程安全，延迟加载
 * 单例加锁只需要第一次初始化时，方法加锁影响性能
 */
class SingletonA {
    private SingletonA() {
    }

    private static SingletonA instance;

    public static synchronized SingletonA getInstance() {
        if (instance == null) {
            instance = new SingletonA();
        }
        return instance;
    }
}

/**
 * 饿汉，线程安全，非延迟加载
 * 简单粗暴，以空间换时间
 */
class SingletonB {
    private SingletonB() {
    }

    private static SingletonB instance = new SingletonB();

    public static SingletonB getInstance() {
        return instance;
    }
}

/**
 * 双重校验，线程安全，延迟加载
 */
class SingletonC {
    private SingletonC() {
    }

    /**
     * volatile 关键字不可少，禁止指令重排序
     * new 一个对象，3个步骤
     * 1.分配内存空间
     * 2.初始化对象
     * 3.对象指向内存空间
     * jvm 可能对指令重排序，变成先将对象指向内存空间，再初始化对象
     * 可能出现，线程1还未完成初始化对象，线程2，由于instance已经不为null
     * 返回了未初始化完成的对象
     */
    private volatile static SingletonC instance;

    public static SingletonC getInstance() {
        if (instance == null) {
            //可能有多个线程同时进入此判断
            synchronized (SingletonC.class) {
                //此处加了类锁，同一时间只有一个线程进入，其他线程进入SingletonC class对象的锁池
                if (instance == null) {
                    //第一个线程可以到达此处，初始化对象，其他线程因为instance已经不为null,直接返回
                    instance = new SingletonC();
                }
            }
        }
        return instance;
    }
}

/**
 * 静态内部类，线程安全，延迟加载
 */
class SingletonD {
    private SingletonD() {
    }

    private static class SingletonHolder {
        private static final SingletonD INSTANCE = new SingletonD();
    }

    public static SingletonD getInstance() {
        return SingletonHolder.INSTANCE;
    }
}

public class SingletonTest {

    public static void main(String[] args) {
        for (int i = 0; i < 100; i++) {
            Thread t = new Thread(() -> {
                SingletonA a = SingletonA.getInstance();
                System.out.println(a.hashCode());
            });
            t.start();
        }
    }

}
