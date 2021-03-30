###Factory method pattern (nhóm khởi tạo)
Khởi tạo đối tượng một cách linh hoạt hơn, có một class cha (super class) với nhiều class con (sub-class). Ví dụ như: animal (cha) và class con là: hổ, hươu, vượn,...  
> Factory Method is a creational design pattern that Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

Được sử dụng khi:  
- Chúng ta có một super class với nhiều class con và dựa trên đầu vào, chúng ta cần trả về một class con. Mô hình này giúp chúng ta đưa trách nhiệm của việc khởi tạo một lớp từ phía người dùng (client) sang lớp Factory.
- Chúng ta không biết sau này sẽ cần đến những lớp con nào nữa. Khi cần mở rộng, hãy tạo ra sub class và implement thêm vào factory method cho việc khởi tạo sub class này.

Ví dụ:   
Super class:
```java
public interface Bank {
    String getBankName();
}
```
Sub classes:  
```java
package com.gpcoder.patterns.creational.factorymethod;
 
public class TPBank implements Bank {
 
    @Override
    public String getBankName() {
        return "TPBank";
    }
 
}
```

```java 
package com.gpcoder.patterns.creational.factorymethod;
 
public class VietcomBank implements Bank {
 
    @Override
    public String getBankName() {
        return "VietcomBank";
    }
 
}
```

Factory class:  
```java 
public class BankFactory {
 
    private BankFactory() {
    }
 
    public static final Bank getBank(BankType bankType) {
        switch (bankType) {
 
        case TPBANK:
            return new TPBank();
 
        case VIETCOMBANK:
            return new VietcomBank();
 
        default:
            throw new IllegalArgumentException("This bank type is unsupported");
        }
    }
 
}
```

Bank type:  
```java 
public enum BankType {
 
    VIETCOMBANK, TPBANK;
 
}
```

Client:  
```java 
public class Client {
 
    public static void main(String[] args) {
        Bank bank = BankFactory.getBank(BankType.TPBANK);
        System.out.println(bank.getBankName()); // TPBank
    }
}
```