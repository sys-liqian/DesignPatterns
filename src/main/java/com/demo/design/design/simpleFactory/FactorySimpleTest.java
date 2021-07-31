package com.demo.design.design.simpleFactory;

/**
 * 简单工厂模式，也叫静态工厂方法模式
 * 缺点：每增加一个Person子类就得修改工厂类
 */
public class FactorySimpleTest {

    public static void main(String[] args) {
        Person person = PersonFactory.getPerson(1);
        person.talk();
        Person person1 = PersonFactory.getPerson(2);
        person1.talk();
    }
}

interface Person {
    void talk();
}

class Student implements Person {

    @Override
    public void talk() {
        System.out.println("我是一个学生！");
    }
}

class Teacher implements Person {

    @Override
    public void talk() {
        System.out.println("我是一个老师!");
    }
}

class PersonFactory {

    public static Person getPerson(int type) {
        switch (type) {
            case 1:
                return new Student();
            case 2:
                return new Teacher();
        }
        return null;
    }
}
