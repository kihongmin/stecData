import os
import re
import pickle

from konlpy.tag import Hannanum


class STECRegularizer:
    def __init__(self, tech_words, double_words, triple_words, syns_words):
        self.puctuation = re.compile('[!"$%&\'()*,-/:;<=>?@[\\]^_`{|}~]')
        self.hannanum = Hannanum()
        with open(tech_words, 'rb') as f:
            self.tech_words = pickle.load(f)
        with open(double_words, 'rb') as f:
            self.double_words = pickle.load(f)
        with open(triple_words, 'rb') as f:
            self.triple_words = pickle.load(f)
        with open(syns_words, 'rb') as f:
            self.syns_words = pickle.load(f)

    def run(self, sent):
        sent = sent.strip()
        sent = sent.replace("\xa0", " ")
        sent = re.sub('[\s]+',' ',sent)
        sent = self.puctuation.sub(u' ', sent)
        sent = sent.lower()

        # pos
        pos = []
        words = self.hannanum.pos(sent)
        for word in words:
            if (word[1] == 'N') or (word[1] == 'F'):
                pos.append(word[0])

        ret = []
        # for num, dic in {3:self.triple_words, 2:self.double_words}.items():
        #     for i, r in enumerate(pos):
        #         word = ' '.join(pos[i:i+num])
        #         tmp = self.syns_words.get(word)
        #         if tmp is not None:
        #             word = tmp
        #         if word in dic:
        #             for j in range(num):
        #                 pos[i+j] = ""
        #             ret.append(word)
        #     pos = [r for r in pos if r != ""]

        # triple
        for i, r in enumerate(pos):
            word = ' '.join(pos[i:i+3])
            tmp = self.syns_words.get(word)
            if tmp is not None:
                word = tmp
            if word in self.triple_words:
                for j in range(3):
                    pos[i+j] = ""
                ret.append(word)
        pos = [r for r in pos if r != ""]

        # double
        for i, r in enumerate(pos):
            word = ' '.join(pos[i:i+2])
            tmp = self.syns_words.get(word)
            if tmp is not None:
                word = tmp
            if word in self.double_words:
                for j in range(2):
                    pos[i+j] = ""
                ret.append(word)
        pos = [r for r in pos if r != ""]

        # normal
        for i, word in enumerate(pos):
            tmp = self.syns_words.get(word)
            if tmp is not None:
                word = tmp
            if word in self.tech_words:
                ret.append(word)

        return ret


regularizer = STECRegularizer(
                    tech_words=os.path.abspath('./main/regularizer/tech_words.pkl'),
                    double_words=os.path.abspath('./main/regularizer/double_words.pkl'),
                    triple_words=os.path.abspath('./main/regularizer/triple_words.pkl'),
                    syns_words=os.path.abspath('./main/regularizer/syns_words.pkl'),)