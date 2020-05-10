class LocalConfig:
    ES_HOST = 'https://search-stec-vsflhepbmhzueqdgrjqiijqmr4.ap-northeast-2.es.amazonaws.com'
    ES_RECRUITMENT = 'recruitments'
    ES_TECH = 'tech_completion'
    ES_COMPANY = 'company_completion'
    ES_DOCTYPE = 'doc'


class DevConfig:
    ES_HOST = 'https://search-stec-vsflhepbmhzueqdgrjqiijqmr4.ap-northeast-2.es.amazonaws.com'
    ES_RECRUITMENT = 'recruitments'
    ES_TECH = 'tech_completion'
    ES_COMPANY = 'company_completion'
    ES_DOCTYPE = 'doc'


get_config = {
    'local': LocalConfig,
    'dev': DevConfig,
}