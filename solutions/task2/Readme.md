# Решение задачи №2

```java
public class Main {

     public static void main(String []args){
        // заданная строка.
        String source = "NEZE ETTPMGEXMSRW EVI XCTMGEPPC GSQTMPIH XS FCXIGSHI XLEX GER VYR SR ERC NEZE ZMVXYEP QEGLMRI VIKEVHPIWW SJ XLI YRHIVPCMRK GSQTYXIV EVGLMXIGXYVI";
        // смещение.
        int k = 4;
        // объявим переменную, в которую сохраним строку, полученную в результате преобразования.
        String result = "";
        // пройдем в цикле по исходной строке и сместим код каждого символа
        // на k влево, учитывая, что последние k символов алфавита
        // уже находятся вначале.
        for (int i = 0; i < source.length(); i++) {
            char symbol = (char)(source.codePointAt(i));
            // пробел не преобразовываем.
            if (symbol == ' ') {
                result += symbol;
                continue;
            }
            symbol -= k;
            // операция 'Z' - 'A' + 1 вернет нам кол-во букв алфавита.
            if (symbol < 'A') {
                symbol += 'Z' - 'A' + 1;
            }
            result += symbol;
        }
        System.out.println("Результат: " + result);
        // в условии задачи сказано, что нужно вывести 
        // 1, 8, 10, 11 и 24 символы (нумерация с 1).
        System.out.print(result.charAt(0));
        System.out.print(result.charAt(7));
        System.out.print(result.charAt(9));
        System.out.print(result.charAt(10));
        System.out.print(result.charAt(23));
     }
}
```
Output:
```
Результат: JAVA APPLICATIONS ARE TYPICALLY COMPILED TO BYTECODE THAT CAN RUN ON ANY JAVA VIRTUAL MACHINE REGARDLESS OF THE UNDERLYING COMPUTER ARCHITECTURE
JPICY
```

## Правильный ответ
__JPICY__