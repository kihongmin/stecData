import re
from datetime import datetime
from datetime import timedelta


class StartDate:
    rocketpunch_compiler = re.compile('등록')

    def transform(date, source=None):
        if source == 'rocketpunch':
            if not StartDate.rocketpunch_compiler.search(date):
                return None

        date = re.sub('[^0-9]', '', date)
        len_date = len(date)
        if len_date == 4:
            date = '2020'+date
        elif len_date == 6:
            date = '20'+date
        elif len_date == 12:
            date = '20'+date[:6]
        elif len_date == 16:
            date = date[:8]
        else:
            return None

        now = datetime.now()
        now -= timedelta(days=1)
        if now.strftime('%Y%m%d') != date:
            return None
        return date