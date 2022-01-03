import java.util.HashMap;

class Solution {
  static int stray(int[] numbers) {
    HashMap<Integer, Integer> numCounter = new HashMap<>();
    
    for (int i = 0; i < numbers.length; i++) {
      Integer counter = numCounter.get(numbers[i]);
      if (counter == null) {
        numCounter.put(numbers[i], 1);
      } else {
        numCounter.put(numbers[i], counter + 1);
      }
    }
    
    for (int key : numCounter.keySet()) {
      if (numCounter.get(key) == 1) {
        return key;
      } else {
        continue;
      }
    }
    return 0;
  }
}