# 设计模式 - Go实践

1. 简单工厂、抽象工厂和工厂函数的区别？
   - 简单工厂 ： 使用一个工厂对象用来生产同一等级结构中的任意产品。（不支持拓展增加产品）
   - 工厂方法 ： 使用多个工厂对象用来生产同一等级结构中对应的固定产品。（支持拓展增加产品）
   - 抽象工厂 ： 使用多个工厂对象用来生产不同产品族的全部产品。（不支持拓展增加产品；支持增加产品族）