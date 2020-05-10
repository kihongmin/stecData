import os

from elasticsearch import Elasticsearch

from .. import config


cfg = config.get_config[os.getenv('STEC_ENV', 'local')]
es = Elasticsearch(cfg.ES_HOST)


class ESRecruitment:
    def index(title, company, job, techs, level, url, start_date):
        body = {
            'title':title,
            'company':company,
            'job':job,
            'techs':techs,
            'level':{
                'gte': level[0],
                'lte': level[-1],
            },
            'url':url,
            'start_date':start_date,
        }
        return es.index(index=cfg.ES_RECRUITMENT, doc_type=cfg.ES_DOCTYPE, body=body, timeout='10s')


class ESTech:
    def index(techs):
        body = {
            'completion': techs,
        }
        return es.index(index=cfg.ES_TECH, doc_type=cfg.ES_DOCTYPE, body=body, timeout='10s')


class ESCompany:
    def index(company):
        body = {
            'completion': company,
        }
        return es.index(index=cfg.ES_COMPANY, doc_type=cfg.ES_DOCTYPE, body=body, timeout='10s')