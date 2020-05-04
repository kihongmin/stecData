class NewbieCodes:
    newbie=10
    expert=90

    def convert(code):
        ret = {
            10:NewbieCodes.newbie,
            90:NewbieCodes.expert,
        }.get(code)
        if ret is None:
            return True, -1
        return False, ret