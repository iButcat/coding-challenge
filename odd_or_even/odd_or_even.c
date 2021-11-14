#include <stddef.h>

const char *odd_or_even(const int *v, size_t sz) {
  int i;
  int sum;
  for (i = 0; i < sz; i++) 
  {
    sum += v[i];
  }
  if (sum % 2 == 0) return "even";
  else return "odd";
  
  return NULL;
}