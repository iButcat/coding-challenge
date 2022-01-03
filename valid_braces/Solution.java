import java.util.Stack;

public class BraceChecker {

  public boolean isValid(String braces) {
    Stack<Character> stack = new Stack<>();
    for (int i = 0; i < braces.length(); i++) {
      char charBrace = braces.charAt(i);
      if (charBrace == '(' || charBrace == '[' || charBrace =='{') {
        stack.push(charBrace);
        continue;
      }  
      
      if (charBrace == ')' || charBrace ==']' || charBrace =='}') {
        if (stack.isEmpty()) {
          return false;
        }
        char last = stack.peek();
        if (charBrace == ')' && last == '(' || charBrace == ']' && last == '[' || charBrace == '}' && last == '{' ) {
           stack.pop();
        } else {
          return false;
        }
      }
    }
     return stack.isEmpty();
  }
}