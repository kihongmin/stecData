import os
from datetime import datetime

from . import STECRegularizer


DOC_PACK = 50
PEOPLE = 3

dic = os.path.abspath('../dataset/dictionary/')
regularizer = STECRegularizer(
                    stop_words=dic+'stop_words.pkl',
                    syns_words=dic+'syns_words.pkl',
                )

old = '../dataset/old/'
new = '../dataset/new/'
new = os.path.abspath(new)
docs = os.listdir(new)
ret = []
for d, doc in enumerate(docs):
    with open(new+doc, 'r') as f:
        lines = f.readlines()
    lines = [line.rstrip('\n\s\t\r\v') for line in lines]
    for line in lines:
        ret += regularizer.run(line)

    # raw text 삭제
    os.remove(new+doc)
    print(f'file [{doc}] in [{new}] is removed!')

    if d%DOC_PACK != 0:
        continue

    # 파싱된 단어 문서 생성
    now = datetime.now()
    now = [now.year, now.month, now.day, now.hour, now.minute, now.second]
    now = ''.join([str(n) for n in now])
    doc = old+d.zfill(3)+'_'+now
    with open(doc, 'w') as f:
        temp = '\n'.join(ret)
        f.write(temp)
    if d==DOC_PACK*PEOPLE:
        break
    ret = []
