int counter = 0;
void *worker() {
   for (int i=0;i<10;i++) {
      counter++;
   }
   return NULL;
}

int main(int argc, char *argv[]) {
   worker();
}