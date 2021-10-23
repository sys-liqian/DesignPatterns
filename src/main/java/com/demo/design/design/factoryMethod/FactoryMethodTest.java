package com.demo.design.design.factoryMethod;

/**
 * 工厂方法模式
 * 优点：Person每增加一种子类只需增加一个新的工厂创建无需改动工厂类
 * 缺点：工厂方法只能产出同种对象（都是Person类型）
 * <p>
 * 主要角色：
 * 抽象工厂（Abstract Factory）：提供了创建对象的接口，调用者通过它访问具体工厂的工厂方法 newProduct() 来创建产品(PersonFactory)（
 * 具体工厂（ConcreteFactory）：主要是实现抽象工厂中的抽象方法，完成具体对象的创建 (StudentFactory，TeacherFactory）
 * 抽象产品（Product）：定义了对象的规范，描述了对象的主要特性和功能。（Person）
 * 具体产品（ConcreteProduct）：实现了抽象产品角色所定义的接口，由具体工厂来创建，它同具体工厂之间一一对应(Student,Teacher)
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

