package com.demo.design.design.abstractFactory;

/**
 * 抽象工厂模式
 * 抽象工厂模式同工厂方法模式一样，也是由抽象工厂、具体工厂、抽象产品和具体产品等 4 个要素构成，但抽象工厂中方法个数不同，抽象产品的个数也不同。现在我们来分析其基本结构和实现方法。
 */
public class AbstractFactoryTest {

    public static void main(String[] args) {

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



class PersonFactory implements BiologyFactory {

    @Override
    public Person getPerson() {
        return null;
    }

    @Override
    public Animal getAnimal() {
        return null;
    }
}



interface Animal {
    void talk();
}

class Cat implements Animal {
    @Override
    public void talk() {
        System.out.println("猫叫");
    }
}

class Dog implements Animal {
    @Override
    public void talk() {
        System.out.println("狗叫");
    }
}

/**
 * 抽象工厂
 */
interface BiologyFactory {
    Person getPerson();

    Animal getAnimal();
}


