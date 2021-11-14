#include <stdbool.h>

bool is_prime(int num)
{
  if (num <= 1) return false;
  if (num == 2) return true;
  int i;
  for (i = 2; i * i <= num; i++) 
  {
    if (num%i==0) return false;
  }
  return true;
}