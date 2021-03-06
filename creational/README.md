## 创建型模式(Creational Pattern)

> 用于描述**“怎样创建对象”**，它的主要特点是**“将对象的创建与使用分离”**。这样可以降低系统的耦合度，使用者不需要关注对象的创建细节，对象的创建由相关的工厂来完成。

分类：

- 单例模式：某个类只能生成一个实例，该类提供了一个全局访问点供外部获取该实例，其拓展是有限多例模式。
- 原型模式：某一个对象作为原型，通过对其进行复制而克隆出多个和原型类似的新实例。
- 工厂方法模式：定义一个用于创建产品的接口，由子类决定生产什么产品。
- 抽象工厂模式：提供一个创建产品族的接口，其每个子类都可以生产一系列相关的产品。
- 建造者模式：将一个复杂对象分解成多个相对简单的部分，然后根据不同需要分别创建它们，最后构建成复杂对象。



工厂模式的定义：

> 定义一个创建产品对象的工厂接口，将产品对象的实际创建工作延迟到具体子工厂类中。这满足创建模式中所要求的“创建与使用分离”的特点。

三中不同的实现方式：

> - 简单工厂模式
> - 工厂方法模式
> - 抽象工厂模式
