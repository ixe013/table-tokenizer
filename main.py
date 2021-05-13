import random
import sys
import time

import tqdm

    
try:
    n = int(sys.argv[1])
except:
    n = 25

def wrap_progress_with_random(progress_func, rand_func):
    def inner_random():
        progress_func.update()
        return rand_func()

    return inner_random

print("Allocating memory...", file=sys.stderr)

table = list(range(n))

progress = tqdm.tqdm(total=len(table))

progress_random = wrap_progress_with_random(progress, random.random)

print("Suffling the table...", file=sys.stderr)
random.shuffle(table, progress_random)

progress.update()
progress.close()

print("Printing the table...", file=sys.stderr)
for r in tqdm.tqdm(table):
    print(r)

