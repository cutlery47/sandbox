import random
import time
import csv

from algs.naive import naiveSearch
from algs.RK import RKSearch
from algs.BM import BMSearch
from algs.KMP import KMPSearch

def onlineAverage(avg: float, val: int, n: int) -> float:
    new_avg = avg + (val - avg) / n
    return new_avg

class Test:
    def __init__(self):
        self.lorem_input = 'inputs/lorem.txt'
        self.lorem_output = 'outputs/lorem.csv'
        self.dna_input = 'inputs/dna.txt'
        self.dna_output = 'outputs/dna.csv'
        self.alice_input = 'inputs/alice.txt'
        self.alice_output = 'outputs/alice.csv'

    def runTest(self, text, pat_len, text_len, runs, variable):
        print(variable)

        naive_avg = 0
        RK_avg = 0
        BM_avg = 0
        KMP_avg = 0
        python_avg = 0

        for run in range(runs):
            pat_start = random.randint(0, text_len - pat_len)
            pat_end = pat_start + pat_len
            pattern = text[pat_start: pat_end]

            # executing naive substring search
            naive_timestamp_start = time.time_ns()
            naive_res = naiveSearch(pattern, text)
            naive_timestamp_finish = time.time_ns()

            if naive_res == -1:
                print(f'Error: naive hasnt found the pattern: {pattern}')

            # executing RK substring search
            RK_timestamp_start = time.time_ns()
            RK_res = RKSearch(pattern, text)
            RK_timestamp_finish = time.time_ns()

            if RK_res == -1:
                print(f'Error: RK hasnt found the pattern: {pattern}')

            # executing BM substring search
            BM_timestamp_start = time.time_ns()
            BM_res = BMSearch(pattern, text)
            BM_timestamp_finish = time.time_ns()

            if BM_res == -1:
                print(f'Error: BM hasnt found the pattern: {pattern}')

            # executing KMP substring search
            KMP_timestamp_start = time.time_ns()
            KMP_res = KMPSearch(pattern, text)
            KMP_timestamp_finish = time.time_ns()

            if KMP_res == -1:
                print(f'Error: KMP hasnt found the pattern: {pattern}')

            # executing default python search
            pyhton_timestamp_start = time.time_ns()
            python_res = text.find(pattern)
            pyhton_timestamp_finish = time.time_ns()

            if python_res == -1:
                print(f'Error: python hasnt found the pattern: {pattern}')

            naive_avg = onlineAverage(naive_avg, naive_timestamp_finish - naive_timestamp_start, run + 1)
            RK_avg = onlineAverage(RK_avg, RK_timestamp_finish - RK_timestamp_start, run + 1)
            BM_avg = onlineAverage(BM_avg, BM_timestamp_finish - BM_timestamp_start, run + 1)
            KMP_avg = onlineAverage(KMP_avg, KMP_timestamp_finish - KMP_timestamp_start, run + 1)
            python_avg = onlineAverage(python_avg, pyhton_timestamp_finish - pyhton_timestamp_start, run + 1)

        return [naive_avg, RK_avg, BM_avg, KMP_avg, python_avg]

    def loremVarPatternTest(self):
        r = open(self.lorem_input, 'r')
        w = open(self.lorem_output, 'w')

        text = r.read()
        writer = csv.writer(w)

        # inserting a header in order to identify each algorithm's performance
        writer.writerow(['len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        for pat_len in range(1, len(text)):
            res = self.runTest(text, pat_len, len(text), 10000, pat_len)
            writer.writerow([pat_len, res[0], res[1], res[2], res[3], res[4]])

    def dnaVarPatternTest(self):
        r = open(self.dna_input, 'r')
        w = open(self.dna_output, 'w')

        text = r.read()
        writer = csv.writer(w)

        # inserting a header in order to identify each algorithm's performance
        writer.writerow(['len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        for pat_len in range(1, 1000):
            res = self.runTest(text, pat_len, len(text), 10000, pat_len)
            writer.writerow([pat_len, res[0], res[1], res[2], res[3], res[4]])

    def aliceVarPatternTest(self):
        r = open(self.alice_input, 'r')
        w = open(self.alice_output, 'w')

        text = r.read()
        writer = csv.writer(w)

        # inserting a header in order to identify each algorithm's performance
        writer.writerow(['len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        for pat_len in range(1, 1000):
            res = self.runTest(text, pat_len, len(text), 1000, pat_len)
            writer.writerow([pat_len, res[0], res[1], res[2], res[3], res[4]])

    # def dialoguesVarPatternTest(self):
    #     r = open(self.fallout_input, 'r')
    #     w = open(self.fallout_output, 'w')
    #
    #     text = r.read()
    #     writer = csv.writer(w)
    #
    #     # inserting a header in order to identify each algorithm's performance
    #     writer.writerow(['len', 'naive', 'RK', 'BM', 'KMP', 'python'])
    #
    #     for pat_len in range(1, 1000):
    #         res = self.runTest(text, pat_len, len(text), 1000, pat_len)
    #         writer.writerow([pat_len, res[0], res[1], res[2], res[3], res[4]])

    def loremVarTextTest(self):
        r = open('inputs/lorem.txt', 'r')
        w = open('outputs/lorem_10.csv', 'w')

        text = r.read()
        writer = csv.writer(w)

        writer.writerow(['pat_len', 'txt_len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        for text_len in range(10 + 1, 455 + 1):
            res = self.runTest(text, 10, text_len, 10000, text_len)
            writer.writerow([10, text_len, res[0], res[1], res[2], res[3], res[4]])

    def dnaVarTextTest(self):
        r = open('inputs/dna.txt', 'r')
        w = open('outputs/dna_2.csv', 'w')

        text = r.read()
        writer = csv.writer(w)

        writer.writerow(['pat_len', 'txt_len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        pat_lens = [10, 256, 1000]
        for pat_len in pat_lens:
            for text_len in range(pat_len + 1, 5000 + pat_len):
                res = self.runTest(text, pat_len, text_len, 1000, text_len)
                writer.writerow([pat_len, text_len, res[0], res[1], res[2], res[3], res[4]])

    def aliceVarTextTest(self):
        r = open('inputs/alice.txt', 'r')
        w = open('outputs/alice_2.csv', 'w')

        text = r.read()
        writer = csv.writer(w)

        writer.writerow(['pat_len', 'txt_len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        pat_lens = [10, 256, 1000]
        for pat_len in pat_lens:
            for text_len in range(pat_len + 1, 5000 + pat_len):
                res = self.runTest(text, pat_len, text_len, 1000, text_len)
                writer.writerow([pat_len, text_len, res[0], res[1], res[2], res[3], res[4]])

    def dialoguesVarTextTest(self):
        r = open('inputs/dialogues.txt', 'r')
        w = open('outputs/dialogues_2.csv', 'w')

        text = r.read()
        writer = csv.writer(w)

        writer.writerow(['pat_len', 'txt_len', 'naive', 'RK', 'BM', 'KMP', 'python'])

        pat_lens = [10, 256, 1000]
        for pat_len in pat_lens:
            for text_len in range(pat_len + 1, 5000 + pat_len):
                res = self.runTest(text, pat_len, text_len, 10000, text_len)
                writer.writerow([pat_len, text_len, res[0], res[1], res[2], res[3], res[4]])



