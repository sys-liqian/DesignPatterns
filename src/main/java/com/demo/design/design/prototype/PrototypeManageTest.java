package com.demo.design.design.prototype;

import java.util.HashMap;

/**
 * 原型模式扩展带管理器的原型模式
 */
public class PrototypeManageTest {

    public static void main(String[] args) {
        ProtoTypeManager pm = new ProtoTypeManager();
        Shape obj1 = pm.getShape("Circle");
        Shape obj2 = pm.getShape("Circle");
        System.out.println(obj1.hashCode() == obj2.hashCode());
        obj1.print();
        Shape obj3 = pm.getShape("Square");
        obj3.print();
    }
}

class ProtoTypeManager {
    private HashMap<String, Shape> ht = new HashMap<>();
    public ProtoTypeManager() {
        ht.put("Circle", new Circle());
        ht.put("Square", new Square());
    }
    public void addShape(String key, Shape obj) {
        ht.put(key, obj);
    }
    public Shape getShape(String key) {
        Shape temp = ht.get(key);
        return (Shape) temp.clone();
    }
}


interface Shape extends Cloneable {
    Object clone();

    void print();
}

class Circle implements Shape {

    @Override
    public Object clone() {
        Circle c = null;
        try {
            c = (Circle) super.clone();
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
        }
        return c;
    }

    @Override
    public void print() {
        System.out.println("这是一个圆形");
    }

}

class Square implements Shape {

    @Override
    public Object clone() {
        Square c = null;
        try {
            c = (Square) super.clone();
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
        }
        return c;
    }

    @Override
    public void print() {
        System.out.println("这是一个方形");
    }
}
