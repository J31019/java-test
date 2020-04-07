# Решение задачи №4

```java
public class Main {

     public static void main(String []args){
        // заданная строка.
        String source = "RRRRRRRRRBBBBBBBBBBBBBBBBBBBBRRRRRRRRRRUURRRRRRRBBBBBBBBBBBRRURRRRRRRRRRRUUUUUUUUUUUUUUUUUUUUUUUUUURRRRRRRRRRLLLLLLLLLLLBBBBBBBBBBBBBBBBBBBBBLLLLLLLLLLLLLLLUUUULLLLLLUUUUUUUUUUBBBBUUUULLLLLLLLLLLLLLLLUUUUUUUUURRRRR";
        // объявляем матрицу размером 50x50.
        char m[][] = new char[50][50];
        // вносим символы в соответствии с условием задачи.
        m[1][19] = 'A';
        m[2][49] = 'C';
        m[9][9] = 'H';
        m[7][39] = 'I';
        m[0][1] = 'M';
        m[42][16] = 'O';
        m[13][17] = 'R';
        m[19][19] = 'S';
        m[30][15] = 'W';
        m[23][26] = 'Z';
        // заведем строковую переменную, в которую будем вносить те символы,
        // которые встретим "по пути".
        String s = "";
        // заведем 2 переменные, в которых будем хранить
        // текущую позицию игрока по оси x и y. 
        int cx = 0, cy = 0;
        // начинаем движение по шагам.
        for (int i = 0; i < source.length(); i++) {
            switch(source.charAt(i)) {
                case 'U':
                    cy--;
                    break;
                case 'B':
                    cy++;
                    break;
                case 'R':
                    cx++;
                    break;
                case 'L':
                    cx--;
                    break;
            }
            if (m[cy][cx] != 0) {
                // проверям, наступал ли игрок на эту клетку до этого.
                if (s.contains(String.valueOf(m[cy][cx]))) {
                    System.out.print(m[cy][cx]);
                }
                else {
                    s += m[cy][cx];
                }
            }
        }
     }
}
```
Output:
```
ZSRHM
```

## Правильный ответ
__ZSRHM__