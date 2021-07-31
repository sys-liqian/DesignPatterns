package com.demo.design.design.prototype;

import java.io.*;
import java.util.List;

/**
 * 原型模式：用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象。
 * <p>
 * 原型模式包含以下主要角色。
 * 抽象原型类：规定了具体原型对象必须实现的接口。
 * 具体原型类：实现抽象原型类的 clone() 方法，它是可被复制的对象。
 * 访问类：使用具体原型类中的 clone() 方法来复制新的对象。
 */
public class PrototypeTest {

    public static void main(String[] args) {
        try {
            protoTypeASimpleClone();
            protoTypeADeepClone();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * 简单原型模式深拷贝测试方法
     */
    public static void protoTypeADeepClone() throws IOException, ClassNotFoundException {
        ProtoTypeACon con = new ProtoTypeACon("a");
        ProtoTypeA protoTypeA = new ProtoTypeA(con);
        ProtoTypeA protoTpeAClone = protoTypeA.deepClone();
        System.out.println("protoTypeA和protoTpeAClone是否相等：" + (protoTypeA.hashCode() == protoTpeAClone.hashCode()));
        System.out.println("protoTypeA和protoTpeAClone中con对象是否相等：" + (protoTypeA.getCon().hashCode() == protoTpeAClone.getCon().hashCode()));
    }

    /**
     * 简单原型模式浅拷贝测试方法
     */
    public static void protoTypeASimpleClone() throws CloneNotSupportedException {
        ProtoTypeACon con = new ProtoTypeACon("a");
        ProtoTypeA protoTypeA = new ProtoTypeA(con);
        ProtoTypeA protoTpeAClone = protoTypeA.clone();
        /**
         * 使用java自带的clone为浅拷贝
         * 浅拷贝，非基本类型属性，仍指向原有属性所指向的对象的内存地址
         * protoTypeA和protoTpeAClone都包含con对象,他俩包含的con是同一个
         */
        con.setName("b");
        System.out.println("protoTypeA和protoTpeAClone是否相等：" + (protoTypeA.hashCode() == protoTpeAClone.hashCode()));
        System.out.println("protoTypeA和protoTpeAClone中con对象是否相等：" + (protoTypeA.getCon().hashCode() == protoTpeAClone.getCon().hashCode()));
    }
}




class ProtoTypeA implements Cloneable, Serializable {

    private ProtoTypeACon con;

    private List<String> list;

    public ProtoTypeA() {
    }

    public ProtoTypeA(ProtoTypeACon con) {
        this.con = con;
    }

    public ProtoTypeACon getCon() {
        return con;
    }

    public void setCon(ProtoTypeACon con) {
        this.con = con;
    }

    public List<String> getList() {
        return list;
    }

    public void setList(List<String> list) {
        this.list = list;
    }

    /**
     * 浅拷贝
     */
    @Override
    protected ProtoTypeA clone() throws CloneNotSupportedException {
        return (ProtoTypeA) super.clone();
    }

    /**
     * 深拷贝
     * 利用序列化实现
     */
    protected ProtoTypeA deepClone() throws IOException, ClassNotFoundException {
        //当前对象写入流
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        ObjectOutputStream oos = new ObjectOutputStream(bos);
        oos.writeObject(this);
        //流中读取对象，原对象任存在于jvm
        ByteArrayInputStream bis = new ByteArrayInputStream(bos.toByteArray());
        ObjectInputStream ois = new ObjectInputStream(bis);
        ProtoTypeA obj = (ProtoTypeA) ois.readObject();
        return obj;
    }


}

class ProtoTypeACon implements Serializable {
    private String name;

    public ProtoTypeACon() {
    }

    public ProtoTypeACon(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
