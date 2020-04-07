# Решение задачи №3

```java
import java.util.HashMap;

public class Main {

     public static void main(String []args){
        // заданная строка.
        String source = "IMPEDFCJYEIPDTUOSWMYSMIAQYDEONTZTWRNGZFXKYBTLFZFJAUJWOYSECCYUJQJXHKNLEGKDMXLOEFVIFGULESWVMRXRUFZGRTUEUWELUFBLTAJAYSKTZNPEJZZRLKATPQBZHLFFHXLDKWOYTIXSGPCTHFOMOTCKSBOFIJBOBUGJYOIQLBAKKRWVHXSSXRZTQRMKAQSXODALBXZCKYZXATCXHBHKMTWGRRHEPQKNGXLIXHYTTAZSWVCZZUBEOXMUXAHEVBILLJRRVSGRIEUIULKBIPJHGXIKWYGYXZTGEFYDGYVKPRRSFJJRLDDQCUEDFFPEEHOTMWAOEMLXWXVIHIJLRCVIQXGZTXDBNXQGZSYNSIYQBPLNTXDKIJVTJOGKKKAYTVUHMVVTWKOCHETOAZBEPNKQQKCBGLJSOQHRRTQILQKFYHQLSACFNMZTCGLGAZIRINEJRLMHPRDGOTNYIOJDNDFYCIGIACALITODHZHJJXJNZWEUDMFWYPNCDYFOKGNTWPONTQLMMXFDNWZVYARYUEBJIPHGCRMFWTMPTXVERDRGXPMOVLWEYJTXWEAPJXKHKCMSWWIPMQFWSVKIWLUVAKKUUQWAOBDDLRGOWHGQDHSCNUPRJOPRMEGLISBUKQWKTSJCNBVDUQZIHUHPYOAZMKSHKIIWMXSZVLCPCJDRVMSJIDXNLNNVLRMNNKXDYSOKKXMWNRXVVOGKYNTHFAUNKSDZUPKUXPUVKNVQCZJRXSIZQYWHDLOXLVLVENIBQPCMDFCGTGZMQMORQLRAKVMXMNLIGOSYVAIAORAIDUPFHHTRFCPPBYQWQCKJYQMKAETCGHBCEJSSFVXVSIPGFMEOPXAXPYKCWZZGMFWIERBIODKFEKLXOXVQDFQOOORODUSNRCVUWLHOMEETODKDUCKVIQUXYYUFPZALYMKPYNBYTWFHVKGAIEZWSXRUOEONGNEMLDWMNJGEVODASJIJKOJWGLKJKEATXBNODCK";
        // объявим коллекцию типа Map, в качестве ключа будем использовать код символа строки,
        // а в качестве значения - кол-во повторений этого символа.
        HashMap<Character, Integer> map = new HashMap<>();
        for (int i = 0; i < source.length(); i++) {
            char symbol = source.charAt(i);
            // проверяем, содержится ли символ в коллекции.
            if (map.containsKey(symbol)) {
                map.put(symbol, map.get(symbol) + 1);
                continue;
            }
            map.put(symbol, 1);
        }
        // необязательный шаг.
        // выведем значения коллекции в отсортированном виде.
        for (int i = source.length(); i > 0; i--) {
            for (Character key : map.keySet()) {
                if (map.get(key) == i) {
                    System.out.print(key + " = ");
                    System.out.println(i);
                }
            }
        }
     }
}
```
Output:
```
K = 55
O = 48
X = 47
I = 44
L = 43
E = 41
R = 41
T = 41
M = 40
D = 39
Y = 39
G = 38
W = 38
J = 37
N = 37
V = 37
F = 35
S = 35
U = 35
A = 34
C = 34
H = 34
P = 34
Q = 34
Z = 34
B = 26
```

## Правильный ответ
__KOXIL__