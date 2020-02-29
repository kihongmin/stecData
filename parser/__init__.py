import pickle

import re
import string
from konlpy.tag import Hannanum


class STECRegularizer:
    '''
    - stop_words.pkl 구조: list of word
    - syns_words.pkl 구조: 쌍방향 딕셔너리
    '''

    def __init__(self, stop_words=None, syns_words=None):
        self.puctuation = re.compile(f'[{string.punctuation}]')
        self.hannanum = Hannanum()
        if stop_words is None:
            self.stop_words = None
        else:
            with open(stop_words, 'rb') as f:
                self.stop_words = pickle.load(f)
        if syns_words is None:
            self.syns_words = None
        else:
            with open(syns_words, 'rb') as f:
                self.syns_words = pickle.load(f)

    def run(self, sent):
        sent = self.puctuation.sub(u' ', sent)
        sent = sent.lower()

        # pos
        ret = []
        words = self.hannanum.pos(sent)
        for word in words:
            if (word[1] == 'N') or (word[1] == 'F'):
                ret.append(word[0])

        # replace syns 

        # remove stop words
        ret = [r for r in ret if r not in self.stop_words]
        return ret