char findMissingLetter(char array[], int arrayLength)
{
  int i;
  char res;
  for (i = 0; i < arrayLength; i++)
  {
    if ((array[i] - 'a') - (array[i+1] - 'a') != -1)
    {
      res = array[i]+1;
      break;
    }
  }
  return res;
}