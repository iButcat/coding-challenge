class Solution {
    public static Boolean checkIfLower(String input) {
      for (Character ele : input.toCharArray()) {
        if (Character.isUpperCase(ele)) {
          return false;
        }
      }
      return true;
    }
    
    public static String camelCase(String input) {
      if (checkIfLower(input)) {
        return input;
      }
      
      String builder = "";
      char[] charArray = input.toCharArray();
      for (int i = 0; i < charArray.length; i++) {
        if (Character.isUpperCase(charArray[i])) {
          builder += " ";
          builder += charArray[i];
        } else {
          builder += charArray[i];
        }
      }
      return builder;
    }
  }