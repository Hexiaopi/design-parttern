### 原型模式(Prototype Pattern)

> 用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象。

**角色**

- 抽象原型：规定了具体原型对象必须实现的接口
- 具体原型：实现抽象原型的clone()方法，是可被复制的对象
- 访问：使用具体原型类中的clone()方法来复制新的对象