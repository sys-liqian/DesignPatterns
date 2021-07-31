package com.demo.design.design.factoryMethod;

/**
 * 工厂方法模式
 * 优点：Person每增加一种子类只需增加一个新的工厂创建无需改动工厂类
 */
public class FactoryMethodTest {

    public static void main(String[] args) {
       StudentFactory sf = new StudentFactory();
       sf.getPerson().talk();
       TeacherFactory tf = new TeacherFactory();
       tf.getPerson().talk();
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

interface PersonFactory {
    Person getPerson();
}

class StudentFactory implements PersonFactory {

    @Override
    public Person getPerson() {
        return new Student();
    }
}

class TeacherFactory implements PersonFactory {

    @Override
    public Person getPerson() {
        return new Teacher();
    }
}

