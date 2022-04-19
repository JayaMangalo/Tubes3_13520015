import random

dnalist = ["A","T","C","G"]
def makeDNA(size):
    dna = ""
    
    for i in range(size):
        dna += str(random.choice(dnalist))

    print(dna)

if __name__ == "__main__":
    makeDNA(30)