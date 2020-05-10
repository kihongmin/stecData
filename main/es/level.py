import re


class Level:
    intern=0
    newbie=10
    unlimited=1000
    # n년차: n0
    # 기본적으로 range로 표현

    # mapper
    string_mapper = {
        '인턴':0,
        '신입':10,
        '경력':1000,
    }

    # programmers
    programmers_no_experience = re.compile('경력 무관')
    programmers_experience = re.compile('경력')

    def programmers2code(text):
        if Level.programmers_no_experience.search(text):
            return (Level.newbie, Level.unlimited)
        if Level.programmers_experience.search(text):
            return (Level.newbie+10, Level.unlimited)
        return []

    def string2code(text):
        return Level.string_mapper.get(text)