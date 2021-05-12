import random
import sys

try:
    import tqdm
    progress = tqdm.tqdm
except ModuleNotFoundError:
    progress = lambda x: x

try:
    n = int(sys.argv[1])
except:
    n = 25

l1 = list(range(n))

random.shuffle(l1)

for l,r in enumerate(progress(l1)):
    print(f"{l},{r}")

