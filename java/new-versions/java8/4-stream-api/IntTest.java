import java.util.*;
import java.util.stream.*;

public class IntTest {
    public static void main(String[] args) {

        IntStream.range(0,10)
            .forEach(System.out::println);

        System.out.println("************with for***************");
        for(int i =0; i < 10; i++){
            System.out.println(i);
        }
        /*
         input  : range()
            |9|
            |8|
            |7|
            |6|
          output :forEach()
         */

        System.out.println("********************************");

        long tedadZoj = IntStream.rangeClosed(1,10)  // outpt: intStream
            .filter(a -> a%2 == 0) // output: intStream    /*  1 2 3 4 5 6 7 8 9 10  -> 2 4 6 8 10 */
            .count();
        System.out.println(tedadZoj);

        /*
         * input : rangeClosed()
            |6|
            |7|
            |8|
            |9|
            |10|
            filter(zoj)
            |6|
            |8|
            |10|
         output :count()
        */

        System.out.println("********************************");

        IntStream.rangeClosed(1,20)
            .filter( a ->
                 IntStream.rangeClosed(1,a)
                 .filter(b -> a%b == 0) // shomarande (maghsum alaih)
                 .count() == 2
                )
            .forEach(a -> System.out.println(a));


        System.out.println("********************************");

        int[] zoj1to10 = IntStream.rangeClosed(1,5)
            .map(i -> 2*i)
            .toArray();

        System.out.println("arr[0] : " + zoj1to10[0]);
        System.out.println(Arrays.toString(zoj1to10));

        System.out.println("********************************");

        Scanner sc = new Scanner(System.in);
        String line = sc.nextLine();
        String[] input_arr =  line.trim().split(" "); // "s","aa","l"
        int[] arr = Arrays.stream(input_arr)  // Stream<String>
            .mapToInt(Integer::parseInt) // IntStream
            .toArray();

        Arrays.stream(arr).forEach( (i) -> System.out.println(i));


        System.out.println("********************************");
        Random random = new Random();
        int randomMin = IntStream.rangeClosed(1,10)
            .map(i -> random.nextInt())
            .reduce((a,b) -> a<b ? a : b) // return optional
            .getAsInt();
        System.out.println("some random integer : "+randomMin);

        /*
          (12 15) 0 -5 42
          (12 0) -5 42
          (0 -5) 42
          (-5 42)
          -5

          ( )
        */

    }
}

// boxed
// allMatch
// noneMatch
// collect
