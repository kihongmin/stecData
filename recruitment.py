## .pkl 파일 X
#from regularizer import regularizer
#from es import ESCompany, ESRecruitment, ESTech


class Recruitment:
    def __init__(self, title, url, company, start_date, level, job, contents=None):
        self.techs = []
        if contents is None:
            self.contents = []
        elif isinstance(contents, list):
            self.contents = contents
        else:
            raise TypeError('class <Recruitment> can not get non-list type <contents> argurment.')
        self.title = title
        self.job = job
        self.url = url
        self.company = company
        self.start_date = start_date
        self.level = level


    def run(self):
        self.regularize()
        self.index2es()


    def get(self, all=False):
        if all:
            return {
                'title': self.title,
                'job':self.job,
                'url': self.url,
                'company': self.company,
                'start_date': self.start_date,
                'level': self.level,
                'techs': self.techs,
                'contents': self.contents,
            }
        return {
            'title': self.title,
            'job':self.job,
            'url': self.url,
            'company': self.company,
            'start_date': self.start_date,
            'level': self.level,
            'techs': self.techs,
        }


    def regularize(self):
        self.techs = []
        for content in self.contents:
            self.techs += regularizer.run(content)
        self.techs = list(set(self.techs))


    def index2es(self):
        items = self.get()
        ret = ESRecruitment.index(**items)
        ret = ESTech.index(techs=self.techs)
        ret = ESCompany.index(company=self.company)
