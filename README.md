# 환경설정

## Dependency

```shell
pip3 install selenium bs4
```

# Data Schema

## ES

### Recruitment
```json
{
  "settings":{
    "analysis":{
      "analyzer":{
        "pos_analyzer":{
          "type":"custom",
          "tokenizer":"hanguel_tokenizer",
            "filter":[
                "lowercase",
                "trim",
            ],
        },
        "completion_analyzer":{
          "type":"custom",
          "char_filter":["jaso_char"],
          "tokenizer":"icu_tokenizer"
        },
      },
      "char_filter":{
        "jaso_char":{
          "type":"icu_normalizer",
          "name":"nfkc_cf",
          "mode":"decompose"
        },
      },
      "tokenizer":{
        "hanguel_tokenizer":{
          "type":"seunjeon_tokenizer",
          "deniflect":"true",
          "decompound":"false",
          "index_eojeol":"false",
          "index_poses":["N", "V", "M", "UNK"],
          "pos_tagging":"false",
          "max_unk_length":8,
        },
      },
    },
  },
  "mappings":{
    "doc":{
      "dynamic":"true",
      "properties": {
        "title":{
          "type":"text",
          "analyzer":"pos_analyzer",
          "copy_to":["title_completion"],
        },
        "title_completion":{
          "type":"completion",
          "analyzer":"completion_analyzer",
        },
        "company":{
          "type":"keyword",
        },
        "job":{
            "type":"keyword",
        },
        "tech":{
            "type":"keyword",
        },
        "level":{
            "type":"integer_range",
        },
        "url":{
          "type":"keyword",
          "index":"false",
          "norms":"false",
        },
        "start_date":{
          "type":"date",
          "format":"yyyyMMdd",
        },
      },
    },
  },
}
```

### completion: techs, company
```json
{
  "settings":{
    "analysis":{
      "analyzer":{
        "completion_analyzer":{
          "type":"custom",
          "char_filter":["jaso_char"],
          "tokenizer":"icu_tokenizer"
        }
      },
      "char_filter":{
        "jaso_char":{
          "type":"icu_normalizer",
          "name":"nfkc_cf",
          "mode":"decompose"
        }
      },
    }
  },
  "mappings":{
    "doc":{
      "dynamic":"true",
      "properties": {
        "tech":{
          "type":"completion",
          "analyzer":"completion_analyzer"
        },
      }
    }
  }
}
```


# 데이터 소스

## Data Description
{
1. url:
    -ex) https://gitlab.com/geekermeter/data/-/edit/master/README.md
2. title
    -ex) [FIFA ONLINE 4] Feature Game Client Engineer
3. origin
    -ex) EA Korea
4. start_date
    -ex) 3.11 (need to be same, but not yet)
5. newbie
    -ex) 신입, 경력
6. content

}

## 크롤링 대상 홈페이지

	- coupang
	- kakao
	- naver
	- ncsoft
	- netmarble
	- nexon
	- programmers
	- rocketpunch
-----------------------------

## 진행상황

|사이트|URL|Title|origin|start_date|newbie|content|auto|
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
|coupang|O|X|X|X|X|X|X|
|kakao|O|X|X|X|X|X|X|
|naver|O|X|X|X|X|X|X|
|ncsoft|O|X|X|X|X|X|X|
|netmarble|O|X|X|X|X|X|X|
|nexon|O|O|O|O|O|O|X|X|
|programmers|O|O|O|O|O|O|O|
|rocketpunch|O|O|O|O|O|O|O|
